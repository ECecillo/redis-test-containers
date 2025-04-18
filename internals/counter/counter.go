package counter

import (
	"context"
	"fmt"
	"redis-connection/example/pkg/repository"
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
		return 0, fmt.Errorf("failed to retrieve counter value of key %s, err: %w", c.key, err)
	}

	return value, err
}

func (c *Counter) Delete() (ok bool, err error) {
	ok, err = c.repository.DeleteCounter(c.ctx, c.key)
	if err != nil {
		return false, fmt.Errorf("failed to delete counter with key %s : %w", c.key, err)
	}
	if !ok {
		return false, fmt.Errorf("query returned no error but got the result is not ok")
	}
	return true, nil
}

func (c *Counter) GetKey() string {
	return c.key
}
