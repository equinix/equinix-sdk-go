# RoutingProtocolIpv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerRoutes** | [**[]RoutingProtocolCustomerRouteIpv6**](RoutingProtocolCustomerRouteIpv6.md) |  | 
**Peerings** | [**[]RoutingProtocolPeeringIpv6**](RoutingProtocolPeeringIpv6.md) |  | 

## Methods

### NewRoutingProtocolIpv6

`func NewRoutingProtocolIpv6(customerRoutes []RoutingProtocolCustomerRouteIpv6, peerings []RoutingProtocolPeeringIpv6, ) *RoutingProtocolIpv6`

NewRoutingProtocolIpv6 instantiates a new RoutingProtocolIpv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolIpv6WithDefaults

`func NewRoutingProtocolIpv6WithDefaults() *RoutingProtocolIpv6`

NewRoutingProtocolIpv6WithDefaults instantiates a new RoutingProtocolIpv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerRoutes

`func (o *RoutingProtocolIpv6) GetCustomerRoutes() []RoutingProtocolCustomerRouteIpv6`

GetCustomerRoutes returns the CustomerRoutes field if non-nil, zero value otherwise.

### GetCustomerRoutesOk

`func (o *RoutingProtocolIpv6) GetCustomerRoutesOk() (*[]RoutingProtocolCustomerRouteIpv6, bool)`

GetCustomerRoutesOk returns a tuple with the CustomerRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoutes

`func (o *RoutingProtocolIpv6) SetCustomerRoutes(v []RoutingProtocolCustomerRouteIpv6)`

SetCustomerRoutes sets CustomerRoutes field to given value.


### GetPeerings

`func (o *RoutingProtocolIpv6) GetPeerings() []RoutingProtocolPeeringIpv6`

GetPeerings returns the Peerings field if non-nil, zero value otherwise.

### GetPeeringsOk

`func (o *RoutingProtocolIpv6) GetPeeringsOk() (*[]RoutingProtocolPeeringIpv6, bool)`

GetPeeringsOk returns a tuple with the Peerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerings

`func (o *RoutingProtocolIpv6) SetPeerings(v []RoutingProtocolPeeringIpv6)`

SetPeerings sets Peerings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


