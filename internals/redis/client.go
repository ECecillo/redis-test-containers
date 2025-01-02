package redis

import (
	"context"
	"fmt"
	"redis-connection/example/pkg/repository"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisClient struct {
	logger *zap.Logger
	conn   *redis.Client
}

var _ repository.Repository = &RedisClient{}

func NewRedisClient(logger *zap.Logger, redisServerAddress string, password string) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisServerAddress,
		Password: password,
		DB:       0,
	})

	client := &RedisClient{
		logger: logger.Named("redis-client"),
		conn:   rdb,
	}

	return client, nil
}

func (redisClient RedisClient) GetCounterValue(ctx context.Context, counterKey string) (counterValue int, err error) {
	val, err := redisClient.conn.Get(ctx, counterKey).Int()
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (redisClient RedisClient) UpsertCounterValue(ctx context.Context, counterKey string) (newCounterValue int, err error) {
	cmdResult := redisClient.conn.Incr(ctx, counterKey)
	err = cmdResult.Err()
	if err != nil {
		return 0, fmt.Errorf("failed to set the counter key %s with value %d : %w", counterKey, newCounterValue, err)
	}

	value, err := cmdResult.Uint64()
	if err != nil {
		return 0, fmt.Errorf("failed to transform redis result into uint64 value : %w", err)
	}

	return int(value), nil
}

func (redisClient RedisClient) DeleteCounter(ctx context.Context, counterKey string) (ok bool, err error) {
	result := redisClient.conn.Del(ctx, counterKey)
	err = result.Err()
	if err != nil {
		return false, fmt.Errorf("tried to delete the following key from redis %s but got err : %w", counterKey, err)
	}
	return true, nil
}

// To test if the connection has been setup correctly
func (redisClient RedisClient) Ping(ctx context.Context) error {
	logger := redisClient.logger.Named("ping")

	logger.Debug("PING")
	pong, err := redisClient.conn.Ping(context.Background()).Result()
	if err != nil {
		// logger.Fatal("could not ping redis server",
		// 	zap.String("adress", redisClient.conn.Options().Network),
		// )
		return err
	}
	logger.Debug(pong)
	return nil
}
