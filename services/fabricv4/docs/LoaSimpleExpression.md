# LoaSimpleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to [**LoaFieldName**](LoaFieldName.md) |  | [optional] 
**Operator** | Pointer to [**LoaSimpleExpressionOperator**](LoaSimpleExpressionOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewLoaSimpleExpression

`func NewLoaSimpleExpression() *LoaSimpleExpression`

NewLoaSimpleExpression instantiates a new LoaSimpleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaSimpleExpressionWithDefaults

`func NewLoaSimpleExpressionWithDefaults() *LoaSimpleExpression`

NewLoaSimpleExpressionWithDefaults instantiates a new LoaSimpleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *LoaSimpleExpression) GetProperty() LoaFieldName`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *LoaSimpleExpression) GetPropertyOk() (*LoaFieldName, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *LoaSimpleExpression) SetProperty(v LoaFieldName)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *LoaSimpleExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *LoaSimpleExpression) GetOperator() LoaSimpleExpressionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LoaSimpleExpression) GetOperatorOk() (*LoaSimpleExpressionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LoaSimpleExpression) SetOperator(v LoaSimpleExpressionOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *LoaSimpleExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *LoaSimpleExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *LoaSimpleExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *LoaSimpleExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *LoaSimpleExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


