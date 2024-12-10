# RouteAggregationRulesBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Customer-provided Route Aggregation Rule description | [optional] 
**Prefix** | **string** |  | 

## Methods

### NewRouteAggregationRulesBase

`func NewRouteAggregationRulesBase(prefix string, ) *RouteAggregationRulesBase`

NewRouteAggregationRulesBase instantiates a new RouteAggregationRulesBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationRulesBaseWithDefaults

`func NewRouteAggregationRulesBaseWithDefaults() *RouteAggregationRulesBase`

NewRouteAggregationRulesBaseWithDefaults instantiates a new RouteAggregationRulesBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RouteAggregationRulesBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteAggregationRulesBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteAggregationRulesBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouteAggregationRulesBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RouteAggregationRulesBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RouteAggregationRulesBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RouteAggregationRulesBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RouteAggregationRulesBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrefix

`func (o *RouteAggregationRulesBase) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RouteAggregationRulesBase) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RouteAggregationRulesBase) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


