package postgres

import (
	"context"

	"github.com/indexer3/ethereum-lake/common/database"
	"github.com/indexer3/ethereum-lake/constant/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ database.IDatabase = (*Postgres)(nil)

type Postgres struct {
	db *gorm.DB
}

func (p *Postgres) Open(ctx context.Context, connectionConfig database.ConnectionConfig) (database.IDatabase, error) {
	return nil, nil
}

func (p *Postgres) BatchWrite(ctx context.Context, tableName string, chunks []any) error {
	batchSize := viper.GetInt(config.PostgresBatchWriteSize)

	return p.db.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(chunks, batchSize).Error
}

func (p *Postgres) Query(ctx context.Context, statement string, args ...any) ([]any, error) {
	return nil, nil
}
