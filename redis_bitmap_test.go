package redis

import (
	"testing"
)

func TestBitmap_GetSetPos(t *testing.T) {
	ctx := newContext()
	bm := makeBitMap(t, "bitmap_getset")

	val, err := bm.SetBit(ctx, 7, 1)
	failIfErr(t, err)
	mustEqual(t, val, int64(0))

	bit, err := bm.GetBit(ctx, 7)
	failIfErr(t, err)
	mustEqual(t, bit, int64(1))

	val, err = bm.SetBit(ctx, 7, 0)
	failIfErr(t, err)
	mustEqual(t, val, int64(1))

	// TODO(oleg): fix
	_, err = bm.BitPos(ctx, 1, 0)
	failIfErr(t, err)
}

func TestBitmap_Count(t *testing.T) {
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

func BenchmarkBitMap(b *testing.B) {
	ctx := newContext()
	bm := makeBitMap(b, "BitMap1")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := bm.SetBit(ctx, int64(i), i&1)
		failIfErr(b, err)

		bitCount, err := bm.BitCountAll(ctx)
		failIfErr(b, err)
		_ = bitCount
	}
}

func makeBitMap(t testing.TB, name string) *BitMap {
	t.Helper()
	removeKey(t, name)
	t.Cleanup(func() { removeKey(t, name) })
	return NewBitMap(name, testClient)
}
