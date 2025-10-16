# DeviceBackupRestore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The unique ID of the backup. | [optional] 
**Name** | Pointer to **string** | The name of the backup. | [optional] 
**Status** | Pointer to **string** | The status of the backup. | [optional] 
**DeleteAllowed** | Pointer to **bool** | Whether you can delete the backup. Only backups in the COMPELTED or FAILED states can be deleted. | [optional] 

## Methods

### NewDeviceBackupRestore

`func NewDeviceBackupRestore() *DeviceBackupRestore`

NewDeviceBackupRestore instantiates a new DeviceBackupRestore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceBackupRestoreWithDefaults

`func NewDeviceBackupRestoreWithDefaults() *DeviceBackupRestore`

NewDeviceBackupRestoreWithDefaults instantiates a new DeviceBackupRestore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *DeviceBackupRestore) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeviceBackupRestore) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeviceBackupRestore) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeviceBackupRestore) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *DeviceBackupRestore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceBackupRestore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceBackupRestore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceBackupRestore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceBackupRestore) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceBackupRestore) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceBackupRestore) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceBackupRestore) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeleteAllowed

`func (o *DeviceBackupRestore) GetDeleteAllowed() bool`

GetDeleteAllowed returns the DeleteAllowed field if non-nil, zero value otherwise.

### GetDeleteAllowedOk

`func (o *DeviceBackupRestore) GetDeleteAllowedOk() (*bool, bool)`

GetDeleteAllowedOk returns a tuple with the DeleteAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAllowed

`func (o *DeviceBackupRestore) SetDeleteAllowed(v bool)`

SetDeleteAllowed sets DeleteAllowed field to given value.

### HasDeleteAllowed

`func (o *DeviceBackupRestore) HasDeleteAllowed() bool`

HasDeleteAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


