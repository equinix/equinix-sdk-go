# PortChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Uniquely identifies a change | [optional] 
**Type** | Pointer to [**PortChangeType**](PortChangeType.md) |  | [optional] 
**Status** | Pointer to [**PortChangeStatus**](PortChangeStatus.md) |  | [optional] 
**Information** | Pointer to **string** | Additional information | [optional] 
**Data** | Pointer to [**PortChangeOperation**](PortChangeOperation.md) |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Set when change flow starts | [optional] 
**UpdatedDateTime** | Pointer to **time.Time** | Set when change object is updated | [optional] 

## Methods

### NewPortChange

`func NewPortChange() *PortChange`

NewPortChange instantiates a new PortChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortChangeWithDefaults

`func NewPortChangeWithDefaults() *PortChange`

NewPortChangeWithDefaults instantiates a new PortChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PortChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PortChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PortChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PortChange) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *PortChange) GetType() PortChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortChange) GetTypeOk() (*PortChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortChange) SetType(v PortChangeType)`

SetType sets Type field to given value.

### HasType

`func (o *PortChange) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *PortChange) GetStatus() PortChangeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortChange) GetStatusOk() (*PortChangeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortChange) SetStatus(v PortChangeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PortChange) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInformation

`func (o *PortChange) GetInformation() string`

GetInformation returns the Information field if non-nil, zero value otherwise.

### GetInformationOk

`func (o *PortChange) GetInformationOk() (*string, bool)`

GetInformationOk returns a tuple with the Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformation

`func (o *PortChange) SetInformation(v string)`

SetInformation sets Information field to given value.

### HasInformation

`func (o *PortChange) HasInformation() bool`

HasInformation returns a boolean if a field has been set.

### GetData

`func (o *PortChange) GetData() PortChangeOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PortChange) GetDataOk() (*PortChangeOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PortChange) SetData(v PortChangeOperation)`

SetData sets Data field to given value.

### HasData

`func (o *PortChange) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *PortChange) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *PortChange) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *PortChange) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *PortChange) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *PortChange) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *PortChange) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *PortChange) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *PortChange) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


