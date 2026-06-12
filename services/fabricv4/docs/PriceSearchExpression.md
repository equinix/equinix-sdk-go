# PriceSearchExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]PriceSearchExpression**](PriceSearchExpression.md) |  | [optional] 
**Or** | Pointer to [**[]PriceSearchExpression**](PriceSearchExpression.md) |  | [optional] 
**Property** | Pointer to **string** |  | [optional] 
**Operator** | Pointer to [**PriceSearchExpressionOperator**](PriceSearchExpressionOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPriceSearchExpression

`func NewPriceSearchExpression() *PriceSearchExpression`

NewPriceSearchExpression instantiates a new PriceSearchExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceSearchExpressionWithDefaults

`func NewPriceSearchExpressionWithDefaults() *PriceSearchExpression`

NewPriceSearchExpressionWithDefaults instantiates a new PriceSearchExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *PriceSearchExpression) GetAnd() []PriceSearchExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *PriceSearchExpression) GetAndOk() (*[]PriceSearchExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *PriceSearchExpression) SetAnd(v []PriceSearchExpression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *PriceSearchExpression) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *PriceSearchExpression) GetOr() []PriceSearchExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *PriceSearchExpression) GetOrOk() (*[]PriceSearchExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *PriceSearchExpression) SetOr(v []PriceSearchExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *PriceSearchExpression) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetProperty

`func (o *PriceSearchExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PriceSearchExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PriceSearchExpression) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *PriceSearchExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *PriceSearchExpression) GetOperator() PriceSearchExpressionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PriceSearchExpression) GetOperatorOk() (*PriceSearchExpressionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PriceSearchExpression) SetOperator(v PriceSearchExpressionOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *PriceSearchExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *PriceSearchExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PriceSearchExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PriceSearchExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *PriceSearchExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


