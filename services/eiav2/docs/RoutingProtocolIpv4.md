# RoutingProtocolIpv4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerRoutes** | [**[]RoutingProtocolCustomerRouteIpv4**](RoutingProtocolCustomerRouteIpv4.md) |  | 
**Peerings** | [**[]RoutingProtocolPeeringIpv4**](RoutingProtocolPeeringIpv4.md) |  | 

## Methods

### NewRoutingProtocolIpv4

`func NewRoutingProtocolIpv4(customerRoutes []RoutingProtocolCustomerRouteIpv4, peerings []RoutingProtocolPeeringIpv4, ) *RoutingProtocolIpv4`

NewRoutingProtocolIpv4 instantiates a new RoutingProtocolIpv4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolIpv4WithDefaults

`func NewRoutingProtocolIpv4WithDefaults() *RoutingProtocolIpv4`

NewRoutingProtocolIpv4WithDefaults instantiates a new RoutingProtocolIpv4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerRoutes

`func (o *RoutingProtocolIpv4) GetCustomerRoutes() []RoutingProtocolCustomerRouteIpv4`

GetCustomerRoutes returns the CustomerRoutes field if non-nil, zero value otherwise.

### GetCustomerRoutesOk

`func (o *RoutingProtocolIpv4) GetCustomerRoutesOk() (*[]RoutingProtocolCustomerRouteIpv4, bool)`

GetCustomerRoutesOk returns a tuple with the CustomerRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoutes

`func (o *RoutingProtocolIpv4) SetCustomerRoutes(v []RoutingProtocolCustomerRouteIpv4)`

SetCustomerRoutes sets CustomerRoutes field to given value.


### GetPeerings

`func (o *RoutingProtocolIpv4) GetPeerings() []RoutingProtocolPeeringIpv4`

GetPeerings returns the Peerings field if non-nil, zero value otherwise.

### GetPeeringsOk

`func (o *RoutingProtocolIpv4) GetPeeringsOk() (*[]RoutingProtocolPeeringIpv4, bool)`

GetPeeringsOk returns a tuple with the Peerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerings

`func (o *RoutingProtocolIpv4) SetPeerings(v []RoutingProtocolPeeringIpv4)`

SetPeerings sets Peerings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


