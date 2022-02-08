package redis

import (
	"context"
	"errors"
)

// Set represents a Redis Set structure.
type Set struct {
	name string
	c    *Client
}

// NewSet instantiates a new Set structure client for Redis.
func NewSet(name string, client *Client) *Set {
	return &Set{name: name, c: client}
}

// Name of the Set structure.
func (set *Set) Name() string { return set.name }

// Add one or more members to a set. Specified members that are already a member of this set are ignored.
// If key does not exist, a new set is created before adding the specified members.
// Returns the number of elements that were added to the set, not including all the elements already present in the set.
//
// See: https://redis.io/commands/sadd
func (set *Set) Get(ctx context.Context, members ...string) (int64, error) {
	req := newRequest("*3\r\n$3\r\nGET\r\n$")
	req.addString2(set.name, members[0])
	return set.c.cmdInt(ctx, req)
}

// Add one or more members to a set. Specified members that are already a member of this set are ignored.
// If key does not exist, a new set is created before adding the specified members.
// Returns the number of elements that were added to the set, not including all the elements already present in the set.
//
// See: https://redis.io/commands/sadd
func (set *Set) Add(ctx context.Context, members ...string) (int64, error) {
	req := newRequestSize(2+len(members), "\r\n$4\r\nSADD\r\n$")
	req.addStringAndStrings(set.name, members)
	return set.c.cmdInt(ctx, req)
}

// Cardinality returns the number of members in a set.
// See: https://redis.io/commands/scard
func (set *Set) Cardinality(ctx context.Context) (int64, error) {
	req := newRequest("*2\r\n$5\r\nSCARD\r\n$")
	req.addString(set.name)
	return set.c.cmdInt(ctx, req)
}

// Diff subtract multiple sets.
// See: https://redis.io/commands/sdiff
func (set *Set) Diff(ctx context.Context, keys ...string) ([]string, error) {
	req := newRequestSize(2+len(keys), "\r\n$5\r\nSDIFF\r\n$")
	req.addStringAndStrings(set.name, keys)
	return set.c.cmdStrings(ctx, req)
}

// DiffStore subtract multiple sets and store the resulting set in a key.
// See: https://redis.io/commands/sdiffstore
func (set *Set) DiffStore(ctx context.Context, dst string, keys ...string) (int64, error) {
	req := newRequestSize(3+len(keys), "\r\n$10\r\nSDIFFSTORE\r\n$")
	req.addString2AndStrings(dst, set.name, keys)
	return set.c.cmdInt(ctx, req)
}

// Inter intersect multiple sets.
// See: https://redis.io/commands/sinter
func (set *Set) Inter(ctx context.Context, keys ...string) ([]string, error) {
	req := newRequestSize(2+len(keys), "\r\n$6\r\nSINTER\r\n$")
	req.addStringAndStrings(set.name, keys)
	return set.c.cmdStrings(ctx, req)
}

// TODO(oleg): new in Redis 7.0.0
// // InterCard intersect multiple sets and store the resulting set in a key.
// // See: https://redis.io/commands/sintercard
// func (set *Set) InterCard(ctx context.Context, dst string, keys ...string) (int64, error) {
// 	req := newRequestSize(3+len(keys), "\r\n$10\r\nSINTERCARD\r\n$")
// 	req.addString2AndStrings(dst, set.name, keys)
// 	return set.c.cmdInt(ctx, req)
// }

// InterStore intersect multiple sets and store the resulting set in a key.
// See: https://redis.io/commands/sinterstore
func (set *Set) InterStore(ctx context.Context, dst string, keys ...string) (int64, error) {
	req := newRequestSize(3+len(keys), "\r\n$11\r\nSINTERSTORE\r\n$")
	req.addString2AndStrings(dst, set.name, keys)
	return set.c.cmdInt(ctx, req)
}

// IsMember returns if member is a member of the set.
// See: https://redis.io/commands/sismember
func (set *Set) IsMember(ctx context.Context, member string) (bool, error) {
	req := newRequest("*3\r\n$9\r\nSISMEMBER\r\n$")
	req.addString2(set.name, member)
	res, err := set.c.cmdInt(ctx, req)
	return res == 1, err
}

// Members returns all the members in a set.
// See: https://redis.io/commands/smembers
func (set *Set) Members(ctx context.Context) ([]string, error) {
	req := newRequest("*2\r\n$8\r\nSMEMBERS\r\n$")
	req.addString(set.name)
	return set.c.cmdStrings(ctx, req)
}

// MultiIsMembers returns whether each member is a member of the set.
// See: https://redis.io/commands/smismember
func (set *Set) MultiIsMembers(ctx context.Context, members ...string) ([]bool, error) {
	req := newRequestSize(2+len(members), "\r\n$10\r\nSMISMEMBER\r\n$")
	req.addStringAndStrings(set.name, members)
	ss, err := set.c.cmdInts(ctx, req)
	if err != nil {
		return nil, err
	}
	res := make([]bool, len(ss))
	for i := range ss {
		res[i] = ss[i] == 1
	}
	return res, nil
}

// Move a member from one set to another (dst).
// See: https://redis.io/commands/smove
func (set *Set) MoveTo(ctx context.Context, dst, member string) (bool, error) {
	req := newRequest("*4\r\n$5\r\nSMOVE\r\n$")
	req.addString3(set.name, dst, member)
	res, err := set.c.cmdInt(ctx, req)
	return res == 1, err
}

// Pop removes and returns one or multiple random members from a set.
// See: https://redis.io/commands/spop
func (set *Set) Pop(ctx context.Context) (string, error) {
	req := newRequest("*2\r\n$4\r\nSPOP\r\n$")
	req.addString(set.name)
	return set.c.cmdString(ctx, req)
}

// RandomMember gets one random member from a set.
// See: https://redis.io/commands/srandmember
func (set *Set) RandomMember(ctx context.Context) (string, error) {
	req := newRequest("*2\r\n$11\r\nSRANDMEMBER\r\n$")
	req.addString(set.name)
	s, err := set.c.cmdString(ctx, req)
	if err != nil && err.Error() != "OK" {
		return "", err
	}
	return s, nil
}

// RandomMembers gets one random member from a set.
// See: https://redis.io/commands/srandmember
func (set *Set) RandomMembers(ctx context.Context, count int64) ([]string, error) {
	req := newRequestSize(int(1+count), "\r\n$11\r\nSRANDMEMBER\r\n$")
	req.addStringInt(set.name, count)
	return set.c.cmdStrings(ctx, req)
}

// Remove one or more members from a set.
// See: https://redis.io/commands/srem
func (set *Set) Remove(ctx context.Context, members ...string) (int64, error) {
	req := newRequestSize(2+len(members), "\r\n$4\r\nSREM\r\n$")
	req.addStringAndStrings(set.name, members)
	return set.c.cmdInt(ctx, req)
}

// Scan incrementally iterate Set elements.
// See: https://redis.io/commands/sscan
func (set *Set) Scan(ctx context.Context, cursor uint64, match string, count int64) (_ []string, _ uint64, _ error) {
	// TODO(oleg): implement
	return nil, 0, errors.New("unimplemented")
}

// Union add multiple sets.
// See: https://redis.io/commands/sunion
func (set *Set) Union(ctx context.Context, keys ...string) ([]string, error) {
	req := newRequestSize(2+len(keys), "\r\n$6\r\nSUNION\r\n$")
	req.addStringAndStrings(set.name, keys)
	return set.c.cmdStrings(ctx, req)
}

// UnionStore add multiple sets and store the resulting set in a key.
// See: https://redis.io/commands/sunionstore
func (set *Set) UnionStore(ctx context.Context, dst string, keys ...string) (int64, error) {
	req := newRequestSize(3+len(keys), "\r\n$11\r\nSUNIONSTORE\r\n$")
	req.addString2AndStrings(dst, set.name, keys)
	return set.c.cmdInt(ctx, req)
}
