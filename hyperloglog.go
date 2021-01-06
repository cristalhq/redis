package redis

import (
	"context"
)

// HyperLogLog data structure can be used in order to count unique elements in a set
// using just a small constant amount of memory,
// specifically 12k bytes for every HyperLogLog (plus a few bytes for the key itself).
type HyperLogLog struct {
	name   string
	client *redisClient
}

// NewHyperLogLog ...
func NewHyperLogLog(name string, client *redisClient) *HyperLogLog {
	return &HyperLogLog{name: name, client: client}
}

// Add the specified elements to the specified HyperLogLog.
func (h *HyperLogLog) Add(ctx context.Context, values ...interface{}) (int64, error) {
	return h.client.PFAdd(ctx, h.name, values...).Result()
}

// Count return the approximated cardinality of the set(s).
func (h *HyperLogLog) Count(ctx context.Context) (int64, error) {
	return h.client.PFCount(ctx, h.name).Result()
}

// CountWith return the approximated cardinality of the set(s).
func (h *HyperLogLog) CountWith(ctx context.Context, keys ...string) (int64, error) {
	return h.client.PFCount(ctx, append([]string{h.name}, keys...)...).Result()
}

// Merge N HyperLogLogs, but with high constant times.
func (h *HyperLogLog) Merge(ctx context.Context, keys ...string) (string, error) {
	return h.client.PFMerge(ctx, h.name, keys...).Result()
}

// MergeInto N HyperLogLogs, but with high constant times.
func (h *HyperLogLog) MergeInto(ctx context.Context, dst string, keys ...string) (string, error) {
	return h.client.PFMerge(ctx, dst, keys...).Result()
}
