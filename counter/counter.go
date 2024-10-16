package counter

import (
	"context"
	"fmt"
	"redis-connection/example/repository"
)

type Counter struct {
	key        string
	ctx        context.Context
	repository repository.Repository
}

func NewCounter(ctx context.Context, encryptedCounterKey string, repository repository.Repository) *Counter {
	return &Counter{
		ctx:        ctx,
		key:        encryptedCounterKey,
		repository: repository,
	}
}

func (c *Counter) Increment() error {
	_, err := c.repository.UpsertCounterValue(c.ctx, c.key)
	if err != nil {
		return fmt.Errorf("failed to increment counter value with key %s : %w", c.key, err)
	}

	return nil
}

func (c *Counter) Get() (int, error) {
	value, err := c.repository.GetCounterValue(c.ctx, c.key)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve counter value of key %s : %w", c.key, err)
	}

	return value, err
}

func (c *Counter) GetKey() string {
	return c.key
}
