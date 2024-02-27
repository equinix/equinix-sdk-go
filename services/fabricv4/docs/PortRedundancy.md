# PortRedundancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Access point redundancy | [optional] 
**Group** | Pointer to **string** | Port UUID of respective primary port | [optional] 
**Priority** | Pointer to [**PortPriority**](PortPriority.md) |  | [optional] 

## Methods

### NewPortRedundancy

`func NewPortRedundancy() *PortRedundancy`

NewPortRedundancy instantiates a new PortRedundancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortRedundancyWithDefaults

`func NewPortRedundancyWithDefaults() *PortRedundancy`

NewPortRedundancyWithDefaults instantiates a new PortRedundancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PortRedundancy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PortRedundancy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PortRedundancy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PortRedundancy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroup

`func (o *PortRedundancy) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PortRedundancy) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PortRedundancy) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PortRedundancy) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetPriority

`func (o *PortRedundancy) GetPriority() PortPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PortRedundancy) GetPriorityOk() (*PortPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PortRedundancy) SetPriority(v PortPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PortRedundancy) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


