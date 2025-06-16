# DeploymentSearchExpressions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to [**SearchDeploymentField**](SearchDeploymentField.md) |  | [optional] 
**Operator** | Pointer to [**DeploymentSearchExpressionsOperator**](DeploymentSearchExpressionsOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeploymentSearchExpressions

`func NewDeploymentSearchExpressions() *DeploymentSearchExpressions`

NewDeploymentSearchExpressions instantiates a new DeploymentSearchExpressions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentSearchExpressionsWithDefaults

`func NewDeploymentSearchExpressionsWithDefaults() *DeploymentSearchExpressions`

NewDeploymentSearchExpressionsWithDefaults instantiates a new DeploymentSearchExpressions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *DeploymentSearchExpressions) GetProperty() SearchDeploymentField`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *DeploymentSearchExpressions) GetPropertyOk() (*SearchDeploymentField, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *DeploymentSearchExpressions) SetProperty(v SearchDeploymentField)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *DeploymentSearchExpressions) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *DeploymentSearchExpressions) GetOperator() DeploymentSearchExpressionsOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *DeploymentSearchExpressions) GetOperatorOk() (*DeploymentSearchExpressionsOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *DeploymentSearchExpressions) SetOperator(v DeploymentSearchExpressionsOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *DeploymentSearchExpressions) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *DeploymentSearchExpressions) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *DeploymentSearchExpressions) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *DeploymentSearchExpressions) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *DeploymentSearchExpressions) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


