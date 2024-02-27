# LinkProtocolQinq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**LinkProtocolType**](LinkProtocolType.md) |  | [optional] 
**InnerTagProtocolId** | **int32** | Inner tag protocol identifier | 
**OuterTagProtocolId** | **int32** | Outer tag protocol identifier | 
**VlanCTag** | **int32** | Inner tag, i.e., C-VLAN tag | 
**VlanSTag** | **int32** | Outer tag, i.e., S-VLAN tag | 
**VlanCTagMin** | Pointer to **int32** | Outer tag Min value specified for QINQ connections | [optional] 
**VlanCTagMax** | Pointer to **int32** | Outer tag Max value specified for QINQ connections | [optional] 
**SubInterface** | Pointer to **int32** | Subinterface identifier | [optional] 

## Methods

### NewLinkProtocolQinq

`func NewLinkProtocolQinq(innerTagProtocolId int32, outerTagProtocolId int32, vlanCTag int32, vlanSTag int32, ) *LinkProtocolQinq`

NewLinkProtocolQinq instantiates a new LinkProtocolQinq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkProtocolQinqWithDefaults

`func NewLinkProtocolQinqWithDefaults() *LinkProtocolQinq`

NewLinkProtocolQinqWithDefaults instantiates a new LinkProtocolQinq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LinkProtocolQinq) GetType() LinkProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkProtocolQinq) GetTypeOk() (*LinkProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkProtocolQinq) SetType(v LinkProtocolType)`

SetType sets Type field to given value.

### HasType

`func (o *LinkProtocolQinq) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInnerTagProtocolId

`func (o *LinkProtocolQinq) GetInnerTagProtocolId() int32`

GetInnerTagProtocolId returns the InnerTagProtocolId field if non-nil, zero value otherwise.

### GetInnerTagProtocolIdOk

`func (o *LinkProtocolQinq) GetInnerTagProtocolIdOk() (*int32, bool)`

GetInnerTagProtocolIdOk returns a tuple with the InnerTagProtocolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerTagProtocolId

`func (o *LinkProtocolQinq) SetInnerTagProtocolId(v int32)`

SetInnerTagProtocolId sets InnerTagProtocolId field to given value.


### GetOuterTagProtocolId

`func (o *LinkProtocolQinq) GetOuterTagProtocolId() int32`

GetOuterTagProtocolId returns the OuterTagProtocolId field if non-nil, zero value otherwise.

### GetOuterTagProtocolIdOk

`func (o *LinkProtocolQinq) GetOuterTagProtocolIdOk() (*int32, bool)`

GetOuterTagProtocolIdOk returns a tuple with the OuterTagProtocolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterTagProtocolId

`func (o *LinkProtocolQinq) SetOuterTagProtocolId(v int32)`

SetOuterTagProtocolId sets OuterTagProtocolId field to given value.


### GetVlanCTag

`func (o *LinkProtocolQinq) GetVlanCTag() int32`

GetVlanCTag returns the VlanCTag field if non-nil, zero value otherwise.

### GetVlanCTagOk

`func (o *LinkProtocolQinq) GetVlanCTagOk() (*int32, bool)`

GetVlanCTagOk returns a tuple with the VlanCTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTag

`func (o *LinkProtocolQinq) SetVlanCTag(v int32)`

SetVlanCTag sets VlanCTag field to given value.


### GetVlanSTag

`func (o *LinkProtocolQinq) GetVlanSTag() int32`

GetVlanSTag returns the VlanSTag field if non-nil, zero value otherwise.

### GetVlanSTagOk

`func (o *LinkProtocolQinq) GetVlanSTagOk() (*int32, bool)`

GetVlanSTagOk returns a tuple with the VlanSTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSTag

`func (o *LinkProtocolQinq) SetVlanSTag(v int32)`

SetVlanSTag sets VlanSTag field to given value.


### GetVlanCTagMin

`func (o *LinkProtocolQinq) GetVlanCTagMin() int32`

GetVlanCTagMin returns the VlanCTagMin field if non-nil, zero value otherwise.

### GetVlanCTagMinOk

`func (o *LinkProtocolQinq) GetVlanCTagMinOk() (*int32, bool)`

GetVlanCTagMinOk returns a tuple with the VlanCTagMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTagMin

`func (o *LinkProtocolQinq) SetVlanCTagMin(v int32)`

SetVlanCTagMin sets VlanCTagMin field to given value.

### HasVlanCTagMin

`func (o *LinkProtocolQinq) HasVlanCTagMin() bool`

HasVlanCTagMin returns a boolean if a field has been set.

### GetVlanCTagMax

`func (o *LinkProtocolQinq) GetVlanCTagMax() int32`

GetVlanCTagMax returns the VlanCTagMax field if non-nil, zero value otherwise.

### GetVlanCTagMaxOk

`func (o *LinkProtocolQinq) GetVlanCTagMaxOk() (*int32, bool)`

GetVlanCTagMaxOk returns a tuple with the VlanCTagMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTagMax

`func (o *LinkProtocolQinq) SetVlanCTagMax(v int32)`

SetVlanCTagMax sets VlanCTagMax field to given value.

### HasVlanCTagMax

`func (o *LinkProtocolQinq) HasVlanCTagMax() bool`

HasVlanCTagMax returns a boolean if a field has been set.

### GetSubInterface

`func (o *LinkProtocolQinq) GetSubInterface() int32`

GetSubInterface returns the SubInterface field if non-nil, zero value otherwise.

### GetSubInterfaceOk

`func (o *LinkProtocolQinq) GetSubInterfaceOk() (*int32, bool)`

GetSubInterfaceOk returns a tuple with the SubInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubInterface

`func (o *LinkProtocolQinq) SetSubInterface(v int32)`

SetSubInterface sets SubInterface field to given value.

### HasSubInterface

`func (o *LinkProtocolQinq) HasSubInterface() bool`

HasSubInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


