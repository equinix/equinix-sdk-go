# VirtualNetworkCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Metro** | Pointer to **string** | The UUID (or metro code) for the Metro in which to create this Virtual Network. | [optional] 
**Vxlan** | Pointer to **int32** | VLAN ID between 2-3999. Must be unique for the project within the Metro in which this Virtual Network is being created. If no value is specified, the next-available VLAN ID in the range 1000-1999 will be automatically selected. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVirtualNetworkCreateInput

`func NewVirtualNetworkCreateInput() *VirtualNetworkCreateInput`

NewVirtualNetworkCreateInput instantiates a new VirtualNetworkCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkCreateInputWithDefaults

`func NewVirtualNetworkCreateInputWithDefaults() *VirtualNetworkCreateInput`

NewVirtualNetworkCreateInputWithDefaults instantiates a new VirtualNetworkCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VirtualNetworkCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualNetworkCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualNetworkCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualNetworkCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetro

`func (o *VirtualNetworkCreateInput) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *VirtualNetworkCreateInput) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *VirtualNetworkCreateInput) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *VirtualNetworkCreateInput) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetVxlan

`func (o *VirtualNetworkCreateInput) GetVxlan() int32`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *VirtualNetworkCreateInput) GetVxlanOk() (*int32, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *VirtualNetworkCreateInput) SetVxlan(v int32)`

SetVxlan sets Vxlan field to given value.

### HasVxlan

`func (o *VirtualNetworkCreateInput) HasVxlan() bool`

HasVxlan returns a boolean if a field has been set.

### GetTags

`func (o *VirtualNetworkCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualNetworkCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualNetworkCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualNetworkCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


