# RoutingProtocolPeeringIpv4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to [**[]RoutingProtocolPeeringConnectionItem**](RoutingProtocolPeeringConnectionItem.md) |  | [optional] 
**CustomerPeerIps** | Pointer to **[]string** |  | [optional] 
**PeerSubnet** | [**RoutingProtocolPeeringIpv4PeerSubnet**](RoutingProtocolPeeringIpv4PeerSubnet.md) |  | 
**VrrpEnabled** | **bool** | Indicates if VRRP is enabled. | 
**EquinixPeerIps** | **[]string** |  | 
**EquinixVRRPIp** | Pointer to **string** |  | [optional] 
**CustomerVRRPIp** | Pointer to **string** |  | [optional] 

## Methods

### NewRoutingProtocolPeeringIpv4

`func NewRoutingProtocolPeeringIpv4(peerSubnet RoutingProtocolPeeringIpv4PeerSubnet, vrrpEnabled bool, equinixPeerIps []string, ) *RoutingProtocolPeeringIpv4`

NewRoutingProtocolPeeringIpv4 instantiates a new RoutingProtocolPeeringIpv4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolPeeringIpv4WithDefaults

`func NewRoutingProtocolPeeringIpv4WithDefaults() *RoutingProtocolPeeringIpv4`

NewRoutingProtocolPeeringIpv4WithDefaults instantiates a new RoutingProtocolPeeringIpv4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *RoutingProtocolPeeringIpv4) GetConnections() []RoutingProtocolPeeringConnectionItem`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *RoutingProtocolPeeringIpv4) GetConnectionsOk() (*[]RoutingProtocolPeeringConnectionItem, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *RoutingProtocolPeeringIpv4) SetConnections(v []RoutingProtocolPeeringConnectionItem)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *RoutingProtocolPeeringIpv4) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetCustomerPeerIps

`func (o *RoutingProtocolPeeringIpv4) GetCustomerPeerIps() []string`

GetCustomerPeerIps returns the CustomerPeerIps field if non-nil, zero value otherwise.

### GetCustomerPeerIpsOk

`func (o *RoutingProtocolPeeringIpv4) GetCustomerPeerIpsOk() (*[]string, bool)`

GetCustomerPeerIpsOk returns a tuple with the CustomerPeerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPeerIps

`func (o *RoutingProtocolPeeringIpv4) SetCustomerPeerIps(v []string)`

SetCustomerPeerIps sets CustomerPeerIps field to given value.

### HasCustomerPeerIps

`func (o *RoutingProtocolPeeringIpv4) HasCustomerPeerIps() bool`

HasCustomerPeerIps returns a boolean if a field has been set.

### GetPeerSubnet

`func (o *RoutingProtocolPeeringIpv4) GetPeerSubnet() RoutingProtocolPeeringIpv4PeerSubnet`

GetPeerSubnet returns the PeerSubnet field if non-nil, zero value otherwise.

### GetPeerSubnetOk

`func (o *RoutingProtocolPeeringIpv4) GetPeerSubnetOk() (*RoutingProtocolPeeringIpv4PeerSubnet, bool)`

GetPeerSubnetOk returns a tuple with the PeerSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSubnet

`func (o *RoutingProtocolPeeringIpv4) SetPeerSubnet(v RoutingProtocolPeeringIpv4PeerSubnet)`

SetPeerSubnet sets PeerSubnet field to given value.


### GetVrrpEnabled

`func (o *RoutingProtocolPeeringIpv4) GetVrrpEnabled() bool`

GetVrrpEnabled returns the VrrpEnabled field if non-nil, zero value otherwise.

### GetVrrpEnabledOk

`func (o *RoutingProtocolPeeringIpv4) GetVrrpEnabledOk() (*bool, bool)`

GetVrrpEnabledOk returns a tuple with the VrrpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrpEnabled

`func (o *RoutingProtocolPeeringIpv4) SetVrrpEnabled(v bool)`

SetVrrpEnabled sets VrrpEnabled field to given value.


### GetEquinixPeerIps

`func (o *RoutingProtocolPeeringIpv4) GetEquinixPeerIps() []string`

GetEquinixPeerIps returns the EquinixPeerIps field if non-nil, zero value otherwise.

### GetEquinixPeerIpsOk

`func (o *RoutingProtocolPeeringIpv4) GetEquinixPeerIpsOk() (*[]string, bool)`

GetEquinixPeerIpsOk returns a tuple with the EquinixPeerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixPeerIps

`func (o *RoutingProtocolPeeringIpv4) SetEquinixPeerIps(v []string)`

SetEquinixPeerIps sets EquinixPeerIps field to given value.


### GetEquinixVRRPIp

`func (o *RoutingProtocolPeeringIpv4) GetEquinixVRRPIp() string`

GetEquinixVRRPIp returns the EquinixVRRPIp field if non-nil, zero value otherwise.

### GetEquinixVRRPIpOk

`func (o *RoutingProtocolPeeringIpv4) GetEquinixVRRPIpOk() (*string, bool)`

GetEquinixVRRPIpOk returns a tuple with the EquinixVRRPIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixVRRPIp

`func (o *RoutingProtocolPeeringIpv4) SetEquinixVRRPIp(v string)`

SetEquinixVRRPIp sets EquinixVRRPIp field to given value.

### HasEquinixVRRPIp

`func (o *RoutingProtocolPeeringIpv4) HasEquinixVRRPIp() bool`

HasEquinixVRRPIp returns a boolean if a field has been set.

### GetCustomerVRRPIp

`func (o *RoutingProtocolPeeringIpv4) GetCustomerVRRPIp() string`

GetCustomerVRRPIp returns the CustomerVRRPIp field if non-nil, zero value otherwise.

### GetCustomerVRRPIpOk

`func (o *RoutingProtocolPeeringIpv4) GetCustomerVRRPIpOk() (*string, bool)`

GetCustomerVRRPIpOk returns a tuple with the CustomerVRRPIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerVRRPIp

`func (o *RoutingProtocolPeeringIpv4) SetCustomerVRRPIp(v string)`

SetCustomerVRRPIp sets CustomerVRRPIp field to given value.

### HasCustomerVRRPIp

`func (o *RoutingProtocolPeeringIpv4) HasCustomerVRRPIp() bool`

HasCustomerVRRPIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


