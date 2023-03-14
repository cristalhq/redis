package redis

import (
	"context"
)

// SortedSet client for sorted set operations in Redis.
// See: https://redis.io/commands#sorted-set
type SortedSet struct {
	name string
	c    *Client
}

// NewSortedSet returns new Redis SortedSet client.
func NewSortedSet(name string, client *Client) SortedSet {
	return SortedSet{name: name, c: client}
}

// SortedSetItem TODO
type SortedSetItem struct {
	Score  float64
	Member string
}

type (
	// SortedSetStore ...
	SortedSetStore = interface{} // redis.ZStore
	// SortedSetRangeBy ...
	SortedSetRangeBy = interface{} // redis.ZRangeBy
	// SortedSetWithKey ...
	SortedSetWithKey = interface{} // redis.ZWithKey
)

// BlockingPopMax ...
// BZMPOP
// See: https://redis.io/commands/bzmpop
func (ss SortedSet) BlockingPop(ctx context.Context) (*SortedSetWithKey, error) {
	panic("redis: SortedSet.BlockingPopMax not implemented")
}

// BlockingPopMax ...
// BZPOPMAX
// See: https://redis.io/commands/bzpopmax
func (ss SortedSet) BlockingPopMax(ctx context.Context) (*SortedSetWithKey, error) {
	panic("redis: SortedSet.BlockingPopMax not implemented")
}

// BlockingPopMin ...
// BZPOPMIN
// See: https://redis.io/commands/bzpopmin
func (ss SortedSet) BlockingPopMin(ctx context.Context) (*SortedSetWithKey, error) {
	panic("redis: SortedSet.BlockingPopMin not implemented")
}

// Add all the specified members with the specified scores to the sorted set.
// See: https://redis.io/commands/zadd
func (ss SortedSet) Add(ctx context.Context, items ...SortedSetItem) (int64, error) {
	return 0, nil
	// req := newRequestSize(2+2*len(items), "\r\n$4\r\nZADD\r\n$")
	// req.addString(ss.name)
	// for range items {
	// 	// req.addStringInt(ss.name)
	// }
	// return ss.c.cmdInt(ctx, req)
}

// Cardinality returns the sorted set cardinality (number of elements).
// See: https://redis.io/commands/zcard
func (ss SortedSet) Cardinality(ctx context.Context) (int64, error) {
	req := newRequest("*2\r\n$5\r\nZCARD\r\n$")
	req.addString(ss.name)
	return ss.c.cmdInt(ctx, req)
}

// Count ...
// See: https://redis.io/commands/zcount
func (ss SortedSet) Count(ctx context.Context, min, max string) (int64, error) {
	req := newRequest("*2\r\n$6\r\nZCOUNT\r\n$")
	req.addString(ss.name)
	return ss.c.cmdInt(ctx, req)
}

// Diff ...
// See: https://redis.io/commands/zdiff
func (ss SortedSet) Diff(ctx context.Context, keys ...string) ([]string, error) {
	req := newRequest("*2\r\n$6\r\nZDIFF\r\n$")
	req.addString(ss.name)
	return ss.c.cmdStrings(ctx, req)
}

// DiffWithScores ...
// --See: https://redis.io/commands/zdiff
func (ss SortedSet) DiffWithScores(ctx context.Context, keys ...string) ([]string, error) {
	req := newRequest("*2\r\n$6\r\nZDIFF\r\n$")
	req.addString(ss.name)
	return ss.c.cmdStrings(ctx, req)
}

// DiffStore ...
// See: https://redis.io/commands/zdiffstore
func (ss SortedSet) DiffStore(ctx context.Context, destKey string, min, max string) (int64, error) {
	req := newRequest("*2\r\n$6\r\nZDIFFSTORE\r\n$")
	req.addString(ss.name)
	return ss.c.cmdInt(ctx, req)
}

// ZINCRBY TODO
// See: https://redis.io/commands/zincrby
func (ss SortedSet) ZINCRBY(ctx context.Context) error { return nil }

// ZINTER TODO
// See: https://redis.io/commands/zinter
func (ss SortedSet) ZINTER(ctx context.Context) error { return nil }

// ZINTERCARD TODO
// See: https://redis.io/commands/zintercard
func (ss SortedSet) ZINTERCARD(ctx context.Context) error { return nil }

// ZINTERSTORE TODO
// See: https://redis.io/commands/zinterstore
func (ss SortedSet) ZINTERSTORE(ctx context.Context) error { return nil }

// LexCount TODO
// See: https://redis.io/commands/zlexcount
func (ss SortedSet) LexCount(ctx context.Context, min, max string) (int64, error) {
	req := newRequest("*4\r\n$9\r\nZLEXCOUNT\r\n$")
	req.addString3(ss.name, min, max)
	return ss.c.cmdInt(ctx, req)
}

// ZMPOP TODO
// See: https://redis.io/commands/zmpop
func (ss SortedSet) ZMPOP(ctx context.Context) error { return nil }

// ZMSCORE TODO
// See: https://redis.io/commands/zmscore
func (ss SortedSet) MultiScore(ctx context.Context, keys ...string) error {
	return nil
}

// PopMax removes and returns up to count members with the highest scores in the sorted set.
// See: https://redis.io/commands/zpopmax
func (ss SortedSet) PopMax(ctx context.Context, count int) ([]string, error) {
	req := newRequest("*3\r\n$7\r\nZPOPMAX\r\n$")
	req.addStringInt(ss.name, int64(count))
	return ss.c.cmdStrings(ctx, req)
}

// PopMin removes and returns up to count members with the lowest scores in the sorted set.
// See: https://redis.io/commands/zpopmin
func (ss SortedSet) PopMin(ctx context.Context, count int) ([]string, error) {
	req := newRequest("*3\r\n$7\r\nZPOPMIN\r\n$")
	req.addStringInt(ss.name, int64(count))
	return ss.c.cmdStrings(ctx, req)
}

// ZRANDMEMBER TODO
// See: https://redis.io/commands/zrandmember
func (ss SortedSet) ZRANDMEMBER(ctx context.Context) error { return nil }

// ZRANGE TODO
// See: https://redis.io/commands/zrange
func (ss SortedSet) ZRANGE(ctx context.Context) error { return nil }

// ZRANGEBYLEX TODO
// See: https://redis.io/commands/zrangebylex
func (ss SortedSet) ZRANGEBYLEX(ctx context.Context) error { return nil }

// ZRANGEBYSCORE TODO
// See: https://redis.io/commands/zrangebyscore
func (ss SortedSet) ZRANGEBYSCORE(ctx context.Context) error { return nil }

// ZRANGESTORE TODO
// See: https://redis.io/commands/zrangestore
func (ss SortedSet) ZRANGESTORE(ctx context.Context) error { return nil }

// ZRANK TODO
// See: https://redis.io/commands/zrank
func (ss SortedSet) ZRANK(ctx context.Context) error { return nil }

// Remove removes the specified members from the sorted set stored.
// Non existing members are ignored.
// See: https://redis.io/commands/zrem
func (ss SortedSet) Remove(ctx context.Context, members ...string) (int64, error) {
	req := newRequestSize(2+len(members), "\r\n$4\r\nZREM\r\n$")
	req.addStringAndStrings(ss.name, members)
	return ss.c.cmdInt(ctx, req)
}

// ZREMRANGEBYLEX TODO
// See: https://redis.io/commands/zremrangebylex
func (ss SortedSet) ZREMRANGEBYLEX(ctx context.Context) error { return nil }

// ZREMRANGEBYRANK TODO
// See: https://redis.io/commands/zremrangebyrank
func (ss SortedSet) ZREMRANGEBYRANK(ctx context.Context) error { return nil }

// ZREMRANGEBYSCORE TODO
// See: https://redis.io/commands/zremrangebyscore
func (ss SortedSet) ZREMRANGEBYSCORE(ctx context.Context) error { return nil }

// ZREVRANGE TODO
// See: https://redis.io/commands/zrevrange
func (ss SortedSet) ZREVRANGE(ctx context.Context) error { return nil }

// ZREVRANGEBYLEX TODO
// See: https://redis.io/commands/zrevrangebylex
func (ss SortedSet) ZREVRANGEBYLEX(ctx context.Context) error { return nil }

// ZREVRANGEBYSCORE TODO
// See: https://redis.io/commands/zrevrangebyscore
func (ss SortedSet) ZREVRANGEBYSCORE(ctx context.Context) error { return nil }

// ZREVRANK TODO
// See: https://redis.io/commands/zrevrank
func (ss SortedSet) ZREVRANK(ctx context.Context) error { return nil }

// ZSCAN TODO
// See: https://redis.io/commands/zscan
func (ss SortedSet) ZSCAN(ctx context.Context) error { return nil }

// Score returns the score of member in the sorted set.
// See: https://redis.io/commands/zscore
func (ss SortedSet) Score(ctx context.Context, member string) (string, error) {
	req := newRequest("*3\r\n$6\r\nZSCORE\r\n$")
	req.addString2(ss.name, member)
	return ss.c.cmdString(ctx, req)
}

// ZUNION TODO
// See: https://redis.io/commands/zunion
func (ss SortedSet) Union(ctx context.Context, keys ...string) ([]string, error) {
	req := newRequestSize(2+len(keys), "\r\n$6\r\nZUNION\r\n$")
	req.addStringAndStrings(ss.name, keys)
	return ss.c.cmdStrings(ctx, req)
}

// ZUNIONSTORE TODO
// See: https://redis.io/commands/zunionstore
func (ss SortedSet) ZUNIONSTORE(ctx context.Context, dst string, keys ...string) (int64, error) {
	req := newRequestSize(3+len(keys), "\r\n$11\r\nSUNIONSTORE\r\n$")
	req.addString2AndStrings(dst, ss.name, keys)
	return ss.c.cmdInt(ctx, req)
}
