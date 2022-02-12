package redis

import (
	"context"
	"errors"
	"strconv"
	"testing"
)

func TestList_BadInput(t *testing.T) {
	ctx := newContext()
	list := NewList("list_bad_input", nil)

	_, err := list.Insert(ctx, "BETWEEN", "", "")
	mustEqual(t, err, errors.New("unknown insert mode: BETWEEN"))

	_, err = list.LeftMove(ctx, "", "CENTER", "LEFT")
	mustEqual(t, err, errors.New("unknown from: CENTER or to: LEFT"))

	_, err = list.LeftMove(ctx, "", "LEFT", "CENTER")
	mustEqual(t, err, errors.New("unknown from: LEFT or to: CENTER"))
}

func TestList_IndexLen(t *testing.T) {
	ctx := newContext()
	list := makeList(t, "list_index")

	size, err := list.LeftPush(ctx, "World")
	failIfErr(t, err)
	mustEqual(t, size, int64(1))

	size, err = list.LeftPush(ctx, "Hello")
	failIfErr(t, err)
	mustEqual(t, size, int64(2))

	got, err := list.Index(ctx, 0)
	failIfErr(t, err)
	mustEqual(t, got, "Hello")

	got, err = list.Index(ctx, -1)
	failIfErr(t, err)
	mustEqual(t, got, "World")

	got, err = list.Index(ctx, 3)
	failIfErr(t, err)
	mustEqual(t, got, "")

	size, err = list.Len(ctx)
	failIfErr(t, err)
	mustEqual(t, size, int64(2))
}

func TestList_Insert(t *testing.T) {
	ctx := newContext()
	list := makeList(t, "list_insert")

	size, err := list.RightPush(ctx, "Hello")
	failIfErr(t, err)
	mustEqual(t, size, int64(1))

	size, err = list.RightPush(ctx, "World")
	failIfErr(t, err)
	mustEqual(t, size, int64(2))

	size, err = list.Insert(ctx, "BEFORE", "World", "Three")
	failIfErr(t, err)
	mustEqual(t, size, int64(3))

	got, err := list.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, got, []string{"Hello", "Three", "World"})
}

func TestList_LeftMove(t *testing.T) {
	ctx := newContext()
	list := makeList(t, "list_leftmove")
	list2 := makeList(t, "list_leftmove_2")

	size, err := list.RightPush(ctx, "one", "two", "three")
	failIfErr(t, err)
	mustEqual(t, size, int64(3))

	got, err := list.LeftMove(ctx, list2.Name(), "RIGHT", "LEFT")
	failIfErr(t, err)
	mustEqual(t, got, "three")

	got, err = list.LeftMove(ctx, list2.Name(), "LEFT", "RIGHT")
	failIfErr(t, err)
	mustEqual(t, got, "one")

	all, err := list.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, all, []string{"two"})

	all, err = list2.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, all, []string{"three", "one"})
}

func TestList_RightPopLeftPush(t *testing.T) {
	ctx := newContext()
	list := makeList(t, "list_rplp")
	list2 := makeList(t, "list_rplp_2")

	size, err := list.RightPush(ctx, "one", "two", "three")
	failIfErr(t, err)
	mustEqual(t, size, int64(3))

	got, err := list.RightPopLeftPush(ctx, list2.Name())
	failIfErr(t, err)
	mustEqual(t, got, "three")

	all, err := list.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, all, []string{"one", "two"})

	all, err = list2.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, all, []string{"three"})
}

func TestList_Remove(t *testing.T) {
	ctx := newContext()
	list := makeList(t, "list_remove")

	size, err := list.RightPush(ctx, "Hello", "Hello", "foo", "Hello")
	failIfErr(t, err)
	mustEqual(t, size, int64(4))

	rem, err := list.Remove(ctx, -2, "Hello")
	failIfErr(t, err)
	mustEqual(t, rem, int64(2))

	got, err := list.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, got, []string{"Hello", "foo"})
}

func TestList_Set(t *testing.T) {
	ctx := newContext()
	list := makeList(t, "list_set")

	size, err := list.RightPush(ctx, "one", "two", "three")
	failIfErr(t, err)
	mustEqual(t, size, int64(3))

	err = list.Set(ctx, 0, "four")
	failIfErr(t, err)

	err = list.Set(ctx, -2, "five")
	failIfErr(t, err)

	got, err := list.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, got, []string{"four", "five", "three"})
}

func TestList_Trim(t *testing.T) {
	ctx := newContext()
	list := makeList(t, "list_trim")

	size, err := list.RightPush(ctx, "one", "two", "three")
	failIfErr(t, err)
	mustEqual(t, size, int64(3))

	err = list.Trim(ctx, 1, -1)
	failIfErr(t, err)

	got, err := list.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, got, []string{"two", "three"})
}

func TestList_PushX(t *testing.T) {
	ctx := newContext()
	list := makeList(t, "list_push_left")

	size, err := list.LeftPush(ctx, "World")
	failIfErr(t, err)
	mustEqual(t, size, int64(1))

	size, err = list.LeftPushX(ctx, "Hello")
	failIfErr(t, err)
	mustEqual(t, size, int64(2))

	list2 := NewList("list_push_left_x", testClient)
	size, err = list2.LeftPushX(ctx, "Hello")
	failIfErr(t, err)
	mustEqual(t, size, int64(0))

	got, err := list.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, got, []string{"Hello", "World"})

	got, err = list2.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, got, []string{})

	list = makeList(t, "list_push_right")

	size, err = list.RightPush(ctx, "Hello")
	failIfErr(t, err)
	mustEqual(t, size, int64(1))

	size, err = list.RightPushX(ctx, "World")
	failIfErr(t, err)
	mustEqual(t, size, int64(2))

	list2 = NewList("list_push_right_x", testClient)
	size, err = list2.RightPushX(ctx, "World")
	failIfErr(t, err)
	mustEqual(t, size, int64(0))

	got, err = list.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, got, []string{"Hello", "World"})

	got, err = list2.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, got, []string{})
}

func TestList_PopX(t *testing.T) {
	ctx := newContext()
	list := makeList(t, "list_left_pop")

	size, err := list.RightPush(ctx, "one", "two", "three", "four", "five")
	failIfErr(t, err)
	mustEqual(t, size, int64(5))

	got, err := list.LeftPop(ctx, 1)
	failIfErr(t, err)
	mustEqual(t, got, []string{"one"})

	got, err = list.LeftPop(ctx, 2)
	failIfErr(t, err)
	mustEqual(t, got, []string{"two", "three"})

	got, err = list.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, got, []string{"four", "five"})

	list = makeList(t, "list_right_pop")
	size, err = list.RightPush(ctx, "one", "two", "three", "four", "five")
	failIfErr(t, err)
	mustEqual(t, size, int64(5))

	got, err = list.RightPop(ctx, 1)
	failIfErr(t, err)
	mustEqual(t, got, []string{"five"})

	got, err = list.RightPop(ctx, 2)
	failIfErr(t, err)
	mustEqual(t, got, []string{"four", "three"})

	got, err = list.Range(ctx, 0, -1)
	failIfErr(t, err)
	mustEqual(t, got, []string{"one", "two"})
}

func BenchmarkList(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	ctx := context.Background()

	hll := makeHLL(b, "bench_list")
	for i := 0; i < b.N; i++ {
		_, err := hll.Add(ctx, strconv.Itoa(i), strconv.Itoa(i+1_000_000))
		failIfErr(b, err)

		_, err = hll.Count(ctx)
		failIfErr(b, err)
	}
}

func makeList(t testing.TB, name string) List {
	t.Helper()
	removeKey(t, name)
	t.Cleanup(func() { removeKey(t, name) })
	return NewList(name, testClient)
}
