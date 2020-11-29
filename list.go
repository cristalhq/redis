package redis

import (
	"time"

	"github.com/go-redis/redis"
)

// List ...
type List struct {
	name   string
	client *redis.Client
}

// NewList ...
func NewList(name string, client *redis.Client) *List {
	return &List{name: name, client: client}
}

// Len ...
func (l *List) Len() (int64, error) {
	return l.client.LLen(l.name).Result()
}

// Index ...
func (l *List) Index(index int64) (string, error) {
	return l.client.LIndex(l.name, index).Result()
}

// LPOS ...
// TODO
func (l *List) LPOS() (int64, error) {
	return 0, nil
}

// Insert ...
func (l *List) Insert(op string, pivot, value interface{}) (int64, error) {
	return l.client.LInsert(l.name, op, pivot, value).Result()
}

// Set ...
func (l *List) Set(index int64, value interface{}) (string, error) {
	return l.client.LSet(l.name, index, value).Result()
}

// Remove ...
func (l *List) Remove(count int64, value interface{}) (int64, error) {
	return l.client.LRem(l.name, count, value).Result()
}

// LeftPop ...
func (l *List) LeftPop() (string, error) {
	return l.client.LPop(l.name).Result()
}

// RightPop ...
func (l *List) RightPop() (string, error) {
	return l.client.RPop(l.name).Result()
}

// LeftPush ...
func (l *List) LeftPush(values ...interface{}) (int64, error) {
	return l.client.LPush(l.name, values...).Result()
}

// RightPush ...
func (l *List) RightPush(values ...interface{}) (int64, error) {
	return l.client.RPush(l.name, values).Result()
}

// LeftPushExist ...
// TODO: must be `values []inteface{}`
func (l *List) LeftPushExist(value interface{}) (int64, error) {
	return l.client.LPushX(l.name, value).Result()
}

// RightPushExist ...
// TODO: must be `values []inteface{}`
func (l *List) RightPushExist(value interface{}) (int64, error) {
	return l.client.RPushX(l.name, value).Result()
}

// Range ...
func (l *List) Range(start, stop int64) ([]string, error) {
	return l.client.LRange(l.name, start, stop).Result()
}

// Trim ...
func (l *List) Trim(start, stop int64) (string, error) {
	return l.client.LTrim(l.name, start, stop).Result()
}

// RightPopLeftPush ...
func (l *List) RightPopLeftPush() (string, error) {
	return l.client.RPopLPush("", "").Result()
}

// BlockingLeftPop ...
func (l *List) BlockingLeftPop(keys ...string) ([]string, error) {
	return l.client.BLPop(time.Second, keys...).Result()
}

// BlockingRightPop ...
func (l *List) BlockingRightPop() ([]string, error) {
	return l.client.BRPop(time.Second, "").Result()
}

// BlockingRightPopLeftPush ...
func (l *List) BlockingRightPopLeftPush() (string, error) {
	return l.client.BRPopLPush("", "", time.Second).Result()
}
