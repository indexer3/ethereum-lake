package cmd

import (
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var ethereumLakeCommand = &cobra.Command{}

func EthereumLakeCommand() *cobra.Command {
	ethereumLakeCommand.AddCommand(indexerCmd)
	ethereumLakeCommand.AddCommand(relayerCmd)

	ethereumLakeCommand.PersistentFlags().String(config.ConfigPathKey, "", "Config path")

	if err := viper.BindPFlags(ethereumLakeCommand.Flags()); err != nil {
		log.Error("failed to bind cmd flag", zap.Error(err))
	}

	return ethereumLakeCommand
}
