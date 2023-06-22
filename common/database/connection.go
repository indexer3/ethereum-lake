package database

import "fmt"

type ConnectionConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func (c ConnectionConfig) DSN(dbType DatabaseType) (string, error) {
	switch dbType {
	case ClickHouse:

	case Postgres:

	case MySQL:

	}

	return "", fmt.Errorf("unsupported database type")
}
