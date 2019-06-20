package redis

import (
	"github.com/go-redis/redis"
)

// Client represents a Redis client.
type Client struct {
	client *redis.Client
}

// NewClient wrapps a redis.Client from github.com/go-redis/redis.
func NewClient(client *redis.Client) (*Client, error) {
	c := &Client{
		client: client,
	}
	return c, nil
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
