package keeper

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/nomercychain/nmxchain/x/truthgpt/types"
)

// Keeper of the truthgpt store
type Keeper struct {
	storeKey   sdk.StoreKey
	memKey     sdk.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

// NewKeeper creates a new truthgpt Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		memKey:        memKey,
		cdc:           cdc,
		paramstore:    ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetDataSource sets a data source
func (k Keeper) SetDataSource(ctx sdk.Context, source types.DataSource) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.DataSourceKey, []byte(source.ID)...)
	value := k.cdc.MustMarshal(&source)
	store.Set(key, value)
}

// GetDataSource returns a data source by ID
func (k Keeper) GetDataSource(ctx sdk.Context, id string) (types.DataSource, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.DataSourceKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.DataSource{}, false
	}

	var source types.DataSource
	k.cdc.MustUnmarshal(value, &source)
	return source, true
}

// DeleteDataSource deletes a data source
func (k Keeper) DeleteDataSource(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.DataSourceKey, []byte(id)...)
	store.Delete(key)
}

// GetAllDataSources returns all data sources
func (k Keeper) GetAllDataSources(ctx sdk.Context) []types.DataSource {
	var sources []types.DataSource
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DataSourceKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var source types.DataSource
		k.cdc.MustUnmarshal(iterator.Value(), &source)
		sources = append(sources, source)
	}

	return sources
}

// SetDataSourceRank sets a data source rank
func (k Keeper) SetDataSourceRank(ctx sdk.Context, rank types.DataSourceRank) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.DataSourceRankKey, []byte(rank.SourceID)...)
	value := k.cdc.MustMarshal(&rank)
	store.Set(key, value)
}

// GetDataSourceRank returns a data source rank by source ID
func (k Keeper) GetDataSourceRank(ctx sdk.Context, sourceID string) (types.DataSourceRank, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.DataSourceRankKey, []byte(sourceID)...)
	value := store.Get(key)
	if value == nil {
		return types.DataSourceRank{}, false
	}

	var rank types.DataSourceRank
	k.cdc.MustUnmarshal(value, &rank)
	return rank, true
}

// GetAllDataSourceRanks returns all data source ranks
func (k Keeper) GetAllDataSourceRanks(ctx sdk.Context) []types.DataSourceRank {
	var ranks []types.DataSourceRank
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DataSourceRankKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var rank types.DataSourceRank
		k.cdc.MustUnmarshal(iterator.Value(), &rank)
		ranks = append(ranks, rank)
	}

	return ranks
}

// SetOracleQuery sets an oracle query
func (k Keeper) SetOracleQuery(ctx sdk.Context, query types.OracleQuery) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.OracleQueryKey, []byte(query.ID)...)
	value := k.cdc.MustMarshal(&query)
	store.Set(key, value)
}

// GetOracleQuery returns an oracle query by ID
func (k Keeper) GetOracleQuery(ctx sdk.Context, id string) (types.OracleQuery, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.OracleQueryKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.OracleQuery{}, false
	}

	var query types.OracleQuery
	k.cdc.MustUnmarshal(value, &query)
	return query, true
}

// GetAllOracleQueries returns all oracle queries
func (k Keeper) GetAllOracleQueries(ctx sdk.Context) []types.OracleQuery {
	var queries []types.OracleQuery
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.OracleQueryKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var query types.OracleQuery
		k.cdc.MustUnmarshal(iterator.Value(), &query)
		queries = append(queries, query)
	}

	return queries
}

// SetOracleResponse sets an oracle response
func (k Keeper) SetOracleResponse(ctx sdk.Context, response types.OracleResponse) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.OracleResponseKey, []byte(response.ID)...)
	value := k.cdc.MustMarshal(&response)
	store.Set(key, value)
}

// GetOracleResponse returns an oracle response by ID
func (k Keeper) GetOracleResponse(ctx sdk.Context, id string) (types.OracleResponse, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.OracleResponseKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.OracleResponse{}, false
	}

	var response types.OracleResponse
	k.cdc.MustUnmarshal(value, &response)
	return response, true
}

// GetAllOracleResponses returns all oracle responses
func (k Keeper) GetAllOracleResponses(ctx sdk.Context) []types.OracleResponse {
	var responses []types.OracleResponse
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.OracleResponseKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var response types.OracleResponse
		k.cdc.MustUnmarshal(iterator.Value(), &response)
		responses = append(responses, response)
	}

	return responses
}

// SetAIModel sets an AI model
func (k Keeper) SetAIModel(ctx sdk.Context, model types.AIModel) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIModelKey, []byte(model.ID)...)
	value := k.cdc.MustMarshal(&model)
	store.Set(key, value)
}

// GetAIModel returns an AI model by ID
func (k Keeper) GetAIModel(ctx sdk.Context, id string) (types.AIModel, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.AIModelKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.AIModel{}, false
	}

	var model types.AIModel
	k.cdc.MustUnmarshal(value, &model)
	return model, true
}

// GetAllAIModels returns all AI models
func (k Keeper) GetAllAIModels(ctx sdk.Context) []types.AIModel {
	var models []types.AIModel
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AIModelKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var model types.AIModel
		k.cdc.MustUnmarshal(iterator.Value(), &model)
		models = append(models, model)
	}

	return models
}

// SetMisinformation sets a misinformation record
func (k Keeper) SetMisinformation(ctx sdk.Context, misinfo types.Misinformation) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.MisinformationKey, []byte(misinfo.ID)...)
	value := k.cdc.MustMarshal(&misinfo)
	store.Set(key, value)
}

// GetMisinformation returns a misinformation record by ID
func (k Keeper) GetMisinformation(ctx sdk.Context, id string) (types.Misinformation, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.MisinformationKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.Misinformation{}, false
	}

	var misinfo types.Misinformation
	k.cdc.MustUnmarshal(value, &misinfo)
	return misinfo, true
}

// GetAllMisinformation returns all misinformation records
func (k Keeper) GetAllMisinformation(ctx sdk.Context) []types.Misinformation {
	var misinfo []types.Misinformation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.MisinformationKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var m types.Misinformation
		k.cdc.MustUnmarshal(iterator.Value(), &m)
		misinfo = append(misinfo, m)
	}

	return misinfo
}

// SetVerificationTask sets a verification task
func (k Keeper) SetVerificationTask(ctx sdk.Context, task types.VerificationTask) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.VerificationTaskKey, []byte(task.ID)...)
	value := k.cdc.MustMarshal(&task)
	store.Set(key, value)
}

// GetVerificationTask returns a verification task by ID
func (k Keeper) GetVerificationTask(ctx sdk.Context, id string) (types.VerificationTask, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.VerificationTaskKey, []byte(id)...)
	value := store.Get(key)
	if value == nil {
		return types.VerificationTask{}, false
	}

	var task types.VerificationTask
	k.cdc.MustUnmarshal(value, &task)
	return task, true
}

// GetAllVerificationTasks returns all verification tasks
func (k Keeper) GetAllVerificationTasks(ctx sdk.Context) []types.VerificationTask {
	var tasks []types.VerificationTask
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.VerificationTaskKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var task types.VerificationTask
		k.cdc.MustUnmarshal(iterator.Value(), &task)
		tasks = append(tasks, task)
	}

	return tasks
}

// CreateDataSource creates a new data source
func (k Keeper) CreateDataSource(ctx sdk.Context, owner sdk.AccAddress, name string, description string, sourceType types.DataSourceType, endpoint string, metadata json.RawMessage) (string, error) {
	// Generate a unique ID for the data source
	id := fmt.Sprintf("%s-%d", owner.String(), ctx.BlockHeight())

	// Create the data source
	source := types.DataSource{
		ID:          id,
		Name:        name,
		Description: description,
		SourceType:  sourceType,
		Endpoint:    endpoint,
		Status:      types.DataSourceStatusPending,
		Owner:       owner,
		CreatedAt:   ctx.BlockTime(),
		UpdatedAt:   ctx.BlockTime(),
		Metadata:    metadata,
	}

	// Store the data source
	k.SetDataSource(ctx, source)

	// Create initial rank
	rank := types.DataSourceRank{
		SourceID:      id,
		Reliability:   sdk.NewDec(5).Quo(sdk.NewDec(10)), // Start at 0.5
		Accuracy:      sdk.NewDec(5).Quo(sdk.NewDec(10)),
		Timeliness:    sdk.NewDec(5).Quo(sdk.NewDec(10)),
		Completeness:  sdk.NewDec(5).Quo(sdk.NewDec(10)),
		TrustScore:    sdk.NewDec(5).Quo(sdk.NewDec(10)),
		LastEvaluated: ctx.BlockTime(),
	}
	k.SetDataSourceRank(ctx, rank)

	return id, nil
}

// SubmitOracleQuery submits a new oracle query
func (k Keeper) SubmitOracleQuery(ctx sdk.Context, requester sdk.AccAddress, queryType string, query string, dataSources []string, fee sdk.Coins, callbackData json.RawMessage) (string, error) {
	// Check if the fee is sufficient
	minFee := sdk.NewCoins(sdk.NewCoin("unmx", sdk.NewInt(100))) // Example minimum fee
	if fee.IsAllLT(minFee) {
		return "", fmt.Errorf("insufficient fee: minimum required is %s", minFee.String())
	}

	// Transfer the fee from the requester to the module
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, requester, types.ModuleName, fee)
	if err != nil {
		return "", err
	}

	// Generate a unique ID for the query
	id := fmt.Sprintf("query-%s-%d", requester.String(), ctx.BlockHeight())

	// Create the query
	oracleQuery := types.OracleQuery{
		ID:          id,
		Requester:   requester,
		QueryType:   queryType,
		Query:       query,
		DataSources: dataSources,
		Status:      types.OracleQueryStatusPending,
		Fee:         fee,
		CreatedAt:   ctx.BlockTime(),
		CallbackData: callbackData,
	}

	// Store the query
	k.SetOracleQuery(ctx, oracleQuery)

	return id, nil
}

// ProcessOracleQuery processes an oracle query
func (k Keeper) ProcessOracleQuery(ctx sdk.Context, queryID string) error {
	// Get the query
	query, found := k.GetOracleQuery(ctx, queryID)
	if !found {
		return fmt.Errorf("query not found: %s", queryID)
	}

	// Check if the query is pending
	if query.Status != types.OracleQueryStatusPending {
		return fmt.Errorf("query is not pending")
	}

	// Update the query status
	query.Status = types.OracleQueryStatusProcessing
	k.SetOracleQuery(ctx, query)

	// In a real implementation, we would:
	// 1. Fetch data from the specified data sources
	// 2. Process the data using AI models
	// 3. Generate a response
	// 4. Update the query status

	// For now, we'll just create a dummy response
	sourceResponses := []types.SourceResponse{}
	for _, sourceID := range query.DataSources {
		source, found := k.GetDataSource(ctx, sourceID)
		if !found {
			continue
		}

		sourceResponse := types.SourceResponse{
			SourceID:  sourceID,
			Response:  json.RawMessage(`{"data": "example data"}`),
			Timestamp: ctx.BlockTime(),
			Status:    "success",
		}
		sourceResponses = append(sourceResponses, sourceResponse)
	}

	// Create the response
	responseID := fmt.Sprintf("response-%s", queryID)
	response := types.OracleResponse{
		ID:              responseID,
		QueryID:         queryID,
		Response:        json.RawMessage(`{"result": "example result"}`),
		SourceResponses: sourceResponses,
		Confidence:      sdk.NewDec(9).Quo(sdk.NewDec(10)), // 0.9
		ProcessedBy:     "default-ai-model",
		CreatedAt:       ctx.BlockTime(),
		Metadata:        json.RawMessage(`{}`),
	}

	// Store the response
	k.SetOracleResponse(ctx, response)

	// Update the query
	query.Status = types.OracleQueryStatusCompleted
	query.CompletedAt = ctx.BlockTime()
	query.ResponseID = responseID
	k.SetOracleQuery(ctx, query)

	return nil
}

// ReportMisinformation reports potential misinformation
func (k Keeper) ReportMisinformation(ctx sdk.Context, reporter sdk.AccAddress, content string, source string, evidence string) (string, error) {
	// Generate a unique ID for the misinformation report
	id := fmt.Sprintf("misinfo-%s-%d", reporter.String(), ctx.BlockHeight())

	// Create the misinformation report
	misinfo := types.Misinformation{
		ID:         id,
		Content:    content,
		Source:     source,
		Reporter:   reporter,
		Confidence: sdk.NewDec(5).Quo(sdk.NewDec(10)), // Start at 0.5
		Evidence:   evidence,
		CreatedAt:  ctx.BlockTime(),
		Status:     "pending",
	}

	// Store the misinformation report
	k.SetMisinformation(ctx, misinfo)

	// Create a verification task
	taskID := fmt.Sprintf("task-%s", id)
	task := types.VerificationTask{
		ID:        taskID,
		Content:   content,
		Source:    source,
		Creator:   reporter,
		Status:    "pending",
		Priority:  1,
		CreatedAt: ctx.BlockTime(),
	}
	k.SetVerificationTask(ctx, task)

	return id, nil
}

// VerifyInformation verifies information using AI models
func (k Keeper) VerifyInformation(ctx sdk.Context, taskID string) error {
	// Get the verification task
	task, found := k.GetVerificationTask(ctx, taskID)
	if !found {
		return fmt.Errorf("verification task not found: %s", taskID)
	}

	// Check if the task is pending
	if task.Status != "pending" {
		return fmt.Errorf("verification task is not pending")
	}

	// In a real implementation, we would:
	// 1. Process the content using AI models
	// 2. Check against trusted data sources
	// 3. Generate a verification result
	// 4. Update the task status

	// For now, we'll just create a dummy result
	task.Status = "completed"
	task.CompletedAt = ctx.BlockTime()
	task.Result = json.RawMessage(`{"verified": true, "confidence": 0.85}`)
	k.SetVerificationTask(ctx, task)

	// If this task was created from a misinformation report, update it
	if strings.HasPrefix(taskID, "task-misinfo-") {
		misinfoID := strings.TrimPrefix(taskID, "task-")
		misinfo, found := k.GetMisinformation(ctx, misinfoID)
		if found {
			// Update based on verification result
			misinfo.Status = "verified"
			misinfo.Confidence = sdk.NewDec(85).Quo(sdk.NewDec(100)) // 0.85
			misinfo.VerifiedBy = append(misinfo.VerifiedBy, "default-ai-model")
			k.SetMisinformation(ctx, misinfo)
		}
	}

	return nil
}

// UpdateDataSourceRankings updates the rankings of data sources
func (k Keeper) UpdateDataSourceRankings(ctx sdk.Context) {
	// Get all data sources
	sources := k.GetAllDataSources(ctx)

	for _, source := range sources {
		// Skip inactive sources
		if source.Status != types.DataSourceStatusActive {
			continue
		}

		// Get the current rank
		rank, found := k.GetDataSourceRank(ctx, source.ID)
		if !found {
			continue
		}

		// In a real implementation, we would:
		// 1. Analyze recent responses from this source
		// 2. Compare with verified information
		// 3. Update the rankings based on performance

		// For now, we'll just make small random adjustments
		// This is just a placeholder for demonstration
		rank.Reliability = adjustRank(rank.Reliability)
		rank.Accuracy = adjustRank(rank.Accuracy)
		rank.Timeliness = adjustRank(rank.Timeliness)
		rank.Completeness = adjustRank(rank.Completeness)

		// Calculate overall trust score
		rank.TrustScore = rank.Reliability.Add(rank.Accuracy).Add(rank.Timeliness).Add(rank.Completeness).Quo(sdk.NewDec(4))
		rank.LastEvaluated = ctx.BlockTime()

		// Store the updated rank
		k.SetDataSourceRank(ctx, rank)
	}
}

// Helper function to adjust a rank slightly
func adjustRank(rank sdk.Dec) sdk.Dec {
	// This is just a placeholder that makes small adjustments
	// In a real implementation, this would be based on actual performance metrics
	adjustment := sdk.NewDec(int64(rand.Intn(10) - 5)).Quo(sdk.NewDec(100)) // -0.05 to +0.05
	newRank := rank.Add(adjustment)
	
	// Ensure the rank stays between 0 and 1
	if newRank.LT(sdk.ZeroDec()) {
		return sdk.ZeroDec()
	}
	if newRank.GT(sdk.OneDec()) {
		return sdk.OneDec()
	}
	
	return newRank
}

// ProcessPendingQueries processes pending oracle queries
func (k Keeper) ProcessPendingQueries(ctx sdk.Context) {
	// Get all pending queries
	queries := k.GetAllOracleQueries(ctx)

	for _, query := range queries {
		// Skip queries that are not pending
		if query.Status != types.OracleQueryStatusPending {
			continue
		}

		// Process the query
		err := k.ProcessOracleQuery(ctx, query.ID)
		if err != nil {
			// Log the error but continue processing other queries
			k.Logger(ctx).Error("Failed to process oracle query", "id", query.ID, "error", err)
		}
	}
}

// ProcessVerificationTasks processes pending verification tasks
func (k Keeper) ProcessVerificationTasks(ctx sdk.Context) {
	// Get all pending verification tasks
	tasks := k.GetAllVerificationTasks(ctx)

	for _, task := range tasks {
		// Skip tasks that are not pending
		if task.Status != "pending" {
			continue
		}

		// Process the task
		err := k.VerifyInformation(ctx, task.ID)
		if err != nil {
			// Log the error but continue processing other tasks
			k.Logger(ctx).Error("Failed to process verification task", "id", task.ID, "error", err)
		}
	}
}