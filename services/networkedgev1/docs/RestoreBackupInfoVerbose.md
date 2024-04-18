# RestoreBackupInfoVerbose

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceBackup** | Pointer to [**RestoreBackupInfoVerboseDeviceBackup**](RestoreBackupInfoVerboseDeviceBackup.md) |  | [optional] 
**Services** | Pointer to [**RestoreBackupInfoVerboseServices**](RestoreBackupInfoVerboseServices.md) |  | [optional] 
**RestoreAllowedAfterDeleteOrEdit** | Pointer to **bool** | If True, the backup is restorable once you perform the recommended opertions. If False, the backup is not restorable. | [optional] 

## Methods

### NewRestoreBackupInfoVerbose

`func NewRestoreBackupInfoVerbose() *RestoreBackupInfoVerbose`

NewRestoreBackupInfoVerbose instantiates a new RestoreBackupInfoVerbose object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreBackupInfoVerboseWithDefaults

`func NewRestoreBackupInfoVerboseWithDefaults() *RestoreBackupInfoVerbose`

NewRestoreBackupInfoVerboseWithDefaults instantiates a new RestoreBackupInfoVerbose object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceBackup

`func (o *RestoreBackupInfoVerbose) GetDeviceBackup() RestoreBackupInfoVerboseDeviceBackup`

GetDeviceBackup returns the DeviceBackup field if non-nil, zero value otherwise.

### GetDeviceBackupOk

`func (o *RestoreBackupInfoVerbose) GetDeviceBackupOk() (*RestoreBackupInfoVerboseDeviceBackup, bool)`

GetDeviceBackupOk returns a tuple with the DeviceBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceBackup

`func (o *RestoreBackupInfoVerbose) SetDeviceBackup(v RestoreBackupInfoVerboseDeviceBackup)`

SetDeviceBackup sets DeviceBackup field to given value.

### HasDeviceBackup

`func (o *RestoreBackupInfoVerbose) HasDeviceBackup() bool`

HasDeviceBackup returns a boolean if a field has been set.

### GetServices

`func (o *RestoreBackupInfoVerbose) GetServices() RestoreBackupInfoVerboseServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *RestoreBackupInfoVerbose) GetServicesOk() (*RestoreBackupInfoVerboseServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *RestoreBackupInfoVerbose) SetServices(v RestoreBackupInfoVerboseServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *RestoreBackupInfoVerbose) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetRestoreAllowedAfterDeleteOrEdit

`func (o *RestoreBackupInfoVerbose) GetRestoreAllowedAfterDeleteOrEdit() bool`

GetRestoreAllowedAfterDeleteOrEdit returns the RestoreAllowedAfterDeleteOrEdit field if non-nil, zero value otherwise.

### GetRestoreAllowedAfterDeleteOrEditOk

`func (o *RestoreBackupInfoVerbose) GetRestoreAllowedAfterDeleteOrEditOk() (*bool, bool)`

GetRestoreAllowedAfterDeleteOrEditOk returns a tuple with the RestoreAllowedAfterDeleteOrEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAllowedAfterDeleteOrEdit

`func (o *RestoreBackupInfoVerbose) SetRestoreAllowedAfterDeleteOrEdit(v bool)`

SetRestoreAllowedAfterDeleteOrEdit sets RestoreAllowedAfterDeleteOrEdit field to given value.

### HasRestoreAllowedAfterDeleteOrEdit

`func (o *RestoreBackupInfoVerbose) HasRestoreAllowedAfterDeleteOrEdit() bool`

HasRestoreAllowedAfterDeleteOrEdit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


