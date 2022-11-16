package redis

import (
	"context"
	"fmt"
)

// Stream client for stream operations in Redis.
// See: https://redis.io/commands#stream
//
// Redis Streams: https://redis.io/topics/streams-intro
type Stream struct {
	name string
	c    *Client
}

// NewStream returns new Redis Stream client.
func NewStream(name string, client *Client) Stream {
	return Stream{name: name, c: client}
}

// StreamMessage ..
type StreamMessage struct {
	ID     string
	Values map[string]interface{}
}

// StreamInfo ...
type StreamInfo struct {
	Stream   string
	Messages []StreamMessage
}

// StreamPending ...
type StreamPending struct {
	Count     int64
	Lower     string
	Higher    string
	Consumers map[string]int64
}

// Name of the Stream structure.
func (st Stream) Name() string { return st.name }

// Ack ...
// See: https://redis.io/commands/xack
func (st Stream) Ack(ctx context.Context, group string, ids ...string) (int64, error) {
	req := newRequestSize(3+len(ids), "\r\n$4\r\nXACK\r\n$")
	req.addString2(st.name, group)
	req.addStrings(ids)
	return st.c.cmdInt(ctx, req)
}

// Add appends the specified stream entry to the stream.
// See: https://redis.io/commands/xadd
func (st Stream) Add(ctx context.Context, values map[string]Value) (string, error) {
	vals := make([]string, 0, 2*len(values))
	for k, v := range values {
		vals = append(vals, k, v)
	}
	req := newRequestSize(3+2*len(values), "\r\n$4\r\nXADD\r\n$")
	req.addString2AndStrings(st.name, "*", vals)
	return st.c.cmdString(ctx, req)
}

// AutoClaim ...
// See: https://redis.io/commands/xautoclaim
func (st Stream) AutoClaim(ctx context.Context) (next string, msg StreamMessage, _ error) {
	req := newRequest("*2\r\n$10\r\nXAUTOCLAIM\r\n$")
	req.addString(st.name)
	ss, err := st.c.cmdStrings(ctx, req)
	if err != nil {
		_ = ss
	}
	return "", StreamMessage{}, nil
}

// Claim ...
// XCLAIM key group consumer min-idle-time id [id ...] [IDLE ms] [TIME unix-time-milliseconds] [RETRYCOUNT count] [FORCE] [JUSTID]
func (st Stream) Claim(ctx context.Context) ([]StreamMessage, error) {
	panic("")
	// return st.client.XClaim(ctx, &redis.XClaimArgs{
	// 	Stream: st.name,
	// }).Result()
}

// Claim ...
func (st Stream) ClaimIDs(ctx context.Context) ([]string, error) {
	req := newRequest("*1\r\n$10\r\nXAUTOCLAIM\r\n$")
	// TODO: +JUSTID
	return st.c.cmdStrings(ctx, req)
}

// Delete the specified entries from a stream, and returns the number of entries deleted.
// See: https://redis.io/commands/xdel
func (st Stream) Delete(ctx context.Context, ids ...string) (int64, error) {
	req := newRequestSize(2+len(ids), "\r\n$4\r\nXDEL\r\n$")
	req.addStringAndStrings(st.name, ids)
	return st.c.cmdInt(ctx, req)
}

// GroupCreate command creates a new consumer group uniquely identified by groupname for the stream.
// See: https://redis.io/commands/xgroup-create
func (st Stream) GroupCreate(ctx context.Context, group, id string) error {
	// TODO(oleg): support [MKSTREAM]
	req := newRequest("*6\r\n$6\r\nXGROUP\r\n$6\r\nCREATE\r\n$")
	req.addString4(st.name, group, id, "MKSTREAM")
	return st.c.cmdSimple(ctx, req)
}

// GroupCreateConsumer create a consumer in the consumer group of the stream.
// Returns the number of destroyed consumer groups (0 or 1).
//
// See: https://redis.io/commands/xgroup-createconsumer
func (st Stream) GroupCreateConsumer(ctx context.Context, group, consumer string) (int, error) {
	req := newRequest("*5\r\n$6\r\nXGROUP\r\n$14\r\nCREATECONSUMER\r\n$")
	req.addString3(st.name, group, consumer)
	i, err := st.c.cmdInt(ctx, req)
	return int(i), err
}

// GroupDeleteConsumer deletes a consumer from the consumer group.
// Returns the number of pending messages that the consumer had before it was deleted.
//
// See: https://redis.io/commands/xgroup-delconsumer
func (st Stream) GroupDeleteConsumer(ctx context.Context, group, consumer string) (int, error) {
	req := newRequest("*5\r\n$6\r\nXGROUP\r\n$11\r\nDELCONSUMER\r\n$")
	req.addString3(st.name, group, consumer)
	i, err := st.c.cmdInt(ctx, req)
	return int(i), err
}

// GroupDestroy completely destroys a consumer group.
// Returns the number of destroyed consumer groups (0 or 1).
//
// See: https://redis.io/commands/xgroup-destroy
func (st Stream) GroupDestroy(ctx context.Context, group string) (int, error) {
	req := newRequest("*4\r\n$6\r\nXGROUP\r\n$7\r\nDESTROY\r\n$")
	req.addString2(st.name, group)
	i, err := st.c.cmdInt(ctx, req)
	return int(i), err
}

// GroupSetID set the last delivered ID for a consumer group.
// See: https://redis.io/commands/xgroup-setid
func (st Stream) GroupSetID(ctx context.Context, group, id string) error {
	req := newRequest("*5\r\n$6\r\nXGROUP\r\n$5\r\nSETID\r\n$")
	req.addString3(st.name, group, id)
	return st.c.cmdSimple(ctx, req)
}

// InfoConsumers TODO
// See: https://redis.io/commands/xinfo-consumers
func (st Stream) InfoConsumers(ctx context.Context, group string) error {
	// XINFO CONSUMERS key groupname
	return nil
}

// InfoGroups TODO
// See: https://redis.io/commands/xinfo-groups
func (st Stream) InfoGroups(ctx context.Context) error {
	// XINFO GROUPS key
	return nil
}

// InfoStream TODO
// See: https://redis.io/commands/xinfo-stream
func (st Stream) InfoStream(ctx context.Context) error {
	// XINFO STREAM key [FULL [COUNT count]]
	return nil
}

// Len кeturns the number of entries inside a stream.
// See: https://redis.io/commands/xlen
func (st Stream) Len(ctx context.Context) (int64, error) {
	req := newRequest("*2\r\n$4\r\nXLEN\r\n$")
	req.addString(st.name)
	return st.c.cmdInt(ctx, req)
}

// Pending ...
func (st Stream) Pending(ctx context.Context, group string) (StreamPending, error) {
	panic("") //return st.client.XPending(ctx, st.name, group).Result()
}

// Get the item by id from stream.
// This operation is not directly supported by Redis, it is a wrapper for XRANGE.
// See: https://redis.io/commands/xrange#fetching-single-items
func (st Stream) Get(ctx context.Context, id string) (StreamMessage, error) {
	req := newRequest("*4\r\n$6\r\nXRANGE\r\n$")
	req.addString3(st.name, id, id)
	ss, err := st.c.cmdStrings(ctx, req)
	_ = ss[0]
	return StreamMessage{}, err
}

// Range returns the stream entries matching a given range of IDs.
// See: https://redis.io/commands/xrange
func (st Stream) Range(ctx context.Context, start, end string, count int64) ([]StreamMessage, error) {
	// TODO: add start/end validation
	req := newRequest("*5\r\n$6\r\nXRANGE\r\n$")
	req.addString3(st.name, start, end)
	req.buf = append(req.buf, '$')
	req.addStringInt("COUNT", count)
	fmt.Printf("req: %s\n", string(req.buf))
	ss, err := st.c.cmdStrings(ctx, req)
	res := make([]StreamMessage, len(ss))
	for i := range ss {
		res[i] = StreamMessage{}
	}
	return res, err
}

// Range ...
func (st Stream) RangeAll(ctx context.Context, start, end string) ([]StreamMessage, error) {
	req := newRequest("*4\r\n$6\r\nXRANGE\r\n$")
	req.addString3(st.name, start, end)
	ss, err := st.c.cmdStrings(ctx, req)
	res := make([]StreamMessage, len(ss))
	for i := range ss {
		res[i] = StreamMessage{}
	}
	return res, err
}

// Read ...
func (st Stream) Read(ctx context.Context) ([]StreamInfo, error) {
	panic("") //return st.client.XRead(ctx, &redis.XReadArgs{Streams: []string{st.name}}).Result()
}

// ReadGroup ...
func (st Stream) ReadGroup(ctx context.Context) ([]StreamInfo, error) {
	panic("") //return st.client.XReadGroup(ctx, &redis.XReadGroupArgs{Streams: []string{st.name}}).Result()
}

// ReverseRange ...
func (st Stream) ReverseRange(ctx context.Context, end, start string) ([]StreamMessage, error) {
	panic("") //return st.client.XRevRange(ctx, st.name, start, end).Result()
}

// ReverseRange ...
// https://redis.io/commands/xrevrange
func (st Stream) ReverseRangeAll(ctx context.Context, end, start string) ([]StreamMessage, error) {
	req := newRequest("*4\r\n$9\r\nXREVRANGE\r\n$")
	req.addString3(st.name, end, start)
	ss, err := st.c.cmdStrings(ctx, req)
	_ = ss
	return nil, err
}

// Note: The XSETID command is an internal command.
// It is used by a Redis master to replicate the last delivered ID of streams.
// https://redis.io/commands/xsetid

// Trim the stream by evicting older entries (entries with lower IDs) if needed.
// See: https://redis.io/commands/xtrim
func (st Stream) Trim(ctx context.Context, maxLen int64) (int64, error) {
	panic("") //return st.client.XRevRange(ctx, st.name, start, end).Result()
	// req := newRequest("*2\r\n$5\r\nXTRIM\r\n$")
	// req.addString(st.name)
	// return st.c.cmdInt(ctx, req)
	// return st.client.XTrim(ctx, st.name, maxLen).Result()
}
