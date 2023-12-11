# VrfMetalGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**Href**](Href.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpReservation** | Pointer to [**VrfIpReservation**](VrfIpReservation.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**State** | Pointer to [**MetalGatewayState**](MetalGatewayState.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VirtualNetwork** | Pointer to [**VirtualNetwork**](VirtualNetwork.md) |  | [optional] 
**Vrf** | Pointer to [**Vrf**](Vrf.md) |  | [optional] 

## Methods

### NewVrfMetalGateway

`func NewVrfMetalGateway() *VrfMetalGateway`

NewVrfMetalGateway instantiates a new VrfMetalGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfMetalGatewayWithDefaults

`func NewVrfMetalGatewayWithDefaults() *VrfMetalGateway`

NewVrfMetalGatewayWithDefaults instantiates a new VrfMetalGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *VrfMetalGateway) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VrfMetalGateway) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VrfMetalGateway) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VrfMetalGateway) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VrfMetalGateway) GetCreatedBy() Href`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VrfMetalGateway) GetCreatedByOk() (*Href, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VrfMetalGateway) SetCreatedBy(v Href)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VrfMetalGateway) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetHref

`func (o *VrfMetalGateway) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VrfMetalGateway) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VrfMetalGateway) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *VrfMetalGateway) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *VrfMetalGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VrfMetalGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VrfMetalGateway) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VrfMetalGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpReservation

`func (o *VrfMetalGateway) GetIpReservation() VrfIpReservation`

GetIpReservation returns the IpReservation field if non-nil, zero value otherwise.

### GetIpReservationOk

`func (o *VrfMetalGateway) GetIpReservationOk() (*VrfIpReservation, bool)`

GetIpReservationOk returns a tuple with the IpReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReservation

`func (o *VrfMetalGateway) SetIpReservation(v VrfIpReservation)`

SetIpReservation sets IpReservation field to given value.

### HasIpReservation

`func (o *VrfMetalGateway) HasIpReservation() bool`

HasIpReservation returns a boolean if a field has been set.

### GetProject

`func (o *VrfMetalGateway) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VrfMetalGateway) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VrfMetalGateway) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *VrfMetalGateway) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *VrfMetalGateway) GetState() MetalGatewayState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VrfMetalGateway) GetStateOk() (*MetalGatewayState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VrfMetalGateway) SetState(v MetalGatewayState)`

SetState sets State field to given value.

### HasState

`func (o *VrfMetalGateway) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VrfMetalGateway) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VrfMetalGateway) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VrfMetalGateway) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VrfMetalGateway) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *VrfMetalGateway) GetVirtualNetwork() VirtualNetwork`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *VrfMetalGateway) GetVirtualNetworkOk() (*VirtualNetwork, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *VrfMetalGateway) SetVirtualNetwork(v VirtualNetwork)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *VrfMetalGateway) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetVrf

`func (o *VrfMetalGateway) GetVrf() Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *VrfMetalGateway) GetVrfOk() (*Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *VrfMetalGateway) SetVrf(v Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *VrfMetalGateway) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


