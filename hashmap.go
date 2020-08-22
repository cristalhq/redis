package redis

import (
	"github.com/go-redis/redis"
)

type HashMap struct {
	name   string
	client *redis.Client
}

func NewHashMap(name string, client *redis.Client) *HashMap {
	return &HashMap{name: name, client: client}
}

func (hm *HashMap) Delete(fields ...string) (int64, error) {
	return hm.client.HDel(hm.name, fields...).Result()
}

func (hm *HashMap) Exists(field string) (bool, error) {
	return hm.client.HExists(hm.name, field).Result()
}

func (hm *HashMap) Get(field string) (string, error) {
	return hm.client.HGet(hm.name, field).Result()
}

func (hm *HashMap) GetAll() (map[string]string, error) {
	return hm.client.HGetAll(hm.name).Result()
}

func (hm *HashMap) IncBy(field string, delta int64) (int64, error) {
	return hm.client.HIncrBy(hm.name, field, delta).Result()
}

func (hm *HashMap) IncByFloat(field string, delta float64) (float64, error) {
	return hm.client.HIncrByFloat(hm.name, field, delta).Result()
}

func (hm *HashMap) Keys() ([]string, error) {
	return hm.client.HKeys(hm.name).Result()
}

func (hm *HashMap) Len() (int64, error) {
	return hm.client.HLen(hm.name).Result()
}

func (hm *HashMap) MultiGet(fields ...string) ([]interface{}, error) {
	return hm.client.HMGet(hm.name, fields...).Result()
}

func (hm *HashMap) MultiSet(fields map[string]interface{}) (string, error) {
	return hm.client.HMSet(hm.name, fields).Result()
}

func (hm *HashMap) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return hm.client.HScan(hm.name, cursor, match, count).Result()
}

func (hm *HashMap) Set(field string, value interface{}) (bool, error) {
	return hm.client.HSet(hm.name, field, value).Result()
}

func (hm *HashMap) SetNotExist(field string, value interface{}) (bool, error) {
	return hm.client.HSetNX(hm.name, field, value).Result()
}

// TODO
func (hm *HashMap) Strlen() (int64, error) {
	return 0, nil
}

func (hm *HashMap) Values() ([]string, error) {
	return hm.client.HVals(hm.name).Result()
}
