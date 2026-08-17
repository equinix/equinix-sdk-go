# AppSubscriptionChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Uniquely identifies a change | [optional] 
**Type** | [**AppSubscriptionChangeType**](AppSubscriptionChangeType.md) |  | 
**Status** | Pointer to [**PortChangeStatus**](PortChangeStatus.md) |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Set when change flow starts | [optional] 
**UpdatedDateTime** | **time.Time** | Set when change object is updated | 
**Data** | Pointer to [**[]AppSubscriptionChangeOperation**](AppSubscriptionChangeOperation.md) |  | [optional] 

## Methods

### NewAppSubscriptionChange

`func NewAppSubscriptionChange(type_ AppSubscriptionChangeType, updatedDateTime time.Time, ) *AppSubscriptionChange`

NewAppSubscriptionChange instantiates a new AppSubscriptionChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubscriptionChangeWithDefaults

`func NewAppSubscriptionChangeWithDefaults() *AppSubscriptionChange`

NewAppSubscriptionChangeWithDefaults instantiates a new AppSubscriptionChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AppSubscriptionChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppSubscriptionChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppSubscriptionChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AppSubscriptionChange) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *AppSubscriptionChange) GetType() AppSubscriptionChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppSubscriptionChange) GetTypeOk() (*AppSubscriptionChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppSubscriptionChange) SetType(v AppSubscriptionChangeType)`

SetType sets Type field to given value.


### GetStatus

`func (o *AppSubscriptionChange) GetStatus() PortChangeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppSubscriptionChange) GetStatusOk() (*PortChangeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppSubscriptionChange) SetStatus(v PortChangeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppSubscriptionChange) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *AppSubscriptionChange) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *AppSubscriptionChange) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *AppSubscriptionChange) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *AppSubscriptionChange) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *AppSubscriptionChange) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *AppSubscriptionChange) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *AppSubscriptionChange) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.


### GetData

`func (o *AppSubscriptionChange) GetData() []AppSubscriptionChangeOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppSubscriptionChange) GetDataOk() (*[]AppSubscriptionChangeOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppSubscriptionChange) SetData(v []AppSubscriptionChangeOperation)`

SetData sets Data field to given value.

### HasData

`func (o *AppSubscriptionChange) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


