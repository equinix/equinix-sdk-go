# RouteAggregationRuleSimpleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/type&#x60; - Route Aggregation Rules Type  * &#x60;/name&#x60; - Route Aggregation Rules Name  * &#x60;/uuid&#x60; - Route Aggregation Rules uuid  * &#x60;/state&#x60; - Route Aggregation Rules status  * &#x60;/prefix&#x60; - Route Aggregation Rule Prefix  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;!&#x3D;&#x60; - not equal  * &#x60;[NOT] LIKE&#x60; - (not) like  * &#x60;[NOT] IN&#x60; - (not) in  * &#x60;ILIKE&#x60; - case-insensitive like  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRouteAggregationRuleSimpleExpression

`func NewRouteAggregationRuleSimpleExpression() *RouteAggregationRuleSimpleExpression`

NewRouteAggregationRuleSimpleExpression instantiates a new RouteAggregationRuleSimpleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationRuleSimpleExpressionWithDefaults

`func NewRouteAggregationRuleSimpleExpressionWithDefaults() *RouteAggregationRuleSimpleExpression`

NewRouteAggregationRuleSimpleExpressionWithDefaults instantiates a new RouteAggregationRuleSimpleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *RouteAggregationRuleSimpleExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *RouteAggregationRuleSimpleExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *RouteAggregationRuleSimpleExpression) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *RouteAggregationRuleSimpleExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *RouteAggregationRuleSimpleExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RouteAggregationRuleSimpleExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RouteAggregationRuleSimpleExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *RouteAggregationRuleSimpleExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *RouteAggregationRuleSimpleExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RouteAggregationRuleSimpleExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RouteAggregationRuleSimpleExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *RouteAggregationRuleSimpleExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


