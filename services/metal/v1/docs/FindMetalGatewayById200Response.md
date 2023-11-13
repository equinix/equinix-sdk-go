# FindMetalGatewayById200Response

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

### NewFindMetalGatewayById200Response

`func NewFindMetalGatewayById200Response() *FindMetalGatewayById200Response`

NewFindMetalGatewayById200Response instantiates a new FindMetalGatewayById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindMetalGatewayById200ResponseWithDefaults

`func NewFindMetalGatewayById200ResponseWithDefaults() *FindMetalGatewayById200Response`

NewFindMetalGatewayById200ResponseWithDefaults instantiates a new FindMetalGatewayById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FindMetalGatewayById200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindMetalGatewayById200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindMetalGatewayById200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FindMetalGatewayById200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FindMetalGatewayById200Response) GetCreatedBy() Href`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FindMetalGatewayById200Response) GetCreatedByOk() (*Href, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FindMetalGatewayById200Response) SetCreatedBy(v Href)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FindMetalGatewayById200Response) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetHref

`func (o *FindMetalGatewayById200Response) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FindMetalGatewayById200Response) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FindMetalGatewayById200Response) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FindMetalGatewayById200Response) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FindMetalGatewayById200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindMetalGatewayById200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindMetalGatewayById200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FindMetalGatewayById200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpReservation

`func (o *FindMetalGatewayById200Response) GetIpReservation() VrfIpReservation`

GetIpReservation returns the IpReservation field if non-nil, zero value otherwise.

### GetIpReservationOk

`func (o *FindMetalGatewayById200Response) GetIpReservationOk() (*VrfIpReservation, bool)`

GetIpReservationOk returns a tuple with the IpReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReservation

`func (o *FindMetalGatewayById200Response) SetIpReservation(v VrfIpReservation)`

SetIpReservation sets IpReservation field to given value.

### HasIpReservation

`func (o *FindMetalGatewayById200Response) HasIpReservation() bool`

HasIpReservation returns a boolean if a field has been set.

### GetProject

`func (o *FindMetalGatewayById200Response) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FindMetalGatewayById200Response) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FindMetalGatewayById200Response) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *FindMetalGatewayById200Response) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *FindMetalGatewayById200Response) GetState() MetalGatewayState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FindMetalGatewayById200Response) GetStateOk() (*MetalGatewayState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FindMetalGatewayById200Response) SetState(v MetalGatewayState)`

SetState sets State field to given value.

### HasState

`func (o *FindMetalGatewayById200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FindMetalGatewayById200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindMetalGatewayById200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindMetalGatewayById200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FindMetalGatewayById200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *FindMetalGatewayById200Response) GetVirtualNetwork() VirtualNetwork`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *FindMetalGatewayById200Response) GetVirtualNetworkOk() (*VirtualNetwork, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *FindMetalGatewayById200Response) SetVirtualNetwork(v VirtualNetwork)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *FindMetalGatewayById200Response) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetVrf

`func (o *FindMetalGatewayById200Response) GetVrf() Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *FindMetalGatewayById200Response) GetVrfOk() (*Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *FindMetalGatewayById200Response) SetVrf(v Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *FindMetalGatewayById200Response) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


