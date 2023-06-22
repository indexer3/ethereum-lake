package postgres

import (
	"context"

	"github.com/indexer3/ethereum-lake/common/database"
	"gorm.io/gorm"
)

var _ database.IDatabase = (*Postgres)(nil)

type Postgres struct {
	db *gorm.DB
}

func (p *Postgres) Open(ctx context.Context, connectionConfig database.ConnectionConfig) (database.IDatabase, error) {
	return nil, nil
}

func (p *Postgres) BatchWrite(ctx context.Context, tableName string, chunks []any) error {
	return nil
}

func (p *Postgres) Query(ctx context.Context, statement string, args ...any) ([]any, error) {
	return nil, nil
}
