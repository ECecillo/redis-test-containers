package main

import (
	"context"
	"log"
	"redis-connection/example/redis"
)

// Just a simple setup to test if we can correctly ping, pong the redis server.
func main() {
	redisClient, err := redis.NewRedisClient("localhost:6379", "test")
	if err != nil {
		log.Fatal(err)
	}
	err = redisClient.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
