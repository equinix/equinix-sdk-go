# Vpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteName** | **string** |  | 
**VirtualDeviceUuid** | **string** | The unique Id of the primary device. | 
**ConfigName** | Pointer to **string** |  | [optional] 
**PeerIp** | **string** |  | 
**PeerSharedKey** | **string** |  | 
**RemoteAsn** | **int64** | Remote Customer ASN - Customer side | 
**RemoteIpAddress** | **string** | Remote Customer IP Address - Customer side | 
**Password** | **string** | BGP Password | 
**LocalAsn** | Pointer to **int64** | Local ASN - Equinix side | [optional] 
**TunnelIp** | **string** | Local Tunnel IP Address in CIDR format | 
**Secondary** | Pointer to [**VpnRequest**](VpnRequest.md) |  | [optional] 

## Methods

### NewVpn

`func NewVpn(siteName string, virtualDeviceUuid string, peerIp string, peerSharedKey string, remoteAsn int64, remoteIpAddress string, password string, tunnelIp string, ) *Vpn`

NewVpn instantiates a new Vpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnWithDefaults

`func NewVpnWithDefaults() *Vpn`

NewVpnWithDefaults instantiates a new Vpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteName

`func (o *Vpn) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *Vpn) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *Vpn) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.


### GetVirtualDeviceUuid

`func (o *Vpn) GetVirtualDeviceUuid() string`

GetVirtualDeviceUuid returns the VirtualDeviceUuid field if non-nil, zero value otherwise.

### GetVirtualDeviceUuidOk

`func (o *Vpn) GetVirtualDeviceUuidOk() (*string, bool)`

GetVirtualDeviceUuidOk returns a tuple with the VirtualDeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceUuid

`func (o *Vpn) SetVirtualDeviceUuid(v string)`

SetVirtualDeviceUuid sets VirtualDeviceUuid field to given value.


### GetConfigName

`func (o *Vpn) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *Vpn) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *Vpn) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *Vpn) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### GetPeerIp

`func (o *Vpn) GetPeerIp() string`

GetPeerIp returns the PeerIp field if non-nil, zero value otherwise.

### GetPeerIpOk

`func (o *Vpn) GetPeerIpOk() (*string, bool)`

GetPeerIpOk returns a tuple with the PeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIp

`func (o *Vpn) SetPeerIp(v string)`

SetPeerIp sets PeerIp field to given value.


### GetPeerSharedKey

`func (o *Vpn) GetPeerSharedKey() string`

GetPeerSharedKey returns the PeerSharedKey field if non-nil, zero value otherwise.

### GetPeerSharedKeyOk

`func (o *Vpn) GetPeerSharedKeyOk() (*string, bool)`

GetPeerSharedKeyOk returns a tuple with the PeerSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSharedKey

`func (o *Vpn) SetPeerSharedKey(v string)`

SetPeerSharedKey sets PeerSharedKey field to given value.


### GetRemoteAsn

`func (o *Vpn) GetRemoteAsn() int64`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *Vpn) GetRemoteAsnOk() (*int64, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *Vpn) SetRemoteAsn(v int64)`

SetRemoteAsn sets RemoteAsn field to given value.


### GetRemoteIpAddress

`func (o *Vpn) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *Vpn) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *Vpn) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.


### GetPassword

`func (o *Vpn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Vpn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Vpn) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetLocalAsn

`func (o *Vpn) GetLocalAsn() int64`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *Vpn) GetLocalAsnOk() (*int64, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *Vpn) SetLocalAsn(v int64)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *Vpn) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetTunnelIp

`func (o *Vpn) GetTunnelIp() string`

GetTunnelIp returns the TunnelIp field if non-nil, zero value otherwise.

### GetTunnelIpOk

`func (o *Vpn) GetTunnelIpOk() (*string, bool)`

GetTunnelIpOk returns a tuple with the TunnelIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelIp

`func (o *Vpn) SetTunnelIp(v string)`

SetTunnelIp sets TunnelIp field to given value.


### GetSecondary

`func (o *Vpn) GetSecondary() VpnRequest`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *Vpn) GetSecondaryOk() (*VpnRequest, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *Vpn) SetSecondary(v VpnRequest)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *Vpn) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


