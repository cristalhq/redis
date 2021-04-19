package redis

import (
	"bufio"
	"strconv"
)

type request struct {
	buf     []byte
	receive chan *bufio.Reader
}

func newEmptyRequest() *request {
	req := &request{
		buf: make([]byte, 256),
	}
	return req
}

func newRequest(prefix string) *request {
	req := newEmptyRequest()
	req.buf = append(req.buf[:0], prefix...)
	return req
}

func (req *request) addBytes(b []byte) {
	req.appendBytes(b)
	req.buf = append(req.buf, '\r', '\n')
}

func (req *request) addString(s string) {
	req.appendString(s)
	req.buf = append(req.buf, '\r', '\n')
}

func (req *request) addInt(v int64) {
	req.appendInt(v)
	req.buf = append(req.buf, '\r', '\n')
}

func (req *request) addBytesSlice(a [][]byte) {
	for _, b := range a {
		req.buf = append(req.buf, '\r', '\n', '$')
		req.appendBytes(b)
	}
	req.buf = append(req.buf, '\r', '\n')
}

func (req *request) addStringSlice(a []string) {
	for _, s := range a {
		req.buf = append(req.buf, '\r', '\n', '$')
		req.appendString(s)
	}
	req.buf = append(req.buf, '\r', '\n')
}

func (req *request) appendBytes(v []byte) {
	req.buf = strconv.AppendUint(req.buf, uint64(len(v)), 10)
	req.buf = append(req.buf, '\r', '\n')
	req.buf = append(req.buf, v...)
}

func (req *request) appendString(v string) {
	req.buf = strconv.AppendUint(req.buf, uint64(len(v)), 10)
	req.buf = append(req.buf, '\r', '\n')
	req.buf = append(req.buf, v...)
}

func (req *request) appendInt(v int64) {
	sizeOffset := len(req.buf)
	sizeSingleDigit := v > -1e8 && v < 1e9
	if sizeSingleDigit {
		req.buf = append(req.buf, 0, '\r', '\n')
	} else {
		req.buf = append(req.buf, 0, 0, '\r', '\n')
	}

	valueOffset := len(req.buf)
	req.buf = strconv.AppendInt(req.buf, v, 10)
	size := len(req.buf) - valueOffset
	if sizeSingleDigit {
		req.buf[sizeOffset] = byte(size + '0')
	} else { // two digits
		req.buf[sizeOffset] = byte(size/10 + '0')
		req.buf[sizeOffset+1] = byte(size%10 + '0')
	}
}
