package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
)

// GetRootQueryCmd returns the root query command for this module
func GetRootQueryCmd() *cobra.Command {
	return GetQueryCmd()
}

// GetRootTxCmd returns the root transaction command for this module
func GetRootTxCmd() *cobra.Command {
	return GetTxCmd()
}