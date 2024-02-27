# LinkProtocolDot1q

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**LinkProtocolType**](LinkProtocolType.md) |  | [optional] 
**TagProtocolId** | Pointer to **string** | Tag protocol identifier | [optional] 
**VlanTag** | **string** | VLAN tag | 
**VlanTagMin** | Pointer to **int32** | VLAN tag Min value specified for DOT1Q connections | [optional] 
**VlanTagMax** | Pointer to **int32** | VLAN tag Max value specified for DOT1Q connections | [optional] 

## Methods

### NewLinkProtocolDot1q

`func NewLinkProtocolDot1q(vlanTag string, ) *LinkProtocolDot1q`

NewLinkProtocolDot1q instantiates a new LinkProtocolDot1q object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkProtocolDot1qWithDefaults

`func NewLinkProtocolDot1qWithDefaults() *LinkProtocolDot1q`

NewLinkProtocolDot1qWithDefaults instantiates a new LinkProtocolDot1q object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LinkProtocolDot1q) GetType() LinkProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkProtocolDot1q) GetTypeOk() (*LinkProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkProtocolDot1q) SetType(v LinkProtocolType)`

SetType sets Type field to given value.

### HasType

`func (o *LinkProtocolDot1q) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTagProtocolId

`func (o *LinkProtocolDot1q) GetTagProtocolId() string`

GetTagProtocolId returns the TagProtocolId field if non-nil, zero value otherwise.

### GetTagProtocolIdOk

`func (o *LinkProtocolDot1q) GetTagProtocolIdOk() (*string, bool)`

GetTagProtocolIdOk returns a tuple with the TagProtocolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagProtocolId

`func (o *LinkProtocolDot1q) SetTagProtocolId(v string)`

SetTagProtocolId sets TagProtocolId field to given value.

### HasTagProtocolId

`func (o *LinkProtocolDot1q) HasTagProtocolId() bool`

HasTagProtocolId returns a boolean if a field has been set.

### GetVlanTag

`func (o *LinkProtocolDot1q) GetVlanTag() string`

GetVlanTag returns the VlanTag field if non-nil, zero value otherwise.

### GetVlanTagOk

`func (o *LinkProtocolDot1q) GetVlanTagOk() (*string, bool)`

GetVlanTagOk returns a tuple with the VlanTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTag

`func (o *LinkProtocolDot1q) SetVlanTag(v string)`

SetVlanTag sets VlanTag field to given value.


### GetVlanTagMin

`func (o *LinkProtocolDot1q) GetVlanTagMin() int32`

GetVlanTagMin returns the VlanTagMin field if non-nil, zero value otherwise.

### GetVlanTagMinOk

`func (o *LinkProtocolDot1q) GetVlanTagMinOk() (*int32, bool)`

GetVlanTagMinOk returns a tuple with the VlanTagMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagMin

`func (o *LinkProtocolDot1q) SetVlanTagMin(v int32)`

SetVlanTagMin sets VlanTagMin field to given value.

### HasVlanTagMin

`func (o *LinkProtocolDot1q) HasVlanTagMin() bool`

HasVlanTagMin returns a boolean if a field has been set.

### GetVlanTagMax

`func (o *LinkProtocolDot1q) GetVlanTagMax() int32`

GetVlanTagMax returns the VlanTagMax field if non-nil, zero value otherwise.

### GetVlanTagMaxOk

`func (o *LinkProtocolDot1q) GetVlanTagMaxOk() (*int32, bool)`

GetVlanTagMaxOk returns a tuple with the VlanTagMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagMax

`func (o *LinkProtocolDot1q) SetVlanTagMax(v int32)`

SetVlanTagMax sets VlanTagMax field to given value.

### HasVlanTagMax

`func (o *LinkProtocolDot1q) HasVlanTagMax() bool`

HasVlanTagMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


