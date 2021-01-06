package redis

import (
	"context"
)

// HashMap ...
type HashMap struct {
	name   string
	client *redisClient
}

// NewHashMap ...
func NewHashMap(name string, client *redisClient) *HashMap {
	return &HashMap{name: name, client: client}
}

// Delete ...
func (hm *HashMap) Delete(ctx context.Context, fields ...string) (int64, error) {
	return hm.client.HDel(ctx, hm.name, fields...).Result()
}

// Exists ...
func (hm *HashMap) Exists(ctx context.Context, field string) (bool, error) {
	return hm.client.HExists(ctx, hm.name, field).Result()
}

// Get ...
func (hm *HashMap) Get(ctx context.Context, field string) (string, error) {
	return hm.client.HGet(ctx, hm.name, field).Result()
}

// GetAll ...
func (hm *HashMap) GetAll(ctx context.Context) (map[string]string, error) {
	return hm.client.HGetAll(ctx, hm.name).Result()
}

// IncBy ...
func (hm *HashMap) IncBy(ctx context.Context, field string, delta int64) (int64, error) {
	return hm.client.HIncrBy(ctx, hm.name, field, delta).Result()
}

// IncByFloat ...
func (hm *HashMap) IncByFloat(ctx context.Context, field string, delta float64) (float64, error) {
	return hm.client.HIncrByFloat(ctx, hm.name, field, delta).Result()
}

// Keys ...
func (hm *HashMap) Keys(ctx context.Context) ([]string, error) {
	return hm.client.HKeys(ctx, hm.name).Result()
}

// Len ...
func (hm *HashMap) Len(ctx context.Context) (int64, error) {
	return hm.client.HLen(ctx, hm.name).Result()
}

// MultiGet ...
func (hm *HashMap) MultiGet(ctx context.Context, fields ...string) ([]interface{}, error) {
	return hm.client.HMGet(ctx, hm.name, fields...).Result()
}

// MultiSet ...
func (hm *HashMap) MultiSet(ctx context.Context, fields map[string]interface{}) (bool, error) {
	return hm.client.HMSet(ctx, hm.name, fields).Result()
}

// Scan ...
func (hm *HashMap) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return hm.client.HScan(ctx, hm.name, cursor, match, count).Result()
}

// Set ...
func (hm *HashMap) Set(ctx context.Context, field string, value interface{}) (int64, error) {
	return hm.client.HSet(ctx, hm.name, field, value).Result()
}

// SetNotExist ...
func (hm *HashMap) SetNotExist(ctx context.Context, field string, value interface{}) (bool, error) {
	return hm.client.HSetNX(ctx, hm.name, field, value).Result()
}

// Strlen ...
// TODO
func (hm *HashMap) Strlen(ctx context.Context) (int64, error) {
	return 0, nil
}

// Values ...
func (hm *HashMap) Values(ctx context.Context) ([]string, error) {
	return hm.client.HVals(ctx, hm.name).Result()
}
