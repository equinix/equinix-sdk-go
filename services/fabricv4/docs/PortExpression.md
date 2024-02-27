# PortExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]PortExpression**](PortExpression.md) |  | [optional] 
**Or** | Pointer to [**[]PortExpression**](PortExpression.md) |  | [optional] 
**Property** | Pointer to [**PortSearchFieldName**](PortSearchFieldName.md) |  | [optional] 
**Operator** | Pointer to [**ServiceTokenSearchExpressionOperator**](ServiceTokenSearchExpressionOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPortExpression

`func NewPortExpression() *PortExpression`

NewPortExpression instantiates a new PortExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortExpressionWithDefaults

`func NewPortExpressionWithDefaults() *PortExpression`

NewPortExpressionWithDefaults instantiates a new PortExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *PortExpression) GetAnd() []PortExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *PortExpression) GetAndOk() (*[]PortExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *PortExpression) SetAnd(v []PortExpression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *PortExpression) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *PortExpression) GetOr() []PortExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *PortExpression) GetOrOk() (*[]PortExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *PortExpression) SetOr(v []PortExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *PortExpression) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetProperty

`func (o *PortExpression) GetProperty() PortSearchFieldName`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PortExpression) GetPropertyOk() (*PortSearchFieldName, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PortExpression) SetProperty(v PortSearchFieldName)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *PortExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *PortExpression) GetOperator() ServiceTokenSearchExpressionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PortExpression) GetOperatorOk() (*ServiceTokenSearchExpressionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PortExpression) SetOperator(v ServiceTokenSearchExpressionOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *PortExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *PortExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PortExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PortExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *PortExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


