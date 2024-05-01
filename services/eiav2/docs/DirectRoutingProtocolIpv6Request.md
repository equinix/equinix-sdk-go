# DirectRoutingProtocolIpv6Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerRoutes** | [**[]CustomerRouteIpv6Request**](CustomerRouteIpv6Request.md) | Collection of customer routes of IP Version 6 (IPv6) | 
**Peerings** | Pointer to [**[]DirectPeeringIpv6Request**](DirectPeeringIpv6Request.md) |  | [optional] 

## Methods

### NewDirectRoutingProtocolIpv6Request

`func NewDirectRoutingProtocolIpv6Request(customerRoutes []CustomerRouteIpv6Request, ) *DirectRoutingProtocolIpv6Request`

NewDirectRoutingProtocolIpv6Request instantiates a new DirectRoutingProtocolIpv6Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectRoutingProtocolIpv6RequestWithDefaults

`func NewDirectRoutingProtocolIpv6RequestWithDefaults() *DirectRoutingProtocolIpv6Request`

NewDirectRoutingProtocolIpv6RequestWithDefaults instantiates a new DirectRoutingProtocolIpv6Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerRoutes

`func (o *DirectRoutingProtocolIpv6Request) GetCustomerRoutes() []CustomerRouteIpv6Request`

GetCustomerRoutes returns the CustomerRoutes field if non-nil, zero value otherwise.

### GetCustomerRoutesOk

`func (o *DirectRoutingProtocolIpv6Request) GetCustomerRoutesOk() (*[]CustomerRouteIpv6Request, bool)`

GetCustomerRoutesOk returns a tuple with the CustomerRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoutes

`func (o *DirectRoutingProtocolIpv6Request) SetCustomerRoutes(v []CustomerRouteIpv6Request)`

SetCustomerRoutes sets CustomerRoutes field to given value.


### GetPeerings

`func (o *DirectRoutingProtocolIpv6Request) GetPeerings() []DirectPeeringIpv6Request`

GetPeerings returns the Peerings field if non-nil, zero value otherwise.

### GetPeeringsOk

`func (o *DirectRoutingProtocolIpv6Request) GetPeeringsOk() (*[]DirectPeeringIpv6Request, bool)`

GetPeeringsOk returns a tuple with the Peerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerings

`func (o *DirectRoutingProtocolIpv6Request) SetPeerings(v []DirectPeeringIpv6Request)`

SetPeerings sets Peerings field to given value.

### HasPeerings

`func (o *DirectRoutingProtocolIpv6Request) HasPeerings() bool`

HasPeerings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


