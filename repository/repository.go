package repository

import "context"

type Repository interface {
	GetCounterValue(ctx context.Context, key string) int
}
