package cmd

import (
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/spf13/cobra"
)

var indexCommand = &cobra.Command{
	Use:   constant.IndexCommandName,
	Short: "index",
}
