# LinkProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**LinkProtocolType**](LinkProtocolType.md) |  | 
**Ipv4** | Pointer to [**LinkProtocolIpv4Ipv6Config**](LinkProtocolIpv4Ipv6Config.md) |  | [optional] 
**Ipv6** | Pointer to [**LinkProtocolIpv4Ipv6Config**](LinkProtocolIpv4Ipv6Config.md) |  | [optional] 
**TagProtocolId** | Pointer to **string** | Tag protocol identifier | [optional] 
**VlanTag** | **string** | VLAN tag | 
**VlanTagMin** | Pointer to **int32** | VLAN tag Min value specified for DOT1Q connections | [optional] 
**VlanTagMax** | Pointer to **int32** | VLAN tag Max value specified for DOT1Q connections | [optional] 
**InnerTagProtocolId** | **int32** | Inner tag protocol identifier | 
**OuterTagProtocolId** | **int32** | Outer tag protocol identifier | 
**VlanCTag** | **int32** | Inner tag, i.e., C-VLAN tag | 
**VlanSTag** | **int32** | Outer tag, i.e., S-VLAN tag | 
**VlanCTagMin** | Pointer to **int32** | Outer tag Min value specified for QINQ connections | [optional] 
**VlanCTagMax** | Pointer to **int32** | Outer tag Max value specified for QINQ connections | [optional] 
**SubInterface** | Pointer to **int32** | Subinterface identifier | [optional] 
**Vni** | **int32** | Virtual Network Identifier | 
**Vnid** | **int32** | Virtual Network Identifier | 
**Type5vni** | **int32** | Type 5 VNI identifier | 

## Methods

### NewLinkProtocol

`func NewLinkProtocol(type_ LinkProtocolType, vlanTag string, innerTagProtocolId int32, outerTagProtocolId int32, vlanCTag int32, vlanSTag int32, vni int32, vnid int32, type5vni int32, ) *LinkProtocol`

NewLinkProtocol instantiates a new LinkProtocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkProtocolWithDefaults

`func NewLinkProtocolWithDefaults() *LinkProtocol`

NewLinkProtocolWithDefaults instantiates a new LinkProtocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LinkProtocol) GetType() LinkProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkProtocol) GetTypeOk() (*LinkProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkProtocol) SetType(v LinkProtocolType)`

SetType sets Type field to given value.


### GetIpv4

`func (o *LinkProtocol) GetIpv4() LinkProtocolIpv4Ipv6Config`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *LinkProtocol) GetIpv4Ok() (*LinkProtocolIpv4Ipv6Config, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *LinkProtocol) SetIpv4(v LinkProtocolIpv4Ipv6Config)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *LinkProtocol) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *LinkProtocol) GetIpv6() LinkProtocolIpv4Ipv6Config`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *LinkProtocol) GetIpv6Ok() (*LinkProtocolIpv4Ipv6Config, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *LinkProtocol) SetIpv6(v LinkProtocolIpv4Ipv6Config)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *LinkProtocol) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetTagProtocolId

`func (o *LinkProtocol) GetTagProtocolId() string`

GetTagProtocolId returns the TagProtocolId field if non-nil, zero value otherwise.

### GetTagProtocolIdOk

`func (o *LinkProtocol) GetTagProtocolIdOk() (*string, bool)`

GetTagProtocolIdOk returns a tuple with the TagProtocolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagProtocolId

`func (o *LinkProtocol) SetTagProtocolId(v string)`

SetTagProtocolId sets TagProtocolId field to given value.

### HasTagProtocolId

`func (o *LinkProtocol) HasTagProtocolId() bool`

HasTagProtocolId returns a boolean if a field has been set.

### GetVlanTag

`func (o *LinkProtocol) GetVlanTag() string`

GetVlanTag returns the VlanTag field if non-nil, zero value otherwise.

### GetVlanTagOk

`func (o *LinkProtocol) GetVlanTagOk() (*string, bool)`

GetVlanTagOk returns a tuple with the VlanTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTag

`func (o *LinkProtocol) SetVlanTag(v string)`

SetVlanTag sets VlanTag field to given value.


### GetVlanTagMin

`func (o *LinkProtocol) GetVlanTagMin() int32`

GetVlanTagMin returns the VlanTagMin field if non-nil, zero value otherwise.

### GetVlanTagMinOk

`func (o *LinkProtocol) GetVlanTagMinOk() (*int32, bool)`

GetVlanTagMinOk returns a tuple with the VlanTagMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagMin

`func (o *LinkProtocol) SetVlanTagMin(v int32)`

SetVlanTagMin sets VlanTagMin field to given value.

### HasVlanTagMin

`func (o *LinkProtocol) HasVlanTagMin() bool`

HasVlanTagMin returns a boolean if a field has been set.

### GetVlanTagMax

`func (o *LinkProtocol) GetVlanTagMax() int32`

GetVlanTagMax returns the VlanTagMax field if non-nil, zero value otherwise.

### GetVlanTagMaxOk

`func (o *LinkProtocol) GetVlanTagMaxOk() (*int32, bool)`

GetVlanTagMaxOk returns a tuple with the VlanTagMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagMax

`func (o *LinkProtocol) SetVlanTagMax(v int32)`

SetVlanTagMax sets VlanTagMax field to given value.

### HasVlanTagMax

`func (o *LinkProtocol) HasVlanTagMax() bool`

HasVlanTagMax returns a boolean if a field has been set.

### GetInnerTagProtocolId

`func (o *LinkProtocol) GetInnerTagProtocolId() int32`

GetInnerTagProtocolId returns the InnerTagProtocolId field if non-nil, zero value otherwise.

### GetInnerTagProtocolIdOk

`func (o *LinkProtocol) GetInnerTagProtocolIdOk() (*int32, bool)`

GetInnerTagProtocolIdOk returns a tuple with the InnerTagProtocolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerTagProtocolId

`func (o *LinkProtocol) SetInnerTagProtocolId(v int32)`

SetInnerTagProtocolId sets InnerTagProtocolId field to given value.


### GetOuterTagProtocolId

`func (o *LinkProtocol) GetOuterTagProtocolId() int32`

GetOuterTagProtocolId returns the OuterTagProtocolId field if non-nil, zero value otherwise.

### GetOuterTagProtocolIdOk

`func (o *LinkProtocol) GetOuterTagProtocolIdOk() (*int32, bool)`

GetOuterTagProtocolIdOk returns a tuple with the OuterTagProtocolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterTagProtocolId

`func (o *LinkProtocol) SetOuterTagProtocolId(v int32)`

SetOuterTagProtocolId sets OuterTagProtocolId field to given value.


### GetVlanCTag

`func (o *LinkProtocol) GetVlanCTag() int32`

GetVlanCTag returns the VlanCTag field if non-nil, zero value otherwise.

### GetVlanCTagOk

`func (o *LinkProtocol) GetVlanCTagOk() (*int32, bool)`

GetVlanCTagOk returns a tuple with the VlanCTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTag

`func (o *LinkProtocol) SetVlanCTag(v int32)`

SetVlanCTag sets VlanCTag field to given value.


### GetVlanSTag

`func (o *LinkProtocol) GetVlanSTag() int32`

GetVlanSTag returns the VlanSTag field if non-nil, zero value otherwise.

### GetVlanSTagOk

`func (o *LinkProtocol) GetVlanSTagOk() (*int32, bool)`

GetVlanSTagOk returns a tuple with the VlanSTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSTag

`func (o *LinkProtocol) SetVlanSTag(v int32)`

SetVlanSTag sets VlanSTag field to given value.


### GetVlanCTagMin

`func (o *LinkProtocol) GetVlanCTagMin() int32`

GetVlanCTagMin returns the VlanCTagMin field if non-nil, zero value otherwise.

### GetVlanCTagMinOk

`func (o *LinkProtocol) GetVlanCTagMinOk() (*int32, bool)`

GetVlanCTagMinOk returns a tuple with the VlanCTagMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTagMin

`func (o *LinkProtocol) SetVlanCTagMin(v int32)`

SetVlanCTagMin sets VlanCTagMin field to given value.

### HasVlanCTagMin

`func (o *LinkProtocol) HasVlanCTagMin() bool`

HasVlanCTagMin returns a boolean if a field has been set.

### GetVlanCTagMax

`func (o *LinkProtocol) GetVlanCTagMax() int32`

GetVlanCTagMax returns the VlanCTagMax field if non-nil, zero value otherwise.

### GetVlanCTagMaxOk

`func (o *LinkProtocol) GetVlanCTagMaxOk() (*int32, bool)`

GetVlanCTagMaxOk returns a tuple with the VlanCTagMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTagMax

`func (o *LinkProtocol) SetVlanCTagMax(v int32)`

SetVlanCTagMax sets VlanCTagMax field to given value.

### HasVlanCTagMax

`func (o *LinkProtocol) HasVlanCTagMax() bool`

HasVlanCTagMax returns a boolean if a field has been set.

### GetSubInterface

`func (o *LinkProtocol) GetSubInterface() int32`

GetSubInterface returns the SubInterface field if non-nil, zero value otherwise.

### GetSubInterfaceOk

`func (o *LinkProtocol) GetSubInterfaceOk() (*int32, bool)`

GetSubInterfaceOk returns a tuple with the SubInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubInterface

`func (o *LinkProtocol) SetSubInterface(v int32)`

SetSubInterface sets SubInterface field to given value.

### HasSubInterface

`func (o *LinkProtocol) HasSubInterface() bool`

HasSubInterface returns a boolean if a field has been set.

### GetVni

`func (o *LinkProtocol) GetVni() int32`

GetVni returns the Vni field if non-nil, zero value otherwise.

### GetVniOk

`func (o *LinkProtocol) GetVniOk() (*int32, bool)`

GetVniOk returns a tuple with the Vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVni

`func (o *LinkProtocol) SetVni(v int32)`

SetVni sets Vni field to given value.


### GetVnid

`func (o *LinkProtocol) GetVnid() int32`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *LinkProtocol) GetVnidOk() (*int32, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *LinkProtocol) SetVnid(v int32)`

SetVnid sets Vnid field to given value.


### GetType5vni

`func (o *LinkProtocol) GetType5vni() int32`

GetType5vni returns the Type5vni field if non-nil, zero value otherwise.

### GetType5vniOk

`func (o *LinkProtocol) GetType5vniOk() (*int32, bool)`

GetType5vniOk returns a tuple with the Type5vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType5vni

`func (o *LinkProtocol) SetType5vni(v int32)`

SetType5vni sets Type5vni field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


