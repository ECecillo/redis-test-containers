package store

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	rdb *redis.Client
}

func NewRedisClient(redisServerAddress string) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisServerAddress,
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
		DB:       0,                                  // use default DB
	})
	return &RedisClient{
		rdb: rdb,
	}
}

func (redisClient *RedisClient) GetCounterValue(counterKey string) (int, error) {
	val, err := redisClient.rdb.Get(ctx, counterKey).Int()
	if err != nil {
		return 0, err
	}
	return val, nil
}

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
		DB:       0,                                  // use default DB
	})
	err := rdb.Set(ctx, "MyFirstKeyForValue1", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "MyFirstKeyForValue1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("MyFirstKeyForValue1", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: MyFirstKeyForValue1 value
	// key2 does not exist
}
