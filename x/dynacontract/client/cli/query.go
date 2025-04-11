package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/nomercychain/nmxchain/x/dynacontract/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	dynacontractQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	dynacontractQueryCmd.AddCommand(
		CmdListContracts(),
		CmdShowContract(),
		CmdQueryContract(),
		CmdGetContractHistory(),
		CmdGetContractState(),
	)

	return dynacontractQueryCmd
}

// CmdListContracts implements the list contracts command
func CmdListContracts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-contracts",
		Short: "List all contracts",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListContractsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ListContracts(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdShowContract implements the show contract command
func CmdShowContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-contract [id]",
		Short: "Shows a contract by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetContractRequest{
				Id: args[0],
			}

			res, err := queryClient.GetContract(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryContract implements the query contract command
func CmdQueryContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-contract [id] [query-data]",
		Short: "Query a contract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryContractRequest{
				Id:        args[0],
				QueryData: args[1],
			}

			res, err := queryClient.QueryContract(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdGetContractHistory implements the get contract history command
func CmdGetContractHistory() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-contract-history [id]",
		Short: "Get contract history",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryContractHistoryRequest{
				Id: args[0],
			}

			res, err := queryClient.GetContractHistory(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdGetContractState implements the get contract state command
func CmdGetContractState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-contract-state [id]",
		Short: "Get contract state",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryContractStateRequest{
				Id: args[0],
			}

			res, err := queryClient.GetContractState(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}