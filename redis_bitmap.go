package redis

import (
	"context"
	"errors"
)

// BitMap client for bitmap operations in Redis.
// See: https://redis.io/commands#bitmap
type BitMap struct {
	name string
	c    *Client
}

// NewBitMap returns new Redis Bitmap client.
func NewBitMap(name string, client *Client) *BitMap {
	return &BitMap{name: name, c: client}
}

// BitMapOp represents operations for BitOp command in Redis.
type BitMapOp string

// BitMapOp possible values.
const (
	AndOp BitMapOp = "and"
	OrOp  BitMapOp = "or"
	XorOp BitMapOp = "xor"
	NotOp BitMapOp = "not"
)

// BitCount TODO
// See: https://redis.io/commands/bitcount
func (bm *BitMap) BitCount(ctx context.Context, start, end int64) (int64, error) {
	req := newRequest("*4\r\n$8\r\nBITCOUNT\r\n$")
	req.addStringInt2(bm.name, start, end)
	return bm.c.cmdInt(ctx, req)
}

// BitCountAll returns the number of set bits in a string.
// See: https://redis.io/commands/bitcount
func (bm *BitMap) BitCountAll(ctx context.Context) (int64, error) {
	req := newRequest("*2\r\n$8\r\nBITCOUNT\r\n$")
	req.addString(bm.name)
	return bm.c.cmdInt(ctx, req)
}

// BitField TODO
// See: https://redis.io/commands/bitfield
func (bm *BitMap) BitField(ctx context.Context) (int64, error) {
	req := newRequest("*2\r\n$8\r\nBITFIELD\r\n$")
	req.addString(bm.name)
	return bm.c.cmdInt(ctx, req)
}

// BitFieldReadOnly TODO
// See: https://redis.io/commands/bitfield_ro
func (bm *BitMap) BitFieldReadOnly(ctx context.Context, enc string, offset int64) (int64, error) {
	req := newRequest("*5\r\n$11\r\nBITFIELD_RO\r\n$")
	req.addStrings(append([]string{bm.name}, "GET"))
	req.addStringInt(enc, offset)
	return bm.c.cmdInt(ctx, req)
}

// BitOp TODO
// See: https://redis.io/commands/bitop
func (bm *BitMap) BitOp(ctx context.Context, op BitMapOp, keys ...string) (int64, error) {
	if op != AndOp && op != OrOp && op != XorOp && op != NotOp {
		return 0, errors.New("unknown BitMap operation")
	}
	if op == NotOp && len(keys) != 1 {
		return 0, errors.New("BitMap Not operation works only with 1 key")
	}
	req := newRequestSize(2+len(keys), "\r\n$5\r\nBITOP\r\n$")
	req.addString2AndStrings(bm.name, string(op), keys)
	return bm.c.cmdInt(ctx, req)
}

// BitPos TODO
// See: https://redis.io/commands/bitpos
func (bm *BitMap) BitPos(ctx context.Context, bit int64, pos ...int64) (int64, error) {
	req := newRequestSize(2+len(pos), "\r\n$6\r\nBITPOS\r\n$")
	req.addStringInt(bm.name, bit)
	return bm.c.cmdInt(ctx, req)
}

// GetBit TODO
// See: https://redis.io/commands/getbit
func (bm *BitMap) GetBit(ctx context.Context, offset int64) (int64, error) {
	req := newRequest("*3\r\n$6\r\nGETBIT\r\n$")
	req.addStringInt(bm.name, offset)
	return bm.c.cmdInt(ctx, req)
}

// SetBit TODO
// See: https://redis.io/commands/setbit
func (bm *BitMap) SetBit(ctx context.Context, offset int64, value int) (int64, error) {
	req := newRequest("*4\r\n$6\r\nSETBIT\r\n$")
	req.addStringInt2(bm.name, offset, int64(value))
	return bm.c.cmdInt(ctx, req)
}
