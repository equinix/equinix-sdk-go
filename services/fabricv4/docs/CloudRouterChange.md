# CloudRouterChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Uniquely identifies a change | [optional] 
**Type** | [**CloudRouterChangeType**](CloudRouterChangeType.md) |  | 
**Status** | Pointer to [**CloudRouterChangeStatus**](CloudRouterChangeStatus.md) |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Set when change flow starts | [optional] 
**UpdatedDateTime** | **time.Time** | Set when change object is updated | 
**Information** | Pointer to **string** | Additional information | [optional] 
**Data** | Pointer to [**CloudRouterChangeOperation**](CloudRouterChangeOperation.md) |  | [optional] 

## Methods

### NewCloudRouterChange

`func NewCloudRouterChange(type_ CloudRouterChangeType, updatedDateTime time.Time, ) *CloudRouterChange`

NewCloudRouterChange instantiates a new CloudRouterChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterChangeWithDefaults

`func NewCloudRouterChangeWithDefaults() *CloudRouterChange`

NewCloudRouterChangeWithDefaults instantiates a new CloudRouterChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *CloudRouterChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CloudRouterChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CloudRouterChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CloudRouterChange) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *CloudRouterChange) GetType() CloudRouterChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudRouterChange) GetTypeOk() (*CloudRouterChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudRouterChange) SetType(v CloudRouterChangeType)`

SetType sets Type field to given value.


### GetStatus

`func (o *CloudRouterChange) GetStatus() CloudRouterChangeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudRouterChange) GetStatusOk() (*CloudRouterChangeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudRouterChange) SetStatus(v CloudRouterChangeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudRouterChange) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *CloudRouterChange) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *CloudRouterChange) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *CloudRouterChange) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *CloudRouterChange) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *CloudRouterChange) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *CloudRouterChange) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *CloudRouterChange) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.


### GetInformation

`func (o *CloudRouterChange) GetInformation() string`

GetInformation returns the Information field if non-nil, zero value otherwise.

### GetInformationOk

`func (o *CloudRouterChange) GetInformationOk() (*string, bool)`

GetInformationOk returns a tuple with the Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformation

`func (o *CloudRouterChange) SetInformation(v string)`

SetInformation sets Information field to given value.

### HasInformation

`func (o *CloudRouterChange) HasInformation() bool`

HasInformation returns a boolean if a field has been set.

### GetData

`func (o *CloudRouterChange) GetData() CloudRouterChangeOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudRouterChange) GetDataOk() (*CloudRouterChangeOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudRouterChange) SetData(v CloudRouterChangeOperation)`

SetData sets Data field to given value.

### HasData

`func (o *CloudRouterChange) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


