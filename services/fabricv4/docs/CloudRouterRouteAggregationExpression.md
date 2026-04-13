# CloudRouterRouteAggregationExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]CloudRouterRouteAggregationExpression**](CloudRouterRouteAggregationExpression.md) |  | [optional] 
**Or** | Pointer to [**[]CloudRouterRouteAggregationExpression**](CloudRouterRouteAggregationExpression.md) |  | [optional] 
**Property** | Pointer to [**CloudRouterRouteAggregationSimpleExpressionProperty**](CloudRouterRouteAggregationSimpleExpressionProperty.md) |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCloudRouterRouteAggregationExpression

`func NewCloudRouterRouteAggregationExpression() *CloudRouterRouteAggregationExpression`

NewCloudRouterRouteAggregationExpression instantiates a new CloudRouterRouteAggregationExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterRouteAggregationExpressionWithDefaults

`func NewCloudRouterRouteAggregationExpressionWithDefaults() *CloudRouterRouteAggregationExpression`

NewCloudRouterRouteAggregationExpressionWithDefaults instantiates a new CloudRouterRouteAggregationExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *CloudRouterRouteAggregationExpression) GetAnd() []CloudRouterRouteAggregationExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *CloudRouterRouteAggregationExpression) GetAndOk() (*[]CloudRouterRouteAggregationExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *CloudRouterRouteAggregationExpression) SetAnd(v []CloudRouterRouteAggregationExpression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *CloudRouterRouteAggregationExpression) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *CloudRouterRouteAggregationExpression) GetOr() []CloudRouterRouteAggregationExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *CloudRouterRouteAggregationExpression) GetOrOk() (*[]CloudRouterRouteAggregationExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *CloudRouterRouteAggregationExpression) SetOr(v []CloudRouterRouteAggregationExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *CloudRouterRouteAggregationExpression) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetProperty

`func (o *CloudRouterRouteAggregationExpression) GetProperty() CloudRouterRouteAggregationSimpleExpressionProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CloudRouterRouteAggregationExpression) GetPropertyOk() (*CloudRouterRouteAggregationSimpleExpressionProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CloudRouterRouteAggregationExpression) SetProperty(v CloudRouterRouteAggregationSimpleExpressionProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *CloudRouterRouteAggregationExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *CloudRouterRouteAggregationExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CloudRouterRouteAggregationExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CloudRouterRouteAggregationExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CloudRouterRouteAggregationExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *CloudRouterRouteAggregationExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CloudRouterRouteAggregationExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CloudRouterRouteAggregationExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *CloudRouterRouteAggregationExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


