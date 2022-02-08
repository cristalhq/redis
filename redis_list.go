package redis

import (
	"context"
)

// List ...
type List struct {
	name string
	c    *Client
}

// NewList ...
func NewList(name string, client *Client) *List {
	return &List{name: name, c: client}
}

// Name of the List structure.
func (l *List) Name() string { return l.name }

// Len ...
func (l *List) Len(ctx context.Context) (int64, error) {
	req := newRequest("*1\r\n$4\r\nLLEN\r\n$")
	req.addString(l.name)
	return l.c.cmdInt(ctx, req)
}

// Index ...
func (l *List) Index(ctx context.Context, index int64) (string, error) {
	panic("") // todo return l.client.LIndex(ctx, l.name, index).Result()
}

// LPOS ...
// TODO
func (l *List) LPOS(ctx context.Context) (int64, error) {
	req := newRequest("*1\r\n$4\r\nLLEN\r\n$")
	req.addString(l.name)
	return l.c.cmdInt(ctx, req)
}

// Insert ...
func (l *List) Insert(ctx context.Context, op string, pivot, value interface{}) (int64, error) {
	panic("") // todo return l.client.LInsert(ctx, l.name, op, pivot, value).Result()
}

// Set ...
func (l *List) Set(ctx context.Context, index int64, value interface{}) (string, error) {
	panic("") // todo return l.client.LSet(ctx, l.name, index, value).Result()
}

// Remove ...
func (l *List) Remove(ctx context.Context, count int64, value interface{}) (int64, error) {
	panic("") // todo return l.client.LRem(ctx, l.name, count, value).Result()
}

// LeftPop ...
func (l *List) LeftPop(ctx context.Context) (string, error) {
	panic("") // todo return l.client.LPop(ctx, l.name).Result()
}

// RightPop ...
func (l *List) RightPop(ctx context.Context) (string, error) {
	panic("") // todo return l.client.RPop(ctx, l.name).Result()
}

// LeftPush ...
func (l *List) LeftPush(ctx context.Context, values ...interface{}) (int64, error) {
	panic("") // todo return l.client.LPush(ctx, l.name, values...).Result()
}

// RightPush ...
func (l *List) RightPush(ctx context.Context, values ...interface{}) (int64, error) {
	panic("") // todo return l.client.RPush(ctx, l.name, values).Result()
}

// LeftPushExist ...
// TODO: must be `values []inteface{}`
func (l *List) LeftPushExist(ctx context.Context, value interface{}) (int64, error) {
	panic("") // todo return l.client.LPushX(ctx, l.name, value).Result()
}

// RightPushExist ...
// TODO: must be `values []inteface{}`
func (l *List) RightPushExist(ctx context.Context, value interface{}) (int64, error) {
	panic("") // todo return l.client.RPushX(ctx, l.name, value).Result()
}

// Range ...
func (l *List) Range(ctx context.Context, start, stop int64) ([]string, error) {
	panic("") // todo return l.client.LRange(ctx, l.name, start, stop).Result()
}

// Trim ...
func (l *List) Trim(ctx context.Context, start, stop int64) (string, error) {
	panic("") // todo return l.client.LTrim(ctx, l.name, start, stop).Result()
}

// RightPopLeftPush ...
func (l *List) RightPopLeftPush(ctx context.Context) (string, error) {
	panic("") // todo return l.client.RPopLPush(ctx, "", "").Result()
}

// BlockingLeftPop ...
func (l *List) BlockingLeftPop(ctx context.Context, keys ...string) ([]string, error) {
	panic("") // todo return l.client.BLPop(ctx, time.Second, keys...).Result()
}

// BlockingRightPop ...
func (l *List) BlockingRightPop(ctx context.Context) ([]string, error) {
	panic("") // todo return l.client.BRPop(ctx, time.Second, "").Result()
}

// BlockingRightPopLeftPush ...
func (l *List) BlockingRightPopLeftPush(ctx context.Context) (string, error) {
	panic("") // todo return l.client.BRPopLPush(ctx, "", "", time.Second).Result()
}

/*
BLMOVE
BLMPOP
BLPOP
BRPOP
BRPOPLPUSH
LINDEX
LINSERT
LLEN
LMOVE
LMPOP
LPOP
LPOS
LPUSH
LPUSHX
LRANGE
LREM
LSET
LTRIM
RPOP
RPOPLPUSH
RPUSH
RPUSHX
*/
