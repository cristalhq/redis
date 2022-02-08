package redis

import (
	"context"
	"testing"
)

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
}

func TestStrings_Count(t *testing.T) {
	ctx := newContext()
	bm := makeBitMap(t, "bitmap_getset")

	val, err := bm.BitCount(ctx, 7, 1)
	failIfErr(t, err)
	mustEqual(t, val, int64(0))

	// 	redis> SET mykey "foobar"
	// "OK"
	// redis> BITCOUNT mykey
	// (integer) 26
	// redis> BITCOUNT mykey 0 0
	// (integer) 4
	// redis> BITCOUNT mykey 1 1
	// (integer) 6
	// redis> BITCOUNT mykey 1 1 BYTE
	// (integer) 6
	// redis> BITCOUNT mykey 5 30 BIT
	// (integer) 17
	// redis>
}

func BenchmarkStrings(b *testing.B) {
	ctx := context.Background()
	str := makeStrings(b)
	key, val := "bench_str_key", "OK" //bench_str_val"
	err := str.Set(ctx, key, val)
	failIfErr(b, err)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// go func() {
		// got, err := str.Get(ctx, key)
		// failIfErr(b, err)
		// mustEqual(b, got, val)
		// }()
		// go func() {
		// 	got, err := str.Get(ctx, key)
		// 	failIfErr(b, err)
		// 	mustEqual(b, got, val)
		// }()

		err := str.Set(ctx, key, val)
		failIfErr(b, err)

		// _, err = str.Get(ctx, key)
		// failIfErr(b, err)
	}
}

func makeStrings(t testing.TB) Strings {
	t.Helper()
	// removeKey(t, name)
	// t.Cleanup(func() { removeKey(t, name) })
	return NewStrings(testClient)
}
