package counter

import (
	"context"
	"crypto/sha256"
	"redis-connection/example/helper"
	"redis-connection/example/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/redis"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestIncrementAndRetrieveCounter(t *testing.T) {
	ctx := context.Background()
	redisEndpoint := setupRedisTestContainer(t, ctx)

	redisClient, err := repository.NewRedisClient(redisEndpoint, "")
	assert.NoError(t, err)

	encryptedCounterKey := helper.EncryptKey("myCounter", "secretKey", sha256.New)
	counter := NewCounter(ctx, encryptedCounterKey, redisClient)

	counter.Increment()
	counter.Increment()
	counter.Increment()

	got, err := counter.Get()
	if err != nil {
		t.Fatalf("unexpected error while retrieving counter value : %v", err)
	}

	expected := 3

	if got != expected {
		t.Errorf("expected value of counter %d is not equal to %d", expected, got)
	}
}

func TestDeleteCounterAfterIncrement(t *testing.T) {
	ctx := context.Background()
	redisEndpoint := setupRedisTestContainer(t, ctx)

	redisClient, err := repository.NewRedisClient(redisEndpoint, "")
	assert.NoError(t, err)

	encryptedCounterKey := helper.EncryptKey("myCounter", "secretKey", sha256.New)
	counter := NewCounter(ctx, encryptedCounterKey, redisClient)

	counter.Increment()
	counter.Increment()
	counter.Increment()

	ok, err := counter.Delete()
	assert.NoError(t, err)
	if !ok {
		t.Fatal("delete should have return an ok with value true")
	}

	_, err = counter.Get()
	if err == nil {
		t.Error("expected an error to be thrown")
	}
}

func setupRedisTestContainer(t testing.TB, ctx context.Context) string {
	t.Helper()
	req := testcontainers.ContainerRequest{
		Image:        "redis:7.4",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
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

func setupRedisTestContainerUsingModule(t testing.TB, ctx context.Context) string {
	t.Helper()
	redisContainer, err := redis.Run(ctx, "redis:7.4")
	assert.NoError(t, err)

	redisEndpoint, err := redisContainer.Endpoint(ctx, "")
	assert.NoError(t, err)

	return redisEndpoint
}
