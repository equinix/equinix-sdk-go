# BGPConnectionIpv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerPeerIp** | **string** | Customer side peering ip | 
**EquinixPeerIp** | Pointer to **string** | Equinix side peering ip | [optional] 
**Enabled** | **bool** | Admin status for the BGP session | 
**OutboundASPrependCount** | Pointer to **int64** | AS path prepend count | [optional] 
**Operation** | Pointer to [**BGPConnectionOperation**](BGPConnectionOperation.md) |  | [optional] 

## Methods

### NewBGPConnectionIpv6

`func NewBGPConnectionIpv6(customerPeerIp string, enabled bool, ) *BGPConnectionIpv6`

NewBGPConnectionIpv6 instantiates a new BGPConnectionIpv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGPConnectionIpv6WithDefaults

`func NewBGPConnectionIpv6WithDefaults() *BGPConnectionIpv6`

NewBGPConnectionIpv6WithDefaults instantiates a new BGPConnectionIpv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerPeerIp

`func (o *BGPConnectionIpv6) GetCustomerPeerIp() string`

GetCustomerPeerIp returns the CustomerPeerIp field if non-nil, zero value otherwise.

### GetCustomerPeerIpOk

`func (o *BGPConnectionIpv6) GetCustomerPeerIpOk() (*string, bool)`

GetCustomerPeerIpOk returns a tuple with the CustomerPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPeerIp

`func (o *BGPConnectionIpv6) SetCustomerPeerIp(v string)`

SetCustomerPeerIp sets CustomerPeerIp field to given value.


### GetEquinixPeerIp

`func (o *BGPConnectionIpv6) GetEquinixPeerIp() string`

GetEquinixPeerIp returns the EquinixPeerIp field if non-nil, zero value otherwise.

### GetEquinixPeerIpOk

`func (o *BGPConnectionIpv6) GetEquinixPeerIpOk() (*string, bool)`

GetEquinixPeerIpOk returns a tuple with the EquinixPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixPeerIp

`func (o *BGPConnectionIpv6) SetEquinixPeerIp(v string)`

SetEquinixPeerIp sets EquinixPeerIp field to given value.

### HasEquinixPeerIp

`func (o *BGPConnectionIpv6) HasEquinixPeerIp() bool`

HasEquinixPeerIp returns a boolean if a field has been set.

### GetEnabled

`func (o *BGPConnectionIpv6) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BGPConnectionIpv6) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BGPConnectionIpv6) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOutboundASPrependCount

`func (o *BGPConnectionIpv6) GetOutboundASPrependCount() int64`

GetOutboundASPrependCount returns the OutboundASPrependCount field if non-nil, zero value otherwise.

### GetOutboundASPrependCountOk

`func (o *BGPConnectionIpv6) GetOutboundASPrependCountOk() (*int64, bool)`

GetOutboundASPrependCountOk returns a tuple with the OutboundASPrependCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundASPrependCount

`func (o *BGPConnectionIpv6) SetOutboundASPrependCount(v int64)`

SetOutboundASPrependCount sets OutboundASPrependCount field to given value.

### HasOutboundASPrependCount

`func (o *BGPConnectionIpv6) HasOutboundASPrependCount() bool`

HasOutboundASPrependCount returns a boolean if a field has been set.

### GetOperation

`func (o *BGPConnectionIpv6) GetOperation() BGPConnectionOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BGPConnectionIpv6) GetOperationOk() (*BGPConnectionOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BGPConnectionIpv6) SetOperation(v BGPConnectionOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *BGPConnectionIpv6) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


