package redis

import (
	"context"
	"testing"
)

func TestStream_Add(t *testing.T) {
	ctx := context.Background()
	st := makeStream(t, "stream_add")

	_, err := st.Add(ctx, map[string]string{"a": "1"})
	failIfErr(t, err)

	id, err := st.Add(ctx, map[string]string{"b": "2"})
	failIfErr(t, err)

	_, err = st.Add(ctx, map[string]string{"c": "3"})
	failIfErr(t, err)

	size, err := st.Len(ctx)
	failIfErr(t, err)
	mustEqual(t, size, int64(3))

	del, err := st.Delete(ctx, id)
	failIfErr(t, err)
	mustEqual(t, del, int64(1))

	msg, err := st.RangeAll(ctx, "-", "+")
	_ = err
	_ = msg
	// failIfErr(t, err)
	// mustEqual(t, len(msg), int(2))
	// 1) 1) 1538561698944-0
	// 2) 1) "a"
	// 	2) "1"
	// 2) 1) 1538561701744-0
	// 2) 1) "c"
	// 	2) "3"
}

func TestStream_Group(t *testing.T) {
	ctx := context.Background()
	st := makeStream(t, "stream_group")

	defer removeKey(t, "group1")
	err := st.GroupCreate(ctx, "group1", "0")
	failIfErr(t, err)

	cons, err := st.GroupCreateConsumer(ctx, "group1", "consumer1")
	failIfErr(t, err)
	mustEqual(t, cons, 1)

	cons, err = st.GroupDeleteConsumer(ctx, "group1", "consumer1")
	failIfErr(t, err)
	mustEqual(t, cons, 0)

	err = st.GroupSetID(ctx, "group1", "0")
	failIfErr(t, err)

	got, err := st.GroupDestroy(ctx, "group1")
	failIfErr(t, err)
	mustEqual(t, got, 1)
}

func makeStream(t testing.TB, name string) Stream {
	t.Helper()
	removeKey(t, name)
	t.Cleanup(func() { removeKey(t, name) })
	return NewStream(name, testClient)
}
