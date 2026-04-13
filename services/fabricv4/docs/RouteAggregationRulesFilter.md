# RouteAggregationRulesFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]RouteAggregationRuleExpression**](RouteAggregationRuleExpression.md) |  | [optional] 
**Or** | Pointer to [**[]RouteAggregationRuleExpression**](RouteAggregationRuleExpression.md) |  | [optional] 

## Methods

### NewRouteAggregationRulesFilter

`func NewRouteAggregationRulesFilter() *RouteAggregationRulesFilter`

NewRouteAggregationRulesFilter instantiates a new RouteAggregationRulesFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationRulesFilterWithDefaults

`func NewRouteAggregationRulesFilterWithDefaults() *RouteAggregationRulesFilter`

NewRouteAggregationRulesFilterWithDefaults instantiates a new RouteAggregationRulesFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *RouteAggregationRulesFilter) GetAnd() []RouteAggregationRuleExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *RouteAggregationRulesFilter) GetAndOk() (*[]RouteAggregationRuleExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *RouteAggregationRulesFilter) SetAnd(v []RouteAggregationRuleExpression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *RouteAggregationRulesFilter) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *RouteAggregationRulesFilter) GetOr() []RouteAggregationRuleExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *RouteAggregationRulesFilter) GetOrOk() (*[]RouteAggregationRuleExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *RouteAggregationRulesFilter) SetOr(v []RouteAggregationRuleExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *RouteAggregationRulesFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


