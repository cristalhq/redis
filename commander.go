package redis

// Commander is a simple Redis client.
type Commander struct {
	c *Client
}

// NewCommander returns new Commander.
func NewCommander(client *Client) Commander {
	return Commander{c: client}
}
