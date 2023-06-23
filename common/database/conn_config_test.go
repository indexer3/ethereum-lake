package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnConnConfig(t *testing.T) {
	t.Run("test on build postgres dsn", func(t *testing.T) {
		connConfig := ConnectionConfig{
			Host:     "localhost",
			Port:     "4000",
			Username: "postgres",
			Password: "password",
			Database: "database",
		}

		dsn, err := connConfig.DSN(Postgres)
		assert.NoError(t, err)

		assert.Equal(t, dsn, "postgres://postgres:password@localhost:4000/database?sslmode=disable")
	})

}
