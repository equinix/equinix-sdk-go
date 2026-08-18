# RoutingProtocolResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerAsn** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**Vlan** | Pointer to **int64** | VLAN ID | [optional] 
**RouteServerAsn** | Pointer to **int64** | Equinix Route Server ASN | [optional] 
**BgpIpv4** | Pointer to [**ExchangeServiceResponseBgp**](ExchangeServiceResponseBgp.md) |  | [optional] 
**BgpIpv6** | Pointer to [**ExchangeServiceResponseBgp**](ExchangeServiceResponseBgp.md) |  | [optional] 
**RouteCollector** | Pointer to [**RouteCollector**](RouteCollector.md) |  | [optional] 

## Methods

### NewRoutingProtocolResponse

`func NewRoutingProtocolResponse() *RoutingProtocolResponse`

NewRoutingProtocolResponse instantiates a new RoutingProtocolResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolResponseWithDefaults

`func NewRoutingProtocolResponseWithDefaults() *RoutingProtocolResponse`

NewRoutingProtocolResponseWithDefaults instantiates a new RoutingProtocolResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerAsn

`func (o *RoutingProtocolResponse) GetCustomerAsn() string`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *RoutingProtocolResponse) GetCustomerAsnOk() (*string, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *RoutingProtocolResponse) SetCustomerAsn(v string)`

SetCustomerAsn sets CustomerAsn field to given value.

### HasCustomerAsn

`func (o *RoutingProtocolResponse) HasCustomerAsn() bool`

HasCustomerAsn returns a boolean if a field has been set.

### GetMacAddress

`func (o *RoutingProtocolResponse) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *RoutingProtocolResponse) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *RoutingProtocolResponse) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *RoutingProtocolResponse) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetVlan

`func (o *RoutingProtocolResponse) GetVlan() int64`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *RoutingProtocolResponse) GetVlanOk() (*int64, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *RoutingProtocolResponse) SetVlan(v int64)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *RoutingProtocolResponse) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetRouteServerAsn

`func (o *RoutingProtocolResponse) GetRouteServerAsn() int64`

GetRouteServerAsn returns the RouteServerAsn field if non-nil, zero value otherwise.

### GetRouteServerAsnOk

`func (o *RoutingProtocolResponse) GetRouteServerAsnOk() (*int64, bool)`

GetRouteServerAsnOk returns a tuple with the RouteServerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteServerAsn

`func (o *RoutingProtocolResponse) SetRouteServerAsn(v int64)`

SetRouteServerAsn sets RouteServerAsn field to given value.

### HasRouteServerAsn

`func (o *RoutingProtocolResponse) HasRouteServerAsn() bool`

HasRouteServerAsn returns a boolean if a field has been set.

### GetBgpIpv4

`func (o *RoutingProtocolResponse) GetBgpIpv4() ExchangeServiceResponseBgp`

GetBgpIpv4 returns the BgpIpv4 field if non-nil, zero value otherwise.

### GetBgpIpv4Ok

`func (o *RoutingProtocolResponse) GetBgpIpv4Ok() (*ExchangeServiceResponseBgp, bool)`

GetBgpIpv4Ok returns a tuple with the BgpIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv4

`func (o *RoutingProtocolResponse) SetBgpIpv4(v ExchangeServiceResponseBgp)`

SetBgpIpv4 sets BgpIpv4 field to given value.

### HasBgpIpv4

`func (o *RoutingProtocolResponse) HasBgpIpv4() bool`

HasBgpIpv4 returns a boolean if a field has been set.

### GetBgpIpv6

`func (o *RoutingProtocolResponse) GetBgpIpv6() ExchangeServiceResponseBgp`

GetBgpIpv6 returns the BgpIpv6 field if non-nil, zero value otherwise.

### GetBgpIpv6Ok

`func (o *RoutingProtocolResponse) GetBgpIpv6Ok() (*ExchangeServiceResponseBgp, bool)`

GetBgpIpv6Ok returns a tuple with the BgpIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv6

`func (o *RoutingProtocolResponse) SetBgpIpv6(v ExchangeServiceResponseBgp)`

SetBgpIpv6 sets BgpIpv6 field to given value.

### HasBgpIpv6

`func (o *RoutingProtocolResponse) HasBgpIpv6() bool`

HasBgpIpv6 returns a boolean if a field has been set.

### GetRouteCollector

`func (o *RoutingProtocolResponse) GetRouteCollector() RouteCollector`

GetRouteCollector returns the RouteCollector field if non-nil, zero value otherwise.

### GetRouteCollectorOk

`func (o *RoutingProtocolResponse) GetRouteCollectorOk() (*RouteCollector, bool)`

GetRouteCollectorOk returns a tuple with the RouteCollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteCollector

`func (o *RoutingProtocolResponse) SetRouteCollector(v RouteCollector)`

SetRouteCollector sets RouteCollector field to given value.

### HasRouteCollector

`func (o *RoutingProtocolResponse) HasRouteCollector() bool`

HasRouteCollector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


