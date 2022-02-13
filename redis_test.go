package redis

import (
	"context"
	"log"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"
)

var testClient *Client

func init() {
	redisAddr := os.Getenv("TEST_REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "127.0.0.1:6379"
	}

	c, err := NewClient(context.Background(), redisAddr)
	if err != nil {
		log.Fatal(err)
	}
	testClient = c

	rand.Seed(time.Now().UnixNano())
}

func TestNames(t *testing.T) {
	mustEqual(t, NewBitMap("bitmap", nil).Name(), "bitmap")
	mustEqual(t, NewList("list", nil).Name(), "list")
	mustEqual(t, NewSet("set", nil).Name(), "set")
}

func TestNotImplemented(t *testing.T) {
	ctx := newContext()

	f := func(fn func()) {
		t.Helper()
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("should fail")
			}
		}()
		fn()
	}

	bm := NewBitMap("bitmap_not_implemented", nil)
	f(func() { bm.BitField(ctx) })
	f(func() { bm.BitFieldReadOnly(ctx) })

	list := NewList("list_not_implemented", nil)
	f(func() { list.BlockingLeftMove(ctx) })
	f(func() { list.BlockingLeftMultiPop(ctx) })
	f(func() { list.BlockingLeftPop(ctx) })
	f(func() { list.BlockingRightPop(ctx) })
	f(func() { list.BlockingRightPopLeftPush(ctx) })
	f(func() { list.LeftMultiPop(ctx) })
	f(func() { list.LeftPos(ctx) })

	set := NewSet("set_not_implemented", nil)
	f(func() { set.Scan(ctx) })

	str := NewStrings(nil)
	f(func() { str.LCS(ctx, "key1", "key2") })
}

func newContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_ = cancel // ok to skip
	return ctx
}

func removeKey(t testing.TB, key string) {
	t.Helper()

	req := newRequest("*2\r\n$3\r\nDEL\r\n$")
	req.addString(key)
	_, err := testClient.cmdInt(context.Background(), req)
	failIfErr(t, err)
}

func failIfErr(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func mustEqual(t testing.TB, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
