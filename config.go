package redis

import (
	"fmt"
	"net/url"
	"strconv"
)

type Config struct {
	Network string
	Address string
	Port    int
}

func (c *Config) Validate() error {
	switch {
	case c.Network != "tcp" && c.Network != "unix":
		return fmt.Errorf("unsupported network: %s", c.Network)
	default:
		return nil
	}
}

func ParseURL(s string) (*Config, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(u.Port())
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Network: u.Scheme,
		Address: u.Host,
		Port:    port,
	}
	return cfg, nil
}
