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

type IDatabase[Conn any] interface {
	Open(ctx context.Context, connectionConfig ConnectionConfig) (IDatabase[Conn], error)
	BatchWrite(ctx context.Context, tableName string, chunks []any) error
	Connection() Conn
}
