# VrfUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpDynamicNeighborsEnabled** | Pointer to **bool** | Toggle to enable the dynamic bgp neighbors feature on the VRF | [optional] 
**BgpDynamicNeighborsExportRouteMap** | Pointer to **bool** | Toggle to export the VRF route-map to the dynamic bgp neighbors | [optional] 
**BgpDynamicNeighborsBfdEnabled** | Pointer to **bool** | Toggle BFD on dynamic bgp neighbors sessions | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IpRanges** | Pointer to **[]string** | A list of CIDR network addresses. Like [\&quot;10.0.0.0/16\&quot;, \&quot;2001:d78::/56\&quot;]. IPv4 blocks must be between /8 and /29 in size. IPv6 blocks must be between /56 and /64. A VRF\\&#39;s IP ranges must be defined in order to create VRF IP Reservations, which can then be used for Metal Gateways or Virtual Circuits. Adding a new CIDR address to the list will result in the creation of a new IP Range for this VRF. Removal of an existing CIDR address from the list will result in the deletion of an existing IP Range for this VRF. Deleting an IP Range will result in the deletion of any VRF IP Reservations contained within the IP Range, as well as the VRF IP Reservation\\&#39;s associated Metal Gateways or Virtual Circuits. If you do not wish to add or remove IP Ranges, either include the full existing list of IP Ranges in the update request, or do not specify the &#x60;ip_ranges&#x60; field in the update request. Specifying a value of &#x60;[]&#x60; will remove all existing IP Ranges from the VRF. | [optional] 
**LocalAsn** | Pointer to **int32** | The new &#x60;local_asn&#x60; value for the VRF. This field cannot be updated when there are active Interconnection Virtual Circuits associated to the VRF, or if any of the VLANs of the VRF&#39;s metal gateway has been assigned on an instance. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVrfUpdateInput

`func NewVrfUpdateInput() *VrfUpdateInput`

NewVrfUpdateInput instantiates a new VrfUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfUpdateInputWithDefaults

`func NewVrfUpdateInputWithDefaults() *VrfUpdateInput`

NewVrfUpdateInputWithDefaults instantiates a new VrfUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpDynamicNeighborsEnabled

`func (o *VrfUpdateInput) GetBgpDynamicNeighborsEnabled() bool`

GetBgpDynamicNeighborsEnabled returns the BgpDynamicNeighborsEnabled field if non-nil, zero value otherwise.

### GetBgpDynamicNeighborsEnabledOk

`func (o *VrfUpdateInput) GetBgpDynamicNeighborsEnabledOk() (*bool, bool)`

GetBgpDynamicNeighborsEnabledOk returns a tuple with the BgpDynamicNeighborsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpDynamicNeighborsEnabled

`func (o *VrfUpdateInput) SetBgpDynamicNeighborsEnabled(v bool)`

SetBgpDynamicNeighborsEnabled sets BgpDynamicNeighborsEnabled field to given value.

### HasBgpDynamicNeighborsEnabled

`func (o *VrfUpdateInput) HasBgpDynamicNeighborsEnabled() bool`

HasBgpDynamicNeighborsEnabled returns a boolean if a field has been set.

### GetBgpDynamicNeighborsExportRouteMap

`func (o *VrfUpdateInput) GetBgpDynamicNeighborsExportRouteMap() bool`

GetBgpDynamicNeighborsExportRouteMap returns the BgpDynamicNeighborsExportRouteMap field if non-nil, zero value otherwise.

### GetBgpDynamicNeighborsExportRouteMapOk

`func (o *VrfUpdateInput) GetBgpDynamicNeighborsExportRouteMapOk() (*bool, bool)`

GetBgpDynamicNeighborsExportRouteMapOk returns a tuple with the BgpDynamicNeighborsExportRouteMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpDynamicNeighborsExportRouteMap

`func (o *VrfUpdateInput) SetBgpDynamicNeighborsExportRouteMap(v bool)`

SetBgpDynamicNeighborsExportRouteMap sets BgpDynamicNeighborsExportRouteMap field to given value.

### HasBgpDynamicNeighborsExportRouteMap

`func (o *VrfUpdateInput) HasBgpDynamicNeighborsExportRouteMap() bool`

HasBgpDynamicNeighborsExportRouteMap returns a boolean if a field has been set.

### GetBgpDynamicNeighborsBfdEnabled

`func (o *VrfUpdateInput) GetBgpDynamicNeighborsBfdEnabled() bool`

GetBgpDynamicNeighborsBfdEnabled returns the BgpDynamicNeighborsBfdEnabled field if non-nil, zero value otherwise.

### GetBgpDynamicNeighborsBfdEnabledOk

`func (o *VrfUpdateInput) GetBgpDynamicNeighborsBfdEnabledOk() (*bool, bool)`

GetBgpDynamicNeighborsBfdEnabledOk returns a tuple with the BgpDynamicNeighborsBfdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpDynamicNeighborsBfdEnabled

`func (o *VrfUpdateInput) SetBgpDynamicNeighborsBfdEnabled(v bool)`

SetBgpDynamicNeighborsBfdEnabled sets BgpDynamicNeighborsBfdEnabled field to given value.

### HasBgpDynamicNeighborsBfdEnabled

`func (o *VrfUpdateInput) HasBgpDynamicNeighborsBfdEnabled() bool`

HasBgpDynamicNeighborsBfdEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *VrfUpdateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VrfUpdateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VrfUpdateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VrfUpdateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpRanges

`func (o *VrfUpdateInput) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *VrfUpdateInput) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *VrfUpdateInput) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *VrfUpdateInput) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetLocalAsn

`func (o *VrfUpdateInput) GetLocalAsn() int32`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *VrfUpdateInput) GetLocalAsnOk() (*int32, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *VrfUpdateInput) SetLocalAsn(v int32)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *VrfUpdateInput) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetName

`func (o *VrfUpdateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VrfUpdateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VrfUpdateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VrfUpdateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *VrfUpdateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VrfUpdateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VrfUpdateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VrfUpdateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


