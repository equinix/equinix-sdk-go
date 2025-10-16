# FabricProviderResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 
**Package** | Pointer to [**CloudRouterPostRequestPackage**](CloudRouterPostRequestPackage.md) |  | [optional] 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 
**Bandwidth** | **int32** |  | 
**Redundancy** | [**ConnectionRedundancy**](ConnectionRedundancy.md) |  | 
**ASide** | [**ConnectionSide**](ConnectionSide.md) |  | 
**ZSide** | [**ConnectionSide**](ConnectionSide.md) |  | 
**BgpIpv4** | [**FabricBGPConnectionIpv4**](FabricBGPConnectionIpv4.md) |  | 
**CustomerAsn** | **int64** | Customer asn | 
**BgpAuthKey** | **string** | BGP authorization key | 
**AsOverrideEnabled** | Pointer to **bool** | Enable AS number override | [optional] 
**Scope** | [**NetworkScope**](NetworkScope.md) |  | 

## Methods

### NewFabricProviderResource

`func NewFabricProviderResource(type_ string, name string, deploymentProperties TopologyProperties, bandwidth int32, redundancy ConnectionRedundancy, aSide ConnectionSide, zSide ConnectionSide, bgpIpv4 FabricBGPConnectionIpv4, customerAsn int64, bgpAuthKey string, scope NetworkScope, ) *FabricProviderResource`

NewFabricProviderResource instantiates a new FabricProviderResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricProviderResourceWithDefaults

`func NewFabricProviderResourceWithDefaults() *FabricProviderResource`

NewFabricProviderResourceWithDefaults instantiates a new FabricProviderResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FabricProviderResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricProviderResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricProviderResource) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *FabricProviderResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricProviderResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricProviderResource) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *FabricProviderResource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FabricProviderResource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FabricProviderResource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FabricProviderResource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetLocation

`func (o *FabricProviderResource) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FabricProviderResource) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FabricProviderResource) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FabricProviderResource) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPackage

`func (o *FabricProviderResource) GetPackage() CloudRouterPostRequestPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *FabricProviderResource) GetPackageOk() (*CloudRouterPostRequestPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *FabricProviderResource) SetPackage(v CloudRouterPostRequestPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *FabricProviderResource) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetDeploymentProperties

`func (o *FabricProviderResource) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *FabricProviderResource) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *FabricProviderResource) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.


### GetBandwidth

`func (o *FabricProviderResource) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *FabricProviderResource) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *FabricProviderResource) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetRedundancy

`func (o *FabricProviderResource) GetRedundancy() ConnectionRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *FabricProviderResource) GetRedundancyOk() (*ConnectionRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *FabricProviderResource) SetRedundancy(v ConnectionRedundancy)`

SetRedundancy sets Redundancy field to given value.


### GetASide

`func (o *FabricProviderResource) GetASide() ConnectionSide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *FabricProviderResource) GetASideOk() (*ConnectionSide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *FabricProviderResource) SetASide(v ConnectionSide)`

SetASide sets ASide field to given value.


### GetZSide

`func (o *FabricProviderResource) GetZSide() ConnectionSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *FabricProviderResource) GetZSideOk() (*ConnectionSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *FabricProviderResource) SetZSide(v ConnectionSide)`

SetZSide sets ZSide field to given value.


### GetBgpIpv4

`func (o *FabricProviderResource) GetBgpIpv4() FabricBGPConnectionIpv4`

GetBgpIpv4 returns the BgpIpv4 field if non-nil, zero value otherwise.

### GetBgpIpv4Ok

`func (o *FabricProviderResource) GetBgpIpv4Ok() (*FabricBGPConnectionIpv4, bool)`

GetBgpIpv4Ok returns a tuple with the BgpIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv4

`func (o *FabricProviderResource) SetBgpIpv4(v FabricBGPConnectionIpv4)`

SetBgpIpv4 sets BgpIpv4 field to given value.


### GetCustomerAsn

`func (o *FabricProviderResource) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *FabricProviderResource) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *FabricProviderResource) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.


### GetBgpAuthKey

`func (o *FabricProviderResource) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *FabricProviderResource) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *FabricProviderResource) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.


### GetAsOverrideEnabled

`func (o *FabricProviderResource) GetAsOverrideEnabled() bool`

GetAsOverrideEnabled returns the AsOverrideEnabled field if non-nil, zero value otherwise.

### GetAsOverrideEnabledOk

`func (o *FabricProviderResource) GetAsOverrideEnabledOk() (*bool, bool)`

GetAsOverrideEnabledOk returns a tuple with the AsOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOverrideEnabled

`func (o *FabricProviderResource) SetAsOverrideEnabled(v bool)`

SetAsOverrideEnabled sets AsOverrideEnabled field to given value.

### HasAsOverrideEnabled

`func (o *FabricProviderResource) HasAsOverrideEnabled() bool`

HasAsOverrideEnabled returns a boolean if a field has been set.

### GetScope

`func (o *FabricProviderResource) GetScope() NetworkScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FabricProviderResource) GetScopeOk() (*NetworkScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FabricProviderResource) SetScope(v NetworkScope)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


