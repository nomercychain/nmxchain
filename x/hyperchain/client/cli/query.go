package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/nomercychain/nmxchain/x/hyperchain/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group hyperchain queries under a subcommand
	cmd := &cobra.Command{
		Use:                        "hyperchain",
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryHyperchain())
	cmd.AddCommand(CmdQueryHyperchains())
	cmd.AddCommand(CmdQueryHyperchainsByCreator())
	cmd.AddCommand(CmdQueryHyperchainsByParent())
	cmd.AddCommand(CmdQueryHyperchainValidator())
	cmd.AddCommand(CmdQueryHyperchainValidators())
	cmd.AddCommand(CmdQueryHyperchainBlock())
	cmd.AddCommand(CmdQueryHyperchainBlocks())
	cmd.AddCommand(CmdQueryHyperchainTransaction())
	cmd.AddCommand(CmdQueryHyperchainTransactions())
	cmd.AddCommand(CmdQueryHyperchainBridge())
	cmd.AddCommand(CmdQueryHyperchainBridges())
	cmd.AddCommand(CmdQueryHyperchainBridgesByChain())
	cmd.AddCommand(CmdQueryHyperchainBridgeTransaction())
	cmd.AddCommand(CmdQueryHyperchainBridgeTransactions())
	cmd.AddCommand(CmdQueryHyperchainPermissions())
	cmd.AddCommand(CmdQueryHyperchainPermission())

	return cmd
}

// CmdQueryHyperchain implements the query hyperchain command
func CmdQueryHyperchain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chain [id]",
		Short: "Query a hyperchain by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryHyperchainRequest{
				Id: args[0],
			}

			res, err := queryClient.Hyperchain(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchains implements the query hyperchains command
func CmdQueryHyperchains() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chains",
		Short: "Query all hyperchains",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryHyperchainsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.Hyperchains(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "hyperchains")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainsByCreator implements the query hyperchains by creator command
func CmdQueryHyperchainsByCreator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chains-by-creator [creator]",
		Short: "Query hyperchains by creator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryHyperchainsByCreatorRequest{
				Creator:    args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.HyperchainsByCreator(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "hyperchains by creator")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainsByParent implements the query hyperchains by parent command
func CmdQueryHyperchainsByParent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chains-by-parent [parent-chain-id]",
		Short: "Query hyperchains by parent chain ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryHyperchainsByParentRequest{
				ParentChainId: args[0],
				Pagination:    pageReq,
			}

			res, err := queryClient.HyperchainsByParent(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "hyperchains by parent")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainValidator implements the query hyperchain validator command
func CmdQueryHyperchainValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator [chain-id] [address]",
		Short: "Query a validator in a hyperchain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryHyperchainValidatorRequest{
				ChainId: args[0],
				Address: args[1],
			}

			res, err := queryClient.HyperchainValidator(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainValidators implements the query hyperchain validators command
func CmdQueryHyperchainValidators() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validators [chain-id]",
		Short: "Query all validators in a hyperchain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryHyperchainValidatorsRequest{
				ChainId:    args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.HyperchainValidators(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "validators")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainBlock implements the query hyperchain block command
func CmdQueryHyperchainBlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block [chain-id] [height]",
		Short: "Query a block in a hyperchain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			height, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryHyperchainBlockRequest{
				ChainId: args[0],
				Height:  height,
			}

			res, err := queryClient.HyperchainBlock(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainBlocks implements the query hyperchain blocks command
func CmdQueryHyperchainBlocks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blocks [chain-id]",
		Short: "Query all blocks in a hyperchain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryHyperchainBlocksRequest{
				ChainId:    args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.HyperchainBlocks(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "blocks")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainTransaction implements the query hyperchain transaction command
func CmdQueryHyperchainTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transaction [id]",
		Short: "Query a transaction in a hyperchain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryHyperchainTransactionRequest{
				Id: args[0],
			}

			res, err := queryClient.HyperchainTransaction(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainTransactions implements the query hyperchain transactions command
func CmdQueryHyperchainTransactions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transactions [chain-id]",
		Short: "Query all transactions in a hyperchain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryHyperchainTransactionsRequest{
				ChainId:    args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.HyperchainTransactions(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "transactions")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainBridge implements the query hyperchain bridge command
func CmdQueryHyperchainBridge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridge [id]",
		Short: "Query a hyperchain bridge by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryHyperchainBridgeRequest{
				Id: args[0],
			}

			res, err := queryClient.HyperchainBridge(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainBridges implements the query hyperchain bridges command
func CmdQueryHyperchainBridges() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridges",
		Short: "Query all hyperchain bridges",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryHyperchainBridgesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.HyperchainBridges(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "bridges")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainBridgesByChain implements the query hyperchain bridges by chain command
func CmdQueryHyperchainBridgesByChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridges-by-chain [chain-id]",
		Short: "Query hyperchain bridges by chain ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryHyperchainBridgesByChainRequest{
				ChainId:    args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.HyperchainBridgesByChain(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "bridges by chain")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainBridgeTransaction implements the query hyperchain bridge transaction command
func CmdQueryHyperchainBridgeTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridge-transaction [id]",
		Short: "Query a hyperchain bridge transaction by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryHyperchainBridgeTransactionRequest{
				Id: args[0],
			}

			res, err := queryClient.HyperchainBridgeTransaction(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainBridgeTransactions implements the query hyperchain bridge transactions command
func CmdQueryHyperchainBridgeTransactions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridge-transactions [bridge-id]",
		Short: "Query hyperchain bridge transactions by bridge ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryHyperchainBridgeTransactionsRequest{
				BridgeId:   args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.HyperchainBridgeTransactions(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "bridge transactions")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainPermissions implements the query hyperchain permissions command
func CmdQueryHyperchainPermissions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "permissions [chain-id]",
		Short: "Query all permissions in a hyperchain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryHyperchainPermissionsRequest{
				ChainId:    args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.HyperchainPermissions(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "permissions")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryHyperchainPermission implements the query hyperchain permission command
func CmdQueryHyperchainPermission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "permission [chain-id] [address] [permission-type]",
		Short: "Query a permission in a hyperchain",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryHyperchainPermissionRequest{
				ChainId:        args[0],
				Address:        args[1],
				PermissionType: args[2],
			}

			res, err := queryClient.HyperchainPermission(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
