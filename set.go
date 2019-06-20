package redis

import "github.com/go-redis/redis"

// Set represents a Redis Set structure.
type Set struct {
	client *redis.Client
}

// SAdd add one or more members to a set.
func (s *Set) SAdd(key string, members ...interface{}) (int64, error) {
	resp := s.client.SAdd(key, members)
	return resp.Result()
}

// SCard get the number of members in a set.
func (s *Set) SCard(key string) (int64, error) {
	resp := s.client.SCard(key)
	return resp.Result()
}

// SDiff subtract multiple sets.
func (s *Set) SDiff(keys ...string) ([]string, error) {
	resp := s.client.SDiff(keys...)
	return resp.Result()
}

// SDIFFSTORE subtract multiple sets and store the resulting set in a key.
func (s *Set) SDIFFSTORE(dst string, keys ...string) (int64, error) {
	resp := s.client.SDiffStore(dst, keys...)
	return resp.Result()
}

// SInter intersect multiple sets.
func (s *Set) SInter(keys ...string) ([]string, error) {
	resp := s.client.SInter(keys...)
	return resp.Result()
}

// SInterstore intersect multiple sets and store the resulting set in a key.
func (s *Set) SInterstore(dst string, keys ...string) (int64, error) {
	resp := s.client.SInterStore(dst, keys...)
	return resp.Result()
}

// SIsMember determine if a given value is a member of a set.
func (s *Set) SIsMember(key string, member interface{}) (bool, error) {
	resp := s.client.SIsMember(key, member)
	return resp.Result()
}

// SMembers get all the members in a set.
func (s *Set) SMembers(key string) ([]string, error) {
	resp := s.client.SMembers(key)
	return resp.Result()
}

// SMove move a member from one set to another.
func (s *Set) SMove(src, dst string, member interface{}) (bool, error) {
	resp := s.client.SMove(src, dst, member)
	return resp.Result()
}

// SPop remove and return one or multiple random members from a set.
func (s *Set) SPop(key string) (string, error) {
	resp := s.client.SPop(key)
	return resp.Result()
}

// SRandMember get one or multiple random members from a set.
func (s *Set) SRandMember(key string) (string, error) {
	resp := s.client.SRandMember(key)
	return resp.Result()
}

// SRem Remove one or more members from a set
func (s *Set) SRem(key string, members ...interface{}) (int64, error) {
	resp := s.client.SRem(key, members)
	return resp.Result()
}

// SScan incrementally iterate Set elements.
func (s *Set) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	resp := s.client.SScan(key, cursor, match, count)
	return resp.Result()
}

// SUnion add multiple sets.s
func (s *Set) SUnion(keys ...string) ([]string, error) {
	resp := s.client.SUnion(keys...)
	return resp.Result()
}

// SUnionStore add multiple sets and store the resulting set in a key.
func (s *Set) SUnionStore(dst string, keys ...string) (int64, error) {
	resp := s.client.SUnionStore(dst, keys...)
	return resp.Result()
}
