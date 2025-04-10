# CloudEventFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]CloudEventSimpleExpression**](CloudEventSimpleExpression.md) |  | [optional] 

## Methods

### NewCloudEventFilters

`func NewCloudEventFilters() *CloudEventFilters`

NewCloudEventFilters instantiates a new CloudEventFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEventFiltersWithDefaults

`func NewCloudEventFiltersWithDefaults() *CloudEventFilters`

NewCloudEventFiltersWithDefaults instantiates a new CloudEventFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *CloudEventFilters) GetAnd() []CloudEventSimpleExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *CloudEventFilters) GetAndOk() (*[]CloudEventSimpleExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *CloudEventFilters) SetAnd(v []CloudEventSimpleExpression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *CloudEventFilters) HasAnd() bool`

HasAnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


