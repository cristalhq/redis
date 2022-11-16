package redis

import (
	"context"
	"strconv"
	"testing"
)

func TestSortedSet_BadInput(t *testing.T) {
	// ctx := newContext()
	// sset := NewList("sortedset_bad_input", nil)
}

func TestSortedSet_Score(t *testing.T) {
	t.Skip()
	ctx := newContext()
	sset := NewSortedSet("sortedset_score", nil)

	size, err := sset.Add(ctx)
	failIfErr(t, err)
	mustEqual(t, size, int64(1))

	score, err := sset.Score(ctx, "one")
	failIfErr(t, err)
	mustEqual(t, score, int64(1))

	size, err = sset.Cardinality(ctx)
	failIfErr(t, err)
	mustEqual(t, size, int64(1))

	// redis> ZADD myzset 1 "one"
	// (integer) 1
	// redis> ZSCORE myzset "one"
	// "1"
	// redis>
}

func BenchmarkSortedSet(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	ctx := context.Background()

	hll := makeHLL(b, "bench_sortedset")
	for i := 0; i < b.N; i++ {
		_, err := hll.Add(ctx, strconv.Itoa(i), strconv.Itoa(i+1_000_000))
		failIfErr(b, err)

		_, err = hll.Count(ctx)
		failIfErr(b, err)
	}
}

func makeSortedSet(t testing.TB, name string) SortedSet {
	t.Helper()
	removeKey(t, name)
	t.Cleanup(func() { removeKey(t, name) })
	return NewSortedSet(name, testClient)
}
