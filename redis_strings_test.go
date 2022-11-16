package redis

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestStrings_BadInput(t *testing.T) {
	ctx := newContext()
	str := NewStrings(nil)

	err := str.MultiSet(ctx, "alone")
	mustEqual(t, err, errors.New("one of the keys does not have a value"))

	_, err = str.MultiSetNotExist(ctx, "alone")
	mustEqual(t, err, errors.New("one of the keys does not have a value"))
}

func TestStrings(t *testing.T) {
	ctx := newContext()
	removeKey(t, "str_getset")
	str := makeStrings(t)

	val, err := str.Get(ctx, "str_getset")
	failIfErr(t, err)
	mustEqual(t, val, "")

	err = str.Set(ctx, "str_getset", "value")
	failIfErr(t, err)

	val, err = str.Get(ctx, "str_getset")
	failIfErr(t, err)
	mustEqual(t, val, "value")

	val, err = str.GetDel(ctx, "str_getset")
	failIfErr(t, err)
	mustEqual(t, val, "value")

	size, err := str.Append(ctx, "str_getset", "app")
	failIfErr(t, err)
	mustEqual(t, size, int64(3))

	val, err = str.GetSet(ctx, "str_getset", "new-val-ue")
	failIfErr(t, err)
	mustEqual(t, val, "app")

	size, err = str.Strlen(ctx, "str_getset")
	failIfErr(t, err)
	mustEqual(t, size, int64(10))

	val, err = str.Substr(ctx, "str_getset", 0, 2)
	failIfErr(t, err)
	mustEqual(t, val, "new")

	b, err := str.SetNotExist(ctx, "str_getset", "newnewnew")
	failIfErr(t, err)
	mustEqual(t, b, false)
}

func TestStrings_Expire(t *testing.T) {
	ctx := newContext()
	removeKey(t, "str_expire")
	str := makeStrings(t)

	{
		err := str.SetExpire(ctx, 20*time.Millisecond, "str_expire", "value1")
		failIfErr(t, err)

		d, _, err := NewKeys(testClient).TTL(ctx, "str_expire")
		failIfErr(t, err)
		if d > time.Duration(20)*time.Millisecond {
			t.Fatalf("too big: %v", d)
		}
		if d < time.Duration(15)*time.Millisecond {
			t.Fatalf("too low: %v", d)
		}

		time.Sleep(30 * time.Millisecond)

		val, err := str.Get(ctx, "str_expire")
		failIfErr(t, err)
		mustEqual(t, val, "")
	}

	{
		err := str.SetExpire(ctx, 20*time.Millisecond, "str_expire", "value2")
		failIfErr(t, err)

		val, err := str.GetPersist(ctx, "str_expire")
		failIfErr(t, err)
		mustEqual(t, val, "value2")

		val, err = str.GetExpire(ctx, "str_expire", 20*time.Millisecond)
		failIfErr(t, err)
		mustEqual(t, val, "value2")

		time.Sleep(30 * time.Millisecond)

		val, err = str.Get(ctx, "str_expire")
		failIfErr(t, err)
		mustEqual(t, val, "")
	}

	{
		err := str.Set(ctx, "str_expire", "value3")
		failIfErr(t, err)

		val, err := str.GetExpireAt(ctx, "str_expire", time.Now().Add(10*time.Millisecond))
		failIfErr(t, err)
		mustEqual(t, val, "value3")

		time.Sleep(100 * time.Millisecond)

		val, err = str.Get(ctx, "str_expire")
		failIfErr(t, err)
		// TODO(oleg): fix
		// mustEqual(t, val, "")
	}
}

func TestStrings_IncDec(t *testing.T) {
	ctx := newContext()
	removeKey(t, "str_count")
	str := makeStrings(t)

	val, err := str.Inc(ctx, "str_count")
	failIfErr(t, err)
	mustEqual(t, val, int64(1))

	val, err = str.IncBy(ctx, "str_count", 41)
	failIfErr(t, err)
	mustEqual(t, val, int64(42))

	val, err = str.Dec(ctx, "str_count")
	failIfErr(t, err)
	mustEqual(t, val, int64(41))

	val, err = str.DecBy(ctx, "str_count", 20)
	failIfErr(t, err)
	mustEqual(t, val, int64(21))

	f, err := str.IncByFloat(ctx, "str_count", 0.001)
	failIfErr(t, err)
	mustEqual(t, f, float64(21.00))
}

func TestStrings_Multi(t *testing.T) {
	ctx := newContext()
	removeKey(t, "str_multi1")
	removeKey(t, "str_multi2")
	str := makeStrings(t)

	err := str.MultiSet(ctx, "str_multi1", "Hello", "str_multi2", "World")
	failIfErr(t, err)

	val1, err := str.Get(ctx, "str_multi1")
	failIfErr(t, err)
	mustEqual(t, val1, "Hello")

	val2, err := str.Get(ctx, "str_multi2")
	failIfErr(t, err)
	mustEqual(t, val2, "World")
}

func TestStrings_MultiNotExist(t *testing.T) {
	ctx := newContext()
	removeKey(t, "str_multix1")
	removeKey(t, "str_multix2")
	removeKey(t, "key3")
	str := makeStrings(t)

	got, err := str.MultiSetNotExist(ctx, "str_multix1", "Hello", "str_multix2", "there")
	failIfErr(t, err)
	mustEqual(t, got, int64(1))

	got, err = str.MultiSetNotExist(ctx, "str_multix2", "new", "key3", "world")
	failIfErr(t, err)
	mustEqual(t, got, int64(0))

	res, err := str.MultiGet(ctx, "str_multix1", "str_multix2", "key3")
	failIfErr(t, err)
	mustEqual(t, res, []string{"Hello", "there", ""})
}

func TestStrings_Range(t *testing.T) {
	ctx := newContext()
	removeKey(t, "str_range")
	str := makeStrings(t)

	err := str.Set(ctx, "str_range", "Hello World")
	failIfErr(t, err)

	size, err := str.SetRange(ctx, "str_range", 6, "Redis")
	failIfErr(t, err)
	mustEqual(t, size, int64(11))

	got, err := str.Get(ctx, "str_range")
	failIfErr(t, err)
	mustEqual(t, got, "Hello Redis")

	got, err = str.GetRange(ctx, "str_range", 0, 3)
	failIfErr(t, err)
	mustEqual(t, got, "Hell")
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
