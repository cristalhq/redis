package redis

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	redisTestClient *redisClient
	testClient      *Client
)

func init() {
	addr, ok := os.LookupEnv("TEST_REDIS_ADDR")
	if !ok {
		addr = ":6379"
	}
	var err error
	redisTestClient = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	testClient, err = NewClient(redisTestClient)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
}

func randomKey(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, rand.Uint64())
}
