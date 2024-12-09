# RouteAggregationChangeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Current outcome of the change flow | [optional] 
**CreatedBy** | Pointer to **string** | Created by User Key | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Set when change flow starts | [optional] 
**UpdatedBy** | Pointer to **string** | Updated by User Key | [optional] 
**UpdatedDateTime** | Pointer to **time.Time** | Set when change object is updated | [optional] 
**Information** | Pointer to **string** | Additional information | [optional] 
**Data** | Pointer to [**RouteAggregationsChangeOperation**](RouteAggregationsChangeOperation.md) |  | [optional] 
**Uuid** | **string** | Uniquely identifies a change | 
**Type** | [**RouteAggregationsChangeType**](RouteAggregationsChangeType.md) |  | 
**Href** | Pointer to **string** | Route AGGREGATION Change URI | [optional] 

## Methods

### NewRouteAggregationChangeData

`func NewRouteAggregationChangeData(uuid string, type_ RouteAggregationsChangeType, ) *RouteAggregationChangeData`

NewRouteAggregationChangeData instantiates a new RouteAggregationChangeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationChangeDataWithDefaults

`func NewRouteAggregationChangeDataWithDefaults() *RouteAggregationChangeData`

NewRouteAggregationChangeDataWithDefaults instantiates a new RouteAggregationChangeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RouteAggregationChangeData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RouteAggregationChangeData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RouteAggregationChangeData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RouteAggregationChangeData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RouteAggregationChangeData) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RouteAggregationChangeData) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RouteAggregationChangeData) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RouteAggregationChangeData) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *RouteAggregationChangeData) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *RouteAggregationChangeData) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *RouteAggregationChangeData) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *RouteAggregationChangeData) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RouteAggregationChangeData) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RouteAggregationChangeData) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RouteAggregationChangeData) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RouteAggregationChangeData) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *RouteAggregationChangeData) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *RouteAggregationChangeData) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *RouteAggregationChangeData) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *RouteAggregationChangeData) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.

### GetInformation

`func (o *RouteAggregationChangeData) GetInformation() string`

GetInformation returns the Information field if non-nil, zero value otherwise.

### GetInformationOk

`func (o *RouteAggregationChangeData) GetInformationOk() (*string, bool)`

GetInformationOk returns a tuple with the Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformation

`func (o *RouteAggregationChangeData) SetInformation(v string)`

SetInformation sets Information field to given value.

### HasInformation

`func (o *RouteAggregationChangeData) HasInformation() bool`

HasInformation returns a boolean if a field has been set.

### GetData

`func (o *RouteAggregationChangeData) GetData() RouteAggregationsChangeOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RouteAggregationChangeData) GetDataOk() (*RouteAggregationsChangeOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RouteAggregationChangeData) SetData(v RouteAggregationsChangeOperation)`

SetData sets Data field to given value.

### HasData

`func (o *RouteAggregationChangeData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetUuid

`func (o *RouteAggregationChangeData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouteAggregationChangeData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouteAggregationChangeData) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *RouteAggregationChangeData) GetType() RouteAggregationsChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteAggregationChangeData) GetTypeOk() (*RouteAggregationsChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteAggregationChangeData) SetType(v RouteAggregationsChangeType)`

SetType sets Type field to given value.


### GetHref

`func (o *RouteAggregationChangeData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteAggregationChangeData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteAggregationChangeData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouteAggregationChangeData) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


