package cli

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/dynacontract/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	dynacontractTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	dynacontractTxCmd.AddCommand(
		CmdCreateContract(),
		CmdUpdateContract(),
		CmdDeleteContract(),
		CmdExecuteContract(),
		CmdUpgradeContract(),
	)

	return dynacontractTxCmd
}

// CmdCreateContract implements the create contract command
func CmdCreateContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-contract [name] [version] [code-file] [init-msg]",
		Short: "Create a new dynamic contract",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Read code file
			codeBytes, err := ioutil.ReadFile(args[2])
			if err != nil {
				return fmt.Errorf("failed to read code file: %w", err)
			}

			msg := types.NewMsgCreateContract(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
				string(codeBytes),
				args[3],
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdUpdateContract implements the update contract command
func CmdUpdateContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-contract [id] [name] [version] [code-file] [update-msg]",
		Short: "Update an existing dynamic contract",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Read code file
			codeBytes, err := ioutil.ReadFile(args[3])
			if err != nil {
				return fmt.Errorf("failed to read code file: %w", err)
			}

			msg := types.NewMsgUpdateContract(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
				args[2],
				string(codeBytes),
				args[4],
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdDeleteContract implements the delete contract command
func CmdDeleteContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-contract [id]",
		Short: "Delete a dynamic contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteContract(
				clientCtx.GetFromAddress().String(),
				args[0],
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdExecuteContract implements the execute contract command
func CmdExecuteContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute-contract [id] [execute-msg] [coins]",
		Short: "Execute a dynamic contract",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var coins sdk.Coins
			if len(args) == 3 && args[2] != "" {
				coins, err = sdk.ParseCoinsNormalized(args[2])
				if err != nil {
					return fmt.Errorf("failed to parse coins: %w", err)
				}
			}

			msg := types.NewMsgExecuteContract(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
				coins,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdUpgradeContract implements the upgrade contract command
func CmdUpgradeContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-contract [id] [version] [code-file] [migrate-msg]",
		Short: "Upgrade a dynamic contract",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Read code file
			codeBytes, err := ioutil.ReadFile(args[2])
			if err != nil {
				return fmt.Errorf("failed to read code file: %w", err)
			}

			msg := types.NewMsgUpgradeContract(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
				string(codeBytes),
				args[3],
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}