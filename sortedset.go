package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// SortedSet ...
type SortedSet struct {
	name   string
	client *redisClient
}

// NewSortedSet instantiates a new SortedSet structure client for Redis.
func NewSortedSet(name string, client *redisClient) *SortedSet {
	return &SortedSet{name: name, client: client}
}

type (
	// SortedSetItem ...
	SortedSetItem = redis.Z
	// SortedSetStore ...
	SortedSetStore = redis.ZStore
	// SortedSetRangeBy ...
	SortedSetRangeBy = redis.ZRangeBy
	// SortedSetWithKey ...
	SortedSetWithKey = redis.ZWithKey
)

// Add ...
func (ss *SortedSet) Add(ctx context.Context, items []*SortedSetItem) (int64, error) {
	return ss.client.ZAdd(ctx, ss.name, items...).Result()
}

// Cardinality ...
func (ss *SortedSet) Cardinality(ctx context.Context) (int64, error) {
	return ss.client.ZCard(ctx, ss.name).Result()
}

// Count ...
func (ss *SortedSet) Count(ctx context.Context, min, max string) (int64, error) {
	return ss.client.ZCount(ctx, ss.name, min, max).Result()
}

// Scan ...
func (ss *SortedSet) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return ss.client.ZScan(ctx, ss.name, cursor, match, count).Result()
}

// Score ...
func (ss *SortedSet) Score(ctx context.Context, member string) (float64, error) {
	return ss.client.ZScore(ctx, ss.name, member).Result()
}

// IncBy ...
func (ss *SortedSet) IncBy(ctx context.Context, delta float64, member string) (float64, error) {
	return ss.client.ZIncrBy(ctx, ss.name, delta, member).Result()
}

// LexCount ...
func (ss *SortedSet) LexCount(ctx context.Context, min, max string) (int64, error) {
	return ss.client.ZLexCount(ctx, ss.name, min, max).Result()
}

// PopMax ...
func (ss *SortedSet) PopMax(ctx context.Context, count ...int64) ([]SortedSetItem, error) {
	return ss.client.ZPopMax(ctx, ss.name, count...).Result()
}

// PopMin ...
func (ss *SortedSet) PopMin(ctx context.Context, count ...int64) ([]SortedSetItem, error) {
	return ss.client.ZPopMin(ctx, ss.name, count...).Result()
}

// Range ...
func (ss *SortedSet) Range(ctx context.Context, start, stop int64) ([]string, error) {
	return ss.client.ZRange(ctx, ss.name, start, stop).Result()
}

// RangeByLex ...
func (ss *SortedSet) RangeByLex(ctx context.Context, opt *SortedSetRangeBy) ([]string, error) {
	return ss.client.ZRangeByLex(ctx, ss.name, opt).Result()
}

// RangeByScore ...
func (ss *SortedSet) RangeByScore(ctx context.Context, opt *SortedSetRangeBy) ([]string, error) {
	return ss.client.ZRangeByScore(ctx, ss.name, opt).Result()
}

// Rank ...
func (ss *SortedSet) Rank(ctx context.Context, member string) (int64, error) {
	return ss.client.ZRank(ctx, ss.name, member).Result()
}

// ReverseRange ...
func (ss *SortedSet) ReverseRange(ctx context.Context, start, stop int64) ([]string, error) {
	return ss.client.ZRevRange(ctx, ss.name, start, stop).Result()
}

// ReverseRangeByLex ...
func (ss *SortedSet) ReverseRangeByLex(ctx context.Context, opt *SortedSetRangeBy) ([]string, error) {
	return ss.client.ZRevRangeByLex(ctx, ss.name, opt).Result()
}

// ReverseRangeByScore ...
func (ss *SortedSet) ReverseRangeByScore(ctx context.Context, opt *SortedSetRangeBy) ([]string, error) {
	return ss.client.ZRevRangeByScore(ctx, ss.name, opt).Result()
}

// ReverseRank ...
func (ss *SortedSet) ReverseRank(ctx context.Context, member string) (int64, error) {
	return ss.client.ZRevRank(ctx, ss.name, member).Result()
}

// Remove ...
func (ss *SortedSet) Remove(ctx context.Context, members ...interface{}) (int64, error) {
	return ss.client.ZRem(ctx, ss.name, members...).Result()
}

// RemoveRangeByLex ...
func (ss *SortedSet) RemoveRangeByLex(ctx context.Context, min, max string) (int64, error) {
	return ss.client.ZRemRangeByLex(ctx, ss.name, min, max).Result()
}

// RemoveRangeByRank ...
func (ss *SortedSet) RemoveRangeByRank(ctx context.Context, start, stop int64) (int64, error) {
	return ss.client.ZRemRangeByRank(ctx, ss.name, start, stop).Result()
}

// RemoveRangeByScore ...
func (ss *SortedSet) RemoveRangeByScore(ctx context.Context, min, max string) (int64, error) {
	return ss.client.ZRemRangeByScore(ctx, ss.name, min, max).Result()
}

// IntersectionStore ...
func (ss *SortedSet) IntersectionStore(ctx context.Context, store *SortedSetStore, _ ...string) (int64, error) {
	return ss.client.ZInterStore(ctx, "", store).Result()
}

// UnionStore ...
func (ss *SortedSet) UnionStore(ctx context.Context, dest string, store *SortedSetStore, _ ...string) (int64, error) {
	return ss.client.ZUnionStore(ctx, "", store).Result()
}

// BlockingPopMax ...
func (ss *SortedSet) BlockingPopMax(ctx context.Context) (*SortedSetWithKey, error) {
	return ss.client.BZPopMax(ctx, time.Second, "").Result()
}

// BlockingPopMin ...
func (ss *SortedSet) BlockingPopMin(ctx context.Context) (*SortedSetWithKey, error) {
	return ss.client.BZPopMin(ctx, time.Second, "").Result()
}
