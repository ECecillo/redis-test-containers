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
)

func TestIncrementAndRetrieveCounter(t *testing.T) {
	counter := setupCounterWithRedis(t)

	counter.Increment()
	counter.Increment()
	counter.Increment()

	got, err := counter.Get()
	assert.NoError(t, err, "unexpected error while retrieving counter value")

	expected := 3

	assert.Equal(t, expected, got, "expected same counter value")
}

func TestDeleteCounterAfterIncrement(t *testing.T) {
	counter := setupCounterWithRedis(t)

	counter.Increment()
	counter.Increment()
	counter.Increment()

	ok, err := counter.Delete()
	assert.NoError(t, err)

	assert.False(t, ok, "delete should have return an ok with value true")

	_, err = counter.Get()
	assert.Error(t, err, "expected an error to be thrown")
}

func TestDeleteCounterWhenNotExisting(t *testing.T) {
	counter := setupCounterWithRedis(t)

	ok, err := counter.Delete()
	assert.NoError(t, err)
	assert.False(t, ok, "delete should have return an ok with value true")
}

// setupCounterWithRedis create a counter with a static key and using a redis client as repository
func setupCounterWithRedis(t testing.TB) *Counter {
	ctx := context.Background()
	redisEndpoint := setupRedisTestContainer(t, ctx)

	redisClient, err := redis.NewRedisClient(redisEndpoint, "")
	assert.NoError(t, err)

	encryptedCounterKey := helper.EncryptKey("myCounter", "secretKey", sha256.New)
	counter := NewCounter(ctx, encryptedCounterKey, redisClient)

	return counter
}

// setupRedisTestContainer setup a redis database with testcontainers library.
func setupRedisTestContainer(t testing.TB, ctx context.Context) string {
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
