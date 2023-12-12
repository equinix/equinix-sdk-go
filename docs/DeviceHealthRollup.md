# DeviceHealthRollup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthRollup** | Pointer to [**DeviceHealthRollupHealthRollup**](DeviceHealthRollupHealthRollup.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Last update of health status. | [optional] [readonly] 

## Methods

### NewDeviceHealthRollup

`func NewDeviceHealthRollup() *DeviceHealthRollup`

NewDeviceHealthRollup instantiates a new DeviceHealthRollup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceHealthRollupWithDefaults

`func NewDeviceHealthRollupWithDefaults() *DeviceHealthRollup`

NewDeviceHealthRollupWithDefaults instantiates a new DeviceHealthRollup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthRollup

`func (o *DeviceHealthRollup) GetHealthRollup() DeviceHealthRollupHealthRollup`

GetHealthRollup returns the HealthRollup field if non-nil, zero value otherwise.

### GetHealthRollupOk

`func (o *DeviceHealthRollup) GetHealthRollupOk() (*DeviceHealthRollupHealthRollup, bool)`

GetHealthRollupOk returns a tuple with the HealthRollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRollup

`func (o *DeviceHealthRollup) SetHealthRollup(v DeviceHealthRollupHealthRollup)`

SetHealthRollup sets HealthRollup field to given value.

### HasHealthRollup

`func (o *DeviceHealthRollup) HasHealthRollup() bool`

HasHealthRollup returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DeviceHealthRollup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeviceHealthRollup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeviceHealthRollup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeviceHealthRollup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


