# RouteFilterRuleSimpleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/type&#x60; - Route Filter Rules Type  * &#x60;/name&#x60; - Route Filter Rules Name  * &#x60;/uuid&#x60; - Route Filter Rules uuid  * &#x60;/state&#x60; - Route Filter Rules status  * &#x60;/prefix&#x60; - Route Filter Rule Prefix  * &#x60;/prefixMatch&#x60; - Route Filter Rule Prefix Match  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;!&#x3D;&#x60; - not equal  * &#x60;[NOT] LIKE&#x60; - (not) like  * &#x60;[NOT] IN&#x60; - (not) in  * &#x60;ILIKE&#x60; - case-insensitive like  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRouteFilterRuleSimpleExpression

`func NewRouteFilterRuleSimpleExpression() *RouteFilterRuleSimpleExpression`

NewRouteFilterRuleSimpleExpression instantiates a new RouteFilterRuleSimpleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFilterRuleSimpleExpressionWithDefaults

`func NewRouteFilterRuleSimpleExpressionWithDefaults() *RouteFilterRuleSimpleExpression`

NewRouteFilterRuleSimpleExpressionWithDefaults instantiates a new RouteFilterRuleSimpleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *RouteFilterRuleSimpleExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *RouteFilterRuleSimpleExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *RouteFilterRuleSimpleExpression) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *RouteFilterRuleSimpleExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *RouteFilterRuleSimpleExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RouteFilterRuleSimpleExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RouteFilterRuleSimpleExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *RouteFilterRuleSimpleExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *RouteFilterRuleSimpleExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RouteFilterRuleSimpleExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RouteFilterRuleSimpleExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *RouteFilterRuleSimpleExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


