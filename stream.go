package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Stream ...
type Stream struct {
	name   string
	client *redisClient
}

type (
	// StreamInfo ...
	StreamInfo = redis.XStream
	// StreamMessage ...
	StreamMessage = redis.XMessage
	// StreamPending ...
	StreamPending = redis.XPending
)

// NewStream instantiates a new Stream structure client for Redis.
func NewStream(name string, client *redisClient) *Stream {
	return &Stream{name: name, client: client}
}

// Ack ...
func (s *Stream) Ack(ctx context.Context, group string, ids ...string) (int64, error) {
	return s.client.XAck(ctx, s.name, group, ids...).Result()
}

// Add ...
func (s *Stream) Add(ctx context.Context, values map[string]interface{}) (string, error) {
	return s.client.XAdd(ctx, &redis.XAddArgs{
		Stream: s.name,
		Values: values,
	}).Result()
}

// AutoClaim ...
// TODO: https://redis.io/commands/xautoclaim
func (s *Stream) AutoClaim(ctx context.Context) ([]StreamMessage, error) {
	return nil, nil
}

// Claim ...
func (s *Stream) Claim(ctx context.Context) ([]StreamMessage, error) {
	return s.client.XClaim(ctx, &redis.XClaimArgs{
		Stream: s.name,
	}).Result()
}

// Delete ...
func (s *Stream) Delete(ctx context.Context, ids ...string) (int64, error) {
	return s.client.XDel(ctx, s.name, ids...).Result()
}

// Group ...
// TODO
func (s *Stream) Group(ctx context.Context) (int64, error) {
	return 0, nil
}

// Info ...
// TODO
func (s *Stream) Info(ctx context.Context) (int64, error) {
	return 0, nil
}

// Len ...
func (s *Stream) Len(ctx context.Context) (int64, error) {
	return s.client.XLen(ctx, s.name).Result()
}

// Pending ...
func (s *Stream) Pending(ctx context.Context, group string) (*StreamPending, error) {
	return s.client.XPending(ctx, s.name, group).Result()
}

// Range ...
func (s *Stream) Range(ctx context.Context, start, stop string) ([]StreamMessage, error) {
	return s.client.XRange(ctx, s.name, start, stop).Result()
}

// Read ...
func (s *Stream) Read(ctx context.Context) ([]StreamInfo, error) {
	return s.client.XRead(ctx, &redis.XReadArgs{Streams: []string{s.name}}).Result()
}

// ReadGroup ...
func (s *Stream) ReadGroup(ctx context.Context) ([]StreamInfo, error) {
	return s.client.XReadGroup(ctx, &redis.XReadGroupArgs{Streams: []string{s.name}}).Result()
}

// ReverseRange ...
func (s *Stream) ReverseRange(ctx context.Context, start, stop string) ([]StreamMessage, error) {
	return s.client.XRevRange(ctx, s.name, start, stop).Result()
}

// Trim ...
func (s *Stream) Trim(ctx context.Context, maxLen int64) (int64, error) {
	return s.client.XTrim(ctx, s.name, maxLen).Result()
}
