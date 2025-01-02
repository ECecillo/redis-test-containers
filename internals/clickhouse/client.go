package clickhouse

import (
	"context"
	"redis-connection/example/pkg/repository"

	"go.uber.org/zap"
)

type ClickHouseClient struct {
	logger *zap.Logger
}

var _ repository.Repository = &ClickHouseClient{}

type Config struct {
	Host string
	Port int

	DatabaseName string
	User         string
	Password     string
}

func NewClickHouseClient(logger *zap.Logger, conf Config) *ClickHouseClient {

	return &ClickHouseClient{
		logger: logger.Named("clickhouse-client"),
		//TODO: add connection
	}
}

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
