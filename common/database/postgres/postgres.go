package postgres

import (
	"context"

	"github.com/indexer3/ethereum-lake/common/config"
	"github.com/indexer3/ethereum-lake/common/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ database.Database = (*Postgres)(nil)

type Postgres struct {
	db *gorm.DB
}

func (p *Postgres) Open(ctx context.Context, connectionConfig database.ConnectionConfig) (database.Database, error) {
	return nil, nil
}

func (p *Postgres) BatchWrite(ctx context.Context, tableName string, chunks []any) error {
	batchSize := int(config.IndexerConf.PostgresConfig.BatchWriteSize)

	return p.db.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(chunks, batchSize).Error
}

func (p *Postgres) Connection() any {
	return p.db
}
