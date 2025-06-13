# FabricRoutingProtocolBGPType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RoutingProtocolBGPTypeType**](RoutingProtocolBGPTypeType.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned route protocol identifier | [optional] 
**BgpIpv4** | [**FabricBGPConnectionIpv4**](FabricBGPConnectionIpv4.md) |  | 
**CustomerAsn** | **int64** | Customer asn | 
**BgpAuthKey** | **string** | BGP authorization key | 
**AsOverrideEnabled** | Pointer to **bool** | Enable AS number override | [optional] 
**DeploymentProperties** | Pointer to [**TopologyProperties**](TopologyProperties.md) |  | [optional] 

## Methods

### NewFabricRoutingProtocolBGPType

`func NewFabricRoutingProtocolBGPType(type_ RoutingProtocolBGPTypeType, bgpIpv4 FabricBGPConnectionIpv4, customerAsn int64, bgpAuthKey string, ) *FabricRoutingProtocolBGPType`

NewFabricRoutingProtocolBGPType instantiates a new FabricRoutingProtocolBGPType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricRoutingProtocolBGPTypeWithDefaults

`func NewFabricRoutingProtocolBGPTypeWithDefaults() *FabricRoutingProtocolBGPType`

NewFabricRoutingProtocolBGPTypeWithDefaults instantiates a new FabricRoutingProtocolBGPType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FabricRoutingProtocolBGPType) GetType() RoutingProtocolBGPTypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricRoutingProtocolBGPType) GetTypeOk() (*RoutingProtocolBGPTypeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricRoutingProtocolBGPType) SetType(v RoutingProtocolBGPTypeType)`

SetType sets Type field to given value.


### GetName

`func (o *FabricRoutingProtocolBGPType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricRoutingProtocolBGPType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricRoutingProtocolBGPType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricRoutingProtocolBGPType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *FabricRoutingProtocolBGPType) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FabricRoutingProtocolBGPType) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FabricRoutingProtocolBGPType) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FabricRoutingProtocolBGPType) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetBgpIpv4

`func (o *FabricRoutingProtocolBGPType) GetBgpIpv4() FabricBGPConnectionIpv4`

GetBgpIpv4 returns the BgpIpv4 field if non-nil, zero value otherwise.

### GetBgpIpv4Ok

`func (o *FabricRoutingProtocolBGPType) GetBgpIpv4Ok() (*FabricBGPConnectionIpv4, bool)`

GetBgpIpv4Ok returns a tuple with the BgpIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv4

`func (o *FabricRoutingProtocolBGPType) SetBgpIpv4(v FabricBGPConnectionIpv4)`

SetBgpIpv4 sets BgpIpv4 field to given value.


### GetCustomerAsn

`func (o *FabricRoutingProtocolBGPType) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *FabricRoutingProtocolBGPType) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *FabricRoutingProtocolBGPType) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.


### GetBgpAuthKey

`func (o *FabricRoutingProtocolBGPType) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *FabricRoutingProtocolBGPType) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *FabricRoutingProtocolBGPType) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.


### GetAsOverrideEnabled

`func (o *FabricRoutingProtocolBGPType) GetAsOverrideEnabled() bool`

GetAsOverrideEnabled returns the AsOverrideEnabled field if non-nil, zero value otherwise.

### GetAsOverrideEnabledOk

`func (o *FabricRoutingProtocolBGPType) GetAsOverrideEnabledOk() (*bool, bool)`

GetAsOverrideEnabledOk returns a tuple with the AsOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOverrideEnabled

`func (o *FabricRoutingProtocolBGPType) SetAsOverrideEnabled(v bool)`

SetAsOverrideEnabled sets AsOverrideEnabled field to given value.

### HasAsOverrideEnabled

`func (o *FabricRoutingProtocolBGPType) HasAsOverrideEnabled() bool`

HasAsOverrideEnabled returns a boolean if a field has been set.

### GetDeploymentProperties

`func (o *FabricRoutingProtocolBGPType) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *FabricRoutingProtocolBGPType) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *FabricRoutingProtocolBGPType) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.

### HasDeploymentProperties

`func (o *FabricRoutingProtocolBGPType) HasDeploymentProperties() bool`

HasDeploymentProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


