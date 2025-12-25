# LinkDeviceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | Pointer to **string** | A device that is part of the device linked group | [optional] 
**DeviceName** | Pointer to **string** | Device name | [optional] 
**MetroCode** | Pointer to **string** | Metro Code | [optional] 
**MetroName** | Pointer to **string** | Name of the metro. | [optional] 
**DeviceTypeCode** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**IpAssigned** | Pointer to **string** |  | [optional] 
**InterfaceId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | The status of the device | [optional] 
**DeviceManagementType** | Pointer to **string** | Device management type | [optional] 
**NetworkScope** | Pointer to **string** |  | [optional] 
**IsDeviceAccessible** | Pointer to **bool** | Whether the device is accessible | [optional] 

## Methods

### NewLinkDeviceResponse

`func NewLinkDeviceResponse() *LinkDeviceResponse`

NewLinkDeviceResponse instantiates a new LinkDeviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkDeviceResponseWithDefaults

`func NewLinkDeviceResponseWithDefaults() *LinkDeviceResponse`

NewLinkDeviceResponseWithDefaults instantiates a new LinkDeviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUuid

`func (o *LinkDeviceResponse) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *LinkDeviceResponse) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *LinkDeviceResponse) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.

### HasDeviceUuid

`func (o *LinkDeviceResponse) HasDeviceUuid() bool`

HasDeviceUuid returns a boolean if a field has been set.

### GetDeviceName

`func (o *LinkDeviceResponse) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *LinkDeviceResponse) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *LinkDeviceResponse) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *LinkDeviceResponse) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetMetroCode

`func (o *LinkDeviceResponse) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *LinkDeviceResponse) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *LinkDeviceResponse) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.

### HasMetroCode

`func (o *LinkDeviceResponse) HasMetroCode() bool`

HasMetroCode returns a boolean if a field has been set.

### GetMetroName

`func (o *LinkDeviceResponse) GetMetroName() string`

GetMetroName returns the MetroName field if non-nil, zero value otherwise.

### GetMetroNameOk

`func (o *LinkDeviceResponse) GetMetroNameOk() (*string, bool)`

GetMetroNameOk returns a tuple with the MetroName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroName

`func (o *LinkDeviceResponse) SetMetroName(v string)`

SetMetroName sets MetroName field to given value.

### HasMetroName

`func (o *LinkDeviceResponse) HasMetroName() bool`

HasMetroName returns a boolean if a field has been set.

### GetDeviceTypeCode

`func (o *LinkDeviceResponse) GetDeviceTypeCode() string`

GetDeviceTypeCode returns the DeviceTypeCode field if non-nil, zero value otherwise.

### GetDeviceTypeCodeOk

`func (o *LinkDeviceResponse) GetDeviceTypeCodeOk() (*string, bool)`

GetDeviceTypeCodeOk returns a tuple with the DeviceTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeCode

`func (o *LinkDeviceResponse) SetDeviceTypeCode(v string)`

SetDeviceTypeCode sets DeviceTypeCode field to given value.

### HasDeviceTypeCode

`func (o *LinkDeviceResponse) HasDeviceTypeCode() bool`

HasDeviceTypeCode returns a boolean if a field has been set.

### GetCategory

`func (o *LinkDeviceResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *LinkDeviceResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *LinkDeviceResponse) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *LinkDeviceResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetIpAssigned

`func (o *LinkDeviceResponse) GetIpAssigned() string`

GetIpAssigned returns the IpAssigned field if non-nil, zero value otherwise.

### GetIpAssignedOk

`func (o *LinkDeviceResponse) GetIpAssignedOk() (*string, bool)`

GetIpAssignedOk returns a tuple with the IpAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAssigned

`func (o *LinkDeviceResponse) SetIpAssigned(v string)`

SetIpAssigned sets IpAssigned field to given value.

### HasIpAssigned

`func (o *LinkDeviceResponse) HasIpAssigned() bool`

HasIpAssigned returns a boolean if a field has been set.

### GetInterfaceId

`func (o *LinkDeviceResponse) GetInterfaceId() int32`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *LinkDeviceResponse) GetInterfaceIdOk() (*int32, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *LinkDeviceResponse) SetInterfaceId(v int32)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *LinkDeviceResponse) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetStatus

`func (o *LinkDeviceResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LinkDeviceResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LinkDeviceResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LinkDeviceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeviceManagementType

`func (o *LinkDeviceResponse) GetDeviceManagementType() string`

GetDeviceManagementType returns the DeviceManagementType field if non-nil, zero value otherwise.

### GetDeviceManagementTypeOk

`func (o *LinkDeviceResponse) GetDeviceManagementTypeOk() (*string, bool)`

GetDeviceManagementTypeOk returns a tuple with the DeviceManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceManagementType

`func (o *LinkDeviceResponse) SetDeviceManagementType(v string)`

SetDeviceManagementType sets DeviceManagementType field to given value.

### HasDeviceManagementType

`func (o *LinkDeviceResponse) HasDeviceManagementType() bool`

HasDeviceManagementType returns a boolean if a field has been set.

### GetNetworkScope

`func (o *LinkDeviceResponse) GetNetworkScope() string`

GetNetworkScope returns the NetworkScope field if non-nil, zero value otherwise.

### GetNetworkScopeOk

`func (o *LinkDeviceResponse) GetNetworkScopeOk() (*string, bool)`

GetNetworkScopeOk returns a tuple with the NetworkScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkScope

`func (o *LinkDeviceResponse) SetNetworkScope(v string)`

SetNetworkScope sets NetworkScope field to given value.

### HasNetworkScope

`func (o *LinkDeviceResponse) HasNetworkScope() bool`

HasNetworkScope returns a boolean if a field has been set.

### GetIsDeviceAccessible

`func (o *LinkDeviceResponse) GetIsDeviceAccessible() bool`

GetIsDeviceAccessible returns the IsDeviceAccessible field if non-nil, zero value otherwise.

### GetIsDeviceAccessibleOk

`func (o *LinkDeviceResponse) GetIsDeviceAccessibleOk() (*bool, bool)`

GetIsDeviceAccessibleOk returns a tuple with the IsDeviceAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeviceAccessible

`func (o *LinkDeviceResponse) SetIsDeviceAccessible(v bool)`

SetIsDeviceAccessible sets IsDeviceAccessible field to given value.

### HasIsDeviceAccessible

`func (o *LinkDeviceResponse) HasIsDeviceAccessible() bool`

HasIsDeviceAccessible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


