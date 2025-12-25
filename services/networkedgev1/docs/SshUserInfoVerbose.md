# SshUserInfoVerbose

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The unique Id of the ssh user. | [optional] 
**Username** | Pointer to **string** | The user name of the ssh user. | [optional] 
**DeviceUuids** | Pointer to **[]string** | The devices associated with this ssh user. | [optional] 
**MetroStatusMap** | Pointer to [**map[string]MetroStatus**](MetroStatus.md) | Status and error messages corresponding to the metros where the user exists | [optional] 

## Methods

### NewSshUserInfoVerbose

`func NewSshUserInfoVerbose() *SshUserInfoVerbose`

NewSshUserInfoVerbose instantiates a new SshUserInfoVerbose object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserInfoVerboseWithDefaults

`func NewSshUserInfoVerboseWithDefaults() *SshUserInfoVerbose`

NewSshUserInfoVerboseWithDefaults instantiates a new SshUserInfoVerbose object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SshUserInfoVerbose) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SshUserInfoVerbose) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SshUserInfoVerbose) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SshUserInfoVerbose) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetUsername

`func (o *SshUserInfoVerbose) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SshUserInfoVerbose) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SshUserInfoVerbose) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SshUserInfoVerbose) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDeviceUuids

`func (o *SshUserInfoVerbose) GetDeviceUuids() []string`

GetDeviceUuids returns the DeviceUuids field if non-nil, zero value otherwise.

### GetDeviceUuidsOk

`func (o *SshUserInfoVerbose) GetDeviceUuidsOk() (*[]string, bool)`

GetDeviceUuidsOk returns a tuple with the DeviceUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuids

`func (o *SshUserInfoVerbose) SetDeviceUuids(v []string)`

SetDeviceUuids sets DeviceUuids field to given value.

### HasDeviceUuids

`func (o *SshUserInfoVerbose) HasDeviceUuids() bool`

HasDeviceUuids returns a boolean if a field has been set.

### GetMetroStatusMap

`func (o *SshUserInfoVerbose) GetMetroStatusMap() map[string]MetroStatus`

GetMetroStatusMap returns the MetroStatusMap field if non-nil, zero value otherwise.

### GetMetroStatusMapOk

`func (o *SshUserInfoVerbose) GetMetroStatusMapOk() (*map[string]MetroStatus, bool)`

GetMetroStatusMapOk returns a tuple with the MetroStatusMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroStatusMap

`func (o *SshUserInfoVerbose) SetMetroStatusMap(v map[string]MetroStatus)`

SetMetroStatusMap sets MetroStatusMap field to given value.

### HasMetroStatusMap

`func (o *SshUserInfoVerbose) HasMetroStatusMap() bool`

HasMetroStatusMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


