# BgpNeighborData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to **int32** | Address Family for IP Address. Accepted values are 4 or 6 | [optional] 
**CustomerAs** | Pointer to **int32** | The customer&#39;s ASN. In a local BGP deployment, this will be an internal ASN used to route within the data center. For a global BGP deployment, this will be the your own ASN, configured when you set up BGP for your project. | [optional] 
**CustomerIp** | Pointer to **string** | The device&#39;s IP address. For an IPv4 BGP session, this is typically the private bond0 address for the device. | [optional] 
**Md5Enabled** | Pointer to **bool** | True if an MD5 password is configured for the project. | [optional] 
**Md5Password** | Pointer to **string** | The MD5 password configured for the project, if set. | [optional] 
**Multihop** | Pointer to **bool** | True when the BGP session should be configured as multihop. | [optional] 
**PeerAs** | Pointer to **int32** | The Peer ASN to use when configuring BGP on your device. | [optional] 
**PeerIps** | Pointer to **[]string** | A list of one or more IP addresses to use for the Peer IP section of your BGP configuration. For non-multihop sessions, this will typically be a single gateway address for the device. For multihop sessions, it will be a list of IPs. | [optional] 
**RoutesIn** | Pointer to [**[]BgpRoute**](BgpRoute.md) | A list of project subnets | [optional] 
**RoutesOut** | Pointer to [**[]BgpRoute**](BgpRoute.md) | A list of outgoing routes. Only populated if the BGP session has default route enabled. | [optional] 

## Methods

### NewBgpNeighborData

`func NewBgpNeighborData() *BgpNeighborData`

NewBgpNeighborData instantiates a new BgpNeighborData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpNeighborDataWithDefaults

`func NewBgpNeighborDataWithDefaults() *BgpNeighborData`

NewBgpNeighborDataWithDefaults instantiates a new BgpNeighborData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *BgpNeighborData) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *BgpNeighborData) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *BgpNeighborData) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *BgpNeighborData) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetCustomerAs

`func (o *BgpNeighborData) GetCustomerAs() int32`

GetCustomerAs returns the CustomerAs field if non-nil, zero value otherwise.

### GetCustomerAsOk

`func (o *BgpNeighborData) GetCustomerAsOk() (*int32, bool)`

GetCustomerAsOk returns a tuple with the CustomerAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAs

`func (o *BgpNeighborData) SetCustomerAs(v int32)`

SetCustomerAs sets CustomerAs field to given value.

### HasCustomerAs

`func (o *BgpNeighborData) HasCustomerAs() bool`

HasCustomerAs returns a boolean if a field has been set.

### GetCustomerIp

`func (o *BgpNeighborData) GetCustomerIp() string`

GetCustomerIp returns the CustomerIp field if non-nil, zero value otherwise.

### GetCustomerIpOk

`func (o *BgpNeighborData) GetCustomerIpOk() (*string, bool)`

GetCustomerIpOk returns a tuple with the CustomerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIp

`func (o *BgpNeighborData) SetCustomerIp(v string)`

SetCustomerIp sets CustomerIp field to given value.

### HasCustomerIp

`func (o *BgpNeighborData) HasCustomerIp() bool`

HasCustomerIp returns a boolean if a field has been set.

### GetMd5Enabled

`func (o *BgpNeighborData) GetMd5Enabled() bool`

GetMd5Enabled returns the Md5Enabled field if non-nil, zero value otherwise.

### GetMd5EnabledOk

`func (o *BgpNeighborData) GetMd5EnabledOk() (*bool, bool)`

GetMd5EnabledOk returns a tuple with the Md5Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Enabled

`func (o *BgpNeighborData) SetMd5Enabled(v bool)`

SetMd5Enabled sets Md5Enabled field to given value.

### HasMd5Enabled

`func (o *BgpNeighborData) HasMd5Enabled() bool`

HasMd5Enabled returns a boolean if a field has been set.

### GetMd5Password

`func (o *BgpNeighborData) GetMd5Password() string`

GetMd5Password returns the Md5Password field if non-nil, zero value otherwise.

### GetMd5PasswordOk

`func (o *BgpNeighborData) GetMd5PasswordOk() (*string, bool)`

GetMd5PasswordOk returns a tuple with the Md5Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Password

`func (o *BgpNeighborData) SetMd5Password(v string)`

SetMd5Password sets Md5Password field to given value.

### HasMd5Password

`func (o *BgpNeighborData) HasMd5Password() bool`

HasMd5Password returns a boolean if a field has been set.

### GetMultihop

`func (o *BgpNeighborData) GetMultihop() bool`

GetMultihop returns the Multihop field if non-nil, zero value otherwise.

### GetMultihopOk

`func (o *BgpNeighborData) GetMultihopOk() (*bool, bool)`

GetMultihopOk returns a tuple with the Multihop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultihop

`func (o *BgpNeighborData) SetMultihop(v bool)`

SetMultihop sets Multihop field to given value.

### HasMultihop

`func (o *BgpNeighborData) HasMultihop() bool`

HasMultihop returns a boolean if a field has been set.

### GetPeerAs

`func (o *BgpNeighborData) GetPeerAs() int32`

GetPeerAs returns the PeerAs field if non-nil, zero value otherwise.

### GetPeerAsOk

`func (o *BgpNeighborData) GetPeerAsOk() (*int32, bool)`

GetPeerAsOk returns a tuple with the PeerAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAs

`func (o *BgpNeighborData) SetPeerAs(v int32)`

SetPeerAs sets PeerAs field to given value.

### HasPeerAs

`func (o *BgpNeighborData) HasPeerAs() bool`

HasPeerAs returns a boolean if a field has been set.

### GetPeerIps

`func (o *BgpNeighborData) GetPeerIps() []string`

GetPeerIps returns the PeerIps field if non-nil, zero value otherwise.

### GetPeerIpsOk

`func (o *BgpNeighborData) GetPeerIpsOk() (*[]string, bool)`

GetPeerIpsOk returns a tuple with the PeerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIps

`func (o *BgpNeighborData) SetPeerIps(v []string)`

SetPeerIps sets PeerIps field to given value.

### HasPeerIps

`func (o *BgpNeighborData) HasPeerIps() bool`

HasPeerIps returns a boolean if a field has been set.

### GetRoutesIn

`func (o *BgpNeighborData) GetRoutesIn() []BgpRoute`

GetRoutesIn returns the RoutesIn field if non-nil, zero value otherwise.

### GetRoutesInOk

`func (o *BgpNeighborData) GetRoutesInOk() (*[]BgpRoute, bool)`

GetRoutesInOk returns a tuple with the RoutesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutesIn

`func (o *BgpNeighborData) SetRoutesIn(v []BgpRoute)`

SetRoutesIn sets RoutesIn field to given value.

### HasRoutesIn

`func (o *BgpNeighborData) HasRoutesIn() bool`

HasRoutesIn returns a boolean if a field has been set.

### GetRoutesOut

`func (o *BgpNeighborData) GetRoutesOut() []BgpRoute`

GetRoutesOut returns the RoutesOut field if non-nil, zero value otherwise.

### GetRoutesOutOk

`func (o *BgpNeighborData) GetRoutesOutOk() (*[]BgpRoute, bool)`

GetRoutesOutOk returns a tuple with the RoutesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutesOut

`func (o *BgpNeighborData) SetRoutesOut(v []BgpRoute)`

SetRoutesOut sets RoutesOut field to given value.

### HasRoutesOut

`func (o *BgpNeighborData) HasRoutesOut() bool`

HasRoutesOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


