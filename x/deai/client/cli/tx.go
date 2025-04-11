package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/deai/types"
)

// GetTxCmd returns the transaction commands for the deai module
func GetTxCmd() *cobra.Command {
	deaiTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "DeAI transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	deaiTxCmd.AddCommand(
		NewCreateAIAgentCmd(),
		NewUpdateAIAgentCmd(),
		NewTrainAIAgentCmd(),
		NewExecuteAIAgentCmd(),
		NewListAIAgentForSaleCmd(),
		NewBuyAIAgentCmd(),
		NewRentAIAgentCmd(),
		NewCancelMarketListingCmd(),
	)

	return deaiTxCmd
}

// NewCreateAIAgentCmd returns a CLI command handler for creating an AI agent
func NewCreateAIAgentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-agent [name] [description] [agent-type] [model-id] [permissions-file] [metadata-file]",
		Short: "Create a new AI agent",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]
			description := args[1]
			agentType := types.AIAgentType(args[2])
			modelID := args[3]
			permissionsFile := args[4]
			metadataFile := args[5]

			// Read permissions from file
			permissionsBytes, err := ioutil.ReadFile(permissionsFile)
			if err != nil {
				return fmt.Errorf("failed to read permissions file: %w", err)
			}

			// Read metadata from file
			metadataBytes, err := ioutil.ReadFile(metadataFile)
			if err != nil {
				return fmt.Errorf("failed to read metadata file: %w", err)
			}

			// Validate JSON
			var permissionsJSON, metadataJSON map[string]interface{}
			if err := json.Unmarshal(permissionsBytes, &permissionsJSON); err != nil {
				return fmt.Errorf("invalid permissions JSON: %w", err)
			}
			if err := json.Unmarshal(metadataBytes, &metadataJSON); err != nil {
				return fmt.Errorf("invalid metadata JSON: %w", err)
			}

			msg := types.NewMsgCreateAIAgent(
				clientCtx.GetFromAddress(),
				name,
				description,
				agentType,
				modelID,
				permissionsBytes,
				metadataBytes,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewUpdateAIAgentCmd returns a CLI command handler for updating an AI agent
func NewUpdateAIAgentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-agent [agent-id] [name] [description] [permissions-file] [metadata-file]",
		Short: "Update an existing AI agent",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			agentID := args[0]
			name := args[1]
			description := args[2]
			permissionsFile := args[3]
			metadataFile := args[4]

			// Read permissions from file
			permissionsBytes, err := ioutil.ReadFile(permissionsFile)
			if err != nil {
				return fmt.Errorf("failed to read permissions file: %w", err)
			}

			// Read metadata from file
			metadataBytes, err := ioutil.ReadFile(metadataFile)
			if err != nil {
				return fmt.Errorf("failed to read metadata file: %w", err)
			}

			// Validate JSON
			var permissionsJSON, metadataJSON map[string]interface{}
			if err := json.Unmarshal(permissionsBytes, &permissionsJSON); err != nil {
				return fmt.Errorf("invalid permissions JSON: %w", err)
			}
			if err := json.Unmarshal(metadataBytes, &metadataJSON); err != nil {
				return fmt.Errorf("invalid metadata JSON: %w", err)
			}

			msg := types.NewMsgUpdateAIAgent(
				clientCtx.GetFromAddress(),
				agentID,
				name,
				description,
				permissionsBytes,
				metadataBytes,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewTrainAIAgentCmd returns a CLI command handler for training an AI agent
func NewTrainAIAgentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "train-agent [agent-id] [data-type] [data-file] [source]",
		Short: "Train an AI agent with new data",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			agentID := args[0]
			dataType := args[1]
			dataFile := args[2]
			source := args[3]

			// Read data from file
			dataBytes, err := ioutil.ReadFile(dataFile)
			if err != nil {
				return fmt.Errorf("failed to read data file: %w", err)
			}

			// Validate JSON
			var dataJSON interface{}
			if err := json.Unmarshal(dataBytes, &dataJSON); err != nil {
				return fmt.Errorf("invalid data JSON: %w", err)
			}

			msg := types.NewMsgTrainAIAgent(
				clientCtx.GetFromAddress(),
				agentID,
				dataType,
				dataBytes,
				source,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewExecuteAIAgentCmd returns a CLI command handler for executing an AI agent
func NewExecuteAIAgentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute-agent [agent-id] [action-type] [data-file] [fee]",
		Short: "Execute an action using an AI agent",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			agentID := args[0]
			actionType := args[1]
			dataFile := args[2]
			fee, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return fmt.Errorf("invalid fee: %w", err)
			}

			// Read data from file
			dataBytes, err := ioutil.ReadFile(dataFile)
			if err != nil {
				return fmt.Errorf("failed to read data file: %w", err)
			}

			// Validate JSON
			var dataJSON interface{}
			if err := json.Unmarshal(dataBytes, &dataJSON); err != nil {
				return fmt.Errorf("invalid data JSON: %w", err)
			}

			msg := types.NewMsgExecuteAIAgent(
				clientCtx.GetFromAddress(),
				agentID,
				actionType,
				dataBytes,
				fee,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewListAIAgentForSaleCmd returns a CLI command handler for listing an AI agent for sale
func NewListAIAgentForSaleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-agent [agent-id] [price] [rental-price] [rental-duration] [listing-type] [expiration-days]",
		Short: "List an AI agent for sale or rent on the marketplace",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			agentID := args[0]
			
			price, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return fmt.Errorf("invalid price: %w", err)
			}
			
			rentalPrice, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return fmt.Errorf("invalid rental price: %w", err)
			}
			
			rentalDuration, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid rental duration: %w", err)
			}
			
			listingType := args[4]
			if listingType != "sale" && listingType != "rent" && listingType != "both" {
				return fmt.Errorf("listing type must be 'sale', 'rent', or 'both'")
			}
			
			expirationDays, err := strconv.ParseUint(args[5], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid expiration days: %w", err)
			}

			msg := types.NewMsgListAIAgentForSale(
				clientCtx.GetFromAddress(),
				agentID,
				price,
				rentalPrice,
				rentalDuration,
				listingType,
				expirationDays,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewBuyAIAgentCmd returns a CLI command handler for buying an AI agent
func NewBuyAIAgentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-agent [listing-id]",
		Short: "Buy an AI agent from the marketplace",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			listingID := args[0]

			msg := types.NewMsgBuyAIAgent(
				clientCtx.GetFromAddress(),
				listingID,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewRentAIAgentCmd returns a CLI command handler for renting an AI agent
func NewRentAIAgentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rent-agent [listing-id] [duration]",
		Short: "Rent an AI agent from the marketplace",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			listingID := args[0]
			
			duration, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid duration: %w", err)
			}

			msg := types.NewMsgRentAIAgent(
				clientCtx.GetFromAddress(),
				listingID,
				duration,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewCancelMarketListingCmd returns a CLI command handler for canceling a marketplace listing
func NewCancelMarketListingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-listing [listing-id]",
		Short: "Cancel a marketplace listing",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			listingID := args[0]

			msg := types.NewMsgCancelMarketListing(
				clientCtx.GetFromAddress(),
				listingID,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}