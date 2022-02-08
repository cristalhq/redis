package redis

import (
	"context"
	"log"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var testClient *Client

func init() {
	addr := "127.0.0.1:6379"

	c, err := NewClient(context.Background(), addr)
	if err != nil {
		log.Fatal(err)
	}
	testClient = c

	rand.Seed(time.Now().UnixNano())
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
