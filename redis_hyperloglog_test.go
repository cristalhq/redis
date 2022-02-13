package redis

import (
	"context"
	"strconv"
	"testing"
)

func TestHyperLogLog_Add(t *testing.T) {
	ctx := context.Background()
	hll := makeHLL(t, "hll_add")

	added, err := hll.Add(ctx, "a", "b", "c", "d", "e", "f", "g")
	failIfErr(t, err)
	mustEqual(t, added, int64(1))

	count, err := hll.Count(ctx)
	failIfErr(t, err)
	mustEqual(t, count, int64(7))
}

func TestHyperLogLog_Count(t *testing.T) {
	ctx := context.Background()
	hll := makeHLL(t, "hll_count")
	hllOther := makeHLL(t, "hll-other_count")

	added, err := hll.Add(ctx, "foo", "bar", "zap")
	failIfErr(t, err)
	mustEqual(t, added, int64(1))

	added, err = hll.Add(ctx, "zap", "zap", "zap")
	failIfErr(t, err)
	mustEqual(t, added, int64(0))

	added, err = hll.Add(ctx, "foo", "bar")
	failIfErr(t, err)
	mustEqual(t, added, int64(0))

	count, err := hll.Count(ctx)
	failIfErr(t, err)
	mustEqual(t, count, int64(3))

	added, err = hllOther.Add(ctx, "1", "2", "3")
	failIfErr(t, err)
	mustEqual(t, added, int64(1))

	count, err = hll.CountWith(ctx, "hll-other_count")
	failIfErr(t, err)
	mustEqual(t, count, int64(6))
}

func TestHyperLogLog_Merge(t *testing.T) {
	ctx := context.Background()
	hll1 := makeHLL(t, "hll_merge1")
	hll2 := makeHLL(t, "hll_merge2")

	added, err := hll1.Add(ctx, "foo", "bar", "zap", "a")
	failIfErr(t, err)
	mustEqual(t, added, int64(1))

	added, err = hll2.Add(ctx, "a", "b", "c", "foo")
	failIfErr(t, err)
	mustEqual(t, added, int64(1))

	err = hll1.MergeInto(ctx, "hll_merge3", "hll_merge1", "hll_merge2")
	failIfErr(t, err)

	hll3 := NewHyperLogLog("hll_merge3", testClient)
	count, err := hll3.Count(ctx)
	failIfErr(t, err)
	mustEqual(t, count, int64(6))

	err = hll1.Merge(ctx, "hll_merge2")
	failIfErr(t, err)
	mustEqual(t, count, int64(6))
}

func BenchmarkHyperLogLog(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	ctx := context.Background()

	hll := makeHLL(b, "bench_hll")
	for i := 0; i < b.N; i++ {
		_, err := hll.Add(ctx, strconv.Itoa(i), strconv.Itoa(i+1_000_000))
		failIfErr(b, err)

		_, err = hll.Count(ctx)
		failIfErr(b, err)
	}
}

func makeHLL(t testing.TB, name string) HyperLogLog {
	t.Helper()
	removeKey(t, name)
	t.Cleanup(func() { removeKey(t, name) })
	return NewHyperLogLog(name, testClient)
}
