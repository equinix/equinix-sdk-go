# VlanVirtualCircuitUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **string** | Speed can be changed only if it is an interconnection on a Dedicated Port | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Vnid** | Pointer to **string** | A Virtual Network record UUID or the VNID of a Metro Virtual Network in your project. | [optional] 

## Methods

### NewVlanVirtualCircuitUpdateInput

`func NewVlanVirtualCircuitUpdateInput() *VlanVirtualCircuitUpdateInput`

NewVlanVirtualCircuitUpdateInput instantiates a new VlanVirtualCircuitUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanVirtualCircuitUpdateInputWithDefaults

`func NewVlanVirtualCircuitUpdateInputWithDefaults() *VlanVirtualCircuitUpdateInput`

NewVlanVirtualCircuitUpdateInputWithDefaults instantiates a new VlanVirtualCircuitUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VlanVirtualCircuitUpdateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VlanVirtualCircuitUpdateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VlanVirtualCircuitUpdateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VlanVirtualCircuitUpdateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *VlanVirtualCircuitUpdateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VlanVirtualCircuitUpdateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VlanVirtualCircuitUpdateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VlanVirtualCircuitUpdateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpeed

`func (o *VlanVirtualCircuitUpdateInput) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VlanVirtualCircuitUpdateInput) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VlanVirtualCircuitUpdateInput) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VlanVirtualCircuitUpdateInput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *VlanVirtualCircuitUpdateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VlanVirtualCircuitUpdateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VlanVirtualCircuitUpdateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VlanVirtualCircuitUpdateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVnid

`func (o *VlanVirtualCircuitUpdateInput) GetVnid() string`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *VlanVirtualCircuitUpdateInput) GetVnidOk() (*string, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *VlanVirtualCircuitUpdateInput) SetVnid(v string)`

SetVnid sets Vnid field to given value.

### HasVnid

`func (o *VlanVirtualCircuitUpdateInput) HasVnid() bool`

HasVnid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


