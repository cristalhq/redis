package redis

import "github.com/go-redis/redis"

// Set represents a Redis Set structure.
type Set struct {
	name   string
	client *redis.Client
}

func NewSet(name string, client *redis.Client) *Set {
	return &Set{name: name, client: client}
}

// Add add one or more members to a set.
func (s *Set) Add(members ...interface{}) (int64, error) {
	return s.client.SAdd(s.name, members).Result()
}

// Cardinality get the number of members in a set.
func (s *Set) Cardinality(key string) (int64, error) {
	return s.client.SCard(key).Result()
}

// IsMember determine if a given value is a member of a set.
func (s *Set) IsMember(member interface{}) (bool, error) {
	return s.client.SIsMember(s.name, member).Result()
}

// Members get all the members in a set.
func (s *Set) Members(key string) ([]string, error) {
	return s.client.SMembers(key).Result()
}

// Move move a member from one set to another.
func (s *Set) Move(src, dst string, member interface{}) (bool, error) {
	return s.client.SMove(src, dst, member).Result()
}

// Pop remove and return one or multiple random members from a set.
func (s *Set) Pop(key string) (string, error) {
	return s.client.SPop(key).Result()
}

// RandomMember get one or multiple random members from a set.
func (s *Set) RandomMember(key string) (string, error) {
	return s.client.SRandMember(key).Result()
}

// Remove Remove one or more members from a set
func (s *Set) Remove(members ...interface{}) (int64, error) {
	return s.client.SRem(s.name, members).Result()
}

// Scan incrementally iterate Set elements.
func (s *Set) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return s.client.SScan(s.name, cursor, match, count).Result()
}

// Difference subtract multiple sets.
func (s *Set) Difference(keys ...string) ([]string, error) {
	return s.client.SDiff(keys...).Result()
}

// DIFFSTORE subtract multiple sets and store the resulting set in a key.
func (s *Set) DIFFSTORE(dst string, keys ...string) (int64, error) {
	return s.client.SDiffStore(dst, keys...).Result()
}

// Intersection intersect multiple sets.
func (s *Set) Intersection(keys ...string) ([]string, error) {
	return s.client.SInter(keys...).Result()
}

// Interstore intersect multiple sets and store the resulting set in a key.
func (s *Set) Interstore(dst string, keys ...string) (int64, error) {
	return s.client.SInterStore(dst, keys...).Result()
}

// Union add multiple sets.
func (s *Set) Union(keys ...string) ([]string, error) {
	return s.client.SUnion(keys...).Result()
}

// UnionStore add multiple sets and store the resulting set in a key.
func (s *Set) UnionStore(dst string, keys ...string) (int64, error) {
	return s.client.SUnionStore(dst, keys...).Result()
}
