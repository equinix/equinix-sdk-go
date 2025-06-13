# ProviderExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]ProviderSearchExpressions**](ProviderSearchExpressions.md) |  | [optional] 
**Or** | Pointer to [**[]ProviderSearchExpressions**](ProviderSearchExpressions.md) |  | [optional] 

## Methods

### NewProviderExpression

`func NewProviderExpression() *ProviderExpression`

NewProviderExpression instantiates a new ProviderExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderExpressionWithDefaults

`func NewProviderExpressionWithDefaults() *ProviderExpression`

NewProviderExpressionWithDefaults instantiates a new ProviderExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *ProviderExpression) GetAnd() []ProviderSearchExpressions`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *ProviderExpression) GetAndOk() (*[]ProviderSearchExpressions, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *ProviderExpression) SetAnd(v []ProviderSearchExpressions)`

SetAnd sets And field to given value.

### HasAnd

`func (o *ProviderExpression) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *ProviderExpression) GetOr() []ProviderSearchExpressions`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *ProviderExpression) GetOrOk() (*[]ProviderSearchExpressions, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *ProviderExpression) SetOr(v []ProviderSearchExpressions)`

SetOr sets Or field to given value.

### HasOr

`func (o *ProviderExpression) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


