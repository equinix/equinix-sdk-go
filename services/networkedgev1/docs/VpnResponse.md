# VpnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteName** | **string** |  | 
**Uuid** | Pointer to **string** |  | [optional] 
**VirtualDeviceUuid** | **string** |  | 
**ConfigName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PeerIp** | **string** |  | 
**PeerSharedKey** | **string** |  | 
**RemoteAsn** | **int64** | Remote Customer ASN - Customer side | 
**RemoteIpAddress** | **string** | Remote Customer IP Address - Customer side | 
**Password** | **string** | BGP Password | 
**LocalAsn** | Pointer to **int64** | Local ASN - Equinix side | [optional] 
**TunnelIp** | **string** | Local Tunnel IP Address in CIDR format | 
**BgpState** | Pointer to **string** |  | [optional] 
**InboundBytes** | Pointer to **string** |  | [optional] 
**InboundPackets** | Pointer to **string** |  | [optional] 
**OutboundBytes** | Pointer to **string** |  | [optional] 
**OutboundPackets** | Pointer to **string** |  | [optional] 
**TunnelStatus** | Pointer to **string** |  | [optional] 
**CustOrgId** | Pointer to **int64** |  | [optional] 
**CreatedDate** | Pointer to **string** |  | [optional] 
**CreatedByFirstName** | Pointer to **string** |  | [optional] 
**CreatedByLastName** | Pointer to **string** |  | [optional] 
**CreatedByEmail** | Pointer to **string** |  | [optional] 
**CreatedByUserKey** | Pointer to **int64** |  | [optional] 
**CreatedByAccountUcmId** | Pointer to **int64** |  | [optional] 
**CreatedByUserName** | Pointer to **string** |  | [optional] 
**CreatedByCustOrgId** | Pointer to **int64** |  | [optional] 
**CreatedByCustOrgName** | Pointer to **string** |  | [optional] 
**CreatedByUserStatus** | Pointer to **string** |  | [optional] 
**CreatedByCompanyName** | Pointer to **string** |  | [optional] 
**LastUpdatedDate** | Pointer to **string** |  | [optional] 
**UpdatedByFirstName** | Pointer to **string** |  | [optional] 
**UpdatedByLastName** | Pointer to **string** |  | [optional] 
**UpdatedByEmail** | Pointer to **string** |  | [optional] 
**UpdatedByUserKey** | Pointer to **int64** |  | [optional] 
**UpdatedByAccountUcmId** | Pointer to **int64** |  | [optional] 
**UpdatedByUserName** | Pointer to **string** |  | [optional] 
**UpdatedByCustOrgId** | Pointer to **int64** |  | [optional] 
**UpdatedByCustOrgName** | Pointer to **string** |  | [optional] 
**UpdatedByUserStatus** | Pointer to **string** |  | [optional] 
**UpdatedByCompanyName** | Pointer to **string** |  | [optional] 

## Methods

### NewVpnResponse

`func NewVpnResponse(siteName string, virtualDeviceUuid string, peerIp string, peerSharedKey string, remoteAsn int64, remoteIpAddress string, password string, tunnelIp string, ) *VpnResponse`

NewVpnResponse instantiates a new VpnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnResponseWithDefaults

`func NewVpnResponseWithDefaults() *VpnResponse`

NewVpnResponseWithDefaults instantiates a new VpnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteName

`func (o *VpnResponse) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *VpnResponse) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *VpnResponse) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.


### GetUuid

`func (o *VpnResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VpnResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VpnResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VpnResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVirtualDeviceUuid

`func (o *VpnResponse) GetVirtualDeviceUuid() string`

GetVirtualDeviceUuid returns the VirtualDeviceUuid field if non-nil, zero value otherwise.

### GetVirtualDeviceUuidOk

`func (o *VpnResponse) GetVirtualDeviceUuidOk() (*string, bool)`

GetVirtualDeviceUuidOk returns a tuple with the VirtualDeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceUuid

`func (o *VpnResponse) SetVirtualDeviceUuid(v string)`

SetVirtualDeviceUuid sets VirtualDeviceUuid field to given value.


### GetConfigName

`func (o *VpnResponse) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *VpnResponse) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *VpnResponse) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *VpnResponse) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### GetStatus

`func (o *VpnResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPeerIp

`func (o *VpnResponse) GetPeerIp() string`

GetPeerIp returns the PeerIp field if non-nil, zero value otherwise.

### GetPeerIpOk

`func (o *VpnResponse) GetPeerIpOk() (*string, bool)`

GetPeerIpOk returns a tuple with the PeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIp

`func (o *VpnResponse) SetPeerIp(v string)`

SetPeerIp sets PeerIp field to given value.


### GetPeerSharedKey

`func (o *VpnResponse) GetPeerSharedKey() string`

GetPeerSharedKey returns the PeerSharedKey field if non-nil, zero value otherwise.

### GetPeerSharedKeyOk

`func (o *VpnResponse) GetPeerSharedKeyOk() (*string, bool)`

GetPeerSharedKeyOk returns a tuple with the PeerSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSharedKey

`func (o *VpnResponse) SetPeerSharedKey(v string)`

SetPeerSharedKey sets PeerSharedKey field to given value.


### GetRemoteAsn

`func (o *VpnResponse) GetRemoteAsn() int64`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *VpnResponse) GetRemoteAsnOk() (*int64, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *VpnResponse) SetRemoteAsn(v int64)`

SetRemoteAsn sets RemoteAsn field to given value.


### GetRemoteIpAddress

`func (o *VpnResponse) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *VpnResponse) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *VpnResponse) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.


### GetPassword

`func (o *VpnResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VpnResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VpnResponse) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetLocalAsn

`func (o *VpnResponse) GetLocalAsn() int64`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *VpnResponse) GetLocalAsnOk() (*int64, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *VpnResponse) SetLocalAsn(v int64)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *VpnResponse) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetTunnelIp

`func (o *VpnResponse) GetTunnelIp() string`

GetTunnelIp returns the TunnelIp field if non-nil, zero value otherwise.

### GetTunnelIpOk

`func (o *VpnResponse) GetTunnelIpOk() (*string, bool)`

GetTunnelIpOk returns a tuple with the TunnelIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelIp

`func (o *VpnResponse) SetTunnelIp(v string)`

SetTunnelIp sets TunnelIp field to given value.


### GetBgpState

`func (o *VpnResponse) GetBgpState() string`

GetBgpState returns the BgpState field if non-nil, zero value otherwise.

### GetBgpStateOk

`func (o *VpnResponse) GetBgpStateOk() (*string, bool)`

GetBgpStateOk returns a tuple with the BgpState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpState

`func (o *VpnResponse) SetBgpState(v string)`

SetBgpState sets BgpState field to given value.

### HasBgpState

`func (o *VpnResponse) HasBgpState() bool`

HasBgpState returns a boolean if a field has been set.

### GetInboundBytes

`func (o *VpnResponse) GetInboundBytes() string`

GetInboundBytes returns the InboundBytes field if non-nil, zero value otherwise.

### GetInboundBytesOk

`func (o *VpnResponse) GetInboundBytesOk() (*string, bool)`

GetInboundBytesOk returns a tuple with the InboundBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundBytes

`func (o *VpnResponse) SetInboundBytes(v string)`

SetInboundBytes sets InboundBytes field to given value.

### HasInboundBytes

`func (o *VpnResponse) HasInboundBytes() bool`

HasInboundBytes returns a boolean if a field has been set.

### GetInboundPackets

`func (o *VpnResponse) GetInboundPackets() string`

GetInboundPackets returns the InboundPackets field if non-nil, zero value otherwise.

### GetInboundPacketsOk

`func (o *VpnResponse) GetInboundPacketsOk() (*string, bool)`

GetInboundPacketsOk returns a tuple with the InboundPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundPackets

`func (o *VpnResponse) SetInboundPackets(v string)`

SetInboundPackets sets InboundPackets field to given value.

### HasInboundPackets

`func (o *VpnResponse) HasInboundPackets() bool`

HasInboundPackets returns a boolean if a field has been set.

### GetOutboundBytes

`func (o *VpnResponse) GetOutboundBytes() string`

GetOutboundBytes returns the OutboundBytes field if non-nil, zero value otherwise.

### GetOutboundBytesOk

`func (o *VpnResponse) GetOutboundBytesOk() (*string, bool)`

GetOutboundBytesOk returns a tuple with the OutboundBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundBytes

`func (o *VpnResponse) SetOutboundBytes(v string)`

SetOutboundBytes sets OutboundBytes field to given value.

### HasOutboundBytes

`func (o *VpnResponse) HasOutboundBytes() bool`

HasOutboundBytes returns a boolean if a field has been set.

### GetOutboundPackets

`func (o *VpnResponse) GetOutboundPackets() string`

GetOutboundPackets returns the OutboundPackets field if non-nil, zero value otherwise.

### GetOutboundPacketsOk

`func (o *VpnResponse) GetOutboundPacketsOk() (*string, bool)`

GetOutboundPacketsOk returns a tuple with the OutboundPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundPackets

`func (o *VpnResponse) SetOutboundPackets(v string)`

SetOutboundPackets sets OutboundPackets field to given value.

### HasOutboundPackets

`func (o *VpnResponse) HasOutboundPackets() bool`

HasOutboundPackets returns a boolean if a field has been set.

### GetTunnelStatus

`func (o *VpnResponse) GetTunnelStatus() string`

GetTunnelStatus returns the TunnelStatus field if non-nil, zero value otherwise.

### GetTunnelStatusOk

`func (o *VpnResponse) GetTunnelStatusOk() (*string, bool)`

GetTunnelStatusOk returns a tuple with the TunnelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelStatus

`func (o *VpnResponse) SetTunnelStatus(v string)`

SetTunnelStatus sets TunnelStatus field to given value.

### HasTunnelStatus

`func (o *VpnResponse) HasTunnelStatus() bool`

HasTunnelStatus returns a boolean if a field has been set.

### GetCustOrgId

`func (o *VpnResponse) GetCustOrgId() int64`

GetCustOrgId returns the CustOrgId field if non-nil, zero value otherwise.

### GetCustOrgIdOk

`func (o *VpnResponse) GetCustOrgIdOk() (*int64, bool)`

GetCustOrgIdOk returns a tuple with the CustOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustOrgId

`func (o *VpnResponse) SetCustOrgId(v int64)`

SetCustOrgId sets CustOrgId field to given value.

### HasCustOrgId

`func (o *VpnResponse) HasCustOrgId() bool`

HasCustOrgId returns a boolean if a field has been set.

### GetCreatedDate

`func (o *VpnResponse) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *VpnResponse) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *VpnResponse) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *VpnResponse) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedByFirstName

`func (o *VpnResponse) GetCreatedByFirstName() string`

GetCreatedByFirstName returns the CreatedByFirstName field if non-nil, zero value otherwise.

### GetCreatedByFirstNameOk

`func (o *VpnResponse) GetCreatedByFirstNameOk() (*string, bool)`

GetCreatedByFirstNameOk returns a tuple with the CreatedByFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByFirstName

`func (o *VpnResponse) SetCreatedByFirstName(v string)`

SetCreatedByFirstName sets CreatedByFirstName field to given value.

### HasCreatedByFirstName

`func (o *VpnResponse) HasCreatedByFirstName() bool`

HasCreatedByFirstName returns a boolean if a field has been set.

### GetCreatedByLastName

`func (o *VpnResponse) GetCreatedByLastName() string`

GetCreatedByLastName returns the CreatedByLastName field if non-nil, zero value otherwise.

### GetCreatedByLastNameOk

`func (o *VpnResponse) GetCreatedByLastNameOk() (*string, bool)`

GetCreatedByLastNameOk returns a tuple with the CreatedByLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByLastName

`func (o *VpnResponse) SetCreatedByLastName(v string)`

SetCreatedByLastName sets CreatedByLastName field to given value.

### HasCreatedByLastName

`func (o *VpnResponse) HasCreatedByLastName() bool`

HasCreatedByLastName returns a boolean if a field has been set.

### GetCreatedByEmail

`func (o *VpnResponse) GetCreatedByEmail() string`

GetCreatedByEmail returns the CreatedByEmail field if non-nil, zero value otherwise.

### GetCreatedByEmailOk

`func (o *VpnResponse) GetCreatedByEmailOk() (*string, bool)`

GetCreatedByEmailOk returns a tuple with the CreatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByEmail

`func (o *VpnResponse) SetCreatedByEmail(v string)`

SetCreatedByEmail sets CreatedByEmail field to given value.

### HasCreatedByEmail

`func (o *VpnResponse) HasCreatedByEmail() bool`

HasCreatedByEmail returns a boolean if a field has been set.

### GetCreatedByUserKey

`func (o *VpnResponse) GetCreatedByUserKey() int64`

GetCreatedByUserKey returns the CreatedByUserKey field if non-nil, zero value otherwise.

### GetCreatedByUserKeyOk

`func (o *VpnResponse) GetCreatedByUserKeyOk() (*int64, bool)`

GetCreatedByUserKeyOk returns a tuple with the CreatedByUserKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserKey

`func (o *VpnResponse) SetCreatedByUserKey(v int64)`

SetCreatedByUserKey sets CreatedByUserKey field to given value.

### HasCreatedByUserKey

`func (o *VpnResponse) HasCreatedByUserKey() bool`

HasCreatedByUserKey returns a boolean if a field has been set.

### GetCreatedByAccountUcmId

`func (o *VpnResponse) GetCreatedByAccountUcmId() int64`

GetCreatedByAccountUcmId returns the CreatedByAccountUcmId field if non-nil, zero value otherwise.

### GetCreatedByAccountUcmIdOk

`func (o *VpnResponse) GetCreatedByAccountUcmIdOk() (*int64, bool)`

GetCreatedByAccountUcmIdOk returns a tuple with the CreatedByAccountUcmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByAccountUcmId

`func (o *VpnResponse) SetCreatedByAccountUcmId(v int64)`

SetCreatedByAccountUcmId sets CreatedByAccountUcmId field to given value.

### HasCreatedByAccountUcmId

`func (o *VpnResponse) HasCreatedByAccountUcmId() bool`

HasCreatedByAccountUcmId returns a boolean if a field has been set.

### GetCreatedByUserName

`func (o *VpnResponse) GetCreatedByUserName() string`

GetCreatedByUserName returns the CreatedByUserName field if non-nil, zero value otherwise.

### GetCreatedByUserNameOk

`func (o *VpnResponse) GetCreatedByUserNameOk() (*string, bool)`

GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserName

`func (o *VpnResponse) SetCreatedByUserName(v string)`

SetCreatedByUserName sets CreatedByUserName field to given value.

### HasCreatedByUserName

`func (o *VpnResponse) HasCreatedByUserName() bool`

HasCreatedByUserName returns a boolean if a field has been set.

### GetCreatedByCustOrgId

`func (o *VpnResponse) GetCreatedByCustOrgId() int64`

GetCreatedByCustOrgId returns the CreatedByCustOrgId field if non-nil, zero value otherwise.

### GetCreatedByCustOrgIdOk

`func (o *VpnResponse) GetCreatedByCustOrgIdOk() (*int64, bool)`

GetCreatedByCustOrgIdOk returns a tuple with the CreatedByCustOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByCustOrgId

`func (o *VpnResponse) SetCreatedByCustOrgId(v int64)`

SetCreatedByCustOrgId sets CreatedByCustOrgId field to given value.

### HasCreatedByCustOrgId

`func (o *VpnResponse) HasCreatedByCustOrgId() bool`

HasCreatedByCustOrgId returns a boolean if a field has been set.

### GetCreatedByCustOrgName

`func (o *VpnResponse) GetCreatedByCustOrgName() string`

GetCreatedByCustOrgName returns the CreatedByCustOrgName field if non-nil, zero value otherwise.

### GetCreatedByCustOrgNameOk

`func (o *VpnResponse) GetCreatedByCustOrgNameOk() (*string, bool)`

GetCreatedByCustOrgNameOk returns a tuple with the CreatedByCustOrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByCustOrgName

`func (o *VpnResponse) SetCreatedByCustOrgName(v string)`

SetCreatedByCustOrgName sets CreatedByCustOrgName field to given value.

### HasCreatedByCustOrgName

`func (o *VpnResponse) HasCreatedByCustOrgName() bool`

HasCreatedByCustOrgName returns a boolean if a field has been set.

### GetCreatedByUserStatus

`func (o *VpnResponse) GetCreatedByUserStatus() string`

GetCreatedByUserStatus returns the CreatedByUserStatus field if non-nil, zero value otherwise.

### GetCreatedByUserStatusOk

`func (o *VpnResponse) GetCreatedByUserStatusOk() (*string, bool)`

GetCreatedByUserStatusOk returns a tuple with the CreatedByUserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserStatus

`func (o *VpnResponse) SetCreatedByUserStatus(v string)`

SetCreatedByUserStatus sets CreatedByUserStatus field to given value.

### HasCreatedByUserStatus

`func (o *VpnResponse) HasCreatedByUserStatus() bool`

HasCreatedByUserStatus returns a boolean if a field has been set.

### GetCreatedByCompanyName

`func (o *VpnResponse) GetCreatedByCompanyName() string`

GetCreatedByCompanyName returns the CreatedByCompanyName field if non-nil, zero value otherwise.

### GetCreatedByCompanyNameOk

`func (o *VpnResponse) GetCreatedByCompanyNameOk() (*string, bool)`

GetCreatedByCompanyNameOk returns a tuple with the CreatedByCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByCompanyName

`func (o *VpnResponse) SetCreatedByCompanyName(v string)`

SetCreatedByCompanyName sets CreatedByCompanyName field to given value.

### HasCreatedByCompanyName

`func (o *VpnResponse) HasCreatedByCompanyName() bool`

HasCreatedByCompanyName returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *VpnResponse) GetLastUpdatedDate() string`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *VpnResponse) GetLastUpdatedDateOk() (*string, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *VpnResponse) SetLastUpdatedDate(v string)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *VpnResponse) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetUpdatedByFirstName

`func (o *VpnResponse) GetUpdatedByFirstName() string`

GetUpdatedByFirstName returns the UpdatedByFirstName field if non-nil, zero value otherwise.

### GetUpdatedByFirstNameOk

`func (o *VpnResponse) GetUpdatedByFirstNameOk() (*string, bool)`

GetUpdatedByFirstNameOk returns a tuple with the UpdatedByFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByFirstName

`func (o *VpnResponse) SetUpdatedByFirstName(v string)`

SetUpdatedByFirstName sets UpdatedByFirstName field to given value.

### HasUpdatedByFirstName

`func (o *VpnResponse) HasUpdatedByFirstName() bool`

HasUpdatedByFirstName returns a boolean if a field has been set.

### GetUpdatedByLastName

`func (o *VpnResponse) GetUpdatedByLastName() string`

GetUpdatedByLastName returns the UpdatedByLastName field if non-nil, zero value otherwise.

### GetUpdatedByLastNameOk

`func (o *VpnResponse) GetUpdatedByLastNameOk() (*string, bool)`

GetUpdatedByLastNameOk returns a tuple with the UpdatedByLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByLastName

`func (o *VpnResponse) SetUpdatedByLastName(v string)`

SetUpdatedByLastName sets UpdatedByLastName field to given value.

### HasUpdatedByLastName

`func (o *VpnResponse) HasUpdatedByLastName() bool`

HasUpdatedByLastName returns a boolean if a field has been set.

### GetUpdatedByEmail

`func (o *VpnResponse) GetUpdatedByEmail() string`

GetUpdatedByEmail returns the UpdatedByEmail field if non-nil, zero value otherwise.

### GetUpdatedByEmailOk

`func (o *VpnResponse) GetUpdatedByEmailOk() (*string, bool)`

GetUpdatedByEmailOk returns a tuple with the UpdatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByEmail

`func (o *VpnResponse) SetUpdatedByEmail(v string)`

SetUpdatedByEmail sets UpdatedByEmail field to given value.

### HasUpdatedByEmail

`func (o *VpnResponse) HasUpdatedByEmail() bool`

HasUpdatedByEmail returns a boolean if a field has been set.

### GetUpdatedByUserKey

`func (o *VpnResponse) GetUpdatedByUserKey() int64`

GetUpdatedByUserKey returns the UpdatedByUserKey field if non-nil, zero value otherwise.

### GetUpdatedByUserKeyOk

`func (o *VpnResponse) GetUpdatedByUserKeyOk() (*int64, bool)`

GetUpdatedByUserKeyOk returns a tuple with the UpdatedByUserKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserKey

`func (o *VpnResponse) SetUpdatedByUserKey(v int64)`

SetUpdatedByUserKey sets UpdatedByUserKey field to given value.

### HasUpdatedByUserKey

`func (o *VpnResponse) HasUpdatedByUserKey() bool`

HasUpdatedByUserKey returns a boolean if a field has been set.

### GetUpdatedByAccountUcmId

`func (o *VpnResponse) GetUpdatedByAccountUcmId() int64`

GetUpdatedByAccountUcmId returns the UpdatedByAccountUcmId field if non-nil, zero value otherwise.

### GetUpdatedByAccountUcmIdOk

`func (o *VpnResponse) GetUpdatedByAccountUcmIdOk() (*int64, bool)`

GetUpdatedByAccountUcmIdOk returns a tuple with the UpdatedByAccountUcmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByAccountUcmId

`func (o *VpnResponse) SetUpdatedByAccountUcmId(v int64)`

SetUpdatedByAccountUcmId sets UpdatedByAccountUcmId field to given value.

### HasUpdatedByAccountUcmId

`func (o *VpnResponse) HasUpdatedByAccountUcmId() bool`

HasUpdatedByAccountUcmId returns a boolean if a field has been set.

### GetUpdatedByUserName

`func (o *VpnResponse) GetUpdatedByUserName() string`

GetUpdatedByUserName returns the UpdatedByUserName field if non-nil, zero value otherwise.

### GetUpdatedByUserNameOk

`func (o *VpnResponse) GetUpdatedByUserNameOk() (*string, bool)`

GetUpdatedByUserNameOk returns a tuple with the UpdatedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserName

`func (o *VpnResponse) SetUpdatedByUserName(v string)`

SetUpdatedByUserName sets UpdatedByUserName field to given value.

### HasUpdatedByUserName

`func (o *VpnResponse) HasUpdatedByUserName() bool`

HasUpdatedByUserName returns a boolean if a field has been set.

### GetUpdatedByCustOrgId

`func (o *VpnResponse) GetUpdatedByCustOrgId() int64`

GetUpdatedByCustOrgId returns the UpdatedByCustOrgId field if non-nil, zero value otherwise.

### GetUpdatedByCustOrgIdOk

`func (o *VpnResponse) GetUpdatedByCustOrgIdOk() (*int64, bool)`

GetUpdatedByCustOrgIdOk returns a tuple with the UpdatedByCustOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByCustOrgId

`func (o *VpnResponse) SetUpdatedByCustOrgId(v int64)`

SetUpdatedByCustOrgId sets UpdatedByCustOrgId field to given value.

### HasUpdatedByCustOrgId

`func (o *VpnResponse) HasUpdatedByCustOrgId() bool`

HasUpdatedByCustOrgId returns a boolean if a field has been set.

### GetUpdatedByCustOrgName

`func (o *VpnResponse) GetUpdatedByCustOrgName() string`

GetUpdatedByCustOrgName returns the UpdatedByCustOrgName field if non-nil, zero value otherwise.

### GetUpdatedByCustOrgNameOk

`func (o *VpnResponse) GetUpdatedByCustOrgNameOk() (*string, bool)`

GetUpdatedByCustOrgNameOk returns a tuple with the UpdatedByCustOrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByCustOrgName

`func (o *VpnResponse) SetUpdatedByCustOrgName(v string)`

SetUpdatedByCustOrgName sets UpdatedByCustOrgName field to given value.

### HasUpdatedByCustOrgName

`func (o *VpnResponse) HasUpdatedByCustOrgName() bool`

HasUpdatedByCustOrgName returns a boolean if a field has been set.

### GetUpdatedByUserStatus

`func (o *VpnResponse) GetUpdatedByUserStatus() string`

GetUpdatedByUserStatus returns the UpdatedByUserStatus field if non-nil, zero value otherwise.

### GetUpdatedByUserStatusOk

`func (o *VpnResponse) GetUpdatedByUserStatusOk() (*string, bool)`

GetUpdatedByUserStatusOk returns a tuple with the UpdatedByUserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserStatus

`func (o *VpnResponse) SetUpdatedByUserStatus(v string)`

SetUpdatedByUserStatus sets UpdatedByUserStatus field to given value.

### HasUpdatedByUserStatus

`func (o *VpnResponse) HasUpdatedByUserStatus() bool`

HasUpdatedByUserStatus returns a boolean if a field has been set.

### GetUpdatedByCompanyName

`func (o *VpnResponse) GetUpdatedByCompanyName() string`

GetUpdatedByCompanyName returns the UpdatedByCompanyName field if non-nil, zero value otherwise.

### GetUpdatedByCompanyNameOk

`func (o *VpnResponse) GetUpdatedByCompanyNameOk() (*string, bool)`

GetUpdatedByCompanyNameOk returns a tuple with the UpdatedByCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByCompanyName

`func (o *VpnResponse) SetUpdatedByCompanyName(v string)`

SetUpdatedByCompanyName sets UpdatedByCompanyName field to given value.

### HasUpdatedByCompanyName

`func (o *VpnResponse) HasUpdatedByCompanyName() bool`

HasUpdatedByCompanyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


