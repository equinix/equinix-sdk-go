# SearchOrExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Or** | [**[]SearchFilterExpression**](SearchFilterExpression.md) |  | 

## Methods

### NewSearchOrExpression

`func NewSearchOrExpression(or []SearchFilterExpression, ) *SearchOrExpression`

NewSearchOrExpression instantiates a new SearchOrExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchOrExpressionWithDefaults

`func NewSearchOrExpressionWithDefaults() *SearchOrExpression`

NewSearchOrExpressionWithDefaults instantiates a new SearchOrExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOr

`func (o *SearchOrExpression) GetOr() []SearchFilterExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *SearchOrExpression) GetOrOk() (*[]SearchFilterExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *SearchOrExpression) SetOr(v []SearchFilterExpression)`

SetOr sets Or field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


