# DirectRoutingProtocolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4** | Pointer to [**DirectRoutingProtocolIpv4Request**](DirectRoutingProtocolIpv4Request.md) |  | [optional] 
**Ipv6** | Pointer to [**DirectRoutingProtocolIpv6Request**](DirectRoutingProtocolIpv6Request.md) |  | [optional] 

## Methods

### NewDirectRoutingProtocolRequest

`func NewDirectRoutingProtocolRequest() *DirectRoutingProtocolRequest`

NewDirectRoutingProtocolRequest instantiates a new DirectRoutingProtocolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectRoutingProtocolRequestWithDefaults

`func NewDirectRoutingProtocolRequestWithDefaults() *DirectRoutingProtocolRequest`

NewDirectRoutingProtocolRequestWithDefaults instantiates a new DirectRoutingProtocolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4

`func (o *DirectRoutingProtocolRequest) GetIpv4() DirectRoutingProtocolIpv4Request`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *DirectRoutingProtocolRequest) GetIpv4Ok() (*DirectRoutingProtocolIpv4Request, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *DirectRoutingProtocolRequest) SetIpv4(v DirectRoutingProtocolIpv4Request)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *DirectRoutingProtocolRequest) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *DirectRoutingProtocolRequest) GetIpv6() DirectRoutingProtocolIpv6Request`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *DirectRoutingProtocolRequest) GetIpv6Ok() (*DirectRoutingProtocolIpv6Request, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *DirectRoutingProtocolRequest) SetIpv6(v DirectRoutingProtocolIpv6Request)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *DirectRoutingProtocolRequest) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


