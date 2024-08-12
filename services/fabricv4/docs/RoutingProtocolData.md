# RoutingProtocolData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Routing protocol type | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**BgpIpv4** | Pointer to [**BGPConnectionIpv4**](BGPConnectionIpv4.md) |  | [optional] 
**BgpIpv6** | Pointer to [**BGPConnectionIpv6**](BGPConnectionIpv6.md) |  | [optional] 
**CustomerAsn** | Pointer to **int64** | Customer asn | [optional] 
**EquinixAsn** | Pointer to **int64** | Equinix asn | [optional] 
**BgpAuthKey** | Pointer to **string** | BGP authorization key | [optional] 
**AsOverrideEnabled** | Pointer to **bool** | Enable AS number override | [optional] 
**Bfd** | Pointer to [**RoutingProtocolBFD**](RoutingProtocolBFD.md) |  | [optional] 
**Href** | Pointer to **string** | Routing Protocol URI | [optional] 
**Uuid** | Pointer to **string** | Routing protocol identifier | [optional] 
**State** | Pointer to [**RoutingProtocolBGPDataState**](RoutingProtocolBGPDataState.md) |  | [optional] 
**Operation** | Pointer to [**RoutingProtocolOperation**](RoutingProtocolOperation.md) |  | [optional] 
**Change** | Pointer to [**RoutingProtocolChange**](RoutingProtocolChange.md) |  | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**DirectIpv4** | Pointer to [**DirectConnectionIpv4**](DirectConnectionIpv4.md) |  | [optional] 
**DirectIpv6** | Pointer to [**DirectConnectionIpv6**](DirectConnectionIpv6.md) |  | [optional] 

## Methods

### NewRoutingProtocolData

`func NewRoutingProtocolData() *RoutingProtocolData`

NewRoutingProtocolData instantiates a new RoutingProtocolData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolDataWithDefaults

`func NewRoutingProtocolDataWithDefaults() *RoutingProtocolData`

NewRoutingProtocolDataWithDefaults instantiates a new RoutingProtocolData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutingProtocolData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoutingProtocolData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *RoutingProtocolData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoutingProtocolData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBgpIpv4

`func (o *RoutingProtocolData) GetBgpIpv4() BGPConnectionIpv4`

GetBgpIpv4 returns the BgpIpv4 field if non-nil, zero value otherwise.

### GetBgpIpv4Ok

`func (o *RoutingProtocolData) GetBgpIpv4Ok() (*BGPConnectionIpv4, bool)`

GetBgpIpv4Ok returns a tuple with the BgpIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv4

`func (o *RoutingProtocolData) SetBgpIpv4(v BGPConnectionIpv4)`

SetBgpIpv4 sets BgpIpv4 field to given value.

### HasBgpIpv4

`func (o *RoutingProtocolData) HasBgpIpv4() bool`

HasBgpIpv4 returns a boolean if a field has been set.

### GetBgpIpv6

`func (o *RoutingProtocolData) GetBgpIpv6() BGPConnectionIpv6`

GetBgpIpv6 returns the BgpIpv6 field if non-nil, zero value otherwise.

### GetBgpIpv6Ok

`func (o *RoutingProtocolData) GetBgpIpv6Ok() (*BGPConnectionIpv6, bool)`

GetBgpIpv6Ok returns a tuple with the BgpIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv6

`func (o *RoutingProtocolData) SetBgpIpv6(v BGPConnectionIpv6)`

SetBgpIpv6 sets BgpIpv6 field to given value.

### HasBgpIpv6

`func (o *RoutingProtocolData) HasBgpIpv6() bool`

HasBgpIpv6 returns a boolean if a field has been set.

### GetCustomerAsn

`func (o *RoutingProtocolData) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *RoutingProtocolData) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *RoutingProtocolData) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.

### HasCustomerAsn

`func (o *RoutingProtocolData) HasCustomerAsn() bool`

HasCustomerAsn returns a boolean if a field has been set.

### GetEquinixAsn

`func (o *RoutingProtocolData) GetEquinixAsn() int64`

GetEquinixAsn returns the EquinixAsn field if non-nil, zero value otherwise.

### GetEquinixAsnOk

`func (o *RoutingProtocolData) GetEquinixAsnOk() (*int64, bool)`

GetEquinixAsnOk returns a tuple with the EquinixAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixAsn

`func (o *RoutingProtocolData) SetEquinixAsn(v int64)`

SetEquinixAsn sets EquinixAsn field to given value.

### HasEquinixAsn

`func (o *RoutingProtocolData) HasEquinixAsn() bool`

HasEquinixAsn returns a boolean if a field has been set.

### GetBgpAuthKey

`func (o *RoutingProtocolData) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *RoutingProtocolData) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *RoutingProtocolData) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.

### HasBgpAuthKey

`func (o *RoutingProtocolData) HasBgpAuthKey() bool`

HasBgpAuthKey returns a boolean if a field has been set.

### GetAsOverrideEnabled

`func (o *RoutingProtocolData) GetAsOverrideEnabled() bool`

GetAsOverrideEnabled returns the AsOverrideEnabled field if non-nil, zero value otherwise.

### GetAsOverrideEnabledOk

`func (o *RoutingProtocolData) GetAsOverrideEnabledOk() (*bool, bool)`

GetAsOverrideEnabledOk returns a tuple with the AsOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOverrideEnabled

`func (o *RoutingProtocolData) SetAsOverrideEnabled(v bool)`

SetAsOverrideEnabled sets AsOverrideEnabled field to given value.

### HasAsOverrideEnabled

`func (o *RoutingProtocolData) HasAsOverrideEnabled() bool`

HasAsOverrideEnabled returns a boolean if a field has been set.

### GetBfd

`func (o *RoutingProtocolData) GetBfd() RoutingProtocolBFD`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *RoutingProtocolData) GetBfdOk() (*RoutingProtocolBFD, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *RoutingProtocolData) SetBfd(v RoutingProtocolBFD)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *RoutingProtocolData) HasBfd() bool`

HasBfd returns a boolean if a field has been set.

### GetHref

`func (o *RoutingProtocolData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RoutingProtocolData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RoutingProtocolData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RoutingProtocolData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *RoutingProtocolData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoutingProtocolData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoutingProtocolData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RoutingProtocolData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *RoutingProtocolData) GetState() RoutingProtocolBGPDataState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RoutingProtocolData) GetStateOk() (*RoutingProtocolBGPDataState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RoutingProtocolData) SetState(v RoutingProtocolBGPDataState)`

SetState sets State field to given value.

### HasState

`func (o *RoutingProtocolData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOperation

`func (o *RoutingProtocolData) GetOperation() RoutingProtocolOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *RoutingProtocolData) GetOperationOk() (*RoutingProtocolOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *RoutingProtocolData) SetOperation(v RoutingProtocolOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *RoutingProtocolData) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetChange

`func (o *RoutingProtocolData) GetChange() RoutingProtocolChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *RoutingProtocolData) GetChangeOk() (*RoutingProtocolChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *RoutingProtocolData) SetChange(v RoutingProtocolChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *RoutingProtocolData) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetChangelog

`func (o *RoutingProtocolData) GetChangelog() Changelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *RoutingProtocolData) GetChangelogOk() (*Changelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *RoutingProtocolData) SetChangelog(v Changelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *RoutingProtocolData) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.

### GetDirectIpv4

`func (o *RoutingProtocolData) GetDirectIpv4() DirectConnectionIpv4`

GetDirectIpv4 returns the DirectIpv4 field if non-nil, zero value otherwise.

### GetDirectIpv4Ok

`func (o *RoutingProtocolData) GetDirectIpv4Ok() (*DirectConnectionIpv4, bool)`

GetDirectIpv4Ok returns a tuple with the DirectIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectIpv4

`func (o *RoutingProtocolData) SetDirectIpv4(v DirectConnectionIpv4)`

SetDirectIpv4 sets DirectIpv4 field to given value.

### HasDirectIpv4

`func (o *RoutingProtocolData) HasDirectIpv4() bool`

HasDirectIpv4 returns a boolean if a field has been set.

### GetDirectIpv6

`func (o *RoutingProtocolData) GetDirectIpv6() DirectConnectionIpv6`

GetDirectIpv6 returns the DirectIpv6 field if non-nil, zero value otherwise.

### GetDirectIpv6Ok

`func (o *RoutingProtocolData) GetDirectIpv6Ok() (*DirectConnectionIpv6, bool)`

GetDirectIpv6Ok returns a tuple with the DirectIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectIpv6

`func (o *RoutingProtocolData) SetDirectIpv6(v DirectConnectionIpv6)`

SetDirectIpv6 sets DirectIpv6 field to given value.

### HasDirectIpv6

`func (o *RoutingProtocolData) HasDirectIpv6() bool`

HasDirectIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


