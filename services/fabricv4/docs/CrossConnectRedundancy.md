# CrossConnectRedundancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Redundancy group identifier. When placing a diverse Cross Connect against an existing one, provide either this identifier or the serial number of the existing Cross Connect.  | [optional] 
**Priority** | Pointer to [**CrossConnectRedundancyPriority**](CrossConnectRedundancyPriority.md) |  | [optional] 

## Methods

### NewCrossConnectRedundancy

`func NewCrossConnectRedundancy() *CrossConnectRedundancy`

NewCrossConnectRedundancy instantiates a new CrossConnectRedundancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectRedundancyWithDefaults

`func NewCrossConnectRedundancyWithDefaults() *CrossConnectRedundancy`

NewCrossConnectRedundancyWithDefaults instantiates a new CrossConnectRedundancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *CrossConnectRedundancy) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CrossConnectRedundancy) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CrossConnectRedundancy) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CrossConnectRedundancy) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetPriority

`func (o *CrossConnectRedundancy) GetPriority() CrossConnectRedundancyPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CrossConnectRedundancy) GetPriorityOk() (*CrossConnectRedundancyPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CrossConnectRedundancy) SetPriority(v CrossConnectRedundancyPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CrossConnectRedundancy) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


