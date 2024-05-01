# DirectPeeringIpv4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | Pointer to [**DirectPeeringIpv4RequestConnection**](DirectPeeringIpv4RequestConnection.md) |  | [optional] 
**EquinixPeerIps** | Pointer to **[]string** | Peering IP addresses in Version 4 (IPv4)  | [optional] 
**EquinixVRRPIp** | Pointer to **string** | Virtual router group IP addresses in Version 4 (IPv4)  | [optional] 

## Methods

### NewDirectPeeringIpv4Request

`func NewDirectPeeringIpv4Request() *DirectPeeringIpv4Request`

NewDirectPeeringIpv4Request instantiates a new DirectPeeringIpv4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectPeeringIpv4RequestWithDefaults

`func NewDirectPeeringIpv4RequestWithDefaults() *DirectPeeringIpv4Request`

NewDirectPeeringIpv4RequestWithDefaults instantiates a new DirectPeeringIpv4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *DirectPeeringIpv4Request) GetConnection() DirectPeeringIpv4RequestConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *DirectPeeringIpv4Request) GetConnectionOk() (*DirectPeeringIpv4RequestConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *DirectPeeringIpv4Request) SetConnection(v DirectPeeringIpv4RequestConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *DirectPeeringIpv4Request) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetEquinixPeerIps

`func (o *DirectPeeringIpv4Request) GetEquinixPeerIps() []string`

GetEquinixPeerIps returns the EquinixPeerIps field if non-nil, zero value otherwise.

### GetEquinixPeerIpsOk

`func (o *DirectPeeringIpv4Request) GetEquinixPeerIpsOk() (*[]string, bool)`

GetEquinixPeerIpsOk returns a tuple with the EquinixPeerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixPeerIps

`func (o *DirectPeeringIpv4Request) SetEquinixPeerIps(v []string)`

SetEquinixPeerIps sets EquinixPeerIps field to given value.

### HasEquinixPeerIps

`func (o *DirectPeeringIpv4Request) HasEquinixPeerIps() bool`

HasEquinixPeerIps returns a boolean if a field has been set.

### GetEquinixVRRPIp

`func (o *DirectPeeringIpv4Request) GetEquinixVRRPIp() string`

GetEquinixVRRPIp returns the EquinixVRRPIp field if non-nil, zero value otherwise.

### GetEquinixVRRPIpOk

`func (o *DirectPeeringIpv4Request) GetEquinixVRRPIpOk() (*string, bool)`

GetEquinixVRRPIpOk returns a tuple with the EquinixVRRPIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixVRRPIp

`func (o *DirectPeeringIpv4Request) SetEquinixVRRPIp(v string)`

SetEquinixVRRPIp sets EquinixVRRPIp field to given value.

### HasEquinixVRRPIp

`func (o *DirectPeeringIpv4Request) HasEquinixVRRPIp() bool`

HasEquinixVRRPIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


