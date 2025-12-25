# VirtualDeviceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceTypeCode** | Pointer to **string** | The type of the device. | [optional] 
**Name** | Pointer to **string** | The name of the device. | [optional] 
**Description** | Pointer to **string** | The description of the device. | [optional] 
**Vendor** | Pointer to **string** | The vendor of the device. | [optional] 
**Category** | Pointer to **string** | The type of the virtual device, whether router or firewall. | [optional] 
**MaxInterfaceCount** | Pointer to **int32** | The maximum available number of interfaces. | [optional] 
**DefaultInterfaceCount** | Pointer to **int32** | The default number of interfaces. | [optional] 
**ClusterMaxInterfaceCount** | Pointer to **int32** | The maximum number of available interfaces in case you are clustering. | [optional] 
**ClusterDefaultInterfaceCount** | Pointer to **int32** | The default number of available interfaces in case you are clustering. | [optional] 
**AvailableMetros** | Pointer to [**[]Metro**](Metro.md) | An array of metros where the device is available. | [optional] 
**SoftwarePackages** | Pointer to [**[]SoftwarePackage**](SoftwarePackage.md) | An array of available software packages | [optional] 
**DeviceManagementTypes** | Pointer to [**VirtualDeviceTypeDeviceManagementTypes**](VirtualDeviceTypeDeviceManagementTypes.md) |  | [optional] 

## Methods

### NewVirtualDeviceType

`func NewVirtualDeviceType() *VirtualDeviceType`

NewVirtualDeviceType instantiates a new VirtualDeviceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDeviceTypeWithDefaults

`func NewVirtualDeviceTypeWithDefaults() *VirtualDeviceType`

NewVirtualDeviceTypeWithDefaults instantiates a new VirtualDeviceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceTypeCode

`func (o *VirtualDeviceType) GetDeviceTypeCode() string`

GetDeviceTypeCode returns the DeviceTypeCode field if non-nil, zero value otherwise.

### GetDeviceTypeCodeOk

`func (o *VirtualDeviceType) GetDeviceTypeCodeOk() (*string, bool)`

GetDeviceTypeCodeOk returns a tuple with the DeviceTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeCode

`func (o *VirtualDeviceType) SetDeviceTypeCode(v string)`

SetDeviceTypeCode sets DeviceTypeCode field to given value.

### HasDeviceTypeCode

`func (o *VirtualDeviceType) HasDeviceTypeCode() bool`

HasDeviceTypeCode returns a boolean if a field has been set.

### GetName

`func (o *VirtualDeviceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualDeviceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualDeviceType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualDeviceType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualDeviceType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualDeviceType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualDeviceType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualDeviceType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVendor

`func (o *VirtualDeviceType) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VirtualDeviceType) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VirtualDeviceType) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *VirtualDeviceType) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetCategory

`func (o *VirtualDeviceType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *VirtualDeviceType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *VirtualDeviceType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *VirtualDeviceType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetMaxInterfaceCount

`func (o *VirtualDeviceType) GetMaxInterfaceCount() int32`

GetMaxInterfaceCount returns the MaxInterfaceCount field if non-nil, zero value otherwise.

### GetMaxInterfaceCountOk

`func (o *VirtualDeviceType) GetMaxInterfaceCountOk() (*int32, bool)`

GetMaxInterfaceCountOk returns a tuple with the MaxInterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInterfaceCount

`func (o *VirtualDeviceType) SetMaxInterfaceCount(v int32)`

SetMaxInterfaceCount sets MaxInterfaceCount field to given value.

### HasMaxInterfaceCount

`func (o *VirtualDeviceType) HasMaxInterfaceCount() bool`

HasMaxInterfaceCount returns a boolean if a field has been set.

### GetDefaultInterfaceCount

`func (o *VirtualDeviceType) GetDefaultInterfaceCount() int32`

GetDefaultInterfaceCount returns the DefaultInterfaceCount field if non-nil, zero value otherwise.

### GetDefaultInterfaceCountOk

`func (o *VirtualDeviceType) GetDefaultInterfaceCountOk() (*int32, bool)`

GetDefaultInterfaceCountOk returns a tuple with the DefaultInterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInterfaceCount

`func (o *VirtualDeviceType) SetDefaultInterfaceCount(v int32)`

SetDefaultInterfaceCount sets DefaultInterfaceCount field to given value.

### HasDefaultInterfaceCount

`func (o *VirtualDeviceType) HasDefaultInterfaceCount() bool`

HasDefaultInterfaceCount returns a boolean if a field has been set.

### GetClusterMaxInterfaceCount

`func (o *VirtualDeviceType) GetClusterMaxInterfaceCount() int32`

GetClusterMaxInterfaceCount returns the ClusterMaxInterfaceCount field if non-nil, zero value otherwise.

### GetClusterMaxInterfaceCountOk

`func (o *VirtualDeviceType) GetClusterMaxInterfaceCountOk() (*int32, bool)`

GetClusterMaxInterfaceCountOk returns a tuple with the ClusterMaxInterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMaxInterfaceCount

`func (o *VirtualDeviceType) SetClusterMaxInterfaceCount(v int32)`

SetClusterMaxInterfaceCount sets ClusterMaxInterfaceCount field to given value.

### HasClusterMaxInterfaceCount

`func (o *VirtualDeviceType) HasClusterMaxInterfaceCount() bool`

HasClusterMaxInterfaceCount returns a boolean if a field has been set.

### GetClusterDefaultInterfaceCount

`func (o *VirtualDeviceType) GetClusterDefaultInterfaceCount() int32`

GetClusterDefaultInterfaceCount returns the ClusterDefaultInterfaceCount field if non-nil, zero value otherwise.

### GetClusterDefaultInterfaceCountOk

`func (o *VirtualDeviceType) GetClusterDefaultInterfaceCountOk() (*int32, bool)`

GetClusterDefaultInterfaceCountOk returns a tuple with the ClusterDefaultInterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDefaultInterfaceCount

`func (o *VirtualDeviceType) SetClusterDefaultInterfaceCount(v int32)`

SetClusterDefaultInterfaceCount sets ClusterDefaultInterfaceCount field to given value.

### HasClusterDefaultInterfaceCount

`func (o *VirtualDeviceType) HasClusterDefaultInterfaceCount() bool`

HasClusterDefaultInterfaceCount returns a boolean if a field has been set.

### GetAvailableMetros

`func (o *VirtualDeviceType) GetAvailableMetros() []Metro`

GetAvailableMetros returns the AvailableMetros field if non-nil, zero value otherwise.

### GetAvailableMetrosOk

`func (o *VirtualDeviceType) GetAvailableMetrosOk() (*[]Metro, bool)`

GetAvailableMetrosOk returns a tuple with the AvailableMetros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMetros

`func (o *VirtualDeviceType) SetAvailableMetros(v []Metro)`

SetAvailableMetros sets AvailableMetros field to given value.

### HasAvailableMetros

`func (o *VirtualDeviceType) HasAvailableMetros() bool`

HasAvailableMetros returns a boolean if a field has been set.

### GetSoftwarePackages

`func (o *VirtualDeviceType) GetSoftwarePackages() []SoftwarePackage`

GetSoftwarePackages returns the SoftwarePackages field if non-nil, zero value otherwise.

### GetSoftwarePackagesOk

`func (o *VirtualDeviceType) GetSoftwarePackagesOk() (*[]SoftwarePackage, bool)`

GetSoftwarePackagesOk returns a tuple with the SoftwarePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwarePackages

`func (o *VirtualDeviceType) SetSoftwarePackages(v []SoftwarePackage)`

SetSoftwarePackages sets SoftwarePackages field to given value.

### HasSoftwarePackages

`func (o *VirtualDeviceType) HasSoftwarePackages() bool`

HasSoftwarePackages returns a boolean if a field has been set.

### GetDeviceManagementTypes

`func (o *VirtualDeviceType) GetDeviceManagementTypes() VirtualDeviceTypeDeviceManagementTypes`

GetDeviceManagementTypes returns the DeviceManagementTypes field if non-nil, zero value otherwise.

### GetDeviceManagementTypesOk

`func (o *VirtualDeviceType) GetDeviceManagementTypesOk() (*VirtualDeviceTypeDeviceManagementTypes, bool)`

GetDeviceManagementTypesOk returns a tuple with the DeviceManagementTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceManagementTypes

`func (o *VirtualDeviceType) SetDeviceManagementTypes(v VirtualDeviceTypeDeviceManagementTypes)`

SetDeviceManagementTypes sets DeviceManagementTypes field to given value.

### HasDeviceManagementTypes

`func (o *VirtualDeviceType) HasDeviceManagementTypes() bool`

HasDeviceManagementTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


