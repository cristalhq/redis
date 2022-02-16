package redis

import (
	"sort"
	"testing"
)

func TestHashMap(t *testing.T) {
	ctx := newContext()
	hmap := makeHashMap(t, "hmap_crud")

	err := hmap.Set(ctx, "field", "value")
	failIfErr(t, err)

	err = hmap.Set(ctx, "field2", "value2")
	failIfErr(t, err)

	val, err := hmap.Get(ctx, "field")
	failIfErr(t, err)
	mustEqual(t, val, "value")

	ok, err := hmap.Exists(ctx, "field")
	failIfErr(t, err)
	mustEqual(t, ok, true)

	all, err := hmap.GetAll(ctx)
	failIfErr(t, err)
	mustEqual(t, all, map[string]string{"field": "value", "field2": "value2"})

	ok, err = hmap.SetNotExist(ctx, "fielddd", "valueee")
	failIfErr(t, err)
	mustEqual(t, ok, true)

	size, err := hmap.Len(ctx)
	failIfErr(t, err)
	mustEqual(t, size, int64(3))

	size, err = hmap.Strlen(ctx, "field")
	failIfErr(t, err)
	mustEqual(t, size, int64(5))

	got, err := hmap.Delete(ctx, "field", "no-field")
	failIfErr(t, err)
	mustEqual(t, got, int64(1))
}

func TestHashMap_KV(t *testing.T) {
	ctx := newContext()
	hmap := makeHashMap(t, "hmap_kv")

	err := hmap.MultiSet(ctx, map[string]string{
		"field1": "value1",
		"field2": "value2",
		"field3": "value3"})
	failIfErr(t, err)

	ks, err := hmap.Keys(ctx)
	failIfErr(t, err)
	sort.Strings(ks)
	mustEqual(t, ks, []string{"field1", "field2", "field3"})

	vs, err := hmap.Values(ctx)
	failIfErr(t, err)
	sort.Strings(vs)
	mustEqual(t, vs, []string{"value1", "value2", "value3"})
}

func TestHashMap_Inc(t *testing.T) {
	ctx := newContext()
	hmap := makeHashMap(t, "hmap_inc")

	val, err := hmap.IncBy(ctx, "field", 10)
	failIfErr(t, err)
	mustEqual(t, val, int64(10))

	val, err = hmap.IncBy(ctx, "field", -7)
	failIfErr(t, err)
	mustEqual(t, val, int64(3))

	f, err := hmap.IncByFloat(ctx, "field", 39.42)
	failIfErr(t, err)
	mustEqual(t, f, float64(42.00))
}

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
