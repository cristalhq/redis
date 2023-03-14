package redis

import (
	"context"
	"net"
)

const defaultMaxIdleConns = 10

type connPool struct {
	cfg    *Config
	connCh chan *conn
}

func newConnPool(ctx context.Context, cfg *Config) (*connPool, error) {
	p := &connPool{
		cfg:    cfg,
		connCh: make(chan *conn, defaultMaxIdleConns),
	}

	var d net.Dialer
	for i := 0; i < defaultMaxIdleConns; i++ {
		c, err := d.DialContext(ctx, cfg.Network, cfg.Address)
		if err != nil {
			return nil, err
		}
		p.connCh <- &conn{Conn: c, p: p.connCh}
	}
	return p, nil
}

func (p *connPool) Close() error {
	for c := range p.connCh {
		c.Close()
	}
	return nil
}

func (p *connPool) Get(ctx context.Context) (*conn, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	case conn := <-p.connCh:
		return conn, nil

		// // TODO: create after a timeout ?
		// default:
		// 	var d net.Dialer
		// 	c, err := d.DialContext(ctx, p.network, p.address)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	return &conn{Conn: c, p: p.connCh}, nil
	}
}

func (p *connPool) Put(c *conn) {
	select {
	case c.p <- c:
	default:
		// we are over pool limit - close conn
		c.Close()
	}
}

type conn struct {
	net.Conn
	p chan *conn
}
