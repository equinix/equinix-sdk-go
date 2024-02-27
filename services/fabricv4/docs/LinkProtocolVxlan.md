# LinkProtocolVxlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**LinkProtocolType**](LinkProtocolType.md) |  | [optional] 
**Vni** | **int32** | Virtual Network Identifier | 

## Methods

### NewLinkProtocolVxlan

`func NewLinkProtocolVxlan(vni int32, ) *LinkProtocolVxlan`

NewLinkProtocolVxlan instantiates a new LinkProtocolVxlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkProtocolVxlanWithDefaults

`func NewLinkProtocolVxlanWithDefaults() *LinkProtocolVxlan`

NewLinkProtocolVxlanWithDefaults instantiates a new LinkProtocolVxlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LinkProtocolVxlan) GetType() LinkProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkProtocolVxlan) GetTypeOk() (*LinkProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkProtocolVxlan) SetType(v LinkProtocolType)`

SetType sets Type field to given value.

### HasType

`func (o *LinkProtocolVxlan) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVni

`func (o *LinkProtocolVxlan) GetVni() int32`

GetVni returns the Vni field if non-nil, zero value otherwise.

### GetVniOk

`func (o *LinkProtocolVxlan) GetVniOk() (*int32, bool)`

GetVniOk returns a tuple with the Vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVni

`func (o *LinkProtocolVxlan) SetVni(v int32)`

SetVni sets Vni field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


