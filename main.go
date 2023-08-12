package main

import (
	"context"

	"github.com/indexer3/ethereum-lake/cmd"
	"github.com/indexer3/ethereum-lake/common/log"
	"go.uber.org/zap"
)

func main() {
	if err := cmd.EthereumLakeCommand().ExecuteContext(context.Background()); err != nil {
		log.Error("failed to execute cmd", zap.Error(err))
	}
}
