package redis

import (
	"encoding/base64"
	"sort"
	"testing"
)

func TestKeys(t *testing.T) {
	ctx := newContext()
	removeKey(t, "keys_key1")
	removeKey(t, "keys_key2")
	keys := NewKeys(testClient)

	err := keys.Set(ctx, "keys_key1", "Hello")
	failIfErr(t, err)

	err = keys.Set(ctx, "keys_key2", "World")
	failIfErr(t, err)

	// TODO(oleg): fix result
	// typ, err := keys.Type(ctx, "keys_key2")
	// failIfErr(t, err)
	// mustEqual(t, typ, "string")

	got, err := keys.Keys(ctx, "keys_key*")
	failIfErr(t, err)
	sort.Strings(got)
	mustEqual(t, got, []string{"keys_key1", "keys_key2"})

	del, err := keys.Delete(ctx, "keys_key1", "keys_key2", "key3")
	failIfErr(t, err)
	mustEqual(t, del, int64(del))

}

func TestKeys_CopyDumpRename(t *testing.T) {
	ctx := newContext()
	removeKey(t, "keys_copy")
	removeKey(t, "keys_dump")
	removeKey(t, "keys_dump")
	removeKey(t, "keys_tmp")
	keys := NewKeys(testClient)

	err := keys.Set(ctx, "keys_copy", "value")
	failIfErr(t, err)

	err = keys.Rename(ctx, "keys_copy", "keys_tmp")
	failIfErr(t, err)

	ok, err := keys.RenameNotExist(ctx, "keys_tmp", "keys_dump")
	failIfErr(t, err)
	mustEqual(t, ok, true)

	val, err := keys.Dump(ctx, "keys_dump")
	failIfErr(t, err)
	encVal := base64.RawURLEncoding.EncodeToString([]byte(val))
	mustEqual(t, encVal, "AAV2YWx1ZQoAgfbBOMGqm24")

	ok, err = keys.Move(ctx, "keys_dump", 1)
	failIfErr(t, err)
	// TODO(oleg): cleanup another DB
	// mustEqual(t, ok, true)
}
