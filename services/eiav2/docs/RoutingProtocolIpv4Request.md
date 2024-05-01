# RoutingProtocolIpv4Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerRoutes** | [**[]CustomerRouteIpv4Request**](CustomerRouteIpv4Request.md) | Collection of customer routes of IP Version 4 (IPv4) | 

## Methods

### NewRoutingProtocolIpv4Request

`func NewRoutingProtocolIpv4Request(customerRoutes []CustomerRouteIpv4Request, ) *RoutingProtocolIpv4Request`

NewRoutingProtocolIpv4Request instantiates a new RoutingProtocolIpv4Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolIpv4RequestWithDefaults

`func NewRoutingProtocolIpv4RequestWithDefaults() *RoutingProtocolIpv4Request`

NewRoutingProtocolIpv4RequestWithDefaults instantiates a new RoutingProtocolIpv4Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerRoutes

`func (o *RoutingProtocolIpv4Request) GetCustomerRoutes() []CustomerRouteIpv4Request`

GetCustomerRoutes returns the CustomerRoutes field if non-nil, zero value otherwise.

### GetCustomerRoutesOk

`func (o *RoutingProtocolIpv4Request) GetCustomerRoutesOk() (*[]CustomerRouteIpv4Request, bool)`

GetCustomerRoutesOk returns a tuple with the CustomerRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoutes

`func (o *RoutingProtocolIpv4Request) SetCustomerRoutes(v []CustomerRouteIpv4Request)`

SetCustomerRoutes sets CustomerRoutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


