# FabricProviderResourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | IPWAN Connection URI | [optional] [readonly] 
**Type** | **string** |  | 
**Uuid** | Pointer to **string** | Equinix-assigned ipwan connection identifier | [optional] 
**State** | [**DeploymentState**](DeploymentState.md) |  | 
**Name** | **string** |  | 
**Location** | [**SimplifiedLocation**](SimplifiedLocation.md) |  | 
**Package** | [**CloudRouterPostRequestPackage**](CloudRouterPostRequestPackage.md) |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 
**Bandwidth** | **int32** |  | 
**Redundancy** | [**ConnectionRedundancy**](ConnectionRedundancy.md) |  | 
**ASide** | [**ConnectionSide**](ConnectionSide.md) |  | 
**ZSide** | [**ConnectionSide**](ConnectionSide.md) |  | 
**BgpIpv4** | [**FabricBGPConnectionIpv4**](FabricBGPConnectionIpv4.md) |  | 
**CustomerAsn** | **int64** | Customer asn | 
**BgpAuthKey** | **string** | BGP authorization key | 
**AsOverrideEnabled** | Pointer to **bool** | Enable AS number override | [optional] 
**Scope** | Pointer to [**NetworkScope**](NetworkScope.md) |  | [optional] 

## Methods

### NewFabricProviderResourceResponse

`func NewFabricProviderResourceResponse(type_ string, state DeploymentState, name string, location SimplifiedLocation, package_ CloudRouterPostRequestPackage, deploymentProperties TopologyProperties, bandwidth int32, redundancy ConnectionRedundancy, aSide ConnectionSide, zSide ConnectionSide, bgpIpv4 FabricBGPConnectionIpv4, customerAsn int64, bgpAuthKey string, ) *FabricProviderResourceResponse`

NewFabricProviderResourceResponse instantiates a new FabricProviderResourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricProviderResourceResponseWithDefaults

`func NewFabricProviderResourceResponseWithDefaults() *FabricProviderResourceResponse`

NewFabricProviderResourceResponseWithDefaults instantiates a new FabricProviderResourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *FabricProviderResourceResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FabricProviderResourceResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FabricProviderResourceResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FabricProviderResourceResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *FabricProviderResourceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricProviderResourceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricProviderResourceResponse) SetType(v string)`

SetType sets Type field to given value.


### GetUuid

`func (o *FabricProviderResourceResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FabricProviderResourceResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FabricProviderResourceResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FabricProviderResourceResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *FabricProviderResourceResponse) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FabricProviderResourceResponse) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FabricProviderResourceResponse) SetState(v DeploymentState)`

SetState sets State field to given value.


### GetName

`func (o *FabricProviderResourceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricProviderResourceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricProviderResourceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *FabricProviderResourceResponse) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FabricProviderResourceResponse) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FabricProviderResourceResponse) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.


### GetPackage

`func (o *FabricProviderResourceResponse) GetPackage() CloudRouterPostRequestPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *FabricProviderResourceResponse) GetPackageOk() (*CloudRouterPostRequestPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *FabricProviderResourceResponse) SetPackage(v CloudRouterPostRequestPackage)`

SetPackage sets Package field to given value.


### GetDeploymentProperties

`func (o *FabricProviderResourceResponse) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *FabricProviderResourceResponse) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *FabricProviderResourceResponse) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.


### GetBandwidth

`func (o *FabricProviderResourceResponse) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *FabricProviderResourceResponse) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *FabricProviderResourceResponse) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetRedundancy

`func (o *FabricProviderResourceResponse) GetRedundancy() ConnectionRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *FabricProviderResourceResponse) GetRedundancyOk() (*ConnectionRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *FabricProviderResourceResponse) SetRedundancy(v ConnectionRedundancy)`

SetRedundancy sets Redundancy field to given value.


### GetASide

`func (o *FabricProviderResourceResponse) GetASide() ConnectionSide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *FabricProviderResourceResponse) GetASideOk() (*ConnectionSide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *FabricProviderResourceResponse) SetASide(v ConnectionSide)`

SetASide sets ASide field to given value.


### GetZSide

`func (o *FabricProviderResourceResponse) GetZSide() ConnectionSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *FabricProviderResourceResponse) GetZSideOk() (*ConnectionSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *FabricProviderResourceResponse) SetZSide(v ConnectionSide)`

SetZSide sets ZSide field to given value.


### GetBgpIpv4

`func (o *FabricProviderResourceResponse) GetBgpIpv4() FabricBGPConnectionIpv4`

GetBgpIpv4 returns the BgpIpv4 field if non-nil, zero value otherwise.

### GetBgpIpv4Ok

`func (o *FabricProviderResourceResponse) GetBgpIpv4Ok() (*FabricBGPConnectionIpv4, bool)`

GetBgpIpv4Ok returns a tuple with the BgpIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv4

`func (o *FabricProviderResourceResponse) SetBgpIpv4(v FabricBGPConnectionIpv4)`

SetBgpIpv4 sets BgpIpv4 field to given value.


### GetCustomerAsn

`func (o *FabricProviderResourceResponse) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *FabricProviderResourceResponse) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *FabricProviderResourceResponse) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.


### GetBgpAuthKey

`func (o *FabricProviderResourceResponse) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *FabricProviderResourceResponse) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *FabricProviderResourceResponse) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.


### GetAsOverrideEnabled

`func (o *FabricProviderResourceResponse) GetAsOverrideEnabled() bool`

GetAsOverrideEnabled returns the AsOverrideEnabled field if non-nil, zero value otherwise.

### GetAsOverrideEnabledOk

`func (o *FabricProviderResourceResponse) GetAsOverrideEnabledOk() (*bool, bool)`

GetAsOverrideEnabledOk returns a tuple with the AsOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOverrideEnabled

`func (o *FabricProviderResourceResponse) SetAsOverrideEnabled(v bool)`

SetAsOverrideEnabled sets AsOverrideEnabled field to given value.

### HasAsOverrideEnabled

`func (o *FabricProviderResourceResponse) HasAsOverrideEnabled() bool`

HasAsOverrideEnabled returns a boolean if a field has been set.

### GetScope

`func (o *FabricProviderResourceResponse) GetScope() NetworkScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FabricProviderResourceResponse) GetScopeOk() (*NetworkScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FabricProviderResourceResponse) SetScope(v NetworkScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *FabricProviderResourceResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


