# StaticRoutingProtocolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4** | Pointer to [**RoutingProtocolIpv4Request**](RoutingProtocolIpv4Request.md) |  | [optional] 
**Ipv6** | Pointer to [**RoutingProtocolIpv6Request**](RoutingProtocolIpv6Request.md) |  | [optional] 

## Methods

### NewStaticRoutingProtocolRequest

`func NewStaticRoutingProtocolRequest() *StaticRoutingProtocolRequest`

NewStaticRoutingProtocolRequest instantiates a new StaticRoutingProtocolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticRoutingProtocolRequestWithDefaults

`func NewStaticRoutingProtocolRequestWithDefaults() *StaticRoutingProtocolRequest`

NewStaticRoutingProtocolRequestWithDefaults instantiates a new StaticRoutingProtocolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4

`func (o *StaticRoutingProtocolRequest) GetIpv4() RoutingProtocolIpv4Request`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *StaticRoutingProtocolRequest) GetIpv4Ok() (*RoutingProtocolIpv4Request, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *StaticRoutingProtocolRequest) SetIpv4(v RoutingProtocolIpv4Request)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *StaticRoutingProtocolRequest) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *StaticRoutingProtocolRequest) GetIpv6() RoutingProtocolIpv6Request`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *StaticRoutingProtocolRequest) GetIpv6Ok() (*RoutingProtocolIpv6Request, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *StaticRoutingProtocolRequest) SetIpv6(v RoutingProtocolIpv6Request)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *StaticRoutingProtocolRequest) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


