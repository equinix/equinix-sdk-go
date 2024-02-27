# SearchExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]SearchExpression**](SearchExpression.md) |  | [optional] 
**Or** | Pointer to [**[]SearchExpression**](SearchExpression.md) |  | [optional] 
**Property** | Pointer to **string** |  | [optional] 
**Operator** | Pointer to [**ExpressionOperator**](ExpressionOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSearchExpression

`func NewSearchExpression() *SearchExpression`

NewSearchExpression instantiates a new SearchExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchExpressionWithDefaults

`func NewSearchExpressionWithDefaults() *SearchExpression`

NewSearchExpressionWithDefaults instantiates a new SearchExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *SearchExpression) GetAnd() []SearchExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *SearchExpression) GetAndOk() (*[]SearchExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *SearchExpression) SetAnd(v []SearchExpression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *SearchExpression) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *SearchExpression) GetOr() []SearchExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *SearchExpression) GetOrOk() (*[]SearchExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *SearchExpression) SetOr(v []SearchExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *SearchExpression) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetProperty

`func (o *SearchExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SearchExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SearchExpression) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *SearchExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *SearchExpression) GetOperator() ExpressionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SearchExpression) GetOperatorOk() (*ExpressionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SearchExpression) SetOperator(v ExpressionOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *SearchExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *SearchExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SearchExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SearchExpression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *SearchExpression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


