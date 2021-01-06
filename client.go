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
func (c *Client) Geo() *Geo {
	geo := &Geo{
		client: c.client,
	}
	return geo
}

// Set returns a client for Redis Set structure.
func (c *Client) Set() *Set {
	set := &Set{
		client: c.client,
	}
	return set
}

// HLL returns a client for Redis HyperLogLog structure.
func (c *Client) HLL() *HLL {
	hll := &HLL{
		client: c.client,
	}
	return hll
}
