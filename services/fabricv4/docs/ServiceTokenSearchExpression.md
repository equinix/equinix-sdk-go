# ServiceTokenSearchExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]ServiceTokenSearchExpression**](ServiceTokenSearchExpression.md) |  | [optional] 
**Property** | Pointer to [**ServiceTokenSearchFieldName**](ServiceTokenSearchFieldName.md) |  | [optional] 
**Operator** | Pointer to [**ServiceTokenSearchExpressionOperator**](ServiceTokenSearchExpressionOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServiceTokenSearchExpression

`func NewServiceTokenSearchExpression() *ServiceTokenSearchExpression`

NewServiceTokenSearchExpression instantiates a new ServiceTokenSearchExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTokenSearchExpressionWithDefaults

`func NewServiceTokenSearchExpressionWithDefaults() *ServiceTokenSearchExpression`

NewServiceTokenSearchExpressionWithDefaults instantiates a new ServiceTokenSearchExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *ServiceTokenSearchExpression) GetAnd() []ServiceTokenSearchExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *ServiceTokenSearchExpression) GetAndOk() (*[]ServiceTokenSearchExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *ServiceTokenSearchExpression) SetAnd(v []ServiceTokenSearchExpression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *ServiceTokenSearchExpression) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetProperty

`func (o *ServiceTokenSearchExpression) GetProperty() ServiceTokenSearchFieldName`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ServiceTokenSearchExpression) GetPropertyOk() (*ServiceTokenSearchFieldName, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ServiceTokenSearchExpression) SetProperty(v ServiceTokenSearchFieldName)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ServiceTokenSearchExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *ServiceTokenSearchExpression) GetOperator() ServiceTokenSearchExpressionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ServiceTokenSearchExpression) GetOperatorOk() (*ServiceTokenSearchExpressionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ServiceTokenSearchExpression) SetOperator(v ServiceTokenSearchExpressionOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ServiceTokenSearchExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *ServiceTokenSearchExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ServiceTokenSearchExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ServiceTokenSearchExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ServiceTokenSearchExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


