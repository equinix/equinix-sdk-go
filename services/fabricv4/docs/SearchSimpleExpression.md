# SearchSimpleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | **string** |  | 
**Operator** | [**SearchSimpleExpressionOperator**](SearchSimpleExpressionOperator.md) |  | 
**Values** | **[]string** |  | 

## Methods

### NewSearchSimpleExpression

`func NewSearchSimpleExpression(property string, operator SearchSimpleExpressionOperator, values []string, ) *SearchSimpleExpression`

NewSearchSimpleExpression instantiates a new SearchSimpleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSimpleExpressionWithDefaults

`func NewSearchSimpleExpressionWithDefaults() *SearchSimpleExpression`

NewSearchSimpleExpressionWithDefaults instantiates a new SearchSimpleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *SearchSimpleExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SearchSimpleExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SearchSimpleExpression) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetOperator

`func (o *SearchSimpleExpression) GetOperator() SearchSimpleExpressionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SearchSimpleExpression) GetOperatorOk() (*SearchSimpleExpressionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SearchSimpleExpression) SetOperator(v SearchSimpleExpressionOperator)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *SearchSimpleExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SearchSimpleExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SearchSimpleExpression) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


