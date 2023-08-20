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

type Database interface {
	BatchWrite(ctx context.Context, tableName string, chunks []any) error
	Connection() any
}
