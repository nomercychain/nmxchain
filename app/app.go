package app

import (
	"io"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cometbft/cometbft/abci/types"
	abci "github.com/cometbft/cometbft/abci/types"
	tmjson "github.com/cometbft/cometbft/libs/json"
	"github.com/cometbft/cometbft/libs/log"
	tmos "github.com/cometbft/cometbft/libs/os"
	dbm "github.com/cometbft/cometbft-db"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	// IBC
	"github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v7/modules/core/02-client"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	// Custom modules
	"github.com/nomercychain/nmxchain/x/deai"
	deaikeeper "github.com/nomercychain/nmxchain/x/deai/keeper"
	deaitypes "github.com/nomercychain/nmxchain/x/deai/types"

	"github.com/nomercychain/nmxchain/x/dynacontracts"
	dynacontractskeeper "github.com/nomercychain/nmxchain/x/dynacontracts/keeper"
	dynacontractstypes "github.com/nomercychain/nmxchain/x/dynacontracts/types"

	"github.com/nomercychain/nmxchain/x/hyperchain"
	hyperchainkeeper "github.com/nomercychain/nmxchain/x/hyperchain/keeper"
	hyperchainstypes "github.com/nomercychain/nmxchain/x/hyperchain/types"

	"github.com/nomercychain/nmxchain/x/hyperchains"
	hyperchainskeeper "github.com/nomercychain/nmxchain/x/hyperchains/keeper"
	hyperchaintypes "github.com/nomercychain/nmxchain/x/hyperchains/types"

	"github.com/nomercychain/nmxchain/x/neuropos"
	neuroposkeeper "github.com/nomercychain/nmxchain/x/neuropos/keeper"
	neuropostypes "github.com/nomercychain/nmxchain/x/neuropos/types"

	"github.com/nomercychain/nmxchain/x/truthgpt"
	truthgptkeeper "github.com/nomercychain/nmxchain/x/truthgpt/keeper"
	truthgpttypes "github.com/nomercychain/nmxchain/x/truthgpt/types"
	// Custom modules
	"github.com/nomercychain/nmxchain/x/neuropos"
	neuroposkeeper "github.com/nomercychain/nmxchain/x/neuropos/keeper"
	neuropostypes "github.com/nomercychain/nmxchain/x/neuropos/types"
	"github.com/nomercychain/nmxchain/x/truthgpt"
	truthgptkeeper "github.com/nomercychain/nmxchain/x/truthgpt/keeper"
	truthgpttypes "github.com/nomercychain/nmxchain/x/truthgpt/types"
	"github.com/nomercychain/nmxchain/x/deai"
	deaikeeper "github.com/nomercychain/nmxchain/x/deai/keeper"
	deaitypes "github.com/nomercychain/nmxchain/x/deai/types"
	"github.com/nomercychain/nmxchain/x/dynacontract"
	dynacontractkeeper "github.com/nomercychain/nmxchain/x/dynacontract/keeper"
	dynacontracttypes "github.com/nomercychain/nmxchain/x/dynacontract/types"
	"github.com/nomercychain/nmxchain/x/hyperchain"
	hyperchainkeeper "github.com/nomercychain/nmxchain/x/hyperchain/keeper"
	hyperchaintypes "github.com/nomercychain/nmxchain/x/hyperchain/types"
)

const (
	AppName = "nmxchain"
)

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distrclient.ProposalHandler, upgradeclient.LegacyProposalHandler, upgradeclient.LegacyCancelProposalHandler,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		// Custom modules
		neuropos.AppModuleBasic{},
		truthgpt.AppModuleBasic{},
		deai.AppModuleBasic{},
		dynacontract.AppModuleBasic{},
		hyperchain.AppModuleBasic{},
	)
)

// NMXApp extends an ABCI application
type NMXApp struct {
	*baseapp.BaseApp
	
	// Keepers
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	EvidenceKeeper   evidencekeeper.Keeper
	
	// Custom keepers
	NeuroPoSKeeper    neuroposkeeper.Keeper
	TruthGPTKeeper    truthgptkeeper.Keeper
	DeAIKeeper        deaikeeper.Keeper
	DynaContractKeeper dynacontractkeeper.Keeper
	HyperChainKeeper  hyperchainkeeper.Keeper
	
	// Codec
	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry
	
	// Keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlockedModuleAccountAddrs returns all the blocked module account addresses.
func (app *NMXApp) BlockedModuleAccountAddrs() map[string]bool {
	modAccAddrs := app.ModuleAccountAddrs()
	delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	// Allow specific module accounts to receive funds
	for acc := range allowedReceivingModAcc {
		delete(modAccAddrs, authtypes.NewModuleAddress(acc).String())
	}

	return modAccAddrs
	// Modules
	mm *module.Manager
	sm *module.SimulationManager
}

// NewNMXApp returns a reference to an initialized NMXApp
func NewNMXApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *NMXApp {
	appCodec := encodingConfig.Codec
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(AppName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		distrtypes.StoreKey, slashingtypes.StoreKey, govtypes.StoreKey,
		paramstypes.StoreKey, upgradetypes.StoreKey, evidencetypes.StoreKey,
		capabilitytypes.StoreKey, neuropostypes.StoreKey, truthgpttypes.StoreKey,
		deaitypes.StoreKey, dynacontracttypes.StoreKey, hyperchaintypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &NMXApp{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	// Initialize params keeper and subspaces
	app.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// Set the subspaces used by the various modules
	authSubspace := app.ParamsKeeper.Subspace(authtypes.ModuleName)
	bankSubspace := app.ParamsKeeper.Subspace(banktypes.ModuleName)
	stakingSubspace := app.ParamsKeeper.Subspace(stakingtypes.ModuleName)
	distrSubspace := app.ParamsKeeper.Subspace(distrtypes.ModuleName)
	slashingSubspace := app.ParamsKeeper.Subspace(slashingtypes.ModuleName)
	govSubspace := app.ParamsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	crisisSubspace := app.ParamsKeeper.Subspace(crisistypes.ModuleName)
	upgradeSubspace := app.ParamsKeeper.Subspace(upgradetypes.ModuleName)
	evidenceSubspace := app.ParamsKeeper.Subspace(evidencetypes.ModuleName)
	
	// Custom module subspaces
	neuroposSubspace := app.ParamsKeeper.Subspace(neuropostypes.ModuleName)
	truthgptSubspace := app.ParamsKeeper.Subspace(truthgpttypes.ModuleName)
	deaiSubspace := app.ParamsKeeper.Subspace(deaitypes.ModuleName)
	dynacontractSubspace := app.ParamsKeeper.Subspace(dynacontracttypes.ModuleName)
	hyperchainSubspace := app.ParamsKeeper.Subspace(hyperchaintypes.ModuleName)

	// Add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], authSubspace, authtypes.ProtoBaseAccount, GetMaccPerms(),
	)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], app.AccountKeeper, bankSubspace, app.ModuleAccountAddrs(),
	)
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec, keys[stakingtypes.StoreKey], app.AccountKeeper, app.BankKeeper, stakingSubspace,
	)
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec, keys[distrtypes.StoreKey], distrSubspace, app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, authtypes.FeeCollectorName,
	)
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec, keys[slashingtypes.StoreKey], &stakingKeeper, slashingSubspace,
	)
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		crisisSubspace, invCheckPeriod, app.BankKeeper, authtypes.FeeCollectorName,
	)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, app.BaseApp)

	// Create evidence keeper with router
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], &app.StakingKeeper, app.SlashingKeeper,
	)
	app.EvidenceKeeper = *evidenceKeeper

	// Create capability keeper and ScopedKeeper
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
	scopedIBCKeeper := app.CapabilityKeeper.ScopedKeeper(authtypes.ModuleName)
	
	// Create static IBC router, add app routes, then set and seal it
	// (This would be implemented if IBC was needed)

	// Create governance keeper
	app.GovKeeper = govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], govSubspace, app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, govtypes.DefaultConfig(),
	)

	// Set staking keeper
	app.StakingKeeper = stakingKeeper

	// Initialize custom module keepers
	app.NeuroPoSKeeper = neuroposkeeper.NewKeeper(
		appCodec,
		keys[neuropostypes.StoreKey],
		neuroposSubspace,
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
	)

	app.TruthGPTKeeper = truthgptkeeper.NewKeeper(
		appCodec,
		keys[truthgpttypes.StoreKey],
		truthgptSubspace,
		app.AccountKeeper,
		app.BankKeeper,
	)

	app.DeAIKeeper = deaikeeper.NewKeeper(
		appCodec,
		keys[deaitypes.StoreKey],
		deaiSubspace,
		app.AccountKeeper,
		app.BankKeeper,
	)

	app.DynaContractKeeper = dynacontractkeeper.NewKeeper(
		appCodec,
		keys[dynacontracttypes.StoreKey],
		dynacontractSubspace,
		app.AccountKeeper,
		app.BankKeeper,
	)

	app.HyperChainKeeper = hyperchainkeeper.NewKeeper(
		appCodec,
		keys[hyperchaintypes.StoreKey],
		hyperchainSubspace,
		app.AccountKeeper,
		app.BankKeeper,
	)

	// Create the module manager
	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		// SDK app modules
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		params.NewAppModule(app.ParamsKeeper),
		
		// Custom app modules
		neuropos.NewAppModule(appCodec, app.NeuroPoSKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		truthgpt.NewAppModule(appCodec, app.TruthGPTKeeper, app.AccountKeeper, app.BankKeeper),
		deai.NewAppModule(appCodec, app.DeAIKeeper, app.AccountKeeper, app.BankKeeper),
		dynacontract.NewAppModule(appCodec, app.DynaContractKeeper, app.AccountKeeper, app.BankKeeper),
		hyperchain.NewAppModule(appCodec, app.HyperChainKeeper, app.AccountKeeper, app.BankKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName, capabilitytypes.ModuleName, crisistypes.ModuleName,
		govtypes.ModuleName, stakingtypes.ModuleName, slashingtypes.ModuleName,
		distrtypes.ModuleName, evidencetypes.ModuleName, authtypes.ModuleName,
		banktypes.ModuleName, genutiltypes.ModuleName, paramstypes.ModuleName,
		// Custom modules
		neuropostypes.ModuleName, truthgpttypes.ModuleName, deaitypes.ModuleName,
		dynacontracttypes.ModuleName, hyperchaintypes.ModuleName,
	)

	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName, govtypes.ModuleName, stakingtypes.ModuleName,
		capabilitytypes.ModuleName, authtypes.ModuleName, banktypes.ModuleName,
		distrtypes.ModuleName, slashingtypes.ModuleName, evidencetypes.ModuleName,
		genutiltypes.ModuleName, paramstypes.ModuleName, upgradetypes.ModuleName,
		// Custom modules
		neuropostypes.ModuleName, truthgpttypes.ModuleName, deaitypes.ModuleName,
		dynacontracttypes.ModuleName, hyperchaintypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName, authtypes.ModuleName, banktypes.ModuleName, distrtypes.ModuleName,
		stakingtypes.ModuleName, slashingtypes.ModuleName, govtypes.ModuleName, crisistypes.ModuleName,
		genutiltypes.ModuleName, evidencetypes.ModuleName, paramstypes.ModuleName, upgradetypes.ModuleName,
		// Custom modules
		neuropostypes.ModuleName, truthgpttypes.ModuleName, deaitypes.ModuleName,
		dynacontracttypes.ModuleName, hyperchaintypes.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.mm.RegisterServices(module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter()))

	// Create the simulation manager and define the order of the modules for deterministic simulations
	app.sm = module.NewSimulationManager(
		auth.NewAppModuleBasic(authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		params.NewAppModule(app.ParamsKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		// Custom modules with simulation support can be added here
	)

	app.sm.RegisterStoreDecoders()

	// Initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// Initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.SetAnteHandler(
		ante.NewAnteHandler(
			app.AccountKeeper,
			app.BankKeeper,
			ante.DefaultSigVerificationGasConsumer,
			encodingConfig.TxConfig.SignModeHandler(),
		),
	)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

// Name returns the name of the App
func (app *NMXApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *NMXApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *NMXApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tx routes
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register legacy and grpc-gateway routes
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register grpc-gateway routes
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
}

// InitChainer application update at chain initialization
func (app *NMXApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *NMXApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *NMXApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range GetMaccPerms() {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlockedModuleAccountAddrs returns all the blocked module account addresses.
func (app *NMXApp) BlockedModuleAccountAddrs() map[string]bool {
	modAccAddrs := app.ModuleAccountAddrs()
	delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	return modAccAddrs
}

// LegacyAmino returns NMXApp's amino codec.
func (app *NMXApp) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns NMXApp's app codec.
func (app *NMXApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns NMXApp's InterfaceRegistry
func (app *NMXApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
func (app *NMXApp) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
func (app *NMXApp) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
func (app *NMXApp) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
func (app *NMXApp) GetSubspace(moduleName string) paramstypes.Subspace {
	return app.ParamsKeeper.GetSubspace(moduleName)
}

// SimulationManager implements the SimulationApp interface
func (app *NMXApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided API server.
func (app *NMXApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	
	// Register legacy tx routes
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	
	// Register new tx routes
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	
	// Register legacy and grpc-gateway routes
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	
	// Register swagger API if enabled
	if apiConfig.Swagger {
		RegisterSwaggerAPI(clientCtx, apiSvr.Router)
	}
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(ctx client.Context, rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *NMXApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *NMXApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		
		// Custom module permissions
		neuropostypes.ModuleName:       {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		truthgpttypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
		deaitypes.ModuleName:           {authtypes.Minter, authtypes.Burner},
		dynacontracttypes.ModuleName:   {authtypes.Minter, authtypes.Burner},
		hyperchaintypes.ModuleName:     {authtypes.Minter, authtypes.Burner},
	}

	return maccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	// SDK module subspaces
	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(upgradetypes.ModuleName)
	paramsKeeper.Subspace(evidencetypes.ModuleName)
	
	// Custom module subspaces
	paramsKeeper.Subspace(neuropostypes.ModuleName)
	paramsKeeper.Subspace(truthgpttypes.ModuleName)
	paramsKeeper.Subspace(deaitypes.ModuleName)
	paramsKeeper.Subspace(dynacontracttypes.ModuleName)
	paramsKeeper.Subspace(hyperchaintypes.ModuleName)

	return paramsKeeper
}

// skipGenesisInvariants is a helper function for the simulation manager
func skipGenesisInvariants() bool {
	return true
}