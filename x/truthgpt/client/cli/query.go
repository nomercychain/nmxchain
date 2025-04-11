package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/truthgpt/types"
)

// GetQueryCmd returns the query commands for this module
func GetQueryCmd() *cobra.Command {
	truthgptQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	truthgptQueryCmd.AddCommand(
		NewQueryParamsCmd(),
		NewQueryOracleProviderCmd(),
		NewQueryOracleProvidersCmd(),
		NewQueryOracleRequestCmd(),
		NewQueryOracleRequestsCmd(),
		NewQueryOracleResponseCmd(),
		NewQueryOracleResponsesCmd(),
		NewQueryRequestResponsesCmd(),
		NewQueryProviderResponsesCmd(),
		NewQueryProviderResponseHistoryCmd(),
		NewQueryDataSourceCmd(),
		NewQueryDataSourcesCmd(),
		NewQueryOracleScriptCmd(),
		NewQueryOracleScriptsCmd(),
		NewQueryPendingRequestsCmd(),
		NewQueryResultCmd(),
		NewQueryProviderRewardsCmd(),
		NewQueryProviderStatsCmd(),
	)

	return truthgptQueryCmd
}

// NewQueryParamsCmd returns a CLI command handler for querying module parameters
func NewQueryParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query the current truthgpt parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryOracleProviderCmd returns a CLI command handler for querying an oracle provider
func NewQueryOracleProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle-provider [address]",
		Short: "Query an oracle provider by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.OracleProvider(cmd.Context(), &types.QueryOracleProviderRequest{
				Address: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryOracleProvidersCmd returns a CLI command handler for querying all oracle providers
func NewQueryOracleProvidersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle-providers",
		Short: "Query all oracle providers",
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

			res, err := queryClient.OracleProviders(cmd.Context(), &types.QueryOracleProvidersRequest{
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "oracle-providers")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryOracleRequestCmd returns a CLI command handler for querying an oracle request
func NewQueryOracleRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle-request [id]",
		Short: "Query an oracle request by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, err := queryClient.OracleRequest(cmd.Context(), &types.QueryOracleRequestRequest{
				Id: id,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryOracleRequestsCmd returns a CLI command handler for querying all oracle requests
func NewQueryOracleRequestsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle-requests",
		Short: "Query all oracle requests",
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

			res, err := queryClient.OracleRequests(cmd.Context(), &types.QueryOracleRequestsRequest{
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "oracle-requests")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryOracleResponseCmd returns a CLI command handler for querying an oracle response
func NewQueryOracleResponseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle-response [id]",
		Short: "Query an oracle response by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, err := queryClient.OracleResponse(cmd.Context(), &types.QueryOracleResponseRequest{
				Id: id,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryOracleResponsesCmd returns a CLI command handler for querying all oracle responses
func NewQueryOracleResponsesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle-responses",
		Short: "Query all oracle responses",
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

			res, err := queryClient.OracleResponses(cmd.Context(), &types.QueryOracleResponsesRequest{
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "oracle-responses")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryRequestResponsesCmd returns a CLI command handler for querying all responses for a request
func NewQueryRequestResponsesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-responses [request-id]",
		Short: "Query all responses for a specific request",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			requestID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.RequestResponses(cmd.Context(), &types.QueryRequestResponsesRequest{
				RequestId:  requestID,
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "request-responses")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryProviderResponsesCmd returns a CLI command handler for querying all responses from a provider
func NewQueryProviderResponsesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider-responses [provider-address]",
		Short: "Query all responses from a specific provider",
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

			res, err := queryClient.ProviderResponses(cmd.Context(), &types.QueryProviderResponsesRequest{
				ProviderAddress: args[0],
				Pagination:      pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "provider-responses")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryProviderResponseHistoryCmd returns a CLI command handler for querying a provider's response history
func NewQueryProviderResponseHistoryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider-response-history [provider-address]",
		Short: "Query a provider's response history",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ProviderResponseHistory(cmd.Context(), &types.QueryProviderResponseHistoryRequest{
				ProviderAddress: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryDataSourceCmd returns a CLI command handler for querying a data source
func NewQueryDataSourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data-source [id]",
		Short: "Query a data source by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, err := queryClient.DataSource(cmd.Context(), &types.QueryDataSourceRequest{
				Id: id,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryDataSourcesCmd returns a CLI command handler for querying all data sources
func NewQueryDataSourcesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data-sources",
		Short: "Query all data sources",
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

			res, err := queryClient.DataSources(cmd.Context(), &types.QueryDataSourcesRequest{
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "data-sources")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryOracleScriptCmd returns a CLI command handler for querying an oracle script
func NewQueryOracleScriptCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle-script [id]",
		Short: "Query an oracle script by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, err := queryClient.OracleScript(cmd.Context(), &types.QueryOracleScriptRequest{
				Id: id,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryOracleScriptsCmd returns a CLI command handler for querying all oracle scripts
func NewQueryOracleScriptsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle-scripts",
		Short: "Query all oracle scripts",
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

			res, err := queryClient.OracleScripts(cmd.Context(), &types.QueryOracleScriptsRequest{
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "oracle-scripts")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryPendingRequestsCmd returns a CLI command handler for querying pending oracle requests
func NewQueryPendingRequestsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pending-requests [provider-address]",
		Short: "Query pending oracle requests for a provider",
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

			res, err := queryClient.PendingRequests(cmd.Context(), &types.QueryPendingRequestsRequest{
				ProviderAddress: args[0],
				Pagination:      pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "pending-requests")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryResultCmd returns a CLI command handler for querying the result of an oracle request
func NewQueryResultCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "result [request-id]",
		Short: "Query the result of an oracle request",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			requestID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, err := queryClient.Result(cmd.Context(), &types.QueryResultRequest{
				RequestId: requestID,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryProviderRewardsCmd returns a CLI command handler for querying a provider's rewards
func NewQueryProviderRewardsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider-rewards [provider-address]",
		Short: "Query an oracle provider's rewards",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ProviderRewards(cmd.Context(), &types.QueryProviderRewardsRequest{
				ProviderAddress: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryProviderStatsCmd returns a CLI command handler for querying a provider's statistics
func NewQueryProviderStatsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider-stats [provider-address]",
		Short: "Query an oracle provider's statistics",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ProviderStats(cmd.Context(), &types.QueryProviderStatsRequest{
				ProviderAddress: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}