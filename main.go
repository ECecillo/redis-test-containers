package main

import (
	"context"
	"log"
	"redis-connection/example/internals/clickhouse"
	"redis-connection/example/internals/redis"

	"go.uber.org/zap"
)

// Just a simple setup to test if we can correctly ping, pong the redis server.
func main() {
	logger := zap.L().Named("root")

	redisClient, err := redis.NewRedisClient(logger, "localhost:6379", "test")
	if err != nil {
		log.Fatal(err)
	}
	err = redisClient.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	clickhouseConfig := clickhouse.Config{
		Host:     "127.0.0.1",
		Port:     9000,
		Database: "default",
		Username: "default",
		Password: "",
	}

	clickhouseClient, err := clickhouse.NewClickHouseClient(logger, clickhouseConfig)
	if err != nil {
		log.Fatal(err)
	}
	_ = clickhouseClient
}
