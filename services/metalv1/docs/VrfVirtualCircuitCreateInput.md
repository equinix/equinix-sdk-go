# VrfVirtualCircuitCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerIp** | Pointer to **string** | An IP address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Md5** | Pointer to **NullableString** | The plaintext BGP peering password shared by neighbors as an MD5 checksum: * must be 10-20 characters long * may not include punctuation * must be a combination of numbers and letters * must contain at least one lowercase, uppercase, and digit character  | [optional] 
**MetalIp** | Pointer to **string** | An IP address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NniVlan** | **int32** |  | 
**PeerAsn** | **int64** | The peer ASN that will be used with the VRF on the Virtual Circuit. | 
**ProjectId** | **string** |  | 
**Speed** | Pointer to **string** | speed can be passed as integer number representing bps speed or string (e.g. &#39;52m&#39; or &#39;100g&#39; or &#39;4 gbps&#39;) | [optional] 
**Subnet** | **string** | The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP. The subnet specified must be contained within an already-defined IP Range for the VRF. | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Vrf** | **string** | The UUID of the VRF that will be associated with the Virtual Circuit. | 

## Methods

### NewVrfVirtualCircuitCreateInput

`func NewVrfVirtualCircuitCreateInput(nniVlan int32, peerAsn int64, projectId string, subnet string, vrf string, ) *VrfVirtualCircuitCreateInput`

NewVrfVirtualCircuitCreateInput instantiates a new VrfVirtualCircuitCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfVirtualCircuitCreateInputWithDefaults

`func NewVrfVirtualCircuitCreateInputWithDefaults() *VrfVirtualCircuitCreateInput`

NewVrfVirtualCircuitCreateInputWithDefaults instantiates a new VrfVirtualCircuitCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerIp

`func (o *VrfVirtualCircuitCreateInput) GetCustomerIp() string`

GetCustomerIp returns the CustomerIp field if non-nil, zero value otherwise.

### GetCustomerIpOk

`func (o *VrfVirtualCircuitCreateInput) GetCustomerIpOk() (*string, bool)`

GetCustomerIpOk returns a tuple with the CustomerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIp

`func (o *VrfVirtualCircuitCreateInput) SetCustomerIp(v string)`

SetCustomerIp sets CustomerIp field to given value.

### HasCustomerIp

`func (o *VrfVirtualCircuitCreateInput) HasCustomerIp() bool`

HasCustomerIp returns a boolean if a field has been set.

### GetDescription

`func (o *VrfVirtualCircuitCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VrfVirtualCircuitCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VrfVirtualCircuitCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VrfVirtualCircuitCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMd5

`func (o *VrfVirtualCircuitCreateInput) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *VrfVirtualCircuitCreateInput) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *VrfVirtualCircuitCreateInput) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *VrfVirtualCircuitCreateInput) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *VrfVirtualCircuitCreateInput) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *VrfVirtualCircuitCreateInput) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetMetalIp

`func (o *VrfVirtualCircuitCreateInput) GetMetalIp() string`

GetMetalIp returns the MetalIp field if non-nil, zero value otherwise.

### GetMetalIpOk

`func (o *VrfVirtualCircuitCreateInput) GetMetalIpOk() (*string, bool)`

GetMetalIpOk returns a tuple with the MetalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalIp

`func (o *VrfVirtualCircuitCreateInput) SetMetalIp(v string)`

SetMetalIp sets MetalIp field to given value.

### HasMetalIp

`func (o *VrfVirtualCircuitCreateInput) HasMetalIp() bool`

HasMetalIp returns a boolean if a field has been set.

### GetName

`func (o *VrfVirtualCircuitCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VrfVirtualCircuitCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VrfVirtualCircuitCreateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VrfVirtualCircuitCreateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNniVlan

`func (o *VrfVirtualCircuitCreateInput) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *VrfVirtualCircuitCreateInput) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *VrfVirtualCircuitCreateInput) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.


### GetPeerAsn

`func (o *VrfVirtualCircuitCreateInput) GetPeerAsn() int64`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *VrfVirtualCircuitCreateInput) GetPeerAsnOk() (*int64, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *VrfVirtualCircuitCreateInput) SetPeerAsn(v int64)`

SetPeerAsn sets PeerAsn field to given value.


### GetProjectId

`func (o *VrfVirtualCircuitCreateInput) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *VrfVirtualCircuitCreateInput) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *VrfVirtualCircuitCreateInput) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSpeed

`func (o *VrfVirtualCircuitCreateInput) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VrfVirtualCircuitCreateInput) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VrfVirtualCircuitCreateInput) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VrfVirtualCircuitCreateInput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetSubnet

`func (o *VrfVirtualCircuitCreateInput) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *VrfVirtualCircuitCreateInput) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *VrfVirtualCircuitCreateInput) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetTags

`func (o *VrfVirtualCircuitCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VrfVirtualCircuitCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VrfVirtualCircuitCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VrfVirtualCircuitCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVrf

`func (o *VrfVirtualCircuitCreateInput) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *VrfVirtualCircuitCreateInput) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *VrfVirtualCircuitCreateInput) SetVrf(v string)`

SetVrf sets Vrf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


