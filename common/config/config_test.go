package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOnLoadConfig(t *testing.T) {
	t.Run("test on load indexer config", func(t *testing.T) {
		err := LoadConfig("../../deploy/config/indexer.yaml", AppTypeIndexer)
		require.Nil(t, err)

		t.Log(IndexerConf)
	})

	t.Run("test on load relayer config", func(t *testing.T) {
		err := LoadConfig("../../deploy/config/relayer.yaml", AppTypeRelayer)
		require.Nil(t, err)

		t.Log(RelayerConf)
	})
}
