# DeviceLinkGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Unique Id of the linked group. | [optional] 
**GroupName** | Pointer to **string** | The name of the linked group | [optional] 
**Subnet** | Pointer to **string** | Subnet of the link group. | [optional] 
**RedundancyType** | Pointer to **string** | Whether the connection is through Fabric&#39;s primary or secondary port. | [optional] 
**Status** | Pointer to **string** | Status of the linked group | [optional] 
**CreatedBy** | Pointer to **string** | Created by username. | [optional] 
**CreatedDate** | Pointer to **string** | Created date. | [optional] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] 
**LastUpdatedDate** | Pointer to **string** |  | [optional] 
**MetroLinks** | Pointer to [**[]LinkInfoResponse**](LinkInfoResponse.md) | An array of links | [optional] 
**LinkDevices** | Pointer to [**[]LinkDeviceResponse**](LinkDeviceResponse.md) | An array of metros and the devices in the metros belonging to the group. | [optional] 

## Methods

### NewDeviceLinkGroupDto

`func NewDeviceLinkGroupDto() *DeviceLinkGroupDto`

NewDeviceLinkGroupDto instantiates a new DeviceLinkGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceLinkGroupDtoWithDefaults

`func NewDeviceLinkGroupDtoWithDefaults() *DeviceLinkGroupDto`

NewDeviceLinkGroupDtoWithDefaults instantiates a new DeviceLinkGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *DeviceLinkGroupDto) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeviceLinkGroupDto) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeviceLinkGroupDto) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeviceLinkGroupDto) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetGroupName

`func (o *DeviceLinkGroupDto) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DeviceLinkGroupDto) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DeviceLinkGroupDto) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *DeviceLinkGroupDto) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetSubnet

`func (o *DeviceLinkGroupDto) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *DeviceLinkGroupDto) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *DeviceLinkGroupDto) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *DeviceLinkGroupDto) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetRedundancyType

`func (o *DeviceLinkGroupDto) GetRedundancyType() string`

GetRedundancyType returns the RedundancyType field if non-nil, zero value otherwise.

### GetRedundancyTypeOk

`func (o *DeviceLinkGroupDto) GetRedundancyTypeOk() (*string, bool)`

GetRedundancyTypeOk returns a tuple with the RedundancyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyType

`func (o *DeviceLinkGroupDto) SetRedundancyType(v string)`

SetRedundancyType sets RedundancyType field to given value.

### HasRedundancyType

`func (o *DeviceLinkGroupDto) HasRedundancyType() bool`

HasRedundancyType returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceLinkGroupDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceLinkGroupDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceLinkGroupDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceLinkGroupDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DeviceLinkGroupDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DeviceLinkGroupDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DeviceLinkGroupDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DeviceLinkGroupDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *DeviceLinkGroupDto) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *DeviceLinkGroupDto) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *DeviceLinkGroupDto) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *DeviceLinkGroupDto) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *DeviceLinkGroupDto) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *DeviceLinkGroupDto) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *DeviceLinkGroupDto) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *DeviceLinkGroupDto) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *DeviceLinkGroupDto) GetLastUpdatedDate() string`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *DeviceLinkGroupDto) GetLastUpdatedDateOk() (*string, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *DeviceLinkGroupDto) SetLastUpdatedDate(v string)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *DeviceLinkGroupDto) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetMetroLinks

`func (o *DeviceLinkGroupDto) GetMetroLinks() []LinkInfoResponse`

GetMetroLinks returns the MetroLinks field if non-nil, zero value otherwise.

### GetMetroLinksOk

`func (o *DeviceLinkGroupDto) GetMetroLinksOk() (*[]LinkInfoResponse, bool)`

GetMetroLinksOk returns a tuple with the MetroLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroLinks

`func (o *DeviceLinkGroupDto) SetMetroLinks(v []LinkInfoResponse)`

SetMetroLinks sets MetroLinks field to given value.

### HasMetroLinks

`func (o *DeviceLinkGroupDto) HasMetroLinks() bool`

HasMetroLinks returns a boolean if a field has been set.

### GetLinkDevices

`func (o *DeviceLinkGroupDto) GetLinkDevices() []LinkDeviceResponse`

GetLinkDevices returns the LinkDevices field if non-nil, zero value otherwise.

### GetLinkDevicesOk

`func (o *DeviceLinkGroupDto) GetLinkDevicesOk() (*[]LinkDeviceResponse, bool)`

GetLinkDevicesOk returns a tuple with the LinkDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDevices

`func (o *DeviceLinkGroupDto) SetLinkDevices(v []LinkDeviceResponse)`

SetLinkDevices sets LinkDevices field to given value.

### HasLinkDevices

`func (o *DeviceLinkGroupDto) HasLinkDevices() bool`

HasLinkDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


