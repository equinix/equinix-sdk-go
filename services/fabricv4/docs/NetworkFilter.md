# NetworkFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]NetworkFilter**](NetworkFilter.md) |  | [optional] 
**Or** | Pointer to [**[]NetworkFilter**](NetworkFilter.md) |  | [optional] 
**Property** | Pointer to [**NetworkSearchFieldName**](NetworkSearchFieldName.md) |  | [optional] 
**Operator** | Pointer to [**NetworkFilterOperator**](NetworkFilterOperator.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNetworkFilter

`func NewNetworkFilter() *NetworkFilter`

NewNetworkFilter instantiates a new NetworkFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFilterWithDefaults

`func NewNetworkFilterWithDefaults() *NetworkFilter`

NewNetworkFilterWithDefaults instantiates a new NetworkFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *NetworkFilter) GetAnd() []NetworkFilter`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *NetworkFilter) GetAndOk() (*[]NetworkFilter, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *NetworkFilter) SetAnd(v []NetworkFilter)`

SetAnd sets And field to given value.

### HasAnd

`func (o *NetworkFilter) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *NetworkFilter) GetOr() []NetworkFilter`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *NetworkFilter) GetOrOk() (*[]NetworkFilter, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *NetworkFilter) SetOr(v []NetworkFilter)`

SetOr sets Or field to given value.

### HasOr

`func (o *NetworkFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetProperty

`func (o *NetworkFilter) GetProperty() NetworkSearchFieldName`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *NetworkFilter) GetPropertyOk() (*NetworkSearchFieldName, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *NetworkFilter) SetProperty(v NetworkSearchFieldName)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *NetworkFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *NetworkFilter) GetOperator() NetworkFilterOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *NetworkFilter) GetOperatorOk() (*NetworkFilterOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *NetworkFilter) SetOperator(v NetworkFilterOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *NetworkFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *NetworkFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *NetworkFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *NetworkFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *NetworkFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


