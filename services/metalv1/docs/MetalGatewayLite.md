# MetalGatewayLite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**GatewayAddress** | Pointer to **string** | The gateway address with subnet CIDR value for this Metal Gateway. For example, a Metal Gateway using an IP reservation with block 10.1.2.0/27 would have a gateway address of 10.1.2.1/27. | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**MetalGatewayState**](MetalGatewayState.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Vlan** | Pointer to **int32** | The VLAN id of the Virtual Network record associated to this Metal Gateway. | [optional] 

## Methods

### NewMetalGatewayLite

`func NewMetalGatewayLite() *MetalGatewayLite`

NewMetalGatewayLite instantiates a new MetalGatewayLite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetalGatewayLiteWithDefaults

`func NewMetalGatewayLiteWithDefaults() *MetalGatewayLite`

NewMetalGatewayLiteWithDefaults instantiates a new MetalGatewayLite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MetalGatewayLite) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetalGatewayLite) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetalGatewayLite) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MetalGatewayLite) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGatewayAddress

`func (o *MetalGatewayLite) GetGatewayAddress() string`

GetGatewayAddress returns the GatewayAddress field if non-nil, zero value otherwise.

### GetGatewayAddressOk

`func (o *MetalGatewayLite) GetGatewayAddressOk() (*string, bool)`

GetGatewayAddressOk returns a tuple with the GatewayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddress

`func (o *MetalGatewayLite) SetGatewayAddress(v string)`

SetGatewayAddress sets GatewayAddress field to given value.

### HasGatewayAddress

`func (o *MetalGatewayLite) HasGatewayAddress() bool`

HasGatewayAddress returns a boolean if a field has been set.

### GetHref

`func (o *MetalGatewayLite) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MetalGatewayLite) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MetalGatewayLite) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *MetalGatewayLite) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *MetalGatewayLite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetalGatewayLite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetalGatewayLite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetalGatewayLite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *MetalGatewayLite) GetState() MetalGatewayState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MetalGatewayLite) GetStateOk() (*MetalGatewayState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MetalGatewayLite) SetState(v MetalGatewayState)`

SetState sets State field to given value.

### HasState

`func (o *MetalGatewayLite) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MetalGatewayLite) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MetalGatewayLite) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MetalGatewayLite) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MetalGatewayLite) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVlan

`func (o *MetalGatewayLite) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *MetalGatewayLite) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *MetalGatewayLite) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *MetalGatewayLite) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


