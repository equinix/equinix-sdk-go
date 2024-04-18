# DeviceBackupInfoVerbose

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The unique Id of the device backup. | [optional] 
**Name** | Pointer to **string** | The name of the backup. | [optional] 
**Version** | Pointer to **string** | Version of the device | [optional] 
**Type** | Pointer to **string** | The type of backup. | [optional] 
**Status** | Pointer to **string** | The status of the backup. | [optional] 
**CreatedBy** | Pointer to **string** | Created by. | [optional] 
**LastUpdatedDate** | Pointer to **string** | Last updated date. | [optional] 
**DownloadUrl** | Pointer to **string** | URL where you can download the backup. | [optional] 
**DeleteAllowed** | Pointer to **bool** | Whether or not you can delete the backup. | [optional] 
**Restores** | Pointer to [**[]PreviousBackups**](PreviousBackups.md) |  | [optional] 
**DeviceUuid** | Pointer to **string** | Unique Id of the device | [optional] 

## Methods

### NewDeviceBackupInfoVerbose

`func NewDeviceBackupInfoVerbose() *DeviceBackupInfoVerbose`

NewDeviceBackupInfoVerbose instantiates a new DeviceBackupInfoVerbose object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceBackupInfoVerboseWithDefaults

`func NewDeviceBackupInfoVerboseWithDefaults() *DeviceBackupInfoVerbose`

NewDeviceBackupInfoVerboseWithDefaults instantiates a new DeviceBackupInfoVerbose object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *DeviceBackupInfoVerbose) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeviceBackupInfoVerbose) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeviceBackupInfoVerbose) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeviceBackupInfoVerbose) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *DeviceBackupInfoVerbose) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceBackupInfoVerbose) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceBackupInfoVerbose) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceBackupInfoVerbose) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *DeviceBackupInfoVerbose) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceBackupInfoVerbose) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceBackupInfoVerbose) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeviceBackupInfoVerbose) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *DeviceBackupInfoVerbose) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceBackupInfoVerbose) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceBackupInfoVerbose) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceBackupInfoVerbose) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceBackupInfoVerbose) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceBackupInfoVerbose) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceBackupInfoVerbose) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceBackupInfoVerbose) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DeviceBackupInfoVerbose) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DeviceBackupInfoVerbose) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DeviceBackupInfoVerbose) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DeviceBackupInfoVerbose) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *DeviceBackupInfoVerbose) GetLastUpdatedDate() string`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *DeviceBackupInfoVerbose) GetLastUpdatedDateOk() (*string, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *DeviceBackupInfoVerbose) SetLastUpdatedDate(v string)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *DeviceBackupInfoVerbose) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *DeviceBackupInfoVerbose) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *DeviceBackupInfoVerbose) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *DeviceBackupInfoVerbose) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *DeviceBackupInfoVerbose) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetDeleteAllowed

`func (o *DeviceBackupInfoVerbose) GetDeleteAllowed() bool`

GetDeleteAllowed returns the DeleteAllowed field if non-nil, zero value otherwise.

### GetDeleteAllowedOk

`func (o *DeviceBackupInfoVerbose) GetDeleteAllowedOk() (*bool, bool)`

GetDeleteAllowedOk returns a tuple with the DeleteAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAllowed

`func (o *DeviceBackupInfoVerbose) SetDeleteAllowed(v bool)`

SetDeleteAllowed sets DeleteAllowed field to given value.

### HasDeleteAllowed

`func (o *DeviceBackupInfoVerbose) HasDeleteAllowed() bool`

HasDeleteAllowed returns a boolean if a field has been set.

### GetRestores

`func (o *DeviceBackupInfoVerbose) GetRestores() []PreviousBackups`

GetRestores returns the Restores field if non-nil, zero value otherwise.

### GetRestoresOk

`func (o *DeviceBackupInfoVerbose) GetRestoresOk() (*[]PreviousBackups, bool)`

GetRestoresOk returns a tuple with the Restores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestores

`func (o *DeviceBackupInfoVerbose) SetRestores(v []PreviousBackups)`

SetRestores sets Restores field to given value.

### HasRestores

`func (o *DeviceBackupInfoVerbose) HasRestores() bool`

HasRestores returns a boolean if a field has been set.

### GetDeviceUuid

`func (o *DeviceBackupInfoVerbose) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *DeviceBackupInfoVerbose) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *DeviceBackupInfoVerbose) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.

### HasDeviceUuid

`func (o *DeviceBackupInfoVerbose) HasDeviceUuid() bool`

HasDeviceUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


