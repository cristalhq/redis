package redis

import (
	"context"
	"testing"
)

func TestStrings_BadInput(t *testing.T) {
}

func TestStrings_GetSet(t *testing.T) {
	ctx := newContext()
	removeKey(t, "key")
	str := makeStrings(t)

	val, err := str.Get(ctx, "key")
	failIfErr(t, err)
	mustEqual(t, val, "")

	err = str.Set(ctx, "key", "value")
	failIfErr(t, err)

	val, err = str.Get(ctx, "key")
	failIfErr(t, err)
	mustEqual(t, val, "value")

	val, err = str.GetDel(ctx, "key")
	failIfErr(t, err)
	mustEqual(t, val, "value")

	size, err := str.Append(ctx, "key", "app")
	failIfErr(t, err)
	mustEqual(t, size, int64(3))

	// GetDel
	// 	redis> SET mykey "Hello"
	// "OK"
	// redis> GETDEL mykey
	// ERR Unknown or disabled command 'GETDEL'
	// redis> GET mykey
	// "Hello"
	// redis>

}

func TestStrings_Multi(t *testing.T) {
	ctx := newContext()
	removeKey(t, "key1")
	removeKey(t, "key2")
	str := makeStrings(t)

	err := str.MultiSet(ctx, "key1", "Hello", "key2", "World")
	failIfErr(t, err)

	val1, err := str.Get(ctx, "key1")
	failIfErr(t, err)
	mustEqual(t, val1, "Hello")

	val2, err := str.Get(ctx, "key2")
	failIfErr(t, err)
	mustEqual(t, val2, "World")
}

func TestStrings_MultiNotExist(t *testing.T) {
	ctx := newContext()
	removeKey(t, "key1")
	removeKey(t, "key2")
	removeKey(t, "key3")
	str := makeStrings(t)

	got, err := str.MultiSetNotExist(ctx, "key1", "Hello", "key2", "there")
	failIfErr(t, err)
	mustEqual(t, got, int64(1))

	got, err = str.MultiSetNotExist(ctx, "key2", "new", "key3", "world")
	failIfErr(t, err)
	mustEqual(t, got, int64(0))

	res, err := str.MultiGet(ctx, "key1", "key2", "key3")
	failIfErr(t, err)
	mustEqual(t, res, []string{"Hello", "there", ""})
}

func BenchmarkStrings(b *testing.B) {
	ctx := context.Background()
	str := makeStrings(b)
	key, val := "bench_str_key", "bench_str_val"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := str.Set(ctx, key, val)
		failIfErr(b, err)

		_, err = str.Get(ctx, key)
		failIfErr(b, err)
	}
}

func makeStrings(t testing.TB) Strings {
	t.Helper()
	return NewStrings(testClient)
}
