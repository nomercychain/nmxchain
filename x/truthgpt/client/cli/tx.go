package cli

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/truthgpt/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	truthgptTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	truthgptTxCmd.AddCommand(
		NewRegisterOracleProviderCmd(),
		NewUpdateOracleProviderCmd(),
		NewDeregisterOracleProviderCmd(),
		NewStakeOracleProviderCmd(),
		NewUnstakeOracleProviderCmd(),
		NewCreateOracleRequestCmd(),
		NewSubmitOracleResponseCmd(),
		NewCancelOracleRequestCmd(),
		NewUpdateProviderReputationCmd(),
		NewCreateDataSourceCmd(),
		NewUpdateDataSourceCmd(),
		NewRemoveDataSourceCmd(),
		NewCreateOracleScriptCmd(),
		NewUpdateOracleScriptCmd(),
		NewRemoveOracleScriptCmd(),
	)

	return truthgptTxCmd
}

// NewRegisterOracleProviderCmd returns a CLI command handler for registering an oracle provider
func NewRegisterOracleProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-oracle-provider [name] [description] [website] [identity] [staked-amount]",
		Short: "Register a new oracle provider",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]
			description := args[1]
			website := args[2]
			identity := args[3]

			stakedAmount, err := sdk.ParseCoinNormalized(args[4])
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterOracleProvider(
				clientCtx.GetFromAddress().String(),
				name,
				description,
				website,
				identity,
				stakedAmount,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewUpdateOracleProviderCmd returns a CLI command handler for updating an oracle provider
func NewUpdateOracleProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-oracle-provider [name] [description] [website] [identity]",
		Short: "Update an existing oracle provider",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]
			description := args[1]
			website := args[2]
			identity := args[3]

			msg := types.NewMsgUpdateOracleProvider(
				clientCtx.GetFromAddress().String(),
				name,
				description,
				website,
				identity,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewDeregisterOracleProviderCmd returns a CLI command handler for deregistering an oracle provider
func NewDeregisterOracleProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deregister-oracle-provider",
		Short: "Deregister an existing oracle provider",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeregisterOracleProvider(
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewStakeOracleProviderCmd returns a CLI command handler for staking tokens to an oracle provider
func NewStakeOracleProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake-oracle-provider [amount]",
		Short: "Stake additional tokens to an oracle provider",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgStakeOracleProvider(
				clientCtx.GetFromAddress().String(),
				amount,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewUnstakeOracleProviderCmd returns a CLI command handler for unstaking tokens from an oracle provider
func NewUnstakeOracleProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unstake-oracle-provider [amount]",
		Short: "Unstake tokens from an oracle provider",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgUnstakeOracleProvider(
				clientCtx.GetFromAddress().String(),
				amount,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewCreateOracleRequestCmd returns a CLI command handler for creating an oracle request
func NewCreateOracleRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-oracle-request [oracle-script-id] [calldata] [ask-count] [min-count] [client-id] [fee-limit] [prepare-gas] [execute-gas] [timeout-blocks]",
		Short: "Create a new oracle request",
		Long: `Create a new oracle request using an oracle script.

The calldata should be encoded in hex format.

The fee-limit is the maximum tokens that will be paid to oracle providers.

The prepare-gas and execute-gas are the gas limits for preparation and execution phases.

The timeout-blocks is the number of blocks after which the request times out.`,
		Args: cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			oracleScriptID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			calldata, err := hex.DecodeString(args[1])
			if err != nil {
				return err
			}

			askCount, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			minCount, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}

			clientID := args[4]

			feeLimit, err := sdk.ParseCoinsNormalized(args[5])
			if err != nil {
				return err
			}

			prepareGas, err := strconv.ParseUint(args[6], 10, 64)
			if err != nil {
				return err
			}

			executeGas, err := strconv.ParseUint(args[7], 10, 64)
			if err != nil {
				return err
			}

			timeoutBlocks, err := strconv.ParseInt(args[8], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateOracleRequest(
				oracleScriptID,
				calldata,
				askCount,
				minCount,
				clientID,
				feeLimit,
				prepareGas,
				executeGas,
				timeoutBlocks,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewSubmitOracleResponseCmd returns a CLI command handler for submitting an oracle response
func NewSubmitOracleResponseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-oracle-response [request-id] [result] [confidence]",
		Short: "Submit a response to an oracle request",
		Long: `Submit a response to an oracle request.

The result should be encoded in hex format.

The confidence is a decimal value between 0 and 1 representing the confidence level of the response.`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			requestID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			result, err := hex.DecodeString(args[1])
			if err != nil {
				return err
			}

			confidence, err := sdk.NewDecFromStr(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitOracleResponse(
				requestID,
				result,
				confidence,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewCancelOracleRequestCmd returns a CLI command handler for canceling an oracle request
func NewCancelOracleRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-oracle-request [request-id]",
		Short: "Cancel an oracle request",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			requestID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelOracleRequest(
				requestID,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewUpdateProviderReputationCmd returns a CLI command handler for updating a provider's reputation
func NewUpdateProviderReputationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-provider-reputation [provider-address] [reputation-change] [reason]",
		Short: "Update an oracle provider's reputation (admin only)",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			providerAddr := args[0]

			reputationChange, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return err
			}

			reason := args[2]

			msg := types.NewMsgUpdateProviderReputation(
				clientCtx.GetFromAddress().String(),
				providerAddr,
				reputationChange,
				reason,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewCreateDataSourceCmd returns a CLI command handler for creating a data source
func NewCreateDataSourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-data-source [name] [description] [executable] [fee] [owner]",
		Short: "Create a new data source",
		Long: `Create a new data source for oracle scripts to use.

The executable should be a path to a file containing the executable code.

The fee is the payment required for using this data source.

The owner is the address that will receive the fee payments.`,
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]
			description := args[1]

			executable, err := ioutil.ReadFile(args[2])
			if err != nil {
				return err
			}

			fee, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}

			owner := args[4]

			msg := types.NewMsgCreateDataSource(
				name,
				description,
				executable,
				fee,
				owner,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewUpdateDataSourceCmd returns a CLI command handler for updating a data source
func NewUpdateDataSourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-data-source [id] [name] [description] [executable] [fee] [owner]",
		Short: "Update an existing data source",
		Long: `Update an existing data source.

The executable should be a path to a file containing the executable code.

The fee is the payment required for using this data source.

The owner is the address that will receive the fee payments.`,
		Args: cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			name := args[1]
			description := args[2]

			executable, err := ioutil.ReadFile(args[3])
			if err != nil {
				return err
			}

			fee, err := sdk.ParseCoinNormalized(args[4])
			if err != nil {
				return err
			}

			owner := args[5]

			msg := types.NewMsgUpdateDataSource(
				id,
				name,
				description,
				executable,
				fee,
				owner,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewRemoveDataSourceCmd returns a CLI command handler for removing a data source
func NewRemoveDataSourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-data-source [id]",
		Short: "Remove an existing data source",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveDataSource(
				id,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewCreateOracleScriptCmd returns a CLI command handler for creating an oracle script
func NewCreateOracleScriptCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-oracle-script [name] [description] [schema] [source-code-url] [code] [owner]",
		Short: "Create a new oracle script",
		Long: `Create a new oracle script for processing data from data sources.

The schema is a JSON string describing the input/output format.

The code should be a path to a file containing the WebAssembly code.

The owner is the address that can update or remove this oracle script.`,
		Args: cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]
			description := args[1]
			schema := args[2]
			sourceCodeURL := args[3]

			code, err := ioutil.ReadFile(args[4])
			if err != nil {
				return err
			}

			owner := args[5]

			msg := types.NewMsgCreateOracleScript(
				name,
				description,
				schema,
				sourceCodeURL,
				code,
				owner,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewUpdateOracleScriptCmd returns a CLI command handler for updating an oracle script
func NewUpdateOracleScriptCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-oracle-script [id] [name] [description] [schema] [source-code-url] [code] [owner]",
		Short: "Update an existing oracle script",
		Long: `Update an existing oracle script.

The schema is a JSON string describing the input/output format.

The code should be a path to a file containing the WebAssembly code.

The owner is the address that can update or remove this oracle script.`,
		Args: cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			name := args[1]
			description := args[2]
			schema := args[3]
			sourceCodeURL := args[4]

			code, err := ioutil.ReadFile(args[5])
			if err != nil {
				return err
			}

			owner := args[6]

			msg := types.NewMsgUpdateOracleScript(
				id,
				name,
				description,
				schema,
				sourceCodeURL,
				code,
				owner,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewRemoveOracleScriptCmd returns a CLI command handler for removing an oracle script
func NewRemoveOracleScriptCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-oracle-script [id]",
		Short: "Remove an existing oracle script",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveOracleScript(
				id,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}