package counter

import (
	"context"
	"crypto/sha256"
	"redis-connection/example/internals/helper"
	"redis-connection/example/internals/redis"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap/zaptest"
)

// This file contains setup function to create databse server
// with testscontainers and return a Counter that use the specified
// database.

// setupCounterWithRedisCluster create a counter with a static key and using a redis client as repository
func setupCounterWithRedisCluster(t testing.TB) *Counter {
	logger := zaptest.NewLogger(t)

	ctx := context.Background()
	redisEndpoint := setupRedisClusterTestContainer(t, ctx)

	redisClient, err := redis.NewRedisClient(logger, redisEndpoint, "")
	assert.NoError(t, err)

	encryptedCounterKey := helper.EncryptKey("myCounter", "secretKey", sha256.New)
	counter := NewCounter(ctx, encryptedCounterKey, redisClient)

	return counter
}

// setupRedisClusterTestContainer setup a redis database with testcontainers library.
func setupRedisClusterTestContainer(t testing.TB, ctx context.Context) string {
	t.Helper()
	req := testcontainers.ContainerRequest{
		Image:        "docker.io/bitnami/redis-cluster:7.4",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Cluster state changed: ok"),
		Env: map[string]string{
			"ALLOW_EMPTY_PASSWORD":      "yes",
			"REDIS_CLUSTER_REPLICAS":    "0",
			"REDIS_NODES":               "127.0.0.1 127.0.0.1 127.0.0.1",
			"REDIS_CLUSTER_CREATOR":     "yes",
			"REDIS_CLUSTER_DYNAMIC_IPS": "no",
			"REDIS_CLUSTER_ANNOUNCE_IP": "127.0.0.1",
		},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)

	redisEndpoint, err := container.Endpoint(ctx, "")
	assert.NoError(t, err)

	return redisEndpoint
}

// TODO: setupCounterWithClickHouse setup a counter component and run a ClickHouse
// server to be used as a repository.
func setupCounterWithClickHouse(t testing.TB) *Counter {
	panic("todo")
	return nil
}

// TODO: setupClickHouseTestContainer setup a ClickHouse databse using testcontainers.
func setupClickHouseTestContainer(t testing.TB, ctx context.Context) string {
	panic("todo")
	return ""
}

// NOTE: This function is an alternative way of declaring a redis database but if you look
// closely to the internal of this module you will see that it is nearly the same.
//
// func setupRedisTestContainerUsingModule(t testing.TB, ctx context.Context) string {
// 	t.Helper()
// 	redisContainer, err := redis.Run(ctx, "redis:7.4")
// 	assert.NoError(t, err)

// 	redisEndpoint, err := redisContainer.Endpoint(ctx, "")
// 	assert.NoError(t, err)

// 	return redisEndpoint
//  }
