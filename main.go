package main

import (
	"context"

	"github.com/indexer3/ethereum-lake/cmd"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	if err := viper.BindPFlags(cmd.RootCommand.Flags()); err != nil {
		log.Error("failed to bind cmd flag", zap.Error(err))
	}

	if err := cmd.RootCommand.ExecuteContext(context.Background()); err != nil {
		log.Error("failed to execute cmd", zap.Error(err))
	}
}
