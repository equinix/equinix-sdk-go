# VrfCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpDynamicNeighborsEnabled** | Pointer to **bool** | Toggle to enable the dynamic bgp neighbors feature on the VRF | [optional] 
**BgpDynamicNeighborsExportRouteMap** | Pointer to **bool** | Toggle to export the VRF route-map to the dynamic bgp neighbors | [optional] 
**BgpDynamicNeighborsBfdEnabled** | Pointer to **bool** | Toggle BFD on dynamic bgp neighbors sessions | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IpRanges** | Pointer to **[]string** | A list of CIDR network addresses. Like [\&quot;10.0.0.0/16\&quot;, \&quot;2001:d78::/59\&quot;]. IPv4 blocks must be between /8 and /29 in size. IPv6 blocks must be between /59 and /64. A VRF\\&#39;s IP ranges must be defined in order to create VRF IP Reservations, which can then be used for Metal Gateways or Virtual Circuits. | [optional] 
**LocalAsn** | Pointer to **int64** |  | [optional] 
**Metro** | **string** | The UUID (or metro code) for the Metro in which to create this VRF. | 
**Name** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVrfCreateInput

`func NewVrfCreateInput(metro string, name string, ) *VrfCreateInput`

NewVrfCreateInput instantiates a new VrfCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfCreateInputWithDefaults

`func NewVrfCreateInputWithDefaults() *VrfCreateInput`

NewVrfCreateInputWithDefaults instantiates a new VrfCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpDynamicNeighborsEnabled

`func (o *VrfCreateInput) GetBgpDynamicNeighborsEnabled() bool`

GetBgpDynamicNeighborsEnabled returns the BgpDynamicNeighborsEnabled field if non-nil, zero value otherwise.

### GetBgpDynamicNeighborsEnabledOk

`func (o *VrfCreateInput) GetBgpDynamicNeighborsEnabledOk() (*bool, bool)`

GetBgpDynamicNeighborsEnabledOk returns a tuple with the BgpDynamicNeighborsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpDynamicNeighborsEnabled

`func (o *VrfCreateInput) SetBgpDynamicNeighborsEnabled(v bool)`

SetBgpDynamicNeighborsEnabled sets BgpDynamicNeighborsEnabled field to given value.

### HasBgpDynamicNeighborsEnabled

`func (o *VrfCreateInput) HasBgpDynamicNeighborsEnabled() bool`

HasBgpDynamicNeighborsEnabled returns a boolean if a field has been set.

### GetBgpDynamicNeighborsExportRouteMap

`func (o *VrfCreateInput) GetBgpDynamicNeighborsExportRouteMap() bool`

GetBgpDynamicNeighborsExportRouteMap returns the BgpDynamicNeighborsExportRouteMap field if non-nil, zero value otherwise.

### GetBgpDynamicNeighborsExportRouteMapOk

`func (o *VrfCreateInput) GetBgpDynamicNeighborsExportRouteMapOk() (*bool, bool)`

GetBgpDynamicNeighborsExportRouteMapOk returns a tuple with the BgpDynamicNeighborsExportRouteMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpDynamicNeighborsExportRouteMap

`func (o *VrfCreateInput) SetBgpDynamicNeighborsExportRouteMap(v bool)`

SetBgpDynamicNeighborsExportRouteMap sets BgpDynamicNeighborsExportRouteMap field to given value.

### HasBgpDynamicNeighborsExportRouteMap

`func (o *VrfCreateInput) HasBgpDynamicNeighborsExportRouteMap() bool`

HasBgpDynamicNeighborsExportRouteMap returns a boolean if a field has been set.

### GetBgpDynamicNeighborsBfdEnabled

`func (o *VrfCreateInput) GetBgpDynamicNeighborsBfdEnabled() bool`

GetBgpDynamicNeighborsBfdEnabled returns the BgpDynamicNeighborsBfdEnabled field if non-nil, zero value otherwise.

### GetBgpDynamicNeighborsBfdEnabledOk

`func (o *VrfCreateInput) GetBgpDynamicNeighborsBfdEnabledOk() (*bool, bool)`

GetBgpDynamicNeighborsBfdEnabledOk returns a tuple with the BgpDynamicNeighborsBfdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpDynamicNeighborsBfdEnabled

`func (o *VrfCreateInput) SetBgpDynamicNeighborsBfdEnabled(v bool)`

SetBgpDynamicNeighborsBfdEnabled sets BgpDynamicNeighborsBfdEnabled field to given value.

### HasBgpDynamicNeighborsBfdEnabled

`func (o *VrfCreateInput) HasBgpDynamicNeighborsBfdEnabled() bool`

HasBgpDynamicNeighborsBfdEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *VrfCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VrfCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VrfCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VrfCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpRanges

`func (o *VrfCreateInput) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *VrfCreateInput) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *VrfCreateInput) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *VrfCreateInput) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetLocalAsn

`func (o *VrfCreateInput) GetLocalAsn() int64`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *VrfCreateInput) GetLocalAsnOk() (*int64, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *VrfCreateInput) SetLocalAsn(v int64)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *VrfCreateInput) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetMetro

`func (o *VrfCreateInput) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *VrfCreateInput) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *VrfCreateInput) SetMetro(v string)`

SetMetro sets Metro field to given value.


### GetName

`func (o *VrfCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VrfCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VrfCreateInput) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *VrfCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VrfCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VrfCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VrfCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


