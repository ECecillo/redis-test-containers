package clickhouse

import (
	"context"
	"redis-connection/example/pkg/repository"
)

type ClickHouseClient struct{}

var _ repository.Repository = &ClickHouseClient{}

// DeleteCounter implements repository.Repository.
func (c *ClickHouseClient) DeleteCounter(ctx context.Context, counterKey string) (ok bool, err error) {
	panic("unimplemented")
}

// GetCounterValue implements repository.Repository.
func (c *ClickHouseClient) GetCounterValue(ctx context.Context, counterKey string) (counterValue int, err error) {
	panic("unimplemented")
}

// UpsertCounterValue implements repository.Repository.
func (c *ClickHouseClient) UpsertCounterValue(ctx context.Context, counterKey string) (newCounterValue int, err error) {
	panic("unimplemented")
}
