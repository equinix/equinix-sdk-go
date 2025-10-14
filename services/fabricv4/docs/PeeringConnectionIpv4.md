# PeeringConnectionIpv4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReverseDnsAddress** | Pointer to **string** |  | [optional] 
**AsSet** | Pointer to **string** |  | [optional] 
**MlpeEnabled** | Pointer to **bool** |  | [optional] 
**AuthKeys** | Pointer to [**[]PeeringConnectionIpv4AuthKeys**](PeeringConnectionIpv4AuthKeys.md) |  | [optional] 
**IpPrefixes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPeeringConnectionIpv4

`func NewPeeringConnectionIpv4() *PeeringConnectionIpv4`

NewPeeringConnectionIpv4 instantiates a new PeeringConnectionIpv4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeeringConnectionIpv4WithDefaults

`func NewPeeringConnectionIpv4WithDefaults() *PeeringConnectionIpv4`

NewPeeringConnectionIpv4WithDefaults instantiates a new PeeringConnectionIpv4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReverseDnsAddress

`func (o *PeeringConnectionIpv4) GetReverseDnsAddress() string`

GetReverseDnsAddress returns the ReverseDnsAddress field if non-nil, zero value otherwise.

### GetReverseDnsAddressOk

`func (o *PeeringConnectionIpv4) GetReverseDnsAddressOk() (*string, bool)`

GetReverseDnsAddressOk returns a tuple with the ReverseDnsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseDnsAddress

`func (o *PeeringConnectionIpv4) SetReverseDnsAddress(v string)`

SetReverseDnsAddress sets ReverseDnsAddress field to given value.

### HasReverseDnsAddress

`func (o *PeeringConnectionIpv4) HasReverseDnsAddress() bool`

HasReverseDnsAddress returns a boolean if a field has been set.

### GetAsSet

`func (o *PeeringConnectionIpv4) GetAsSet() string`

GetAsSet returns the AsSet field if non-nil, zero value otherwise.

### GetAsSetOk

`func (o *PeeringConnectionIpv4) GetAsSetOk() (*string, bool)`

GetAsSetOk returns a tuple with the AsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsSet

`func (o *PeeringConnectionIpv4) SetAsSet(v string)`

SetAsSet sets AsSet field to given value.

### HasAsSet

`func (o *PeeringConnectionIpv4) HasAsSet() bool`

HasAsSet returns a boolean if a field has been set.

### GetMlpeEnabled

`func (o *PeeringConnectionIpv4) GetMlpeEnabled() bool`

GetMlpeEnabled returns the MlpeEnabled field if non-nil, zero value otherwise.

### GetMlpeEnabledOk

`func (o *PeeringConnectionIpv4) GetMlpeEnabledOk() (*bool, bool)`

GetMlpeEnabledOk returns a tuple with the MlpeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlpeEnabled

`func (o *PeeringConnectionIpv4) SetMlpeEnabled(v bool)`

SetMlpeEnabled sets MlpeEnabled field to given value.

### HasMlpeEnabled

`func (o *PeeringConnectionIpv4) HasMlpeEnabled() bool`

HasMlpeEnabled returns a boolean if a field has been set.

### GetAuthKeys

`func (o *PeeringConnectionIpv4) GetAuthKeys() []PeeringConnectionIpv4AuthKeys`

GetAuthKeys returns the AuthKeys field if non-nil, zero value otherwise.

### GetAuthKeysOk

`func (o *PeeringConnectionIpv4) GetAuthKeysOk() (*[]PeeringConnectionIpv4AuthKeys, bool)`

GetAuthKeysOk returns a tuple with the AuthKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKeys

`func (o *PeeringConnectionIpv4) SetAuthKeys(v []PeeringConnectionIpv4AuthKeys)`

SetAuthKeys sets AuthKeys field to given value.

### HasAuthKeys

`func (o *PeeringConnectionIpv4) HasAuthKeys() bool`

HasAuthKeys returns a boolean if a field has been set.

### GetIpPrefixes

`func (o *PeeringConnectionIpv4) GetIpPrefixes() []string`

GetIpPrefixes returns the IpPrefixes field if non-nil, zero value otherwise.

### GetIpPrefixesOk

`func (o *PeeringConnectionIpv4) GetIpPrefixesOk() (*[]string, bool)`

GetIpPrefixesOk returns a tuple with the IpPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPrefixes

`func (o *PeeringConnectionIpv4) SetIpPrefixes(v []string)`

SetIpPrefixes sets IpPrefixes field to given value.

### HasIpPrefixes

`func (o *PeeringConnectionIpv4) HasIpPrefixes() bool`

HasIpPrefixes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


