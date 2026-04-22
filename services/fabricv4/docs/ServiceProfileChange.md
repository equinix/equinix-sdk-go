# ServiceProfileChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Uniquely identifies a change | [optional] 
**Type** | **string** | Type of change | 
**Status** | Pointer to [**ServiceProfileChangeStatus**](ServiceProfileChangeStatus.md) |  | [optional] 
**CreatedDateTime** | **time.Time** | Set when change flow starts | 
**UpdatedDateTime** | Pointer to **time.Time** | Set when change object is updated | [optional] 
**Information** | Pointer to **string** | Additional information | [optional] 
**Data** | Pointer to [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | [optional] 

## Methods

### NewServiceProfileChange

`func NewServiceProfileChange(type_ string, createdDateTime time.Time, ) *ServiceProfileChange`

NewServiceProfileChange instantiates a new ServiceProfileChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileChangeWithDefaults

`func NewServiceProfileChangeWithDefaults() *ServiceProfileChange`

NewServiceProfileChangeWithDefaults instantiates a new ServiceProfileChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ServiceProfileChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceProfileChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceProfileChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceProfileChange) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *ServiceProfileChange) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceProfileChange) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceProfileChange) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *ServiceProfileChange) GetStatus() ServiceProfileChangeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceProfileChange) GetStatusOk() (*ServiceProfileChangeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceProfileChange) SetStatus(v ServiceProfileChangeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceProfileChange) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *ServiceProfileChange) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *ServiceProfileChange) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *ServiceProfileChange) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.


### GetUpdatedDateTime

`func (o *ServiceProfileChange) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *ServiceProfileChange) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *ServiceProfileChange) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *ServiceProfileChange) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.

### GetInformation

`func (o *ServiceProfileChange) GetInformation() string`

GetInformation returns the Information field if non-nil, zero value otherwise.

### GetInformationOk

`func (o *ServiceProfileChange) GetInformationOk() (*string, bool)`

GetInformationOk returns a tuple with the Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformation

`func (o *ServiceProfileChange) SetInformation(v string)`

SetInformation sets Information field to given value.

### HasInformation

`func (o *ServiceProfileChange) HasInformation() bool`

HasInformation returns a boolean if a field has been set.

### GetData

`func (o *ServiceProfileChange) GetData() []JsonPatchOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceProfileChange) GetDataOk() (*[]JsonPatchOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceProfileChange) SetData(v []JsonPatchOperation)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceProfileChange) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


