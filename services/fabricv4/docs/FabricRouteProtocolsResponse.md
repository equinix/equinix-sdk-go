# FabricRouteProtocolsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Route Protocol URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned route protocol identifier | [optional] 
**State** | [**DeploymentState**](DeploymentState.md) |  | 
**Type** | [**RoutingProtocolBGPTypeType**](RoutingProtocolBGPTypeType.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**BgpIpv4** | [**FabricBGPConnectionIpv4**](FabricBGPConnectionIpv4.md) |  | 
**CustomerAsn** | **int64** | Customer asn | 
**BgpAuthKey** | **string** | BGP authorization key | 
**AsOverrideEnabled** | Pointer to **bool** | Enable AS number override | [optional] 
**DeploymentProperties** | Pointer to [**TopologyProperties**](TopologyProperties.md) |  | [optional] 

## Methods

### NewFabricRouteProtocolsResponse

`func NewFabricRouteProtocolsResponse(state DeploymentState, type_ RoutingProtocolBGPTypeType, bgpIpv4 FabricBGPConnectionIpv4, customerAsn int64, bgpAuthKey string, ) *FabricRouteProtocolsResponse`

NewFabricRouteProtocolsResponse instantiates a new FabricRouteProtocolsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricRouteProtocolsResponseWithDefaults

`func NewFabricRouteProtocolsResponseWithDefaults() *FabricRouteProtocolsResponse`

NewFabricRouteProtocolsResponseWithDefaults instantiates a new FabricRouteProtocolsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *FabricRouteProtocolsResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FabricRouteProtocolsResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FabricRouteProtocolsResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FabricRouteProtocolsResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *FabricRouteProtocolsResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FabricRouteProtocolsResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FabricRouteProtocolsResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FabricRouteProtocolsResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *FabricRouteProtocolsResponse) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FabricRouteProtocolsResponse) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FabricRouteProtocolsResponse) SetState(v DeploymentState)`

SetState sets State field to given value.


### GetType

`func (o *FabricRouteProtocolsResponse) GetType() RoutingProtocolBGPTypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricRouteProtocolsResponse) GetTypeOk() (*RoutingProtocolBGPTypeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricRouteProtocolsResponse) SetType(v RoutingProtocolBGPTypeType)`

SetType sets Type field to given value.


### GetName

`func (o *FabricRouteProtocolsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricRouteProtocolsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricRouteProtocolsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricRouteProtocolsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBgpIpv4

`func (o *FabricRouteProtocolsResponse) GetBgpIpv4() FabricBGPConnectionIpv4`

GetBgpIpv4 returns the BgpIpv4 field if non-nil, zero value otherwise.

### GetBgpIpv4Ok

`func (o *FabricRouteProtocolsResponse) GetBgpIpv4Ok() (*FabricBGPConnectionIpv4, bool)`

GetBgpIpv4Ok returns a tuple with the BgpIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv4

`func (o *FabricRouteProtocolsResponse) SetBgpIpv4(v FabricBGPConnectionIpv4)`

SetBgpIpv4 sets BgpIpv4 field to given value.


### GetCustomerAsn

`func (o *FabricRouteProtocolsResponse) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *FabricRouteProtocolsResponse) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *FabricRouteProtocolsResponse) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.


### GetBgpAuthKey

`func (o *FabricRouteProtocolsResponse) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *FabricRouteProtocolsResponse) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *FabricRouteProtocolsResponse) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.


### GetAsOverrideEnabled

`func (o *FabricRouteProtocolsResponse) GetAsOverrideEnabled() bool`

GetAsOverrideEnabled returns the AsOverrideEnabled field if non-nil, zero value otherwise.

### GetAsOverrideEnabledOk

`func (o *FabricRouteProtocolsResponse) GetAsOverrideEnabledOk() (*bool, bool)`

GetAsOverrideEnabledOk returns a tuple with the AsOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOverrideEnabled

`func (o *FabricRouteProtocolsResponse) SetAsOverrideEnabled(v bool)`

SetAsOverrideEnabled sets AsOverrideEnabled field to given value.

### HasAsOverrideEnabled

`func (o *FabricRouteProtocolsResponse) HasAsOverrideEnabled() bool`

HasAsOverrideEnabled returns a boolean if a field has been set.

### GetDeploymentProperties

`func (o *FabricRouteProtocolsResponse) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *FabricRouteProtocolsResponse) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *FabricRouteProtocolsResponse) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.

### HasDeploymentProperties

`func (o *FabricRouteProtocolsResponse) HasDeploymentProperties() bool`

HasDeploymentProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


