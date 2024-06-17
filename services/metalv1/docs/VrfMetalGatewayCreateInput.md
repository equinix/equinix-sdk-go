# VrfMetalGatewayCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpReservationId** | **string** | The UUID an a VRF IP Reservation that belongs to the same project as the one in which the Metal Gateway is to be created. Additionally, the VRF IP Reservation and the Virtual Network must reside in the same Metro. | 
**VirtualNetworkId** | **string** | THe UUID of a Metro Virtual Network that belongs to the same project as the one in which the Metal Gateway is to be created. Additionally, the Virtual Network and the VRF IP Reservation must reside in the same metro. In the case of the IP reservation being an IPv6 based VRF IP Reservation, the Virtual Network must not already have an associated IPv6 based VRF IP Reservation. There can be exactly one IPv6 based VRF IP Reservation associated to a Virtual Network. | 

## Methods

### NewVrfMetalGatewayCreateInput

`func NewVrfMetalGatewayCreateInput(ipReservationId string, virtualNetworkId string, ) *VrfMetalGatewayCreateInput`

NewVrfMetalGatewayCreateInput instantiates a new VrfMetalGatewayCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfMetalGatewayCreateInputWithDefaults

`func NewVrfMetalGatewayCreateInputWithDefaults() *VrfMetalGatewayCreateInput`

NewVrfMetalGatewayCreateInputWithDefaults instantiates a new VrfMetalGatewayCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpReservationId

`func (o *VrfMetalGatewayCreateInput) GetIpReservationId() string`

GetIpReservationId returns the IpReservationId field if non-nil, zero value otherwise.

### GetIpReservationIdOk

`func (o *VrfMetalGatewayCreateInput) GetIpReservationIdOk() (*string, bool)`

GetIpReservationIdOk returns a tuple with the IpReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReservationId

`func (o *VrfMetalGatewayCreateInput) SetIpReservationId(v string)`

SetIpReservationId sets IpReservationId field to given value.


### GetVirtualNetworkId

`func (o *VrfMetalGatewayCreateInput) GetVirtualNetworkId() string`

GetVirtualNetworkId returns the VirtualNetworkId field if non-nil, zero value otherwise.

### GetVirtualNetworkIdOk

`func (o *VrfMetalGatewayCreateInput) GetVirtualNetworkIdOk() (*string, bool)`

GetVirtualNetworkIdOk returns a tuple with the VirtualNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkId

`func (o *VrfMetalGatewayCreateInput) SetVirtualNetworkId(v string)`

SetVirtualNetworkId sets VirtualNetworkId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


