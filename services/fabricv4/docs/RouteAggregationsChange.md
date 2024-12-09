# RouteAggregationsChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Uniquely identifies a change | 
**Type** | [**RouteAggregationsChangeType**](RouteAggregationsChangeType.md) |  | 
**Href** | Pointer to **string** | Route AGGREGATION Change URI | [optional] 

## Methods

### NewRouteAggregationsChange

`func NewRouteAggregationsChange(uuid string, type_ RouteAggregationsChangeType, ) *RouteAggregationsChange`

NewRouteAggregationsChange instantiates a new RouteAggregationsChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationsChangeWithDefaults

`func NewRouteAggregationsChangeWithDefaults() *RouteAggregationsChange`

NewRouteAggregationsChangeWithDefaults instantiates a new RouteAggregationsChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RouteAggregationsChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouteAggregationsChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouteAggregationsChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *RouteAggregationsChange) GetType() RouteAggregationsChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteAggregationsChange) GetTypeOk() (*RouteAggregationsChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteAggregationsChange) SetType(v RouteAggregationsChangeType)`

SetType sets Type field to given value.


### GetHref

`func (o *RouteAggregationsChange) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteAggregationsChange) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteAggregationsChange) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouteAggregationsChange) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


