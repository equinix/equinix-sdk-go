# DeviceUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysPxe** | Pointer to **bool** |  | [optional] 
**BillingCycle** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**Description** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**FirmwareSetId** | Pointer to **string** |  | [optional] 
**IpxeScriptUrl** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** | Whether the device should be locked, preventing accidental deletion. | [optional] 
**NetworkFrozen** | Pointer to **bool** | If true, this instance can not be converted to a different network type. | [optional] 
**SpotInstance** | Pointer to **bool** | Can be set to false to convert a spot-market instance to on-demand. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Userdata** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceUpdateInput

`func NewDeviceUpdateInput() *DeviceUpdateInput`

NewDeviceUpdateInput instantiates a new DeviceUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceUpdateInputWithDefaults

`func NewDeviceUpdateInputWithDefaults() *DeviceUpdateInput`

NewDeviceUpdateInputWithDefaults instantiates a new DeviceUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysPxe

`func (o *DeviceUpdateInput) GetAlwaysPxe() bool`

GetAlwaysPxe returns the AlwaysPxe field if non-nil, zero value otherwise.

### GetAlwaysPxeOk

`func (o *DeviceUpdateInput) GetAlwaysPxeOk() (*bool, bool)`

GetAlwaysPxeOk returns a tuple with the AlwaysPxe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysPxe

`func (o *DeviceUpdateInput) SetAlwaysPxe(v bool)`

SetAlwaysPxe sets AlwaysPxe field to given value.

### HasAlwaysPxe

`func (o *DeviceUpdateInput) HasAlwaysPxe() bool`

HasAlwaysPxe returns a boolean if a field has been set.

### GetBillingCycle

`func (o *DeviceUpdateInput) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *DeviceUpdateInput) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *DeviceUpdateInput) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *DeviceUpdateInput) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetCustomdata

`func (o *DeviceUpdateInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *DeviceUpdateInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *DeviceUpdateInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *DeviceUpdateInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceUpdateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceUpdateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceUpdateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceUpdateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostname

`func (o *DeviceUpdateInput) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DeviceUpdateInput) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DeviceUpdateInput) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DeviceUpdateInput) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetFirmwareSetId

`func (o *DeviceUpdateInput) GetFirmwareSetId() string`

GetFirmwareSetId returns the FirmwareSetId field if non-nil, zero value otherwise.

### GetFirmwareSetIdOk

`func (o *DeviceUpdateInput) GetFirmwareSetIdOk() (*string, bool)`

GetFirmwareSetIdOk returns a tuple with the FirmwareSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareSetId

`func (o *DeviceUpdateInput) SetFirmwareSetId(v string)`

SetFirmwareSetId sets FirmwareSetId field to given value.

### HasFirmwareSetId

`func (o *DeviceUpdateInput) HasFirmwareSetId() bool`

HasFirmwareSetId returns a boolean if a field has been set.

### GetIpxeScriptUrl

`func (o *DeviceUpdateInput) GetIpxeScriptUrl() string`

GetIpxeScriptUrl returns the IpxeScriptUrl field if non-nil, zero value otherwise.

### GetIpxeScriptUrlOk

`func (o *DeviceUpdateInput) GetIpxeScriptUrlOk() (*string, bool)`

GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpxeScriptUrl

`func (o *DeviceUpdateInput) SetIpxeScriptUrl(v string)`

SetIpxeScriptUrl sets IpxeScriptUrl field to given value.

### HasIpxeScriptUrl

`func (o *DeviceUpdateInput) HasIpxeScriptUrl() bool`

HasIpxeScriptUrl returns a boolean if a field has been set.

### GetLocked

`func (o *DeviceUpdateInput) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *DeviceUpdateInput) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *DeviceUpdateInput) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *DeviceUpdateInput) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetNetworkFrozen

`func (o *DeviceUpdateInput) GetNetworkFrozen() bool`

GetNetworkFrozen returns the NetworkFrozen field if non-nil, zero value otherwise.

### GetNetworkFrozenOk

`func (o *DeviceUpdateInput) GetNetworkFrozenOk() (*bool, bool)`

GetNetworkFrozenOk returns a tuple with the NetworkFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFrozen

`func (o *DeviceUpdateInput) SetNetworkFrozen(v bool)`

SetNetworkFrozen sets NetworkFrozen field to given value.

### HasNetworkFrozen

`func (o *DeviceUpdateInput) HasNetworkFrozen() bool`

HasNetworkFrozen returns a boolean if a field has been set.

### GetSpotInstance

`func (o *DeviceUpdateInput) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *DeviceUpdateInput) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *DeviceUpdateInput) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.

### HasSpotInstance

`func (o *DeviceUpdateInput) HasSpotInstance() bool`

HasSpotInstance returns a boolean if a field has been set.

### GetTags

`func (o *DeviceUpdateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceUpdateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceUpdateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceUpdateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUserdata

`func (o *DeviceUpdateInput) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *DeviceUpdateInput) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *DeviceUpdateInput) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *DeviceUpdateInput) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


