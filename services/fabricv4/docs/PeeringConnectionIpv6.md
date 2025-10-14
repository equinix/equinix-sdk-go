# PeeringConnectionIpv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReverseDnsAddress** | Pointer to **string** |  | [optional] 
**AsSet** | Pointer to **string** |  | [optional] 
**MlpeEnabled** | Pointer to **bool** |  | [optional] 
**AuthKeys** | Pointer to [**[]PeeringConnectionIpv4AuthKeys**](PeeringConnectionIpv4AuthKeys.md) |  | [optional] 
**IpPrefixes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPeeringConnectionIpv6

`func NewPeeringConnectionIpv6() *PeeringConnectionIpv6`

NewPeeringConnectionIpv6 instantiates a new PeeringConnectionIpv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeeringConnectionIpv6WithDefaults

`func NewPeeringConnectionIpv6WithDefaults() *PeeringConnectionIpv6`

NewPeeringConnectionIpv6WithDefaults instantiates a new PeeringConnectionIpv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReverseDnsAddress

`func (o *PeeringConnectionIpv6) GetReverseDnsAddress() string`

GetReverseDnsAddress returns the ReverseDnsAddress field if non-nil, zero value otherwise.

### GetReverseDnsAddressOk

`func (o *PeeringConnectionIpv6) GetReverseDnsAddressOk() (*string, bool)`

GetReverseDnsAddressOk returns a tuple with the ReverseDnsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseDnsAddress

`func (o *PeeringConnectionIpv6) SetReverseDnsAddress(v string)`

SetReverseDnsAddress sets ReverseDnsAddress field to given value.

### HasReverseDnsAddress

`func (o *PeeringConnectionIpv6) HasReverseDnsAddress() bool`

HasReverseDnsAddress returns a boolean if a field has been set.

### GetAsSet

`func (o *PeeringConnectionIpv6) GetAsSet() string`

GetAsSet returns the AsSet field if non-nil, zero value otherwise.

### GetAsSetOk

`func (o *PeeringConnectionIpv6) GetAsSetOk() (*string, bool)`

GetAsSetOk returns a tuple with the AsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsSet

`func (o *PeeringConnectionIpv6) SetAsSet(v string)`

SetAsSet sets AsSet field to given value.

### HasAsSet

`func (o *PeeringConnectionIpv6) HasAsSet() bool`

HasAsSet returns a boolean if a field has been set.

### GetMlpeEnabled

`func (o *PeeringConnectionIpv6) GetMlpeEnabled() bool`

GetMlpeEnabled returns the MlpeEnabled field if non-nil, zero value otherwise.

### GetMlpeEnabledOk

`func (o *PeeringConnectionIpv6) GetMlpeEnabledOk() (*bool, bool)`

GetMlpeEnabledOk returns a tuple with the MlpeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlpeEnabled

`func (o *PeeringConnectionIpv6) SetMlpeEnabled(v bool)`

SetMlpeEnabled sets MlpeEnabled field to given value.

### HasMlpeEnabled

`func (o *PeeringConnectionIpv6) HasMlpeEnabled() bool`

HasMlpeEnabled returns a boolean if a field has been set.

### GetAuthKeys

`func (o *PeeringConnectionIpv6) GetAuthKeys() []PeeringConnectionIpv4AuthKeys`

GetAuthKeys returns the AuthKeys field if non-nil, zero value otherwise.

### GetAuthKeysOk

`func (o *PeeringConnectionIpv6) GetAuthKeysOk() (*[]PeeringConnectionIpv4AuthKeys, bool)`

GetAuthKeysOk returns a tuple with the AuthKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKeys

`func (o *PeeringConnectionIpv6) SetAuthKeys(v []PeeringConnectionIpv4AuthKeys)`

SetAuthKeys sets AuthKeys field to given value.

### HasAuthKeys

`func (o *PeeringConnectionIpv6) HasAuthKeys() bool`

HasAuthKeys returns a boolean if a field has been set.

### GetIpPrefixes

`func (o *PeeringConnectionIpv6) GetIpPrefixes() []string`

GetIpPrefixes returns the IpPrefixes field if non-nil, zero value otherwise.

### GetIpPrefixesOk

`func (o *PeeringConnectionIpv6) GetIpPrefixesOk() (*[]string, bool)`

GetIpPrefixesOk returns a tuple with the IpPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPrefixes

`func (o *PeeringConnectionIpv6) SetIpPrefixes(v []string)`

SetIpPrefixes sets IpPrefixes field to given value.

### HasIpPrefixes

`func (o *PeeringConnectionIpv6) HasIpPrefixes() bool`

HasIpPrefixes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


