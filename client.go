package redis

import (
	"github.com/go-redis/redis/v8"
)

type redisClient = redis.Client

// Client represents a Redis client.
type Client struct {
	client *redisClient
}

// NewClient wrapps a redis.Client from github.com/go-redis/redis.
func NewClient(client *redisClient) (*Client, error) {
	c := &Client{
		client: client,
	}
	return c, nil
}

// Geo returns a client for Redis Geo structure.
func (c *Client) Geo(name string) *Geo {
	return NewGeo(name, c.client)
}

// Set returns a client for Redis Set structure.
func (c *Client) Set(name string) *Set {
	return NewSet(name, c.client)
}

// HyperLogLog returns a client for Redis HyperLogLog structure.
func (c *Client) HyperLogLog(name string) *HyperLogLog {
	return NewHyperLogLog(name, c.client)
}
