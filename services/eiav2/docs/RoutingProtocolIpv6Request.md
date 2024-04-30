# RoutingProtocolIpv6Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerRoutes** | [**[]CustomerRouteIpv6Request**](CustomerRouteIpv6Request.md) | Collection of customer routes of IP Version 6 (IPv6) | 

## Methods

### NewRoutingProtocolIpv6Request

`func NewRoutingProtocolIpv6Request(customerRoutes []CustomerRouteIpv6Request, ) *RoutingProtocolIpv6Request`

NewRoutingProtocolIpv6Request instantiates a new RoutingProtocolIpv6Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolIpv6RequestWithDefaults

`func NewRoutingProtocolIpv6RequestWithDefaults() *RoutingProtocolIpv6Request`

NewRoutingProtocolIpv6RequestWithDefaults instantiates a new RoutingProtocolIpv6Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerRoutes

`func (o *RoutingProtocolIpv6Request) GetCustomerRoutes() []CustomerRouteIpv6Request`

GetCustomerRoutes returns the CustomerRoutes field if non-nil, zero value otherwise.

### GetCustomerRoutesOk

`func (o *RoutingProtocolIpv6Request) GetCustomerRoutesOk() (*[]CustomerRouteIpv6Request, bool)`

GetCustomerRoutesOk returns a tuple with the CustomerRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoutes

`func (o *RoutingProtocolIpv6Request) SetCustomerRoutes(v []CustomerRouteIpv6Request)`

SetCustomerRoutes sets CustomerRoutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


