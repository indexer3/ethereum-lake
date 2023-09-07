package config

import (
	"errors"
	"strings"

	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type AppType string

const (
	AppTypeIndexer AppType = "indexer"
	AppTypeRelayer AppType = "relayer"
)

func (a AppType) String() string {
	return string(a)
}

func LoadConfig(path string, app AppType) error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)
	viper.SetEnvPrefix("CONFIG_ENV")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Error("failed to read in config", zap.Error(err))
		return err
	}

	switch app {
	case AppTypeIndexer:
		if err := viper.Unmarshal(&IndexerConf); err != nil {
			log.Error("failed to unmarshal config", zap.Error(err))
			return err
		}
	case AppTypeRelayer:
		if err := viper.Unmarshal(&RelayerConf); err != nil {
			log.Error("failed to unmarshal config", zap.Error(err))
			return err
		}
	default:
		log.Error("unknown app type", zap.String("app", app.String()))
		return errors.New("unknown app type")
	}

	return nil
}
