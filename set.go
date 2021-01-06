package redis

import (
	"context"
)

// Set represents a Redis Set structure.
type Set struct {
	name   string
	client *redisClient
}

// NewSet instantiates a new Set structure client for Redis.
func NewSet(name string, client *redisClient) *Set {
	return &Set{name: name, client: client}
}

// Add one or more members to a set.
func (s *Set) Add(ctx context.Context, members ...interface{}) (int64, error) {
	return s.client.SAdd(ctx, s.name, members).Result()
}

// Cardinality returns the number of members in a set.
func (s *Set) Cardinality(ctx context.Context, key string) (int64, error) {
	return s.client.SCard(ctx, key).Result()
}

// IsMember determines if a given value is a member of a set.
func (s *Set) IsMember(ctx context.Context, member interface{}) (bool, error) {
	return s.client.SIsMember(ctx, s.name, member).Result()
}

// Members returns all the members in a set.
func (s *Set) Members(ctx context.Context, key string) ([]string, error) {
	return s.client.SMembers(ctx, key).Result()
}

// Move a member from one set to another.
func (s *Set) Move(ctx context.Context, src, dst string, member interface{}) (bool, error) {
	return s.client.SMove(ctx, src, dst, member).Result()
}

// Pop removes and returns one or multiple random members from a set.
func (s *Set) Pop(ctx context.Context, key string) (string, error) {
	return s.client.SPop(ctx, key).Result()
}

// RandomMember gets one random member from a set.
func (s *Set) RandomMember(ctx context.Context, key string) (string, error) {
	return s.client.SRandMember(ctx, key).Result()
}

// RandomMembers gets one random member from a set.
func (s *Set) RandomMembers(ctx context.Context, key string, count int64) ([]string, error) {
	return s.client.SRandMemberN(ctx, key, count).Result()
}

// Remove one or more members from a set.
func (s *Set) Remove(ctx context.Context, members ...interface{}) (int64, error) {
	return s.client.SRem(ctx, s.name, members).Result()
}

// Scan incrementally iterate Set elements.
func (s *Set) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return s.client.SScan(ctx, s.name, cursor, match, count).Result()
}

// Difference subtract multiple sets.
func (s *Set) Difference(ctx context.Context, keys ...string) ([]string, error) {
	return s.client.SDiff(ctx, keys...).Result()
}

// DifferenceStore subtract multiple sets and store the resulting set in a key.
func (s *Set) DifferenceStore(ctx context.Context, dst string, keys ...string) (int64, error) {
	return s.client.SDiffStore(ctx, dst, keys...).Result()
}

// Intersection intersect multiple sets.
func (s *Set) Intersection(ctx context.Context, keys ...string) ([]string, error) {
	return s.client.SInter(ctx, keys...).Result()
}

// InterectionStore intersect multiple sets and store the resulting set in a key.
func (s *Set) InterectionStore(ctx context.Context, dst string, keys ...string) (int64, error) {
	return s.client.SInterStore(ctx, dst, keys...).Result()
}

// Union add multiple sets.
func (s *Set) Union(ctx context.Context, keys ...string) ([]string, error) {
	return s.client.SUnion(ctx, keys...).Result()
}

// UnionStore add multiple sets and store the resulting set in a key.
func (s *Set) UnionStore(ctx context.Context, dst string, keys ...string) (int64, error) {
	return s.client.SUnionStore(ctx, dst, keys...).Result()
}
