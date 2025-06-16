# ProviderSearchExpressions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to [**SearchProviderField**](SearchProviderField.md) |  | [optional] 
**Operator** | Pointer to [**DeploymentSearchExpressionsOperator**](DeploymentSearchExpressionsOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProviderSearchExpressions

`func NewProviderSearchExpressions() *ProviderSearchExpressions`

NewProviderSearchExpressions instantiates a new ProviderSearchExpressions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSearchExpressionsWithDefaults

`func NewProviderSearchExpressionsWithDefaults() *ProviderSearchExpressions`

NewProviderSearchExpressionsWithDefaults instantiates a new ProviderSearchExpressions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *ProviderSearchExpressions) GetProperty() SearchProviderField`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ProviderSearchExpressions) GetPropertyOk() (*SearchProviderField, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ProviderSearchExpressions) SetProperty(v SearchProviderField)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ProviderSearchExpressions) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *ProviderSearchExpressions) GetOperator() DeploymentSearchExpressionsOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ProviderSearchExpressions) GetOperatorOk() (*DeploymentSearchExpressionsOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ProviderSearchExpressions) SetOperator(v DeploymentSearchExpressionsOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ProviderSearchExpressions) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *ProviderSearchExpressions) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ProviderSearchExpressions) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ProviderSearchExpressions) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ProviderSearchExpressions) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


