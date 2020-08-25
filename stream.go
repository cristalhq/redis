package redis

import (
	"github.com/go-redis/redis"
)

type Stream struct {
	name   string
	client *redis.Client
}

type (
	StreamInfo    = redis.XStream
	StreamMessage = redis.XMessage
	StreamPending = redis.XPending
)

// NewStream instantiates a new Stream structure client for Redis.
func NewStream(name string, client *redis.Client) *Stream {
	return &Stream{name: name, client: client}
}

func (s *Stream) Ack(group string, ids ...string) (int64, error) {
	return s.client.XAck(s.name, group, ids...).Result()
}

func (s *Stream) Add(values map[string]interface{}) (string, error) {
	return s.client.XAdd(&redis.XAddArgs{
		Stream: s.name,
		Values: values,
	}).Result()
}

func (s *Stream) Claim() ([]StreamMessage, error) {
	return s.client.XClaim(&redis.XClaimArgs{
		Stream: s.name,
	}).Result()
}

func (s *Stream) Delete(ids ...string) (int64, error) {
	return s.client.XDel(s.name, ids...).Result()
}

// TODO
func (s *Stream) Group() (int64, error) {
	return 0, nil
}

// TODO
func (s *Stream) INFO() (int64, error) {
	return 0, nil
}

func (s *Stream) Len() (int64, error) {
	return s.client.XLen(s.name).Result()
}

func (s *Stream) Pending(group string) (*StreamPending, error) {
	return s.client.XPending(s.name, group).Result()
}

func (s *Stream) Range(start, stop string) ([]StreamMessage, error) {
	return s.client.XRange(s.name, start, stop).Result()
}

func (s *Stream) Read() ([]StreamInfo, error) {
	return s.client.XRead(&redis.XReadArgs{
		Streams: []string{s.name},
	}).Result()
}

func (s *Stream) ReadGroup() ([]StreamInfo, error) {
	return s.client.XReadGroup(&redis.XReadGroupArgs{
		Streams: []string{s.name},
	}).Result()
}

func (s *Stream) ReverseRange(start, stop string) ([]StreamMessage, error) {
	return s.client.XRevRange(s.name, start, stop).Result()
}

func (s *Stream) Trim(maxLen int64) (int64, error) {
	return s.client.XTrim(s.name, maxLen).Result()
}
