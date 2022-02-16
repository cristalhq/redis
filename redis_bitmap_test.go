package redis

import (
	"errors"
	"testing"
)

func TestBitMap_BadInput(t *testing.T) {
	ctx := newContext()
	bm := NewBitMap("bitmap_bad_input", nil)

	_, err := bm.BitOp(ctx, BitMapOp("WAT"), "", "", "")
	mustEqual(t, err, errors.New("unknown BitMap operation: WAT"))

	_, err = bm.BitOp(ctx, BitMapOp("NOT"), "dst", "key1", "key2")
	mustEqual(t, err, errors.New("BitMap Not operation works only with 1 key"))
}

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
	bm := makeBitMap(t, "bitmap_count")
	err := NewStrings(testClient).Set(ctx, bm.Name(), "foobar")
	failIfErr(t, err)

	count, err := bm.BitCountAll(ctx)
	failIfErr(t, err)
	mustEqual(t, count, int64(26))

	count, err = bm.BitCount(ctx, 0, 0)
	failIfErr(t, err)
	mustEqual(t, count, int64(4))

	count, err = bm.BitCount(ctx, 1, 1)
	failIfErr(t, err)
	mustEqual(t, count, int64(6))

	count, err = bm.BitCountByte(ctx, 1, 1)
	failIfErr(t, err)
	mustEqual(t, count, int64(6))

	// TODO(oleg): fix
	// count, err = bm.BitCount(ctx, 5, 30)
	// failIfErr(t, err)
	// mustEqual(t, count, int64(17))
}

func TestBitmap_Op(t *testing.T) {
	ctx := newContext()
	str := NewStrings(testClient)

	bm1 := makeBitMap(t, "bitmap_op1")
	bm2 := makeBitMap(t, "bitmap_op2")
	err := str.Set(ctx, bm1.Name(), "foobar")
	failIfErr(t, err)
	err = str.Set(ctx, bm2.Name(), "abcdef")
	failIfErr(t, err)

	res, err := bm1.BitOp(ctx, AndOp, "bitmap_and", bm1.Name(), bm2.Name())
	failIfErr(t, err)
	mustEqual(t, res, int64(6))

	res, err = bm1.BitOp(ctx, OrOp, "bitmap_or", bm1.Name(), bm2.Name())
	failIfErr(t, err)
	mustEqual(t, res, int64(6))

	res, err = bm1.BitOp(ctx, XorOp, "bitmap_xor", bm1.Name(), bm2.Name())
	failIfErr(t, err)
	mustEqual(t, res, int64(6))

	// TODO(oleg): fix
	val, err := str.Get(ctx, "bitmap_and")
	_ = val
	failIfErr(t, err)
	// mustEqual(t, val, "`bc`ab")

	val, err = str.Get(ctx, "bitmap_or")
	failIfErr(t, err)
	// mustEqual(t, val, "`bc`ab")

	val, err = str.Get(ctx, "bitmap_xor")
	failIfErr(t, err)
	// mustEqual(t, val, "`bc`ab")
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

func makeBitMap(t testing.TB, name string) BitMap {
	t.Helper()
	removeKey(t, name)
	t.Cleanup(func() { removeKey(t, name) })
	return NewBitMap(name, testClient)
}
