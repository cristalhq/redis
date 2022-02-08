package redis

import "context"

type Script struct {
	name string
	c    *Client
}

func NewScript(name string, client *Client) *Script {
	return &Script{name: name, c: client}
}

/*
EVAL
EVALSHA
EVALSHA_RO
EVAL_RO
FCALL
FCALL_RO
FUNCTION DELETE
FUNCTION DUMP
FUNCTION FLUSH
FUNCTION KILL
FUNCTION LIST
FUNCTION LOAD
FUNCTION RESTORE
FUNCTION STATS
*/

func (s *Script) DEBUG(ctx context.Context, mode string) error { return nil }
func (s *Script) EXISTS(ctx context.Context) error             { return nil }
func (s *Script) FLUSH(ctx context.Context) error              { return nil }
func (s *Script) KILL(ctx context.Context) error               { return nil }
func (s *Script) LOAD(ctx context.Context) error               { return nil }
