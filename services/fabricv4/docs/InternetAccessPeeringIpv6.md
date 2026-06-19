# InternetAccessPeeringIpv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** | IPv6 prefix for the peering connection | [optional] 
**PrefixLength** | Pointer to **int32** | Determines the size of subnet | [optional] 
**EquinixPeerIp** | Pointer to **string** | IPv6 peering IP address for Equinix side | [optional] 
**CustomerPeerIp** | Pointer to **string** | IPv6 peering IP address for customer side | [optional] 
**EquinixVrrpIp** | Pointer to **string** | IPv6 address for Equinix VRRP IP | [optional] 
**CustomerVrrpIp** | Pointer to **string** | IPv6 address for customer VRRP IP | [optional] 

## Methods

### NewInternetAccessPeeringIpv6

`func NewInternetAccessPeeringIpv6() *InternetAccessPeeringIpv6`

NewInternetAccessPeeringIpv6 instantiates a new InternetAccessPeeringIpv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessPeeringIpv6WithDefaults

`func NewInternetAccessPeeringIpv6WithDefaults() *InternetAccessPeeringIpv6`

NewInternetAccessPeeringIpv6WithDefaults instantiates a new InternetAccessPeeringIpv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *InternetAccessPeeringIpv6) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *InternetAccessPeeringIpv6) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *InternetAccessPeeringIpv6) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *InternetAccessPeeringIpv6) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPrefixLength

`func (o *InternetAccessPeeringIpv6) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *InternetAccessPeeringIpv6) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *InternetAccessPeeringIpv6) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *InternetAccessPeeringIpv6) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### GetEquinixPeerIp

`func (o *InternetAccessPeeringIpv6) GetEquinixPeerIp() string`

GetEquinixPeerIp returns the EquinixPeerIp field if non-nil, zero value otherwise.

### GetEquinixPeerIpOk

`func (o *InternetAccessPeeringIpv6) GetEquinixPeerIpOk() (*string, bool)`

GetEquinixPeerIpOk returns a tuple with the EquinixPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixPeerIp

`func (o *InternetAccessPeeringIpv6) SetEquinixPeerIp(v string)`

SetEquinixPeerIp sets EquinixPeerIp field to given value.

### HasEquinixPeerIp

`func (o *InternetAccessPeeringIpv6) HasEquinixPeerIp() bool`

HasEquinixPeerIp returns a boolean if a field has been set.

### GetCustomerPeerIp

`func (o *InternetAccessPeeringIpv6) GetCustomerPeerIp() string`

GetCustomerPeerIp returns the CustomerPeerIp field if non-nil, zero value otherwise.

### GetCustomerPeerIpOk

`func (o *InternetAccessPeeringIpv6) GetCustomerPeerIpOk() (*string, bool)`

GetCustomerPeerIpOk returns a tuple with the CustomerPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPeerIp

`func (o *InternetAccessPeeringIpv6) SetCustomerPeerIp(v string)`

SetCustomerPeerIp sets CustomerPeerIp field to given value.

### HasCustomerPeerIp

`func (o *InternetAccessPeeringIpv6) HasCustomerPeerIp() bool`

HasCustomerPeerIp returns a boolean if a field has been set.

### GetEquinixVrrpIp

`func (o *InternetAccessPeeringIpv6) GetEquinixVrrpIp() string`

GetEquinixVrrpIp returns the EquinixVrrpIp field if non-nil, zero value otherwise.

### GetEquinixVrrpIpOk

`func (o *InternetAccessPeeringIpv6) GetEquinixVrrpIpOk() (*string, bool)`

GetEquinixVrrpIpOk returns a tuple with the EquinixVrrpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixVrrpIp

`func (o *InternetAccessPeeringIpv6) SetEquinixVrrpIp(v string)`

SetEquinixVrrpIp sets EquinixVrrpIp field to given value.

### HasEquinixVrrpIp

`func (o *InternetAccessPeeringIpv6) HasEquinixVrrpIp() bool`

HasEquinixVrrpIp returns a boolean if a field has been set.

### GetCustomerVrrpIp

`func (o *InternetAccessPeeringIpv6) GetCustomerVrrpIp() string`

GetCustomerVrrpIp returns the CustomerVrrpIp field if non-nil, zero value otherwise.

### GetCustomerVrrpIpOk

`func (o *InternetAccessPeeringIpv6) GetCustomerVrrpIpOk() (*string, bool)`

GetCustomerVrrpIpOk returns a tuple with the CustomerVrrpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerVrrpIp

`func (o *InternetAccessPeeringIpv6) SetCustomerVrrpIp(v string)`

SetCustomerVrrpIp sets CustomerVrrpIp field to given value.

### HasCustomerVrrpIp

`func (o *InternetAccessPeeringIpv6) HasCustomerVrrpIp() bool`

HasCustomerVrrpIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


