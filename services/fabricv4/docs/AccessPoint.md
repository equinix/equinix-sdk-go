# AccessPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AccessPointType**](AccessPointType.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 
**Port** | Pointer to [**SimplifiedPort**](SimplifiedPort.md) |  | [optional] 
**Profile** | Pointer to [**SimplifiedServiceProfile**](SimplifiedServiceProfile.md) |  | [optional] 
**Router** | Pointer to [**CloudRouter**](CloudRouter.md) |  | [optional] 
**LinkProtocol** | Pointer to [**SimplifiedLinkProtocol**](SimplifiedLinkProtocol.md) |  | [optional] 
**VirtualDevice** | Pointer to [**VirtualDevice**](VirtualDevice.md) |  | [optional] 
**Interface** | Pointer to [**Interface**](Interface.md) |  | [optional] 
**Network** | Pointer to [**SimplifiedNetwork**](SimplifiedNetwork.md) |  | [optional] 
**SellerRegion** | Pointer to **string** | Access point seller region | [optional] 
**PeeringType** | Pointer to [**PeeringType**](PeeringType.md) |  | [optional] 
**AuthenticationKey** | Pointer to **string** | Access point authentication key | [optional] 
**ProviderConnectionId** | Pointer to **string** | Provider assigned Connection Id | [optional] 
**VirtualNetwork** | Pointer to [**VirtualNetwork**](VirtualNetwork.md) |  | [optional] 
**Interconnection** | Pointer to [**MetalInterconnection**](MetalInterconnection.md) |  | [optional] 
**VpicInterface** | Pointer to [**VpicInterface**](VpicInterface.md) |  | [optional] 

## Methods

### NewAccessPoint

`func NewAccessPoint() *AccessPoint`

NewAccessPoint instantiates a new AccessPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPointWithDefaults

`func NewAccessPointWithDefaults() *AccessPoint`

NewAccessPointWithDefaults instantiates a new AccessPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccessPoint) GetType() AccessPointType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessPoint) GetTypeOk() (*AccessPointType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessPoint) SetType(v AccessPointType)`

SetType sets Type field to given value.

### HasType

`func (o *AccessPoint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccount

`func (o *AccessPoint) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccessPoint) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccessPoint) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AccessPoint) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLocation

`func (o *AccessPoint) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AccessPoint) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AccessPoint) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AccessPoint) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPort

`func (o *AccessPoint) GetPort() SimplifiedPort`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AccessPoint) GetPortOk() (*SimplifiedPort, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AccessPoint) SetPort(v SimplifiedPort)`

SetPort sets Port field to given value.

### HasPort

`func (o *AccessPoint) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProfile

`func (o *AccessPoint) GetProfile() SimplifiedServiceProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AccessPoint) GetProfileOk() (*SimplifiedServiceProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AccessPoint) SetProfile(v SimplifiedServiceProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AccessPoint) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetRouter

`func (o *AccessPoint) GetRouter() CloudRouter`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *AccessPoint) GetRouterOk() (*CloudRouter, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *AccessPoint) SetRouter(v CloudRouter)`

SetRouter sets Router field to given value.

### HasRouter

`func (o *AccessPoint) HasRouter() bool`

HasRouter returns a boolean if a field has been set.

### GetLinkProtocol

`func (o *AccessPoint) GetLinkProtocol() SimplifiedLinkProtocol`

GetLinkProtocol returns the LinkProtocol field if non-nil, zero value otherwise.

### GetLinkProtocolOk

`func (o *AccessPoint) GetLinkProtocolOk() (*SimplifiedLinkProtocol, bool)`

GetLinkProtocolOk returns a tuple with the LinkProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkProtocol

`func (o *AccessPoint) SetLinkProtocol(v SimplifiedLinkProtocol)`

SetLinkProtocol sets LinkProtocol field to given value.

### HasLinkProtocol

`func (o *AccessPoint) HasLinkProtocol() bool`

HasLinkProtocol returns a boolean if a field has been set.

### GetVirtualDevice

`func (o *AccessPoint) GetVirtualDevice() VirtualDevice`

GetVirtualDevice returns the VirtualDevice field if non-nil, zero value otherwise.

### GetVirtualDeviceOk

`func (o *AccessPoint) GetVirtualDeviceOk() (*VirtualDevice, bool)`

GetVirtualDeviceOk returns a tuple with the VirtualDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDevice

`func (o *AccessPoint) SetVirtualDevice(v VirtualDevice)`

SetVirtualDevice sets VirtualDevice field to given value.

### HasVirtualDevice

`func (o *AccessPoint) HasVirtualDevice() bool`

HasVirtualDevice returns a boolean if a field has been set.

### GetInterface

`func (o *AccessPoint) GetInterface() Interface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *AccessPoint) GetInterfaceOk() (*Interface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *AccessPoint) SetInterface(v Interface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *AccessPoint) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetNetwork

`func (o *AccessPoint) GetNetwork() SimplifiedNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AccessPoint) GetNetworkOk() (*SimplifiedNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AccessPoint) SetNetwork(v SimplifiedNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *AccessPoint) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSellerRegion

`func (o *AccessPoint) GetSellerRegion() string`

GetSellerRegion returns the SellerRegion field if non-nil, zero value otherwise.

### GetSellerRegionOk

`func (o *AccessPoint) GetSellerRegionOk() (*string, bool)`

GetSellerRegionOk returns a tuple with the SellerRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerRegion

`func (o *AccessPoint) SetSellerRegion(v string)`

SetSellerRegion sets SellerRegion field to given value.

### HasSellerRegion

`func (o *AccessPoint) HasSellerRegion() bool`

HasSellerRegion returns a boolean if a field has been set.

### GetPeeringType

`func (o *AccessPoint) GetPeeringType() PeeringType`

GetPeeringType returns the PeeringType field if non-nil, zero value otherwise.

### GetPeeringTypeOk

`func (o *AccessPoint) GetPeeringTypeOk() (*PeeringType, bool)`

GetPeeringTypeOk returns a tuple with the PeeringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeringType

`func (o *AccessPoint) SetPeeringType(v PeeringType)`

SetPeeringType sets PeeringType field to given value.

### HasPeeringType

`func (o *AccessPoint) HasPeeringType() bool`

HasPeeringType returns a boolean if a field has been set.

### GetAuthenticationKey

`func (o *AccessPoint) GetAuthenticationKey() string`

GetAuthenticationKey returns the AuthenticationKey field if non-nil, zero value otherwise.

### GetAuthenticationKeyOk

`func (o *AccessPoint) GetAuthenticationKeyOk() (*string, bool)`

GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKey

`func (o *AccessPoint) SetAuthenticationKey(v string)`

SetAuthenticationKey sets AuthenticationKey field to given value.

### HasAuthenticationKey

`func (o *AccessPoint) HasAuthenticationKey() bool`

HasAuthenticationKey returns a boolean if a field has been set.

### GetProviderConnectionId

`func (o *AccessPoint) GetProviderConnectionId() string`

GetProviderConnectionId returns the ProviderConnectionId field if non-nil, zero value otherwise.

### GetProviderConnectionIdOk

`func (o *AccessPoint) GetProviderConnectionIdOk() (*string, bool)`

GetProviderConnectionIdOk returns a tuple with the ProviderConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConnectionId

`func (o *AccessPoint) SetProviderConnectionId(v string)`

SetProviderConnectionId sets ProviderConnectionId field to given value.

### HasProviderConnectionId

`func (o *AccessPoint) HasProviderConnectionId() bool`

HasProviderConnectionId returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *AccessPoint) GetVirtualNetwork() VirtualNetwork`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *AccessPoint) GetVirtualNetworkOk() (*VirtualNetwork, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *AccessPoint) SetVirtualNetwork(v VirtualNetwork)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *AccessPoint) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetInterconnection

`func (o *AccessPoint) GetInterconnection() MetalInterconnection`

GetInterconnection returns the Interconnection field if non-nil, zero value otherwise.

### GetInterconnectionOk

`func (o *AccessPoint) GetInterconnectionOk() (*MetalInterconnection, bool)`

GetInterconnectionOk returns a tuple with the Interconnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnection

`func (o *AccessPoint) SetInterconnection(v MetalInterconnection)`

SetInterconnection sets Interconnection field to given value.

### HasInterconnection

`func (o *AccessPoint) HasInterconnection() bool`

HasInterconnection returns a boolean if a field has been set.

### GetVpicInterface

`func (o *AccessPoint) GetVpicInterface() VpicInterface`

GetVpicInterface returns the VpicInterface field if non-nil, zero value otherwise.

### GetVpicInterfaceOk

`func (o *AccessPoint) GetVpicInterfaceOk() (*VpicInterface, bool)`

GetVpicInterfaceOk returns a tuple with the VpicInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpicInterface

`func (o *AccessPoint) SetVpicInterface(v VpicInterface)`

SetVpicInterface sets VpicInterface field to given value.

### HasVpicInterface

`func (o *AccessPoint) HasVpicInterface() bool`

HasVpicInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


