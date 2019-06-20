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
