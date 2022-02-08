package redis

import "context"

type Strings struct {
	c *Client
}

func NewStrings(client *Client) Strings {
	return Strings{c: client}
}

// Get TODO
// See: https://redis.io/commands/get
func (s Strings) Get(ctx context.Context, key string) (string, error) {
	req := newRequest("*2\r\n$3\r\nGET\r\n$")
	req.addString(key)
	return s.c.cmdString(ctx, req)
}

// Set key to hold the string value.
// If key already holds a value, it is overwritten, regardless of its type.
// Any previous time to live associated with the key is discarded on successful SET operation.
//
// See: https://redis.io/commands/set
func (s Strings) Set(ctx context.Context, key, value string) error {
	req := newRequest("*3\r\n$3\r\nSET\r\n$")
	req.addString2(key, value)
	_, err := s.c.cmdString(ctx, req)
	return err
}

/*
APPEND
DECR
DECRBY
GET
GETDEL
GETEX
GETRANGE
GETSET
INCR
INCRBY
INCRBYFLOAT
LCS
MGET
MSET
MSETNX
PSETEX
SETEX
SETNX
SETRANGE
STRLEN
SUBSTR
*/
