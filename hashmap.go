package redis

import (
	"github.com/go-redis/redis"
)

// HashMap ...
type HashMap struct {
	name   string
	client *redis.Client
}

// NewHashMap ...
func NewHashMap(name string, client *redis.Client) *HashMap {
	return &HashMap{name: name, client: client}
}

// Delete ...
func (hm *HashMap) Delete(fields ...string) (int64, error) {
	return hm.client.HDel(hm.name, fields...).Result()
}

// Exists ...
func (hm *HashMap) Exists(field string) (bool, error) {
	return hm.client.HExists(hm.name, field).Result()
}

// Get ...
func (hm *HashMap) Get(field string) (string, error) {
	return hm.client.HGet(hm.name, field).Result()
}

// GetAll ...
func (hm *HashMap) GetAll() (map[string]string, error) {
	return hm.client.HGetAll(hm.name).Result()
}

// IncBy ...
func (hm *HashMap) IncBy(field string, delta int64) (int64, error) {
	return hm.client.HIncrBy(hm.name, field, delta).Result()
}

// IncByFloat ...
func (hm *HashMap) IncByFloat(field string, delta float64) (float64, error) {
	return hm.client.HIncrByFloat(hm.name, field, delta).Result()
}

// Keys ...
func (hm *HashMap) Keys() ([]string, error) {
	return hm.client.HKeys(hm.name).Result()
}

// Len ...
func (hm *HashMap) Len() (int64, error) {
	return hm.client.HLen(hm.name).Result()
}

// MultiGet ...
func (hm *HashMap) MultiGet(fields ...string) ([]interface{}, error) {
	return hm.client.HMGet(hm.name, fields...).Result()
}

// MultiSet ...
func (hm *HashMap) MultiSet(fields map[string]interface{}) (string, error) {
	return hm.client.HMSet(hm.name, fields).Result()
}

// Scan ...
func (hm *HashMap) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return hm.client.HScan(hm.name, cursor, match, count).Result()
}

// Set ...
func (hm *HashMap) Set(field string, value interface{}) (bool, error) {
	return hm.client.HSet(hm.name, field, value).Result()
}

// SetNotExist ...
func (hm *HashMap) SetNotExist(field string, value interface{}) (bool, error) {
	return hm.client.HSetNX(hm.name, field, value).Result()
}

// Strlen ...
// TODO
func (hm *HashMap) Strlen() (int64, error) {
	return 0, nil
}

// Values ...
func (hm *HashMap) Values() ([]string, error) {
	return hm.client.HVals(hm.name).Result()
}
