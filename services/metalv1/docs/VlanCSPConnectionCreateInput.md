# VlanCSPConnectionCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to **string** | The preferred email used for communication and notifications about the Equinix Fabric interconnection. Optional and defaults to the primary user email address when using a User API key or the organization owner email address when using a Project API key. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Project** | **string** |  | 
**Metro** | **string** | A Metro ID or code. When creating Fabric VCs (Metal Billed), this is where interconnection will be originating from, as we pre-authorize the use of one of our shared ports as the origin of the interconnection using A-Side service tokens. We only allow local connections for Fabric VCs (Metal Billed), so the destination location must be the same as the origin. For Fabric VCs (Fabric Billed), or shared connections, this will be the destination of the interconnection. We allow remote connections for Fabric VCs (Fabric Billed), so the origin of the interconnection can be a different metro set here. | 
**Speed** | Pointer to **string** | A interconnection speed, in bps, mbps, or gbps. For Fabric VCs, this represents the maximum speed of the interconnection. For Fabric VCs (Metal Billed), this can only be one of the following: &#39;&#39;50mbps&#39;&#39;, &#39;&#39;200mbps&#39;&#39;, &#39;&#39;500mbps&#39;&#39;, &#39;&#39;1gbps&#39;&#39;, &#39;&#39;2gbps&#39;&#39;, &#39;&#39;5gbps&#39;&#39; or &#39;&#39;10gbps&#39;&#39;, and is required for creation. For Fabric VCs (Fabric Billed), this field will always default to &#39;&#39;10gbps&#39;&#39; even if it is not provided. For example, &#39;&#39;500000000&#39;&#39;, &#39;&#39;50m&#39;&#39;, or&#39; &#39;&#39;500mbps&#39;&#39; will all work as valid inputs. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | [**VlanCSPConnectionCreateInputType**](VlanCSPConnectionCreateInputType.md) |  | 
**Vlans** | **[]int32** | A list of one or two metro-based VLANs that will be set on the virtual circuits of primary and/or secondary interconnections respectively when creating Fabric VCs. VLANs can also be set after the interconnection is created, but are required to fully activate the virtual circuits. | 
**FabricProvider** | [**AWSFabricProvider**](AWSFabricProvider.md) |  | 

## Methods

### NewVlanCSPConnectionCreateInput

`func NewVlanCSPConnectionCreateInput(name string, project string, metro string, type_ VlanCSPConnectionCreateInputType, vlans []int32, fabricProvider AWSFabricProvider, ) *VlanCSPConnectionCreateInput`

NewVlanCSPConnectionCreateInput instantiates a new VlanCSPConnectionCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanCSPConnectionCreateInputWithDefaults

`func NewVlanCSPConnectionCreateInputWithDefaults() *VlanCSPConnectionCreateInput`

NewVlanCSPConnectionCreateInputWithDefaults instantiates a new VlanCSPConnectionCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *VlanCSPConnectionCreateInput) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *VlanCSPConnectionCreateInput) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *VlanCSPConnectionCreateInput) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *VlanCSPConnectionCreateInput) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetDescription

`func (o *VlanCSPConnectionCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VlanCSPConnectionCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VlanCSPConnectionCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VlanCSPConnectionCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *VlanCSPConnectionCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VlanCSPConnectionCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VlanCSPConnectionCreateInput) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *VlanCSPConnectionCreateInput) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VlanCSPConnectionCreateInput) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VlanCSPConnectionCreateInput) SetProject(v string)`

SetProject sets Project field to given value.


### GetMetro

`func (o *VlanCSPConnectionCreateInput) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *VlanCSPConnectionCreateInput) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *VlanCSPConnectionCreateInput) SetMetro(v string)`

SetMetro sets Metro field to given value.


### GetSpeed

`func (o *VlanCSPConnectionCreateInput) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VlanCSPConnectionCreateInput) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VlanCSPConnectionCreateInput) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VlanCSPConnectionCreateInput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *VlanCSPConnectionCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VlanCSPConnectionCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VlanCSPConnectionCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VlanCSPConnectionCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *VlanCSPConnectionCreateInput) GetType() VlanCSPConnectionCreateInputType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VlanCSPConnectionCreateInput) GetTypeOk() (*VlanCSPConnectionCreateInputType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VlanCSPConnectionCreateInput) SetType(v VlanCSPConnectionCreateInputType)`

SetType sets Type field to given value.


### GetVlans

`func (o *VlanCSPConnectionCreateInput) GetVlans() []int32`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *VlanCSPConnectionCreateInput) GetVlansOk() (*[]int32, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *VlanCSPConnectionCreateInput) SetVlans(v []int32)`

SetVlans sets Vlans field to given value.


### GetFabricProvider

`func (o *VlanCSPConnectionCreateInput) GetFabricProvider() AWSFabricProvider`

GetFabricProvider returns the FabricProvider field if non-nil, zero value otherwise.

### GetFabricProviderOk

`func (o *VlanCSPConnectionCreateInput) GetFabricProviderOk() (*AWSFabricProvider, bool)`

GetFabricProviderOk returns a tuple with the FabricProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricProvider

`func (o *VlanCSPConnectionCreateInput) SetFabricProvider(v AWSFabricProvider)`

SetFabricProvider sets FabricProvider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


