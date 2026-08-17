# AppServiceChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Uniquely identifies a change | [optional] 
**Type** | [**AppServiceChangeType**](AppServiceChangeType.md) |  | 
**Status** | Pointer to [**PortChangeStatus**](PortChangeStatus.md) |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Set when change flow starts | [optional] 
**UpdatedDateTime** | **time.Time** | Set when change object is updated | 
**Data** | Pointer to [**[]AppServiceChangeOperation**](AppServiceChangeOperation.md) |  | [optional] 

## Methods

### NewAppServiceChange

`func NewAppServiceChange(type_ AppServiceChangeType, updatedDateTime time.Time, ) *AppServiceChange`

NewAppServiceChange instantiates a new AppServiceChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceChangeWithDefaults

`func NewAppServiceChangeWithDefaults() *AppServiceChange`

NewAppServiceChangeWithDefaults instantiates a new AppServiceChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AppServiceChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppServiceChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppServiceChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AppServiceChange) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *AppServiceChange) GetType() AppServiceChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppServiceChange) GetTypeOk() (*AppServiceChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppServiceChange) SetType(v AppServiceChangeType)`

SetType sets Type field to given value.


### GetStatus

`func (o *AppServiceChange) GetStatus() PortChangeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppServiceChange) GetStatusOk() (*PortChangeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppServiceChange) SetStatus(v PortChangeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppServiceChange) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *AppServiceChange) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *AppServiceChange) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *AppServiceChange) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *AppServiceChange) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *AppServiceChange) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *AppServiceChange) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *AppServiceChange) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.


### GetData

`func (o *AppServiceChange) GetData() []AppServiceChangeOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppServiceChange) GetDataOk() (*[]AppServiceChangeOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppServiceChange) SetData(v []AppServiceChangeOperation)`

SetData sets Data field to given value.

### HasData

`func (o *AppServiceChange) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


