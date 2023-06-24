package mysql

import (
	"context"

	"github.com/indexer3/ethereum-lake/common/database"
	"github.com/indexer3/ethereum-lake/constant/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ database.IDatabase = (*MySQL)(nil)

type MySQL struct {
	db *gorm.DB
}

func Open(ctx context.Context, connectionConfig database.ConnectionConfig) (database.IDatabase, error) {
	return nil, nil
}

func (m *MySQL) BatchWrite(ctx context.Context, tableName string, chunks []any) error {
	batchSize := viper.GetInt(config.MySQLBatchWriteSize)

	return m.db.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(chunks, batchSize).Error
}

func (m *MySQL) Connection() any {
	return m.db
}
