# VirtualNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedTo** | Pointer to [**Project**](Project.md) |  | [optional] 
**AssignedToVirtualCircuit** | Pointer to **bool** | True if the virtual network is attached to a virtual circuit. False if not. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to [**Href**](Href.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Instances** | Pointer to [**[]Device**](Device.md) | A list of instances with ports currently associated to this Virtual Network. | [optional] 
**MetalGateways** | Pointer to [**[]MetalGatewayLite**](MetalGatewayLite.md) | A list of metal gateways currently associated to this Virtual Network. | [optional] 
**Metro** | Pointer to [**Metro**](Metro.md) |  | [optional] 
**MetroCode** | Pointer to **string** | The Metro code of the metro in which this Virtual Network is defined. | [optional] 
**Vxlan** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVirtualNetwork

`func NewVirtualNetwork() *VirtualNetwork`

NewVirtualNetwork instantiates a new VirtualNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkWithDefaults

`func NewVirtualNetworkWithDefaults() *VirtualNetwork`

NewVirtualNetworkWithDefaults instantiates a new VirtualNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedTo

`func (o *VirtualNetwork) GetAssignedTo() Project`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *VirtualNetwork) GetAssignedToOk() (*Project, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *VirtualNetwork) SetAssignedTo(v Project)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *VirtualNetwork) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetAssignedToVirtualCircuit

`func (o *VirtualNetwork) GetAssignedToVirtualCircuit() bool`

GetAssignedToVirtualCircuit returns the AssignedToVirtualCircuit field if non-nil, zero value otherwise.

### GetAssignedToVirtualCircuitOk

`func (o *VirtualNetwork) GetAssignedToVirtualCircuitOk() (*bool, bool)`

GetAssignedToVirtualCircuitOk returns a tuple with the AssignedToVirtualCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToVirtualCircuit

`func (o *VirtualNetwork) SetAssignedToVirtualCircuit(v bool)`

SetAssignedToVirtualCircuit sets AssignedToVirtualCircuit field to given value.

### HasAssignedToVirtualCircuit

`func (o *VirtualNetwork) HasAssignedToVirtualCircuit() bool`

HasAssignedToVirtualCircuit returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *VirtualNetwork) GetFacility() Href`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *VirtualNetwork) GetFacilityOk() (*Href, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *VirtualNetwork) SetFacility(v Href)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *VirtualNetwork) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetHref

`func (o *VirtualNetwork) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VirtualNetwork) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VirtualNetwork) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *VirtualNetwork) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VirtualNetwork) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VirtualNetwork) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VirtualNetwork) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VirtualNetwork) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *VirtualNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualNetwork) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstances

`func (o *VirtualNetwork) GetInstances() []Device`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *VirtualNetwork) GetInstancesOk() (*[]Device, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *VirtualNetwork) SetInstances(v []Device)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *VirtualNetwork) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMetalGateways

`func (o *VirtualNetwork) GetMetalGateways() []MetalGatewayLite`

GetMetalGateways returns the MetalGateways field if non-nil, zero value otherwise.

### GetMetalGatewaysOk

`func (o *VirtualNetwork) GetMetalGatewaysOk() (*[]MetalGatewayLite, bool)`

GetMetalGatewaysOk returns a tuple with the MetalGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateways

`func (o *VirtualNetwork) SetMetalGateways(v []MetalGatewayLite)`

SetMetalGateways sets MetalGateways field to given value.

### HasMetalGateways

`func (o *VirtualNetwork) HasMetalGateways() bool`

HasMetalGateways returns a boolean if a field has been set.

### GetMetro

`func (o *VirtualNetwork) GetMetro() Metro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *VirtualNetwork) GetMetroOk() (*Metro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *VirtualNetwork) SetMetro(v Metro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *VirtualNetwork) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetMetroCode

`func (o *VirtualNetwork) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *VirtualNetwork) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *VirtualNetwork) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.

### HasMetroCode

`func (o *VirtualNetwork) HasMetroCode() bool`

HasMetroCode returns a boolean if a field has been set.

### GetVxlan

`func (o *VirtualNetwork) GetVxlan() int32`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *VirtualNetwork) GetVxlanOk() (*int32, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *VirtualNetwork) SetVxlan(v int32)`

SetVxlan sets Vxlan field to given value.

### HasVxlan

`func (o *VirtualNetwork) HasVxlan() bool`

HasVxlan returns a boolean if a field has been set.

### GetTags

`func (o *VirtualNetwork) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualNetwork) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualNetwork) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualNetwork) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


