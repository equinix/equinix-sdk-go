# PortCapacityExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**PortCapacityExpression**](PortCapacityExpression.md) |  | [optional] 
**Or** | Pointer to [**PortCapacityExpression**](PortCapacityExpression.md) |  | [optional] 
**Property** | Pointer to [**PortCapacitySearchFieldName**](PortCapacitySearchFieldName.md) |  | [optional] 
**Operator** | Pointer to [**ServiceTokenSearchExpressionOperator**](ServiceTokenSearchExpressionOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPortCapacityExpression

`func NewPortCapacityExpression() *PortCapacityExpression`

NewPortCapacityExpression instantiates a new PortCapacityExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortCapacityExpressionWithDefaults

`func NewPortCapacityExpressionWithDefaults() *PortCapacityExpression`

NewPortCapacityExpressionWithDefaults instantiates a new PortCapacityExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *PortCapacityExpression) GetAnd() PortCapacityExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *PortCapacityExpression) GetAndOk() (*PortCapacityExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *PortCapacityExpression) SetAnd(v PortCapacityExpression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *PortCapacityExpression) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *PortCapacityExpression) GetOr() PortCapacityExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *PortCapacityExpression) GetOrOk() (*PortCapacityExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *PortCapacityExpression) SetOr(v PortCapacityExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *PortCapacityExpression) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetProperty

`func (o *PortCapacityExpression) GetProperty() PortCapacitySearchFieldName`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PortCapacityExpression) GetPropertyOk() (*PortCapacitySearchFieldName, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PortCapacityExpression) SetProperty(v PortCapacitySearchFieldName)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *PortCapacityExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *PortCapacityExpression) GetOperator() ServiceTokenSearchExpressionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PortCapacityExpression) GetOperatorOk() (*ServiceTokenSearchExpressionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PortCapacityExpression) SetOperator(v ServiceTokenSearchExpressionOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *PortCapacityExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *PortCapacityExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PortCapacityExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PortCapacityExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *PortCapacityExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


