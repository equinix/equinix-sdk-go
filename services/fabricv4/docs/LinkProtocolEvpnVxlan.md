# LinkProtocolEvpnVxlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**LinkProtocolType**](LinkProtocolType.md) |  | [optional] 
**Vnid** | **int32** | Virtual Network Identifier | 
**Type5vni** | **int32** | Type 5 VNI identifier | 

## Methods

### NewLinkProtocolEvpnVxlan

`func NewLinkProtocolEvpnVxlan(vnid int32, type5vni int32, ) *LinkProtocolEvpnVxlan`

NewLinkProtocolEvpnVxlan instantiates a new LinkProtocolEvpnVxlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkProtocolEvpnVxlanWithDefaults

`func NewLinkProtocolEvpnVxlanWithDefaults() *LinkProtocolEvpnVxlan`

NewLinkProtocolEvpnVxlanWithDefaults instantiates a new LinkProtocolEvpnVxlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LinkProtocolEvpnVxlan) GetType() LinkProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkProtocolEvpnVxlan) GetTypeOk() (*LinkProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkProtocolEvpnVxlan) SetType(v LinkProtocolType)`

SetType sets Type field to given value.

### HasType

`func (o *LinkProtocolEvpnVxlan) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVnid

`func (o *LinkProtocolEvpnVxlan) GetVnid() int32`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *LinkProtocolEvpnVxlan) GetVnidOk() (*int32, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *LinkProtocolEvpnVxlan) SetVnid(v int32)`

SetVnid sets Vnid field to given value.


### GetType5vni

`func (o *LinkProtocolEvpnVxlan) GetType5vni() int32`

GetType5vni returns the Type5vni field if non-nil, zero value otherwise.

### GetType5vniOk

`func (o *LinkProtocolEvpnVxlan) GetType5vniOk() (*int32, bool)`

GetType5vniOk returns a tuple with the Type5vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType5vni

`func (o *LinkProtocolEvpnVxlan) SetType5vni(v int32)`

SetType5vni sets Type5vni field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


