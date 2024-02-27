# Expression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]Expression**](Expression.md) |  | [optional] 
**Or** | Pointer to [**[]Expression**](Expression.md) |  | [optional] 
**Property** | Pointer to [**SearchFieldName**](SearchFieldName.md) |  | [optional] 
**Operator** | Pointer to [**ExpressionOperator**](ExpressionOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewExpression

`func NewExpression() *Expression`

NewExpression instantiates a new Expression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpressionWithDefaults

`func NewExpressionWithDefaults() *Expression`

NewExpressionWithDefaults instantiates a new Expression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *Expression) GetAnd() []Expression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *Expression) GetAndOk() (*[]Expression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *Expression) SetAnd(v []Expression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *Expression) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *Expression) GetOr() []Expression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *Expression) GetOrOk() (*[]Expression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *Expression) SetOr(v []Expression)`

SetOr sets Or field to given value.

### HasOr

`func (o *Expression) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetProperty

`func (o *Expression) GetProperty() SearchFieldName`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *Expression) GetPropertyOk() (*SearchFieldName, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *Expression) SetProperty(v SearchFieldName)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *Expression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *Expression) GetOperator() ExpressionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Expression) GetOperatorOk() (*ExpressionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Expression) SetOperator(v ExpressionOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Expression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *Expression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Expression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Expression) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *Expression) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


