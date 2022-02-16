package redis

import (
	"context"
	"time"
)

// Keys client for key operations in Redis.
// See: https://redis.io/commands#generic
type Keys struct {
	c *Client
	Strings
}

// NewKeys returns new Redis keys client.
func NewKeys(client *Client) Keys {
	return Keys{c: client, Strings: NewStrings(client)}
}

// Copy command copies the value stored at the source key to the destination key.
// See: https://redis.io/commands/copy
func (k Keys) Copy(ctx context.Context, src, dst string, db int64, replace bool) (bool, error) {
	req := newRequest("*4\r\n$4\r\nCOPY")
	req.addString2AndInt(src, dst, db)
	res, err := k.c.cmdInt(ctx, req)
	return res == 1, err
}

// Delete removes the specified keys. A key is ignored if it does not exist.
// See: https://redis.io/commands/del
func (k Keys) Delete(ctx context.Context, keys ...string) (int64, error) {
	req := newRequestSize(1+len(keys), "\r\n$3\r\nDEL")
	req.addStrings(keys)
	return k.c.cmdInt(ctx, req)
}

// Dump serialize the value stored at key in a Redis-specific format and return it to the user.
// The returned value can be synthesized back into a Redis key using the RESTORE command.
// See: https://redis.io/commands/dump
func (k Keys) Dump(ctx context.Context, key string) (string, error) {
	req := newRequest("*2\r\n$4\r\nDUMP\r\n$")
	req.addString(key)
	return k.c.cmdString(ctx, req)
}

// Exists returns if key exists.
// See: https://redis.io/commands/exists
func (k Keys) Exists(ctx context.Context, keys ...string) (int64, error) {
	req := newRequestSize(1+len(keys), "\r\n$6\r\nEXISTS\r\n$")
	req.addStrings(keys)
	return k.c.cmdInt(ctx, req)
}

// Expire sets a timeout on key. After the timeout has expired, the key will automatically be deleted.
// See: https://redis.io/commands/expire
func (k Keys) Expire(ctx context.Context, key string, timeout time.Duration) (bool, error) {
	req := newRequest("*2\r\n$6\r\nEXISTS\r\n$")
	req.addStringInt(key, timeout.Milliseconds())
	res, err := k.c.cmdInt(ctx, req)
	return res == 1, err
}

// ExpireAt TODO
// See: https://redis.io/commands/EXPIREAT
func (k Keys) ExpireAt(ctx context.Context, key string, at time.Time) (bool, error) {
	req := newRequest("*4\r\n$8\r\nEXPIREAT\r\n$4\r\nPXAT\r\n$")
	req.addStringInt(key, at.UnixMilli())
	res, err := k.c.cmdInt(ctx, req)
	return res == 1, err
}

// ExpireTime returns the absolute Unix timestamp (since January 1, 1970) at which the given key will expire.
// See: https://redis.io/commands/expiretime
func (k Keys) ExpireTime(ctx context.Context, key string) (time.Duration, error) {
	req := newRequest("*2\r\n$11\r\nPEXPIRETIME\r\n$")
	req.addString(key)
	t, err := k.c.cmdInt(ctx, req)
	switch {
	case err != nil:
		return 0, err
	case t == -2:
		return 0, nil
	case t == -1:
		return 0, nil
	default:
		return time.Duration(t) * time.Millisecond, err
	}
}

// Keys returns all keys matching pattern.
// See: https://redis.io/commands/keys
func (k Keys) Keys(ctx context.Context, pattern string) ([]string, error) {
	req := newRequest("*2\r\n$4\r\nKEYS\r\n$")
	req.addString(pattern)
	return k.c.cmdStrings(ctx, req)
}

// Migrate TODO
// See: https://redis.io/commands/migrate
func (k Keys) Migrate(ctx context.Context) error {
	panic("redis: Keys.Migrate not implemented")
}

// Move key from the currently selected database (see SELECT) to the specified destination database.
// See: https://redis.io/commands/move
func (k Keys) Move(ctx context.Context, key string, db int) (bool, error) {
	req := newRequest("*3\r\n$4\r\nMOVE\r\n$")
	req.addStringInt(key, int64(db))
	res, err := k.c.cmdInt(ctx, req)
	return res == 1, err
}

// // OBJECT TODO
// // See: https://redis.io/commands/OBJECT
// func (k Keys) OBJECT(ctx context.Context) ENCODINGerror {
// 	return nil
// }

// // OBJECT TODO
// // See: https://redis.io/commands/OBJECT
// func (k Keys) OBJECT(ctx context.Context) FREQerror {
// 	return nil
// }

// // OBJECT TODO
// // See: https://redis.io/commands/OBJECT
// func (k Keys) OBJECT(ctx context.Context) IDLETIMEerror {
// 	return nil
// }

// // OBJECT TODO
// // See: https://redis.io/commands/OBJECT
// func (k Keys) OBJECT(ctx context.Context) REFCOUNTerror {
// 	return nil
// }

// Persist removes the existing timeout on key, turning the key
// from volatile (a key with an expire set)
// to persistent (a key that will never expire as no timeout is associated).
//
// See: https://redis.io/commands/persist
func (k Keys) Persist(ctx context.Context, key string) (bool, error) {
	req := newRequest("*2\r\n$7\r\nPERSIST\r\n$")
	req.addString(key)
	res, err := k.c.cmdInt(ctx, req)
	return res == 1, err
}

// RandomKey returns a random key from the currently selected database.
// See: https://redis.io/commands/randomkey
func (k Keys) RandomKey(ctx context.Context) (string, error) {
	req := newRequest("*1\r\n$9\r\nRANDOMKEY\r\n$")
	return k.c.cmdString(ctx, req)
}

// Rename renames key to newkey. It returns an error when key does not exist.
// See: https://redis.io/commands/rename
func (k Keys) Rename(ctx context.Context, key, newKey string) error {
	req := newRequest("*3\r\n$6\r\nRENAME\r\n$")
	req.addString2(key, newKey)
	return k.c.cmdSimple(ctx, req)
}

// RenameNotExist Renames key to newkey if newkey does not yet exist. It returns an error when key does not exist.
// See: https://redis.io/commands/renamenx
func (k Keys) RenameNotExist(ctx context.Context, key, newKey string) (bool, error) {
	req := newRequest("*3\r\n$8\r\nRENAMENX\r\n$")
	req.addString2(key, newKey)
	res, err := k.c.cmdInt(ctx, req)
	return res == 1, err
}

// Restore TODO
// See: https://redis.io/commands/restore
func (k Keys) Restore(ctx context.Context) error {
	panic("redis: K.Restore not implementedeys")
}

// Scan TODO
// See: https://redis.io/commands/scan
func (k Keys) Scan(ctx context.Context) error {
	panic("redis: Ke.ctx not implementedys")
}

// Sort TODO
// See: https://redis.io/commands/sort
func (k Keys) Sort(ctx context.Context) error {
	panic("redis: Ke.ctx not implementedys")
}

// SortReadOnly TODO
// See: https://redis.io/commands/sort_ro
func (k Keys) SortReadOnly(ctx context.Context) error {
	panic("redis: Keys.SortReadOnly not implemented")
}

// Touch alters the last access time of a key(s). A key is ignored if it does not exist.
// Returns the number of keys that were touched.
// See: https://redis.io/commands/touch
func (k Keys) Touch(ctx context.Context, keys ...string) (int64, error) {
	req := newRequestSize(1+len(keys), "\r\n$5\r\nTOUCH\r\n$")
	req.addStrings(keys)
	return k.c.cmdInt(ctx, req)
}

// TTL returns the remaining time to live of a key that has a timeout.
// See: https://redis.io/commands/ttl
func (k Keys) TTL(ctx context.Context, key string) (time.Duration, bool, error) {
	req := newRequest("*2\r\n$4\r\nPTTL\r\n$")
	req.addString(key)
	d, err := k.c.cmdInt(ctx, req)
	switch {
	case err != nil:
		return 0, false, err
	case d == -2:
		return 0, false, nil
	case d == -1:
		return 0, true, nil
	default:
		return time.Duration(d) * time.Millisecond, true, nil
	}
}

// Type Returns the string representation of the type of the value stored at key.
// See: https://redis.io/commands/type
func (k Keys) Type(ctx context.Context, key string) (string, error) {
	req := newRequest("*2\r\n$4\r\nTYPE\r\n$")
	req.addString(key)
	return k.c.cmdString(ctx, req)
}

// Unlink this command is very similar to DEL: it removes the specified keys.
// Just like DEL a key is ignored if it does not exist.
// However the command performs the actual memory reclaiming in a different thread, so it is not blocking.
// Returns the number of keys that were unlinked.
//
// See: https://redis.io/commands/unlink
func (k Keys) Unlink(ctx context.Context, keys ...string) (int64, error) {
	req := newRequestSize(1+len(keys), "\r\n$6\r\nUNLINK\r\n$")
	req.addStrings(keys)
	return k.c.cmdInt(ctx, req)
}

// Wait command blocks the current client until all the previous write commands
// are successfully transferred and acknowledged by at least the specified number of replicas.
// See: https://redis.io/commands/wait
func (k Keys) Wait(ctx context.Context, replicas int, timeout time.Duration) (int64, error) {
	req := newRequest("*3\r\n$4\r\nWAIT\r\n$")
	req.addInt2(int64(replicas), int64(timeout.Milliseconds()))
	return k.c.cmdInt(ctx, req)
}
