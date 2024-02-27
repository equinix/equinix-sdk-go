# RoutingProtocolBGPData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**RoutingProtocolBGPTypeType**](RoutingProtocolBGPTypeType.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**BgpIpv4** | Pointer to [**BGPConnectionIpv4**](BGPConnectionIpv4.md) |  | [optional] 
**BgpIpv6** | Pointer to [**BGPConnectionIpv6**](BGPConnectionIpv6.md) |  | [optional] 
**CustomerAsn** | Pointer to **int64** | Customer asn | [optional] 
**EquinixAsn** | Pointer to **int64** | Equinix asn | [optional] 
**BgpAuthKey** | Pointer to **string** | BGP authorization key | [optional] 
**Bfd** | Pointer to [**RoutingProtocolBFD**](RoutingProtocolBFD.md) |  | [optional] 
**Href** | Pointer to **string** | Routing Protocol URI | [optional] 
**Uuid** | Pointer to **string** | Routing protocol identifier | [optional] 
**State** | Pointer to [**RoutingProtocolBGPDataState**](RoutingProtocolBGPDataState.md) |  | [optional] 
**Operation** | Pointer to [**RoutingProtocolOperation**](RoutingProtocolOperation.md) |  | [optional] 
**Change** | Pointer to [**RoutingProtocolChange**](RoutingProtocolChange.md) |  | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewRoutingProtocolBGPData

`func NewRoutingProtocolBGPData() *RoutingProtocolBGPData`

NewRoutingProtocolBGPData instantiates a new RoutingProtocolBGPData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolBGPDataWithDefaults

`func NewRoutingProtocolBGPDataWithDefaults() *RoutingProtocolBGPData`

NewRoutingProtocolBGPDataWithDefaults instantiates a new RoutingProtocolBGPData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutingProtocolBGPData) GetType() RoutingProtocolBGPTypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolBGPData) GetTypeOk() (*RoutingProtocolBGPTypeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolBGPData) SetType(v RoutingProtocolBGPTypeType)`

SetType sets Type field to given value.

### HasType

`func (o *RoutingProtocolBGPData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *RoutingProtocolBGPData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolBGPData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolBGPData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoutingProtocolBGPData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBgpIpv4

`func (o *RoutingProtocolBGPData) GetBgpIpv4() BGPConnectionIpv4`

GetBgpIpv4 returns the BgpIpv4 field if non-nil, zero value otherwise.

### GetBgpIpv4Ok

`func (o *RoutingProtocolBGPData) GetBgpIpv4Ok() (*BGPConnectionIpv4, bool)`

GetBgpIpv4Ok returns a tuple with the BgpIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv4

`func (o *RoutingProtocolBGPData) SetBgpIpv4(v BGPConnectionIpv4)`

SetBgpIpv4 sets BgpIpv4 field to given value.

### HasBgpIpv4

`func (o *RoutingProtocolBGPData) HasBgpIpv4() bool`

HasBgpIpv4 returns a boolean if a field has been set.

### GetBgpIpv6

`func (o *RoutingProtocolBGPData) GetBgpIpv6() BGPConnectionIpv6`

GetBgpIpv6 returns the BgpIpv6 field if non-nil, zero value otherwise.

### GetBgpIpv6Ok

`func (o *RoutingProtocolBGPData) GetBgpIpv6Ok() (*BGPConnectionIpv6, bool)`

GetBgpIpv6Ok returns a tuple with the BgpIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv6

`func (o *RoutingProtocolBGPData) SetBgpIpv6(v BGPConnectionIpv6)`

SetBgpIpv6 sets BgpIpv6 field to given value.

### HasBgpIpv6

`func (o *RoutingProtocolBGPData) HasBgpIpv6() bool`

HasBgpIpv6 returns a boolean if a field has been set.

### GetCustomerAsn

`func (o *RoutingProtocolBGPData) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *RoutingProtocolBGPData) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *RoutingProtocolBGPData) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.

### HasCustomerAsn

`func (o *RoutingProtocolBGPData) HasCustomerAsn() bool`

HasCustomerAsn returns a boolean if a field has been set.

### GetEquinixAsn

`func (o *RoutingProtocolBGPData) GetEquinixAsn() int64`

GetEquinixAsn returns the EquinixAsn field if non-nil, zero value otherwise.

### GetEquinixAsnOk

`func (o *RoutingProtocolBGPData) GetEquinixAsnOk() (*int64, bool)`

GetEquinixAsnOk returns a tuple with the EquinixAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixAsn

`func (o *RoutingProtocolBGPData) SetEquinixAsn(v int64)`

SetEquinixAsn sets EquinixAsn field to given value.

### HasEquinixAsn

`func (o *RoutingProtocolBGPData) HasEquinixAsn() bool`

HasEquinixAsn returns a boolean if a field has been set.

### GetBgpAuthKey

`func (o *RoutingProtocolBGPData) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *RoutingProtocolBGPData) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *RoutingProtocolBGPData) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.

### HasBgpAuthKey

`func (o *RoutingProtocolBGPData) HasBgpAuthKey() bool`

HasBgpAuthKey returns a boolean if a field has been set.

### GetBfd

`func (o *RoutingProtocolBGPData) GetBfd() RoutingProtocolBFD`

GetBfd returns the Bfd field if non-nil, zero value otherwise.

### GetBfdOk

`func (o *RoutingProtocolBGPData) GetBfdOk() (*RoutingProtocolBFD, bool)`

GetBfdOk returns a tuple with the Bfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfd

`func (o *RoutingProtocolBGPData) SetBfd(v RoutingProtocolBFD)`

SetBfd sets Bfd field to given value.

### HasBfd

`func (o *RoutingProtocolBGPData) HasBfd() bool`

HasBfd returns a boolean if a field has been set.

### GetHref

`func (o *RoutingProtocolBGPData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RoutingProtocolBGPData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RoutingProtocolBGPData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RoutingProtocolBGPData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *RoutingProtocolBGPData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoutingProtocolBGPData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoutingProtocolBGPData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RoutingProtocolBGPData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *RoutingProtocolBGPData) GetState() RoutingProtocolBGPDataState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RoutingProtocolBGPData) GetStateOk() (*RoutingProtocolBGPDataState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RoutingProtocolBGPData) SetState(v RoutingProtocolBGPDataState)`

SetState sets State field to given value.

### HasState

`func (o *RoutingProtocolBGPData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOperation

`func (o *RoutingProtocolBGPData) GetOperation() RoutingProtocolOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *RoutingProtocolBGPData) GetOperationOk() (*RoutingProtocolOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *RoutingProtocolBGPData) SetOperation(v RoutingProtocolOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *RoutingProtocolBGPData) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetChange

`func (o *RoutingProtocolBGPData) GetChange() RoutingProtocolChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *RoutingProtocolBGPData) GetChangeOk() (*RoutingProtocolChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *RoutingProtocolBGPData) SetChange(v RoutingProtocolChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *RoutingProtocolBGPData) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetChangelog

`func (o *RoutingProtocolBGPData) GetChangelog() Changelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *RoutingProtocolBGPData) GetChangelogOk() (*Changelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *RoutingProtocolBGPData) SetChangelog(v Changelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *RoutingProtocolBGPData) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


