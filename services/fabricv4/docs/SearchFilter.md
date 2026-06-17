# SearchFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | [**[]SearchFilterExpression**](SearchFilterExpression.md) |  | 
**Or** | [**[]SearchFilterExpression**](SearchFilterExpression.md) |  | 

## Methods

### NewSearchFilter

`func NewSearchFilter(and []SearchFilterExpression, or []SearchFilterExpression, ) *SearchFilter`

NewSearchFilter instantiates a new SearchFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFilterWithDefaults

`func NewSearchFilterWithDefaults() *SearchFilter`

NewSearchFilterWithDefaults instantiates a new SearchFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *SearchFilter) GetAnd() []SearchFilterExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *SearchFilter) GetAndOk() (*[]SearchFilterExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *SearchFilter) SetAnd(v []SearchFilterExpression)`

SetAnd sets And field to given value.


### GetOr

`func (o *SearchFilter) GetOr() []SearchFilterExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *SearchFilter) GetOrOk() (*[]SearchFilterExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *SearchFilter) SetOr(v []SearchFilterExpression)`

SetOr sets Or field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


