# OpticalConnectRedundancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to [**OpticalConnectRedundancyPriority**](OpticalConnectRedundancyPriority.md) |  | [optional] 
**Group** | Pointer to **string** | Redundancy group identifier | [optional] 

## Methods

### NewOpticalConnectRedundancy

`func NewOpticalConnectRedundancy() *OpticalConnectRedundancy`

NewOpticalConnectRedundancy instantiates a new OpticalConnectRedundancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpticalConnectRedundancyWithDefaults

`func NewOpticalConnectRedundancyWithDefaults() *OpticalConnectRedundancy`

NewOpticalConnectRedundancyWithDefaults instantiates a new OpticalConnectRedundancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *OpticalConnectRedundancy) GetPriority() OpticalConnectRedundancyPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OpticalConnectRedundancy) GetPriorityOk() (*OpticalConnectRedundancyPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OpticalConnectRedundancy) SetPriority(v OpticalConnectRedundancyPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *OpticalConnectRedundancy) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetGroup

`func (o *OpticalConnectRedundancy) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *OpticalConnectRedundancy) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *OpticalConnectRedundancy) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *OpticalConnectRedundancy) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


