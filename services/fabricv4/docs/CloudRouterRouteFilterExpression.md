# CloudRouterRouteFilterExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]CloudRouterRouteFilterExpression**](CloudRouterRouteFilterExpression.md) |  | [optional] 
**Or** | Pointer to [**[]CloudRouterRouteFilterExpression**](CloudRouterRouteFilterExpression.md) |  | [optional] 
**Property** | Pointer to [**CloudRouterRouteFilterSimpleExpressionProperty**](CloudRouterRouteFilterSimpleExpressionProperty.md) |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCloudRouterRouteFilterExpression

`func NewCloudRouterRouteFilterExpression() *CloudRouterRouteFilterExpression`

NewCloudRouterRouteFilterExpression instantiates a new CloudRouterRouteFilterExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterRouteFilterExpressionWithDefaults

`func NewCloudRouterRouteFilterExpressionWithDefaults() *CloudRouterRouteFilterExpression`

NewCloudRouterRouteFilterExpressionWithDefaults instantiates a new CloudRouterRouteFilterExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *CloudRouterRouteFilterExpression) GetAnd() []CloudRouterRouteFilterExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *CloudRouterRouteFilterExpression) GetAndOk() (*[]CloudRouterRouteFilterExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *CloudRouterRouteFilterExpression) SetAnd(v []CloudRouterRouteFilterExpression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *CloudRouterRouteFilterExpression) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *CloudRouterRouteFilterExpression) GetOr() []CloudRouterRouteFilterExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *CloudRouterRouteFilterExpression) GetOrOk() (*[]CloudRouterRouteFilterExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *CloudRouterRouteFilterExpression) SetOr(v []CloudRouterRouteFilterExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *CloudRouterRouteFilterExpression) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetProperty

`func (o *CloudRouterRouteFilterExpression) GetProperty() CloudRouterRouteFilterSimpleExpressionProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CloudRouterRouteFilterExpression) GetPropertyOk() (*CloudRouterRouteFilterSimpleExpressionProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CloudRouterRouteFilterExpression) SetProperty(v CloudRouterRouteFilterSimpleExpressionProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *CloudRouterRouteFilterExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *CloudRouterRouteFilterExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CloudRouterRouteFilterExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CloudRouterRouteFilterExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CloudRouterRouteFilterExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *CloudRouterRouteFilterExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CloudRouterRouteFilterExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CloudRouterRouteFilterExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *CloudRouterRouteFilterExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


