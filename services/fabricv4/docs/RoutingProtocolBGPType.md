# RoutingProtocolBGPType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RoutingProtocolBGPTypeType**](RoutingProtocolBGPTypeType.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**BgpIpv4** | Pointer to [**BGPConnectionIpv4**](BGPConnectionIpv4.md) |  | [optional] 
**BgpIpv6** | Pointer to [**BGPConnectionIpv6**](BGPConnectionIpv6.md) |  | [optional] 
**CustomerAsn** | Pointer to **int64** | Customer asn | [optional] 
**BgpAuthKey** | Pointer to **string** | BGP authorization key | [optional] 
**AsOverrideEnabled** | Pointer to **bool** | Enable AS number override | [optional] 
**Bfd** | Pointer to [**RoutingProtocolBFD**](RoutingProtocolBFD.md) |  | [optional] 

## Methods

### NewRoutingProtocolBGPType

`func NewRoutingProtocolBGPType(type_ RoutingProtocolBGPTypeType, ) *RoutingProtocolBGPType`

NewRoutingProtocolBGPType instantiates a new RoutingProtocolBGPType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolBGPTypeWithDefaults

`func NewRoutingProtocolBGPTypeWithDefaults() *RoutingProtocolBGPType`

NewRoutingProtocolBGPTypeWithDefaults instantiates a new RoutingProtocolBGPType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutingProtocolBGPType) GetType() RoutingProtocolBGPTypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolBGPType) GetTypeOk() (*RoutingProtocolBGPTypeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolBGPType) SetType(v RoutingProtocolBGPTypeType)`

SetType sets Type field to given value.


### GetName

`func (o *RoutingProtocolBGPType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolBGPType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolBGPType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoutingProtocolBGPType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBgpIpv4

`func (o *RoutingProtocolBGPType) GetBgpIpv4() BGPConnectionIpv4`

GetBgpIpv4 returns the BgpIpv4 field if non-nil, zero value otherwise.

### GetBgpIpv4Ok

`func (o *RoutingProtocolBGPType) GetBgpIpv4Ok() (*BGPConnectionIpv4, bool)`

GetBgpIpv4Ok returns a tuple with the BgpIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv4

`func (o *RoutingProtocolBGPType) SetBgpIpv4(v BGPConnectionIpv4)`

SetBgpIpv4 sets BgpIpv4 field to given value.

### HasBgpIpv4

`func (o *RoutingProtocolBGPType) HasBgpIpv4() bool`

HasBgpIpv4 returns a boolean if a field has been set.

### GetBgpIpv6

`func (o *RoutingProtocolBGPType) GetBgpIpv6() BGPConnectionIpv6`

GetBgpIpv6 returns the BgpIpv6 field if non-nil, zero value otherwise.

### GetBgpIpv6Ok

`func (o *RoutingProtocolBGPType) GetBgpIpv6Ok() (*BGPConnectionIpv6, bool)`

GetBgpIpv6Ok returns a tuple with the BgpIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv6

`func (o *RoutingProtocolBGPType) SetBgpIpv6(v BGPConnectionIpv6)`

SetBgpIpv6 sets BgpIpv6 field to given value.

### HasBgpIpv6

`func (o *RoutingProtocolBGPType) HasBgpIpv6() bool`

HasBgpIpv6 returns a boolean if a field has been set.

### GetCustomerAsn

`func (o *RoutingProtocolBGPType) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *RoutingProtocolBGPType) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *RoutingProtocolBGPType) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.

### HasCustomerAsn

`func (o *RoutingProtocolBGPType) HasCustomerAsn() bool`

HasCustomerAsn returns a boolean if a field has been set.

### GetBgpAuthKey

`func (o *RoutingProtocolBGPType) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *RoutingProtocolBGPType) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *RoutingProtocolBGPType) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.

### HasBgpAuthKey

`func (o *RoutingProtocolBGPType) HasBgpAuthKey() bool`

HasBgpAuthKey returns a boolean if a field has been set.

### GetAsOverrideEnabled

`func (o *RoutingProtocolBGPType) GetAsOverrideEnabled() bool`

GetAsOverrideEnabled returns the AsOverrideEnabled field if non-nil, zero value otherwise.

### GetAsOverrideEnabledOk

`func (o *RoutingProtocolBGPType) GetAsOverrideEnabledOk() (*bool, bool)`

GetAsOverrideEnabledOk returns a tuple with the AsOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOverrideEnabled

`func (o *RoutingProtocolBGPType) SetAsOverrideEnabled(v bool)`

SetAsOverrideEnabled sets AsOverrideEnabled field to given value.

### HasAsOverrideEnabled

`func (o *RoutingProtocolBGPType) HasAsOverrideEnabled() bool`

HasAsOverrideEnabled returns a boolean if a field has been set.

### GetBfd

`func (o *RoutingProtocolBGPType) GetBfd() RoutingProtocolBFD`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *RoutingProtocolBGPType) GetBfdOk() (*RoutingProtocolBFD, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *RoutingProtocolBGPType) SetBfd(v RoutingProtocolBFD)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *RoutingProtocolBGPType) HasBfd() bool`

HasBfd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


