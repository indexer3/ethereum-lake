package cmd

import (
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/spf13/cobra"
)

var relayerCmd = &cobra.Command{
	Use: constant.RelayerCommandName,
}
