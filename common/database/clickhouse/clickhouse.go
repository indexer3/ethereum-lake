package clickhouse

import (
	"context"
	"fmt"

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

func (c *ClickHouse) Open(ctx context.Context, connectionConfig database.ConnectionConfig) (database.IDatabase, error) {
	return nil, nil
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

func (c *ClickHouse) Query(ctx context.Context, statement string, args ...any) ([]any, error) {
	return nil, nil
}
