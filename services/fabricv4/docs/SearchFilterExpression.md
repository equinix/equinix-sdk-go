# SearchFilterExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | [**[]SearchFilterExpression**](SearchFilterExpression.md) |  | 
**Or** | [**[]SearchFilterExpression**](SearchFilterExpression.md) |  | 
**Property** | **string** |  | 
**Operator** | [**SearchSimpleExpressionOperator**](SearchSimpleExpressionOperator.md) |  | 
**Values** | **[]string** |  | 

## Methods

### NewSearchFilterExpression

`func NewSearchFilterExpression(and []SearchFilterExpression, or []SearchFilterExpression, property string, operator SearchSimpleExpressionOperator, values []string, ) *SearchFilterExpression`

NewSearchFilterExpression instantiates a new SearchFilterExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFilterExpressionWithDefaults

`func NewSearchFilterExpressionWithDefaults() *SearchFilterExpression`

NewSearchFilterExpressionWithDefaults instantiates a new SearchFilterExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *SearchFilterExpression) GetAnd() []SearchFilterExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *SearchFilterExpression) GetAndOk() (*[]SearchFilterExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *SearchFilterExpression) SetAnd(v []SearchFilterExpression)`

SetAnd sets And field to given value.


### GetOr

`func (o *SearchFilterExpression) GetOr() []SearchFilterExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *SearchFilterExpression) GetOrOk() (*[]SearchFilterExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *SearchFilterExpression) SetOr(v []SearchFilterExpression)`

SetOr sets Or field to given value.


### GetProperty

`func (o *SearchFilterExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SearchFilterExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SearchFilterExpression) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetOperator

`func (o *SearchFilterExpression) GetOperator() SearchSimpleExpressionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SearchFilterExpression) GetOperatorOk() (*SearchSimpleExpressionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SearchFilterExpression) SetOperator(v SearchSimpleExpressionOperator)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *SearchFilterExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SearchFilterExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SearchFilterExpression) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


