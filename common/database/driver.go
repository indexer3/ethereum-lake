package database

import (
	"context"
)

type DatabaseType string

const (
	Postgres   DatabaseType = "POSTGRES"
	ClickHouse DatabaseType = "CLICKHOUSE"
	MySQL      DatabaseType = "MYSQL"
)

type IDatabase interface {
	Open(ctx context.Context, connectionConfig ConnectionConfig) (IDatabase, error)
	BatchWrite(ctx context.Context, tableName string, chunks []any) error
	Query(ctx context.Context, statement string, args ...any) ([]any, error)
}
