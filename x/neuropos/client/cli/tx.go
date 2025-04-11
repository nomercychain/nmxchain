package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/neuropos/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	neuroposTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	neuroposTxCmd.AddCommand(
		NewCreateValidatorCmd(),
		NewEditValidatorCmd(),
		NewDelegateCmd(),
		NewUndelegateCmd(),
		NewRedelegateCmd(),
		NewUnjailCmd(),
		NewUpdateNeuralNetworkCmd(),
		NewTrainNeuralNetworkCmd(),
		NewSubmitNeuralPredictionCmd(),
		NewUpdateValidatorReputationCmd(),
	)

	return neuroposTxCmd
}

// NewCreateValidatorCmd returns a CLI command handler for creating a validator
func NewCreateValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-validator [pubkey] [amount] [commission-rate] [commission-max-rate] [commission-max-change-rate] [min-self-delegation] [moniker] [identity] [website] [security-contact] [details]",
		Short: "Create a new validator",
		Args:  cobra.ExactArgs(11),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pubkey := args[0]

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			commissionRate, err := sdk.NewDecFromStr(args[2])
			if err != nil {
				return err
			}

			commissionMaxRate, err := sdk.NewDecFromStr(args[3])
			if err != nil {
				return err
			}

			commissionMaxChangeRate, err := sdk.NewDecFromStr(args[4])
			if err != nil {
				return err
			}

			minSelfDelegation, ok := sdk.NewIntFromString(args[5])
			if !ok {
				return fmt.Errorf("invalid min self delegation amount: %s", args[5])
			}

			moniker := args[6]
			identity := args[7]
			website := args[8]
			securityContact := args[9]
			details := args[10]

			valAddr := clientCtx.GetFromAddress().String()

			msg := types.NewMsgCreateValidator(
				sdk.ValAddress(clientCtx.GetFromAddress()),
				pubkey,
				amount,
				types.Description{
					Moniker:         moniker,
					Identity:        identity,
					Website:         website,
					SecurityContact: securityContact,
					Details:         details,
				},
				types.CommissionRates{
					Rate:          commissionRate,
					MaxRate:       commissionMaxRate,
					MaxChangeRate: commissionMaxChangeRate,
				},
				minSelfDelegation,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewEditValidatorCmd returns a CLI command handler for editing a validator
func NewEditValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-validator [commission-rate] [min-self-delegation] [moniker] [identity] [website] [security-contact] [details]",
		Short: "Edit an existing validator",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var commissionRate *sdk.Dec
			if args[0] != "" {
				rate, err := sdk.NewDecFromStr(args[0])
				if err != nil {
					return err
				}
				commissionRate = &rate
			}

			var minSelfDelegation *sdk.Int
			if args[1] != "" {
				msd, ok := sdk.NewIntFromString(args[1])
				if !ok {
					return fmt.Errorf("invalid min self delegation amount: %s", args[1])
				}
				minSelfDelegation = &msd
			}

			moniker := args[2]
			identity := args[3]
			website := args[4]
			securityContact := args[5]
			details := args[6]

			valAddr := sdk.ValAddress(clientCtx.GetFromAddress())

			msg := types.NewMsgEditValidator(
				valAddr,
				types.Description{
					Moniker:         moniker,
					Identity:        identity,
					Website:         website,
					SecurityContact: securityContact,
					Details:         details,
				},
				commissionRate,
				minSelfDelegation,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewDelegateCmd returns a CLI command handler for delegating tokens to a validator
func NewDelegateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegate [validator-addr] [amount]",
		Short: "Delegate tokens to a validator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			valAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			delAddr := clientCtx.GetFromAddress()

			msg := types.NewMsgDelegate(delAddr, valAddr, amount)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewUndelegateCmd returns a CLI command handler for undelegating tokens from a validator
func NewUndelegateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "undelegate [validator-addr] [amount]",
		Short: "Undelegate tokens from a validator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			valAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			delAddr := clientCtx.GetFromAddress()

			msg := types.NewMsgUndelegate(delAddr, valAddr, amount)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewRedelegateCmd returns a CLI command handler for redelegating tokens from one validator to another
func NewRedelegateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redelegate [src-validator-addr] [dst-validator-addr] [amount]",
		Short: "Redelegate tokens from one validator to another",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			valSrcAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			valDstAddr, err := sdk.ValAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			delAddr := clientCtx.GetFromAddress()

			msg := types.NewMsgBeginRedelegate(delAddr, valSrcAddr, valDstAddr, amount)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewUnjailCmd returns a CLI command handler for unjailing a validator
func NewUnjailCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unjail",
		Short: "Unjail a jailed validator",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			valAddr := sdk.ValAddress(clientCtx.GetFromAddress())

			msg := types.NewMsgUnjail(valAddr)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewUpdateNeuralNetworkCmd returns a CLI command handler for updating a neural network
func NewUpdateNeuralNetworkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-neural-network [network-id] [architecture] [layers-json] [weights-file] [metadata]",
		Short: "Update or create a neural network",
		Long: `Update an existing neural network or create a new one if network-id is empty.

The layers-json should be a JSON array of layer objects with the following format:

[{"type":"dense","input_size":10,"output_size":5,"activation":"relu"},...]

The weights-file should be a path to a file containing the neural network weights in JSON format.

The metadata is an optional string containing additional information about the neural network.`,
		Args: cobra.RangeArgs(3, 5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			networkID := ""
			if len(args) > 0 && args[0] != "" {
				networkID = args[0]
			}

			architecture := args[1]

			// Parse layers JSON
			var layersData []map[string]interface{}
			err = json.Unmarshal([]byte(args[2]), &layersData)
			if err != nil {
				return fmt.Errorf("failed to parse layers JSON: %w", err)
			}

			// Convert to proto layers
			var layers []types.Layer
			for _, layerData := range layersData {
				layerType, _ := layerData["type"].(string)
				inputSize, _ := layerData["input_size"].(float64)
				outputSize, _ := layerData["output_size"].(float64)
				activation, _ := layerData["activation"].(string)

				layers = append(layers, types.Layer{
					Type:       layerType,
					InputSize:  uint32(inputSize),
					OutputSize: uint32(outputSize),
					Activation: activation,
				})
			}

			// Read weights file if provided
			var weights json.RawMessage
			if len(args) > 3 && args[3] != "" {
				weightsBytes, err := os.ReadFile(args[3])
				if err != nil {
					return fmt.Errorf("failed to read weights file: %w", err)
				}
				weights = weightsBytes
			}

			// Get metadata if provided
			var metadata []byte
			if len(args) > 4 && args[4] != "" {
				metadata = []byte(args[4])
			}

			valAddr := sdk.ValAddress(clientCtx.GetFromAddress())

			msg := types.NewMsgUpdateNeuralNetwork(
				valAddr,
				networkID,
				architecture,
				layers,
				weights,
				metadata,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewTrainNeuralNetworkCmd returns a CLI command handler for training a neural network
func NewTrainNeuralNetworkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "train-neural-network [network-id] [features-file] [labels-file] [epochs] [learning-rate] [metadata]",
		Short: "Train a neural network",
		Long: `Train an existing neural network with the provided features and labels.

The features-file and labels-file should be paths to files containing the training data in JSON format.

The epochs parameter specifies the number of training epochs.

The learning-rate parameter specifies the learning rate for training.

The metadata is an optional string containing additional information about the training.`,
		Args: cobra.RangeArgs(5, 6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			networkID := args[0]

			// Read features file
			featuresBytes, err := os.ReadFile(args[1])
			if err != nil {
				return fmt.Errorf("failed to read features file: %w", err)
			}

			// Read labels file
			labelsBytes, err := os.ReadFile(args[2])
			if err != nil {
				return fmt.Errorf("failed to read labels file: %w", err)
			}

			// Parse epochs
			epochs, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid epochs value: %w", err)
			}

			// Parse learning rate
			learningRate, err := sdk.NewDecFromStr(args[4])
			if err != nil {
				return fmt.Errorf("invalid learning rate: %w", err)
			}

			// Get metadata if provided
			var metadata []byte
			if len(args) > 5 && args[5] != "" {
				metadata = []byte(args[5])
			}

			valAddr := sdk.ValAddress(clientCtx.GetFromAddress())

			msg := types.NewMsgTrainNeuralNetwork(
				valAddr,
				networkID,
				featuresBytes,
				labelsBytes,
				epochs,
				learningRate,
				metadata,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewSubmitNeuralPredictionCmd returns a CLI command handler for submitting a neural prediction
func NewSubmitNeuralPredictionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-neural-prediction [network-id] [input-file] [output-file] [confidence] [metadata]",
		Short: "Submit a neural prediction",
		Long: `Submit a prediction from a neural network.

The input-file should be a path to a file containing the input data in JSON format.

The output-file should be a path to a file containing the output data in JSON format.

The confidence parameter specifies the confidence level of the prediction (0-1).

The metadata is an optional string containing additional information about the prediction.`,
		Args: cobra.RangeArgs(4, 5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			networkID := args[0]

			// Read input file
			inputBytes, err := os.ReadFile(args[1])
			if err != nil {
				return fmt.Errorf("failed to read input file: %w", err)
			}

			// Read output file
			outputBytes, err := os.ReadFile(args[2])
			if err != nil {
				return fmt.Errorf("failed to read output file: %w", err)
			}

			// Parse confidence
			confidence, err := sdk.NewDecFromStr(args[3])
			if err != nil {
				return fmt.Errorf("invalid confidence value: %w", err)
			}

			// Get metadata if provided
			var metadata []byte
			if len(args) > 4 && args[4] != "" {
				metadata = []byte(args[4])
			}

			valAddr := sdk.ValAddress(clientCtx.GetFromAddress())

			msg := types.NewMsgSubmitNeuralPrediction(
				valAddr,
				networkID,
				inputBytes,
				outputBytes,
				confidence,
				metadata,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewUpdateValidatorReputationCmd returns a CLI command handler for updating a validator's reputation
func NewUpdateValidatorReputationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-validator-reputation [validator-addr] [reputation-change] [reason]",
		Short: "Update a validator's reputation",
		Long: `Update a validator's reputation (admin only).

The validator-addr is the address of the validator to update.

The reputation-change is the amount to change the reputation by (-1 to 1).

The reason is a string describing the reason for the reputation change.`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			valAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			// Parse reputation change
			reputationChange, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return fmt.Errorf("invalid reputation change value: %w", err)
			}

			reason := args[2]

			adminAddr := clientCtx.GetFromAddress()

			msg := types.NewMsgUpdateValidatorReputation(
				adminAddr,
				valAddr,
				reputationChange,
				reason,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}