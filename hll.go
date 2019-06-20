package redis

import "github.com/go-redis/redis"

// HLL (the HyperLogLog) data structure can be used in order to count unique elements in a set
// using just a small constant amount of memory,
// specifically 12k bytes for every HyperLogLog (plus a few bytes for the key itself).
type HLL struct {
	client *redis.Client
}

// PFAdd adds the specified elements to the specified HyperLogLog.
func (h *HLL) PFAdd(key string, values ...interface{}) (int64, error) {
	resp := h.client.PFAdd(key, values...)
	return resp.Result()
}

// PFCount return the approximated cardinality of the set(s).
func (h *HLL) PFCount(keys ...string) (int64, error) {
	resp := h.client.PFCount(keys...)
	return resp.Result()
}

// PFMerge merge N HyperLogLogs, but with high constant times.
func (h *HLL) PFMerge(dest string, keys ...string) error {
	resp := h.client.PFMerge(dest, keys...)
	return resp.Err()
}
