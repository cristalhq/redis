package redis

import (
	"bufio"
	"context"
)

var defaultConfig = &Config{
	Network: "tcp",
	Address: ":6379",
}

type Client struct {
	pool *connPool
}

func NewClient(ctx context.Context, cfg *Config) (*Client, error) {
	if cfg == nil {
		cfg = defaultConfig
	}

	pool, err := newConnPool(ctx, cfg)
	if err != nil {
		return nil, err
	}

	c := &Client{
		pool: pool,
	}
	return c, nil
}

func (c *Client) Close() error {
	c.pool.Close()
	return nil
}

func (c *Client) send(ctx context.Context, req *request) (*conn, *bufio.Reader, error) {
	conn, err := c.pool.Get(ctx)
	if err != nil {
		return nil, nil, err
	}

	if _, err := conn.Write(req.buf); err != nil {
		c.pool.Put(conn)
		return nil, nil, err
	}

	br := responsePool.Get(conn, 1256)
	return conn, br, nil
}

func (c *Client) release(conn *conn, req *request, resp *bufio.Reader) {
	c.pool.Put(conn)
	requestPool.Put(req)
	responsePool.Put(resp)
}

func (c *Client) cmdSimple(ctx context.Context, req *request) error {
	conn, resp, err := c.send(ctx, req)
	if err != nil {
		return err
	}
	defer c.release(conn, req, resp)

	s, err := responseDecodeString(resp)
	if err != nil {
		if err == errOK || err == errNull {
			return nil
		}
		return err
	}
	if s == "OK" {
		return nil
	}
	return err
}

func (c *Client) cmdInt(ctx context.Context, req *request) (int64, error) {
	conn, resp, err := c.send(ctx, req)
	if err != nil {
		return 0, err
	}
	defer c.release(conn, req, resp)

	return responseDecodeInt(resp)
}

func (c *Client) cmdInts(ctx context.Context, req *request) ([]int64, error) {
	conn, resp, err := c.send(ctx, req)
	if err != nil {
		return nil, err
	}
	defer c.release(conn, req, resp)

	return responseDecodeInts(resp)
}

func (c *Client) cmdFloat(ctx context.Context, req *request) (float64, error) {
	conn, resp, err := c.send(ctx, req)
	if err != nil {
		return 0, err
	}
	defer c.release(conn, req, resp)

	return responseDecodeFloat(resp)
}

func (c *Client) cmdString(ctx context.Context, req *request) (string, error) {
	conn, resp, err := c.send(ctx, req)
	if err != nil {
		return "", err
	}
	defer c.release(conn, req, resp)

	s, err := responseDecodeString(resp)
	if err != nil {
		if err == errOK || err == errNull {
			return "", nil
		}
		return "", err
	}
	return s, err
}

func (c *Client) cmdStrings(ctx context.Context, req *request) ([]string, error) {
	conn, resp, err := c.send(ctx, req)
	if err != nil {
		return nil, err
	}
	defer c.release(conn, req, resp)

	return responseDecodeStrings(resp)
}
