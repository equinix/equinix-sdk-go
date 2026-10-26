# LiveEventFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]LiveEventSimpleExpression**](LiveEventSimpleExpression.md) |  | [optional] 

## Methods

### NewLiveEventFilters

`func NewLiveEventFilters() *LiveEventFilters`

NewLiveEventFilters instantiates a new LiveEventFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveEventFiltersWithDefaults

`func NewLiveEventFiltersWithDefaults() *LiveEventFilters`

NewLiveEventFiltersWithDefaults instantiates a new LiveEventFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *LiveEventFilters) GetAnd() []LiveEventSimpleExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *LiveEventFilters) GetAndOk() (*[]LiveEventSimpleExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *LiveEventFilters) SetAnd(v []LiveEventSimpleExpression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *LiveEventFilters) HasAnd() bool`

HasAnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


