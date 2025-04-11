package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/nomercychain/nmxchain/x/deai/types"
)

// GetQueryCmd returns the query commands for the deai module
func GetQueryCmd() *cobra.Command {
	deaiQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the DeAI module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	deaiQueryCmd.AddCommand(
		GetCmdQueryAIAgent(),
		GetCmdQueryAIAgents(),
		GetCmdQueryAIAgentState(),
		GetCmdQueryAIAgentActions(),
		GetCmdQueryAIAgentAction(),
		GetCmdQueryAIAgentModels(),
		GetCmdQueryAIAgentModel(),
		GetCmdQueryAIAgentTrainingData(),
		GetCmdQueryMarketplaceListings(),
		GetCmdQueryMarketplaceListing(),
	)

	return deaiQueryCmd
}

// GetCmdQueryAIAgent returns the command to query a specific AI agent
func GetCmdQueryAIAgent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent [id]",
		Short: "Query a specific AI agent by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryAIAgentRequest{
				ID: args[0],
			}

			res, err := queryClient.AIAgent(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAIAgents returns the command to query all AI agents
func GetCmdQueryAIAgents() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agents [owner]",
		Short: "Query all AI agents, optionally filtered by owner",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryAIAgentsRequest{}
			if len(args) > 0 {
				req.Owner = args[0]
			}

			res, err := queryClient.AIAgents(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAIAgentState returns the command to query the state of an AI agent
func GetCmdQueryAIAgentState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent-state [agent-id]",
		Short: "Query the state of an AI agent",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryAIAgentStateRequest{
				AgentID: args[0],
			}

			res, err := queryClient.AIAgentState(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAIAgentActions returns the command to query actions for an AI agent
func GetCmdQueryAIAgentActions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent-actions [agent-id]",
		Short: "Query actions for an AI agent",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryAIAgentActionsRequest{
				AgentID: args[0],
			}

			res, err := queryClient.AIAgentActions(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAIAgentAction returns the command to query a specific AI agent action
func GetCmdQueryAIAgentAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent-action [id]",
		Short: "Query a specific AI agent action by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryAIAgentActionRequest{
				ID: args[0],
			}

			res, err := queryClient.AIAgentAction(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAIAgentModels returns the command to query all AI agent models
func GetCmdQueryAIAgentModels() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent-models",
		Short: "Query all AI agent models",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryAIAgentModelsRequest{}

			res, err := queryClient.AIAgentModels(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAIAgentModel returns the command to query a specific AI agent model
func GetCmdQueryAIAgentModel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent-model [id]",
		Short: "Query a specific AI agent model by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryAIAgentModelRequest{
				ID: args[0],
			}

			res, err := queryClient.AIAgentModel(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAIAgentTrainingData returns the command to query training data for an AI agent
func GetCmdQueryAIAgentTrainingData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent-training-data [agent-id]",
		Short: "Query training data for an AI agent",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryAIAgentTrainingDataRequest{
				AgentID: args[0],
			}

			res, err := queryClient.AIAgentTrainingData(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryMarketplaceListings returns the command to query marketplace listings
func GetCmdQueryMarketplaceListings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "marketplace-listings [status] [listing-type]",
		Short: "Query marketplace listings, optionally filtered by status and listing type",
		Args:  cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryMarketplaceListingsRequest{}
			if len(args) > 0 {
				req.Status = args[0]
			}
			if len(args) > 1 {
				req.ListingType = args[1]
			}

			res, err := queryClient.MarketplaceListings(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryMarketplaceListing returns the command to query a specific marketplace listing
func GetCmdQueryMarketplaceListing() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "marketplace-listing [id]",
		Short: "Query a specific marketplace listing by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryMarketplaceListingRequest{
				ID: args[0],
			}

			res, err := queryClient.MarketplaceListing(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}