package redis

import (
	"github.com/go-redis/redis"
)

// Stream ...
type Stream struct {
	name   string
	client *redis.Client
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
func NewStream(name string, client *redis.Client) *Stream {
	return &Stream{name: name, client: client}
}

// Ack ...
func (s *Stream) Ack(group string, ids ...string) (int64, error) {
	return s.client.XAck(s.name, group, ids...).Result()
}

// Add ...
func (s *Stream) Add(values map[string]interface{}) (string, error) {
	return s.client.XAdd(&redis.XAddArgs{
		Stream: s.name,
		Values: values,
	}).Result()
}

// Claim ...
func (s *Stream) Claim() ([]StreamMessage, error) {
	return s.client.XClaim(&redis.XClaimArgs{
		Stream: s.name,
	}).Result()
}

// Delete ...
func (s *Stream) Delete(ids ...string) (int64, error) {
	return s.client.XDel(s.name, ids...).Result()
}

// Group ...
// TODO
func (s *Stream) Group() (int64, error) {
	return 0, nil
}

// Info ...
// TODO
func (s *Stream) Info() (int64, error) {
	return 0, nil
}

// Len ...
func (s *Stream) Len() (int64, error) {
	return s.client.XLen(s.name).Result()
}

// Pending ...
func (s *Stream) Pending(group string) (*StreamPending, error) {
	return s.client.XPending(s.name, group).Result()
}

// Range ...
func (s *Stream) Range(start, stop string) ([]StreamMessage, error) {
	return s.client.XRange(s.name, start, stop).Result()
}

// Read ...
func (s *Stream) Read() ([]StreamInfo, error) {
	return s.client.XRead(&redis.XReadArgs{
		Streams: []string{s.name},
	}).Result()
}

// ReadGroup ...
func (s *Stream) ReadGroup() ([]StreamInfo, error) {
	return s.client.XReadGroup(&redis.XReadGroupArgs{
		Streams: []string{s.name},
	}).Result()
}

// ReverseRange ...
func (s *Stream) ReverseRange(start, stop string) ([]StreamMessage, error) {
	return s.client.XRevRange(s.name, start, stop).Result()
}

// Trim ...
func (s *Stream) Trim(maxLen int64) (int64, error) {
	return s.client.XTrim(s.name, maxLen).Result()
}
