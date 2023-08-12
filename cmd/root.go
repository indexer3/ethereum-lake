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

	ethereumLakeCommand.PersistentFlags().String(config.ClickHouseHost, "localhost", "ClickHouse host")
	ethereumLakeCommand.PersistentFlags().Int(config.ClickHousePort, 9000, "ClickHouse port")
	ethereumLakeCommand.PersistentFlags().String(config.ClickHouseUser, "default", "ClickHouse user")
	ethereumLakeCommand.PersistentFlags().String(config.ClickHousePassword, "", "ClickHouse password")
	ethereumLakeCommand.PersistentFlags().String(config.ClickHouseBatchWriteSize, "default", "ClickHouse database")

	ethereumLakeCommand.PersistentFlags().String(config.MySQLHost, "localhost", "MySQL host")
	ethereumLakeCommand.PersistentFlags().Int(config.MySQLPort, 8080, "MySQL port")
	ethereumLakeCommand.PersistentFlags().String(config.MySQLUser, "default", "MySQL user")
	ethereumLakeCommand.PersistentFlags().String(config.MySQLPassword, "", "MySQL password")
	ethereumLakeCommand.PersistentFlags().String(config.MySQLBatchWriteSize, "default", "MySQL database")

	ethereumLakeCommand.PersistentFlags().String(config.PostgresHost, "localhost", "Postgres host")
	ethereumLakeCommand.PersistentFlags().Int(config.PostgresPort, 8081, "MySQL port")
	ethereumLakeCommand.PersistentFlags().String(config.PostgresUser, "default", "Postgres user")
	ethereumLakeCommand.PersistentFlags().String(config.PostgresPassword, "", "Postgres password")
	ethereumLakeCommand.PersistentFlags().String(config.PostgresBatchWriteSize, "default", "Postgres database")

	ethereumLakeCommand.PersistentFlags().Int(config.RPCMaxRetry, 5, "RPC max retry")

	ethereumLakeCommand.PersistentFlags().String(config.ChainbaseSQLAPIKey, "demo", "Chainbase SQL API key")

	ethereumLakeCommand.PersistentFlags().String(config.EthereumRPCUrl, "", "Ethereum RPC URL")
	ethereumLakeCommand.PersistentFlags().String(config.OptimismRPCUrl, "", "Optimism RPC URL")
	ethereumLakeCommand.PersistentFlags().String(config.ArbitrumRPCUrl, "", "Arbitrum RPC URL")
	ethereumLakeCommand.PersistentFlags().String(config.BSCRPCUrl, "", "BSC RPC URL")
	ethereumLakeCommand.PersistentFlags().String(config.PolygonRPCUrl, "", "Polygon RPC URL")
	ethereumLakeCommand.PersistentFlags().String(config.BaseRPCUrl, "", "Base RPC URL")
	ethereumLakeCommand.PersistentFlags().String(config.GoerliRPCUrl, "", "Goerli RPC URL")

	if err := viper.BindPFlags(ethereumLakeCommand.Flags()); err != nil {
		log.Error("failed to bind cmd flag", zap.Error(err))
	}

	return ethereumLakeCommand
}
