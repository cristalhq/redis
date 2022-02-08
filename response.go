package redis

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/gobwas/pool"
)

type Value = string

func responseDecodeInt(r *bufio.Reader) (int64, error) {
	line, err := responseReadLine(r)
	if err != nil {
		return 0, err
	}
	if len(line) > 3 && line[0] == ':' {
		return parseInt(line[1 : len(line)-2]), nil
	}
	return 0, responseReadError(r, line, "integer")
}

func responseDecodeInts(r *bufio.Reader) ([]int64, error) {
	l, err := responseReadArrayLen(r)
	if err != nil {
		return nil, err
	}

	array := make([]int64, 0, l)
	for len(array) < cap(array) {
		s, err := responseDecodeInt(r)
		switch err {
		case nil:
			array = append(array, s)
		case errNull:
			array = append(array, 0)
		default:
			return nil, err
		}
	}
	return array, nil
}

func responseDecodeFloat(r *bufio.Reader) (float64, error) {
	line, err := responseReadLine(r)
	if err != nil {
		return 0, err
	}
	if len(line) > 3 && line[0] == ':' {
		return strconv.ParseFloat(string(line[1:len(line)-2]), 64)
	}
	return 0, responseReadError(r, line, "integer")
}

func responseReadError(r *bufio.Reader, line []byte, want string) error {
	switch {
	case len(line) > 3 && line[0] == '-':
		return errors.New(string(line[1 : len(line)-2]))
	case len(line) > 3 && line[0] == '!':
		l := parseInt(line[1 : len(line)-2])
		if l < 0 || l > sizeMax {
			return nil
		}
		s, err := responseReadStringSize(r, int(l))
		if err != nil {
			return fmt.Errorf("%w; blob error unavailable", err)
		}
		return errors.New(s)
	case len(line) == 5 && line[0] == '+' && line[1] == 'O' && line[2] == 'K':
		return errOK
	default:
		return fmt.Errorf("%w; %s expected-received %.40q", errProtocol, want, line)
	}
}

func responseDecodeString(r *bufio.Reader) (string, error) {
	l, err := responseReadBlobLen(r)
	if err != nil {
		return "", err
	}
	return responseReadStringSize(r, l)
}

func responseReadBlobLen(r *bufio.Reader) (int, error) {
	line, err := responseReadLine(r)
	if err != nil {
		return 0, err
	}

	if len(line) > 3 && line[0] == '$' {
		l := parseInt(line[1 : len(line)-2])
		switch {
		case l >= 0 && l <= sizeMax:
			return int(l), nil
		case l == -1:
			return 0, errNull
		}
	}
	return 0, responseReadError(r, line, "blob")
}

func responseDecodeStrings(r *bufio.Reader) ([]string, error) {
	l, err := responseReadArrayLen(r)
	if err != nil {
		return nil, err
	}
	array := make([]string, 0, l)

	for len(array) < cap(array) {
		s, err := responseDecodeString(r)
		switch err {
		case nil:
			array = append(array, s)
		case errNull:
			array = append(array, "")
		default:
			return nil, err
		}
	}
	return array, nil
}

func responseReadArrayLen(r *bufio.Reader) (int64, error) {
	line, err := responseReadLine(r)
	if err != nil {
		return 0, err
	}

	if len(line) > 3 && line[0] == '*' {
		l := parseInt(line[1 : len(line)-2])
		switch {
		case l >= 0 && l <= elementMax:
			return l, nil
		case l == -1:
			return 0, errNull
		}
	}
	return 0, responseReadError(r, line, "array")
}

func responseReadLine(r *bufio.Reader) (line []byte, err error) {
	line, err = r.ReadSlice('\n')
	if err != nil {
		if err == bufio.ErrBufferFull {
			err = fmt.Errorf("%w; LF exceeds %d bytes: %.40q…", errProtocol, r.Size(), line)
		}
		return nil, err
	}
	return line, nil
}

func responseReadStringSize(r *bufio.Reader, size int) (string, error) {
	slice, err := r.Peek(size)
	switch err {
	case nil:
		s := string(slice)
		_, err = r.Discard(size + 2)
		return s, err
	case bufio.ErrBufferFull:
		// pass
	default:
		return "", err
	}

	var blob strings.Builder
	blob.Grow(size)
	blob.Write(slice)
	for {
		size -= len(slice)
		slice, err = r.Peek(size)
		blob.Write(slice)
		switch err {
		case nil:
			_, err = r.Discard(size + 2) // skip CRLF
			return blob.String(), err
		case bufio.ErrBufferFull:
			_, _ = r.Discard(len(slice)) // guaranteed to succeed
		default:
			return "", err
		}
	}
}

func parseInt(b []byte) int64 {
	if len(b) == 0 {
		return 0
	}
	u := uint64(b[0])

	neg := false
	if u == '-' {
		neg = true
		u = 0
	} else {
		u -= '0'
	}

	for i := 1; i < len(b); i++ {
		u = u*10 + uint64(b[i]-'0')
	}

	value := int64(u)
	if neg {
		value = -value
	}
	return value
}

var responsePool = NewReaderPool(1256, 65536)

// getResponse returns bufio.Reader whose buffer has at least size bytes. It returns
// its capacity for further pass to Put().
// Note that size could be ceiled to the next power of two.
// getResponse is a wrapper around responsePool.Get().
func getResponse(w io.Reader, size int) *bufio.Reader { return responsePool.Get(w, size) }

// ReaderPool contains logic of *bufio.Reader reuse with various size.
type ReaderPool struct {
	pool *pool.Pool
}

// NewReaderPool creates new ReaderPool that reuses writers which size is in
// logarithmic range [min, max].
func NewReaderPool(min, max int) *ReaderPool {
	return &ReaderPool{pool.New(min, max)}
}

// Get returns bufio.Reader whose buffer has at least size bytes.
func (rp *ReaderPool) Get(r io.Reader, size int) *bufio.Reader {
	v, n := rp.pool.Get(size)
	if v != nil {
		br := v.(*bufio.Reader)
		br.Reset(r)
		return br
	}
	return bufio.NewReaderSize(r, n)
}

// Put takes ownership of bufio.Reader for further reuse.
func (rp *ReaderPool) Put(br *bufio.Reader) {
	br.Reset(nil)
	rp.pool.Put(br, br.Size())
}

const defaultClientReadBufferSize = 4096

var (
	errProtocol = errors.New("redis: protocol violation")
	errNull     = errors.New("redis: null")
	errOK       = errors.New("redis: OK")
)

const (
	sizeMax    = 512 << 20
	elementMax = 1<<32 - 1
)
