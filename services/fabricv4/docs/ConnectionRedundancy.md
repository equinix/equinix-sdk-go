# ConnectionRedundancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Redundancy group identifier (UUID of primary connection) | [optional] 
**Priority** | Pointer to [**ConnectionPriority**](ConnectionPriority.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided name for the redundancy group | [optional] 
**Description** | Pointer to **string** | Customer-provided description for the redundancy group | [optional] 

## Methods

### NewConnectionRedundancy

`func NewConnectionRedundancy() *ConnectionRedundancy`

NewConnectionRedundancy instantiates a new ConnectionRedundancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionRedundancyWithDefaults

`func NewConnectionRedundancyWithDefaults() *ConnectionRedundancy`

NewConnectionRedundancyWithDefaults instantiates a new ConnectionRedundancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ConnectionRedundancy) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ConnectionRedundancy) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ConnectionRedundancy) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ConnectionRedundancy) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetPriority

`func (o *ConnectionRedundancy) GetPriority() ConnectionPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ConnectionRedundancy) GetPriorityOk() (*ConnectionPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ConnectionRedundancy) SetPriority(v ConnectionPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ConnectionRedundancy) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetName

`func (o *ConnectionRedundancy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionRedundancy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionRedundancy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionRedundancy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectionRedundancy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectionRedundancy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectionRedundancy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectionRedundancy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


