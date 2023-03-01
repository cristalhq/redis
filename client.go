package redis

import (
	"bufio"
	"context"
	"net"
)

type Client struct {
	pool *connPool
}

func NewClient(ctx context.Context, addr string) (*Client, error) {
	pool, err := newConnPool(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}

	c := &Client{
		pool: pool,
	}
	return c, nil
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

const defaultMaxIdleConns = 10

// connPool is a connection pool.
type connPool struct {
	network string
	address string
	connCh  chan *conn
}

// newConnPool creates a new connection pool.
func newConnPool(ctx context.Context, network, address string) (*connPool, error) {
	p := &connPool{
		network: network,
		address: address,
		connCh:  make(chan *conn, defaultMaxIdleConns),
	}

	var d net.Dialer
	for i := 0; i < defaultMaxIdleConns; i++ {
		c, err := d.DialContext(ctx, p.network, p.address)
		if err != nil {
			return nil, err
		}
		p.connCh <- &conn{Conn: c, p: p.connCh}
	}
	return p, nil
}

func (p *connPool) Get(ctx context.Context) (*conn, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	case conn := <-p.connCh:
		return conn, nil

	// TODO: create after a timeout ?
	default:
		var d net.Dialer
		c, err := d.DialContext(ctx, p.network, p.address)
		if err != nil {
			return nil, err
		}
		return &conn{Conn: c, p: p.connCh}, nil
	}
}

func (p *connPool) Put(c *conn) {
	select {
	// BUG: not everything is read from a conn
	// case c.p <- c:
	default:
		// we are over pool limit - close conn
		c.Close()
	}
}

type conn struct {
	net.Conn
	p chan *conn
}
