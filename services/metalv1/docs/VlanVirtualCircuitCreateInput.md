# VlanVirtualCircuitCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NniVlan** | Pointer to **int32** |  | [optional] 
**ProjectId** | **string** |  | 
**Speed** | Pointer to **string** | speed can be passed as integer number representing bps speed or string (e.g. &#39;52m&#39; or &#39;100g&#39; or &#39;4 gbps&#39;) | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Vnid** | Pointer to **string** | A Virtual Network record UUID or the VNID of a Metro Virtual Network in your project (sent as integer). | [optional] 

## Methods

### NewVlanVirtualCircuitCreateInput

`func NewVlanVirtualCircuitCreateInput(projectId string, ) *VlanVirtualCircuitCreateInput`

NewVlanVirtualCircuitCreateInput instantiates a new VlanVirtualCircuitCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanVirtualCircuitCreateInputWithDefaults

`func NewVlanVirtualCircuitCreateInputWithDefaults() *VlanVirtualCircuitCreateInput`

NewVlanVirtualCircuitCreateInputWithDefaults instantiates a new VlanVirtualCircuitCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VlanVirtualCircuitCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VlanVirtualCircuitCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VlanVirtualCircuitCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VlanVirtualCircuitCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *VlanVirtualCircuitCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VlanVirtualCircuitCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VlanVirtualCircuitCreateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VlanVirtualCircuitCreateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNniVlan

`func (o *VlanVirtualCircuitCreateInput) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *VlanVirtualCircuitCreateInput) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *VlanVirtualCircuitCreateInput) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.

### HasNniVlan

`func (o *VlanVirtualCircuitCreateInput) HasNniVlan() bool`

HasNniVlan returns a boolean if a field has been set.

### GetProjectId

`func (o *VlanVirtualCircuitCreateInput) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *VlanVirtualCircuitCreateInput) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *VlanVirtualCircuitCreateInput) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSpeed

`func (o *VlanVirtualCircuitCreateInput) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VlanVirtualCircuitCreateInput) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VlanVirtualCircuitCreateInput) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VlanVirtualCircuitCreateInput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *VlanVirtualCircuitCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VlanVirtualCircuitCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VlanVirtualCircuitCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VlanVirtualCircuitCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVnid

`func (o *VlanVirtualCircuitCreateInput) GetVnid() string`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *VlanVirtualCircuitCreateInput) GetVnidOk() (*string, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *VlanVirtualCircuitCreateInput) SetVnid(v string)`

SetVnid sets Vnid field to given value.

### HasVnid

`func (o *VlanVirtualCircuitCreateInput) HasVnid() bool`

HasVnid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


