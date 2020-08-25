package redis

import (
	"errors"

	"github.com/go-redis/redis"
)

type BitMap struct {
	name   string
	client *redis.Client
}

type BitMapOp string

const (
	AndOp BitMapOp = "and"
	OrOp  BitMapOp = "or"
	XorOp BitMapOp = "xor"
	NotOp BitMapOp = "not"
)

func NewBitMap(name string, client *redis.Client) *BitMap {
	return &BitMap{name: name, client: client}
}

func (bm *BitMap) GetBit(offset int64) (int64, error) {
	return bm.client.GetBit(bm.name, offset).Result()
}

func (bm *BitMap) SetBit(offset int64, value int) (int64, error) {
	return bm.client.SetBit(bm.name, offset, value).Result()
}

func (bm *BitMap) BitCount(start, stop int64) (int64, error) {
	return bm.client.BitCount(bm.name, &redis.BitCount{Start: start, End: stop}).Result()
}

func (bm *BitMap) BitOp(op BitMapOp, keys ...string) (int64, error) {
	switch op {
	case AndOp:
		return bm.client.BitOpAnd(bm.name, keys...).Result()
	case OrOp:
		return bm.client.BitOpOr(bm.name, keys...).Result()
	case XorOp:
		return bm.client.BitOpXor(bm.name, keys...).Result()
	case NotOp:
		return bm.client.BitOpNot(bm.name, keys[0]).Result()
	default:
		return 0, errors.New("unknown BitMap operation")
	}
}

func (bm *BitMap) BitPos(bit int64, pos ...int64) (int64, error) {
	return bm.client.BitPos(bm.name, bit, pos...).Result()
}
