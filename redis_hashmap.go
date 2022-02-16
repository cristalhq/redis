package redis

import (
	"context"
)

// HashMap client for hashmap operations in Redis.
// See: https://redis.io/commands#hash
type HashMap struct {
	name string
	c    *Client
}

// NewHashMap instantiates new Redis HashMap client.
func NewHashMap(name string, client *Client) HashMap {
	return HashMap{name: name, c: client}
}

// Name of the HashMap structure.
func (hm HashMap) Name() string { return hm.name }

// Delete the specified fields from the hash.
// See: https://redis.io/commands/hdel
func (hm HashMap) Delete(ctx context.Context, fields ...string) (int64, error) {
	req := newRequestSize(2+len(fields), "\r\n$4\r\nHDEL\r\n$")
	req.addStringAndStrings(hm.name, fields)
	return hm.c.cmdInt(ctx, req)
}

// Exists returns if field is an existing field in the hash.
// See: https://redis.io/commands/hexists
func (hm HashMap) Exists(ctx context.Context, field string) (bool, error) {
	req := newRequest("*3\r\n$7\r\nHEXISTS\r\n$")
	req.addString2(hm.name, field)
	res, err := hm.c.cmdInt(ctx, req)
	return res == 1, err
}

// Get returns the value associated with field in the hash.
// See: https://redis.io/commands/hget
func (hm HashMap) Get(ctx context.Context, field string) (string, error) {
	req := newRequest("*3\r\n$4\r\nHGET\r\n$")
	req.addString2(hm.name, field)
	return hm.c.cmdString(ctx, req)
}

// GetAll returns all fields and values of the hash.
// See: https://redis.io/commands/hgetall
func (hm HashMap) GetAll(ctx context.Context) (map[string]string, error) {
	req := newRequest("*2\r\n$7\r\nHGETALL\r\n$")
	req.addString(hm.name)
	ss, err := hm.c.cmdStrings(ctx, req)
	if err != nil {
		return nil, err
	}
	res := make(map[string]string, len(ss)/2)
	for i := 0; i < len(ss); i += 2 {
		res[ss[i]] = ss[i+1]
	}
	return res, nil
}

// IncBy increments the number stored at field in the hash.
// See: https://redis.io/commands/hincrby
func (hm HashMap) IncBy(ctx context.Context, field string, delta int64) (int64, error) {
	req := newRequest("*4\r\n$7\r\nHINCRBY\r\n$")
	req.addString2AndInt(hm.name, field, delta)
	return hm.c.cmdInt(ctx, req)
}

// IncByFloat increments the number stored at field in the hash.
// See: https://redis.io/commands/hincrbyfloat
func (hm HashMap) IncByFloat(ctx context.Context, field string, delta float64) (float64, error) {
	req := newRequest("*4\r\n$12\r\nHINCRBYFLOAT\r\n$")
	// TODO(oleg): must encode as float
	req.addString2AndInt(hm.name, field, int64(delta))
	return hm.c.cmdFloat(ctx, req)
}

// Keys returns all field names in the hash.
// See: https://redis.io/commands/hkeys
func (hm HashMap) Keys(ctx context.Context) ([]string, error) {
	req := newRequest("*2\r\n$5\r\nHKEYS\r\n$")
	req.addString(hm.name)
	return hm.c.cmdStrings(ctx, req)
}

// Len returns the number of fields contained in the hash.
// See: https://redis.io/commands/hlen
func (hm HashMap) Len(ctx context.Context) (int64, error) {
	req := newRequest("*2\r\n$4\r\nHLEN\r\n$")
	req.addString(hm.name)
	return hm.c.cmdInt(ctx, req)
}

// MultiGet returns the values associated with the specified fields in the hash.
// See: https://redis.io/commands/hmget
func (hm HashMap) MultiGet(ctx context.Context, fields ...string) ([]Value, error) {
	req := newRequestSize(2+len(fields), "\r\n$5\r\nHMGET\r\n$")
	req.addStringAndStrings(hm.name, fields)
	return hm.c.cmdStrings(ctx, req)
}

// MultiSet the specified fields to their respective values in the hash.
// This command overwrites any specified fields already existing in the hash.
//
// Deprecation notice: as of Redis version 4.0.0 this command is considered as deprecated.
// While it is unlikely that it will be completely removed, prefer using `HSET` with multiple field-value pairs in its stead.
//
// See: https://redis.io/commands/hmset
func (hm HashMap) MultiSet(ctx context.Context, fields map[string]Value) error {
	req := newRequestSize(2+2*len(fields), "\r\n$5\r\nHMSET\r\n$")
	req.addStringAndMap(hm.name, fields)
	return hm.c.cmdSimple(ctx, req)
}

// RandomField TODO
// See: https://redis.io/commands/hrandfield
func (hm HashMap) RandomField(ctx context.Context, cursor uint64, match string, count int64) (_ []string, _ uint64, _ error) {
	panic("redis: HashMap.RandomField not implemented")
}

// Scan TODO
// See: https://redis.io/commands/hscan
func (hm HashMap) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	panic("redis: HashMap.Scan not implemented")
}

// Set field in the hash to value.
// If set does not exist, a new set is created.
// If field already exists in the hash, it is overwritten.
// See: https://redis.io/commands/hset
func (hm HashMap) Set(ctx context.Context, field string, value Value) error {
	req := newRequest("*4\r\n$4\r\nHSET\r\n$")
	req.addString3(hm.name, field, value)
	_, err := hm.c.cmdInt(ctx, req)
	return err
}

// SetNotExist field in the hash to value, only if field does not yet exist.
// If set does not exist, a new set is created.
// If field already exists, this operation has no effect.
// See: https://redis.io/commands/hsetnx
func (hm HashMap) SetNotExist(ctx context.Context, field string, value Value) (bool, error) {
	req := newRequest("*4\r\n$6\r\nHSETNX\r\n$")
	req.addString3(hm.name, field, value)
	res, err := hm.c.cmdInt(ctx, req)
	return res == 1, err
}

// Strlen returns the string length of the value associated with field in the hash.
// If the key or the field do not exist, 0 is returned.
// See: https://redis.io/commands/hstrlen
func (hm HashMap) Strlen(ctx context.Context, key string) (int64, error) {
	req := newRequest("*3\r\n$7\r\nHSTRLEN\r\n$")
	req.addString2(hm.name, key)
	return hm.c.cmdInt(ctx, req)
}

// Values returns all values in the hash.
// See: https://redis.io/commands/hvals
func (hm HashMap) Values(ctx context.Context) ([]string, error) {
	req := newRequest("*2\r\n$5\r\nHVALS\r\n$")
	req.addString(hm.name)
	return hm.c.cmdStrings(ctx, req)
}
