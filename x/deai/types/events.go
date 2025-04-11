package types

// Event types for the deai module
const (
	EventTypeCreateAIAgent        = "create_ai_agent"
	EventTypeUpdateAIAgent        = "update_ai_agent"
	EventTypeTrainAIAgent         = "train_ai_agent"
	EventTypeExecuteAIAgent       = "execute_ai_agent"
	EventTypeListAIAgentForSale   = "list_ai_agent_for_sale"
	EventTypeBuyAIAgent           = "buy_ai_agent"
	EventTypeRentAIAgent          = "rent_ai_agent"
	EventTypeCancelMarketListing  = "cancel_market_listing"
	EventTypeTrainingCompleted    = "training_completed"
	EventTypeMarketplaceExpired   = "marketplace_expired"
	
	// Attribute keys
	AttributeKeyAgentID        = "agent_id"
	AttributeKeyCreator        = "creator"
	AttributeKeyOwner          = "owner"
	AttributeKeyName           = "name"
	AttributeKeyModelID        = "model_id"
	AttributeKeyActionID       = "action_id"
	AttributeKeyActionType     = "action_type"
	AttributeKeySender         = "sender"
	AttributeKeyFee            = "fee"
	AttributeKeyListingID      = "listing_id"
	AttributeKeyListingType    = "listing_type"
	AttributeKeyPrice          = "price"
	AttributeKeyRentalPrice    = "rental_price"
	AttributeKeyRentalDuration = "rental_duration"
	AttributeKeyBuyer          = "buyer"
	AttributeKeyRenter         = "renter"
	AttributeKeyExpiresAt      = "expires_at"
	AttributeKeyTrainingDataID = "training_data_id"
	AttributeKeyDataType       = "data_type"
	AttributeKeyTimestamp      = "timestamp"
	AttributeKeyStatus         = "status"
	AttributeKeyResult         = "result"
)