package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnConnConfig(t *testing.T) {
	t.Run("test on build postgres dsn", func(t *testing.T) {
		connConfig := ConnectionConfig{
			
		}

		dsn, err := connConfig.DSN(Postgres)
		assert.NoError(t, err)

		assert.Equal(t, dsn, "")
	})



}
