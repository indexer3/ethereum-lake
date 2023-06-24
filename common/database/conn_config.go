package database

import "fmt"

type ConnectionConfig struct {
	DBType   DatabaseType
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func (c ConnectionConfig) DSN() (string, error) {
	switch c.DBType {
	case ClickHouse:
		return fmt.Sprintf("tcp://%s:%s@%s:%s?database=%s", c.Username, c.Password, c.Host, c.Port, c.Database), nil
	case Postgres:
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.Username, c.Password, c.Host, c.Port, c.Database), nil
	case MySQL:
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", c.Username, c.Password, c.Host, c.Port, c.Database), nil
	}

	return "", fmt.Errorf("unsupported database type")
}
