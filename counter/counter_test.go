package counter

import (
	"crypto/sha256"
	"redis-connection/example/helper"
	"redis-connection/example/store"
	"testing"
)

// Calling an Increment function will update a counter in our redis database
func TestIncrementCounter(t *testing.T) {
	//FIXME: For now not working because we don't have any key with counterKey
	// we need to setup testcontainers in order to make this work, or check redis github for testing locally.

	redisClient := store.NewRedisClient("localhost:6379")
	encryptedCounterKey := helper.EncryptKey("myCounter", "secretKey", sha256.New)
	counter := &Counter{
		Key: encryptedCounterKey,
	}
	counter.Increment()

	got, err := redisClient.GetCounterValue(encryptedCounterKey)
	if err != nil {
		t.Fatalf("unexpected error while retrieving counter value : %v", err)
	}

	expected := 1

	if got != expected {
		t.Errorf("expected value of counter %d is not equal to %d", expected, got)
	}
}
