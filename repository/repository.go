package repository

import "context"

type Repository interface {
	UpsertCounterValue(ctx context.Context, counterKey string) (newCounterValue int, err error)
	GetCounterValue(ctx context.Context, key string) (counterValue int, err error)
}
