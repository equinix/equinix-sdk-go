# CreateMetalGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpReservationId** | **string** | The UUID an a VRF IP Reservation that belongs to the same project as the one in which the Metal Gateway is to be created. Additionally, the VRF IP Reservation and the Virtual Network must reside in the same Metro. | 
**PrivateIpv4SubnetSize** | Pointer to **int32** | The subnet size (8, 16, 32, 64, or 128) of the private IPv4 reservation that will be created for the metal gateway. This field is required unless a project IP reservation was specified.           Please keep in mind that the number of private metal gateway ranges are limited per project. If you would like to increase the limit per project, please contact support for assistance. | [optional] 
**VirtualNetworkId** | **string** | THe UUID of a Metro Virtual Network that belongs to the same project as the one in which the Metal Gateway is to be created. Additionally, the Virtual Network and the VRF IP Reservation must reside in the same metro. In the case of the IP reservation being an IPv6 based VRF IP Reservation, the Virtual Network must not already have an associated IPv6 based VRF IP Reservation. There can be exactly one IPv6 based VRF IP Reservation associated to a Virtual Network. | 

## Methods

### NewCreateMetalGatewayRequest

`func NewCreateMetalGatewayRequest(ipReservationId string, virtualNetworkId string, ) *CreateMetalGatewayRequest`

NewCreateMetalGatewayRequest instantiates a new CreateMetalGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMetalGatewayRequestWithDefaults

`func NewCreateMetalGatewayRequestWithDefaults() *CreateMetalGatewayRequest`

NewCreateMetalGatewayRequestWithDefaults instantiates a new CreateMetalGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpReservationId

`func (o *CreateMetalGatewayRequest) GetIpReservationId() string`

GetIpReservationId returns the IpReservationId field if non-nil, zero value otherwise.

### GetIpReservationIdOk

`func (o *CreateMetalGatewayRequest) GetIpReservationIdOk() (*string, bool)`

GetIpReservationIdOk returns a tuple with the IpReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReservationId

`func (o *CreateMetalGatewayRequest) SetIpReservationId(v string)`

SetIpReservationId sets IpReservationId field to given value.


### GetPrivateIpv4SubnetSize

`func (o *CreateMetalGatewayRequest) GetPrivateIpv4SubnetSize() int32`

GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPrivateIpv4SubnetSizeOk

`func (o *CreateMetalGatewayRequest) GetPrivateIpv4SubnetSizeOk() (*int32, bool)`

GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv4SubnetSize

`func (o *CreateMetalGatewayRequest) SetPrivateIpv4SubnetSize(v int32)`

SetPrivateIpv4SubnetSize sets PrivateIpv4SubnetSize field to given value.

### HasPrivateIpv4SubnetSize

`func (o *CreateMetalGatewayRequest) HasPrivateIpv4SubnetSize() bool`

HasPrivateIpv4SubnetSize returns a boolean if a field has been set.

### GetVirtualNetworkId

`func (o *CreateMetalGatewayRequest) GetVirtualNetworkId() string`

GetVirtualNetworkId returns the VirtualNetworkId field if non-nil, zero value otherwise.

### GetVirtualNetworkIdOk

`func (o *CreateMetalGatewayRequest) GetVirtualNetworkIdOk() (*string, bool)`

GetVirtualNetworkIdOk returns a tuple with the VirtualNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkId

`func (o *CreateMetalGatewayRequest) SetVirtualNetworkId(v string)`

SetVirtualNetworkId sets VirtualNetworkId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


