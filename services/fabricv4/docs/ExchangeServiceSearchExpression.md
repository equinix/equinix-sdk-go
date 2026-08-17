# ExchangeServiceSearchExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | [**[]ExchangeServiceSearchExpression**](ExchangeServiceSearchExpression.md) |  | 
**Or** | [**[]ExchangeServiceSearchExpression**](ExchangeServiceSearchExpression.md) |  | 
**Property** | [**ExchangeServicePropertyExpressionProperty**](ExchangeServicePropertyExpressionProperty.md) |  | 
**Operator** | [**ExchangeServicePropertyExpressionOperator**](ExchangeServicePropertyExpressionOperator.md) |  | 
**Values** | **[]string** |  | 

## Methods

### NewExchangeServiceSearchExpression

`func NewExchangeServiceSearchExpression(and []ExchangeServiceSearchExpression, or []ExchangeServiceSearchExpression, property ExchangeServicePropertyExpressionProperty, operator ExchangeServicePropertyExpressionOperator, values []string, ) *ExchangeServiceSearchExpression`

NewExchangeServiceSearchExpression instantiates a new ExchangeServiceSearchExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeServiceSearchExpressionWithDefaults

`func NewExchangeServiceSearchExpressionWithDefaults() *ExchangeServiceSearchExpression`

NewExchangeServiceSearchExpressionWithDefaults instantiates a new ExchangeServiceSearchExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *ExchangeServiceSearchExpression) GetAnd() []ExchangeServiceSearchExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *ExchangeServiceSearchExpression) GetAndOk() (*[]ExchangeServiceSearchExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *ExchangeServiceSearchExpression) SetAnd(v []ExchangeServiceSearchExpression)`

SetAnd sets And field to given value.


### GetOr

`func (o *ExchangeServiceSearchExpression) GetOr() []ExchangeServiceSearchExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *ExchangeServiceSearchExpression) GetOrOk() (*[]ExchangeServiceSearchExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *ExchangeServiceSearchExpression) SetOr(v []ExchangeServiceSearchExpression)`

SetOr sets Or field to given value.


### GetProperty

`func (o *ExchangeServiceSearchExpression) GetProperty() ExchangeServicePropertyExpressionProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ExchangeServiceSearchExpression) GetPropertyOk() (*ExchangeServicePropertyExpressionProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ExchangeServiceSearchExpression) SetProperty(v ExchangeServicePropertyExpressionProperty)`

SetProperty sets Property field to given value.


### GetOperator

`func (o *ExchangeServiceSearchExpression) GetOperator() ExchangeServicePropertyExpressionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ExchangeServiceSearchExpression) GetOperatorOk() (*ExchangeServicePropertyExpressionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ExchangeServiceSearchExpression) SetOperator(v ExchangeServicePropertyExpressionOperator)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *ExchangeServiceSearchExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ExchangeServiceSearchExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ExchangeServiceSearchExpression) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


