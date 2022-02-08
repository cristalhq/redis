package redis

import (
	"context"
	"fmt"
)

// Stream TODO
// See: https://redis.io/topics/streams-intro
type Stream struct {
	name string
	c    *Client
}

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

// NewStream instantiates a new Stream structure client for Redis.
func NewStream(name string, client *Client) *Stream {
	return &Stream{name: name, c: client}
}

// Name of the Stream structure.
func (st *Stream) Name() string { return st.name }

// Ack ...
// See: https://redis.io/commands/xack
func (st *Stream) Ack(ctx context.Context, group string, ids ...string) (int64, error) {
	req := newRequestSize(3+len(ids), "\r\n$4\r\nXACK\r\n$")
	req.addString2(st.name, group)
	req.addStrings(ids)
	return st.c.cmdInt(ctx, req)
}

// Add ...
// See: https://redis.io/commands/xadd
func (st *Stream) Add(ctx context.Context, values map[string]Value) (string, error) {
	// "*5\r\n$4\r\nXADD\r\n$8\r\nmystream\r\n$1\r\n*\r\n$1\r\na\r\n$1\r\n1\r\n"
	vals := make([]string, 0, len(values)*2)
	for k, v := range values {
		vals = append(vals, k, v)
	}

	req := newRequestSize(3+2*len(values), "\r\n$4\r\nXADD\r\n$")
	req.addString2AndStrings(st.name, "*", vals)
	return st.c.cmdString(ctx, req)
	// panic("")
	// return st.client.XAdd(ctx, &redis.XAddArgs{
	// 	Stream: st.name,
	// 	Values: values,
	// }).Result()
}

// AutoClaim ...
// See: https://redis.io/commands/xautoclaim
func (st *Stream) AutoClaim(ctx context.Context) (next string, msg StreamMessage, _ error) {
	req := newRequest("*2\r\n$10\r\nXAUTOCLAIM\r\n$")
	req.addString(st.name)
	ss, err := st.c.cmdStrings(ctx, req)
	if err != nil {
		_ = ss
	}
	return "", StreamMessage{}, nil
}

// Claim ...
func (st *Stream) Claim(ctx context.Context) ([]StreamMessage, error) {
	panic("")
	// return st.client.XClaim(ctx, &redis.XClaimArgs{
	// 	Stream: st.name,
	// }).Result()
}

// Claim ...
func (st *Stream) ClaimIDs(ctx context.Context) ([]string, error) {
	req := newRequest("*1\r\n$10\r\nXAUTOCLAIM\r\n$")
	// TODO: +JUSTID
	return st.c.cmdStrings(ctx, req)
}

// Delete the specified entries from a stream, and returns the number of entries deleted.
// See: https://redis.io/commands/xdel
func (st *Stream) Delete(ctx context.Context, ids ...string) (int64, error) {
	req := newRequestSize(2+len(ids), "\r\n$4\r\nXDEL\r\n$")
	req.addStringAndStrings(st.name, ids)
	return st.c.cmdInt(ctx, req)
}

// Group ...
// XGROUP CREATE
// XGROUP CREATECONSUMER
// XGROUP DELCONSUMER
// XGROUP DESTROY
// XGROUP SETID
func (st *Stream) Group(ctx context.Context) (int64, error) {
	return 0, nil
}

// Info ...
// XINFO CONSUMERS
// XINFO GROUPS
// XINFO STREAM
func (st *Stream) Info(ctx context.Context) (int64, error) {
	return 0, nil
}

// Len кeturns the number of entries inside a stream.
// See: https://redis.io/commands/xlen
func (st *Stream) Len(ctx context.Context) (int64, error) {
	req := newRequest("*2\r\n$4\r\nXLEN\r\n$")
	req.addString(st.name)
	return st.c.cmdInt(ctx, req)
}

// Pending ...
func (st *Stream) Pending(ctx context.Context, group string) (*StreamPending, error) {
	panic("") //return st.client.XPending(ctx, st.name, group).Result()
}

// Get the item by id from stream.
// This operation is not directly supported by Redis, it is a wrapper for XRANGE.
// See: https://redis.io/commands/xrange#fetching-single-items
func (st *Stream) Get(ctx context.Context, id string) (StreamMessage, error) {
	req := newRequest("*4\r\n$6\r\nXRANGE\r\n$")
	req.addString3(st.name, id, id)
	ss, err := st.c.cmdStrings(ctx, req)
	_ = ss[0]
	return StreamMessage{}, err
}

// Range ...
// TODO: add start/end validation
func (st *Stream) Range(ctx context.Context, start, end string, count int64) ([]StreamMessage, error) {
	req := newRequest("*5\r\n$6\r\nXRANGE\r\n$")
	req.addString3(st.name, start, end)
	req.buf = append(req.buf, '$')
	req.addStringInt("COUNT", count)
	fmt.Printf("req: %s\n", string(req.buf))
	ss, err := st.c.cmdStrings(ctx, req)
	_ = ss
	return nil, err
}

// Range ...
func (st *Stream) RangeAll(ctx context.Context, start, end string) ([]StreamMessage, error) {
	req := newRequest("*4\r\n$6\r\nXRANGE\r\n$")
	req.addString3(st.name, start, end)
	ss, err := st.c.cmdStrings(ctx, req)
	_ = ss
	return nil, err
}

// Read ...
func (st *Stream) Read(ctx context.Context) ([]StreamInfo, error) {
	panic("") //return st.client.XRead(ctx, &redis.XReadArgs{Streams: []string{st.name}}).Result()
}

// ReadGroup ...
func (st *Stream) ReadGroup(ctx context.Context) ([]StreamInfo, error) {
	panic("") //return st.client.XReadGroup(ctx, &redis.XReadGroupArgs{Streams: []string{st.name}}).Result()
}

// ReverseRange ...
func (st *Stream) ReverseRange(ctx context.Context, end, start string) ([]StreamMessage, error) {
	panic("") //return st.client.XRevRange(ctx, st.name, start, end).Result()
}

// ReverseRange ...
// https://redis.io/commands/xrevrange
func (st *Stream) ReverseRangeAll(ctx context.Context, end, start string) ([]StreamMessage, error) {
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
func (st *Stream) Trim(ctx context.Context, maxLen int64) (int64, error) {
	panic("") //return st.client.XRevRange(ctx, st.name, start, end).Result()
	// req := newRequest("*2\r\n$5\r\nXTRIM\r\n$")
	// req.addString(st.name)
	// return st.c.cmdInt(ctx, req)
	// return st.client.XTrim(ctx, st.name, maxLen).Result()
}
