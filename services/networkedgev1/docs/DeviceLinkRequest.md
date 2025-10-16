# DeviceLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **string** | Group name. | 
**Subnet** | **string** | Subnet of the link group. | 
**RedundancyType** | Pointer to **string** | Whether the connection should be created through Fabric&#39;s primary or secondary port. | [optional] 
**LinkDevices** | Pointer to [**[]LinkDeviceInfo**](LinkDeviceInfo.md) | An array of devices to link. | [optional] 
**MetroLinks** | Pointer to [**[]LinkInfo**](LinkInfo.md) | An array of links. | [optional] 

## Methods

### NewDeviceLinkRequest

`func NewDeviceLinkRequest(groupName string, subnet string, ) *DeviceLinkRequest`

NewDeviceLinkRequest instantiates a new DeviceLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceLinkRequestWithDefaults

`func NewDeviceLinkRequestWithDefaults() *DeviceLinkRequest`

NewDeviceLinkRequestWithDefaults instantiates a new DeviceLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *DeviceLinkRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DeviceLinkRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DeviceLinkRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetSubnet

`func (o *DeviceLinkRequest) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *DeviceLinkRequest) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *DeviceLinkRequest) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetRedundancyType

`func (o *DeviceLinkRequest) GetRedundancyType() string`

GetRedundancyType returns the RedundancyType field if non-nil, zero value otherwise.

### GetRedundancyTypeOk

`func (o *DeviceLinkRequest) GetRedundancyTypeOk() (*string, bool)`

GetRedundancyTypeOk returns a tuple with the RedundancyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyType

`func (o *DeviceLinkRequest) SetRedundancyType(v string)`

SetRedundancyType sets RedundancyType field to given value.

### HasRedundancyType

`func (o *DeviceLinkRequest) HasRedundancyType() bool`

HasRedundancyType returns a boolean if a field has been set.

### GetLinkDevices

`func (o *DeviceLinkRequest) GetLinkDevices() []LinkDeviceInfo`

GetLinkDevices returns the LinkDevices field if non-nil, zero value otherwise.

### GetLinkDevicesOk

`func (o *DeviceLinkRequest) GetLinkDevicesOk() (*[]LinkDeviceInfo, bool)`

GetLinkDevicesOk returns a tuple with the LinkDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDevices

`func (o *DeviceLinkRequest) SetLinkDevices(v []LinkDeviceInfo)`

SetLinkDevices sets LinkDevices field to given value.

### HasLinkDevices

`func (o *DeviceLinkRequest) HasLinkDevices() bool`

HasLinkDevices returns a boolean if a field has been set.

### GetMetroLinks

`func (o *DeviceLinkRequest) GetMetroLinks() []LinkInfo`

GetMetroLinks returns the MetroLinks field if non-nil, zero value otherwise.

### GetMetroLinksOk

`func (o *DeviceLinkRequest) GetMetroLinksOk() (*[]LinkInfo, bool)`

GetMetroLinksOk returns a tuple with the MetroLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroLinks

`func (o *DeviceLinkRequest) SetMetroLinks(v []LinkInfo)`

SetMetroLinks sets MetroLinks field to given value.

### HasMetroLinks

`func (o *DeviceLinkRequest) HasMetroLinks() bool`

HasMetroLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


