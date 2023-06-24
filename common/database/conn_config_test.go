package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnConnConfig(t *testing.T) {
	t.Run("test on build postgres dsn", func(t *testing.T) {
		connConfig := ConnectionConfig{
			DBType:   Postgres,
			Host:     "localhost",
			Port:     "4000",
			Username: "postgres",
			Password: "password",
			Database: "database",
		}

		dsn, err := connConfig.DSN()
		assert.NoError(t, err)

		assert.Equal(t, dsn, "postgres://postgres:password@localhost:4000/database?sslmode=disable")
	})

	t.Run("test on build mysql dsn", func(t *testing.T) {
		connConfig := ConnectionConfig{
			DBType:   MySQL,
			Host:     "localhost",
			Port:     "4000",
			Username: "user",
			Password: "password",
			Database: "database",
		}
		dsn, err := connConfig.DSN()
		assert.NoError(t, err)

		assert.Equal(t, dsn, "user:password@tcp(localhost:4000)/database?charset=utf8mb4&parseTime=True")
	})

	t.Run("test on build clickhouse dsn", func(t *testing.T) {
		connConfig := ConnectionConfig{
			DBType:   ClickHouse,
			Host:     "localhost",
			Port:     "4000",
			Username: "user",
			Password: "password",
			Database: "database",
		}

		dsn, err := connConfig.DSN()
		assert.NoError(t, err)

		assert.Equal(t, dsn, "tcp://user:password@localhost:4000?database=database")
	})

}
