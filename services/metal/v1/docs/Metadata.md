# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**Facility** | Pointer to **string** | The facility code of the instance | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Iqn** | Pointer to **string** |  | [optional] 
**Metro** | Pointer to **string** | The metro code of the instance | [optional] 
**Network** | Pointer to [**MetadataNetwork**](MetadataNetwork.md) |  | [optional] 
**OperatingSystem** | Pointer to **map[string]interface{}** |  | [optional] 
**Plan** | Pointer to **string** | The plan slug of the instance | [optional] 
**PrivateSubnets** | Pointer to **[]string** | An array of the private subnets | [optional] 
**Reserved** | Pointer to **bool** |  | [optional] 
**Specs** | Pointer to **map[string]interface{}** | The specs of the plan version of the instance | [optional] 
**SshKeys** | Pointer to **[]string** |  | [optional] 
**SwitchShortId** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**DeviceState**](DeviceState.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Volumes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *Metadata) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Metadata) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Metadata) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *Metadata) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCustomdata

`func (o *Metadata) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *Metadata) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *Metadata) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *Metadata) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetFacility

`func (o *Metadata) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *Metadata) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *Metadata) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *Metadata) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetHostname

`func (o *Metadata) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Metadata) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Metadata) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Metadata) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *Metadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Metadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Metadata) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Metadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIqn

`func (o *Metadata) GetIqn() string`

GetIqn returns the Iqn field if non-nil, zero value otherwise.

### GetIqnOk

`func (o *Metadata) GetIqnOk() (*string, bool)`

GetIqnOk returns a tuple with the Iqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqn

`func (o *Metadata) SetIqn(v string)`

SetIqn sets Iqn field to given value.

### HasIqn

`func (o *Metadata) HasIqn() bool`

HasIqn returns a boolean if a field has been set.

### GetMetro

`func (o *Metadata) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *Metadata) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *Metadata) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *Metadata) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNetwork

`func (o *Metadata) GetNetwork() MetadataNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Metadata) GetNetworkOk() (*MetadataNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Metadata) SetNetwork(v MetadataNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Metadata) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *Metadata) GetOperatingSystem() map[string]interface{}`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Metadata) GetOperatingSystemOk() (*map[string]interface{}, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Metadata) SetOperatingSystem(v map[string]interface{})`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *Metadata) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetPlan

`func (o *Metadata) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Metadata) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Metadata) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Metadata) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *Metadata) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *Metadata) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *Metadata) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *Metadata) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetReserved

`func (o *Metadata) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *Metadata) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *Metadata) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *Metadata) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetSpecs

`func (o *Metadata) GetSpecs() map[string]interface{}`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *Metadata) GetSpecsOk() (*map[string]interface{}, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *Metadata) SetSpecs(v map[string]interface{})`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *Metadata) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetSshKeys

`func (o *Metadata) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *Metadata) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *Metadata) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *Metadata) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetSwitchShortId

`func (o *Metadata) GetSwitchShortId() string`

GetSwitchShortId returns the SwitchShortId field if non-nil, zero value otherwise.

### GetSwitchShortIdOk

`func (o *Metadata) GetSwitchShortIdOk() (*string, bool)`

GetSwitchShortIdOk returns a tuple with the SwitchShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchShortId

`func (o *Metadata) SetSwitchShortId(v string)`

SetSwitchShortId sets SwitchShortId field to given value.

### HasSwitchShortId

`func (o *Metadata) HasSwitchShortId() bool`

HasSwitchShortId returns a boolean if a field has been set.

### GetState

`func (o *Metadata) GetState() DeviceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Metadata) GetStateOk() (*DeviceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Metadata) SetState(v DeviceState)`

SetState sets State field to given value.

### HasState

`func (o *Metadata) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *Metadata) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Metadata) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Metadata) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Metadata) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVolumes

`func (o *Metadata) GetVolumes() []string`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *Metadata) GetVolumesOk() (*[]string, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *Metadata) SetVolumes(v []string)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *Metadata) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


