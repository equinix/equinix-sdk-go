# RouteAggregationRulesPatchRequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** | Handy shortcut for operation name | 
**Path** | **string** | path to change | 
**Value** | **map[string]interface{}** | new value for updated parameter | 

## Methods

### NewRouteAggregationRulesPatchRequestItem

`func NewRouteAggregationRulesPatchRequestItem(op string, path string, value map[string]interface{}, ) *RouteAggregationRulesPatchRequestItem`

NewRouteAggregationRulesPatchRequestItem instantiates a new RouteAggregationRulesPatchRequestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationRulesPatchRequestItemWithDefaults

`func NewRouteAggregationRulesPatchRequestItemWithDefaults() *RouteAggregationRulesPatchRequestItem`

NewRouteAggregationRulesPatchRequestItemWithDefaults instantiates a new RouteAggregationRulesPatchRequestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *RouteAggregationRulesPatchRequestItem) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RouteAggregationRulesPatchRequestItem) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RouteAggregationRulesPatchRequestItem) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *RouteAggregationRulesPatchRequestItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RouteAggregationRulesPatchRequestItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RouteAggregationRulesPatchRequestItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *RouteAggregationRulesPatchRequestItem) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RouteAggregationRulesPatchRequestItem) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RouteAggregationRulesPatchRequestItem) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


