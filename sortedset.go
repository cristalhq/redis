package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type SortedSet struct {
	name   string
	client *redis.Client
}

// NewSortedSet instantiates a new SortedSet structure client for Redis.
func NewSortedSet(name string, client *redis.Client) *SortedSet {
	return &SortedSet{name: name, client: client}
}

type (
	SortedSetItem    = redis.Z
	SortedSetStore   = redis.ZStore
	SortedSetRangeBy = redis.ZRangeBy
	SortedSetWithKey = redis.ZWithKey
)

func (ss *SortedSet) Add(items []SortedSetItem) (int64, error) {
	return ss.client.ZAdd(ss.name, items...).Result()
}

func (ss *SortedSet) Cardinality() (int64, error) {
	return ss.client.ZCard(ss.name).Result()
}

func (ss *SortedSet) Count(min, max string) (int64, error) {
	return ss.client.ZCount(ss.name, min, max).Result()
}

func (ss *SortedSet) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return ss.client.ZScan(ss.name, cursor, match, count).Result()
}

func (ss *SortedSet) Score(member string) (float64, error) {
	return ss.client.ZScore(ss.name, member).Result()
}

func (ss *SortedSet) IncBy(delta float64, member string) (float64, error) {
	return ss.client.ZIncrBy(ss.name, delta, member).Result()
}

func (ss *SortedSet) LexCount(min, max string) (int64, error) {
	return ss.client.ZLexCount(ss.name, min, max).Result()
}

func (ss *SortedSet) PopMax(count ...int64) ([]SortedSetItem, error) {
	return ss.client.ZPopMax(ss.name, count...).Result()
}

func (ss *SortedSet) PopMin(count ...int64) ([]SortedSetItem, error) {
	return ss.client.ZPopMin(ss.name, count...).Result()
}

func (ss *SortedSet) Range(start, stop int64) ([]string, error) {
	return ss.client.ZRange(ss.name, start, stop).Result()
}

func (ss *SortedSet) RangeByLex(opt SortedSetRangeBy) ([]string, error) {
	return ss.client.ZRangeByLex(ss.name, opt).Result()
}

func (ss *SortedSet) RangeByScore(opt SortedSetRangeBy) ([]string, error) {
	return ss.client.ZRangeByScore(ss.name, opt).Result()
}

func (ss *SortedSet) Rank(member string) (int64, error) {
	return ss.client.ZRank(ss.name, member).Result()
}

func (ss *SortedSet) ReverseRange(start, stop int64) ([]string, error) {
	return ss.client.ZRevRange(ss.name, start, stop).Result()
}

func (ss *SortedSet) ReverseRangeByLex(opt SortedSetRangeBy) ([]string, error) {
	return ss.client.ZRevRangeByLex(ss.name, opt).Result()
}

func (ss *SortedSet) ReverseRangeByScore(opt SortedSetRangeBy) ([]string, error) {
	return ss.client.ZRevRangeByScore(ss.name, opt).Result()
}

func (ss *SortedSet) ReverseRank(member string) (int64, error) {
	return ss.client.ZRevRank(ss.name, member).Result()
}

func (ss *SortedSet) Remove(members ...interface{}) (int64, error) {
	return ss.client.ZRem(ss.name, members...).Result()
}

func (ss *SortedSet) RemoveRangeByLex(min, max string) (int64, error) {
	return ss.client.ZRemRangeByLex(ss.name, min, max).Result()
}

func (ss *SortedSet) RemoveRangeByRank(start, stop int64) (int64, error) {
	return ss.client.ZRemRangeByRank(ss.name, start, stop).Result()
}

func (ss *SortedSet) RemoveRangeByScore(min, max string) (int64, error) {
	return ss.client.ZRemRangeByScore(ss.name, min, max).Result()
}

func (ss *SortedSet) IntersectionStore(store SortedSetStore, keys ...string) (int64, error) {
	return ss.client.ZInterStore("", store, keys...).Result()
}

func (ss *SortedSet) UnionStore(dest string, store SortedSetStore, keys ...string) (int64, error) {
	return ss.client.ZUnionStore("", store, keys...).Result()
}

func (ss *SortedSet) BlockingPopMax() (SortedSetWithKey, error) {
	return ss.client.BZPopMax(time.Second, "").Result()
}

func (ss *SortedSet) BlockingPopMin() (SortedSetWithKey, error) {
	return ss.client.BZPopMin(time.Second, "").Result()
}
