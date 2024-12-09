# RouteFilterChangeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Current outcome of the change flow | [optional] 
**CreatedBy** | Pointer to **string** | Created by user key | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Set when change flow starts | [optional] 
**UpdatedBy** | Pointer to **string** | Updated by user key | [optional] 
**UpdatedDateTime** | Pointer to **time.Time** | Set when change object is updated | [optional] 
**Information** | Pointer to **string** | Additional information | [optional] 
**Data** | Pointer to [**RouteFiltersChangeOperation**](RouteFiltersChangeOperation.md) |  | [optional] 
**Uuid** | **string** | Uniquely identifies a change | 
**Type** | [**RouteFiltersChangeType**](RouteFiltersChangeType.md) |  | 
**Href** | Pointer to **string** | Route filter change URI | [optional] 

## Methods

### NewRouteFilterChangeData

`func NewRouteFilterChangeData(uuid string, type_ RouteFiltersChangeType, ) *RouteFilterChangeData`

NewRouteFilterChangeData instantiates a new RouteFilterChangeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFilterChangeDataWithDefaults

`func NewRouteFilterChangeDataWithDefaults() *RouteFilterChangeData`

NewRouteFilterChangeDataWithDefaults instantiates a new RouteFilterChangeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RouteFilterChangeData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RouteFilterChangeData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RouteFilterChangeData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RouteFilterChangeData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RouteFilterChangeData) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RouteFilterChangeData) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RouteFilterChangeData) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RouteFilterChangeData) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *RouteFilterChangeData) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *RouteFilterChangeData) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *RouteFilterChangeData) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *RouteFilterChangeData) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RouteFilterChangeData) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RouteFilterChangeData) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RouteFilterChangeData) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RouteFilterChangeData) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *RouteFilterChangeData) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *RouteFilterChangeData) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *RouteFilterChangeData) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *RouteFilterChangeData) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.

### GetInformation

`func (o *RouteFilterChangeData) GetInformation() string`

GetInformation returns the Information field if non-nil, zero value otherwise.

### GetInformationOk

`func (o *RouteFilterChangeData) GetInformationOk() (*string, bool)`

GetInformationOk returns a tuple with the Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformation

`func (o *RouteFilterChangeData) SetInformation(v string)`

SetInformation sets Information field to given value.

### HasInformation

`func (o *RouteFilterChangeData) HasInformation() bool`

HasInformation returns a boolean if a field has been set.

### GetData

`func (o *RouteFilterChangeData) GetData() RouteFiltersChangeOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RouteFilterChangeData) GetDataOk() (*RouteFiltersChangeOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RouteFilterChangeData) SetData(v RouteFiltersChangeOperation)`

SetData sets Data field to given value.

### HasData

`func (o *RouteFilterChangeData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetUuid

`func (o *RouteFilterChangeData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouteFilterChangeData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouteFilterChangeData) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *RouteFilterChangeData) GetType() RouteFiltersChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteFilterChangeData) GetTypeOk() (*RouteFiltersChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteFilterChangeData) SetType(v RouteFiltersChangeType)`

SetType sets Type field to given value.


### GetHref

`func (o *RouteFilterChangeData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteFilterChangeData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteFilterChangeData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouteFilterChangeData) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


