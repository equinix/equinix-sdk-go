# SimplifiedLinkProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**LinkProtocolType**](LinkProtocolType.md) |  | [optional] 
**VlanTag** | Pointer to **int32** | vlanTag value specified for DOT1Q connections | [optional] 
**VlanSTag** | Pointer to **int32** | vlanSTag value specified for QINQ connections | [optional] 
**VlanCTag** | Pointer to **int32** | vlanCTag value specified for QINQ connections | [optional] 

## Methods

### NewSimplifiedLinkProtocol

`func NewSimplifiedLinkProtocol() *SimplifiedLinkProtocol`

NewSimplifiedLinkProtocol instantiates a new SimplifiedLinkProtocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedLinkProtocolWithDefaults

`func NewSimplifiedLinkProtocolWithDefaults() *SimplifiedLinkProtocol`

NewSimplifiedLinkProtocolWithDefaults instantiates a new SimplifiedLinkProtocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SimplifiedLinkProtocol) GetType() LinkProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedLinkProtocol) GetTypeOk() (*LinkProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedLinkProtocol) SetType(v LinkProtocolType)`

SetType sets Type field to given value.

### HasType

`func (o *SimplifiedLinkProtocol) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlanTag

`func (o *SimplifiedLinkProtocol) GetVlanTag() int32`

GetVlanTag returns the VlanTag field if non-nil, zero value otherwise.

### GetVlanTagOk

`func (o *SimplifiedLinkProtocol) GetVlanTagOk() (*int32, bool)`

GetVlanTagOk returns a tuple with the VlanTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTag

`func (o *SimplifiedLinkProtocol) SetVlanTag(v int32)`

SetVlanTag sets VlanTag field to given value.

### HasVlanTag

`func (o *SimplifiedLinkProtocol) HasVlanTag() bool`

HasVlanTag returns a boolean if a field has been set.

### GetVlanSTag

`func (o *SimplifiedLinkProtocol) GetVlanSTag() int32`

GetVlanSTag returns the VlanSTag field if non-nil, zero value otherwise.

### GetVlanSTagOk

`func (o *SimplifiedLinkProtocol) GetVlanSTagOk() (*int32, bool)`

GetVlanSTagOk returns a tuple with the VlanSTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSTag

`func (o *SimplifiedLinkProtocol) SetVlanSTag(v int32)`

SetVlanSTag sets VlanSTag field to given value.

### HasVlanSTag

`func (o *SimplifiedLinkProtocol) HasVlanSTag() bool`

HasVlanSTag returns a boolean if a field has been set.

### GetVlanCTag

`func (o *SimplifiedLinkProtocol) GetVlanCTag() int32`

GetVlanCTag returns the VlanCTag field if non-nil, zero value otherwise.

### GetVlanCTagOk

`func (o *SimplifiedLinkProtocol) GetVlanCTagOk() (*int32, bool)`

GetVlanCTagOk returns a tuple with the VlanCTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCTag

`func (o *SimplifiedLinkProtocol) SetVlanCTag(v int32)`

SetVlanCTag sets VlanCTag field to given value.

### HasVlanCTag

`func (o *SimplifiedLinkProtocol) HasVlanCTag() bool`

HasVlanCTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


