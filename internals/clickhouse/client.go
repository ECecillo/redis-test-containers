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
	conn   driver.Conn
}

var _ repository.Repository = &ClickHouseClient{}

// NewClickHouseClient is a constructor for a ClickHouseClient, it will setup for us
// the connection between our app and the specified database server.
func NewClickHouseClient(logger *zap.Logger, conf Config) (*ClickHouseClient, error) {

	conn, err := connect(conf)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to clickhouse, err: %w", err)
	}

	return &ClickHouseClient{
		logger: logger.Named("clickhouse-client"),
		conn:   conn,
	}, nil
}

// DeleteCounter implements repository.Repository.
func (c *ClickHouseClient) DeleteCounter(ctx context.Context, counterKey string) (ok bool, err error) {

	logger := c.logger.Named("delete").With(
		zap.String("counter_key", counterKey),
	)

	logger.Info("start to delete counter value")

	if err = c.conn.Exec(ctx, "DELETE FROM counter WHERE Id = ?", counterKey); err != nil {
		logger.Warn("unable to delete counter",
			zap.Error(err),
		)
		return false, fmt.Errorf("unable to delete counter, err: %w", err)
	}

	logger.Debug("successfully removed counter")

	return true, nil
}

// GetCounterValue implements repository.Repository.
func (c *ClickHouseClient) GetCounterValue(ctx context.Context, counterKey string) (counterValue int, err error) {

	var ckCounter Counter

	logger := c.logger.Named("get").With(
		zap.String("counter_key", counterKey),
	)

	if err = c.conn.Select(ctx, &ckCounter, "SELECT Id, CounterValue FROM counter WHERE Id = ?", counterKey); err != nil {
		logger.Warn("unable to retrieve counter value",
			zap.Error(err),
		)
		return 0, fmt.Errorf("unable to retrieve counter value, err: %w", err)
	}

	return ckCounter.CounterValue, nil
}

type Counter struct {
	Id           string
	CounterValue int
}

// UpsertCounterValue implements repository.Repository.
func (c *ClickHouseClient) UpsertCounterValue(ctx context.Context, counterKey string) (newCounterValue int, err error) {
	// Since upsert command does not exist in clickhouse we are going to make a SELECT followed by INSERT with new value.
	logger := c.logger.Named("upsert").With(
		zap.String("counter_key", counterKey),
	)

	actualValue, err := c.GetCounterValue(ctx, counterKey)
	if err != nil {
		logger.Warn("unable to get counter value",
			zap.Error(err),
		)
		return 0, fmt.Errorf("unable to get counter value, err: %w", err)
	}

	newCounterValue = actualValue + 1

	logger = logger.With(
		zap.Int("new_counter_value", newCounterValue),
		zap.Int("previous_counter_value", actualValue),
	)

	logger.Debug("calculated new counter value")

	err = c.conn.Exec(ctx, "INSERT INTO counter VALUES (?, ?)", counterKey, newCounterValue)
	if err != nil {
		logger.Warn("unable to insert new counter value",
			zap.Error(err),
		)
		return 0, fmt.Errorf("unable to insert new counter value, err: %w", err)
	}

	logger.Info("new counter value inserted")

	return newCounterValue, nil
}

// Ping just check if we can
func (c ClickHouseClient) Ping(ctx context.Context) error {
	logger := c.logger.Named("ping")

	logger.Info("PING")
	if err := c.conn.Ping(ctx); err != nil {
		logger.Fatal("unable to ping clickhouse server")
		return fmt.Errorf("unable to ping clickhouse server, err: %w", err)
	}
	logger.Info("PONG")

	return nil
}

func connect(conf Config) (driver.Conn, error) {

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
		return nil, fmt.Errorf("unable to establish connection to clickhouse, err: %w", err)
	}
	v, err := conn.ServerVersion()
	if err != nil {
		return nil, fmt.Errorf("unable to get clickhouse server verison, err: %w", err)
	}
	fmt.Println("clickhouse server version ", v)

	return conn, nil
}
