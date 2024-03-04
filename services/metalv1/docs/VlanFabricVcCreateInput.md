# VlanFabricVcCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to **string** | The preferred email used for communication and notifications about the Equinix Fabric interconnection. Required when using a Project API key. Optional and defaults to the primary user email address when using a User API key. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FacilityId** | Pointer to **string** |  | [optional] 
**Metro** | **string** | A Metro ID or code. When creating Fabric VCs (Metal Billed), this is where interconnection will be originating from, as we pre-authorize the use of one of our shared ports as the origin of the interconnection using A-Side service tokens. We only allow local connections for Fabric VCs (Metal Billed), so the destination location must be the same as the origin. For Fabric VCs (Fabric Billed), or shared connections, this will be the destination of the interconnection. We allow remote connections for Fabric VCs (Fabric Billed), so the origin of the interconnection can be a different metro set here. | 
**Name** | **string** |  | 
**Project** | Pointer to **string** |  | [optional] 
**Redundancy** | **string** | Either &#39;primary&#39; or &#39;redundant&#39;. | 
**ServiceTokenType** | [**VlanFabricVcCreateInputServiceTokenType**](VlanFabricVcCreateInputServiceTokenType.md) |  | 
**Speed** | Pointer to **string** | A interconnection speed, in bps, mbps, or gbps. For Fabric VCs, this represents the maximum speed of the interconnection. For Fabric VCs (Metal Billed), this can only be one of the following:  &#39;&#39;50mbps&#39;&#39;, &#39;&#39;200mbps&#39;&#39;, &#39;&#39;500mbps&#39;&#39;, &#39;&#39;1gbps&#39;&#39;, &#39;&#39;2gbps&#39;&#39;, &#39;&#39;5gbps&#39;&#39; or &#39;&#39;10gbps&#39;&#39;, and is required for creation. For Fabric VCs (Fabric Billed), this field will always default to &#39;&#39;10gbps&#39;&#39; even if it is not provided. For example, &#39;&#39;500000000&#39;&#39;, &#39;&#39;50m&#39;&#39;, or&#39; &#39;&#39;500mbps&#39;&#39; will all work as valid inputs. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | [**VlanFabricVcCreateInputType**](VlanFabricVcCreateInputType.md) |  | 
**Vlans** | Pointer to **[]int32** | A list of one or two metro-based VLANs that will be set on the virtual circuits of primary and/or secondary (if redundant) interconnections respectively when creating Fabric VCs. VLANs can also be set after the interconnection is created, but are required to fully activate the virtual circuits. | [optional] 

## Methods

### NewVlanFabricVcCreateInput

`func NewVlanFabricVcCreateInput(metro string, name string, redundancy string, serviceTokenType VlanFabricVcCreateInputServiceTokenType, type_ VlanFabricVcCreateInputType, ) *VlanFabricVcCreateInput`

NewVlanFabricVcCreateInput instantiates a new VlanFabricVcCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanFabricVcCreateInputWithDefaults

`func NewVlanFabricVcCreateInputWithDefaults() *VlanFabricVcCreateInput`

NewVlanFabricVcCreateInputWithDefaults instantiates a new VlanFabricVcCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *VlanFabricVcCreateInput) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *VlanFabricVcCreateInput) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *VlanFabricVcCreateInput) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *VlanFabricVcCreateInput) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetDescription

`func (o *VlanFabricVcCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VlanFabricVcCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VlanFabricVcCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VlanFabricVcCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacilityId

`func (o *VlanFabricVcCreateInput) GetFacilityId() string`

GetFacilityId returns the FacilityId field if non-nil, zero value otherwise.

### GetFacilityIdOk

`func (o *VlanFabricVcCreateInput) GetFacilityIdOk() (*string, bool)`

GetFacilityIdOk returns a tuple with the FacilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityId

`func (o *VlanFabricVcCreateInput) SetFacilityId(v string)`

SetFacilityId sets FacilityId field to given value.

### HasFacilityId

`func (o *VlanFabricVcCreateInput) HasFacilityId() bool`

HasFacilityId returns a boolean if a field has been set.

### GetMetro

`func (o *VlanFabricVcCreateInput) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *VlanFabricVcCreateInput) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *VlanFabricVcCreateInput) SetMetro(v string)`

SetMetro sets Metro field to given value.


### GetName

`func (o *VlanFabricVcCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VlanFabricVcCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VlanFabricVcCreateInput) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *VlanFabricVcCreateInput) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VlanFabricVcCreateInput) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VlanFabricVcCreateInput) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *VlanFabricVcCreateInput) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRedundancy

`func (o *VlanFabricVcCreateInput) GetRedundancy() string`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *VlanFabricVcCreateInput) GetRedundancyOk() (*string, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *VlanFabricVcCreateInput) SetRedundancy(v string)`

SetRedundancy sets Redundancy field to given value.


### GetServiceTokenType

`func (o *VlanFabricVcCreateInput) GetServiceTokenType() VlanFabricVcCreateInputServiceTokenType`

GetServiceTokenType returns the ServiceTokenType field if non-nil, zero value otherwise.

### GetServiceTokenTypeOk

`func (o *VlanFabricVcCreateInput) GetServiceTokenTypeOk() (*VlanFabricVcCreateInputServiceTokenType, bool)`

GetServiceTokenTypeOk returns a tuple with the ServiceTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenType

`func (o *VlanFabricVcCreateInput) SetServiceTokenType(v VlanFabricVcCreateInputServiceTokenType)`

SetServiceTokenType sets ServiceTokenType field to given value.


### GetSpeed

`func (o *VlanFabricVcCreateInput) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VlanFabricVcCreateInput) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VlanFabricVcCreateInput) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VlanFabricVcCreateInput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *VlanFabricVcCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VlanFabricVcCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VlanFabricVcCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VlanFabricVcCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *VlanFabricVcCreateInput) GetType() VlanFabricVcCreateInputType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VlanFabricVcCreateInput) GetTypeOk() (*VlanFabricVcCreateInputType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VlanFabricVcCreateInput) SetType(v VlanFabricVcCreateInputType)`

SetType sets Type field to given value.


### GetVlans

`func (o *VlanFabricVcCreateInput) GetVlans() []int32`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *VlanFabricVcCreateInput) GetVlansOk() (*[]int32, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *VlanFabricVcCreateInput) SetVlans(v []int32)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *VlanFabricVcCreateInput) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


