package redis

import (
	"context"
)

// SortedSet ...
type SortedSet struct {
	name string
	c    *Client
}

// NewSortedSet instantiates a new SortedSet structure client for Redis.
func NewSortedSet(name string, client *Client) *SortedSet {
	return &SortedSet{name: name, c: client}
}

type (
	// SortedSetItem ...
	SortedSetItem = interface{} // redis.Z
	// SortedSetStore ...
	SortedSetStore = interface{} // redis.ZStore
	// SortedSetRangeBy ...
	SortedSetRangeBy = interface{} // redis.ZRangeBy
	// SortedSetWithKey ...
	SortedSetWithKey = interface{} // redis.ZWithKey
)

// Add ...
func (ss *SortedSet) Add(ctx context.Context, items ...*SortedSetItem) (int64, error) {
	req := newRequestSize(2+len(items), "\r\n$4\r\nZADD\r\n$")
	req.addString(ss.name)
	return ss.c.cmdInt(ctx, req)
}

// Cardinality ...
func (ss *SortedSet) Cardinality(ctx context.Context) (int64, error) {
	req := newRequest("*2\r\n$5\r\nZCARD\r\n$")
	req.addString(ss.name)
	return ss.c.cmdInt(ctx, req)
}

// Count ...
func (ss *SortedSet) Count(ctx context.Context, min, max string) (int64, error) {
	req := newRequest("*2\r\n$6\r\nZCOUNT\r\n$")
	req.addString(ss.name)
	return ss.c.cmdInt(ctx, req)
}

// Scan ...
func (ss *SortedSet) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	panic("redis: SortedSet.Scan not implemented")
}

// Score ...
func (ss *SortedSet) Score(ctx context.Context, member string) (float64, error) {
	panic("redis: SortedSet.Score not implemented")
}

// IncBy ...
func (ss *SortedSet) IncBy(ctx context.Context, delta float64, member string) (float64, error) {
	panic("redis: SortedSet.IncBy not implemented")
}

// LexCount ...
func (ss *SortedSet) LexCount(ctx context.Context, min, max string) (int64, error) {
	panic("redis: SortedSet.LexCount not implemented")
}

// PopMax ...
func (ss *SortedSet) PopMax(ctx context.Context, count ...int64) ([]SortedSetItem, error) {
	panic("redis: SortedSet.PopMax not implemented")
}

// PopMin ...
func (ss *SortedSet) PopMin(ctx context.Context, count ...int64) ([]SortedSetItem, error) {
	panic("redis: SortedSet.PopMin not implemented")
}

// Range ...
func (ss *SortedSet) Range(ctx context.Context, start, stop int64) ([]string, error) {
	panic("redis: SortedSet.Range not implemented")
}

// RangeByLex ...
func (ss *SortedSet) RangeByLex(ctx context.Context, opt *SortedSetRangeBy) ([]string, error) {
	panic("redis: SortedSet.RangeByLex not implemented")
}

// RangeByScore ...
func (ss *SortedSet) RangeByScore(ctx context.Context, opt *SortedSetRangeBy) ([]string, error) {
	panic("redis: SortedSet.RangeByScore not implemented")
}

// Rank ...
func (ss *SortedSet) Rank(ctx context.Context, member string) (int64, error) {
	panic("redis: SortedSet.Rank not implemented")
}

// ReverseRange ...
func (ss *SortedSet) ReverseRange(ctx context.Context, start, stop int64) ([]string, error) {
	panic("redis: SortedSet.ReverseRange not implemented")
}

// ReverseRangeByLex ...
func (ss *SortedSet) ReverseRangeByLex(ctx context.Context, opt *SortedSetRangeBy) ([]string, error) {
	panic("redis: SortedSet.ReverseRangeByLex not implemented")
}

// ReverseRangeByScore ...
func (ss *SortedSet) ReverseRangeByScore(ctx context.Context, opt *SortedSetRangeBy) ([]string, error) {
	panic("redis: SortedSet.ReverseRangeByScore not implemented")
}

// ReverseRank ...
func (ss *SortedSet) ReverseRank(ctx context.Context, member string) (int64, error) {
	panic("redis: SortedSet.ReverseRank not implemented")
}

// Remove ...
func (ss *SortedSet) Remove(ctx context.Context, members ...interface{}) (int64, error) {
	panic("redis: SortedSet.Remove not implemented")
}

// RemoveRangeByLex ...
func (ss *SortedSet) RemoveRangeByLex(ctx context.Context, min, max string) (int64, error) {
	panic("redis: SortedSet.RemoveRangeByLex not implemented")
}

// RemoveRangeByRank ...
func (ss *SortedSet) RemoveRangeByRank(ctx context.Context, start, stop int64) (int64, error) {
	panic("redis: SortedSet.RemoveRangeByRank not implemented")
}

// RemoveRangeByScore ...
func (ss *SortedSet) RemoveRangeByScore(ctx context.Context, min, max string) (int64, error) {
	panic("redis: SortedSet.RemoveRangeByScore not implemented")
}

// IntersectionStore ...
func (ss *SortedSet) IntersectionStore(ctx context.Context, store *SortedSetStore, _ ...string) (int64, error) {
	panic("redis: SortedSet.IntersectionStore not implemented")
}

// UnionStore ...
func (ss *SortedSet) UnionStore(ctx context.Context, dest string, store *SortedSetStore, _ ...string) (int64, error) {
	panic("redis: SortedSet.UnionStore not implemented")
}

// BlockingPopMax ...
func (ss *SortedSet) BlockingPopMax(ctx context.Context) (*SortedSetWithKey, error) {
	panic("redis: SortedSet.BlockingPopMax not implemented")
}

// BlockingPopMin ...
func (ss *SortedSet) BlockingPopMin(ctx context.Context) (*SortedSetWithKey, error) {
	panic("redis: SortedSet.BlockingPopMin not implemented")
}

/*
BZMPOP
BZPOPMAX
BZPOPMIN

ZCARD
ZCOUNT
ZDIFF
ZDIFFSTORE
ZINCRBY
ZINTER
ZINTERCARD
ZINTERSTORE
ZLEXCOUNT
ZMPOP
ZMSCORE
ZPOPMAX
ZPOPMIN
ZRANDMEMBER
ZRANGE
ZRANGEBYLEX
ZRANGEBYSCORE
ZRANGESTORE
ZRANK
ZREM
ZREMRANGEBYLEX
ZREMRANGEBYRANK
ZREMRANGEBYSCORE
ZREVRANGE
ZREVRANGEBYLEX
ZREVRANGEBYSCORE
ZREVRANK
ZSCAN
ZSCORE
ZUNION
ZUNIONSTORE
*/
