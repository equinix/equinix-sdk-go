# RoutingProtocolChangeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Current outcome of the change flow | [optional] 
**CreatedBy** | Pointer to **string** | Created by User Key | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Set when change flow starts | [optional] 
**UpdatedBy** | Pointer to **string** | Updated by User Key | [optional] 
**UpdatedDateTime** | Pointer to **time.Time** | Set when change object is updated | [optional] 
**Information** | Pointer to **string** | Additional information | [optional] 
**Data** | Pointer to [**RoutingProtocolChangeOperation**](RoutingProtocolChangeOperation.md) |  | [optional] 
**Uuid** | **string** | Uniquely identifies a change | 
**Type** | [**RoutingProtocolChangeType**](RoutingProtocolChangeType.md) |  | 
**Href** | Pointer to **string** | Routing Protocol Change URI | [optional] 

## Methods

### NewRoutingProtocolChangeData

`func NewRoutingProtocolChangeData(uuid string, type_ RoutingProtocolChangeType, ) *RoutingProtocolChangeData`

NewRoutingProtocolChangeData instantiates a new RoutingProtocolChangeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolChangeDataWithDefaults

`func NewRoutingProtocolChangeDataWithDefaults() *RoutingProtocolChangeData`

NewRoutingProtocolChangeDataWithDefaults instantiates a new RoutingProtocolChangeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RoutingProtocolChangeData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RoutingProtocolChangeData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RoutingProtocolChangeData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RoutingProtocolChangeData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RoutingProtocolChangeData) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RoutingProtocolChangeData) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RoutingProtocolChangeData) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RoutingProtocolChangeData) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *RoutingProtocolChangeData) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *RoutingProtocolChangeData) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *RoutingProtocolChangeData) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *RoutingProtocolChangeData) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RoutingProtocolChangeData) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RoutingProtocolChangeData) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RoutingProtocolChangeData) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RoutingProtocolChangeData) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *RoutingProtocolChangeData) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *RoutingProtocolChangeData) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *RoutingProtocolChangeData) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *RoutingProtocolChangeData) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.

### GetInformation

`func (o *RoutingProtocolChangeData) GetInformation() string`

GetInformation returns the Information field if non-nil, zero value otherwise.

### GetInformationOk

`func (o *RoutingProtocolChangeData) GetInformationOk() (*string, bool)`

GetInformationOk returns a tuple with the Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformation

`func (o *RoutingProtocolChangeData) SetInformation(v string)`

SetInformation sets Information field to given value.

### HasInformation

`func (o *RoutingProtocolChangeData) HasInformation() bool`

HasInformation returns a boolean if a field has been set.

### GetData

`func (o *RoutingProtocolChangeData) GetData() RoutingProtocolChangeOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RoutingProtocolChangeData) GetDataOk() (*RoutingProtocolChangeOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RoutingProtocolChangeData) SetData(v RoutingProtocolChangeOperation)`

SetData sets Data field to given value.

### HasData

`func (o *RoutingProtocolChangeData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetUuid

`func (o *RoutingProtocolChangeData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoutingProtocolChangeData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoutingProtocolChangeData) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *RoutingProtocolChangeData) GetType() RoutingProtocolChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolChangeData) GetTypeOk() (*RoutingProtocolChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolChangeData) SetType(v RoutingProtocolChangeType)`

SetType sets Type field to given value.


### GetHref

`func (o *RoutingProtocolChangeData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RoutingProtocolChangeData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RoutingProtocolChangeData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RoutingProtocolChangeData) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


