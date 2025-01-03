package repository

import "context"

type Repository interface {
	UpsertCounterValue(ctx context.Context, counterKey string) (newCounterValue int, err error)

	GetCounterValue(ctx context.Context, counterKey string) (counterValue int, err error)

	// DeleteCounter delete the specified counter key from repository.
	//
	// Return a boolean that indicate if the operation was successful.
	DeleteCounter(ctx context.Context, counterKey string) (ok bool, err error)
}
