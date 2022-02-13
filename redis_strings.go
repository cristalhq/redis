package redis

import (
	"context"
	"errors"
	"time"
)

type Strings struct {
	c *Client
}

func NewStrings(client *Client) Strings {
	return Strings{c: client}
}

// Append this command appends the value at the end of the string.
// See: https://redis.io/commands/append
func (str Strings) Append(ctx context.Context, key, value string) (int64, error) {
	req := newRequest("*3\r\n$6\r\nAPPEND\r\n$")
	req.addString2(key, value)
	return str.c.cmdInt(ctx, req)
}

// Dec decrements the number stored at key by one.
// See: https://redis.io/commands/decr
func (str Strings) Dec(ctx context.Context, key string) (int64, error) {
	req := newRequest("*2\r\n$4\r\nDECR\r\n$")
	req.addString(key)
	return str.c.cmdInt(ctx, req)
}

// DecBy decrements the number stored at key by delta.
// See: https://redis.io/commands/decrby
func (str Strings) DecBy(ctx context.Context, key string, delta int64) (int64, error) {
	req := newRequest("*3\r\n$4\r\nDECRBY\r\n$")
	req.addStringInt(key, delta)
	return str.c.cmdInt(ctx, req)
}

// Get the value of key.
// See: https://redis.io/commands/GET
func (str Strings) Get(ctx context.Context, key string) (string, error) {
	req := newRequest("*2\r\n$3\r\nGET\r\n$")
	req.addString(key)
	return str.c.cmdString(ctx, req)
}

// GetDel get the value of key and delete the key.
// See: https://redis.io/commands/GETDEL
func (str Strings) GetDel(ctx context.Context, key string) (string, error) {
	req := newRequest("*2\r\n$6\r\nGETDEL\r\n$")
	req.addString(key)
	return str.c.cmdString(ctx, req)
}

// GetExpire TODO
// See: https://redis.io/commands/getex
func (str Strings) GetExpire(ctx context.Context, d time.Duration, key string) (string, error) {
	req := newRequest("*4\r\n$5\r\nGETEX\r\n$2\r\nPX\r\n$")
	req.addStringInt(key, int64(d.Milliseconds()))
	return str.c.cmdString(ctx, req)
}

// GetExpireAt TODO
// See: https://redis.io/commands/getex
func (str Strings) GetExpireAt(ctx context.Context, key string) (string, error) {
	req := newRequest("*3\r\n$5\r\nGETEX\r\n$7\r\nPERSIST\r\n$")
	req.addString(key)
	return str.c.cmdString(ctx, req)
}

// GetPersist remove the time to live associated with the key.
// See: https://redis.io/commands/getex
func (str Strings) GetPersist(ctx context.Context, t time.Time, key string) (string, error) {
	req := newRequest("*4\r\n$5\r\nGETEX\r\n$4\r\nPXAT\r\n$")
	req.addStringInt(key, int64(t.UnixMilli()))
	return str.c.cmdString(ctx, req)
}

// GetRange returns the substring of the string value stored at key, determined by the offsets start and end (both are inclusive).
// See: https://redis.io/commands/getrange
func (str Strings) GetRange(ctx context.Context, key string, start, end int64) (string, error) {
	req := newRequest("*4\r\n$7\r\nGETRANGE\r\n$")
	req.addStringInt2(key, start, end)
	return str.c.cmdString(ctx, req)
}

// GetSet atomically sets key to value and returns the old value stored at key.
// See: https://redis.io/commands/GETSET
func (str Strings) GetSet(ctx context.Context, key, value string) (string, error) {
	req := newRequest("*3\r\n$6\r\nGETSET\r\n$")
	req.addString2(key, value)
	return str.c.cmdString(ctx, req)
}

// Inc increments the number stored at key by one.
// See: https://redis.io/commands/incr
func (str Strings) Inc(ctx context.Context, key string) (int64, error) {
	req := newRequest("*2\r\n$4\r\nINCR\r\n$")
	req.addString(key)
	return str.c.cmdInt(ctx, req)
}

// IncBy increments the number stored at key by delta.
// See: https://redis.io/commands/incrby
func (str Strings) IncBy(ctx context.Context, key string, delta int64) (int64, error) {
	req := newRequest("*3\r\n$4\r\nINCRBY\r\n$")
	req.addStringInt(key, delta)
	return str.c.cmdInt(ctx, req)
}

// IncByFloat increments the number stored at key by delta.
// See: https://redis.io/commands/incrbyfloat
func (str Strings) IncByFloat(ctx context.Context, key string, delta float64) (float64, error) {
	req := newRequest("*4\r\n$12\r\nHINCRBYFLOAT\r\n$")
	// TODO(oleg): must encode as float
	req.addStringInt(key, int64(delta))
	return str.c.cmdFloat(ctx, req)
}

// LCS command implements the longest common subsequence algorithm.
// See: https://redis.io/commands/LCS
func (str Strings) LCS(ctx context.Context, key1, key2 string) error {
	panic("redis: Strings.LCS not implemented")
}

// MultiGet returns the values of all specified keys.
// See: https://redis.io/commands/MGET
func (str Strings) MultiGet(ctx context.Context, keys ...string) ([]string, error) {
	req := newRequestSize(1+len(keys), "\r\n$4\r\nMGET")
	req.addStrings(keys)
	return str.c.cmdStrings(ctx, req)
}

// MultiSet sets the given keys to their respective values
// See: https://redis.io/commands/MSET
func (str Strings) MultiSet(ctx context.Context, pairs ...string) error {
	if len(pairs)%2 == 1 {
		return errors.New("one of the keys does not have a value")
	}
	req := newRequestSize(1+len(pairs), "\r\n$4\r\nMSET")
	req.addStrings(pairs)
	_, err := str.c.cmdString(ctx, req)
	return err
}

// MultiSetNotExist sets the given keys to their respective values.
// See: https://redis.io/commands/msetnx
func (str Strings) MultiSetNotExist(ctx context.Context, pairs ...string) (int64, error) {
	if len(pairs)%2 == 1 {
		return 0, errors.New("one of the keys does not have a value")
	}
	req := newRequestSize(1+len(pairs), "\r\n$6\r\nMSETNX")
	req.addStrings(pairs)
	return str.c.cmdInt(ctx, req)
}

// Set key to hold the string value.
// If key already holds a value, it is overwritten, regardless of its type.
// Any previous time to live associated with the key is discarded on successful SET operation.
//
// See: https://redis.io/commands/set
func (str Strings) Set(ctx context.Context, key, value string) error {
	req := newRequest("*3\r\n$3\r\nSET\r\n$")
	req.addString2(key, value)
	_, err := str.c.cmdString(ctx, req)
	return err
}

// SetExpire key to hold the string value and set key to timeout after a given number of seconds.
// See: https://redis.io/commands/SETEX
func (str Strings) SetExpire(ctx context.Context, d time.Duration, key, value string) error {
	req := newRequest("*4\r\n$6\r\nPSETEX\r\n$")
	req.addStringIntString(key, int64(d.Milliseconds()), value)
	_, err := str.c.cmdString(ctx, req)
	return err
}

// SetNotExist key to hold string value if key does not exist.
// In that case, it is equal to SET. When key already holds a value, no operation is performed.
// See: https://redis.io/commands/SETNX
func (str Strings) SetNotExist(ctx context.Context, key, value string) error {
	req := newRequest("*3\r\n$5\r\nSETNX\r\n$")
	req.addString2(key, value)
	_, err := str.c.cmdString(ctx, req)
	return err
}

// SetRange overwrites part of the string stored at key, starting at the specified offset, for the entire length of value.
// See: https://redis.io/commands/setrange
func (str Strings) SetRange(ctx context.Context, key string, offset int64, value string) (int64, error) {
	req := newRequest("*4\r\n$8\r\nSETRANGE\r\n$")
	req.addStringIntString(key, offset, value)
	return str.c.cmdInt(ctx, req)
}

// Strlen returns the length of the string value stored at key.
// See: https://redis.io/commands/strlen
func (str Strings) Strlen(ctx context.Context, key string) (int64, error) {
	req := newRequest("*2\r\n$6\r\nSTRLEN\r\n$")
	req.addString(key)
	return str.c.cmdInt(ctx, req)
}

// Substr returns the substring of the string value stored at key.
// See: https://redis.io/commands/substr
func (str Strings) Substr(ctx context.Context, key string, start, end int64) (string, error) {
	req := newRequest("*4\r\n$6\r\nSUBSTR\r\n$")
	req.addStringInt2(key, start, end)
	return str.c.cmdString(ctx, req)
}
