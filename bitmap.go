package redis

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

// BitMap ...
type BitMap struct {
	name   string
	client *redisClient
}

// BitMapOp ...
type BitMapOp string

// BitMapOp consts.
const (
	AndOp BitMapOp = "and"
	OrOp  BitMapOp = "or"
	XorOp BitMapOp = "xor"
	NotOp BitMapOp = "not"
)

// NewBitMap ...
func NewBitMap(name string, client *redisClient) *BitMap {
	return &BitMap{name: name, client: client}
}

// GetBit ...
func (bm *BitMap) GetBit(ctx context.Context, offset int64) (int64, error) {
	return bm.client.GetBit(ctx, bm.name, offset).Result()
}

// SetBit ...
func (bm *BitMap) SetBit(ctx context.Context, offset int64, value int) (int64, error) {
	return bm.client.SetBit(ctx, bm.name, offset, value).Result()
}

// BitCount ...
func (bm *BitMap) BitCount(ctx context.Context, start, stop int64) (int64, error) {
	return bm.client.BitCount(ctx, bm.name, &redis.BitCount{Start: start, End: stop}).Result()
}

// BitOp ...
func (bm *BitMap) BitOp(ctx context.Context, op BitMapOp, keys ...string) (int64, error) {
	switch op {
	case AndOp:
		return bm.client.BitOpAnd(ctx, bm.name, keys...).Result()
	case OrOp:
		return bm.client.BitOpOr(ctx, bm.name, keys...).Result()
	case XorOp:
		return bm.client.BitOpXor(ctx, bm.name, keys...).Result()
	case NotOp:
		return bm.client.BitOpNot(ctx, bm.name, keys[0]).Result()
	default:
		return 0, errors.New("unknown BitMap operation")
	}
}

// BitPos ...
func (bm *BitMap) BitPos(ctx context.Context, bit int64, pos ...int64) (int64, error) {
	return bm.client.BitPos(ctx, bm.name, bit, pos...).Result()
}
