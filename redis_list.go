package redis

import (
	"context"
	"fmt"
)

// List client for list operations in Redis.
// See: https://redis.io/commands#list
type List struct {
	name string
	c    *Client
}

// NewList returns new Redis List client.
func NewList(name string, client *Client) List {
	return List{name: name, c: client}
}

// Name of the List structure.
func (list List) Name() string { return list.name }

// BlockingLeftMove TODO
// See: https://redis.io/commands/blmove
func (list List) BlockingLeftMove(ctx context.Context) error {
	panic("redis: List.BlockingLeftMove not implemented")
}

// BlockingLeftMultiPop TODO
// See: https://redis.io/commands/blmpop
func (list List) BlockingLeftMultiPop(ctx context.Context) error {
	panic("redis: List.BlockingLeftMultiPop not implemented")
}

// BlockingLeftPop TODO
// See: https://redis.io/commands/blpop
func (list List) BlockingLeftPop(ctx context.Context) error {
	panic("redis: List.BlockingLeftPop not implemented")
}

// BlockingRightPop TODO
// See: https://redis.io/commands/brpop
func (list List) BlockingRightPop(ctx context.Context) error {
	panic("redis: List.BlockingRightPop not implemented")
}

// BlockingRightPopLeftPush TODO
// See: https://redis.io/commands/brpoplpush
func (list List) BlockingRightPopLeftPush(ctx context.Context) error {
	panic("redis: List.BlockingRightPopLeftPush not implemented")
}

// Index returns the element at index index in the list.
// See: https://redis.io/commands/lindex
func (list List) Index(ctx context.Context, index int64) (string, error) {
	req := newRequest("*3\r\n$6\r\nLINDEX\r\n$")
	req.addStringInt(list.name, index)
	return list.c.cmdString(ctx, req)
}

// Insert element in the list either before or after the reference value pivot.
// See: https://redis.io/commands/linsert
func (list List) Insert(ctx context.Context, mode string, pivot, element string) (int64, error) {
	if mode != "BEFORE" && mode != "AFTER" {
		return 0, fmt.Errorf("unknown insert mode: %s", mode)
	}
	req := newRequest("*5\r\n$7\r\nLINSERT\r\n$")
	req.addString4(list.name, mode, pivot, element)
	return list.c.cmdInt(ctx, req)
}

// Len returns the length of the list.
// See: https://redis.io/commands/llen
func (list List) Len(ctx context.Context) (int64, error) {
	req := newRequest("*2\r\n$4\r\nLLEN\r\n$")
	req.addString(list.name)
	return list.c.cmdInt(ctx, req)
}

// LeftMove atomically returns and removes the first/last element (head/tail depending on the wherefrom argument) of the list
// and pushes the element at the first/last element (head/tail depending on the whereto argument) of the list stored at destination.
// Returns the element being popped and pushed.
//
// See: https://redis.io/commands/lmove
func (list List) LeftMove(ctx context.Context, destination, from, to string) (string, error) {
	if !((from == "LEFT" && to == "RIGHT") || (from == "RIGHT" && to == "LEFT")) {
		return "", fmt.Errorf("unknown from: %s or to: %s", from, to)
	}
	req := newRequest("*5\r\n$5\r\nLMOVE\r\n$")
	req.addString4(list.name, destination, from, to)
	return list.c.cmdString(ctx, req)
}

// LeftMultiPop TODO
// See: https://redis.io/commands/lmpop
func (list List) LeftMultiPop(ctx context.Context) error {
	panic("redis: List.LeftMultiPop not implemented")
}

// LeftPop removes and returns the first elements of the list.
// See: https://redis.io/commands/lpop
func (list List) LeftPop(ctx context.Context, count int64) ([]string, error) {
	req := newRequest("*3\r\n$4\r\nLPOP\r\n$")
	req.addStringInt(list.name, count)
	return list.c.cmdStrings(ctx, req)
}

// LeftPos TODO
// See: https://redis.io/commands/lpos
func (list List) LeftPos(ctx context.Context) error {
	panic("redis: List.LeftPos not implemented")
}

// LeftPush insert all the specified values at the head of the list.
// Returns the length of the list after the push operation.
//
// See: https://redis.io/commands/lpush
func (list List) LeftPush(ctx context.Context, elements ...string) (int64, error) {
	req := newRequestSize(2+len(elements), "\r\n$5\r\nLPUSH\r\n$")
	req.addStringAndStrings(list.name, elements)
	return list.c.cmdInt(ctx, req)
}

// LeftPushX inserts specified values at the head of the list, only if list already exists.
// In contrary to LPUSH, no operation will be performed when key does not yet exist.
// Returns the length of the list after the push operation.
//
// See: https://redis.io/commands/lpushx
func (list List) LeftPushX(ctx context.Context, elements ...string) (int64, error) {
	req := newRequestSize(2+len(elements), "\r\n$6\r\nLPUSHX\r\n$")
	req.addStringAndStrings(list.name, elements)
	return list.c.cmdInt(ctx, req)
}

// Range returns the specified elements of the list.
// See: https://redis.io/commands/lrange
func (list List) Range(ctx context.Context, start, stop int64) ([]string, error) {
	req := newRequest("*4\r\n$6\r\nLRANGE\r\n$")
	req.addStringInt2(list.name, start, stop)
	return list.c.cmdStrings(ctx, req)
}

// Remove the first count occurrences of elements equal to element from the list.
// See: https://redis.io/commands/lrem
func (list List) Remove(ctx context.Context, count int64, element string) (int64, error) {
	req := newRequest("*4\r\n$4\r\nLREM\r\n$")
	req.addStringIntString(list.name, count, element)
	return list.c.cmdInt(ctx, req)
}

// Set the list element at index to element.
// See: https://redis.io/commands/lset
func (list List) Set(ctx context.Context, index int64, element string) error {
	req := newRequest("*4\r\n$4\r\nLSET\r\n$")
	req.addStringIntString(list.name, index, element)
	_, err := list.c.cmdString(ctx, req)
	return err
}

// Trim an existing list so that it will contain only the specified range of elements specified.
// See: https://redis.io/commands/ltrim
func (list List) Trim(ctx context.Context, start, stop int64) error {
	req := newRequest("*4\r\n$5\r\nLTRIM\r\n$")
	req.addStringInt2(list.name, start, stop)
	_, err := list.c.cmdString(ctx, req)
	return err
}

// RightPop removes and returns the last elements of the list.
// See: https://redis.io/commands/rpop
func (list List) RightPop(ctx context.Context, count int64) ([]string, error) {
	req := newRequest("*3\r\n$4\r\nRPOP\r\n$")
	req.addStringInt(list.name, count)
	return list.c.cmdStrings(ctx, req)
}

// RightPopLeftPush atomically returns and removes the last element (tail) of the list stored
// and pushes the element at the first element (head) of the list stored at destination.
// Returns the element being popped and pushed.
//
// See: https://redis.io/commands/rpoplpush
func (list List) RightPopLeftPush(ctx context.Context, destination string) (string, error) {
	req := newRequest("*3\r\n$9\r\nRPOPLPUSH\r\n$")
	req.addString2(list.name, destination)
	return list.c.cmdString(ctx, req)
}

// RightPush insert all the specified values at the tail of the list.
// Returns the length of the list after the push operation.
//
// See: https://redis.io/commands/rpush
func (list List) RightPush(ctx context.Context, elements ...string) (int64, error) {
	req := newRequestSize(2+len(elements), "\r\n$5\r\nRPUSH\r\n$")
	req.addStringAndStrings(list.name, elements)
	return list.c.cmdInt(ctx, req)
}

// RightPushX inserts specified values at the tail of the list, only if list already exists.
// In contrary to RPUSH, no operation will be performed when key does not yet exist.
// Returns the length of the list after the push operation.
//
// See: https://redis.io/commands/rpushx
func (list List) RightPushX(ctx context.Context, elements ...string) (int64, error) {
	req := newRequestSize(2+len(elements), "\r\n$6\r\nRPUSHX\r\n$")
	req.addStringAndStrings(list.name, elements)
	return list.c.cmdInt(ctx, req)
}
