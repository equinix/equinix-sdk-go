# DirectPeeringIpv6Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | Pointer to [**DirectPeeringIpv4RequestConnection**](DirectPeeringIpv4RequestConnection.md) |  | [optional] 
**EquinixPeerIps** | Pointer to **[]string** | Peering IP addresses in Version 6 (IPv6)  | [optional] 
**EquinixVRRPIp** | Pointer to **string** | Virtual router group IP addresses in Version 6 (IPv6)  | [optional] 

## Methods

### NewDirectPeeringIpv6Request

`func NewDirectPeeringIpv6Request() *DirectPeeringIpv6Request`

NewDirectPeeringIpv6Request instantiates a new DirectPeeringIpv6Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectPeeringIpv6RequestWithDefaults

`func NewDirectPeeringIpv6RequestWithDefaults() *DirectPeeringIpv6Request`

NewDirectPeeringIpv6RequestWithDefaults instantiates a new DirectPeeringIpv6Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *DirectPeeringIpv6Request) GetConnection() DirectPeeringIpv4RequestConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *DirectPeeringIpv6Request) GetConnectionOk() (*DirectPeeringIpv4RequestConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *DirectPeeringIpv6Request) SetConnection(v DirectPeeringIpv4RequestConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *DirectPeeringIpv6Request) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetEquinixPeerIps

`func (o *DirectPeeringIpv6Request) GetEquinixPeerIps() []string`

GetEquinixPeerIps returns the EquinixPeerIps field if non-nil, zero value otherwise.

### GetEquinixPeerIpsOk

`func (o *DirectPeeringIpv6Request) GetEquinixPeerIpsOk() (*[]string, bool)`

GetEquinixPeerIpsOk returns a tuple with the EquinixPeerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixPeerIps

`func (o *DirectPeeringIpv6Request) SetEquinixPeerIps(v []string)`

SetEquinixPeerIps sets EquinixPeerIps field to given value.

### HasEquinixPeerIps

`func (o *DirectPeeringIpv6Request) HasEquinixPeerIps() bool`

HasEquinixPeerIps returns a boolean if a field has been set.

### GetEquinixVRRPIp

`func (o *DirectPeeringIpv6Request) GetEquinixVRRPIp() string`

GetEquinixVRRPIp returns the EquinixVRRPIp field if non-nil, zero value otherwise.

### GetEquinixVRRPIpOk

`func (o *DirectPeeringIpv6Request) GetEquinixVRRPIpOk() (*string, bool)`

GetEquinixVRRPIpOk returns a tuple with the EquinixVRRPIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixVRRPIp

`func (o *DirectPeeringIpv6Request) SetEquinixVRRPIp(v string)`

SetEquinixVRRPIp sets EquinixVRRPIp field to given value.

### HasEquinixVRRPIp

`func (o *DirectPeeringIpv6Request) HasEquinixVRRPIp() bool`

HasEquinixVRRPIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


