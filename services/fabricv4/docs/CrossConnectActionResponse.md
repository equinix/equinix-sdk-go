# CrossConnectActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Action URI. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Action Identifier. | [optional] [readonly] 
**Type** | Pointer to [**CrossConnectActionType**](CrossConnectActionType.md) |  | [optional] [readonly] 
**State** | Pointer to [**CrossConnectActionState**](CrossConnectActionState.md) |  | [optional] [readonly] 
**Data** | Pointer to [**CrossConnectActionResponseData**](CrossConnectActionResponseData.md) |  | [optional] [readonly] 
**ChangeLog** | Pointer to [**CrossConnectChangelog**](CrossConnectChangelog.md) |  | [optional] [readonly] 

## Methods

### NewCrossConnectActionResponse

`func NewCrossConnectActionResponse() *CrossConnectActionResponse`

NewCrossConnectActionResponse instantiates a new CrossConnectActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectActionResponseWithDefaults

`func NewCrossConnectActionResponseWithDefaults() *CrossConnectActionResponse`

NewCrossConnectActionResponseWithDefaults instantiates a new CrossConnectActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CrossConnectActionResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CrossConnectActionResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CrossConnectActionResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CrossConnectActionResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *CrossConnectActionResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CrossConnectActionResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CrossConnectActionResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CrossConnectActionResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *CrossConnectActionResponse) GetType() CrossConnectActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossConnectActionResponse) GetTypeOk() (*CrossConnectActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossConnectActionResponse) SetType(v CrossConnectActionType)`

SetType sets Type field to given value.

### HasType

`func (o *CrossConnectActionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *CrossConnectActionResponse) GetState() CrossConnectActionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CrossConnectActionResponse) GetStateOk() (*CrossConnectActionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CrossConnectActionResponse) SetState(v CrossConnectActionState)`

SetState sets State field to given value.

### HasState

`func (o *CrossConnectActionResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *CrossConnectActionResponse) GetData() CrossConnectActionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CrossConnectActionResponse) GetDataOk() (*CrossConnectActionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CrossConnectActionResponse) SetData(v CrossConnectActionResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CrossConnectActionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetChangeLog

`func (o *CrossConnectActionResponse) GetChangeLog() CrossConnectChangelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *CrossConnectActionResponse) GetChangeLogOk() (*CrossConnectChangelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *CrossConnectActionResponse) SetChangeLog(v CrossConnectChangelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *CrossConnectActionResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


