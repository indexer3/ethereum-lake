package cmd

import (
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   constant.LakeCommandName,
	Short: "etl command",
}

func init() {
	RootCommand.AddCommand(indexCommand)
}
