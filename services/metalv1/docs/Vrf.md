# Vrf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Optional field that can be set to describe the VRF | [optional] 
**Bill** | Pointer to **bool** | True if the VRF is being billed. Usage will start when the first VRF Virtual Circuit is active, and will only stop when the VRF has been deleted. | [optional] [default to false]
**BgpDynamicNeighborsEnabled** | Pointer to **bool** | Toggle to enable the dynamic bgp neighbors feature on the VRF | [optional] 
**BgpDynamicNeighborsExportRouteMap** | Pointer to **bool** | Toggle to export the VRF route-map to the dynamic bgp neighbors | [optional] 
**BgpDynamicNeighborsBfdEnabled** | Pointer to **bool** | Toggle BFD on dynamic bgp neighbors sessions | [optional] 
**LocalAsn** | Pointer to **int64** | A 4-byte ASN associated with the VRF. | [optional] 
**VirtualCircuits** | Pointer to [**[]VrfVirtualCircuit**](VrfVirtualCircuit.md) | Virtual circuits that are in the VRF | [optional] 
**IpRanges** | Pointer to **[]string** | A list of CIDR network addresses. Like [\&quot;10.0.0.0/16\&quot;, \&quot;2001:d78::/59\&quot;]. | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Metro** | Pointer to [**Metro**](Metro.md) |  | [optional] 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVrf

`func NewVrf() *Vrf`

NewVrf instantiates a new Vrf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfWithDefaults

`func NewVrfWithDefaults() *Vrf`

NewVrfWithDefaults instantiates a new Vrf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vrf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vrf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vrf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Vrf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Vrf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vrf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vrf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vrf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Vrf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Vrf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Vrf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Vrf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBill

`func (o *Vrf) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *Vrf) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *Vrf) SetBill(v bool)`

SetBill sets Bill field to given value.

### HasBill

`func (o *Vrf) HasBill() bool`

HasBill returns a boolean if a field has been set.

### GetBgpDynamicNeighborsEnabled

`func (o *Vrf) GetBgpDynamicNeighborsEnabled() bool`

GetBgpDynamicNeighborsEnabled returns the BgpDynamicNeighborsEnabled field if non-nil, zero value otherwise.

### GetBgpDynamicNeighborsEnabledOk

`func (o *Vrf) GetBgpDynamicNeighborsEnabledOk() (*bool, bool)`

GetBgpDynamicNeighborsEnabledOk returns a tuple with the BgpDynamicNeighborsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpDynamicNeighborsEnabled

`func (o *Vrf) SetBgpDynamicNeighborsEnabled(v bool)`

SetBgpDynamicNeighborsEnabled sets BgpDynamicNeighborsEnabled field to given value.

### HasBgpDynamicNeighborsEnabled

`func (o *Vrf) HasBgpDynamicNeighborsEnabled() bool`

HasBgpDynamicNeighborsEnabled returns a boolean if a field has been set.

### GetBgpDynamicNeighborsExportRouteMap

`func (o *Vrf) GetBgpDynamicNeighborsExportRouteMap() bool`

GetBgpDynamicNeighborsExportRouteMap returns the BgpDynamicNeighborsExportRouteMap field if non-nil, zero value otherwise.

### GetBgpDynamicNeighborsExportRouteMapOk

`func (o *Vrf) GetBgpDynamicNeighborsExportRouteMapOk() (*bool, bool)`

GetBgpDynamicNeighborsExportRouteMapOk returns a tuple with the BgpDynamicNeighborsExportRouteMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpDynamicNeighborsExportRouteMap

`func (o *Vrf) SetBgpDynamicNeighborsExportRouteMap(v bool)`

SetBgpDynamicNeighborsExportRouteMap sets BgpDynamicNeighborsExportRouteMap field to given value.

### HasBgpDynamicNeighborsExportRouteMap

`func (o *Vrf) HasBgpDynamicNeighborsExportRouteMap() bool`

HasBgpDynamicNeighborsExportRouteMap returns a boolean if a field has been set.

### GetBgpDynamicNeighborsBfdEnabled

`func (o *Vrf) GetBgpDynamicNeighborsBfdEnabled() bool`

GetBgpDynamicNeighborsBfdEnabled returns the BgpDynamicNeighborsBfdEnabled field if non-nil, zero value otherwise.

### GetBgpDynamicNeighborsBfdEnabledOk

`func (o *Vrf) GetBgpDynamicNeighborsBfdEnabledOk() (*bool, bool)`

GetBgpDynamicNeighborsBfdEnabledOk returns a tuple with the BgpDynamicNeighborsBfdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpDynamicNeighborsBfdEnabled

`func (o *Vrf) SetBgpDynamicNeighborsBfdEnabled(v bool)`

SetBgpDynamicNeighborsBfdEnabled sets BgpDynamicNeighborsBfdEnabled field to given value.

### HasBgpDynamicNeighborsBfdEnabled

`func (o *Vrf) HasBgpDynamicNeighborsBfdEnabled() bool`

HasBgpDynamicNeighborsBfdEnabled returns a boolean if a field has been set.

### GetLocalAsn

`func (o *Vrf) GetLocalAsn() int64`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *Vrf) GetLocalAsnOk() (*int64, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *Vrf) SetLocalAsn(v int64)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *Vrf) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetVirtualCircuits

`func (o *Vrf) GetVirtualCircuits() []VrfVirtualCircuit`

GetVirtualCircuits returns the VirtualCircuits field if non-nil, zero value otherwise.

### GetVirtualCircuitsOk

`func (o *Vrf) GetVirtualCircuitsOk() (*[]VrfVirtualCircuit, bool)`

GetVirtualCircuitsOk returns a tuple with the VirtualCircuits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCircuits

`func (o *Vrf) SetVirtualCircuits(v []VrfVirtualCircuit)`

SetVirtualCircuits sets VirtualCircuits field to given value.

### HasVirtualCircuits

`func (o *Vrf) HasVirtualCircuits() bool`

HasVirtualCircuits returns a boolean if a field has been set.

### GetIpRanges

`func (o *Vrf) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *Vrf) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *Vrf) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *Vrf) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetProject

`func (o *Vrf) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Vrf) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Vrf) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Vrf) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetMetro

`func (o *Vrf) GetMetro() Metro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *Vrf) GetMetroOk() (*Metro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *Vrf) SetMetro(v Metro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *Vrf) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Vrf) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Vrf) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Vrf) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Vrf) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetHref

`func (o *Vrf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Vrf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Vrf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Vrf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Vrf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Vrf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Vrf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Vrf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Vrf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Vrf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Vrf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Vrf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTags

`func (o *Vrf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Vrf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Vrf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Vrf) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


