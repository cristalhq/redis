package redis

import (
	"strconv"
)

type request struct {
	buf []byte
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


func (req *request) addString(a string) {
	req.buf = strconv.AppendUint(req.buf, uint64(len(a)), 10)
	req.buf = append(req.buf, '\r', '\n')
	req.buf = append(req.buf, a...)
	req.buf = append(req.buf, '\r', '\n')
}

