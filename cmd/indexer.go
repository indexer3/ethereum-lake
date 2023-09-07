package cmd

import (
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/spf13/cobra"
)

var indexerCmd = &cobra.Command{
	Use: constant.IndexerCommandName,
}
