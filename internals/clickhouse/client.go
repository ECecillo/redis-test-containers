package clickhouse

import (
	"context"
	"fmt"
	"redis-connection/example/pkg/repository"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"go.uber.org/zap"
)

type ClickHouseClient struct {
	logger *zap.Logger
	conn   *driver.Conn
}

var _ repository.Repository = &ClickHouseClient{}

func NewClickHouseClient(logger *zap.Logger, conf Config) (*ClickHouseClient, error) {

	conn, err := connect(conf)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to ClickHouse, err: %w", err)
	}

	return &ClickHouseClient{
		logger: logger.Named("clickhouse-client"),
		conn:   conn,
	}, nil
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

func connect(conf Config) (*driver.Conn, error) {

	//NOTE: simple setup not a production one.
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", conf.Host, conf.Port)},
		Auth: clickhouse.Auth{
			Database: conf.Database,
			Username: conf.Username,
			Password: conf.Password,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("unable to establish connection to ClickHouse, err: %w", err)
	}
	v, err := conn.ServerVersion()
	if err != nil {
		return nil, fmt.Errorf("unable to get ClickHouse server verison, err: %w", err)
	}
	fmt.Println("ClickHouse server version ", v)

	return &conn, nil
}
