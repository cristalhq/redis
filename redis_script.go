package redis

import (
	"context"
	"fmt"
)

// Scripting client for script operations in Redis.
// See: https://redis.io/commands#scripting
//
// Redis functions: https://redis.io/topics/functions-intro
type Scripting struct {
	c *Client
}

// NewScripting returns new Redis Script client.
func NewScripting(client *Client) Scripting {
	return Scripting{c: client}
}

// Eval invokes the execution of a server-side Lua script.
// See: https://redis.io/commands/eval
func (sc Scripting) Eval(ctx context.Context, script string, keys []string, args ...string) ([]string, error) {
	req := newRequestSize(3+len(keys)+len(args), "\r\n$4\r\nEVAL\r\n$")
	req.addStringIntStrings(script, int64(len(keys)), append(keys, args...))
	return sc.c.cmdStrings(ctx, req)
}

// EvalSHA evaluates a script from the server's cache by its SHA1 digest.
// See: https://redis.io/commands/evalsha
func (sc Scripting) EvalSHA(ctx context.Context, hash string, keys []string, args ...string) ([]string, error) {
	req := newRequestSize(3+len(keys)+len(args), "\r\n$7\r\nEVALSHA\r\n$")
	req.addStringIntStrings(hash, int64(len(keys)), append(keys, args...))
	return sc.c.cmdStrings(ctx, req)
}

// EvalSHAReadOnly is a read-only variant of the EVALSHA command that cannot execute commands that modify data.
// See: https://redis.io/commands/evalsha_ro
func (sc Scripting) EvalSHAReadOnly(ctx context.Context, hash string, keys []string, args ...string) ([]string, error) {
	req := newRequestSize(3+len(keys)+len(args), "\r\n$9\r\nEVALSHA_RO\r\n$")
	req.addStringIntStrings(hash, int64(len(keys)), append(keys, args...))
	return sc.c.cmdStrings(ctx, req)
}

// EvalReadOnly is a read-only variant of the EVAL command that cannot execute commands that modify data.
// See: https://redis.io/commands/eval_ro
func (sc Scripting) EvalReadOnly(ctx context.Context, script string, keys []string, args ...string) ([]string, error) {
	req := newRequestSize(3+len(keys)+len(args), "\r\n$7\r\nEVAL_RO\r\n$")
	req.addStringIntStrings(script, int64(len(keys)), append(keys, args...))
	return sc.c.cmdStrings(ctx, req)
}

// Call invokes a function.
// See: https://redis.io/commands/fcall
func (sc Scripting) Call(ctx context.Context, script string, keys []string, args ...string) ([]string, error) {
	req := newRequestSize(3+len(keys)+len(args), "\r\n$5\r\nFCALL\r\n$")
	req.addStringIntStrings(script, int64(len(keys)), append(keys, args...))
	return sc.c.cmdStrings(ctx, req)
}

// CallReadOnly this is a read-only variant of the FCALL command that cannot execute commands that modify data.
// See: https://redis.io/commands/fcall_ro
func (sc Scripting) CallReadOnly(ctx context.Context, script string, keys []string, args ...string) ([]string, error) {
	req := newRequestSize(3+len(keys)+len(args), "\r\n$8\r\nFCALL_RO\r\n$")
	req.addStringIntStrings(script, int64(len(keys)), append(keys, args...))
	return sc.c.cmdStrings(ctx, req)
}

// FunctionDelete deletes a library and all its functions.
// See: https://redis.io/commands/function-delete
func (sc Scripting) FunctionDelete(ctx context.Context, libName string) error {
	req := newRequest("*3\r\n$8\r\nFUNCTION\r\n$6\r\nDELETE\r\n$")
	req.addString(libName)
	return sc.c.cmdSimple(ctx, req)
}

// FunctionDump returns the serialized payload of loaded libraries.
// You can restore the serialized payload later with the FUNCTION RESTORE command.
// See: https://redis.io/commands/function-dump
func (sc Scripting) FunctionDump(ctx context.Context) error {
	req := newRequest("*2\r\n$8\r\nFUNCTION\r\n$4\r\nDUMP\r\n$")
	return sc.c.cmdSimple(ctx, req)
}

// FunctionFlush deletes all the libraries.
// See: https://redis.io/commands/function-flush
func (sc Scripting) FunctionFlush(ctx context.Context, mode string) error {
	if mode != "ASYNC" && mode != "SYNC" {
		return fmt.Errorf("unknown script flush mode: %s", mode)
	}
	req := newRequest("*3\r\n$8\r\nFUNCTION\r\n$5\r\nFLUSH\r\n$")
	req.addString(mode)
	return sc.c.cmdSimple(ctx, req)
}

// FunctionKill kills a function that is currently executing.
// See: https://redis.io/commands/function-kill
func (sc Scripting) FunctionKill(ctx context.Context) error {
	req := newRequest("*2\r\n$8\r\nFUNCTION\r\n$4\r\nKILL\r\n$")
	return sc.c.cmdSimple(ctx, req)
}

// FunctionList returns information about the functions and libraries.
// See: https://redis.io/commands/function-list
func (sc Scripting) FunctionList(ctx context.Context) error {
	// FUNCTION LIST [LIBRARYNAME library-name-pattern] [WITHCODE]
	return nil
}

// FunctionLoad loads a library to Redis.
// See: https://redis.io/commands/function-load
func (sc Scripting) FunctionLoad(ctx context.Context, lib, script string) error {
	// FUNCTION LOAD engine-name library-name [REPLACE] [DESCRIPTION library-description] function-code
	return nil
}

// FunctionRestore restores libraries from the serialized payload.
// See: https://redis.io/commands/function-restore
func (sc Scripting) FunctionRestore(ctx context.Context) error {
	// FUNCTION RESTORE serialized-value [FLUSH|APPEND|REPLACE]
	return nil
}

// FunctionStats returns information about the function that's currently running and information about the available execution engines.
// See: https://redis.io/commands/function-stats
func (sc Scripting) FunctionStats(ctx context.Context) ([]string, error) {
	req := newRequest("*2\r\n$6\r\nSCRIPT\r\n$5\r\nSTATS\r\n$")
	return sc.c.cmdStrings(ctx, req)
}

// ScriptDebug sets the debug mode for subsequent scripts executed with EVAL.
// See: https://redis.io/commands/script-debug
func (sc Scripting) ScriptDebug(ctx context.Context, mode string) error {
	if mode != "YES" && mode != "SYNC" && mode != "NO" {
		return fmt.Errorf("unknown script debug mode: %s", mode)
	}
	req := newRequest("*3\r\n$6\r\nSCRIPT\r\n$5\r\nDEBUG\r\n$")
	req.addString(mode)
	return sc.c.cmdSimple(ctx, req)
}

// ScriptExists returns information about the existence of the scripts in the script cache.
// See: https://redis.io/commands/script-exists
func (sc Scripting) ScriptExists(ctx context.Context, hashes ...string) ([]bool, error) {
	req := newRequestSize(2+len(hashes), "\r\n$6\r\nSCRIPT\r\n$6\r\nEXISTS\r\n$")
	req.addStrings(hashes)
	ss, err := sc.c.cmdInts(ctx, req)
	if err != nil {
		return nil, err
	}
	res := make([]bool, len(ss))
	for i := range ss {
		res[i] = ss[i] == 1
	}
	return res, nil
}

// ScriptFlush flushes the Lua scripts cache.
// See: https://redis.io/commands/script-flush
func (sc Scripting) ScriptFlush(ctx context.Context, mode string) error {
	if mode != "ASYNC" && mode != "SYNC" {
		return fmt.Errorf("unknown script flush mode: %s", mode)
	}
	req := newRequest("*3\r\n$6\r\nSCRIPT\r\n$5\r\nFLUSH\r\n$")
	req.addString(mode)
	return sc.c.cmdSimple(ctx, req)
}

// ScriptKill kills the currently executing EVAL script, assuming no write operation was yet performed by the script.
// See: https://redis.io/commands/script-kill
func (sc Scripting) ScriptKill(ctx context.Context) error {
	req := newRequest("*2\r\n$6\r\nSCRIPT\r\n$4\r\nKILL\r\n$")
	return sc.c.cmdSimple(ctx, req)
}

// ScriptLoad loads a script into the scripts cache, without executing it.
// After the specified command is loaded into the script cache it will be callable using EVALSHA with the correct SHA1 digest of the script,
// exactly like after the first successful invocation of EVAL.
//
// See: https://redis.io/commands/script-load
func (sc Scripting) ScriptLoad(ctx context.Context, script string) (string, error) {
	req := newRequest("*3\r\n$6\r\nSCRIPT\r\n$4\r\nLOAD\r\n$")
	req.addString(script)
	return sc.c.cmdString(ctx, req)
}
