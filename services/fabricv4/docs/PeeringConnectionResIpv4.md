# PeeringConnectionResIpv4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerPeerIp** | Pointer to **string** | Customer peer IPv4 address | [optional] 
**PrimaryEquinixPeerIp** | Pointer to **string** | Primary Equinix peer IPv4 address | [optional] 
**SecondaryEquinixPeerIp** | Pointer to **string** | Secondary Equinix peer IPv4 address | [optional] 
**ReverseDnsAddress** | Pointer to **string** | Reverse DNS address for the BGP session | [optional] 
**AsSet** | Pointer to **string** | Autonomous System Set for the BGP session | [optional] 
**MlpeEnabled** | Pointer to **bool** | Whether MLPE is enabled for the BGP session | [optional] 
**AuthKeys** | Pointer to [**[]PeeringConnectionResIpv4AuthKeys**](PeeringConnectionResIpv4AuthKeys.md) |  | [optional] 
**IpPrefixes** | Pointer to **[]string** | List of IP prefixes for the BGP session | [optional] 
**Enabled** | Pointer to **bool** | Whether BGP IPv4 is enabled | [optional] 

## Methods

### NewPeeringConnectionResIpv4

`func NewPeeringConnectionResIpv4() *PeeringConnectionResIpv4`

NewPeeringConnectionResIpv4 instantiates a new PeeringConnectionResIpv4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeeringConnectionResIpv4WithDefaults

`func NewPeeringConnectionResIpv4WithDefaults() *PeeringConnectionResIpv4`

NewPeeringConnectionResIpv4WithDefaults instantiates a new PeeringConnectionResIpv4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerPeerIp

`func (o *PeeringConnectionResIpv4) GetCustomerPeerIp() string`

GetCustomerPeerIp returns the CustomerPeerIp field if non-nil, zero value otherwise.

### GetCustomerPeerIpOk

`func (o *PeeringConnectionResIpv4) GetCustomerPeerIpOk() (*string, bool)`

GetCustomerPeerIpOk returns a tuple with the CustomerPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPeerIp

`func (o *PeeringConnectionResIpv4) SetCustomerPeerIp(v string)`

SetCustomerPeerIp sets CustomerPeerIp field to given value.

### HasCustomerPeerIp

`func (o *PeeringConnectionResIpv4) HasCustomerPeerIp() bool`

HasCustomerPeerIp returns a boolean if a field has been set.

### GetPrimaryEquinixPeerIp

`func (o *PeeringConnectionResIpv4) GetPrimaryEquinixPeerIp() string`

GetPrimaryEquinixPeerIp returns the PrimaryEquinixPeerIp field if non-nil, zero value otherwise.

### GetPrimaryEquinixPeerIpOk

`func (o *PeeringConnectionResIpv4) GetPrimaryEquinixPeerIpOk() (*string, bool)`

GetPrimaryEquinixPeerIpOk returns a tuple with the PrimaryEquinixPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEquinixPeerIp

`func (o *PeeringConnectionResIpv4) SetPrimaryEquinixPeerIp(v string)`

SetPrimaryEquinixPeerIp sets PrimaryEquinixPeerIp field to given value.

### HasPrimaryEquinixPeerIp

`func (o *PeeringConnectionResIpv4) HasPrimaryEquinixPeerIp() bool`

HasPrimaryEquinixPeerIp returns a boolean if a field has been set.

### GetSecondaryEquinixPeerIp

`func (o *PeeringConnectionResIpv4) GetSecondaryEquinixPeerIp() string`

GetSecondaryEquinixPeerIp returns the SecondaryEquinixPeerIp field if non-nil, zero value otherwise.

### GetSecondaryEquinixPeerIpOk

`func (o *PeeringConnectionResIpv4) GetSecondaryEquinixPeerIpOk() (*string, bool)`

GetSecondaryEquinixPeerIpOk returns a tuple with the SecondaryEquinixPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryEquinixPeerIp

`func (o *PeeringConnectionResIpv4) SetSecondaryEquinixPeerIp(v string)`

SetSecondaryEquinixPeerIp sets SecondaryEquinixPeerIp field to given value.

### HasSecondaryEquinixPeerIp

`func (o *PeeringConnectionResIpv4) HasSecondaryEquinixPeerIp() bool`

HasSecondaryEquinixPeerIp returns a boolean if a field has been set.

### GetReverseDnsAddress

`func (o *PeeringConnectionResIpv4) GetReverseDnsAddress() string`

GetReverseDnsAddress returns the ReverseDnsAddress field if non-nil, zero value otherwise.

### GetReverseDnsAddressOk

`func (o *PeeringConnectionResIpv4) GetReverseDnsAddressOk() (*string, bool)`

GetReverseDnsAddressOk returns a tuple with the ReverseDnsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseDnsAddress

`func (o *PeeringConnectionResIpv4) SetReverseDnsAddress(v string)`

SetReverseDnsAddress sets ReverseDnsAddress field to given value.

### HasReverseDnsAddress

`func (o *PeeringConnectionResIpv4) HasReverseDnsAddress() bool`

HasReverseDnsAddress returns a boolean if a field has been set.

### GetAsSet

`func (o *PeeringConnectionResIpv4) GetAsSet() string`

GetAsSet returns the AsSet field if non-nil, zero value otherwise.

### GetAsSetOk

`func (o *PeeringConnectionResIpv4) GetAsSetOk() (*string, bool)`

GetAsSetOk returns a tuple with the AsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsSet

`func (o *PeeringConnectionResIpv4) SetAsSet(v string)`

SetAsSet sets AsSet field to given value.

### HasAsSet

`func (o *PeeringConnectionResIpv4) HasAsSet() bool`

HasAsSet returns a boolean if a field has been set.

### GetMlpeEnabled

`func (o *PeeringConnectionResIpv4) GetMlpeEnabled() bool`

GetMlpeEnabled returns the MlpeEnabled field if non-nil, zero value otherwise.

### GetMlpeEnabledOk

`func (o *PeeringConnectionResIpv4) GetMlpeEnabledOk() (*bool, bool)`

GetMlpeEnabledOk returns a tuple with the MlpeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlpeEnabled

`func (o *PeeringConnectionResIpv4) SetMlpeEnabled(v bool)`

SetMlpeEnabled sets MlpeEnabled field to given value.

### HasMlpeEnabled

`func (o *PeeringConnectionResIpv4) HasMlpeEnabled() bool`

HasMlpeEnabled returns a boolean if a field has been set.

### GetAuthKeys

`func (o *PeeringConnectionResIpv4) GetAuthKeys() []PeeringConnectionResIpv4AuthKeys`

GetAuthKeys returns the AuthKeys field if non-nil, zero value otherwise.

### GetAuthKeysOk

`func (o *PeeringConnectionResIpv4) GetAuthKeysOk() (*[]PeeringConnectionResIpv4AuthKeys, bool)`

GetAuthKeysOk returns a tuple with the AuthKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKeys

`func (o *PeeringConnectionResIpv4) SetAuthKeys(v []PeeringConnectionResIpv4AuthKeys)`

SetAuthKeys sets AuthKeys field to given value.

### HasAuthKeys

`func (o *PeeringConnectionResIpv4) HasAuthKeys() bool`

HasAuthKeys returns a boolean if a field has been set.

### GetIpPrefixes

`func (o *PeeringConnectionResIpv4) GetIpPrefixes() []string`

GetIpPrefixes returns the IpPrefixes field if non-nil, zero value otherwise.

### GetIpPrefixesOk

`func (o *PeeringConnectionResIpv4) GetIpPrefixesOk() (*[]string, bool)`

GetIpPrefixesOk returns a tuple with the IpPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPrefixes

`func (o *PeeringConnectionResIpv4) SetIpPrefixes(v []string)`

SetIpPrefixes sets IpPrefixes field to given value.

### HasIpPrefixes

`func (o *PeeringConnectionResIpv4) HasIpPrefixes() bool`

HasIpPrefixes returns a boolean if a field has been set.

### GetEnabled

`func (o *PeeringConnectionResIpv4) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PeeringConnectionResIpv4) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PeeringConnectionResIpv4) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PeeringConnectionResIpv4) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


