# DeploymentExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]DeploymentSearchExpressions**](DeploymentSearchExpressions.md) |  | [optional] 
**Or** | Pointer to [**[]DeploymentSearchExpressions**](DeploymentSearchExpressions.md) |  | [optional] 

## Methods

### NewDeploymentExpression

`func NewDeploymentExpression() *DeploymentExpression`

NewDeploymentExpression instantiates a new DeploymentExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentExpressionWithDefaults

`func NewDeploymentExpressionWithDefaults() *DeploymentExpression`

NewDeploymentExpressionWithDefaults instantiates a new DeploymentExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *DeploymentExpression) GetAnd() []DeploymentSearchExpressions`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *DeploymentExpression) GetAndOk() (*[]DeploymentSearchExpressions, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *DeploymentExpression) SetAnd(v []DeploymentSearchExpressions)`

SetAnd sets And field to given value.

### HasAnd

`func (o *DeploymentExpression) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *DeploymentExpression) GetOr() []DeploymentSearchExpressions`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *DeploymentExpression) GetOrOk() (*[]DeploymentSearchExpressions, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *DeploymentExpression) SetOr(v []DeploymentSearchExpressions)`

SetOr sets Or field to given value.

### HasOr

`func (o *DeploymentExpression) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


