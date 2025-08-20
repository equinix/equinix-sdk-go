# EquipmentInstallRequestServiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceLocation** | **string** | Device Location | 
**ElevationDrawingAttached** | **bool** | Elevation drawing attached? | 
**InstallationPoint** | **string** | Installation Point | 
**InstalledEquipmentPhotoRequired** | **bool** | Installed Equipment Photo Required? | 
**MountHardwareIncluded** | **bool** | Mount hardware included? | 
**PatchDevices** | **bool** | Patch Devices? | 
**PatchingInfo** | Pointer to **string** | Patching info | [optional] 
**PowerItOn** | **bool** | Power it on? | 
**NeedSupportFromASubmarineCableStationEngineer** | Pointer to **bool** | This flag is only applicable to Submarine Cable IBXs | [optional] 
**ScopeOfWork** | **string** | Enter any additional details that will help our technicians execute your request. You may also attach your scope of work as a document if you exceed the character limit in this field. | 

## Methods

### NewEquipmentInstallRequestServiceDetails

`func NewEquipmentInstallRequestServiceDetails(deviceLocation string, elevationDrawingAttached bool, installationPoint string, installedEquipmentPhotoRequired bool, mountHardwareIncluded bool, patchDevices bool, powerItOn bool, scopeOfWork string, ) *EquipmentInstallRequestServiceDetails`

NewEquipmentInstallRequestServiceDetails instantiates a new EquipmentInstallRequestServiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentInstallRequestServiceDetailsWithDefaults

`func NewEquipmentInstallRequestServiceDetailsWithDefaults() *EquipmentInstallRequestServiceDetails`

NewEquipmentInstallRequestServiceDetailsWithDefaults instantiates a new EquipmentInstallRequestServiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceLocation

`func (o *EquipmentInstallRequestServiceDetails) GetDeviceLocation() string`

GetDeviceLocation returns the DeviceLocation field if non-nil, zero value otherwise.

### GetDeviceLocationOk

`func (o *EquipmentInstallRequestServiceDetails) GetDeviceLocationOk() (*string, bool)`

GetDeviceLocationOk returns a tuple with the DeviceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocation

`func (o *EquipmentInstallRequestServiceDetails) SetDeviceLocation(v string)`

SetDeviceLocation sets DeviceLocation field to given value.


### GetElevationDrawingAttached

`func (o *EquipmentInstallRequestServiceDetails) GetElevationDrawingAttached() bool`

GetElevationDrawingAttached returns the ElevationDrawingAttached field if non-nil, zero value otherwise.

### GetElevationDrawingAttachedOk

`func (o *EquipmentInstallRequestServiceDetails) GetElevationDrawingAttachedOk() (*bool, bool)`

GetElevationDrawingAttachedOk returns a tuple with the ElevationDrawingAttached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevationDrawingAttached

`func (o *EquipmentInstallRequestServiceDetails) SetElevationDrawingAttached(v bool)`

SetElevationDrawingAttached sets ElevationDrawingAttached field to given value.


### GetInstallationPoint

`func (o *EquipmentInstallRequestServiceDetails) GetInstallationPoint() string`

GetInstallationPoint returns the InstallationPoint field if non-nil, zero value otherwise.

### GetInstallationPointOk

`func (o *EquipmentInstallRequestServiceDetails) GetInstallationPointOk() (*string, bool)`

GetInstallationPointOk returns a tuple with the InstallationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationPoint

`func (o *EquipmentInstallRequestServiceDetails) SetInstallationPoint(v string)`

SetInstallationPoint sets InstallationPoint field to given value.


### GetInstalledEquipmentPhotoRequired

`func (o *EquipmentInstallRequestServiceDetails) GetInstalledEquipmentPhotoRequired() bool`

GetInstalledEquipmentPhotoRequired returns the InstalledEquipmentPhotoRequired field if non-nil, zero value otherwise.

### GetInstalledEquipmentPhotoRequiredOk

`func (o *EquipmentInstallRequestServiceDetails) GetInstalledEquipmentPhotoRequiredOk() (*bool, bool)`

GetInstalledEquipmentPhotoRequiredOk returns a tuple with the InstalledEquipmentPhotoRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledEquipmentPhotoRequired

`func (o *EquipmentInstallRequestServiceDetails) SetInstalledEquipmentPhotoRequired(v bool)`

SetInstalledEquipmentPhotoRequired sets InstalledEquipmentPhotoRequired field to given value.


### GetMountHardwareIncluded

`func (o *EquipmentInstallRequestServiceDetails) GetMountHardwareIncluded() bool`

GetMountHardwareIncluded returns the MountHardwareIncluded field if non-nil, zero value otherwise.

### GetMountHardwareIncludedOk

`func (o *EquipmentInstallRequestServiceDetails) GetMountHardwareIncludedOk() (*bool, bool)`

GetMountHardwareIncludedOk returns a tuple with the MountHardwareIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountHardwareIncluded

`func (o *EquipmentInstallRequestServiceDetails) SetMountHardwareIncluded(v bool)`

SetMountHardwareIncluded sets MountHardwareIncluded field to given value.


### GetPatchDevices

`func (o *EquipmentInstallRequestServiceDetails) GetPatchDevices() bool`

GetPatchDevices returns the PatchDevices field if non-nil, zero value otherwise.

### GetPatchDevicesOk

`func (o *EquipmentInstallRequestServiceDetails) GetPatchDevicesOk() (*bool, bool)`

GetPatchDevicesOk returns a tuple with the PatchDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchDevices

`func (o *EquipmentInstallRequestServiceDetails) SetPatchDevices(v bool)`

SetPatchDevices sets PatchDevices field to given value.


### GetPatchingInfo

`func (o *EquipmentInstallRequestServiceDetails) GetPatchingInfo() string`

GetPatchingInfo returns the PatchingInfo field if non-nil, zero value otherwise.

### GetPatchingInfoOk

`func (o *EquipmentInstallRequestServiceDetails) GetPatchingInfoOk() (*string, bool)`

GetPatchingInfoOk returns a tuple with the PatchingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchingInfo

`func (o *EquipmentInstallRequestServiceDetails) SetPatchingInfo(v string)`

SetPatchingInfo sets PatchingInfo field to given value.

### HasPatchingInfo

`func (o *EquipmentInstallRequestServiceDetails) HasPatchingInfo() bool`

HasPatchingInfo returns a boolean if a field has been set.

### GetPowerItOn

`func (o *EquipmentInstallRequestServiceDetails) GetPowerItOn() bool`

GetPowerItOn returns the PowerItOn field if non-nil, zero value otherwise.

### GetPowerItOnOk

`func (o *EquipmentInstallRequestServiceDetails) GetPowerItOnOk() (*bool, bool)`

GetPowerItOnOk returns a tuple with the PowerItOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerItOn

`func (o *EquipmentInstallRequestServiceDetails) SetPowerItOn(v bool)`

SetPowerItOn sets PowerItOn field to given value.


### GetNeedSupportFromASubmarineCableStationEngineer

`func (o *EquipmentInstallRequestServiceDetails) GetNeedSupportFromASubmarineCableStationEngineer() bool`

GetNeedSupportFromASubmarineCableStationEngineer returns the NeedSupportFromASubmarineCableStationEngineer field if non-nil, zero value otherwise.

### GetNeedSupportFromASubmarineCableStationEngineerOk

`func (o *EquipmentInstallRequestServiceDetails) GetNeedSupportFromASubmarineCableStationEngineerOk() (*bool, bool)`

GetNeedSupportFromASubmarineCableStationEngineerOk returns a tuple with the NeedSupportFromASubmarineCableStationEngineer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedSupportFromASubmarineCableStationEngineer

`func (o *EquipmentInstallRequestServiceDetails) SetNeedSupportFromASubmarineCableStationEngineer(v bool)`

SetNeedSupportFromASubmarineCableStationEngineer sets NeedSupportFromASubmarineCableStationEngineer field to given value.

### HasNeedSupportFromASubmarineCableStationEngineer

`func (o *EquipmentInstallRequestServiceDetails) HasNeedSupportFromASubmarineCableStationEngineer() bool`

HasNeedSupportFromASubmarineCableStationEngineer returns a boolean if a field has been set.

### GetScopeOfWork

`func (o *EquipmentInstallRequestServiceDetails) GetScopeOfWork() string`

GetScopeOfWork returns the ScopeOfWork field if non-nil, zero value otherwise.

### GetScopeOfWorkOk

`func (o *EquipmentInstallRequestServiceDetails) GetScopeOfWorkOk() (*string, bool)`

GetScopeOfWorkOk returns a tuple with the ScopeOfWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOfWork

`func (o *EquipmentInstallRequestServiceDetails) SetScopeOfWork(v string)`

SetScopeOfWork sets ScopeOfWork field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


