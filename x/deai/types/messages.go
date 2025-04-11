package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Message types
const (
	TypeMsgCreateAIAgent        = "create_ai_agent"
	TypeMsgUpdateAIAgent        = "update_ai_agent"
	TypeMsgTrainAIAgent         = "train_ai_agent"
	TypeMsgExecuteAIAgent       = "execute_ai_agent"
	TypeMsgListAIAgentForSale   = "list_ai_agent_for_sale"
	TypeMsgBuyAIAgent           = "buy_ai_agent"
	TypeMsgRentAIAgent          = "rent_ai_agent"
	TypeMsgCancelMarketListing  = "cancel_market_listing"
)

var (
	_ sdk.Msg = &MsgCreateAIAgent{}
	_ sdk.Msg = &MsgUpdateAIAgent{}
	_ sdk.Msg = &MsgTrainAIAgent{}
	_ sdk.Msg = &MsgExecuteAIAgent{}
	_ sdk.Msg = &MsgListAIAgentForSale{}
	_ sdk.Msg = &MsgBuyAIAgent{}
	_ sdk.Msg = &MsgRentAIAgent{}
	_ sdk.Msg = &MsgCancelMarketListing{}
)

// MsgCreateAIAgent defines a message to create a new AI agent
type MsgCreateAIAgent struct {
	Creator     sdk.AccAddress  `json:"creator"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	AgentType   AIAgentType     `json:"agent_type"`
	ModelID     string          `json:"model_id"`
	Permissions json.RawMessage `json:"permissions"`
	Metadata    json.RawMessage `json:"metadata"`
}

// NewMsgCreateAIAgent creates a new MsgCreateAIAgent instance
func NewMsgCreateAIAgent(
	creator sdk.AccAddress,
	name string,
	description string,
	agentType AIAgentType,
	modelID string,
	permissions json.RawMessage,
	metadata json.RawMessage,
) *MsgCreateAIAgent {
	return &MsgCreateAIAgent{
		Creator:     creator,
		Name:        name,
		Description: description,
		AgentType:   agentType,
		ModelID:     modelID,
		Permissions: permissions,
		Metadata:    metadata,
	}
}

// Route returns the message route
func (msg MsgCreateAIAgent) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg MsgCreateAIAgent) Type() string {
	return TypeMsgCreateAIAgent
}

// ValidateBasic performs basic validation
func (msg MsgCreateAIAgent) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator address cannot be empty")
	}
	if msg.Name == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}
	if msg.ModelID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "model ID cannot be empty")
	}
	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgCreateAIAgent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the signers
func (msg MsgCreateAIAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// MsgUpdateAIAgent defines a message to update an AI agent
type MsgUpdateAIAgent struct {
	Owner       sdk.AccAddress  `json:"owner"`
	AgentID     string          `json:"agent_id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Permissions json.RawMessage `json:"permissions"`
	Metadata    json.RawMessage `json:"metadata"`
}

// NewMsgUpdateAIAgent creates a new MsgUpdateAIAgent instance
func NewMsgUpdateAIAgent(
	owner sdk.AccAddress,
	agentID string,
	name string,
	description string,
	permissions json.RawMessage,
	metadata json.RawMessage,
) *MsgUpdateAIAgent {
	return &MsgUpdateAIAgent{
		Owner:       owner,
		AgentID:     agentID,
		Name:        name,
		Description: description,
		Permissions: permissions,
		Metadata:    metadata,
	}
}

// Route returns the message route
func (msg MsgUpdateAIAgent) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg MsgUpdateAIAgent) Type() string {
	return TypeMsgUpdateAIAgent
}

// ValidateBasic performs basic validation
func (msg MsgUpdateAIAgent) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner address cannot be empty")
	}
	if msg.AgentID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "agent ID cannot be empty")
	}
	if msg.Name == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}
	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgUpdateAIAgent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the signers
func (msg MsgUpdateAIAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgTrainAIAgent defines a message to train an AI agent
type MsgTrainAIAgent struct {
	Owner     sdk.AccAddress  `json:"owner"`
	AgentID   string          `json:"agent_id"`
	DataType  string          `json:"data_type"`
	Data      json.RawMessage `json:"data"`
	Source    string          `json:"source"`
}

// NewMsgTrainAIAgent creates a new MsgTrainAIAgent instance
func NewMsgTrainAIAgent(
	owner sdk.AccAddress,
	agentID string,
	dataType string,
	data json.RawMessage,
	source string,
) *MsgTrainAIAgent {
	return &MsgTrainAIAgent{
		Owner:     owner,
		AgentID:   agentID,
		DataType:  dataType,
		Data:      data,
		Source:    source,
	}
}

// Route returns the message route
func (msg MsgTrainAIAgent) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg MsgTrainAIAgent) Type() string {
	return TypeMsgTrainAIAgent
}

// ValidateBasic performs basic validation
func (msg MsgTrainAIAgent) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner address cannot be empty")
	}
	if msg.AgentID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "agent ID cannot be empty")
	}
	if msg.DataType == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "data type cannot be empty")
	}
	if len(msg.Data) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "data cannot be empty")
	}
	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgTrainAIAgent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the signers
func (msg MsgTrainAIAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgExecuteAIAgent defines a message to execute an AI agent
type MsgExecuteAIAgent struct {
	Sender     sdk.AccAddress  `json:"sender"`
	AgentID    string          `json:"agent_id"`
	ActionType string          `json:"action_type"`
	Data       json.RawMessage `json:"data"`
	Fee        sdk.Coin        `json:"fee"`
}

// NewMsgExecuteAIAgent creates a new MsgExecuteAIAgent instance
func NewMsgExecuteAIAgent(
	sender sdk.AccAddress,
	agentID string,
	actionType string,
	data json.RawMessage,
	fee sdk.Coin,
) *MsgExecuteAIAgent {
	return &MsgExecuteAIAgent{
		Sender:     sender,
		AgentID:    agentID,
		ActionType: actionType,
		Data:       data,
		Fee:        fee,
	}
}

// Route returns the message route
func (msg MsgExecuteAIAgent) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg MsgExecuteAIAgent) Type() string {
	return TypeMsgExecuteAIAgent
}

// ValidateBasic performs basic validation
func (msg MsgExecuteAIAgent) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	if msg.AgentID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "agent ID cannot be empty")
	}
	if msg.ActionType == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "action type cannot be empty")
	}
	if msg.Fee.IsZero() || !msg.Fee.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "fee must be valid and non-zero")
	}
	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgExecuteAIAgent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the signers
func (msg MsgExecuteAIAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// MsgListAIAgentForSale defines a message to list an AI agent for sale
type MsgListAIAgentForSale struct {
	Seller          sdk.AccAddress `json:"seller"`
	AgentID         string         `json:"agent_id"`
	Price           sdk.Coins      `json:"price"`
	RentalPrice     sdk.Coins      `json:"rental_price,omitempty"`
	RentalDuration  uint64         `json:"rental_duration,omitempty"`
	ListingType     string         `json:"listing_type"` // "sale", "rent", "both"
	ExpirationDays  uint64         `json:"expiration_days"`
}

// NewMsgListAIAgentForSale creates a new MsgListAIAgentForSale instance
func NewMsgListAIAgentForSale(
	seller sdk.AccAddress,
	agentID string,
	price sdk.Coins,
	rentalPrice sdk.Coins,
	rentalDuration uint64,
	listingType string,
	expirationDays uint64,
) *MsgListAIAgentForSale {
	return &MsgListAIAgentForSale{
		Seller:          seller,
		AgentID:         agentID,
		Price:           price,
		RentalPrice:     rentalPrice,
		RentalDuration:  rentalDuration,
		ListingType:     listingType,
		ExpirationDays:  expirationDays,
	}
}

// Route returns the message route
func (msg MsgListAIAgentForSale) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg MsgListAIAgentForSale) Type() string {
	return TypeMsgListAIAgentForSale
}

// ValidateBasic performs basic validation
func (msg MsgListAIAgentForSale) ValidateBasic() error {
	if msg.Seller.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "seller address cannot be empty")
	}
	if msg.AgentID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "agent ID cannot be empty")
	}
	if msg.ListingType != "sale" && msg.ListingType != "rent" && msg.ListingType != "both" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "listing type must be 'sale', 'rent', or 'both'")
	}
	if (msg.ListingType == "sale" || msg.ListingType == "both") && !msg.Price.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "price must be valid for sale listings")
	}
	if (msg.ListingType == "rent" || msg.ListingType == "both") && (!msg.RentalPrice.IsValid() || msg.RentalDuration == 0) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "rental price and duration must be valid for rental listings")
	}
	if msg.ExpirationDays == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "expiration days cannot be zero")
	}
	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgListAIAgentForSale) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the signers
func (msg MsgListAIAgentForSale) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Seller}
}

// MsgBuyAIAgent defines a message to buy an AI agent
type MsgBuyAIAgent struct {
	Buyer     sdk.AccAddress `json:"buyer"`
	ListingID string         `json:"listing_id"`
}

// NewMsgBuyAIAgent creates a new MsgBuyAIAgent instance
func NewMsgBuyAIAgent(
	buyer sdk.AccAddress,
	listingID string,
) *MsgBuyAIAgent {
	return &MsgBuyAIAgent{
		Buyer:     buyer,
		ListingID: listingID,
	}
}

// Route returns the message route
func (msg MsgBuyAIAgent) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg MsgBuyAIAgent) Type() string {
	return TypeMsgBuyAIAgent
}

// ValidateBasic performs basic validation
func (msg MsgBuyAIAgent) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "buyer address cannot be empty")
	}
	if msg.ListingID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "listing ID cannot be empty")
	}
	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgBuyAIAgent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the signers
func (msg MsgBuyAIAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

// MsgRentAIAgent defines a message to rent an AI agent
type MsgRentAIAgent struct {
	Renter    sdk.AccAddress `json:"renter"`
	ListingID string         `json:"listing_id"`
	Duration  uint64         `json:"duration"`
}

// NewMsgRentAIAgent creates a new MsgRentAIAgent instance
func NewMsgRentAIAgent(
	renter sdk.AccAddress,
	listingID string,
	duration uint64,
) *MsgRentAIAgent {
	return &MsgRentAIAgent{
		Renter:    renter,
		ListingID: listingID,
		Duration:  duration,
	}
}

// Route returns the message route
func (msg MsgRentAIAgent) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg MsgRentAIAgent) Type() string {
	return TypeMsgRentAIAgent
}

// ValidateBasic performs basic validation
func (msg MsgRentAIAgent) ValidateBasic() error {
	if msg.Renter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "renter address cannot be empty")
	}
	if msg.ListingID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "listing ID cannot be empty")
	}
	if msg.Duration == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "duration cannot be zero")
	}
	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgRentAIAgent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the signers
func (msg MsgRentAIAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Renter}
}

// MsgCancelMarketListing defines a message to cancel a marketplace listing
type MsgCancelMarketListing struct {
	Owner     sdk.AccAddress `json:"owner"`
	ListingID string         `json:"listing_id"`
}

// NewMsgCancelMarketListing creates a new MsgCancelMarketListing instance
func NewMsgCancelMarketListing(
	owner sdk.AccAddress,
	listingID string,
) *MsgCancelMarketListing {
	return &MsgCancelMarketListing{
		Owner:     owner,
		ListingID: listingID,
	}
}

// Route returns the message route
func (msg MsgCancelMarketListing) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg MsgCancelMarketListing) Type() string {
	return TypeMsgCancelMarketListing
}

// ValidateBasic performs basic validation
func (msg MsgCancelMarketListing) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner address cannot be empty")
	}
	if msg.ListingID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "listing ID cannot be empty")
	}
	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgCancelMarketListing) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the signers
func (msg MsgCancelMarketListing) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}