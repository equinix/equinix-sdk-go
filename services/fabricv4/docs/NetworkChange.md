# NetworkChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Network URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Uniquely identifies a change | [optional] 
**Type** | Pointer to [**NetworkChangeType**](NetworkChangeType.md) |  | [optional] 
**Status** | Pointer to [**NetworkChangeStatus**](NetworkChangeStatus.md) |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Set when change flow starts | [optional] 
**UpdatedDateTime** | Pointer to **time.Time** | Set when change object is updated | [optional] 
**Data** | Pointer to [**[]NetworkChangeOperation**](NetworkChangeOperation.md) |  | [optional] 

## Methods

### NewNetworkChange

`func NewNetworkChange() *NetworkChange`

NewNetworkChange instantiates a new NetworkChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkChangeWithDefaults

`func NewNetworkChangeWithDefaults() *NetworkChange`

NewNetworkChangeWithDefaults instantiates a new NetworkChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NetworkChange) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NetworkChange) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NetworkChange) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *NetworkChange) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *NetworkChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NetworkChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NetworkChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NetworkChange) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *NetworkChange) GetType() NetworkChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkChange) GetTypeOk() (*NetworkChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkChange) SetType(v NetworkChangeType)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkChange) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkChange) GetStatus() NetworkChangeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkChange) GetStatusOk() (*NetworkChangeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkChange) SetStatus(v NetworkChangeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkChange) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *NetworkChange) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *NetworkChange) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *NetworkChange) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *NetworkChange) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *NetworkChange) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *NetworkChange) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *NetworkChange) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *NetworkChange) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.

### GetData

`func (o *NetworkChange) GetData() []NetworkChangeOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkChange) GetDataOk() (*[]NetworkChangeOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkChange) SetData(v []NetworkChangeOperation)`

SetData sets Data field to given value.

### HasData

`func (o *NetworkChange) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


