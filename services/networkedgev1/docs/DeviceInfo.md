# DeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aside** | Pointer to [**JsonNode**](JsonNode.md) |  | [optional] 
**Category** | Pointer to **string** | Category of the device. | [optional] 
**CloudProfileProvisioningStatus** | Pointer to **string** |  | [optional] 
**ConnectionStatus** | Pointer to **string** |  | [optional] 
**ConnectionUuid** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** | Name of the device. | [optional] 
**DeviceTypeCode** | Pointer to **string** | Device type code. | [optional] 
**DeviceUUID** | Pointer to **string** | Unique Id of the device. | [optional] 
**InterfaceId** | Pointer to **string** |  | [optional] 
**InterfaceOverlayStatus** | Pointer to **string** |  | [optional] 
**InterfaceUUID** | Pointer to **string** | Unique Id of the interface used to link the device. | [optional] 
**IpAssigned** | Pointer to **string** | Assigned IP address of the device. | [optional] 
**NetworkDetails** | Pointer to [**JsonNode**](JsonNode.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the device. | [optional] 
**Throughput** | Pointer to **string** | Throughput of the device. | [optional] 
**ThroughputUnit** | Pointer to **string** | Throughput unit of the device. | [optional] 
**Vxlan** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceInfo

`func NewDeviceInfo() *DeviceInfo`

NewDeviceInfo instantiates a new DeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInfoWithDefaults

`func NewDeviceInfoWithDefaults() *DeviceInfo`

NewDeviceInfoWithDefaults instantiates a new DeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAside

`func (o *DeviceInfo) GetAside() JsonNode`

GetAside returns the Aside field if non-nil, zero value otherwise.

### GetAsideOk

`func (o *DeviceInfo) GetAsideOk() (*JsonNode, bool)`

GetAsideOk returns a tuple with the Aside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAside

`func (o *DeviceInfo) SetAside(v JsonNode)`

SetAside sets Aside field to given value.

### HasAside

`func (o *DeviceInfo) HasAside() bool`

HasAside returns a boolean if a field has been set.

### GetCategory

`func (o *DeviceInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DeviceInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DeviceInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *DeviceInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCloudProfileProvisioningStatus

`func (o *DeviceInfo) GetCloudProfileProvisioningStatus() string`

GetCloudProfileProvisioningStatus returns the CloudProfileProvisioningStatus field if non-nil, zero value otherwise.

### GetCloudProfileProvisioningStatusOk

`func (o *DeviceInfo) GetCloudProfileProvisioningStatusOk() (*string, bool)`

GetCloudProfileProvisioningStatusOk returns a tuple with the CloudProfileProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProfileProvisioningStatus

`func (o *DeviceInfo) SetCloudProfileProvisioningStatus(v string)`

SetCloudProfileProvisioningStatus sets CloudProfileProvisioningStatus field to given value.

### HasCloudProfileProvisioningStatus

`func (o *DeviceInfo) HasCloudProfileProvisioningStatus() bool`

HasCloudProfileProvisioningStatus returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *DeviceInfo) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *DeviceInfo) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *DeviceInfo) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *DeviceInfo) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetConnectionUuid

`func (o *DeviceInfo) GetConnectionUuid() string`

GetConnectionUuid returns the ConnectionUuid field if non-nil, zero value otherwise.

### GetConnectionUuidOk

`func (o *DeviceInfo) GetConnectionUuidOk() (*string, bool)`

GetConnectionUuidOk returns a tuple with the ConnectionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUuid

`func (o *DeviceInfo) SetConnectionUuid(v string)`

SetConnectionUuid sets ConnectionUuid field to given value.

### HasConnectionUuid

`func (o *DeviceInfo) HasConnectionUuid() bool`

HasConnectionUuid returns a boolean if a field has been set.

### GetDeviceName

`func (o *DeviceInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DeviceInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceTypeCode

`func (o *DeviceInfo) GetDeviceTypeCode() string`

GetDeviceTypeCode returns the DeviceTypeCode field if non-nil, zero value otherwise.

### GetDeviceTypeCodeOk

`func (o *DeviceInfo) GetDeviceTypeCodeOk() (*string, bool)`

GetDeviceTypeCodeOk returns a tuple with the DeviceTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeCode

`func (o *DeviceInfo) SetDeviceTypeCode(v string)`

SetDeviceTypeCode sets DeviceTypeCode field to given value.

### HasDeviceTypeCode

`func (o *DeviceInfo) HasDeviceTypeCode() bool`

HasDeviceTypeCode returns a boolean if a field has been set.

### GetDeviceUUID

`func (o *DeviceInfo) GetDeviceUUID() string`

GetDeviceUUID returns the DeviceUUID field if non-nil, zero value otherwise.

### GetDeviceUUIDOk

`func (o *DeviceInfo) GetDeviceUUIDOk() (*string, bool)`

GetDeviceUUIDOk returns a tuple with the DeviceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUUID

`func (o *DeviceInfo) SetDeviceUUID(v string)`

SetDeviceUUID sets DeviceUUID field to given value.

### HasDeviceUUID

`func (o *DeviceInfo) HasDeviceUUID() bool`

HasDeviceUUID returns a boolean if a field has been set.

### GetInterfaceId

`func (o *DeviceInfo) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *DeviceInfo) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *DeviceInfo) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *DeviceInfo) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceOverlayStatus

`func (o *DeviceInfo) GetInterfaceOverlayStatus() string`

GetInterfaceOverlayStatus returns the InterfaceOverlayStatus field if non-nil, zero value otherwise.

### GetInterfaceOverlayStatusOk

`func (o *DeviceInfo) GetInterfaceOverlayStatusOk() (*string, bool)`

GetInterfaceOverlayStatusOk returns a tuple with the InterfaceOverlayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceOverlayStatus

`func (o *DeviceInfo) SetInterfaceOverlayStatus(v string)`

SetInterfaceOverlayStatus sets InterfaceOverlayStatus field to given value.

### HasInterfaceOverlayStatus

`func (o *DeviceInfo) HasInterfaceOverlayStatus() bool`

HasInterfaceOverlayStatus returns a boolean if a field has been set.

### GetInterfaceUUID

`func (o *DeviceInfo) GetInterfaceUUID() string`

GetInterfaceUUID returns the InterfaceUUID field if non-nil, zero value otherwise.

### GetInterfaceUUIDOk

`func (o *DeviceInfo) GetInterfaceUUIDOk() (*string, bool)`

GetInterfaceUUIDOk returns a tuple with the InterfaceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceUUID

`func (o *DeviceInfo) SetInterfaceUUID(v string)`

SetInterfaceUUID sets InterfaceUUID field to given value.

### HasInterfaceUUID

`func (o *DeviceInfo) HasInterfaceUUID() bool`

HasInterfaceUUID returns a boolean if a field has been set.

### GetIpAssigned

`func (o *DeviceInfo) GetIpAssigned() string`

GetIpAssigned returns the IpAssigned field if non-nil, zero value otherwise.

### GetIpAssignedOk

`func (o *DeviceInfo) GetIpAssignedOk() (*string, bool)`

GetIpAssignedOk returns a tuple with the IpAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAssigned

`func (o *DeviceInfo) SetIpAssigned(v string)`

SetIpAssigned sets IpAssigned field to given value.

### HasIpAssigned

`func (o *DeviceInfo) HasIpAssigned() bool`

HasIpAssigned returns a boolean if a field has been set.

### GetNetworkDetails

`func (o *DeviceInfo) GetNetworkDetails() JsonNode`

GetNetworkDetails returns the NetworkDetails field if non-nil, zero value otherwise.

### GetNetworkDetailsOk

`func (o *DeviceInfo) GetNetworkDetailsOk() (*JsonNode, bool)`

GetNetworkDetailsOk returns a tuple with the NetworkDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDetails

`func (o *DeviceInfo) SetNetworkDetails(v JsonNode)`

SetNetworkDetails sets NetworkDetails field to given value.

### HasNetworkDetails

`func (o *DeviceInfo) HasNetworkDetails() bool`

HasNetworkDetails returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThroughput

`func (o *DeviceInfo) GetThroughput() string`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *DeviceInfo) GetThroughputOk() (*string, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *DeviceInfo) SetThroughput(v string)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *DeviceInfo) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetThroughputUnit

`func (o *DeviceInfo) GetThroughputUnit() string`

GetThroughputUnit returns the ThroughputUnit field if non-nil, zero value otherwise.

### GetThroughputUnitOk

`func (o *DeviceInfo) GetThroughputUnitOk() (*string, bool)`

GetThroughputUnitOk returns a tuple with the ThroughputUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputUnit

`func (o *DeviceInfo) SetThroughputUnit(v string)`

SetThroughputUnit sets ThroughputUnit field to given value.

### HasThroughputUnit

`func (o *DeviceInfo) HasThroughputUnit() bool`

HasThroughputUnit returns a boolean if a field has been set.

### GetVxlan

`func (o *DeviceInfo) GetVxlan() string`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *DeviceInfo) GetVxlanOk() (*string, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *DeviceInfo) SetVxlan(v string)`

SetVxlan sets Vxlan field to given value.

### HasVxlan

`func (o *DeviceInfo) HasVxlan() bool`

HasVxlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


