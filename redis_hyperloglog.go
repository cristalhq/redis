package redis

import (
	"context"
)

// HyperLogLog data structure can be used in order to count unique elements in a set
// using just a small constant amount of memory,
// specifically 12k bytes for every HyperLogLog (plus a few bytes for the key itself).
type HyperLogLog struct {
	name string
	c    *Client
}

// NewHyperLogLog instantiates new Redis HyperLogLog client.
func NewHyperLogLog(name string, client *Client) HyperLogLog {
	return HyperLogLog{name: name, c: client}
}

// Add the specified elements to the specified HyperLogLog.
// See: https://redis.io/commands/pfadd
func (h HyperLogLog) Add(ctx context.Context, values ...string) (int64, error) {
	req := newRequestSize(2+len(values), "\r\n$5\r\nPFADD\r\n$")
	req.addStringAndStrings(h.name, values)
	return h.c.cmdInt(ctx, req)
}

// Count return the approximated cardinality of the set.
// See: https://redis.io/commands/pfcount
func (h HyperLogLog) Count(ctx context.Context) (int64, error) {
	req := newRequest("*2\r\n$7\r\nPFCOUNT\r\n$")
	req.addString(h.name)
	return h.c.cmdInt(ctx, req)
}

// CountWith return the approximated cardinality of the sets.
// See: https://redis.io/commands/pfcount
func (h HyperLogLog) CountWith(ctx context.Context, keys ...string) (int64, error) {
	req := newRequestSize(2+len(keys), "\r\n$7\r\nPFCOUNT\r\n$")
	req.addStringAndStrings(h.name, keys)
	return h.c.cmdInt(ctx, req)
}

// Merge N HyperLogLogs, but with high constant times.
// See: https://redis.io/commands/pfmerge
func (h HyperLogLog) Merge(ctx context.Context, keys ...string) error {
	req := newRequestSize(1+len(keys), "\r\n$7\r\nPFMERGE\r\n$")
	req.addStringAndStrings(h.name, keys)
	return h.c.cmdSimple(ctx, req)
}

// MergeInto N HyperLogLogs, but with high constant times.
// See: https://redis.io/commands/pfmerge
func (h HyperLogLog) MergeInto(ctx context.Context, destination string, keys ...string) error {
	req := newRequestSize(2+len(keys), "\r\n$7\r\nPFMERGE\r\n$")
	req.addStringAndStrings(destination, keys)
	return h.c.cmdSimple(ctx, req)
}
