# InterconnectFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]InterconnectFilter**](InterconnectFilter.md) |  | [optional] 
**Or** | Pointer to [**[]InterconnectFilter**](InterconnectFilter.md) |  | [optional] 
**Property** | Pointer to [**InterconnectSearchFieldName**](InterconnectSearchFieldName.md) |  | [optional] 
**Operator** | Pointer to [**InterconnectFilterOperator**](InterconnectFilterOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInterconnectFilter

`func NewInterconnectFilter() *InterconnectFilter`

NewInterconnectFilter instantiates a new InterconnectFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectFilterWithDefaults

`func NewInterconnectFilterWithDefaults() *InterconnectFilter`

NewInterconnectFilterWithDefaults instantiates a new InterconnectFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *InterconnectFilter) GetAnd() []InterconnectFilter`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *InterconnectFilter) GetAndOk() (*[]InterconnectFilter, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *InterconnectFilter) SetAnd(v []InterconnectFilter)`

SetAnd sets And field to given value.

### HasAnd

`func (o *InterconnectFilter) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *InterconnectFilter) GetOr() []InterconnectFilter`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *InterconnectFilter) GetOrOk() (*[]InterconnectFilter, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *InterconnectFilter) SetOr(v []InterconnectFilter)`

SetOr sets Or field to given value.

### HasOr

`func (o *InterconnectFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetProperty

`func (o *InterconnectFilter) GetProperty() InterconnectSearchFieldName`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *InterconnectFilter) GetPropertyOk() (*InterconnectSearchFieldName, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *InterconnectFilter) SetProperty(v InterconnectSearchFieldName)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *InterconnectFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *InterconnectFilter) GetOperator() InterconnectFilterOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *InterconnectFilter) GetOperatorOk() (*InterconnectFilterOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *InterconnectFilter) SetOperator(v InterconnectFilterOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *InterconnectFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *InterconnectFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InterconnectFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InterconnectFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *InterconnectFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


