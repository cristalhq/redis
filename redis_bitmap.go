package redis

import (
	"context"
	"errors"
	"fmt"
)

// BitMap client for bitmap operations in Redis.
// See: https://redis.io/commands#bitmap
type BitMap struct {
	name string
	c    *Client
}

// NewBitMap returns new Redis Bitmap client.
func NewBitMap(name string, client *Client) BitMap {
	return BitMap{name: name, c: client}
}

// Name of the BitMap structure.
func (bm BitMap) Name() string { return bm.name }

// BitMapOp represents operations for BitOp command in Redis.
type BitMapOp string

// BitMapOp possible values.
const (
	AndOp BitMapOp = "AND"
	OrOp  BitMapOp = "OR"
	XorOp BitMapOp = "XOR"
	NotOp BitMapOp = "NOT"
)

// BitCount returns the number of set bits in a string between start and end indexes.
// See: https://redis.io/commands/bitcount
func (bm BitMap) BitCount(ctx context.Context, start, end int64) (int64, error) {
	req := newRequest("*4\r\n$8\r\nBITCOUNT\r\n$")
	req.addStringInt2(bm.name, start, end)
	return bm.c.cmdInt(ctx, req)
}

// BitCountByte returns the number of set bits in a string between start and end indexes.
// See: https://redis.io/commands/bitcount
func (bm BitMap) BitCountByte(ctx context.Context, start, end int64) (int64, error) {
	req := newRequest("*4\r\n$8\r\nBITCOUNT\r\n$")
	req.addStringInt2String(bm.name, start, end, "BYTE")
	return bm.c.cmdInt(ctx, req)
}

// BitCountAll returns the number of set bits in a string.
// See: https://redis.io/commands/bitcount
func (bm BitMap) BitCountAll(ctx context.Context) (int64, error) {
	req := newRequest("*2\r\n$8\r\nBITCOUNT\r\n$")
	req.addString(bm.name)
	return bm.c.cmdInt(ctx, req)
}

// BitField TODO
// See: https://redis.io/commands/bitfield
func (bm BitMap) BitField(ctx context.Context) error {
	panic("redis: BitMap.BitField not implemented")
}

// BitFieldReadOnly TODO
// See: https://redis.io/commands/bitfield_ro
func (bm BitMap) BitFieldReadOnly(ctx context.Context) error {
	panic("redis: BitMap.BitFieldReadOnly not implemented")
}

// BitOp perform a bitwise operation between multiple keys (containing string values)
// and store the result in the destkey key.
//
// See: https://redis.io/commands/bitop
func (bm BitMap) BitOp(ctx context.Context, op BitMapOp, destination string, keys ...string) (int64, error) {
	if op != AndOp && op != OrOp && op != XorOp && op != NotOp {
		return 0, fmt.Errorf("unknown BitMap operation: %s", op)
	}
	if op == NotOp && len(keys) != 1 {
		return 0, errors.New("BitMap Not operation works only with 1 key")
	}
	req := newRequestSize(2+len(keys), "\r\n$5\r\nBITOP\r\n$")
	req.addStringAndStrings(string(op), append([]string{destination}, keys...))
	return bm.c.cmdInt(ctx, req)
}

// BitPos returns the position of the first bit set to 1 or 0 in a string.
// See: https://redis.io/commands/bitpos
func (bm BitMap) BitPos(ctx context.Context, bit int64, pos ...int64) (int64, error) {
	req := newRequestSize(2+len(pos), "\r\n$6\r\nBITPOS\r\n$")
	req.addStringInt(bm.name, bit)
	return bm.c.cmdInt(ctx, req)
}

// GetBit returns the bit value at offset.
// See: https://redis.io/commands/getbit
func (bm BitMap) GetBit(ctx context.Context, offset int64) (int64, error) {
	req := newRequest("*3\r\n$6\r\nGETBIT\r\n$")
	req.addStringInt(bm.name, offset)
	return bm.c.cmdInt(ctx, req)
}

// SetBit sets or clears the bit at offset.
// See: https://redis.io/commands/setbit
func (bm BitMap) SetBit(ctx context.Context, offset int64, value int) (int64, error) {
	req := newRequest("*4\r\n$6\r\nSETBIT\r\n$")
	req.addStringInt2(bm.name, offset, int64(value))
	return bm.c.cmdInt(ctx, req)
}
