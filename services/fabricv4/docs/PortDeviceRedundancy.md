# PortDeviceRedundancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Device redundancy group | [optional] 
**Priority** | Pointer to [**PortDeviceRedundancyPriority**](PortDeviceRedundancyPriority.md) |  | [optional] 

## Methods

### NewPortDeviceRedundancy

`func NewPortDeviceRedundancy() *PortDeviceRedundancy`

NewPortDeviceRedundancy instantiates a new PortDeviceRedundancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortDeviceRedundancyWithDefaults

`func NewPortDeviceRedundancyWithDefaults() *PortDeviceRedundancy`

NewPortDeviceRedundancyWithDefaults instantiates a new PortDeviceRedundancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *PortDeviceRedundancy) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PortDeviceRedundancy) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PortDeviceRedundancy) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PortDeviceRedundancy) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetPriority

`func (o *PortDeviceRedundancy) GetPriority() PortDeviceRedundancyPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PortDeviceRedundancy) GetPriorityOk() (*PortDeviceRedundancyPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PortDeviceRedundancy) SetPriority(v PortDeviceRedundancyPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PortDeviceRedundancy) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


