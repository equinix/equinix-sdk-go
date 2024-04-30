# RoutingProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RoutingProtocolType**](RoutingProtocolType.md) |  | 
**Ipv4** | Pointer to [**RoutingProtocolIpBlock**](RoutingProtocolIpBlock.md) |  | [optional] 
**Ipv6** | Pointer to [**RoutingProtocolIpBlock**](RoutingProtocolIpBlock.md) |  | [optional] 

## Methods

### NewRoutingProtocol

`func NewRoutingProtocol(type_ RoutingProtocolType, ) *RoutingProtocol`

NewRoutingProtocol instantiates a new RoutingProtocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolWithDefaults

`func NewRoutingProtocolWithDefaults() *RoutingProtocol`

NewRoutingProtocolWithDefaults instantiates a new RoutingProtocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutingProtocol) GetType() RoutingProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocol) GetTypeOk() (*RoutingProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocol) SetType(v RoutingProtocolType)`

SetType sets Type field to given value.


### GetIpv4

`func (o *RoutingProtocol) GetIpv4() RoutingProtocolIpBlock`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *RoutingProtocol) GetIpv4Ok() (*RoutingProtocolIpBlock, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *RoutingProtocol) SetIpv4(v RoutingProtocolIpBlock)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *RoutingProtocol) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *RoutingProtocol) GetIpv6() RoutingProtocolIpBlock`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *RoutingProtocol) GetIpv6Ok() (*RoutingProtocolIpBlock, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *RoutingProtocol) SetIpv6(v RoutingProtocolIpBlock)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *RoutingProtocol) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


