# DirectRoutingProtocolIpv4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerRoutes** | [**[]CustomerRouteIpv4Request**](CustomerRouteIpv4Request.md) | Collection of customer routes of IP Version 4 (IPv4) | 
**Peerings** | Pointer to [**[]DirectPeeringIpv4Request**](DirectPeeringIpv4Request.md) |  | [optional] 

## Methods

### NewDirectRoutingProtocolIpv4Request

`func NewDirectRoutingProtocolIpv4Request(customerRoutes []CustomerRouteIpv4Request, ) *DirectRoutingProtocolIpv4Request`

NewDirectRoutingProtocolIpv4Request instantiates a new DirectRoutingProtocolIpv4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectRoutingProtocolIpv4RequestWithDefaults

`func NewDirectRoutingProtocolIpv4RequestWithDefaults() *DirectRoutingProtocolIpv4Request`

NewDirectRoutingProtocolIpv4RequestWithDefaults instantiates a new DirectRoutingProtocolIpv4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerRoutes

`func (o *DirectRoutingProtocolIpv4Request) GetCustomerRoutes() []CustomerRouteIpv4Request`

GetCustomerRoutes returns the CustomerRoutes field if non-nil, zero value otherwise.

### GetCustomerRoutesOk

`func (o *DirectRoutingProtocolIpv4Request) GetCustomerRoutesOk() (*[]CustomerRouteIpv4Request, bool)`

GetCustomerRoutesOk returns a tuple with the CustomerRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoutes

`func (o *DirectRoutingProtocolIpv4Request) SetCustomerRoutes(v []CustomerRouteIpv4Request)`

SetCustomerRoutes sets CustomerRoutes field to given value.


### GetPeerings

`func (o *DirectRoutingProtocolIpv4Request) GetPeerings() []DirectPeeringIpv4Request`

GetPeerings returns the Peerings field if non-nil, zero value otherwise.

### GetPeeringsOk

`func (o *DirectRoutingProtocolIpv4Request) GetPeeringsOk() (*[]DirectPeeringIpv4Request, bool)`

GetPeeringsOk returns a tuple with the Peerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerings

`func (o *DirectRoutingProtocolIpv4Request) SetPeerings(v []DirectPeeringIpv4Request)`

SetPeerings sets Peerings field to given value.

### HasPeerings

`func (o *DirectRoutingProtocolIpv4Request) HasPeerings() bool`

HasPeerings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


