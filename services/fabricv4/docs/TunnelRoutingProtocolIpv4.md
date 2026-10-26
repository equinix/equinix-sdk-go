# TunnelRoutingProtocolIpv4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable or disable the Routing Protocol on IPv4 | [optional] 
**LocalPeerIp** | Pointer to **string** | The local IP address | [optional] 
**RemotePeerIp** | Pointer to **string** | The remote peer&#39;s IP address | [optional] 

## Methods

### NewTunnelRoutingProtocolIpv4

`func NewTunnelRoutingProtocolIpv4() *TunnelRoutingProtocolIpv4`

NewTunnelRoutingProtocolIpv4 instantiates a new TunnelRoutingProtocolIpv4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelRoutingProtocolIpv4WithDefaults

`func NewTunnelRoutingProtocolIpv4WithDefaults() *TunnelRoutingProtocolIpv4`

NewTunnelRoutingProtocolIpv4WithDefaults instantiates a new TunnelRoutingProtocolIpv4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TunnelRoutingProtocolIpv4) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TunnelRoutingProtocolIpv4) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TunnelRoutingProtocolIpv4) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TunnelRoutingProtocolIpv4) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLocalPeerIp

`func (o *TunnelRoutingProtocolIpv4) GetLocalPeerIp() string`

GetLocalPeerIp returns the LocalPeerIp field if non-nil, zero value otherwise.

### GetLocalPeerIpOk

`func (o *TunnelRoutingProtocolIpv4) GetLocalPeerIpOk() (*string, bool)`

GetLocalPeerIpOk returns a tuple with the LocalPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPeerIp

`func (o *TunnelRoutingProtocolIpv4) SetLocalPeerIp(v string)`

SetLocalPeerIp sets LocalPeerIp field to given value.

### HasLocalPeerIp

`func (o *TunnelRoutingProtocolIpv4) HasLocalPeerIp() bool`

HasLocalPeerIp returns a boolean if a field has been set.

### GetRemotePeerIp

`func (o *TunnelRoutingProtocolIpv4) GetRemotePeerIp() string`

GetRemotePeerIp returns the RemotePeerIp field if non-nil, zero value otherwise.

### GetRemotePeerIpOk

`func (o *TunnelRoutingProtocolIpv4) GetRemotePeerIpOk() (*string, bool)`

GetRemotePeerIpOk returns a tuple with the RemotePeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePeerIp

`func (o *TunnelRoutingProtocolIpv4) SetRemotePeerIp(v string)`

SetRemotePeerIp sets RemotePeerIp field to given value.

### HasRemotePeerIp

`func (o *TunnelRoutingProtocolIpv4) HasRemotePeerIp() bool`

HasRemotePeerIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


