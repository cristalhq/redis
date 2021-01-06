package redis

import (
	"bytes"
	"math"
	"testing"
	"time"
)

func TestConnWriter(t *testing.T) {
	f := func(args []interface{}, want string) {
		var buf bytes.Buffer
		w := newWriter(&buf)

		if err := w.WriteArgs(args); err != nil {
			t.Error(err)
		}

		if err := w.Flush(); err != nil {
			t.Error(err)
		}

		if actual := buf.String(); actual != want {
			t.Errorf("%v, want %q, got %q", args, want, actual)
		}
	}

	f(
		[]interface{}{"SET", "key", "value"}, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n",
	)
	f(
		[]interface{}{"SET", "key", "value"}, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n",
	)
	f(
		[]interface{}{"SET", "key", "value"}, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n",
	)
	f(
		[]interface{}{"SET", "key", byte(100)}, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$3\r\n100\r\n",
	)
	f(
		[]interface{}{"SET", "key", 100}, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$3\r\n100\r\n",
	)
	f(
		[]interface{}{"SET", "key", int64(math.MinInt64)}, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$20\r\n-9223372036854775808\r\n",
	)
	f(
		[]interface{}{"SET", "key", 1337.31337}, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$10\r\n1337.31337\r\n",
	)
	f(
		[]interface{}{"SET", "key", float64(133731337.133731337)}, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$18\r\n133731337.13373134\r\n",
	)
	f(
		[]interface{}{"SET", "key", ""}, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$0\r\n\r\n",
	)
	f(
		[]interface{}{"SET", "key", nil}, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$0\r\n\r\n",
	)
	f(
		[]interface{}{"SET", "key", 2*time.Hour + 12*time.Minute + 42*time.Second}, "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$13\r\n7962000000000\r\n",
	)
	f(
		[]interface{}{"ECHO", true, false}, "*3\r\n$4\r\nECHO\r\n$1\r\n1\r\n$1\r\n0\r\n",
	)
}
