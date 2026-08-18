# LoaFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to [**LoaFieldName**](LoaFieldName.md) |  | [optional] 
**Operator** | Pointer to [**LoaSimpleExpressionOperator**](LoaSimpleExpressionOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**Or** | Pointer to [**[]LoaSimpleExpression**](LoaSimpleExpression.md) |  | [optional] 

## Methods

### NewLoaFilter

`func NewLoaFilter() *LoaFilter`

NewLoaFilter instantiates a new LoaFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaFilterWithDefaults

`func NewLoaFilterWithDefaults() *LoaFilter`

NewLoaFilterWithDefaults instantiates a new LoaFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *LoaFilter) GetProperty() LoaFieldName`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *LoaFilter) GetPropertyOk() (*LoaFieldName, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *LoaFilter) SetProperty(v LoaFieldName)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *LoaFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *LoaFilter) GetOperator() LoaSimpleExpressionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LoaFilter) GetOperatorOk() (*LoaSimpleExpressionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LoaFilter) SetOperator(v LoaSimpleExpressionOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *LoaFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *LoaFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *LoaFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *LoaFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *LoaFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOr

`func (o *LoaFilter) GetOr() []LoaSimpleExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *LoaFilter) GetOrOk() (*[]LoaSimpleExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *LoaFilter) SetOr(v []LoaSimpleExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *LoaFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


