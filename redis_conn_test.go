package redis

import "testing"

func TestConn_Ping(t *testing.T) {
	ctx := newContext()
	cs := NewClients(testClient)

	s, err := cs.Ping(ctx, "")
	// failIfErr(t, err)
	// mustEqual(t, s, "PONG")

	s, err = cs.Ping(ctx, "HELLO")
	failIfErr(t, err)
	mustEqual(t, s, "HELLO")

	s, err = cs.Echo(ctx, "HELLO WORLD")
	failIfErr(t, err)
	mustEqual(t, s, "HELLO WORLD")
}
