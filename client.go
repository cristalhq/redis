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

// BitMap returns a client for Redis BitMap structure.
func (c *Client) BitMap(name string) *BitMap {
	return NewBitMap(name, c.client)
}

// Geo returns a client for Redis Geo structure.
func (c *Client) Geo(name string) *Geo {
	return NewGeo(name, c.client)
}

// HashMap returns a client for Redis HashMap structure.
func (c *Client) HashMap(name string) *HashMap {
	return NewHashMap(name, c.client)
}

// HyperLogLog returns a client for Redis HyperLogLog structure.
func (c *Client) HyperLogLog(name string) *HyperLogLog {
	return NewHyperLogLog(name, c.client)
}

// List returns a client for Redis List structure.
func (c *Client) List(name string) *List {
	return NewList(name, c.client)
}

// SortedSet returns a client for Redis SortedSet structure.
func (c *Client) SortedSet(name string) *SortedSet {
	return NewSortedSet(name, c.client)
}

// Set returns a client for Redis Set structure.
func (c *Client) Set(name string) *Set {
	return NewSet(name, c.client)
}

// Stream returns a client for Redis Stream structure.
func (c *Client) Stream(name string) *Stream {
	return NewStream(name, c.client)
}
