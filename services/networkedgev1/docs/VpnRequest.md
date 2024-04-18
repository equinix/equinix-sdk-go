# VpnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigName** | Pointer to **string** |  | [optional] 
**PeerIp** | **string** |  | 
**PeerSharedKey** | **string** |  | 
**RemoteAsn** | **int64** | Remote Customer ASN - Customer side | 
**RemoteIpAddress** | **string** | Remote Customer IP Address - Customer side | 
**Password** | **string** | BGP Password | 
**LocalAsn** | Pointer to **int64** | Local ASN - Equinix side | [optional] 
**TunnelIp** | **string** | Local Tunnel IP Address in CIDR format | 

## Methods

### NewVpnRequest

`func NewVpnRequest(peerIp string, peerSharedKey string, remoteAsn int64, remoteIpAddress string, password string, tunnelIp string, ) *VpnRequest`

NewVpnRequest instantiates a new VpnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnRequestWithDefaults

`func NewVpnRequestWithDefaults() *VpnRequest`

NewVpnRequestWithDefaults instantiates a new VpnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigName

`func (o *VpnRequest) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *VpnRequest) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *VpnRequest) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *VpnRequest) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### GetPeerIp

`func (o *VpnRequest) GetPeerIp() string`

GetPeerIp returns the PeerIp field if non-nil, zero value otherwise.

### GetPeerIpOk

`func (o *VpnRequest) GetPeerIpOk() (*string, bool)`

GetPeerIpOk returns a tuple with the PeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIp

`func (o *VpnRequest) SetPeerIp(v string)`

SetPeerIp sets PeerIp field to given value.


### GetPeerSharedKey

`func (o *VpnRequest) GetPeerSharedKey() string`

GetPeerSharedKey returns the PeerSharedKey field if non-nil, zero value otherwise.

### GetPeerSharedKeyOk

`func (o *VpnRequest) GetPeerSharedKeyOk() (*string, bool)`

GetPeerSharedKeyOk returns a tuple with the PeerSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSharedKey

`func (o *VpnRequest) SetPeerSharedKey(v string)`

SetPeerSharedKey sets PeerSharedKey field to given value.


### GetRemoteAsn

`func (o *VpnRequest) GetRemoteAsn() int64`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *VpnRequest) GetRemoteAsnOk() (*int64, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *VpnRequest) SetRemoteAsn(v int64)`

SetRemoteAsn sets RemoteAsn field to given value.


### GetRemoteIpAddress

`func (o *VpnRequest) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *VpnRequest) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *VpnRequest) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.


### GetPassword

`func (o *VpnRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VpnRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VpnRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetLocalAsn

`func (o *VpnRequest) GetLocalAsn() int64`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *VpnRequest) GetLocalAsnOk() (*int64, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *VpnRequest) SetLocalAsn(v int64)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *VpnRequest) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetTunnelIp

`func (o *VpnRequest) GetTunnelIp() string`

GetTunnelIp returns the TunnelIp field if non-nil, zero value otherwise.

### GetTunnelIpOk

`func (o *VpnRequest) GetTunnelIpOk() (*string, bool)`

GetTunnelIpOk returns a tuple with the TunnelIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelIp

`func (o *VpnRequest) SetTunnelIp(v string)`

SetTunnelIp sets TunnelIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


