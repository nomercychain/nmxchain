package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeRegisterDataSource = "register_data_source"
	TypeUpdateDataSource   = "update_data_source"
	TypeRemoveDataSource   = "remove_data_source"
	TypeCreateOracleQuery  = "create_oracle_query"
	TypeSubmitOracleResponse = "submit_oracle_response"
	TypeRegisterAIModel    = "register_ai_model"
	TypeUpdateAIModel      = "update_ai_model"
	TypeRemoveAIModel      = "remove_ai_model"
	TypeReportMisinformation = "report_misinformation"
	TypeCreateVerificationTask = "create_verification_task"
	TypeCompleteVerificationTask = "complete_verification_task"
)

var _ sdk.Msg = &MsgRegisterDataSource{}

// MsgRegisterDataSource defines a message to register a new data source
type MsgRegisterDataSource struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	SourceType  DataSourceType `json:"source_type"`
	Endpoint    string         `json:"endpoint"`
	Metadata    string         `json:"metadata"`
	Owner       string         `json:"owner"`
}

// NewMsgRegisterDataSource creates a new MsgRegisterDataSource instance
func NewMsgRegisterDataSource(
	name string,
	description string,
	sourceType DataSourceType,
	endpoint string,
	metadata string,
	owner string,
) *MsgRegisterDataSource {
	return &MsgRegisterDataSource{
		Name:        name,
		Description: description,
		SourceType:  sourceType,
		Endpoint:    endpoint,
		Metadata:    metadata,
		Owner:       owner,
	}
}

// Route returns the message route
func (msg MsgRegisterDataSource) Route() string { return RouterKey }

// Type returns the message type
func (msg MsgRegisterDataSource) Type() string { return TypeRegisterDataSource }

// ValidateBasic performs basic validation
func (msg MsgRegisterDataSource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address: %s", err)
	}

	if msg.Name == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}

	if len(msg.Name) > 100 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name too long")
	}

	if len(msg.Description) > 500 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "description too long")
	}

	if msg.Endpoint == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "endpoint cannot be empty")
	}

	valid := false
	validTypes := []DataSourceType{
		DataSourceTypeAPI,
		DataSourceTypeWebsite,
		DataSourceTypeBlockchain,
		DataSourceTypeIPFS,
		DataSourceTypeCustom,
	}

	for _, t := range validTypes {
		if msg.SourceType == t {
			valid = true
			break
		}
	}

	if !valid {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid source type")
	}

	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgRegisterDataSource) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the signers
func (msg MsgRegisterDataSource) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

var _ sdk.Msg = &MsgUpdateDataSource{}

// MsgUpdateDataSource defines a message to update an existing data source
type MsgUpdateDataSource struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	SourceType  DataSourceType `json:"source_type"`
	Endpoint    string         `json:"endpoint"`
	Metadata    string         `json:"metadata"`
	Owner       string         `json:"owner"`
}

// NewMsgUpdateDataSource creates a new MsgUpdateDataSource instance
func NewMsgUpdateDataSource(
	id string,
	name string,
	description string,
	sourceType DataSourceType,
	endpoint string,
	metadata string,
	owner string,
) *MsgUpdateDataSource {
	return &MsgUpdateDataSource{
		ID:          id,
		Name:        name,
		Description: description,
		SourceType:  sourceType,
		Endpoint:    endpoint,
		Metadata:    metadata,
		Owner:       owner,
	}
}

// Route returns the message route
func (msg MsgUpdateDataSource) Route() string { return RouterKey }

// Type returns the message type
func (msg MsgUpdateDataSource) Type() string { return TypeUpdateDataSource }

// ValidateBasic performs basic validation
func (msg MsgUpdateDataSource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address: %s", err)
	}

	if msg.ID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "id cannot be empty")
	}

	if msg.Name == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}

	if len(msg.Name) > 100 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name too long")
	}

	if len(msg.Description) > 500 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "description too long")
	}

	if msg.Endpoint == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "endpoint cannot be empty")
	}

	valid := false
	validTypes := []DataSourceType{
		DataSourceTypeAPI,
		DataSourceTypeWebsite,
		DataSourceTypeBlockchain,
		DataSourceTypeIPFS,
		DataSourceTypeCustom,
	}

	for _, t := range validTypes {
		if msg.SourceType == t {
			valid = true
			break
		}
	}

	if !valid {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid source type")
	}

	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgUpdateDataSource) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the signers
func (msg MsgUpdateDataSource) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

var _ sdk.Msg = &MsgRemoveDataSource{}

// MsgRemoveDataSource defines a message to remove a data source
type MsgRemoveDataSource struct {
	ID    string `json:"id"`
	Owner string `json:"owner"`
}

// NewMsgRemoveDataSource creates a new MsgRemoveDataSource instance
func NewMsgRemoveDataSource(
	id string,
	owner string,
) *MsgRemoveDataSource {
	return &MsgRemoveDataSource{
		ID:    id,
		Owner: owner,
	}
}

// Route returns the message route
func (msg MsgRemoveDataSource) Route() string { return RouterKey }

// Type returns the message type
func (msg MsgRemoveDataSource) Type() string { return TypeRemoveDataSource }

// ValidateBasic performs basic validation
func (msg MsgRemoveDataSource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address: %s", err)
	}

	if msg.ID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "id cannot be empty")
	}

	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgRemoveDataSource) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the signers
func (msg MsgRemoveDataSource) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

var _ sdk.Msg = &MsgCreateOracleQuery{}

// MsgCreateOracleQuery defines a message to create a new oracle query
type MsgCreateOracleQuery struct {
	Requester    string    `json:"requester"`
	QueryType    string    `json:"query_type"`
	Query        string    `json:"query"`
	DataSources  []string  `json:"data_sources,omitempty"`
	Fee          sdk.Coins `json:"fee"`
	CallbackData string    `json:"callback_data,omitempty"`
}

// NewMsgCreateOracleQuery creates a new MsgCreateOracleQuery instance
func NewMsgCreateOracleQuery(
	requester string,
	queryType string,
	query string,
	dataSources []string,
	fee sdk.Coins,
	callbackData string,
) *MsgCreateOracleQuery {
	return &MsgCreateOracleQuery{
		Requester:    requester,
		QueryType:    queryType,
		Query:        query,
		DataSources:  dataSources,
		Fee:          fee,
		CallbackData: callbackData,
	}
}

// Route returns the message route
func (msg MsgCreateOracleQuery) Route() string { return RouterKey }

// Type returns the message type
func (msg MsgCreateOracleQuery) Type() string { return TypeCreateOracleQuery }

// ValidateBasic performs basic validation
func (msg MsgCreateOracleQuery) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Requester)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid requester address: %s", err)
	}

	if msg.QueryType == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "query type cannot be empty")
	}

	if msg.Query == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "query cannot be empty")
	}

	if len(msg.Query) > 1000 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "query too long")
	}

	if !msg.Fee.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Fee.String())
	}

	if msg.Fee.IsAnyNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Fee.String())
	}

	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgCreateOracleQuery) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the signers
func (msg MsgCreateOracleQuery) GetSigners() []sdk.AccAddress {
	requester, err := sdk.AccAddressFromBech32(msg.Requester)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{requester}
}

var _ sdk.Msg = &MsgSubmitOracleResponse{}

// MsgSubmitOracleResponse defines a message to submit a response to an oracle query
type MsgSubmitOracleResponse struct {
	QueryID     string `json:"query_id"`
	Response    string `json:"response"`
	Confidence  string `json:"confidence"`
	Responder   string `json:"responder"`
	SourceID    string `json:"source_id"`
}

// NewMsgSubmitOracleResponse creates a new MsgSubmitOracleResponse instance
func NewMsgSubmitOracleResponse(
	queryID string,
	response string,
	confidence string,
	responder string,
	sourceID string,
) *MsgSubmitOracleResponse {
	return &MsgSubmitOracleResponse{
		QueryID:    queryID,
		Response:   response,
		Confidence: confidence,
		Responder:  responder,
		SourceID:   sourceID,
	}
}

// Route returns the message route
func (msg MsgSubmitOracleResponse) Route() string { return RouterKey }

// Type returns the message type
func (msg MsgSubmitOracleResponse) Type() string { return TypeSubmitOracleResponse }

// ValidateBasic performs basic validation
func (msg MsgSubmitOracleResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Responder)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid responder address: %s", err)
	}

	if msg.QueryID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "query ID cannot be empty")
	}

	if msg.Response == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "response cannot be empty")
	}

	if msg.SourceID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "source ID cannot be empty")
	}

	// Validate confidence (should be a decimal between 0 and 1)
	confidence, err := sdk.NewDecFromStr(msg.Confidence)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid confidence format")
	}

	if confidence.IsNegative() || confidence.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "confidence must be between 0 and 1")
	}

	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgSubmitOracleResponse) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the signers
func (msg MsgSubmitOracleResponse) GetSigners() []sdk.AccAddress {
	responder, err := sdk.AccAddressFromBech32(msg.Responder)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{responder}
}

var _ sdk.Msg = &MsgRegisterAIModel{}

// MsgRegisterAIModel defines a message to register a new AI model
type MsgRegisterAIModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ModelType   string `json:"model_type"`
	ModelURL    string `json:"model_url"`
	ModelHash   string `json:"model_hash"`
	Metadata    string `json:"metadata"`
	Owner       string `json:"owner"`
}

// NewMsgRegisterAIModel creates a new MsgRegisterAIModel instance
func NewMsgRegisterAIModel(
	name string,
	description string,
	modelType string,
	modelURL string,
	modelHash string,
	metadata string,
	owner string,
) *MsgRegisterAIModel {
	return &MsgRegisterAIModel{
		Name:        name,
		Description: description,
		ModelType:   modelType,
		ModelURL:    modelURL,
		ModelHash:   modelHash,
		Metadata:    metadata,
		Owner:       owner,
	}
}

// Route returns the message route
func (msg MsgRegisterAIModel) Route() string { return RouterKey }

// Type returns the message type
func (msg MsgRegisterAIModel) Type() string { return TypeRegisterAIModel }

// ValidateBasic performs basic validation
func (msg MsgRegisterAIModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address: %s", err)
	}

	if msg.Name == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}

	if len(msg.Name) > 100 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name too long")
	}

	if len(msg.Description) > 500 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "description too long")
	}

	if msg.ModelType == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "model type cannot be empty")
	}

	if msg.ModelURL == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "model URL cannot be empty")
	}

	if msg.ModelHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "model hash cannot be empty")
	}

	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgRegisterAIModel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the signers
func (msg MsgRegisterAIModel) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

var _ sdk.Msg = &MsgUpdateAIModel{}

// MsgUpdateAIModel defines a message to update an existing AI model
type MsgUpdateAIModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ModelType   string `json:"model_type"`
	ModelURL    string `json:"model_url"`
	ModelHash   string `json:"model_hash"`
	Metadata    string `json:"metadata"`
	Owner       string `json:"owner"`
}

// NewMsgUpdateAIModel creates a new MsgUpdateAIModel instance
func NewMsgUpdateAIModel(
	id string,
	name string,
	description string,
	modelType string,
	modelURL string,
	modelHash string,
	metadata string,
	owner string,
) *MsgUpdateAIModel {
	return &MsgUpdateAIModel{
		ID:          id,
		Name:        name,
		Description: description,
		ModelType:   modelType,
		ModelURL:    modelURL,
		ModelHash:   modelHash,
		Metadata:    metadata,
		Owner:       owner,
	}
}

// Route returns the message route
func (msg MsgUpdateAIModel) Route() string { return RouterKey }

// Type returns the message type
func (msg MsgUpdateAIModel) Type() string { return TypeUpdateAIModel }

// ValidateBasic performs basic validation
func (msg MsgUpdateAIModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address: %s", err)
	}

	if msg.ID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "id cannot be empty")
	}

	if msg.Name == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}

	if len(msg.Name) > 100 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "name too long")
	}

	if len(msg.Description) > 500 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "description too long")
	}

	if msg.ModelType == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "model type cannot be empty")
	}

	if msg.ModelURL == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "model URL cannot be empty")
	}

	if msg.ModelHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "model hash cannot be empty")
	}

	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgUpdateAIModel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the signers
func (msg MsgUpdateAIModel) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

var _ sdk.Msg = &MsgRemoveAIModel{}

// MsgRemoveAIModel defines a message to remove an AI model
type MsgRemoveAIModel struct {
	ID    string `json:"id"`
	Owner string `json:"owner"`
}

// NewMsgRemoveAIModel creates a new MsgRemoveAIModel instance
func NewMsgRemoveAIModel(
	id string,
	owner string,
) *MsgRemoveAIModel {
	return &MsgRemoveAIModel{
		ID:    id,
		Owner: owner,
	}
}

// Route returns the message route
func (msg MsgRemoveAIModel) Route() string { return RouterKey }

// Type returns the message type
func (msg MsgRemoveAIModel) Type() string { return TypeRemoveAIModel }

// ValidateBasic performs basic validation
func (msg MsgRemoveAIModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address: %s", err)
	}

	if msg.ID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "id cannot be empty")
	}

	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgRemoveAIModel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the signers
func (msg MsgRemoveAIModel) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

var _ sdk.Msg = &MsgReportMisinformation{}

// MsgReportMisinformation defines a message to report misinformation
type MsgReportMisinformation struct {
	Content    string `json:"content"`
	Source     string `json:"source"`
	Evidence   string `json:"evidence"`
	Confidence string `json:"confidence"`
	Reporter   string `json:"reporter"`
}

// NewMsgReportMisinformation creates a new MsgReportMisinformation instance
func NewMsgReportMisinformation(
	content string,
	source string,
	evidence string,
	confidence string,
	reporter string,
) *MsgReportMisinformation {
	return &MsgReportMisinformation{
		Content:    content,
		Source:     source,
		Evidence:   evidence,
		Confidence: confidence,
		Reporter:   reporter,
	}
}

// Route returns the message route
func (msg MsgReportMisinformation) Route() string { return RouterKey }

// Type returns the message type
func (msg MsgReportMisinformation) Type() string { return TypeReportMisinformation }

// ValidateBasic performs basic validation
func (msg MsgReportMisinformation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Reporter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid reporter address: %s", err)
	}

	if msg.Content == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "content cannot be empty")
	}

	if len(msg.Content) > 1000 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "content too long")
	}

	if msg.Source == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "source cannot be empty")
	}

	if msg.Evidence == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "evidence cannot be empty")
	}

	// Validate confidence (should be a decimal between 0 and 1)
	confidence, err := sdk.NewDecFromStr(msg.Confidence)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid confidence format")
	}

	if confidence.IsNegative() || confidence.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "confidence must be between 0 and 1")
	}

	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgReportMisinformation) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the signers
func (msg MsgReportMisinformation) GetSigners() []sdk.AccAddress {
	reporter, err := sdk.AccAddressFromBech32(msg.Reporter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{reporter}
}

var _ sdk.Msg = &MsgCreateVerificationTask{}

// MsgCreateVerificationTask defines a message to create a verification task
type MsgCreateVerificationTask struct {
	Content  string `json:"content"`
	Source   string `json:"source"`
	Priority uint64 `json:"priority"`
	Creator  string `json:"creator"`
}

// NewMsgCreateVerificationTask creates a new MsgCreateVerificationTask instance
func NewMsgCreateVerificationTask(
	content string,
	source string,
	priority uint64,
	creator string,
) *MsgCreateVerificationTask {
	return &MsgCreateVerificationTask{
		Content:  content,
		Source:   source,
		Priority: priority,
		Creator:  creator,
	}
}

// Route returns the message route
func (msg MsgCreateVerificationTask) Route() string { return RouterKey }

// Type returns the message type
func (msg MsgCreateVerificationTask) Type() string { return TypeCreateVerificationTask }

// ValidateBasic performs basic validation
func (msg MsgCreateVerificationTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	if msg.Content == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "content cannot be empty")
	}

	if len(msg.Content) > 1000 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "content too long")
	}

	if msg.Source == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "source cannot be empty")
	}

	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgCreateVerificationTask) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the signers
func (msg MsgCreateVerificationTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

var _ sdk.Msg = &MsgCompleteVerificationTask{}

// MsgCompleteVerificationTask defines a message to complete a verification task
type MsgCompleteVerificationTask struct {
	TaskID   string `json:"task_id"`
	Result   string `json:"result"`
	Verifier string `json:"verifier"`
}

// NewMsgCompleteVerificationTask creates a new MsgCompleteVerificationTask instance
func NewMsgCompleteVerificationTask(
	taskID string,
	result string,
	verifier string,
) *MsgCompleteVerificationTask {
	return &MsgCompleteVerificationTask{
		TaskID:   taskID,
		Result:   result,
		Verifier: verifier,
	}
}

// Route returns the message route
func (msg MsgCompleteVerificationTask) Route() string { return RouterKey }

// Type returns the message type
func (msg MsgCompleteVerificationTask) Type() string { return TypeCompleteVerificationTask }

// ValidateBasic performs basic validation
func (msg MsgCompleteVerificationTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Verifier)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid verifier address: %s", err)
	}

	if msg.TaskID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "task ID cannot be empty")
	}

	if msg.Result == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "result cannot be empty")
	}

	return nil
}

// GetSignBytes returns the bytes to sign
func (msg MsgCompleteVerificationTask) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the signers
func (msg MsgCompleteVerificationTask) GetSigners() []sdk.AccAddress {
	verifier, err := sdk.AccAddressFromBech32(msg.Verifier)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{verifier}
}