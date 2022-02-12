package redis

import (
	"strconv"
	"sync"
)

type request struct {
	buf []byte
}

// TODO(oleg): make adaptive
var requestPool = sync.Pool{
	New: func() interface{} {
		return &request{buf: make([]byte, 0, 256)}
	},
}

func newRequest(prefix string) *request {
	req := requestPool.Get().(*request)
	req.buf = append(req.buf[:0], prefix...)
	return req
}

func newRequestSize(n int, prefix string) *request {
	req := requestPool.Get().(*request)
	req.buf = append(req.buf[:0], '*')
	req.buf = strconv.AppendUint(req.buf, uint64(n), 10)
	req.buf = append(req.buf, prefix...)
	return req
}

func (r *request) str(v string) {
	r.buf = strconv.AppendUint(r.buf, uint64(len(v)), 10)
	r.buf = append(r.buf, "\r\n"...)
	r.buf = append(r.buf, v...)
}

func (r *request) int(v int64) {
	sizeOffset := len(r.buf)
	sizeSingleDigit := v > -1e8 && v < 1e9
	if sizeSingleDigit {
		r.buf = append(r.buf, 0, '\r', '\n')
	} else {
		r.buf = append(r.buf, 0, 0, '\r', '\n')
	}

	valueOffset := len(r.buf)
	r.buf = strconv.AppendInt(r.buf, v, 10)
	size := len(r.buf) - valueOffset
	if sizeSingleDigit {
		r.buf[sizeOffset] = byte(size + '0')
	} else { // two digits
		r.buf[sizeOffset] = byte(size/10 + '0')
		r.buf[sizeOffset+1] = byte(size%10 + '0')
	}
}

func (r *request) addString(a string) {
	r.str(a)
	r.buf = append(r.buf, "\r\n"...)
}

func (r *request) addString2(a, b string) {
	r.str(a)
	r.buf = append(r.buf, "\r\n$"...)
	r.str(b)
	r.buf = append(r.buf, "\r\n"...)
}

func (r *request) addString3(a, b, c string) {
	r.str(a)
	r.buf = append(r.buf, "\r\n$"...)
	r.str(b)
	r.buf = append(r.buf, "\r\n$"...)
	r.str(c)
	r.buf = append(r.buf, "\r\n"...)
}

func (r *request) addString4(a, b, c, d string) {
	r.str(a)
	r.buf = append(r.buf, "\r\n$"...)
	r.str(b)
	r.buf = append(r.buf, "\r\n$"...)
	r.str(c)
	r.buf = append(r.buf, "\r\n$"...)
	r.str(d)
	r.buf = append(r.buf, "\r\n"...)
}

func (r *request) addStringIntString(a string, b int64, c string) {
	r.str(a)
	r.buf = append(r.buf, "\r\n$"...)
	r.int(b)
	r.buf = append(r.buf, "\r\n$"...)
	r.str(c)
	r.buf = append(r.buf, "\r\n"...)
}

func (r *request) addString2AndInt(a, b string, c int64) {
	r.str(a)
	r.buf = append(r.buf, "\r\n$"...)
	r.str(b)
	r.buf = append(r.buf, "\r\n$"...)
	r.int(c)
	r.buf = append(r.buf, "\r\n"...)
}

func (r *request) addStringAndStrings(a string, b []string) {
	r.str(a)
	for _, s := range b {
		r.buf = append(r.buf, "\r\n$"...)
		r.str(s)
	}
	r.buf = append(r.buf, "\r\n"...)
}

func (r *request) addString2AndStrings(a, c string, b []string) {
	r.str(a)
	r.buf = append(r.buf, "\r\n$"...)
	r.str(c)
	for _, s := range b {
		r.buf = append(r.buf, "\r\n$"...)
		r.str(s)
	}
	r.buf = append(r.buf, "\r\n"...)
}

func (r *request) addStrings(a []string) {
	for _, s := range a {
		r.buf = append(r.buf, "\r\n$"...)
		r.str(s)
	}
	r.buf = append(r.buf, "\r\n"...)
}

func (r *request) addStringInt(a1 string, a2 int64) {
	r.str(a1)
	r.buf = append(r.buf, "\r\n$"...)
	r.int(a2)
	r.buf = append(r.buf, "\r\n"...)
}

func (r *request) addStringInt2(a1 string, a2, a3 int64) {
	r.str(a1)
	r.buf = append(r.buf, "\r\n$"...)
	r.int(a2)
	r.buf = append(r.buf, "\r\n$"...)
	r.int(a3)
	r.buf = append(r.buf, "\r\n"...)
}

func (r *request) addStringAndMap(a string, b map[string]Value) {
	r.str(a)
	for k, v := range b {
		r.buf = append(r.buf, "\r\n$"...)
		r.str(k)
		r.buf = append(r.buf, "\r\n$"...)
		r.str(v)
	}
	r.buf = append(r.buf, "\r\n"...)
}
