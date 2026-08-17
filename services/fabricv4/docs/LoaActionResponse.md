# LoaActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Action URI | [optional] 
**Uuid** | Pointer to **string** | Action Identifier | [optional] 
**State** | Pointer to [**LoaActionState**](LoaActionState.md) |  | [optional] 
**Type** | Pointer to [**LoaActionType**](LoaActionType.md) |  | [optional] 
**Data** | Pointer to [**LoaActionData**](LoaActionData.md) |  | [optional] 
**ChangeLog** | Pointer to [**LoaChangelog**](LoaChangelog.md) |  | [optional] 

## Methods

### NewLoaActionResponse

`func NewLoaActionResponse() *LoaActionResponse`

NewLoaActionResponse instantiates a new LoaActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaActionResponseWithDefaults

`func NewLoaActionResponseWithDefaults() *LoaActionResponse`

NewLoaActionResponseWithDefaults instantiates a new LoaActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *LoaActionResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LoaActionResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LoaActionResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LoaActionResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *LoaActionResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LoaActionResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LoaActionResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LoaActionResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *LoaActionResponse) GetState() LoaActionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LoaActionResponse) GetStateOk() (*LoaActionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LoaActionResponse) SetState(v LoaActionState)`

SetState sets State field to given value.

### HasState

`func (o *LoaActionResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *LoaActionResponse) GetType() LoaActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoaActionResponse) GetTypeOk() (*LoaActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoaActionResponse) SetType(v LoaActionType)`

SetType sets Type field to given value.

### HasType

`func (o *LoaActionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *LoaActionResponse) GetData() LoaActionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoaActionResponse) GetDataOk() (*LoaActionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoaActionResponse) SetData(v LoaActionData)`

SetData sets Data field to given value.

### HasData

`func (o *LoaActionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetChangeLog

`func (o *LoaActionResponse) GetChangeLog() LoaChangelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *LoaActionResponse) GetChangeLogOk() (*LoaChangelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *LoaActionResponse) SetChangeLog(v LoaChangelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *LoaActionResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


