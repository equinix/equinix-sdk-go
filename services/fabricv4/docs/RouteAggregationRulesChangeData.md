# RouteAggregationRulesChangeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Current outcome of the change flow | [optional] 
**CreatedBy** | Pointer to **string** | Created by User Key | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Set when change flow starts | [optional] 
**UpdatedBy** | Pointer to **string** | Updated by User Key | [optional] 
**UpdatedDateTime** | Pointer to **time.Time** | Set when change object is updated | [optional] 
**Data** | Pointer to [**RouteAggregationRulesChangeOperation**](RouteAggregationRulesChangeOperation.md) |  | [optional] 
**Uuid** | **string** | Uniquely identifies a change | 
**Type** | [**RouteAggregationRulesChangeType**](RouteAggregationRulesChangeType.md) |  | 
**Href** | Pointer to **string** | Route Aggregation Change URI | [optional] 

## Methods

### NewRouteAggregationRulesChangeData

`func NewRouteAggregationRulesChangeData(uuid string, type_ RouteAggregationRulesChangeType, ) *RouteAggregationRulesChangeData`

NewRouteAggregationRulesChangeData instantiates a new RouteAggregationRulesChangeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationRulesChangeDataWithDefaults

`func NewRouteAggregationRulesChangeDataWithDefaults() *RouteAggregationRulesChangeData`

NewRouteAggregationRulesChangeDataWithDefaults instantiates a new RouteAggregationRulesChangeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RouteAggregationRulesChangeData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RouteAggregationRulesChangeData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RouteAggregationRulesChangeData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RouteAggregationRulesChangeData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RouteAggregationRulesChangeData) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RouteAggregationRulesChangeData) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RouteAggregationRulesChangeData) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RouteAggregationRulesChangeData) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *RouteAggregationRulesChangeData) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *RouteAggregationRulesChangeData) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *RouteAggregationRulesChangeData) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *RouteAggregationRulesChangeData) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RouteAggregationRulesChangeData) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RouteAggregationRulesChangeData) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RouteAggregationRulesChangeData) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RouteAggregationRulesChangeData) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *RouteAggregationRulesChangeData) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *RouteAggregationRulesChangeData) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *RouteAggregationRulesChangeData) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *RouteAggregationRulesChangeData) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.

### GetData

`func (o *RouteAggregationRulesChangeData) GetData() RouteAggregationRulesChangeOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RouteAggregationRulesChangeData) GetDataOk() (*RouteAggregationRulesChangeOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RouteAggregationRulesChangeData) SetData(v RouteAggregationRulesChangeOperation)`

SetData sets Data field to given value.

### HasData

`func (o *RouteAggregationRulesChangeData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetUuid

`func (o *RouteAggregationRulesChangeData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouteAggregationRulesChangeData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouteAggregationRulesChangeData) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *RouteAggregationRulesChangeData) GetType() RouteAggregationRulesChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteAggregationRulesChangeData) GetTypeOk() (*RouteAggregationRulesChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteAggregationRulesChangeData) SetType(v RouteAggregationRulesChangeType)`

SetType sets Type field to given value.


### GetHref

`func (o *RouteAggregationRulesChangeData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteAggregationRulesChangeData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteAggregationRulesChangeData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouteAggregationRulesChangeData) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


