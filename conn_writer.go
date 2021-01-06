package redis

import (
	"bufio"
	"encoding"
	"fmt"
	"io"
	"strconv"
	"time"
)

type connWriter struct {
	wr     *bufio.Writer
	lenBuf []byte
	numBuf []byte
}

func newWriter(wr io.Writer) *connWriter {
	return &connWriter{
		wr:     bufio.NewWriter(wr),
		lenBuf: make([]byte, 64),
		numBuf: make([]byte, 64),
	}
}

func (w *connWriter) Reset(wr io.Writer) {
	w.wr.Reset(wr)
}

func (w *connWriter) Flush() error {
	return w.wr.Flush()
}

func (w *connWriter) WriteArgs(args []interface{}) error {
	err := w.wr.WriteByte('*')
	if err != nil {
		return err
	}

	err = w.writeLen(len(args))
	if err != nil {
		return err
	}

	for _, arg := range args {
		err := w.writeArg(arg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *connWriter) writeLen(n int) error {
	w.lenBuf = strconv.AppendUint(w.lenBuf[:0], uint64(n), 10)
	w.lenBuf = append(w.lenBuf, '\r', '\n')
	_, err := w.wr.Write(w.lenBuf)
	return err
}

func (w *connWriter) writeArg(v interface{}) error {
	switch v := v.(type) {
	case nil:
		return w.writeString("")

	case encoding.BinaryMarshaler:
		b, err := v.MarshalBinary()
		if err != nil {
			return err
		}
		return w.writeBytes(b)

	case string:
		return w.writeString(v)
	case []byte:
		return w.writeBytes(v)

	case int:
		return w.writeInt(int64(v))
	case int8:
		return w.writeInt(int64(v))
	case int16:
		return w.writeInt(int64(v))
	case int32:
		return w.writeInt(int64(v))
	case int64:
		return w.writeInt(v)

	case uint:
		return w.writeUint(uint64(v))
	case uint8:
		return w.writeUint(uint64(v))
	case uint16:
		return w.writeUint(uint64(v))
	case uint32:
		return w.writeUint(uint64(v))
	case uint64:
		return w.writeUint(v)

	case float32:
		return w.writeFloat(float64(v))
	case float64:
		return w.writeFloat(v)

	case bool:
		if v {
			return w.writeInt(1)
		}
		return w.writeInt(0)

	case time.Time:
		return w.writeInt(v.Unix())
	case time.Duration:
		return w.writeInt(int64(v))

	default:
		return fmt.Errorf("redis: can't marshal %T", v)
	}
}

func (w *connWriter) writeBytes(b []byte) error {
	err := w.wr.WriteByte('$')
	if err != nil {
		return err
	}

	err = w.writeLen(len(b))
	if err != nil {
		return err
	}

	_, err = w.wr.Write(b)
	if err != nil {
		return err
	}
	return w.crlf()
}

func (w *connWriter) writeString(s string) error {
	return w.writeBytes(s2b(s))
}

func (w *connWriter) writeInt(n int64) error {
	w.numBuf = strconv.AppendInt(w.numBuf[:0], n, 10)
	return w.writeBytes(w.numBuf)
}

func (w *connWriter) writeUint(n uint64) error {
	w.numBuf = strconv.AppendUint(w.numBuf[:0], n, 10)
	return w.writeBytes(w.numBuf)
}

func (w *connWriter) writeFloat(f float64) error {
	w.numBuf = strconv.AppendFloat(w.numBuf[:0], f, 'f', -1, 64)
	return w.writeBytes(w.numBuf)
}

func (w *connWriter) crlf() error {
	err := w.wr.WriteByte('\r')
	if err != nil {
		return err
	}
	return w.wr.WriteByte('\n')
}
