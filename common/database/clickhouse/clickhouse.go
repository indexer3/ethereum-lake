package clickhouse

import (
	"context"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/indexer3/ethereum-lake/common/database"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant/config"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var _ database.IDatabase = (*ClickHouse)(nil)

type ClickHouse struct {
	db driver.Conn
}

func Open(ctx context.Context, connectionConfig database.ConnectionConfig) (*ClickHouse, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{
			fmt.Sprintf("%s:%s", connectionConfig.Host, connectionConfig.Port),
		},
		Auth: clickhouse.Auth{
			Database: connectionConfig.Database,
			Username: connectionConfig.Username,
			Password: connectionConfig.Password,
		},
		ClientInfo: clickhouse.ClientInfo{
			Products: []struct {
				Name    string
				Version string
			}{{
				Name:    "ethereum-lake[clickhouse]",
				Version: "v0.0.0",
			}},
		},
		Debugf: func(format string, v ...any) {
			log.Info("ethereum-lake[clickhouse info]", zap.Any("info", fmt.Sprintf(format, v)))
		},
		Debug:           false,
		DialTimeout:     time.Minute * 10,
		MaxOpenConns:    50,
		MaxIdleConns:    50,
		ConnMaxLifetime: time.Hour * 48,
	})
	if err != nil {
		log.Error("failed to connect to clickhouse", zap.Error(err))
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		log.Error("failed to ping clickhouse", zap.Error(err))
		return nil, err
	}

	return &ClickHouse{db: conn}, nil
}

func (c *ClickHouse) BatchWrite(ctx context.Context, tableName string, dataArr []any) error {
	batchSize := viper.GetInt(config.ClickHouseBatchWriteSize)
	chunks := lo.Chunk(dataArr, batchSize)

	for _, chunk := range chunks {
		batch, err := c.db.PrepareBatch(ctx, fmt.Sprintf("insert into %s", tableName))
		if err != nil {
			log.Error("failed to preprare batch in clickhouse", zap.Error(err))
			return err
		}

		for _, c := range chunk {
			err = batch.AppendStruct(c)
			if err != nil {
				log.Error("failed to append batch struct in clickhouse", zap.Error(err))
				return err
			}
		}

		if err = batch.Send(); err != nil {
			log.Error("failed to send batch write data to clickhouse", zap.Error(err))
			return err
		}
	}

	return nil
}

func (c *ClickHouse) Connection() any {
	return c.db
}
