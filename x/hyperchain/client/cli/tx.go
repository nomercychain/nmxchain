package cli

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/hyperchain/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "hyperchain",
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateHyperchain())
	cmd.AddCommand(CmdUpdateHyperchain())
	cmd.AddCommand(CmdJoinHyperchainAsValidator())
	cmd.AddCommand(CmdLeaveHyperchain())
	cmd.AddCommand(CmdCreateHyperchainBridge())
	cmd.AddCommand(CmdUpdateHyperchainBridge())
	cmd.AddCommand(CmdRegisterHyperchainBridgeRelayer())
	cmd.AddCommand(CmdRemoveHyperchainBridgeRelayer())
	cmd.AddCommand(CmdInitiateHyperchainBridgeTransaction())
	cmd.AddCommand(CmdApproveHyperchainBridgeTransaction())
	cmd.AddCommand(CmdSubmitHyperchainBlock())
	cmd.AddCommand(CmdSubmitHyperchainTransaction())
	cmd.AddCommand(CmdGrantHyperchainPermission())
	cmd.AddCommand(CmdRevokeHyperchainPermission())

	return cmd
}

// CmdCreateHyperchain implements the create hyperchain command
func CmdCreateHyperchain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-chain [name] [description] [chain-type] [consensus-type] [max-validators] [min-stake] [deposit] [supported-tokens] [supported-modules] [parent-chain-id] [agent-id]",
		Short: "Create a new hyperchain",
		Args:  cobra.RangeArgs(7, 11),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]
			description := args[1]

			// Parse chain type
			chainTypeStr := strings.ToUpper(args[2])
			var chainType types.HyperchainType
			switch chainTypeStr {
			case "STANDARD":
				chainType = types.HyperchainTypeStandard
			case "SPECIALIZED":
				chainType = types.HyperchainTypeSpecialized
			case "PRIVATE":
				chainType = types.HyperchainTypePrivate
			case "ENTERPRISE":
				chainType = types.HyperchainTypeEnterprise
			case "CONSORTIUM":
				chainType = types.HyperchainTypeConsortium
			default:
				return fmt.Errorf("invalid chain type: %s", chainTypeStr)
			}

			// Parse consensus type
			consensusTypeStr := strings.ToUpper(args[3])
			var consensusType types.ConsensusType
			switch consensusTypeStr {
			case "TENDERMINT":
				consensusType = types.ConsensusTypeTendermint
			case "NEUROPOS":
				consensusType = types.ConsensusTypeNeuroPOS
			case "POA":
				consensusType = types.ConsensusTypePOA
			case "HYBRID":
				consensusType = types.ConsensusTypeHybrid
			case "CUSTOM":
				consensusType = types.ConsensusTypeCustom
			default:
				return fmt.Errorf("invalid consensus type: %s", consensusTypeStr)
			}

			// Parse max validators
			maxValidators, err := strconv.ParseUint(args[4], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid max validators: %s", err)
			}

			// Parse min stake
			minStake, err := strconv.ParseUint(args[5], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid min stake: %s", err)
			}

			// Parse deposit
			deposit, err := sdk.ParseCoinNormalized(args[6])
			if err != nil {
				return fmt.Errorf("invalid deposit: %s", err)
			}

			// Parse supported tokens
			var supportedTokens []string
			if len(args) > 7 && args[7] != "" {
				supportedTokens = strings.Split(args[7], ",")
			}

			// Parse supported modules
			var supportedModules []string
			if len(args) > 8 && args[8] != "" {
				supportedModules = strings.Split(args[8], ",")
			}

			// Parse parent chain ID
			parentChainID := ""
			if len(args) > 9 && args[9] != "" {
				parentChainID = args[9]
			}

			// Parse agent ID
			agentID := ""
			if len(args) > 10 && args[10] != "" {
				agentID = args[10]
			}

			// Get genesis config from file
			genesisConfigFile, _ := cmd.Flags().GetString("genesis-config")
			var genesisConfig []byte
			if genesisConfigFile != "" {
				genesisConfig, err = json.Marshal(genesisConfigFile)
				if err != nil {
					return fmt.Errorf("invalid genesis config: %s", err)
				}
			}

			// Get chain config from file
			chainConfigFile, _ := cmd.Flags().GetString("chain-config")
			var chainConfig []byte
			if chainConfigFile != "" {
				chainConfig, err = json.Marshal(chainConfigFile)
				if err != nil {
					return fmt.Errorf("invalid chain config: %s", err)
				}
			}

			// Get metadata from file
			metadataFile, _ := cmd.Flags().GetString("metadata")
			var metadata []byte
			if metadataFile != "" {
				metadata, err = json.Marshal(metadataFile)
				if err != nil {
					return fmt.Errorf("invalid metadata: %s", err)
				}
			}

			msg := types.MsgCreateHyperchain{
				Creator:          clientCtx.GetFromAddress().String(),
				Name:             name,
				Description:      description,
				ChainType:        chainType,
				ConsensusType:    consensusType,
				GenesisConfig:    genesisConfig,
				ChainConfig:      chainConfig,
				Metadata:         metadata,
				MaxValidators:    maxValidators,
				MinStake:         minStake,
				ParentChainId:    parentChainID,
				SupportedTokens:  supportedTokens,
				SupportedModules: supportedModules,
				AgentId:          agentID,
				Deposit:          deposit,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String("genesis-config", "", "Genesis configuration file path")
	cmd.Flags().String("chain-config", "", "Chain configuration file path")
	cmd.Flags().String("metadata", "", "Metadata file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdUpdateHyperchain implements the update hyperchain command
func CmdUpdateHyperchain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-chain [chain-id] [name] [description] [max-validators] [min-stake] [supported-tokens] [supported-modules] [agent-id]",
		Short: "Update an existing hyperchain",
		Args:  cobra.RangeArgs(3, 8),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainID := args[0]
			name := args[1]
			description := args[2]

			// Parse max validators
			var maxValidators uint64
			if len(args) > 3 && args[3] != "" {
				maxValidators, err = strconv.ParseUint(args[3], 10, 64)
				if err != nil {
					return fmt.Errorf("invalid max validators: %s", err)
				}
			}

			// Parse min stake
			var minStake uint64
			if len(args) > 4 && args[4] != "" {
				minStake, err = strconv.ParseUint(args[4], 10, 64)
				if err != nil {
					return fmt.Errorf("invalid min stake: %s", err)
				}
			}

			// Parse supported tokens
			var supportedTokens []string
			if len(args) > 5 && args[5] != "" {
				supportedTokens = strings.Split(args[5], ",")
			}

			// Parse supported modules
			var supportedModules []string
			if len(args) > 6 && args[6] != "" {
				supportedModules = strings.Split(args[6], ",")
			}

			// Parse agent ID
			agentID := ""
			if len(args) > 7 && args[7] != "" {
				agentID = args[7]
			}

			// Get chain config from file
			chainConfigFile, _ := cmd.Flags().GetString("chain-config")
			var chainConfig []byte
			if chainConfigFile != "" {
				chainConfig, err = json.Marshal(chainConfigFile)
				if err != nil {
					return fmt.Errorf("invalid chain config: %s", err)
				}
			}

			// Get metadata from file
			metadataFile, _ := cmd.Flags().GetString("metadata")
			var metadata []byte
			if metadataFile != "" {
				metadata, err = json.Marshal(metadataFile)
				if err != nil {
					return fmt.Errorf("invalid metadata: %s", err)
				}
			}

			msg := types.MsgUpdateHyperchain{
				Admin:            clientCtx.GetFromAddress().String(),
				ChainId:          chainID,
				Name:             name,
				Description:      description,
				ChainConfig:      chainConfig,
				Metadata:         metadata,
				MaxValidators:    maxValidators,
				MinStake:         minStake,
				SupportedTokens:  supportedTokens,
				SupportedModules: supportedModules,
				AgentId:          agentID,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String("chain-config", "", "Chain configuration file path")
	cmd.Flags().String("metadata", "", "Metadata file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdJoinHyperchainAsValidator implements the join hyperchain as validator command
func CmdJoinHyperchainAsValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "join-chain [chain-id] [pubkey] [stake]",
		Short: "Join a hyperchain as a validator",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainID := args[0]
			pubkey := args[1]

			// Parse stake
			stake, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return fmt.Errorf("invalid stake: %s", err)
			}

			// Get metadata from file
			metadataFile, _ := cmd.Flags().GetString("metadata")
			var metadata []byte
			if metadataFile != "" {
				metadata, err = json.Marshal(metadataFile)
				if err != nil {
					return fmt.Errorf("invalid metadata: %s", err)
				}
			}

			msg := types.MsgJoinHyperchainAsValidator{
				Validator: clientCtx.GetFromAddress().String(),
				ChainId:   chainID,
				Pubkey:    pubkey,
				Stake:     stake,
				Metadata:  metadata,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String("metadata", "", "Metadata file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdLeaveHyperchain implements the leave hyperchain command
func CmdLeaveHyperchain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "leave-chain [chain-id]",
		Short: "Leave a hyperchain as a validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainID := args[0]

			msg := types.MsgLeaveHyperchain{
				Validator: clientCtx.GetFromAddress().String(),
				ChainId:   chainID,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdCreateHyperchainBridge implements the create hyperchain bridge command
func CmdCreateHyperchainBridge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-bridge [source-chain-id] [target-chain-id] [min-relayers] [supported-tokens]",
		Short: "Create a new hyperchain bridge",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sourceChainID := args[0]
			targetChainID := args[1]

			// Parse min relayers
			minRelayers, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid min relayers: %s", err)
			}

			// Parse supported tokens
			supportedTokens := strings.Split(args[3], ",")

			// Get metadata from file
			metadataFile, _ := cmd.Flags().GetString("metadata")
			var metadata []byte
			if metadataFile != "" {
				metadata, err = json.Marshal(metadataFile)
				if err != nil {
					return fmt.Errorf("invalid metadata: %s", err)
				}
			}

			msg := types.MsgCreateHyperchainBridge{
				Creator:         clientCtx.GetFromAddress().String(),
				SourceChainId:   sourceChainID,
				TargetChainId:   targetChainID,
				MinRelayers:     minRelayers,
				SupportedTokens: supportedTokens,
				Metadata:        metadata,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String("metadata", "", "Metadata file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdUpdateHyperchainBridge implements the update hyperchain bridge command
func CmdUpdateHyperchainBridge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-bridge [bridge-id] [min-relayers] [supported-tokens]",
		Short: "Update an existing hyperchain bridge",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			bridgeID := args[0]

			// Parse min relayers
			minRelayers, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid min relayers: %s", err)
			}

			// Parse supported tokens
			supportedTokens := strings.Split(args[2], ",")

			// Get metadata from file
			metadataFile, _ := cmd.Flags().GetString("metadata")
			var metadata []byte
			if metadataFile != "" {
				metadata, err = json.Marshal(metadataFile)
				if err != nil {
					return fmt.Errorf("invalid metadata: %s", err)
				}
			}

			msg := types.MsgUpdateHyperchainBridge{
				Admin:           clientCtx.GetFromAddress().String(),
				BridgeId:        bridgeID,
				MinRelayers:     minRelayers,
				SupportedTokens: supportedTokens,
				Metadata:        metadata,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String("metadata", "", "Metadata file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdRegisterHyperchainBridgeRelayer implements the register hyperchain bridge relayer command
func CmdRegisterHyperchainBridgeRelayer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-relayer [bridge-id] [relayer]",
		Short: "Register a relayer for a hyperchain bridge",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			bridgeID := args[0]
			relayer := args[1]

			msg := types.MsgRegisterHyperchainBridgeRelayer{
				Admin:    clientCtx.GetFromAddress().String(),
				BridgeId: bridgeID,
				Relayer:  relayer,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdRemoveHyperchainBridgeRelayer implements the remove hyperchain bridge relayer command
func CmdRemoveHyperchainBridgeRelayer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-relayer [bridge-id] [relayer]",
		Short: "Remove a relayer from a hyperchain bridge",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			bridgeID := args[0]
			relayer := args[1]

			msg := types.MsgRemoveHyperchainBridgeRelayer{
				Admin:    clientCtx.GetFromAddress().String(),
				BridgeId: bridgeID,
				Relayer:  relayer,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdInitiateHyperchainBridgeTransaction implements the initiate hyperchain bridge transaction command
func CmdInitiateHyperchainBridgeTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "initiate-bridge-tx [bridge-id] [recipient] [amount] [source-tx-id]",
		Short: "Initiate a transaction through a hyperchain bridge",
		Args:  cobra.RangeArgs(3, 4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			bridgeID := args[0]
			recipient := args[1]

			// Parse amount
			amount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return fmt.Errorf("invalid amount: %s", err)
			}

			// Parse source transaction ID
			sourceTxID := ""
			if len(args) > 3 && args[3] != "" {
				sourceTxID = args[3]
			}

			// Get metadata from file
			metadataFile, _ := cmd.Flags().GetString("metadata")
			var metadata []byte
			if metadataFile != "" {
				metadata, err = json.Marshal(metadataFile)
				if err != nil {
					return fmt.Errorf("invalid metadata: %s", err)
				}
			}

			msg := types.MsgInitiateHyperchainBridgeTransaction{
				Sender:     clientCtx.GetFromAddress().String(),
				BridgeId:   bridgeID,
				Recipient:  recipient,
				Amount:     amount,
				SourceTxId: sourceTxID,
				Metadata:   metadata,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String("metadata", "", "Metadata file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdApproveHyperchainBridgeTransaction implements the approve hyperchain bridge transaction command
func CmdApproveHyperchainBridgeTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-bridge-tx [bridge-id] [tx-id]",
		Short: "Approve a transaction through a hyperchain bridge",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			bridgeID := args[0]
			txID := args[1]

			msg := types.MsgApproveHyperchainBridgeTransaction{
				Relayer:  clientCtx.GetFromAddress().String(),
				BridgeId: bridgeID,
				TxId:     txID,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdSubmitHyperchainBlock implements the submit hyperchain block command
func CmdSubmitHyperchainBlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-block [chain-id] [height] [parent-hash] [num-txs]",
		Short: "Submit a block to a hyperchain",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainID := args[0]

			// Parse height
			height, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid height: %s", err)
			}

			parentHash := args[2]

			// Parse num txs
			numTxs, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid num txs: %s", err)
			}

			// Get data from file
			dataFile, _ := cmd.Flags().GetString("data")
			var data []byte
			if dataFile != "" {
				data, err = json.Marshal(dataFile)
				if err != nil {
					return fmt.Errorf("invalid data: %s", err)
				}
			}

			msg := types.MsgSubmitHyperchainBlock{
				Proposer:   clientCtx.GetFromAddress().String(),
				ChainId:    chainID,
				Height:     height,
				ParentHash: parentHash,
				NumTxs:     numTxs,
				Data:       data,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String("data", "", "Block data file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdSubmitHyperchainTransaction implements the submit hyperchain transaction command
func CmdSubmitHyperchainTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-tx [chain-id] [recipient] [amount]",
		Short: "Submit a transaction to a hyperchain",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainID := args[0]
			recipient := args[1]

			// Parse amount
			amount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return fmt.Errorf("invalid amount: %s", err)
			}

			// Get data from file
			dataFile, _ := cmd.Flags().GetString("data")
			var data []byte
			if dataFile != "" {
				data, err = json.Marshal(dataFile)
				if err != nil {
					return fmt.Errorf("invalid data: %s", err)
				}
			}

			msg := types.MsgSubmitHyperchainTransaction{
				Sender:    clientCtx.GetFromAddress().String(),
				ChainId:   chainID,
				Recipient: recipient,
				Amount:    amount,
				Data:      data,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String("data", "", "Transaction data file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdGrantHyperchainPermission implements the grant hyperchain permission command
func CmdGrantHyperchainPermission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant-permission [chain-id] [address] [permission-type] [expiration-days]",
		Short: "Grant permission to a hyperchain",
		Args:  cobra.RangeArgs(3, 4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainID := args[0]
			address := args[1]
			permissionType := args[2]

			// Parse expiration days
			var expirationDays uint64
			if len(args) > 3 && args[3] != "" {
				expirationDays, err = strconv.ParseUint(args[3], 10, 64)
				if err != nil {
					return fmt.Errorf("invalid expiration days: %s", err)
				}
			}

			// Get metadata from file
			metadataFile, _ := cmd.Flags().GetString("metadata")
			var metadata []byte
			if metadataFile != "" {
				metadata, err = json.Marshal(metadataFile)
				if err != nil {
					return fmt.Errorf("invalid metadata: %s", err)
				}
			}

			msg := types.MsgGrantHyperchainPermission{
				Admin:          clientCtx.GetFromAddress().String(),
				ChainId:        chainID,
				Address:        address,
				PermissionType: permissionType,
				ExpirationDays: expirationDays,
				Metadata:       metadata,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String("metadata", "", "Metadata file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdRevokeHyperchainPermission implements the revoke hyperchain permission command
func CmdRevokeHyperchainPermission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-permission [chain-id] [address] [permission-type]",
		Short: "Revoke permission from a hyperchain",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainID := args[0]
			address := args[1]
			permissionType := args[2]

			msg := types.MsgRevokeHyperchainPermission{
				Admin:          clientCtx.GetFromAddress().String(),
				ChainId:        chainID,
				Address:        address,
				PermissionType: permissionType,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}