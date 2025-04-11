# NoMercyChain Development Guide

This comprehensive guide will walk you through the process of setting up your development environment, exploring the codebase, implementing modules, developing the frontend, testing, and deploying a testnet for the NoMercyChain project.

## 1. Set Up Your Development Environment

### Install Go (version 1.18 or higher)

#### For Windows:
1. Download the installer from https://go.dev/dl/go1.18.10.windows-amd64.msi
2. Run the installer and follow the prompts
3. Open a new PowerShell window and verify the installation:
   ```powershell
   go version
   ```
4. Set up environment variables (if not set by the installer):
   - Right-click on "This PC" or "My Computer" and select "Properties"
   - Click on "Advanced system settings"
   - Click on "Environment Variables"
   - Under "System variables", find "Path" and click "Edit"
   - Add `C:\Go\bin` to the list
   - Add a new system variable named "GOPATH" with the value `%USERPROFILE%\go`
   - Add `%USERPROFILE%\go\bin` to the Path variable

#### For Linux:
```bash
wget https://go.dev/dl/go1.18.10.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.18.10.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

Add the following to your `~/.profile` or `~/.bashrc`:
```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

#### For macOS:
```bash
brew install go@1.18
```

Verify installation:
```bash
go version
```

### Install Node.js and npm (for the frontend)

#### For Windows:
1. Download the installer from https://nodejs.org/dist/v16.20.0/node-v16.20.0-x64.msi
2. Run the installer and follow the prompts
3. Open a new PowerShell window and verify the installation:
   ```powershell
   node -v
   npm -v
   ```

#### For Linux:
```bash
curl -fsSL https://deb.nodesource.com/setup_16.x | sudo -E bash -
sudo apt-get install -y nodejs
```

#### For macOS:
```bash
brew install node@16
```

Verify installation:
```bash
node -v
npm -v
```

### Clone the Repository

```bash
git clone https://github.com/nomercychain/nmxchain.git
cd nmxchain
```

### Set Up the Project

#### For Windows (PowerShell):
```powershell
# Install Go dependencies
go mod tidy

# Install frontend dependencies
cd client
npm install
cd ..

# Run the setup script
.\scripts\setup_dev_env.ps1
```

#### For Linux/macOS:
```bash
# Install Go dependencies
go mod tidy

# Install frontend dependencies
cd client
npm install
cd ..

# Run the setup script
chmod +x scripts/setup_dev_env.sh
./scripts/setup_dev_env.sh
```

## 2. Explore the Codebase

### Review Module Implementations

The NoMercyChain project has four core modules located in the `x/` directory:

1. **DynaContracts Module** (`x/dynacontracts/`): AI-powered smart contracts
   - Key files:
     - `types/`: Data structures and messages
     - `keeper/`: State management
     - `client/`: CLI commands

2. **DeAI Module** (`x/deai/`): User-controlled AI agents
   - Key files:
     - `types/`: Data structures and messages
     - `keeper/`: State management
     - `client/`: CLI commands

3. **TruthGPT Module** (`x/truthgpt/`): AI-powered oracles
   - Key files:
     - `types/`: Data structures and messages
     - `keeper/`: State management
     - `client/`: CLI commands
     - `abci.go`: BeginBlocker and EndBlocker implementations

4. **HyperChains Module** (`x/hyperchains/`): AI-generated Layer 3 chains
   - Key files:
     - `types/`: Data structures and messages
     - `keeper/`: State management
     - `client/`: CLI commands
     - `abci.go`: BeginBlocker and EndBlocker implementations

### Explore Frontend Components

The frontend is a React application located in the `client/src/` directory:

1. **Components** (`client/src/components/`):
   - `Navbar.js`: Top navigation bar
   - `Sidebar.js`: Side navigation menu
   - `Footer.js`: Page footer

2. **Pages** (`client/src/pages/`):
   - `Dashboard.js`: Main dashboard
   - `Wallet.js`: Wallet management
   - `AIAgents.js`: AI agent management
   - (Other pages to be implemented)

3. **Context** (`client/src/context/`):
   - `WalletContext.js`: Wallet state management

4. **Styles** (`client/src/App.css`, `client/src/index.css`):
   - Global styles and layout

### Read Documentation

Review the following documentation files:
- `SETUP.md`: Setup instructions
- `NEXT_STEPS.md`: Development roadmap and tasks

## 3. Complete Module Implementations

### Implement Message Handlers

For each module, you need to implement message handlers in the respective `handler.go` files:

#### Example for DynaContracts Module:

Create a file at `x/dynacontracts/handler.go`:

```go
package dynacontracts

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nomercychain/nmxchain/x/dynacontracts/keeper"
	"github.com/nomercychain/nmxchain/x/dynacontracts/types"
)

// NewHandler creates a new handler for dynacontracts module
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateContract:
			res, err := msgServer.CreateContract(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdateContract:
			res, err := msgServer.UpdateContract(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgExecuteContract:
			res, err := msgServer.ExecuteContract(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
```

Repeat this process for each module.

### Create Query Handlers

Implement query handlers in the respective `keeper/querier.go` files:

#### Example for DeAI Module:

Create a file at `x/deai/keeper/querier.go`:

```go
package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nomercychain/nmxchain/x/deai/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier creates a new querier for deai module
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryGetAgent:
			return getAgent(ctx, path[1:], k, legacyQuerierCdc)
		case types.QueryListAgents:
			return listAgents(ctx, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}
	}
}

// getAgent handles the query for a single agent
func getAgent(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	if len(path) != 1 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid agent id")
	}

	agentID := path[0]
	agent, found := k.GetAgent(ctx, agentID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "agent %s not found", agentID)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, agent)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// listAgents handles the query for all agents
func listAgents(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	agents := k.GetAllAgents(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, agents)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// QueryServerImpl defines the gRPC querier service
type QueryServerImpl struct {
	Keeper
}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return &QueryServerImpl{Keeper: k}
}

// Agent implements the Query/Agent gRPC method
func (q QueryServerImpl) Agent(c context.Context, req *types.QueryAgentRequest) (*types.QueryAgentResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	agent, found := q.GetAgent(ctx, req.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "agent %s not found", req.Id)
	}

	return &types.QueryAgentResponse{Agent: agent}, nil
}

// Agents implements the Query/Agents gRPC method
func (q QueryServerImpl) Agents(c context.Context, req *types.QueryAgentsRequest) (*types.QueryAgentsResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	agents := q.GetAllAgents(ctx)

	return &types.QueryAgentsResponse{Agents: agents}, nil
}
```

Repeat this process for each module.

### Develop CLI Commands

Create CLI commands in the respective `client/cli/` directories:

#### Example for TruthGPT Module:

Create files at `x/truthgpt/client/cli/tx.go` and `x/truthgpt/client/cli/query.go`:

**tx.go**:
```go
package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/nomercychain/nmxchain/x/truthgpt/types"
)

// GetTxCmd returns the transaction commands for the truthgpt module
func GetTxCmd() *cobra.Command {
	truthgptTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	truthgptTxCmd.AddCommand(
		NewCreateQueryCmd(),
		NewVerifyDataCmd(),
		NewSubmitResponseCmd(),
	)

	return truthgptTxCmd
}

// NewCreateQueryCmd implements the create-query command
func NewCreateQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-query [content] [reward]",
		Short: "Create a new oracle query",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			content := args[0]
			reward, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateQuery(
				clientCtx.GetFromAddress().String(),
				content,
				reward,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewVerifyDataCmd implements the verify-data command
func NewVerifyDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-data [query-id] [data]",
		Short: "Verify data for an oracle query",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryID := args[0]
			data := args[1]

			msg := types.NewMsgVerifyData(
				clientCtx.GetFromAddress().String(),
				queryID,
				data,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewSubmitResponseCmd implements the submit-response command
func NewSubmitResponseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-response [query-id] [response]",
		Short: "Submit a response to an oracle query",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryID := args[0]
			response := args[1]

			msg := types.NewMsgSubmitResponse(
				clientCtx.GetFromAddress().String(),
				queryID,
				response,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
```

**query.go**:
```go
package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/nomercychain/nmxchain/x/truthgpt/types"
)

// GetQueryCmd returns the query commands for the truthgpt module
func GetQueryCmd() *cobra.Command {
	truthgptQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	truthgptQueryCmd.AddCommand(
		NewQueryOracleQueryCmd(),
		NewQueryOracleQueriesCmd(),
		NewQueryDataSourceCmd(),
		NewQueryDataSourcesCmd(),
	)

	return truthgptQueryCmd
}

// NewQueryOracleQueryCmd implements the query oracle-query command
func NewQueryOracleQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle-query [id]",
		Short: "Query a specific oracle query by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.OracleQuery(
				cmd.Context(),
				&types.QueryOracleQueryRequest{Id: args[0]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// NewQueryOracleQueriesCmd implements the query oracle-queries command
func NewQueryOracleQueriesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle-queries",
		Short: "Query all oracle queries",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.OracleQueries(
				cmd.Context(),
				&types.QueryOracleQueriesRequest{},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// NewQueryDataSourceCmd implements the query data-source command
func NewQueryDataSourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data-source [id]",
		Short: "Query a specific data source by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.DataSource(
				cmd.Context(),
				&types.QueryDataSourceRequest{Id: args[0]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// NewQueryDataSourcesCmd implements the query data-sources command
func NewQueryDataSourcesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data-sources",
		Short: "Query all data sources",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.DataSources(
				cmd.Context(),
				&types.QueryDataSourcesRequest{},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
```

Repeat this process for each module.

## 4. Develop the Frontend

### Create Remaining Pages

Create the following pages in the `client/src/pages/` directory:

1. **Staking.js**: For staking tokens and managing validators
2. **Governance.js**: For participating in governance proposals
3. **SmartContracts.js**: For managing DynaContracts
4. **HyperChains.js**: For creating and managing HyperChains
5. **Oracle.js**: For interacting with the TruthGPT Oracle
6. **Settings.js**: For user settings

#### Example for HyperChains.js:

```jsx
import React, { useState, useContext } from 'react';
import { 
  Box, 
  Container, 
  Grid, 
  Paper, 
  Typography, 
  Button, 
  TextField, 
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Card,
  CardContent,
  CardActions,
  Chip,
  Divider,
  LinearProgress,
  Tab,
  Tabs
} from '@mui/material';
import { styled } from '@mui/material/styles';
import AddIcon from '@mui/icons-material/Add';
import LayersIcon from '@mui/icons-material/Layers';
import SettingsIcon from '@mui/icons-material/Settings';
import PlayArrowIcon from '@mui/icons-material/PlayArrow';
import { WalletContext } from '../context/WalletContext';

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  borderRadius: 12,
  background: 'linear-gradient(135deg, rgba(30,30,30,1) 0%, rgba(20,20,20,1) 100%)',
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
}));

const ChainCard = styled(Card)(({ theme }) => ({
  height: '100%',
  display: 'flex',
  flexDirection: 'column',
  borderRadius: 12,
  background: 'linear-gradient(135deg, rgba(30,30,30,1) 0%, rgba(20,20,20,1) 100%)',
  boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
  transition: 'transform 0.2s',
  '&:hover': {
    transform: 'translateY(-4px)',
  },
}));

const StyledTab = styled(Tab)(({ theme }) => ({
  textTransform: 'none',
  fontWeight: 600,
  fontSize: '1rem',
}));

const HyperChains = () => {
  const { account } = useContext(WalletContext);
  const [tabValue, setTabValue] = useState(0);
  const [createDialogOpen, setCreateDialogOpen] = useState(false);
  const [promptDialogOpen, setPromptDialogOpen] = useState(false);
  const [prompt, setPrompt] = useState('');

  const handleTabChange = (event, newValue) => {
    setTabValue(newValue);
  };

  // Mock chains data
  const myChains = [
    {
      id: 'chain1',
      name: 'GameFi Chain',
      description: 'Low-latency chain optimized for NFT gaming',
      type: 'appchain',
      status: 'active',
      createdAt: '2023-05-15',
      validators: 10,
      transactions: 1245678,
      tps: 1200,
    },
    {
      id: 'chain2',
      name: 'DeFi Privacy Chain',
      description: 'Privacy-focused chain for DeFi applications',
      type: 'zkevm',
      status: 'deploying',
      createdAt: '2023-06-01',
      validators: 5,
      transactions: 0,
      tps: 0,
    },
    {
      id: 'chain3',
      name: 'Data Oracle Chain',
      description: 'Chain for decentralized data verification',
      type: 'rollup',
      status: 'active',
      createdAt: '2023-04-22',
      validators: 15,
      transactions: 987654,
      tps: 850,
    },
  ];

  // Mock templates data
  const templates = [
    {
      id: 'template1',
      name: 'Gaming Chain',
      description: 'Optimized for gaming applications with low latency',
      type: 'appchain',
    },
    {
      id: 'template2',
      name: 'DeFi Chain',
      description: 'Optimized for financial applications with high throughput',
      type: 'zkevm',
    },
    {
      id: 'template3',
      name: 'Data Chain',
      description: 'Optimized for data storage and verification',
      type: 'rollup',
    },
    {
      id: 'template4',
      name: 'Custom Chain',
      description: 'Fully customizable chain with your choice of modules',
      type: 'custom',
    },
  ];

  const handleCreateFromPrompt = () => {
    // In a real implementation, this would call the API to create a chain from the prompt
    console.log('Creating chain from prompt:', prompt);
    setPromptDialogOpen(false);
    setPrompt('');
  };

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 4 }}>
        <Typography variant="h4">
          HyperChains
        </Typography>
        <Box>
          <Button 
            variant="contained" 
            startIcon={<AddIcon />} 
            sx={{ mr: 2 }}
            onClick={() => setPromptDialogOpen(true)}
          >
            Create from Prompt
          </Button>
          <Button 
            variant="contained" 
            startIcon={<AddIcon />}
            onClick={() => setCreateDialogOpen(true)}
          >
            Create Chain
          </Button>
        </Box>
      </Box>

      <StyledPaper sx={{ mb: 4 }}>
        <Typography variant="h6" gutterBottom>
          What are HyperChains?
        </Typography>
        <Typography variant="body1" paragraph>
          HyperChains are AI-generated Layer 3 chains that can be created from natural language prompts. 
          They provide customized blockchain environments optimized for specific use cases.
        </Typography>
        <Typography variant="body1">
          You can create a HyperChain by describing what you want in natural language, or by selecting from pre-defined templates.
          Each HyperChain comes with its own set of validators, governance, and customized modules.
        </Typography>
      </StyledPaper>

      <Box sx={{ mb: 4 }}>
        <Tabs 
          value={tabValue} 
          onChange={handleTabChange} 
          sx={{ 
            mb: 3,
            '& .MuiTabs-indicator': {
              backgroundColor: 'primary.main',
            }
          }}
        >
          <StyledTab label="My Chains" />
          <StyledTab label="Templates" />
          <StyledTab label="Explorer" />
        </Tabs>

        {tabValue === 0 && (
          <>
            {account ? (
              <Grid container spacing={4}>
                {myChains.map((chain) => (
                  <Grid item xs={12} md={4} key={chain.id}>
                    <ChainCard>
                      <CardContent sx={{ flexGrow: 1 }}>
                        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', mb: 2 }}>
                          <Typography variant="h6">
                            {chain.name}
                          </Typography>
                          <Chip 
                            label={chain.status} 
                            size="small" 
                            sx={{ 
                              bgcolor: chain.status === 'active' ? 'success.main' : 
                                      chain.status === 'deploying' ? 'warning.main' : 'error.main',
                              color: 'white'
                            }} 
                          />
                        </Box>
                        <Chip 
                          label={chain.type} 
                          size="small" 
                          sx={{ 
                            mb: 2,
                            bgcolor: chain.type === 'zkevm' ? 'primary.main' : 
                                    chain.type === 'rollup' ? 'secondary.main' : 
                                    chain.type === 'appchain' ? 'warning.main' : 'info.main',
                            color: 'white'
                          }} 
                        />
                        <Typography variant="body2" color="text.secondary" paragraph>
                          {chain.description}
                        </Typography>
                        <Divider sx={{ my: 2 }} />
                        <Grid container spacing={2}>
                          <Grid item xs={4}>
                            <Typography variant="caption" color="text.secondary" display="block">
                              Validators
                            </Typography>
                            <Typography variant="body2">
                              {chain.validators}
                            </Typography>
                          </Grid>
                          <Grid item xs={4}>
                            <Typography variant="caption" color="text.secondary" display="block">
                              Transactions
                            </Typography>
                            <Typography variant="body2">
                              {chain.transactions.toLocaleString()}
                            </Typography>
                          </Grid>
                          <Grid item xs={4}>
                            <Typography variant="caption" color="text.secondary" display="block">
                              TPS
                            </Typography>
                            <Typography variant="body2">
                              {chain.tps}
                            </Typography>
                          </Grid>
                        </Grid>
                        <Box sx={{ mt: 2 }}>
                          <Typography variant="caption" color="text.secondary" display="block" gutterBottom>
                            Status: {chain.status}
                          </Typography>
                          <LinearProgress 
                            variant="determinate" 
                            value={chain.status === 'active' ? 100 : chain.status === 'deploying' ? 50 : 0} 
                            sx={{ 
                              height: 6, 
                              borderRadius: 3,
                              bgcolor: 'rgba(255,255,255,0.1)',
                              '& .MuiLinearProgress-bar': {
                                bgcolor: chain.status === 'active' ? 'success.main' : 
                                        chain.status === 'deploying' ? 'warning.main' : 'error.main',
                              }
                            }} 
                          />
                        </Box>
                      </CardContent>
                      <CardActions sx={{ justifyContent: 'space-between', p: 2 }}>
                        <Button size="small" startIcon={<SettingsIcon />}>
                          Manage
                        </Button>
                        <Button size="small" variant="outlined">
                          Explorer
                        </Button>
                      </CardActions>
                    </ChainCard>
                  </Grid>
                ))}
              </Grid>
            ) : (
              <StyledPaper sx={{ p: 4, textAlign: 'center' }}>
                <Typography variant="h6" gutterBottom>
                  Connect Your Wallet
                </Typography>
                <Typography variant="body1" color="text.secondary" paragraph>
                  Please connect your wallet to view and manage your HyperChains.
                </Typography>
                <Button 
                  variant="contained" 
                  size="large" 
                  onClick={() => {}}
                  sx={{ 
                    mt: 2, 
                    background: 'linear-gradient(45deg, #6200ea 30%, #00e676 90%)',
                    '&:hover': {
                      background: 'linear-gradient(45deg, #5000d6 30%, #00c060 90%)',
                    },
                  }}
                >
                  Connect Wallet
                </Button>
              </StyledPaper>
            )}
          </>
        )}

        {tabValue === 1 && (
          <Grid container spacing={4}>
            {templates.map((template) => (
              <Grid item xs={12} sm={6} md={3} key={template.id}>
                <ChainCard>
                  <CardContent>
                    <Typography variant="h6" gutterBottom>
                      {template.name}
                    </Typography>
                    <Chip 
                      label={template.type} 
                      size="small" 
                      sx={{ 
                        mb: 2,
                        bgcolor: template.type === 'zkevm' ? 'primary.main' : 
                                template.type === 'rollup' ? 'secondary.main' : 
                                template.type === 'appchain' ? 'warning.main' : 'info.main',
                        color: 'white'
                      }} 
                    />
                    <Typography variant="body2" color="text.secondary">
                      {template.description}
                    </Typography>
                  </CardContent>
                  <Box sx={{ flexGrow: 1 }} />
                  <CardActions>
                    <Button size="small" fullWidth variant="contained">
                      Use Template
                    </Button>
                  </CardActions>
                </ChainCard>
              </Grid>
            ))}
          </Grid>
        )}

        {tabValue === 2 && (
          <StyledPaper>
            <Typography variant="h6" gutterBottom>
              HyperChains Explorer
            </Typography>
            <Divider sx={{ my: 2 }} />
            <Typography variant="body1" paragraph>
              Explore all active HyperChains in the NoMercyChain ecosystem.
            </Typography>
            <Button variant="contained" startIcon={<LayersIcon />}>
              Browse Chains
            </Button>
          </StyledPaper>
        )}
      </Box>

      {/* Create Chain Dialog */}
      <Dialog open={createDialogOpen} onClose={() => setCreateDialogOpen(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Create New HyperChain</DialogTitle>
        <DialogContent>
          <TextField
            label="Chain Name"
            fullWidth
            margin="normal"
            variant="outlined"
            placeholder="Enter a name for your chain"
          />
          <TextField
            label="Description"
            fullWidth
            margin="normal"
            variant="outlined"
            placeholder="Describe your chain's purpose"
            multiline
            rows={3}
          />
          <TextField
            label="Chain Type"
            fullWidth
            margin="normal"
            variant="outlined"
            select
            SelectProps={{
              native: true,
            }}
          >
            <option value="">Select a chain type</option>
            <option value="zkevm">zkEVM</option>
            <option value="optimistic">Optimistic Rollup</option>
            <option value="rollup">Standard Rollup</option>
            <option value="appchain">Application Chain</option>
            <option value="custom">Custom</option>
          </TextField>
          <TextField
            label="Template"
            fullWidth
            margin="normal"
            variant="outlined"
            select
            SelectProps={{
              native: true,
            }}
          >
            <option value="">Select a template</option>
            <option value="template1">Gaming Chain</option>
            <option value="template2">DeFi Chain</option>
            <option value="template3">Data Chain</option>
            <option value="template4">Custom Chain</option>
          </TextField>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setCreateDialogOpen(false)}>Cancel</Button>
          <Button variant="contained">Create Chain</Button>
        </DialogActions>
      </Dialog>

      {/* Create from Prompt Dialog */}
      <Dialog open={promptDialogOpen} onClose={() => setPromptDialogOpen(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Create Chain from Prompt</DialogTitle>
        <DialogContent>
          <Typography variant="body2" color="text.secondary" paragraph sx={{ mt: 2 }}>
            Describe the chain you want to create in natural language. Our AI will generate a chain based on your description.
          </Typography>
          <TextField
            label="Prompt"
            fullWidth
            margin="normal"
            variant="outlined"
            placeholder="E.g., Create a low-latency NFT game chain with high throughput and minimal fees"
            multiline
            rows={4}
            value={prompt}
            onChange={(e) => setPrompt(e.target.value)}
          />
          <Typography variant="caption" color="text.secondary">
            Examples: "Create a privacy-focused DeFi chain", "Build a chain for decentralized social media"
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setPromptDialogOpen(false)}>Cancel</Button>
          <Button 
            variant="contained" 
            onClick={handleCreateFromPrompt}
            disabled={!prompt.trim()}
          >
            Generate Chain
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
};

export default HyperChains;
```

### Implement State Management

Enhance the application's state management by creating additional context providers:

#### Create a ChainContext:

```jsx
// client/src/context/ChainContext.js
import { createContext, useState, useEffect } from 'react';

export const ChainContext = createContext({
  chains: [],
  templates: [],
  loading: false,
  error: null,
  fetchChains: () => {},
  fetchTemplates: () => {},
  createChain: () => {},
  createChainFromPrompt: () => {},
});

export const ChainProvider = ({ children }) => {
  const [chains, setChains] = useState([]);
  const [templates, setTemplates] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchChains = async () => {
    setLoading(true);
    try {
      // In a real implementation, this would call the API
      // const response = await api.get('/chains');
      // setChains(response.data);
      
      // Mock data for now
      setChains([
        {
          id: 'chain1',
          name: 'GameFi Chain',
          description: 'Low-latency chain optimized for NFT gaming',
          type: 'appchain',
          status: 'active',
          createdAt: '2023-05-15',
          validators: 10,
          transactions: 1245678,
          tps: 1200,
        },
        // ... other chains
      ]);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const fetchTemplates = async () => {
    setLoading(true);
    try {
      // In a real implementation, this would call the API
      // const response = await api.get('/templates');
      // setTemplates(response.data);
      
      // Mock data for now
      setTemplates([
        {
          id: 'template1',
          name: 'Gaming Chain',
          description: 'Optimized for gaming applications with low latency',
          type: 'appchain',
        },
        // ... other templates
      ]);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const createChain = async (chainData) => {
    setLoading(true);
    try {
      // In a real implementation, this would call the API
      // const response = await api.post('/chains', chainData);
      // return response.data;
      
      // Mock response for now
      return {
        id: 'new-chain-' + Date.now(),
        ...chainData,
        status: 'deploying',
        createdAt: new Date().toISOString(),
      };
    } catch (err) {
      setError(err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  const createChainFromPrompt = async (prompt) => {
    setLoading(true);
    try {
      // In a real implementation, this would call the API
      // const response = await api.post('/chains/prompt', { prompt });
      // return response.data;
      
      // Mock response for now
      return {
        id: 'prompt-chain-' + Date.now(),
        name: `Chain from prompt: ${prompt.substring(0, 20)}...`,
        description: prompt,
        type: prompt.includes('privacy') ? 'zkevm' : 
              prompt.includes('game') ? 'appchain' : 'rollup',
        status: 'deploying',
        createdAt: new Date().toISOString(),
      };
    } catch (err) {
      setError(err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchChains();
    fetchTemplates();
  }, []);

  return (
    <ChainContext.Provider
      value={{
        chains,
        templates,
        loading,
        error,
        fetchChains,
        fetchTemplates,
        createChain,
        createChainFromPrompt,
      }}
    >
      {children}
    </ChainContext.Provider>
  );
};
```

### Add API Integration

Create an API service to connect with the blockchain:

```javascript
// client/src/services/api.js
import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:1317';

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add request interceptor to include auth token
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('auth_token');
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Chain API
export const chainApi = {
  getChains: () => api.get('/hyperchains/chains'),
  getChain: (id) => api.get(`/hyperchains/chains/${id}`),
  createChain: (data) => api.post('/hyperchains/chains', data),
  createChainFromPrompt: (prompt) => api.post('/hyperchains/chains/prompt', { prompt }),
  getTemplates: () => api.get('/hyperchains/templates'),
  getTemplate: (id) => api.get(`/hyperchains/templates/${id}`),
};

// Agent API
export const agentApi = {
  getAgents: () => api.get('/deai/agents'),
  getAgent: (id) => api.get(`/deai/agents/${id}`),
  createAgent: (data) => api.post('/deai/agents', data),
  updateAgent: (id, data) => api.put(`/deai/agents/${id}`, data),
  deleteAgent: (id) => api.delete(`/deai/agents/${id}`),
};

// Oracle API
export const oracleApi = {
  getQueries: () => api.get('/truthgpt/queries'),
  getQuery: (id) => api.get(`/truthgpt/queries/${id}`),
  createQuery: (data) => api.post('/truthgpt/queries', data),
  verifyData: (queryId, data) => api.post(`/truthgpt/queries/${queryId}/verify`, { data }),
  submitResponse: (queryId, response) => api.post(`/truthgpt/queries/${queryId}/respond`, { response }),
};

// Smart Contract API
export const contractApi = {
  getContracts: () => api.get('/dynacontracts/contracts'),
  getContract: (id) => api.get(`/dynacontracts/contracts/${id}`),
  createContract: (data) => api.post('/dynacontracts/contracts', data),
  updateContract: (id, data) => api.put(`/dynacontracts/contracts/${id}`, data),
  executeContract: (id, data) => api.post(`/dynacontracts/contracts/${id}/execute`, data),
};

export default api;
```

## 5. Test Your Implementation

### Write Unit Tests for Modules

Create unit tests for each module in their respective `tests` directories:

#### Example for HyperChains Module:

```go
// x/hyperchains/keeper/keeper_test.go
package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/app"
	"github.com/nomercychain/nmxchain/x/hyperchains/keeper"
	"github.com/nomercychain/nmxchain/x/hyperchains/types"
)

func setupKeeper(t testing.TB) (*app.App, keeper.Keeper, sdk.Context) {
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, Time: time.Now()})
	return app, app.HyperchainsKeeper, ctx
}

func TestCreateChain(t *testing.T) {
	_, k, ctx := setupKeeper(t)
	
	creator := sdk.AccAddress([]byte("creator"))
	name := "Test Chain"
	description := "Test chain description"
	chainType := types.ChainTypeRollup
	
	// Create a chain
	id, err := k.CreateChainFromTemplate(ctx, creator, "template1", name, description, []byte(`{}`))
	require.NoError(t, err)
	require.NotEmpty(t, id)
	
	// Verify the chain was created
	chain, found := k.GetChain(ctx, id)
	require.True(t, found)
	require.Equal(t, name, chain.Name)
	require.Equal(t, description, chain.Description)
	require.Equal(t, creator, chain.Creator)
	require.Equal(t, chainType, chain.ChainType)
	require.Equal(t, types.ChainStatusProposed, chain.Status)
}

func TestCreateChainFromPrompt(t *testing.T) {
	_, k, ctx := setupKeeper(t)
	
	creator := sdk.AccAddress([]byte("creator"))
	prompt := "Create a privacy-focused DeFi chain"
	
	// Create a chain from prompt
	id, err := k.CreateChainFromPrompt(ctx, creator, prompt)
	require.NoError(t, err)
	require.NotEmpty(t, id)
	
	// Verify the chain was created
	chain, found := k.GetChain(ctx, id)
	require.True(t, found)
	require.Contains(t, chain.Name, prompt)
	require.Equal(t, prompt, chain.Description)
	require.Equal(t, creator, chain.Creator)
	require.Equal(t, types.ChainTypeZkEVM, chain.ChainType)
	require.Equal(t, types.ChainStatusProposed, chain.Status)
}

func TestDeployChain(t *testing.T) {
	_, k, ctx := setupKeeper(t)
	
	creator := sdk.AccAddress([]byte("creator"))
	name := "Test Chain"
	description := "Test chain description"
	
	// Create a chain
	id, err := k.CreateChainFromTemplate(ctx, creator, "template1", name, description, []byte(`{}`))
	require.NoError(t, err)
	
	// Deploy the chain
	err = k.DeployChain(ctx, id, creator)
	require.NoError(t, err)
	
	// Verify the chain status was updated
	chain, found := k.GetChain(ctx, id)
	require.True(t, found)
	require.Equal(t, types.ChainStatusDeploying, chain.Status)
	
	// Verify a deployment was created
	deployments := k.GetChainDeploymentsByChain(ctx, id)
	require.Len(t, deployments, 1)
	require.Equal(t, id, deployments[0].ChainID)
	require.Equal(t, creator, deployments[0].Deployer)
	require.Equal(t, "in_progress", deployments[0].Status)
}
```

### Create Integration Tests

Create integration tests to verify module interactions:

```go
// tests/integration/hyperchains_test.go
package integration_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/app"
	"github.com/nomercychain/nmxchain/x/hyperchains/types"
)

func TestHyperChainsWorkflow(t *testing.T) {
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, Time: time.Now()})
	
	// Create a user
	user := sdk.AccAddress([]byte("user"))
	
	// 1. Create a chain from prompt
	prompt := "Create a low-latency NFT game chain"
	chainID, err := app.HyperchainsKeeper.CreateChainFromPrompt(ctx, user, prompt)
	require.NoError(t, err)
	require.NotEmpty(t, chainID)
	
	// 2. Verify the chain was created with the correct type
	chain, found := app.HyperchainsKeeper.GetChain(ctx, chainID)
	require.True(t, found)
	require.Equal(t, types.ChainTypeAppChain, chain.ChainType)
	
	// 3. Deploy the chain
	err = app.HyperchainsKeeper.DeployChain(ctx, chainID, user)
	require.NoError(t, err)
	
	// 4. Process a block to simulate deployment progress
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(ctx.BlockTime().Add(time.Minute))
	app.HyperchainsKeeper.ProcessPendingDeployments(ctx)
	
	// 5. Verify the deployment status was updated
	deployments := app.HyperchainsKeeper.GetChainDeploymentsByChain(ctx, chainID)
	require.Len(t, deployments, 1)
	require.Equal(t, "in_progress", deployments[0].Status)
	
	// 6. Process another block to complete the deployment
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(ctx.BlockTime().Add(time.Minute))
	app.HyperchainsKeeper.ProcessPendingDeployments(ctx)
	
	// 7. Verify the deployment was completed
	deployments = app.HyperchainsKeeper.GetChainDeploymentsByChain(ctx, chainID)
	require.Len(t, deployments, 1)
	require.Equal(t, "completed", deployments[0].Status)
	
	// 8. Verify the chain status was updated
	chain, found = app.HyperchainsKeeper.GetChain(ctx, chainID)
	require.True(t, found)
	require.Equal(t, types.ChainStatusActive, chain.Status)
	
	// 9. Register a validator for the chain
	validatorAddr := sdk.ValAddress([]byte("validator"))
	operatorAddr := sdk.AccAddress([]byte("operator"))
	err = app.HyperchainsKeeper.RegisterChainValidator(ctx, chainID, validatorAddr, operatorAddr)
	require.NoError(t, err)
	
	// 10. Verify the validator was registered
	validators := app.HyperchainsKeeper.GetChainValidatorsByChain(ctx, chainID)
	require.Len(t, validators, 1)
	require.Equal(t, validatorAddr, validators[0].ValidatorAddress)
	require.Equal(t, operatorAddr, validators[0].OperatorAddress)
}
```

### Test Frontend-Backend Integration

Create end-to-end tests for the frontend-backend integration:

```javascript
// client/src/tests/integration/hyperchains.test.js
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import { ChainProvider } from '../../context/ChainContext';
import { WalletContext } from '../../context/WalletContext';
import HyperChains from '../../pages/HyperChains';
import { chainApi } from '../../services/api';

// Mock the API
jest.mock('../../services/api', () => ({
  chainApi: {
    getChains: jest.fn(),
    getTemplates: jest.fn(),
    createChainFromPrompt: jest.fn(),
  },
}));

const mockWalletContext = {
  account: '0x1234567890abcdef1234567890abcdef12345678',
  balance: '1000',
  connectWallet: jest.fn(),
  disconnectWallet: jest.fn(),
  isConnecting: false,
};

describe('HyperChains Integration Tests', () => {
  beforeEach(() => {
    // Reset mocks
    jest.clearAllMocks();
    
    // Mock API responses
    chainApi.getChains.mockResolvedValue({
      data: [
        {
          id: 'chain1',
          name: 'GameFi Chain',
          description: 'Low-latency chain optimized for NFT gaming',
          type: 'appchain',
          status: 'active',
          createdAt: '2023-05-15',
          validators: 10,
          transactions: 1245678,
          tps: 1200,
        },
      ],
    });
    
    chainApi.getTemplates.mockResolvedValue({
      data: [
        {
          id: 'template1',
          name: 'Gaming Chain',
          description: 'Optimized for gaming applications with low latency',
          type: 'appchain',
        },
      ],
    });
    
    chainApi.createChainFromPrompt.mockResolvedValue({
      data: {
        id: 'new-chain-1',
        name: 'Chain from prompt: Create a gaming chain',
        description: 'Create a gaming chain',
        type: 'appchain',
        status: 'deploying',
        createdAt: '2023-06-15T12:00:00Z',
      },
    });
  });
  
  test('renders HyperChains page and loads data', async () => {
    render(
      <BrowserRouter>
        <WalletContext.Provider value={mockWalletContext}>
          <ChainProvider>
            <HyperChains />
          </ChainProvider>
        </WalletContext.Provider>
      </BrowserRouter>
    );
    
    // Check that the page title is rendered
    expect(screen.getByText('HyperChains')).toBeInTheDocument();
    
    // Wait for chains to load
    await waitFor(() => {
      expect(chainApi.getChains).toHaveBeenCalled();
      expect(screen.getByText('GameFi Chain')).toBeInTheDocument();
    });
    
    // Switch to Templates tab
    fireEvent.click(screen.getByText('Templates'));
    
    // Wait for templates to load
    await waitFor(() => {
      expect(chainApi.getTemplates).toHaveBeenCalled();
      expect(screen.getByText('Gaming Chain')).toBeInTheDocument();
    });
  });
  
  test('creates a chain from prompt', async () => {
    render(
      <BrowserRouter>
        <WalletContext.Provider value={mockWalletContext}>
          <ChainProvider>
            <HyperChains />
          </ChainProvider>
        </WalletContext.Provider>
      </BrowserRouter>
    );
    
    // Click the "Create from Prompt" button
    fireEvent.click(screen.getByText('Create from Prompt'));
    
    // Enter a prompt
    fireEvent.change(screen.getByPlaceholderText(/E.g., Create a low-latency NFT game chain/i), {
      target: { value: 'Create a gaming chain' },
    });
    
    // Click the "Generate Chain" button
    fireEvent.click(screen.getByText('Generate Chain'));
    
    // Wait for the API call
    await waitFor(() => {
      expect(chainApi.createChainFromPrompt).toHaveBeenCalledWith('Create a gaming chain');
    });
  });
});
```

## 6. Deploy a Testnet

### Set Up Validator Nodes

1. **Prepare the validator machine**:
   ```bash
   # Update system packages
   sudo apt update && sudo apt upgrade -y
   
   # Install required packages
   sudo apt install -y build-essential git curl jq
   
   # Install Go
   wget https://go.dev/dl/go1.18.10.linux-amd64.tar.gz
   sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.18.10.linux-amd64.tar.gz
   export PATH=$PATH:/usr/local/go/bin
   echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
   source ~/.profile
   ```

2. **Build and install the NoMercyChain binary**:
   ```bash
   git clone https://github.com/nomercychain/nmxchain.git
   cd nmxchain
   make install
   ```

3. **Initialize the validator node**:
   ```bash
   nmxchaind init <validator-name> --chain-id nomercychain-testnet-1
   ```

4. **Create a validator key**:
   ```bash
   nmxchaind keys add validator --keyring-backend test
   ```

5. **Add genesis account**:
   ```bash
   nmxchaind add-genesis-account $(nmxchaind keys show validator -a --keyring-backend test) 10000000000unmx
   ```

6. **Create validator transaction**:
   ```bash
   nmxchaind gentx validator 5000000000unmx \
     --chain-id nomercychain-testnet-1 \
     --moniker="<validator-name>" \
     --commission-rate="0.10" \
     --commission-max-rate="0.20" \
     --commission-max-change-rate="0.01" \
     --min-self-delegation="1" \
     --keyring-backend test
   ```

### Configure Genesis File

1. **Collect gentx files from all validators**:
   ```bash
   # On the coordinator node
   mkdir -p ~/gentx
   # Copy all gentx files to this directory
   ```

2. **Add gentx files to genesis**:
   ```bash
   nmxchaind collect-gentxs --gentx-dir ~/gentx
   ```

3. **Validate genesis file**:
   ```bash
   nmxchaind validate-genesis
   ```

4. **Customize genesis parameters**:
   ```bash
   # Edit the genesis file
   nano ~/.nmxchain/config/genesis.json
   
   # Set custom parameters for modules
   # For example, adjust staking parameters:
   # "unbonding_time": "1814400s",  # 21 days
   # "max_validators": 100,
   # "max_entries": 7,
   # "historical_entries": 10000,
   ```

### Launch the Testnet

#### For Windows:

1. **Configure P2P settings**:
   ```powershell
   # Edit config.toml using Notepad or any text editor
   notepad $env:USERPROFILE\.nmxchain\config\config.toml
   
   # Set persistent peers
   # persistent_peers = "validator1_node_id@validator1_ip:26656,validator2_node_id@validator2_ip:26656"
   ```

#### For Linux/macOS:

1. **Configure P2P settings**:
   ```bash
   # Edit config.toml
   nano ~/.nmxchain/config/config.toml
   
   # Set persistent peers
   # persistent_peers = "validator1_node_id@validator1_ip:26656,validator2_node_id@validator2_ip:26656"
   ```

2. **Configure API and RPC settings**:
   ```bash
   # Edit app.toml
   nano ~/.nmxchain/config/app.toml
   
   # Enable the API server
   [api]
   enable = true
   swagger = true
   address = "tcp://0.0.0.0:1317"
   
   # Enable the gRPC server
   [grpc]
   enable = true
   address = "0.0.0.0:9090"
   ```

3. **Start the validator node**:
   ```bash
   nmxchaind start
   ```

4. **Set up a systemd service** (for production):
   ```bash
   sudo nano /etc/systemd/system/nmxchaind.service
   ```
   
   Add the following content:
   ```
   [Unit]
   Description=NoMercyChain Node
   After=network-online.target
   
   [Service]
   User=<your-user>
   ExecStart=/home/<your-user>/go/bin/nmxchaind start
   Restart=always
   RestartSec=3
   LimitNOFILE=4096
   
   [Install]
   WantedBy=multi-user.target
   ```
   
   Enable and start the service:
   ```bash
   sudo systemctl enable nmxchaind
   sudo systemctl start nmxchaind
   ```

5. **Monitor the node**:
   ```bash
   # Check logs
   sudo journalctl -u nmxchaind -f
   
   # Check node status
   nmxchaind status
   ```

## Conclusion

This development guide provides a comprehensive roadmap for setting up your environment, implementing modules, developing the frontend, testing, and deploying a testnet for the NoMercyChain project.

By following these steps, you'll be able to build a fully functional blockchain platform with AI-powered smart contracts, decentralized AI agents, and Layer 3 solutions.

Remember to refer to the `NEXT_STEPS.md` file for a prioritized list of tasks and the project timeline.

Happy coding!