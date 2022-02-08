package redis

import "context"

// PubSub TODO
// See: https://redis.io/topics/pubsub
type PubSub struct {
	name string
	c    *Client
}

func NewPubSub(name string) *PubSub {
	return &PubSub{}
}

func (ps *PubSub) PSUBSCRIBE(ctx context.Context) error { return nil }

// PUBLISH TODO
// See: https://redis.io/commands/publish
func (ps *PubSub) PUBLISH(ctx context.Context, msg string) (int64, error) {
	req := newRequest("*3\r\n$7\r\nPUBLISH\r\n$")
	req.addString2(ps.name, msg)
	return ps.c.cmdInt(ctx, req)
}

// CHANNELS TODO
// See: https://redis.io/commands/pubsub-channels
func (ps *PubSub) CHANNELS(ctx context.Context, pattern string) ([]string, error) {
	var req *request
	if pattern == "" {
		req = newRequest("*1\r\n$8\r\nCHANNELS\r\n$")
	} else {
		req = newRequest("*2\r\n$8\r\nCHANNELS\r\n$")
		req.addString(pattern)
	}
	return ps.c.cmdStrings(ctx, req)
}

func (ps *PubSub) PUBSUB_NUMPAT(ctx context.Context) error        { return nil }
func (ps *PubSub) PUBSUB_NUMSUB(ctx context.Context) error        { return nil }
func (ps *PubSub) PUBSUB_SHARDCHANNELS(ctx context.Context) error { return nil }
func (ps *PubSub) PUNSUBSCRIBE(ctx context.Context) error         { return nil }
func (ps *PubSub) SPUBLISH(ctx context.Context) error             { return nil }
func (ps *PubSub) SSUBSCRIBE(ctx context.Context) error           { return nil }
func (ps *PubSub) SUBSCRIBE(ctx context.Context) error            { return nil }
func (ps *PubSub) SUNSUBSCRIBE(ctx context.Context) error         { return nil }
func (ps *PubSub) UNSUBSCRIBE(ctx context.Context) error          { return nil }

type PubSubListener struct{}

func (psl *PubSubListener) SUBSCRIBE(ctx context.Context) error    { return nil }
func (psl *PubSubListener) SSUBSCRIBE(ctx context.Context) error   { return nil }
func (psl *PubSubListener) PSUBSCRIBE(ctx context.Context) error   { return nil }
func (psl *PubSubListener) UNSUBSCRIBE(ctx context.Context) error  { return nil }
func (psl *PubSubListener) SUNSUBSCRIBE(ctx context.Context) error { return nil }
func (psl *PubSubListener) PUNSUBSCRIBE(ctx context.Context) error { return nil }
func (psl *PubSubListener) PING(ctx context.Context) error         { return nil }
func (psl *PubSubListener) RESET(ctx context.Context) error        { return nil }
func (psl *PubSubListener) QUIT(ctx context.Context) error         { return nil }
