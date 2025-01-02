package repository

import "context"

type Repository interface {
	UpsertCounterValue(ctx context.Context, counterKey string) (newCounterValue int, err error)
	GetCounterValue(ctx context.Context, counterKey string) (counterValue int, err error)
	DeleteCounter(ctx context.Context, counterKey string) (ok bool, err error)
}

var (
	_ Repository = &RedisClient{}
)
