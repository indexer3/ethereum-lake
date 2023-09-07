package mysql

import (
	"context"

	"github.com/indexer3/ethereum-lake/common/config"
	"github.com/indexer3/ethereum-lake/common/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ database.Database = (*MySQL)(nil)

type MySQL struct {
	db *gorm.DB
}

func Open(ctx context.Context, connectionConfig database.ConnectionConfig) (database.Database, error) {
	return nil, nil
}

func (m *MySQL) BatchWrite(ctx context.Context, tableName string, chunks []any) error {
	batchSize := int(config.IndexerConf.MySQLConfig.BatchWriteSize)

	return m.db.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(chunks, batchSize).Error
}

func (m *MySQL) Connection() any {
	return m.db
}
