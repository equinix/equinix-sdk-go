# RoutingProtocolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerAsn** | Pointer to **int32** | Customer Autonomous System Number (ASN) | [optional] 
**MacAddress** | Pointer to **string** | MAC address of the customer interface | [optional] 
**BgpIpv4** | Pointer to [**ExchangeServiceRequestBgp**](ExchangeServiceRequestBgp.md) |  | [optional] 
**BgpIpv6** | Pointer to [**ExchangeServiceRequestBgp**](ExchangeServiceRequestBgp.md) |  | [optional] 

## Methods

### NewRoutingProtocolRequest

`func NewRoutingProtocolRequest() *RoutingProtocolRequest`

NewRoutingProtocolRequest instantiates a new RoutingProtocolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolRequestWithDefaults

`func NewRoutingProtocolRequestWithDefaults() *RoutingProtocolRequest`

NewRoutingProtocolRequestWithDefaults instantiates a new RoutingProtocolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerAsn

`func (o *RoutingProtocolRequest) GetCustomerAsn() int32`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *RoutingProtocolRequest) GetCustomerAsnOk() (*int32, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *RoutingProtocolRequest) SetCustomerAsn(v int32)`

SetCustomerAsn sets CustomerAsn field to given value.

### HasCustomerAsn

`func (o *RoutingProtocolRequest) HasCustomerAsn() bool`

HasCustomerAsn returns a boolean if a field has been set.

### GetMacAddress

`func (o *RoutingProtocolRequest) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *RoutingProtocolRequest) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *RoutingProtocolRequest) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *RoutingProtocolRequest) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetBgpIpv4

`func (o *RoutingProtocolRequest) GetBgpIpv4() ExchangeServiceRequestBgp`

GetBgpIpv4 returns the BgpIpv4 field if non-nil, zero value otherwise.

### GetBgpIpv4Ok

`func (o *RoutingProtocolRequest) GetBgpIpv4Ok() (*ExchangeServiceRequestBgp, bool)`

GetBgpIpv4Ok returns a tuple with the BgpIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv4

`func (o *RoutingProtocolRequest) SetBgpIpv4(v ExchangeServiceRequestBgp)`

SetBgpIpv4 sets BgpIpv4 field to given value.

### HasBgpIpv4

`func (o *RoutingProtocolRequest) HasBgpIpv4() bool`

HasBgpIpv4 returns a boolean if a field has been set.

### GetBgpIpv6

`func (o *RoutingProtocolRequest) GetBgpIpv6() ExchangeServiceRequestBgp`

GetBgpIpv6 returns the BgpIpv6 field if non-nil, zero value otherwise.

### GetBgpIpv6Ok

`func (o *RoutingProtocolRequest) GetBgpIpv6Ok() (*ExchangeServiceRequestBgp, bool)`

GetBgpIpv6Ok returns a tuple with the BgpIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv6

`func (o *RoutingProtocolRequest) SetBgpIpv6(v ExchangeServiceRequestBgp)`

SetBgpIpv6 sets BgpIpv6 field to given value.

### HasBgpIpv6

`func (o *RoutingProtocolRequest) HasBgpIpv6() bool`

HasBgpIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


