# RouteAggregationRulesChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Uniquely identifies a change | 
**Type** | [**RouteAggregationRulesChangeType**](RouteAggregationRulesChangeType.md) |  | 
**Href** | Pointer to **string** | Route Aggregation Change URI | [optional] 

## Methods

### NewRouteAggregationRulesChange

`func NewRouteAggregationRulesChange(uuid string, type_ RouteAggregationRulesChangeType, ) *RouteAggregationRulesChange`

NewRouteAggregationRulesChange instantiates a new RouteAggregationRulesChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationRulesChangeWithDefaults

`func NewRouteAggregationRulesChangeWithDefaults() *RouteAggregationRulesChange`

NewRouteAggregationRulesChangeWithDefaults instantiates a new RouteAggregationRulesChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RouteAggregationRulesChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouteAggregationRulesChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouteAggregationRulesChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *RouteAggregationRulesChange) GetType() RouteAggregationRulesChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteAggregationRulesChange) GetTypeOk() (*RouteAggregationRulesChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteAggregationRulesChange) SetType(v RouteAggregationRulesChangeType)`

SetType sets Type field to given value.


### GetHref

`func (o *RouteAggregationRulesChange) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteAggregationRulesChange) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteAggregationRulesChange) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouteAggregationRulesChange) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


