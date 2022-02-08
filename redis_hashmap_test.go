package redis

import (
	"testing"
)

func TestHashMap_MultiGetSet(t *testing.T) {
	ctx := newContext()
	hmap := makeHashMap(t, "hmap_hmget")

	err := hmap.Set(ctx, "field1", "Hello")
	failIfErr(t, err)
	err = hmap.Set(ctx, "field2", "World")
	failIfErr(t, err)

	vals, err := hmap.MultiGet(ctx, "field1", "field2", "nofield")
	failIfErr(t, err)
	mustEqual(t, vals, []string{"Hello", "World", ""})

	hmap = makeHashMap(t, "hmap_hmset")
	err = hmap.MultiSet(ctx, map[string]string{"field1": "Hello", "field2": "World"})
	failIfErr(t, err)

	got, err := hmap.Get(ctx, "field1")
	failIfErr(t, err)
	mustEqual(t, got, "Hello")

	got, err = hmap.Get(ctx, "field2")
	failIfErr(t, err)
	mustEqual(t, got, "World")
}

func makeHashMap(t testing.TB, name string) HashMap {
	t.Helper()
	removeKey(t, name)
	t.Cleanup(func() { removeKey(t, name) })
	return NewHashMap(name, testClient)
}
