# InternetAccessPeeringIpv4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** | IPv4 prefix for the peering connection | [optional] 
**PrefixLength** | Pointer to **int32** | Determines the size of subnet | [optional] 
**EquinixPeerIp** | Pointer to **string** | IPv4 peering IP address for Equinix side | [optional] 
**CustomerPeerIp** | Pointer to **string** | IPv4 peering IP address for customer side | [optional] 
**EquinixVrrpIp** | Pointer to **string** | IPv4 address for Equinix VRRP IP | [optional] 
**CustomerVrrpIp** | Pointer to **string** | IPv4 address for customer VRRP IP | [optional] 

## Methods

### NewInternetAccessPeeringIpv4

`func NewInternetAccessPeeringIpv4() *InternetAccessPeeringIpv4`

NewInternetAccessPeeringIpv4 instantiates a new InternetAccessPeeringIpv4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessPeeringIpv4WithDefaults

`func NewInternetAccessPeeringIpv4WithDefaults() *InternetAccessPeeringIpv4`

NewInternetAccessPeeringIpv4WithDefaults instantiates a new InternetAccessPeeringIpv4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *InternetAccessPeeringIpv4) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *InternetAccessPeeringIpv4) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *InternetAccessPeeringIpv4) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *InternetAccessPeeringIpv4) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPrefixLength

`func (o *InternetAccessPeeringIpv4) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *InternetAccessPeeringIpv4) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *InternetAccessPeeringIpv4) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *InternetAccessPeeringIpv4) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### GetEquinixPeerIp

`func (o *InternetAccessPeeringIpv4) GetEquinixPeerIp() string`

GetEquinixPeerIp returns the EquinixPeerIp field if non-nil, zero value otherwise.

### GetEquinixPeerIpOk

`func (o *InternetAccessPeeringIpv4) GetEquinixPeerIpOk() (*string, bool)`

GetEquinixPeerIpOk returns a tuple with the EquinixPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixPeerIp

`func (o *InternetAccessPeeringIpv4) SetEquinixPeerIp(v string)`

SetEquinixPeerIp sets EquinixPeerIp field to given value.

### HasEquinixPeerIp

`func (o *InternetAccessPeeringIpv4) HasEquinixPeerIp() bool`

HasEquinixPeerIp returns a boolean if a field has been set.

### GetCustomerPeerIp

`func (o *InternetAccessPeeringIpv4) GetCustomerPeerIp() string`

GetCustomerPeerIp returns the CustomerPeerIp field if non-nil, zero value otherwise.

### GetCustomerPeerIpOk

`func (o *InternetAccessPeeringIpv4) GetCustomerPeerIpOk() (*string, bool)`

GetCustomerPeerIpOk returns a tuple with the CustomerPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPeerIp

`func (o *InternetAccessPeeringIpv4) SetCustomerPeerIp(v string)`

SetCustomerPeerIp sets CustomerPeerIp field to given value.

### HasCustomerPeerIp

`func (o *InternetAccessPeeringIpv4) HasCustomerPeerIp() bool`

HasCustomerPeerIp returns a boolean if a field has been set.

### GetEquinixVrrpIp

`func (o *InternetAccessPeeringIpv4) GetEquinixVrrpIp() string`

GetEquinixVrrpIp returns the EquinixVrrpIp field if non-nil, zero value otherwise.

### GetEquinixVrrpIpOk

`func (o *InternetAccessPeeringIpv4) GetEquinixVrrpIpOk() (*string, bool)`

GetEquinixVrrpIpOk returns a tuple with the EquinixVrrpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixVrrpIp

`func (o *InternetAccessPeeringIpv4) SetEquinixVrrpIp(v string)`

SetEquinixVrrpIp sets EquinixVrrpIp field to given value.

### HasEquinixVrrpIp

`func (o *InternetAccessPeeringIpv4) HasEquinixVrrpIp() bool`

HasEquinixVrrpIp returns a boolean if a field has been set.

### GetCustomerVrrpIp

`func (o *InternetAccessPeeringIpv4) GetCustomerVrrpIp() string`

GetCustomerVrrpIp returns the CustomerVrrpIp field if non-nil, zero value otherwise.

### GetCustomerVrrpIpOk

`func (o *InternetAccessPeeringIpv4) GetCustomerVrrpIpOk() (*string, bool)`

GetCustomerVrrpIpOk returns a tuple with the CustomerVrrpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerVrrpIp

`func (o *InternetAccessPeeringIpv4) SetCustomerVrrpIp(v string)`

SetCustomerVrrpIp sets CustomerVrrpIp field to given value.

### HasCustomerVrrpIp

`func (o *InternetAccessPeeringIpv4) HasCustomerVrrpIp() bool`

HasCustomerVrrpIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


