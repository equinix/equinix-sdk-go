# AppDomainChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Uniquely identifies a change | [optional] 
**Type** | [**AppDomainChangeType**](AppDomainChangeType.md) |  | 
**Status** | Pointer to [**PortChangeStatus**](PortChangeStatus.md) |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Set when change flow starts | [optional] 
**UpdatedDateTime** | **time.Time** | Set when change object is updated | 
**Data** | Pointer to [**[]AppDomainChangeOperation**](AppDomainChangeOperation.md) |  | [optional] 

## Methods

### NewAppDomainChange

`func NewAppDomainChange(type_ AppDomainChangeType, updatedDateTime time.Time, ) *AppDomainChange`

NewAppDomainChange instantiates a new AppDomainChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDomainChangeWithDefaults

`func NewAppDomainChangeWithDefaults() *AppDomainChange`

NewAppDomainChangeWithDefaults instantiates a new AppDomainChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AppDomainChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppDomainChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppDomainChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AppDomainChange) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *AppDomainChange) GetType() AppDomainChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppDomainChange) GetTypeOk() (*AppDomainChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppDomainChange) SetType(v AppDomainChangeType)`

SetType sets Type field to given value.


### GetStatus

`func (o *AppDomainChange) GetStatus() PortChangeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppDomainChange) GetStatusOk() (*PortChangeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppDomainChange) SetStatus(v PortChangeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppDomainChange) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *AppDomainChange) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *AppDomainChange) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *AppDomainChange) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *AppDomainChange) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *AppDomainChange) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *AppDomainChange) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *AppDomainChange) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.


### GetData

`func (o *AppDomainChange) GetData() []AppDomainChangeOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppDomainChange) GetDataOk() (*[]AppDomainChangeOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppDomainChange) SetData(v []AppDomainChangeOperation)`

SetData sets Data field to given value.

### HasData

`func (o *AppDomainChange) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


