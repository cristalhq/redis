package redis

import (
	"context"
	"fmt"
	"time"
)

type Clients struct {
	c *Client
}

func NewClients(client *Client) Clients {
	return Clients{c: client}
}

// Auth command authenticates the current connection in two cases.
// See: https://redis.io/commands/auth
func (c Clients) Auth(ctx context.Context, username, password string) error {
	var req *request
	if username == "" {
		req = newRequest("*2\r\n$4\r\nAUTH\r\n$")
		req.addString(password)
	} else {
		req = newRequest("*3\r\n$4\r\nAUTH\r\n$")
		req.addString2(username, password)
	}
	return c.c.cmdSimple(ctx, req)
}

func (c Clients) ClientCaching(ctx context.Context) error {
	return nil
}

// ClientGetName returns the name of the current connection as set by ClientSetName.
// See: https://redis.io/commands/client-getname
func (c Clients) ClientGetName(ctx context.Context) (string, error) {
	req := newRequest("*2\r\n$6\r\nCLIENT\r\n$7\r\nGETNAME\r\n")
	return c.c.cmdString(ctx, req)
}

func (c Clients) ClientGetRedir(ctx context.Context) error {
	return nil
}

// ClientID returns the ID of the current connection.
// See: https://redis.io/commands/client-id
func (c Clients) ClientID(ctx context.Context) (int64, error) {
	req := newRequest("*2\r\n$6\r\nCLIENT\r\n$2\r\nID\r\n$")
	return c.c.cmdInt(ctx, req)
}

// See: https://redis.io/commands/client-info
func (c Clients) ClientInfo(ctx context.Context) (string, error) {
	req := newRequest("*2\r\n$6\r\nCLIENT\r\n$4\r\nINFO\r\n$")
	return c.c.cmdString(ctx, req)
}

// See: https://redis.io/commands/client-kill
func (c Clients) ClientKill(ctx context.Context) error {
	return nil
}

// ClientList returns information and statistics about the client connections server.
// See: https://redis.io/commands/client-list
func (c Clients) ClientList(ctx context.Context) ([]string, error) {
	// CLIENT LIST [TYPE NORMAL|MASTER|REPLICA|PUBSUB] [ID client-id [client-id ...]]
	return nil, nil
}

// ClientNoEvict sets the client eviction mode for the current connection.
// See: https://redis.io/commands/client-no-evict
func (c Clients) ClientNoEvict(ctx context.Context, mode string) error {
	if mode != "ON" && mode != "OFF" {
		return fmt.Errorf("unknown ClientNoEvict mode: %s", mode)
	}
	req := newRequest("*4\r\n$6\r\nCLIENT\r\n$8\r\nNO-EVICT\r\n$")
	req.addString(mode)
	return c.c.cmdSimple(ctx, req)
}

// ClientPause is a connections control command able to suspend all the Redis clients for the specified amount of time.
// See: https://redis.io/commands/client-pause
func (c Clients) ClientPause(ctx context.Context, d time.Duration, mode string) error {
	if mode != "WRITE" && mode != "ALL" {
		return fmt.Errorf("unknown ClientPause mode: %s", mode)
	}
	req := newRequest("*4\r\n$6\r\nCLIENT\r\n$5\r\nPAUSE\r\n$")
	req.addIntString(d.Milliseconds(), mode)
	return c.c.cmdSimple(ctx, req)
}

// ClientReply
// See: https://redis.io/commands/client-reply
func (c Clients) ClientReply(ctx context.Context, mode string) error {
	if mode != "ON" && mode != "OFF" && mode != "SKIP" {
		return fmt.Errorf("unknown ClientReply mode: %s", mode)
	}
	req := newRequest("*3\r\n$6\r\nCLIENT\r\n$5\r\nREPLY\r\n$")
	req.addString(mode)
	return c.c.cmdSimple(ctx, req)
}

// ClientSetName TODO
// See: https://redis.io/commands/client-setname
func (c Clients) ClientSetName(ctx context.Context, name string) error {
	req := newRequest("*3\r\n$6\r\nCLIENT\r\n$7\r\nSETNAME\r\n$")
	req.addString(name)
	return c.c.cmdSimple(ctx, req)
}

// ClientTracking TODO
// See: https://redis.io/commands/client-tracking
func (c Clients) ClientTracking(ctx context.Context) error {
	// CLIENT TRACKING ON|OFF [REDIRECT client-id] [PREFIX prefix [PREFIX prefix ...]] [BCAST] [OPTIN] [OPTOUT] [NOLOOP]
	req := newRequest("*2\r\n$6\r\nCLIENT\r\n$8\r\nTRACKING\r\n$")
	return c.c.cmdSimple(ctx, req)
}

// ClientTrackingInfo command returns information about the current client connection's use of the server assisted client side caching feature.
// See: https://redis.io/commands/client-trackinginfo
func (c Clients) ClientTrackingInfo(ctx context.Context) ([]string, error) {
	req := newRequest("*2\r\n$6\r\nCLIENT\r\n$12\r\nTRACKINGINFO\r\n$")
	return c.c.cmdStrings(ctx, req)
}

// ClientUnblock TODO
// See: https://redis.io/commands/client-unblock
func (c Clients) ClientUnblock(ctx context.Context, id string) error {
	// CLIENT  client-id [TIMEOUT|ERROR]
	req := newRequest("*3\r\n$6\r\nCLIENT\r\n$7\r\nUNBLOCK\r\n$")
	req.addString(id)
	return c.c.cmdSimple(ctx, req)
}

// ClientUnpauses used to resume command processing for all clients that were paused by ClientPause.
// See: https://redis.io/commands/client-unpause
func (c Clients) ClientUnpause(ctx context.Context) error {
	req := newRequest("*2\r\n$6\r\nCLIENT\r\n$7\r\nUNPAUSE\r\n$")
	return c.c.cmdSimple(ctx, req)
}

// Echo returns message.
// See: https://redis.io/commands/echo
func (c Clients) Echo(ctx context.Context, msg string) (string, error) {
	req := newRequest("*2\r\n$4\r\nECHO\r\n$")
	req.addString(msg)
	return c.c.cmdString(ctx, req)
}

// Hello TODO
// https://redis.io/commands/hello
func (c Clients) Hello(ctx context.Context) error {
	// HELLO [protover [AUTH username password] [SETNAME clientname]]
	return nil
}

// Ping returns PONG if no argument is provided, otherwise return a copy of the argument
// See: https://redis.io/commands/ping
func (c Clients) Ping(ctx context.Context, msg string) (string, error) {
	var req *request
	if msg == "" {
		req = newRequest("*1\r\n$4\r\nPING\r\n$")
	} else {
		req = newRequest("*2\r\n$4\r\nPING\r\n$")
		req.addString(msg)
	}
	return c.c.cmdString(ctx, req)
}

// Quit asks the server to close the connection.
// See: https://redis.io/commands/quit
func (c Clients) Quit(ctx context.Context) error {
	req := newRequest("*1\r\n$4\r\nQUIT\r\n$")
	return c.c.cmdSimple(ctx, req)
}

// This command performs a full reset of the connection's server-side context, mimicking the effect of disconnecting and reconnecting again.
// See: https://redis.io/commands/reset
func (c Clients) Reset(ctx context.Context) error {
	req := newRequest("*1\r\n$5\r\nRESET\r\n$")
	return c.c.cmdSimple(ctx, req)
}

// Select the Redis logical database having the specified zero-based numeric index
// See: https://redis.io/commands/select
func (c Clients) Select(ctx context.Context, db int) error {
	req := newRequest("*2\r\n$6\r\nSELECT\r\n$")
	req.int(int64(db))
	return c.c.cmdSimple(ctx, req)
}
