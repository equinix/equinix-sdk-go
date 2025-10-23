# RoutingProtocolBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Routing protocol type | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**BgpIpv4** | Pointer to [**BGPConnectionIpv4**](BGPConnectionIpv4.md) |  | [optional] 
**BgpIpv6** | Pointer to [**BGPConnectionIpv6**](BGPConnectionIpv6.md) |  | [optional] 
**CustomerAsn** | Pointer to **int64** | Customer asn | [optional] 
**BgpAuthKey** | Pointer to **string** | BGP authorization key | [optional] 
**AsOverrideEnabled** | Pointer to **bool** | Enable AS number override | [optional] 
**Bfd** | Pointer to [**RoutingProtocolBFD**](RoutingProtocolBFD.md) |  | [optional] 
**DirectIpv4** | Pointer to [**DirectConnectionIpv4**](DirectConnectionIpv4.md) |  | [optional] 
**DirectIpv6** | Pointer to [**DirectConnectionIpv6**](DirectConnectionIpv6.md) |  | [optional] 

## Methods

### NewRoutingProtocolBase

`func NewRoutingProtocolBase() *RoutingProtocolBase`

NewRoutingProtocolBase instantiates a new RoutingProtocolBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolBaseWithDefaults

`func NewRoutingProtocolBaseWithDefaults() *RoutingProtocolBase`

NewRoutingProtocolBaseWithDefaults instantiates a new RoutingProtocolBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutingProtocolBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolBase) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoutingProtocolBase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *RoutingProtocolBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoutingProtocolBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBgpIpv4

`func (o *RoutingProtocolBase) GetBgpIpv4() BGPConnectionIpv4`

GetBgpIpv4 returns the BgpIpv4 field if non-nil, zero value otherwise.

### GetBgpIpv4Ok

`func (o *RoutingProtocolBase) GetBgpIpv4Ok() (*BGPConnectionIpv4, bool)`

GetBgpIpv4Ok returns a tuple with the BgpIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv4

`func (o *RoutingProtocolBase) SetBgpIpv4(v BGPConnectionIpv4)`

SetBgpIpv4 sets BgpIpv4 field to given value.

### HasBgpIpv4

`func (o *RoutingProtocolBase) HasBgpIpv4() bool`

HasBgpIpv4 returns a boolean if a field has been set.

### GetBgpIpv6

`func (o *RoutingProtocolBase) GetBgpIpv6() BGPConnectionIpv6`

GetBgpIpv6 returns the BgpIpv6 field if non-nil, zero value otherwise.

### GetBgpIpv6Ok

`func (o *RoutingProtocolBase) GetBgpIpv6Ok() (*BGPConnectionIpv6, bool)`

GetBgpIpv6Ok returns a tuple with the BgpIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv6

`func (o *RoutingProtocolBase) SetBgpIpv6(v BGPConnectionIpv6)`

SetBgpIpv6 sets BgpIpv6 field to given value.

### HasBgpIpv6

`func (o *RoutingProtocolBase) HasBgpIpv6() bool`

HasBgpIpv6 returns a boolean if a field has been set.

### GetCustomerAsn

`func (o *RoutingProtocolBase) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *RoutingProtocolBase) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *RoutingProtocolBase) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.

### HasCustomerAsn

`func (o *RoutingProtocolBase) HasCustomerAsn() bool`

HasCustomerAsn returns a boolean if a field has been set.

### GetBgpAuthKey

`func (o *RoutingProtocolBase) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *RoutingProtocolBase) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *RoutingProtocolBase) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.

### HasBgpAuthKey

`func (o *RoutingProtocolBase) HasBgpAuthKey() bool`

HasBgpAuthKey returns a boolean if a field has been set.

### GetAsOverrideEnabled

`func (o *RoutingProtocolBase) GetAsOverrideEnabled() bool`

GetAsOverrideEnabled returns the AsOverrideEnabled field if non-nil, zero value otherwise.

### GetAsOverrideEnabledOk

`func (o *RoutingProtocolBase) GetAsOverrideEnabledOk() (*bool, bool)`

GetAsOverrideEnabledOk returns a tuple with the AsOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOverrideEnabled

`func (o *RoutingProtocolBase) SetAsOverrideEnabled(v bool)`

SetAsOverrideEnabled sets AsOverrideEnabled field to given value.

### HasAsOverrideEnabled

`func (o *RoutingProtocolBase) HasAsOverrideEnabled() bool`

HasAsOverrideEnabled returns a boolean if a field has been set.

### GetBfd

`func (o *RoutingProtocolBase) GetBfd() RoutingProtocolBFD`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *RoutingProtocolBase) GetBfdOk() (*RoutingProtocolBFD, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *RoutingProtocolBase) SetBfd(v RoutingProtocolBFD)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *RoutingProtocolBase) HasBfd() bool`

HasBfd returns a boolean if a field has been set.

### GetDirectIpv4

`func (o *RoutingProtocolBase) GetDirectIpv4() DirectConnectionIpv4`

GetDirectIpv4 returns the DirectIpv4 field if non-nil, zero value otherwise.

### GetDirectIpv4Ok

`func (o *RoutingProtocolBase) GetDirectIpv4Ok() (*DirectConnectionIpv4, bool)`

GetDirectIpv4Ok returns a tuple with the DirectIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectIpv4

`func (o *RoutingProtocolBase) SetDirectIpv4(v DirectConnectionIpv4)`

SetDirectIpv4 sets DirectIpv4 field to given value.

### HasDirectIpv4

`func (o *RoutingProtocolBase) HasDirectIpv4() bool`

HasDirectIpv4 returns a boolean if a field has been set.

### GetDirectIpv6

`func (o *RoutingProtocolBase) GetDirectIpv6() DirectConnectionIpv6`

GetDirectIpv6 returns the DirectIpv6 field if non-nil, zero value otherwise.

### GetDirectIpv6Ok

`func (o *RoutingProtocolBase) GetDirectIpv6Ok() (*DirectConnectionIpv6, bool)`

GetDirectIpv6Ok returns a tuple with the DirectIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectIpv6

`func (o *RoutingProtocolBase) SetDirectIpv6(v DirectConnectionIpv6)`

SetDirectIpv6 sets DirectIpv6 field to given value.

### HasDirectIpv6

`func (o *RoutingProtocolBase) HasDirectIpv6() bool`

HasDirectIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


