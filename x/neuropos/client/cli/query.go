package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/nomercychain/nmxchain/x/neuropos/types"
)

// GetQueryCmd returns the query commands for this module
func GetQueryCmd() *cobra.Command {
	neuroposQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	neuroposQueryCmd.AddCommand(
		NewQueryParamsCmd(),
		NewQueryValidatorCmd(),
		NewQueryValidatorsCmd(),
		NewQueryDelegationCmd(),
		NewQueryDelegatorDelegationsCmd(),
		NewQueryValidatorDelegationsCmd(),
		NewQueryUnbondingDelegationCmd(),
		NewQueryDelegatorUnbondingDelegationsCmd(),
		NewQueryValidatorUnbondingDelegationsCmd(),
		NewQueryRedelegationCmd(),
		NewQueryDelegatorRedelegationsCmd(),
		NewQueryValidatorRedelegationsCmd(),
		NewQueryNeuralNetworkCmd(),
		NewQueryNeuralNetworksCmd(),
		NewQueryNeuralNetworkWeightsCmd(),
		NewQueryTrainingDataCmd(),
		NewQueryNetworkTrainingDataCmd(),
		NewQueryNeuralPredictionCmd(),
		NewQueryNetworkNeuralPredictionsCmd(),
		NewQueryValidatorPerformanceCmd(),
		NewQueryValidatorPerformancesCmd(),
		NewQueryValidatorReputationCmd(),
		NewQueryValidatorReputationsCmd(),
		NewQueryValidatorSlashEventsCmd(),
		NewQueryNetworkStateCmd(),
		NewQueryAnomalyReportsCmd(),
	)

	return neuroposQueryCmd
}

// NewQueryParamsCmd returns a CLI command handler for querying module parameters
func NewQueryParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query the current neuropos parameters",
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

// NewQueryValidatorCmd returns a CLI command handler for querying a validator
func NewQueryValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator [validator-addr]",
		Short: "Query a validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Validator(cmd.Context(), &types.QueryValidatorRequest{
				ValidatorAddr: args[0],
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

// NewQueryValidatorsCmd returns a CLI command handler for querying all validators
func NewQueryValidatorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validators",
		Short: "Query all validators",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Validators(cmd.Context(), &types.QueryValidatorsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryDelegationCmd returns a CLI command handler for querying a delegation
func NewQueryDelegationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegation [delegator-addr] [validator-addr]",
		Short: "Query a delegation",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Delegation(cmd.Context(), &types.QueryDelegationRequest{
				DelegatorAddr: args[0],
				ValidatorAddr: args[1],
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

// NewQueryDelegatorDelegationsCmd returns a CLI command handler for querying all delegations for a delegator
func NewQueryDelegatorDelegationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegator-delegations [delegator-addr]",
		Short: "Query all delegations for a delegator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.DelegatorDelegations(cmd.Context(), &types.QueryDelegatorDelegationsRequest{
				DelegatorAddr: args[0],
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

// NewQueryValidatorDelegationsCmd returns a CLI command handler for querying all delegations to a validator
func NewQueryValidatorDelegationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-delegations [validator-addr]",
		Short: "Query all delegations to a validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ValidatorDelegations(cmd.Context(), &types.QueryValidatorDelegationsRequest{
				ValidatorAddr: args[0],
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

// NewQueryUnbondingDelegationCmd returns a CLI command handler for querying an unbonding delegation
func NewQueryUnbondingDelegationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unbonding-delegation [delegator-addr] [validator-addr]",
		Short: "Query an unbonding delegation",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.UnbondingDelegation(cmd.Context(), &types.QueryUnbondingDelegationRequest{
				DelegatorAddr: args[0],
				ValidatorAddr: args[1],
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

// NewQueryDelegatorUnbondingDelegationsCmd returns a CLI command handler for querying all unbonding delegations for a delegator
func NewQueryDelegatorUnbondingDelegationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegator-unbonding-delegations [delegator-addr]",
		Short: "Query all unbonding delegations for a delegator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.DelegatorUnbondingDelegations(cmd.Context(), &types.QueryDelegatorUnbondingDelegationsRequest{
				DelegatorAddr: args[0],
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

// NewQueryValidatorUnbondingDelegationsCmd returns a CLI command handler for querying all unbonding delegations from a validator
func NewQueryValidatorUnbondingDelegationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-unbonding-delegations [validator-addr]",
		Short: "Query all unbonding delegations from a validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ValidatorUnbondingDelegations(cmd.Context(), &types.QueryValidatorUnbondingDelegationsRequest{
				ValidatorAddr: args[0],
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

// NewQueryRedelegationCmd returns a CLI command handler for querying a redelegation
func NewQueryRedelegationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redelegation [delegator-addr] [src-validator-addr] [dst-validator-addr]",
		Short: "Query a redelegation",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Redelegation(cmd.Context(), &types.QueryRedelegationRequest{
				DelegatorAddr:    args[0],
				SrcValidatorAddr: args[1],
				DstValidatorAddr: args[2],
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

// NewQueryDelegatorRedelegationsCmd returns a CLI command handler for querying all redelegations for a delegator
func NewQueryDelegatorRedelegationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegator-redelegations [delegator-addr]",
		Short: "Query all redelegations for a delegator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.DelegatorRedelegations(cmd.Context(), &types.QueryDelegatorRedelegationsRequest{
				DelegatorAddr: args[0],
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

// NewQueryValidatorRedelegationsCmd returns a CLI command handler for querying all redelegations from a validator
func NewQueryValidatorRedelegationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-redelegations [validator-addr]",
		Short: "Query all redelegations from a validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ValidatorRedelegations(cmd.Context(), &types.QueryValidatorRedelegationsRequest{
				ValidatorAddr: args[0],
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

// NewQueryNeuralNetworkCmd returns a CLI command handler for querying a neural network
func NewQueryNeuralNetworkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "neural-network [network-id]",
		Short: "Query a neural network",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.NeuralNetwork(cmd.Context(), &types.QueryNeuralNetworkRequest{
				NetworkId: args[0],
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

// NewQueryNeuralNetworksCmd returns a CLI command handler for querying all neural networks
func NewQueryNeuralNetworksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "neural-networks",
		Short: "Query all neural networks",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.NeuralNetworks(cmd.Context(), &types.QueryNeuralNetworksRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryNeuralNetworkWeightsCmd returns a CLI command handler for querying neural network weights
func NewQueryNeuralNetworkWeightsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "neural-network-weights [network-id] [version]",
		Short: "Query neural network weights",
		Long:  "Query neural network weights. If version is not provided, the latest version will be returned.",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			var version uint64
			if len(args) > 1 {
				version, err = strconv.ParseUint(args[1], 10, 64)
				if err != nil {
					return fmt.Errorf("invalid version: %w", err)
				}
			}

			res, err := queryClient.NeuralNetworkWeights(cmd.Context(), &types.QueryNeuralNetworkWeightsRequest{
				NetworkId: args[0],
				Version:   version,
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

// NewQueryTrainingDataCmd returns a CLI command handler for querying training data
func NewQueryTrainingDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "training-data [data-id]",
		Short: "Query training data",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.TrainingData(cmd.Context(), &types.QueryTrainingDataRequest{
				DataId: args[0],
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

// NewQueryNetworkTrainingDataCmd returns a CLI command handler for querying all training data for a neural network
func NewQueryNetworkTrainingDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "network-training-data [network-id]",
		Short: "Query all training data for a neural network",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.NetworkTrainingData(cmd.Context(), &types.QueryNetworkTrainingDataRequest{
				NetworkId: args[0],
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

// NewQueryNeuralPredictionCmd returns a CLI command handler for querying a neural prediction
func NewQueryNeuralPredictionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "neural-prediction [prediction-id]",
		Short: "Query a neural prediction",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.NeuralPrediction(cmd.Context(), &types.QueryNeuralPredictionRequest{
				PredictionId: args[0],
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

// NewQueryNetworkNeuralPredictionsCmd returns a CLI command handler for querying all neural predictions for a neural network
func NewQueryNetworkNeuralPredictionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "network-neural-predictions [network-id]",
		Short: "Query all neural predictions for a neural network",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.NetworkNeuralPredictions(cmd.Context(), &types.QueryNetworkNeuralPredictionsRequest{
				NetworkId: args[0],
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

// NewQueryValidatorPerformanceCmd returns a CLI command handler for querying a validator's performance
func NewQueryValidatorPerformanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-performance [validator-addr]",
		Short: "Query a validator's performance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ValidatorPerformance(cmd.Context(), &types.QueryValidatorPerformanceRequest{
				ValidatorAddr: args[0],
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

// NewQueryValidatorPerformancesCmd returns a CLI command handler for querying all validator performances
func NewQueryValidatorPerformancesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-performances",
		Short: "Query all validator performances",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ValidatorPerformances(cmd.Context(), &types.QueryValidatorPerformancesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryValidatorReputationCmd returns a CLI command handler for querying a validator's reputation
func NewQueryValidatorReputationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-reputation [validator-addr]",
		Short: "Query a validator's reputation",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ValidatorReputation(cmd.Context(), &types.QueryValidatorReputationRequest{
				ValidatorAddr: args[0],
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

// NewQueryValidatorReputationsCmd returns a CLI command handler for querying all validator reputations
func NewQueryValidatorReputationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-reputations",
		Short: "Query all validator reputations",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ValidatorReputations(cmd.Context(), &types.QueryValidatorReputationsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryValidatorSlashEventsCmd returns a CLI command handler for querying a validator's slash events
func NewQueryValidatorSlashEventsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-slash-events [validator-addr]",
		Short: "Query a validator's slash events",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ValidatorSlashEvents(cmd.Context(), &types.QueryValidatorSlashEventsRequest{
				ValidatorAddr: args[0],
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

// NewQueryNetworkStateCmd returns a CLI command handler for querying the network state
func NewQueryNetworkStateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "network-state",
		Short: "Query the current network state",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.NetworkState(cmd.Context(), &types.QueryNetworkStateRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// NewQueryAnomalyReportsCmd returns a CLI command handler for querying anomaly reports
func NewQueryAnomalyReportsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "anomaly-reports",
		Short: "Query all anomaly reports",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AnomalyReports(cmd.Context(), &types.QueryAnomalyReportsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
