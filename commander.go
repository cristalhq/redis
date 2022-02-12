package redis

import "context"

// Commander is a simple Redis client.
type Commander struct {
	c *Client
}

// NewCommander returns new Commander.
func NewCommander(client *Client) Commander {
	return Commander{c: client}
}

// BitCount does BITCOUNT command, see: https://redis.io/commands/bitcount
func (c Commander) BitCount(ctx context.Context, key string, start, end int64) (int64, error) {
	return NewBitMap(key, c.c).BitCount(ctx, start, end)
}

// BitCountAll does BITCOUNT command, see: https://redis.io/commands/bitcount
func (c Commander) BitCountAll(ctx context.Context, key string) (int64, error) {
	return NewBitMap(key, c.c).BitCountAll(ctx)
}

// BitField does BITFIELD command, see: https://redis.io/commands/bitfield
func (c Commander) BitField(ctx context.Context, key string) error {
	return NewBitMap(key, c.c).BitField(ctx)
}

// BitField does BITFIELD_RO command, see: https://redis.io/commands/bitfield_ro
func (c Commander) BitFieldReadOnly(ctx context.Context, key string) error {
	return NewBitMap(key, c.c).BitFieldReadOnly(ctx)
}

// BitOp does BITOP command, see: https://redis.io/commands/bitop
func (c Commander) BitOp(ctx context.Context, op BitMapOp, destKey string, keys ...string) (int64, error) {
	return NewBitMap(destKey, c.c).BitOp(ctx, op, destKey, keys...)
}

// BitPos does BITPOS command, see: https://redis.io/commands/bitpos
func (c Commander) BitPos(ctx context.Context, key string, bit int64, pos ...int64) (int64, error) {
	return NewBitMap(key, c.c).BitPos(ctx, bit, pos...)
}

// GetBit does GETBIT command, see: https://redis.io/commands/getbit
func (c Commander) GetBit(ctx context.Context, key string, offset int64) (int64, error) {
	return NewBitMap(key, c.c).GetBit(ctx, offset)
}

// SetBit does SETBIT command, see: https://redis.io/commands/setbit
func (c Commander) SetBit(ctx context.Context, key string, offset int64, value int) (int64, error) {
	return NewBitMap(key, c.c).SetBit(ctx, offset, value)
}

// BLMove does BLMOVE, see https://redis.io/commands/blmove
func (c *Commander) BLMove(ctx context.Context) error {
	return NewList("", c.c).BlockingLeftMove(ctx)
}

// BLMPop does BLMPOP, see https://redis.io/commands/blmpop
func (c *Commander) BLMPop(ctx context.Context) error {
	return NewList("", c.c).BlockingLeftMultiPop(ctx)
}

// BLPop does BLPOP, see https://redis.io/commands/blPop
func (c *Commander) BLPop(ctx context.Context) error {
	return NewList("", c.c).BlockingLeftPop(ctx)
}

// BRPop does BRPOP, see https://redis.io/commands/brpop
func (c *Commander) BRPop(ctx context.Context) error {
	return NewList("", c.c).BlockingRightPop(ctx)
}

// BRPopLPush does BRPOPLPUSH, see https://redis.io/commands/brpoplpush
func (c *Commander) BRPopLPush(ctx context.Context) error {
	return NewList("", c.c).BlockingRightPopLeftPush(ctx)
}

// LIndex does LINDEX, see https://redis.io/commands/lindex
func (c *Commander) LIndex(ctx context.Context, key string, index int64) (string, error) {
	return NewList(key, c.c).Index(ctx, index)
}

// LInsert does LINSERT, see https://redis.io/commands/linsert
func (c *Commander) LInsert(ctx context.Context, key, op string, pivot, value string) (int64, error) {
	return NewList(key, c.c).Insert(ctx, op, pivot, value)
}

// LLen does LLEN, see https://redis.io/commands/llen
func (c *Commander) LLen(ctx context.Context, key string) (int64, error) {
	return NewList(key, c.c).Len(ctx)
}

// LMove does LMOVE, see https://redis.io/commands/lmove
func (c *Commander) LMove(ctx context.Context, src, dst, srcpos, destpos string) (string, error) {
	return NewList(src, c.c).LeftMove(ctx, dst, srcpos, destpos)
}

// LMPop does LMPOP, see https://redis.io/commands/lmpop
func (c *Commander) LMPop(ctx context.Context) error {
	return NewList("", c.c).LeftMultiPop(ctx)
}

// LPop does LPOP, see https://redis.io/commands/lpop
func (c *Commander) LPop(ctx context.Context, key string) (string, error) {
	res, err := NewList(key, c.c).LeftPop(ctx, 1)
	return res[0], err
}

// LPos does LPOS, see https://redis.io/commands/lpos
func (c *Commander) LPos(ctx context.Context, key string) error {
	return NewList(key, c.c).LeftPos(ctx)
}

// LPush does LPUSH, see https://redis.io/commands/lpush
func (c *Commander) LPush(ctx context.Context, key string, elements ...string) (int64, error) {
	return NewList(key, c.c).LeftPush(ctx, elements...)
}

// LPushX does LPUSHX, see https://redis.io/commands/lpushx
func (c *Commander) LPushX(ctx context.Context, key string, elements ...string) (int64, error) {
	return NewList(key, c.c).LeftPushX(ctx, elements...)
}

// LRange does LRANGE, see https://redis.io/commands/lrange
func (c *Commander) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return NewList(key, c.c).Range(ctx, start, stop)
}

// LRem does LREM, see https://redis.io/commands/lrem
func (c *Commander) LRem(ctx context.Context, key string, count int64, value string) (int64, error) {
	return NewList(key, c.c).Remove(ctx, count, value)
}

// LSet does LSET, see https://redis.io/commands/lset
func (c *Commander) LSet(ctx context.Context, key string, index int64, value string) error {
	return NewList(key, c.c).Set(ctx, index, value)
}

// LTrim does LTRIM, see https://redis.io/commands/ltrim
func (c *Commander) LTrim(ctx context.Context, key string, start, stop int64) error {
	return NewList(key, c.c).Trim(ctx, start, stop)
}

// RPop does RPOP, see https://redis.io/commands/rpop
func (c *Commander) RPop(ctx context.Context, key string) (string, error) {
	res, err := NewList(key, c.c).RightPop(ctx, 1)
	return res[0], err
}

// RPopLPush does RPOPLPUSH, see https://redis.io/commands/rpoplpush
func (c *Commander) RPopLPush(ctx context.Context, src, dst string) (string, error) {
	return NewList(src, c.c).RightPopLeftPush(ctx, dst)
}

// RPush does RPUSH, see https://redis.io/commands/rpush
func (c *Commander) RPush(ctx context.Context, key string, elements ...string) (int64, error) {
	return NewList(key, c.c).RightPush(ctx, elements...)
}

// RPushX does RPUSHX, see https://redis.io/commands/rpushx
func (c *Commander) RPushX(ctx context.Context, key string, elements ...string) (int64, error) {
	return NewList(key, c.c).RightPushX(ctx, elements...)
}
