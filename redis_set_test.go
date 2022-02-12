package redis

import (
	"context"
	"sort"
	"testing"
)

func TestSet_Add(t *testing.T) {
	ctx := newContext()
	set := makeSet(t, "set_add")

	added, err := set.Add(ctx, "Hello")
	failIfErr(t, err)
	mustEqual(t, added, int64(1))

	added, err = set.Add(ctx, "World")
	failIfErr(t, err)
	mustEqual(t, added, int64(1))

	added, err = set.Add(ctx, "World")
	failIfErr(t, err)
	mustEqual(t, added, int64(0))

	size, err := set.Cardinality(ctx)
	failIfErr(t, err)
	mustEqual(t, size, int64(2))

	ms, err := set.Members(ctx)
	failIfErr(t, err)
	sort.Strings(ms)
	mustEqual(t, ms, []string{"Hello", "World"})

	is, err := set.IsMember(ctx, "Hello")
	failIfErr(t, err)
	mustEqual(t, is, true)

	is, err = set.IsMember(ctx, "Goodbye")
	failIfErr(t, err)
	mustEqual(t, is, false)

	iss, err := set.MultiIsMembers(ctx, "Hello", "Goodbuy", "World")
	failIfErr(t, err)
	mustEqual(t, iss, []bool{true, false, true})
}

func TestSet_RandomMember(t *testing.T) {
	ctx := newContext()
	set := makeSet(t, "set_rand")

	added, err := set.Add(ctx, "one", "two", "three")
	failIfErr(t, err)
	mustEqual(t, added, int64(3))

	m, err := set.RandomMember(ctx)
	failIfErr(t, err)
	if m != "one" && m != "two" && m != "three" {
		t.Fatalf("got %q", m)
	}

	ms, err := set.RandomMembers(ctx, 2)
	failIfErr(t, err)
	mustEqual(t, len(ms), int(2))

	// TODO: must be -5
	// ms, err = set.RandomMembers(ctx, 5)
	// failIfErr(t, err)
	// mustEqual(t, len(ms), int(2))
}

func TestSet_MovePopRemove(t *testing.T) {
	ctx := newContext()
	set := makeSet(t, "set_move")
	set2 := makeSet(t, "otherset_move")

	added, err := set.Add(ctx, "one", "two")
	failIfErr(t, err)
	mustEqual(t, added, int64(2))

	added, err = set2.Add(ctx, "three")
	failIfErr(t, err)
	mustEqual(t, added, int64(1))

	_, err = set.MoveTo(ctx, set2.Name(), "two")
	failIfErr(t, err)
	mustEqual(t, added, int64(1))

	ms, err := set.Members(ctx)
	failIfErr(t, err)
	mustEqual(t, ms, []string{"one"})

	ms, err = set2.Members(ctx)
	failIfErr(t, err)
	sort.Strings(ms)
	mustEqual(t, ms, []string{"three", "two"})

	pop, err := set2.Pop(ctx)
	failIfErr(t, err)
	if pop != "two" && pop != "three" {
		t.Fatalf("got %s", pop)
	}

	rem, err := set.Remove(ctx, "one", "10")
	failIfErr(t, err)
	mustEqual(t, rem, int64(1))

	added, err = set.Add(ctx, "one", "two", "three", "four")
	failIfErr(t, err)

	ps, err := set.Pops(ctx, 2)
	failIfErr(t, err)
	mustEqual(t, len(ps), int(2))
}

func TestSet_Diff(t *testing.T) {
	ctx := newContext()
	set1 := makeSet(t, "set_diff1")
	set2 := makeSet(t, "set_diff2")

	added, err := set1.Add(ctx, "a", "b", "c")
	failIfErr(t, err)
	mustEqual(t, added, int64(3))
	added, err = set2.Add(ctx, "c", "d", "e")
	failIfErr(t, err)
	mustEqual(t, added, int64(3))

	defer removeKey(t, "set_diff3")
	total, err := set1.DiffStore(ctx, "set_diff3", set2.Name())
	failIfErr(t, err)
	mustEqual(t, total, int64(2))

	set3 := NewSet("set_diff3", testClient)
	res, err := set3.Members(ctx)
	failIfErr(t, err)
	sort.Strings(res)
	mustEqual(t, res, []string{"a", "b"})

	res, err = set1.Diff(ctx, set2.Name())
	failIfErr(t, err)
	sort.Strings(res)
	mustEqual(t, res, []string{"a", "b"})
}

func TestSet_Inter(t *testing.T) {
	ctx := newContext()
	set1 := makeSet(t, "set_inter1")
	set2 := makeSet(t, "set_inter2")

	added, err := set1.Add(ctx, "a", "b", "c")
	failIfErr(t, err)
	mustEqual(t, added, int64(3))
	added, err = set2.Add(ctx, "c", "d", "e")
	failIfErr(t, err)
	mustEqual(t, added, int64(3))

	defer removeKey(t, "set_inter3")
	total, err := set1.InterStore(ctx, "set_inter3", set2.Name())
	failIfErr(t, err)
	mustEqual(t, total, int64(1))

	set3 := NewSet("set_inter3", testClient)
	res, err := set3.Members(ctx)
	failIfErr(t, err)
	sort.Strings(res)
	mustEqual(t, res, []string{"c"})

	res, err = set1.Inter(ctx, set2.Name())
	failIfErr(t, err)
	sort.Strings(res)
	mustEqual(t, res, []string{"c"})
}

func TestSet_Union(t *testing.T) {
	ctx := newContext()
	set1 := makeSet(t, "set_union1")
	set2 := makeSet(t, "set_union2")

	added, err := set1.Add(ctx, "a", "b", "c")
	failIfErr(t, err)
	mustEqual(t, added, int64(3))
	added, err = set2.Add(ctx, "c", "d", "e")
	failIfErr(t, err)
	mustEqual(t, added, int64(3))

	defer removeKey(t, "set_union3")
	total, err := set1.UnionStore(ctx, "set_union3", set2.Name())
	failIfErr(t, err)
	mustEqual(t, total, int64(5))

	set3 := NewSet("set_union3", testClient)
	res, err := set3.Members(ctx)
	failIfErr(t, err)
	sort.Strings(res)
	mustEqual(t, res, []string{"a", "b", "c", "d", "e"})

	res, err = set1.Union(ctx, set2.Name())
	failIfErr(t, err)
	sort.Strings(res)
	mustEqual(t, res, []string{"a", "b", "c", "d", "e"})
}

func BenchmarkSet_Add(b *testing.B) {
	ctx := context.Background()
	set := makeSet(b, "bench_set")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := set.Add(ctx, "1", "2", "3")
		failIfErr(b, err)

		count, err := set.Cardinality(ctx)
		failIfErr(b, err)
		if count != 3 {
			b.Fatalf("got %d", count)
		}

		_, err = set.IsMember(ctx, "1")
		failIfErr(b, err)
	}
}

func makeSet(t testing.TB, name string) Set {
	t.Helper()
	removeKey(t, name)
	t.Cleanup(func() { removeKey(t, name) })
	return NewSet(name, testClient)
}
