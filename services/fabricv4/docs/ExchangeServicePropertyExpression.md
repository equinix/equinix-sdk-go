# ExchangeServicePropertyExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**ExchangeServicePropertyExpressionProperty**](ExchangeServicePropertyExpressionProperty.md) |  | 
**Operator** | [**ExchangeServicePropertyExpressionOperator**](ExchangeServicePropertyExpressionOperator.md) |  | 
**Values** | **[]string** |  | 

## Methods

### NewExchangeServicePropertyExpression

`func NewExchangeServicePropertyExpression(property ExchangeServicePropertyExpressionProperty, operator ExchangeServicePropertyExpressionOperator, values []string, ) *ExchangeServicePropertyExpression`

NewExchangeServicePropertyExpression instantiates a new ExchangeServicePropertyExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeServicePropertyExpressionWithDefaults

`func NewExchangeServicePropertyExpressionWithDefaults() *ExchangeServicePropertyExpression`

NewExchangeServicePropertyExpressionWithDefaults instantiates a new ExchangeServicePropertyExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *ExchangeServicePropertyExpression) GetProperty() ExchangeServicePropertyExpressionProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ExchangeServicePropertyExpression) GetPropertyOk() (*ExchangeServicePropertyExpressionProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ExchangeServicePropertyExpression) SetProperty(v ExchangeServicePropertyExpressionProperty)`

SetProperty sets Property field to given value.


### GetOperator

`func (o *ExchangeServicePropertyExpression) GetOperator() ExchangeServicePropertyExpressionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ExchangeServicePropertyExpression) GetOperatorOk() (*ExchangeServicePropertyExpressionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ExchangeServicePropertyExpression) SetOperator(v ExchangeServicePropertyExpressionOperator)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *ExchangeServicePropertyExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ExchangeServicePropertyExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ExchangeServicePropertyExpression) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


