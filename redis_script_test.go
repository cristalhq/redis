package redis

import "testing"

func TestScript_Eval(t *testing.T) {
	ctx := newContext()
	scr := makeScripting(t)

	script := "return {KEYS[1],ARGV[1]}"
	args := []string{"hello"}
	val, err := scr.Eval(ctx, script, []string{"key"}, args...)
	failIfErr(t, err)
	mustEqual(t, val, []string{"key", "hello"})

	// TODO(oleg): only for Redis 7.0
	// val, err = scr.EvalReadOnly(ctx, script, []string{"key"}, args...)
	// failIfErr(t, err)
	// mustEqual(t, val, []string{"key", "hello"})

	sha := "bfbf458525d6a0b19200bfd6db3af481156b367b"
	got, err := scr.ScriptLoad(ctx, script)
	failIfErr(t, err)
	mustEqual(t, got, sha)

	val, err = scr.EvalSHA(ctx, sha, []string{"key"}, args...)
	failIfErr(t, err)
	mustEqual(t, val, []string{"key", "hello"})
}

func TestScript_Function(t *testing.T) {
	t.Skip("only for Redis 7.0")
	ctx := newContext()
	scr := makeScripting(t)

	err := scr.FunctionLoad(ctx, "mylib", "redis.register_function('myfunc', function(keys, args) return args[1] end)")
	failIfErr(t, err)

	res, err := scr.Call(ctx, "myfunc", nil, "hello")
	failIfErr(t, err)
	mustEqual(t, res, []string{"hello"})

	res, err = scr.CallReadOnly(ctx, "myfunc", nil, "hello")
	failIfErr(t, err)
	mustEqual(t, res, []string{"hello"})

	scr.FunctionLoad(ctx, "mylib", "redis.register_function('myfunc', function(keys, args) return args[1] end)")
	// redis> FUNCTION LOAD Lua mylib "redis.register_function('myfunc', function(keys, args) return args[1] end)"
	// OK
	// redis> FCALL myfunc 0 hello
	// "hello"

	// redis> FUNCTION LOAD Lua mylib "redis.register_function('myfunc', function(keys, args) return 'hello' end)"
	// OK
	// redis> FCALL myfunc 0
	// "hello"
	// redis> FUNCTION DELETE mylib
	// OK
	// redis> FCALL myfunc 0
	// (error) ERR Function not found

	err = scr.FunctionDelete(ctx, "mylib")
	failIfErr(t, err)

	_, err = scr.Call(ctx, "myfunc", nil, "hello")
	// failIfErr(t, err)
}

func makeScripting(t testing.TB) Scripting {
	return NewScripting(testClient)
}
