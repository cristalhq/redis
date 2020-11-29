package redis

import "github.com/go-redis/redis"

// HyperLogLog data structure can be used in order to count unique elements in a set
// using just a small constant amount of memory,
// specifically 12k bytes for every HyperLogLog (plus a few bytes for the key itself).
type HyperLogLog struct {
	name   string
	client *redis.Client
}

// NewHyperLogLog ...
func NewHyperLogLog(name string, client *redis.Client) *HyperLogLog {
	return &HyperLogLog{name: name, client: client}
}

// Add the specified elements to the specified HyperLogLog.
func (h *HyperLogLog) Add(values ...interface{}) (int64, error) {
	return h.client.PFAdd(h.name, values...).Result()
}

// Count return the approximated cardinality of the set(s).
func (h *HyperLogLog) Count() (int64, error) {
	return h.client.PFCount(h.name).Result()
}

// CountWith return the approximated cardinality of the set(s).
func (h *HyperLogLog) CountWith(keys ...string) (int64, error) {
	return h.client.PFCount(append([]string{h.name}, keys...)...).Result()
}

// Merge N HyperLogLogs, but with high constant times.
func (h *HyperLogLog) Merge(keys ...string) (string, error) {
	return h.client.PFMerge(h.name, keys...).Result()
}

// MergeInto N HyperLogLogs, but with high constant times.
func (h *HyperLogLog) MergeInto(dst string, keys ...string) (string, error) {
	return h.client.PFMerge(dst, keys...).Result()
}
