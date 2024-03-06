# SharedPortVCVlanCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to **string** | The preferred email used for communication and notifications about the Equinix Fabric interconnection. Required when using a Project API key. Optional and defaults to the primary user email address when using a User API key. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Project** | **string** |  | 
**Metro** | **string** | A Metro ID or code. When creating Fabric VCs (Metal Billed), this is where interconnection will be originating from, as we pre-authorize the use of one of our shared ports as the origin of the interconnection using A-Side service tokens. We only allow local connections for Fabric VCs (Metal Billed), so the destination location must be the same as the origin. For Fabric VCs (Fabric Billed), or shared connections, this will be the destination of the interconnection. We allow remote connections for Fabric VCs (Fabric Billed), so the origin of the interconnection can be a different metro set here. | 
**Speed** | Pointer to **string** | A interconnection speed, in bps, mbps, or gbps. For Fabric VCs, this represents the maximum speed of the interconnection. For Fabric VCs (Metal Billed), this can only be one of the following: &#39;&#39;50mbps&#39;&#39;, &#39;&#39;200mbps&#39;&#39;, &#39;&#39;500mbps&#39;&#39;, &#39;&#39;1gbps&#39;&#39;, &#39;&#39;2gbps&#39;&#39;, &#39;&#39;5gbps&#39;&#39; or &#39;&#39;10gbps&#39;&#39;, and is required for creation. For Fabric VCs (Fabric Billed), this field will always default to &#39;&#39;10gbps&#39;&#39; even if it is not provided. For example, &#39;&#39;500000000&#39;&#39;, &#39;&#39;50m&#39;&#39;, or&#39; &#39;&#39;500mbps&#39;&#39; will all work as valid inputs. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | [**SharedPortVCVlanCreateInputType**](SharedPortVCVlanCreateInputType.md) |  | 
**Vlans** | **[]int32** | A list of one or two metro-based VLANs that will be set on the virtual circuits of primary and/or secondary interconnections respectively when creating Fabric VCs. VLANs can also be set after the interconnection is created, but are required to fully activate the virtual circuits. | 

## Methods

### NewSharedPortVCVlanCreateInput

`func NewSharedPortVCVlanCreateInput(name string, project string, metro string, type_ SharedPortVCVlanCreateInputType, vlans []int32, ) *SharedPortVCVlanCreateInput`

NewSharedPortVCVlanCreateInput instantiates a new SharedPortVCVlanCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedPortVCVlanCreateInputWithDefaults

`func NewSharedPortVCVlanCreateInputWithDefaults() *SharedPortVCVlanCreateInput`

NewSharedPortVCVlanCreateInputWithDefaults instantiates a new SharedPortVCVlanCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *SharedPortVCVlanCreateInput) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *SharedPortVCVlanCreateInput) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *SharedPortVCVlanCreateInput) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *SharedPortVCVlanCreateInput) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetDescription

`func (o *SharedPortVCVlanCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SharedPortVCVlanCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SharedPortVCVlanCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SharedPortVCVlanCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SharedPortVCVlanCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedPortVCVlanCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedPortVCVlanCreateInput) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *SharedPortVCVlanCreateInput) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SharedPortVCVlanCreateInput) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SharedPortVCVlanCreateInput) SetProject(v string)`

SetProject sets Project field to given value.


### GetMetro

`func (o *SharedPortVCVlanCreateInput) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *SharedPortVCVlanCreateInput) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *SharedPortVCVlanCreateInput) SetMetro(v string)`

SetMetro sets Metro field to given value.


### GetSpeed

`func (o *SharedPortVCVlanCreateInput) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *SharedPortVCVlanCreateInput) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *SharedPortVCVlanCreateInput) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *SharedPortVCVlanCreateInput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *SharedPortVCVlanCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SharedPortVCVlanCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SharedPortVCVlanCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SharedPortVCVlanCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *SharedPortVCVlanCreateInput) GetType() SharedPortVCVlanCreateInputType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SharedPortVCVlanCreateInput) GetTypeOk() (*SharedPortVCVlanCreateInputType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SharedPortVCVlanCreateInput) SetType(v SharedPortVCVlanCreateInputType)`

SetType sets Type field to given value.


### GetVlans

`func (o *SharedPortVCVlanCreateInput) GetVlans() []int32`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *SharedPortVCVlanCreateInput) GetVlansOk() (*[]int32, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *SharedPortVCVlanCreateInput) SetVlans(v []int32)`

SetVlans sets Vlans field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


