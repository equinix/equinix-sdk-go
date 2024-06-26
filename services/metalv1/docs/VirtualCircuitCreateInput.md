# VirtualCircuitCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NniVlan** | **int32** |  | 
**ProjectId** | **string** |  | 
**Speed** | Pointer to **string** | speed can be passed as integer number representing bps speed or string (e.g. &#39;52m&#39; or &#39;100g&#39; or &#39;4 gbps&#39;) | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Vnid** | Pointer to **string** | A Virtual Network record UUID or the VNID of a Metro Virtual Network in your project (sent as integer). | [optional] 
**CustomerIp** | Pointer to **string** | An IPv4 address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used. | [optional] 
**Md5** | Pointer to **NullableString** | The plaintext BGP peering password shared by neighbors as an MD5 checksum: * must be 10-20 characters long * may not include punctuation * must be a combination of numbers and letters * must contain at least one lowercase, uppercase, and digit character  | [optional] 
**MetalIp** | Pointer to **string** | An IPv4 address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used. | [optional] 
**PeerAsn** | **int64** | The peer ASN that will be used with the VRF on the Virtual Circuit. | 
**Subnet** | **string** | The /30 or /31 IPv4 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP. The subnet specified must be contained within an already-defined IP Range for the VRF. | 
**Vrf** | **string** | The UUID of the VRF that will be associated with the Virtual Circuit. | 
**SubnetIpv6** | Pointer to **string** | The /126 or /127 IPv6 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IPv6 and Customer IPv6 must be IPs from this subnet. For /126 subnets, the network and broadcast IPs cannot be used as the Metal IPv6 or Customer IPv6. The subnet specified must be contained within an already-defined IP Range for the VRF. | [optional] 
**CustomerIpv6** | Pointer to **string** | An IPv6 address from the subnet IPv6 that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet IPv6 as the Metal IPv6. By default, the last usable IP address in the subnet IPv6 will be used. | [optional] 
**MetalIpv6** | Pointer to **string** | An IPv6 address from the subnet IPv6 that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IPv6 address in the subnet IPv6 as the Customer IP. By default, the first usable IPv6 address in the subnet IPv6 will be used. | [optional] 

## Methods

### NewVirtualCircuitCreateInput

`func NewVirtualCircuitCreateInput(nniVlan int32, projectId string, peerAsn int64, subnet string, vrf string, ) *VirtualCircuitCreateInput`

NewVirtualCircuitCreateInput instantiates a new VirtualCircuitCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCircuitCreateInputWithDefaults

`func NewVirtualCircuitCreateInputWithDefaults() *VirtualCircuitCreateInput`

NewVirtualCircuitCreateInputWithDefaults instantiates a new VirtualCircuitCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VirtualCircuitCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualCircuitCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualCircuitCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualCircuitCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *VirtualCircuitCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualCircuitCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualCircuitCreateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualCircuitCreateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNniVlan

`func (o *VirtualCircuitCreateInput) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *VirtualCircuitCreateInput) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *VirtualCircuitCreateInput) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.


### GetProjectId

`func (o *VirtualCircuitCreateInput) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *VirtualCircuitCreateInput) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *VirtualCircuitCreateInput) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSpeed

`func (o *VirtualCircuitCreateInput) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VirtualCircuitCreateInput) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VirtualCircuitCreateInput) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VirtualCircuitCreateInput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *VirtualCircuitCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualCircuitCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualCircuitCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualCircuitCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVnid

`func (o *VirtualCircuitCreateInput) GetVnid() string`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *VirtualCircuitCreateInput) GetVnidOk() (*string, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *VirtualCircuitCreateInput) SetVnid(v string)`

SetVnid sets Vnid field to given value.

### HasVnid

`func (o *VirtualCircuitCreateInput) HasVnid() bool`

HasVnid returns a boolean if a field has been set.

### GetCustomerIp

`func (o *VirtualCircuitCreateInput) GetCustomerIp() string`

GetCustomerIp returns the CustomerIp field if non-nil, zero value otherwise.

### GetCustomerIpOk

`func (o *VirtualCircuitCreateInput) GetCustomerIpOk() (*string, bool)`

GetCustomerIpOk returns a tuple with the CustomerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIp

`func (o *VirtualCircuitCreateInput) SetCustomerIp(v string)`

SetCustomerIp sets CustomerIp field to given value.

### HasCustomerIp

`func (o *VirtualCircuitCreateInput) HasCustomerIp() bool`

HasCustomerIp returns a boolean if a field has been set.

### GetMd5

`func (o *VirtualCircuitCreateInput) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *VirtualCircuitCreateInput) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *VirtualCircuitCreateInput) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *VirtualCircuitCreateInput) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *VirtualCircuitCreateInput) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *VirtualCircuitCreateInput) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetMetalIp

`func (o *VirtualCircuitCreateInput) GetMetalIp() string`

GetMetalIp returns the MetalIp field if non-nil, zero value otherwise.

### GetMetalIpOk

`func (o *VirtualCircuitCreateInput) GetMetalIpOk() (*string, bool)`

GetMetalIpOk returns a tuple with the MetalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalIp

`func (o *VirtualCircuitCreateInput) SetMetalIp(v string)`

SetMetalIp sets MetalIp field to given value.

### HasMetalIp

`func (o *VirtualCircuitCreateInput) HasMetalIp() bool`

HasMetalIp returns a boolean if a field has been set.

### GetPeerAsn

`func (o *VirtualCircuitCreateInput) GetPeerAsn() int64`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *VirtualCircuitCreateInput) GetPeerAsnOk() (*int64, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *VirtualCircuitCreateInput) SetPeerAsn(v int64)`

SetPeerAsn sets PeerAsn field to given value.


### GetSubnet

`func (o *VirtualCircuitCreateInput) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *VirtualCircuitCreateInput) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *VirtualCircuitCreateInput) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetVrf

`func (o *VirtualCircuitCreateInput) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *VirtualCircuitCreateInput) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *VirtualCircuitCreateInput) SetVrf(v string)`

SetVrf sets Vrf field to given value.


### GetSubnetIpv6

`func (o *VirtualCircuitCreateInput) GetSubnetIpv6() string`

GetSubnetIpv6 returns the SubnetIpv6 field if non-nil, zero value otherwise.

### GetSubnetIpv6Ok

`func (o *VirtualCircuitCreateInput) GetSubnetIpv6Ok() (*string, bool)`

GetSubnetIpv6Ok returns a tuple with the SubnetIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIpv6

`func (o *VirtualCircuitCreateInput) SetSubnetIpv6(v string)`

SetSubnetIpv6 sets SubnetIpv6 field to given value.

### HasSubnetIpv6

`func (o *VirtualCircuitCreateInput) HasSubnetIpv6() bool`

HasSubnetIpv6 returns a boolean if a field has been set.

### GetCustomerIpv6

`func (o *VirtualCircuitCreateInput) GetCustomerIpv6() string`

GetCustomerIpv6 returns the CustomerIpv6 field if non-nil, zero value otherwise.

### GetCustomerIpv6Ok

`func (o *VirtualCircuitCreateInput) GetCustomerIpv6Ok() (*string, bool)`

GetCustomerIpv6Ok returns a tuple with the CustomerIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIpv6

`func (o *VirtualCircuitCreateInput) SetCustomerIpv6(v string)`

SetCustomerIpv6 sets CustomerIpv6 field to given value.

### HasCustomerIpv6

`func (o *VirtualCircuitCreateInput) HasCustomerIpv6() bool`

HasCustomerIpv6 returns a boolean if a field has been set.

### GetMetalIpv6

`func (o *VirtualCircuitCreateInput) GetMetalIpv6() string`

GetMetalIpv6 returns the MetalIpv6 field if non-nil, zero value otherwise.

### GetMetalIpv6Ok

`func (o *VirtualCircuitCreateInput) GetMetalIpv6Ok() (*string, bool)`

GetMetalIpv6Ok returns a tuple with the MetalIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalIpv6

`func (o *VirtualCircuitCreateInput) SetMetalIpv6(v string)`

SetMetalIpv6 sets MetalIpv6 field to given value.

### HasMetalIpv6

`func (o *VirtualCircuitCreateInput) HasMetalIpv6() bool`

HasMetalIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


