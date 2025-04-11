package keeper

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nomercychain/nmxchain/x/neuropos/types"
)

// SetNeuralNetwork sets a neural network in the store
func (k Keeper) SetNeuralNetwork(ctx sdk.Context, network types.NeuralNetwork) {
	store := ctx.KVStore(k.storeKey)
	key := types.NeuralNetworkKey(network.ID)
	value := k.cdc.MustMarshal(&network)
	store.Set(key, value)
}

// GetNeuralNetwork returns a neural network by ID
func (k Keeper) GetNeuralNetwork(ctx sdk.Context, id string) (types.NeuralNetwork, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.NeuralNetworkKey(id)
	value := store.Get(key)
	if value == nil {
		return types.NeuralNetwork{}, false
	}

	var network types.NeuralNetwork
	k.cdc.MustUnmarshal(value, &network)
	return network, true
}

// DeleteNeuralNetwork deletes a neural network
func (k Keeper) DeleteNeuralNetwork(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	key := types.NeuralNetworkKey(id)
	store.Delete(key)
}

// GetAllNeuralNetworks returns all neural networks
func (k Keeper) GetAllNeuralNetworks(ctx sdk.Context) []types.NeuralNetwork {
	var networks []types.NeuralNetwork
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.NeuralNetworkKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var network types.NeuralNetwork
		k.cdc.MustUnmarshal(iterator.Value(), &network)
		networks = append(networks, network)
	}

	return networks
}

// SetNeuralNetworkWeights sets neural network weights in the store
func (k Keeper) SetNeuralNetworkWeights(ctx sdk.Context, weights types.NeuralNetworkWeights) {
	store := ctx.KVStore(k.storeKey)
	key := types.NeuralNetworkWeightKey(weights.NetworkID, weights.Version)
	value := k.cdc.MustMarshal(&weights)
	store.Set(key, value)
}

// GetNeuralNetworkWeights returns neural network weights by network ID and version
func (k Keeper) GetNeuralNetworkWeights(ctx sdk.Context, networkID string, version uint64) (types.NeuralNetworkWeights, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.NeuralNetworkWeightKey(networkID, version)
	value := store.Get(key)
	if value == nil {
		return types.NeuralNetworkWeights{}, false
	}

	var weights types.NeuralNetworkWeights
	k.cdc.MustUnmarshal(value, &weights)
	return weights, true
}

// GetLatestNeuralNetworkWeights returns the latest neural network weights by network ID
func (k Keeper) GetLatestNeuralNetworkWeights(ctx sdk.Context, networkID string) (types.NeuralNetworkWeights, bool) {
	store := ctx.KVStore(k.storeKey)
	prefix := append(types.NeuralNetworkWeightKeyPrefix, []byte(networkID+"/")...)
	iterator := sdk.KVStoreReversePrefixIterator(store, prefix)
	defer iterator.Close()

	if iterator.Valid() {
		var weights types.NeuralNetworkWeights
		k.cdc.MustUnmarshal(iterator.Value(), &weights)
		return weights, true
	}

	return types.NeuralNetworkWeights{}, false
}

// SetTrainingData sets training data in the store
func (k Keeper) SetTrainingData(ctx sdk.Context, data types.TrainingData) {
	store := ctx.KVStore(k.storeKey)
	key := types.TrainingDataKey(data.ID)
	value := k.cdc.MustMarshal(&data)
	store.Set(key, value)
}

// GetTrainingData returns training data by ID
func (k Keeper) GetTrainingData(ctx sdk.Context, id string) (types.TrainingData, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.TrainingDataKey(id)
	value := store.Get(key)
	if value == nil {
		return types.TrainingData{}, false
	}

	var data types.TrainingData
	k.cdc.MustUnmarshal(value, &data)
	return data, true
}

// GetTrainingDataByNetwork returns all training data for a specific network
func (k Keeper) GetTrainingDataByNetwork(ctx sdk.Context, networkID string) []types.TrainingData {
	var dataList []types.TrainingData
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.NeuralNetworkTrainingDataKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var data types.TrainingData
		k.cdc.MustUnmarshal(iterator.Value(), &data)
		if data.NetworkID == networkID {
			dataList = append(dataList, data)
		}
	}

	return dataList
}

// SetNeuralPrediction sets a neural prediction in the store
func (k Keeper) SetNeuralPrediction(ctx sdk.Context, prediction types.NeuralPrediction) {
	store := ctx.KVStore(k.storeKey)
	key := types.NeuralPredictionKey(prediction.ID)
	value := k.cdc.MustMarshal(&prediction)
	store.Set(key, value)
}

// GetNeuralPrediction returns a neural prediction by ID
func (k Keeper) GetNeuralPrediction(ctx sdk.Context, id string) (types.NeuralPrediction, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.NeuralPredictionKey(id)
	value := store.Get(key)
	if value == nil {
		return types.NeuralPrediction{}, false
	}

	var prediction types.NeuralPrediction
	k.cdc.MustUnmarshal(value, &prediction)
	return prediction, true
}

// GetNeuralPredictionsByNetwork returns all neural predictions for a specific network
func (k Keeper) GetNeuralPredictionsByNetwork(ctx sdk.Context, networkID string) []types.NeuralPrediction {
	var predictions []types.NeuralPrediction
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.NeuralNetworkPredictionKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var prediction types.NeuralPrediction
		k.cdc.MustUnmarshal(iterator.Value(), &prediction)
		if prediction.NetworkID == networkID {
			predictions = append(predictions, prediction)
		}
	}

	return predictions
}

// CreateNeuralNetwork creates a new neural network
func (k Keeper) CreateNeuralNetwork(ctx sdk.Context, architecture string, layers []types.Layer, metadata []byte) (types.NeuralNetwork, error) {
	// Generate a unique ID for the network
	id := fmt.Sprintf("nn-%d-%s", ctx.BlockHeight(), ctx.TxHash())

	// Create the neural network
	network := types.NeuralNetwork{
		ID:              id,
		Architecture:    architecture,
		Layers:          layers,
		Status:          types.NeuralNetworkStatusActive,
		Accuracy:        sdk.ZeroDec(),
		Loss:            sdk.NewDec(1),
		LastTrainedTime: time.Time{},
		LastUpdatedTime: ctx.BlockTime(),
		CreatedTime:     ctx.BlockTime(),
		Metadata:        metadata,
	}

	// Validate the neural network
	if !isValidArchitecture(architecture) {
		return types.NeuralNetwork{}, types.ErrInvalidNeuralNetworkArchitecture
	}

	if len(layers) == 0 {
		return types.NeuralNetwork{}, types.ErrInvalidNeuralNetworkArchitecture
	}

	// Save the neural network
	k.SetNeuralNetwork(ctx, network)

	// Create initial weights (zeros)
	initialWeights, err := createInitialWeights(layers)
	if err != nil {
		return types.NeuralNetwork{}, err
	}

	// Save the weights
	weights := types.NeuralNetworkWeights{
		NetworkID: id,
		Weights:   initialWeights,
		UpdatedAt: ctx.BlockTime(),
		Version:   1,
	}
	k.SetNeuralNetworkWeights(ctx, weights)

	return network, nil
}

// UpdateNeuralNetwork updates an existing neural network
func (k Keeper) UpdateNeuralNetwork(ctx sdk.Context, networkID string, architecture string, layers []types.Layer, weights json.RawMessage, metadata []byte) error {
	// Get the existing neural network
	network, found := k.GetNeuralNetwork(ctx, networkID)
	if !found {
		return types.ErrNoNeuralNetworkFound
	}

	// Check if the network is currently being updated or trained
	if network.Status == types.NeuralNetworkStatusUpdating || network.Status == types.NeuralNetworkStatusTraining {
		return types.ErrNeuralNetworkUpdateInProgress
	}

	// Validate the neural network
	if !isValidArchitecture(architecture) {
		return types.ErrInvalidNeuralNetworkArchitecture
	}

	if len(layers) == 0 {
		return types.ErrInvalidNeuralNetworkArchitecture
	}

	if len(weights) == 0 {
		return types.ErrInvalidNeuralNetworkWeights
	}

	// Update the neural network
	network.Architecture = architecture
	network.Layers = layers
	network.Status = types.NeuralNetworkStatusActive
	network.LastUpdatedTime = ctx.BlockTime()
	if metadata != nil {
		network.Metadata = metadata
	}

	// Save the neural network
	k.SetNeuralNetwork(ctx, network)

	// Get the latest weights version
	latestWeights, found := k.GetLatestNeuralNetworkWeights(ctx, networkID)
	var version uint64 = 1
	if found {
		version = latestWeights.Version + 1
	}

	// Save the new weights
	newWeights := types.NeuralNetworkWeights{
		NetworkID: networkID,
		Weights:   weights,
		UpdatedAt: ctx.BlockTime(),
		Version:   version,
	}
	k.SetNeuralNetworkWeights(ctx, newWeights)

	return nil
}

// TrainNeuralNetwork trains a neural network with provided data
func (k Keeper) TrainNeuralNetwork(ctx sdk.Context, networkID string, features json.RawMessage, labels json.RawMessage, epochs uint64, learningRate sdk.Dec, metadata []byte) error {
	// Get the existing neural network
	network, found := k.GetNeuralNetwork(ctx, networkID)
	if !found {
		return types.ErrNoNeuralNetworkFound
	}

	// Check if the network is currently being updated or trained
	if network.Status == types.NeuralNetworkStatusUpdating || network.Status == types.NeuralNetworkStatusTraining {
		return types.ErrNeuralNetworkTrainingInProgress
	}

	// Validate the training data
	if len(features) == 0 {
		return types.ErrInvalidTrainingData
	}

	if len(labels) == 0 {
		return types.ErrInvalidTrainingData
	}

	if epochs == 0 {
		return types.ErrInvalidEpochs
	}

	if learningRate.IsNegative() || learningRate.GT(sdk.OneDec()) {
		return types.ErrInvalidLearningRate
	}

	// Save the training data
	trainingDataID := fmt.Sprintf("td-%d-%s", ctx.BlockHeight(), ctx.TxHash())
	trainingData := types.TrainingData{
		ID:        trainingDataID,
		NetworkID: networkID,
		Features:  features,
		Labels:    labels,
		CreatedAt: ctx.BlockTime(),
		Metadata:  metadata,
	}
	k.SetTrainingData(ctx, trainingData)

	// Update the neural network status
	network.Status = types.NeuralNetworkStatusTraining
	network.LastTrainedTime = ctx.BlockTime()
	k.SetNeuralNetwork(ctx, network)

	// In a real implementation, this would trigger the actual training process
	// For now, we'll simulate training by updating the accuracy and loss
	// This would typically be done in an EndBlocker or a separate process

	// Simulate training results
	network.Accuracy = sdk.NewDecWithPrec(85, 2) // 0.85
	network.Loss = sdk.NewDecWithPrec(15, 2)     // 0.15
	network.Status = types.NeuralNetworkStatusActive
	k.SetNeuralNetwork(ctx, network)

	return nil
}

// SubmitNeuralPrediction submits a prediction from a neural network
func (k Keeper) SubmitNeuralPrediction(ctx sdk.Context, networkID string, input json.RawMessage, output json.RawMessage, confidence sdk.Dec, validatorAddr sdk.ValAddress, metadata []byte) (types.NeuralPrediction, error) {
	// Get the existing neural network
	network, found := k.GetNeuralNetwork(ctx, networkID)
	if !found {
		return types.NeuralPrediction{}, types.ErrNoNeuralNetworkFound
	}

	// Validate the prediction data
	if len(input) == 0 {
		return types.NeuralPrediction{}, types.ErrInvalidPredictionInput
	}

	if len(output) == 0 {
		return types.NeuralPrediction{}, types.ErrInvalidPredictionInput
	}

	if confidence.IsNegative() || confidence.GT(sdk.OneDec()) {
		return types.NeuralPrediction{}, types.ErrInvalidPredictionInput
	}

	// Create the prediction
	predictionID := fmt.Sprintf("pred-%d-%s", ctx.BlockHeight(), ctx.TxHash())
	prediction := types.NeuralPrediction{
		ID:           predictionID,
		NetworkID:    networkID,
		Input:        input,
		Output:       output,
		Confidence:   confidence,
		Timestamp:    ctx.BlockTime(),
		ValidatorSet: []string{validatorAddr.String()},
		Metadata:     metadata,
	}

	// Save the prediction
	k.SetNeuralPrediction(ctx, prediction)

	return prediction, nil
}

// Helper functions

// isValidArchitecture checks if the architecture is valid
func isValidArchitecture(architecture string) bool {
	validArchitectures := map[string]bool{
		types.NeuralNetworkArchitectureMLP:       true,
		types.NeuralNetworkArchitectureCNN:       true,
		types.NeuralNetworkArchitectureRNN:       true,
		types.NeuralNetworkArchitectureLSTM:      true,
		types.NeuralNetworkArchitectureGRU:       true,
		types.NeuralNetworkArchitectureTransformer: true,
		types.NeuralNetworkArchitectureAutoencoder: true,
		types.NeuralNetworkArchitectureGAN:       true,
	}

	_, found := validArchitectures[architecture]
	return found
}

// createInitialWeights creates initial weights for a neural network
func createInitialWeights(layers []types.Layer) (json.RawMessage, error) {
	// In a real implementation, this would create appropriate initial weights
	// based on the network architecture and layers
	// For now, we'll create a simple placeholder

	// Create a simple structure to represent weights
	weights := make([]map[string]interface{}, len(layers))
	for i, layer := range layers {
		weights[i] = map[string]interface{}{
			"layer_type":  layer.Type,
			"input_size":  layer.InputSize,
			"output_size": layer.OutputSize,
			"weights":     make([]float64, layer.InputSize*layer.OutputSize), // Initialize with zeros
			"biases":      make([]float64, layer.OutputSize),                 // Initialize with zeros
		}
	}

	// Convert to JSON
	weightsJSON, err := json.Marshal(weights)
	if err != nil {
		return nil, err
	}

	return weightsJSON, nil
}