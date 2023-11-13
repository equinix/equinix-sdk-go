# PortVlanAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Native** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to [**Href**](Href.md) |  | [optional] 
**State** | Pointer to [**PortVlanAssignmentState**](PortVlanAssignmentState.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VirtualNetwork** | Pointer to [**Href**](Href.md) |  | [optional] 
**Vlan** | Pointer to **int32** |  | [optional] 

## Methods

### NewPortVlanAssignment

`func NewPortVlanAssignment() *PortVlanAssignment`

NewPortVlanAssignment instantiates a new PortVlanAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortVlanAssignmentWithDefaults

`func NewPortVlanAssignmentWithDefaults() *PortVlanAssignment`

NewPortVlanAssignmentWithDefaults instantiates a new PortVlanAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PortVlanAssignment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortVlanAssignment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortVlanAssignment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortVlanAssignment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *PortVlanAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortVlanAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortVlanAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortVlanAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNative

`func (o *PortVlanAssignment) GetNative() bool`

GetNative returns the Native field if non-nil, zero value otherwise.

### GetNativeOk

`func (o *PortVlanAssignment) GetNativeOk() (*bool, bool)`

GetNativeOk returns a tuple with the Native field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNative

`func (o *PortVlanAssignment) SetNative(v bool)`

SetNative sets Native field to given value.

### HasNative

`func (o *PortVlanAssignment) HasNative() bool`

HasNative returns a boolean if a field has been set.

### GetPort

`func (o *PortVlanAssignment) GetPort() Href`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PortVlanAssignment) GetPortOk() (*Href, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PortVlanAssignment) SetPort(v Href)`

SetPort sets Port field to given value.

### HasPort

`func (o *PortVlanAssignment) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetState

`func (o *PortVlanAssignment) GetState() PortVlanAssignmentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PortVlanAssignment) GetStateOk() (*PortVlanAssignmentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PortVlanAssignment) SetState(v PortVlanAssignmentState)`

SetState sets State field to given value.

### HasState

`func (o *PortVlanAssignment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortVlanAssignment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortVlanAssignment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortVlanAssignment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortVlanAssignment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *PortVlanAssignment) GetVirtualNetwork() Href`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *PortVlanAssignment) GetVirtualNetworkOk() (*Href, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *PortVlanAssignment) SetVirtualNetwork(v Href)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *PortVlanAssignment) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetVlan

`func (o *PortVlanAssignment) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PortVlanAssignment) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PortVlanAssignment) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PortVlanAssignment) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


