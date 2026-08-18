# AppLinkChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Uniquely identifies a change | [optional] 
**Type** | [**AppLinkChangeType**](AppLinkChangeType.md) |  | 
**Status** | Pointer to [**PortChangeStatus**](PortChangeStatus.md) |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Set when change flow starts | [optional] 
**UpdatedDateTime** | **time.Time** | Set when change object is updated | 
**Data** | Pointer to [**[]AppLinkChangeOperation**](AppLinkChangeOperation.md) |  | [optional] 

## Methods

### NewAppLinkChange

`func NewAppLinkChange(type_ AppLinkChangeType, updatedDateTime time.Time, ) *AppLinkChange`

NewAppLinkChange instantiates a new AppLinkChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkChangeWithDefaults

`func NewAppLinkChangeWithDefaults() *AppLinkChange`

NewAppLinkChangeWithDefaults instantiates a new AppLinkChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AppLinkChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppLinkChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppLinkChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AppLinkChange) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *AppLinkChange) GetType() AppLinkChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppLinkChange) GetTypeOk() (*AppLinkChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppLinkChange) SetType(v AppLinkChangeType)`

SetType sets Type field to given value.


### GetStatus

`func (o *AppLinkChange) GetStatus() PortChangeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppLinkChange) GetStatusOk() (*PortChangeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppLinkChange) SetStatus(v PortChangeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppLinkChange) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *AppLinkChange) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *AppLinkChange) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *AppLinkChange) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *AppLinkChange) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *AppLinkChange) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *AppLinkChange) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *AppLinkChange) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.


### GetData

`func (o *AppLinkChange) GetData() []AppLinkChangeOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppLinkChange) GetDataOk() (*[]AppLinkChangeOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppLinkChange) SetData(v []AppLinkChangeOperation)`

SetData sets Data field to given value.

### HasData

`func (o *AppLinkChange) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


