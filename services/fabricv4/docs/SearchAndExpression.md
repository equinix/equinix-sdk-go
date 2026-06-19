# SearchAndExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | [**[]SearchFilterExpression**](SearchFilterExpression.md) |  | 

## Methods

### NewSearchAndExpression

`func NewSearchAndExpression(and []SearchFilterExpression, ) *SearchAndExpression`

NewSearchAndExpression instantiates a new SearchAndExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAndExpressionWithDefaults

`func NewSearchAndExpressionWithDefaults() *SearchAndExpression`

NewSearchAndExpressionWithDefaults instantiates a new SearchAndExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *SearchAndExpression) GetAnd() []SearchFilterExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *SearchAndExpression) GetAndOk() (*[]SearchFilterExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *SearchAndExpression) SetAnd(v []SearchFilterExpression)`

SetAnd sets And field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


