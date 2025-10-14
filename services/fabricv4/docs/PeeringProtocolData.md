# PeeringProtocolData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Peering Protocol URI | [optional] 
**Uuid** | Pointer to **string** | Peering protocol identifier | [optional] 
**Type** | Pointer to [**PeeringProtocolDataType**](PeeringProtocolDataType.md) |  | [optional] 
**Name** | Pointer to **string** | Protocol Name | [optional] 
**Description** | Pointer to **string** | Protocol Description | [optional] 
**CustomerAsn** | Pointer to **int64** | Customer ASN | [optional] 
**EquinixAsn** | Pointer to **int64** | Equinix ASN | [optional] 
**State** | Pointer to [**PeeringProtocolDataState**](PeeringProtocolDataState.md) |  | [optional] 
**MacAddress** | Pointer to **string** | MAC Address of The Peering Protocol | [optional] 
**BgpIpv4** | Pointer to [**PeeringConnectionResIpv4**](PeeringConnectionResIpv4.md) |  | [optional] 
**BgpIpv6** | Pointer to [**PeeringConnectionResIpv6**](PeeringConnectionResIpv6.md) |  | [optional] 
**RouteCollectors** | Pointer to [**PeeringProtocolDataRouteCollectors**](PeeringProtocolDataRouteCollectors.md) |  | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewPeeringProtocolData

`func NewPeeringProtocolData() *PeeringProtocolData`

NewPeeringProtocolData instantiates a new PeeringProtocolData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeeringProtocolDataWithDefaults

`func NewPeeringProtocolDataWithDefaults() *PeeringProtocolData`

NewPeeringProtocolDataWithDefaults instantiates a new PeeringProtocolData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PeeringProtocolData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PeeringProtocolData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PeeringProtocolData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PeeringProtocolData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *PeeringProtocolData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PeeringProtocolData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PeeringProtocolData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PeeringProtocolData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *PeeringProtocolData) GetType() PeeringProtocolDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PeeringProtocolData) GetTypeOk() (*PeeringProtocolDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PeeringProtocolData) SetType(v PeeringProtocolDataType)`

SetType sets Type field to given value.

### HasType

`func (o *PeeringProtocolData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *PeeringProtocolData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PeeringProtocolData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PeeringProtocolData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PeeringProtocolData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PeeringProtocolData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PeeringProtocolData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PeeringProtocolData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PeeringProtocolData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomerAsn

`func (o *PeeringProtocolData) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *PeeringProtocolData) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *PeeringProtocolData) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.

### HasCustomerAsn

`func (o *PeeringProtocolData) HasCustomerAsn() bool`

HasCustomerAsn returns a boolean if a field has been set.

### GetEquinixAsn

`func (o *PeeringProtocolData) GetEquinixAsn() int64`

GetEquinixAsn returns the EquinixAsn field if non-nil, zero value otherwise.

### GetEquinixAsnOk

`func (o *PeeringProtocolData) GetEquinixAsnOk() (*int64, bool)`

GetEquinixAsnOk returns a tuple with the EquinixAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixAsn

`func (o *PeeringProtocolData) SetEquinixAsn(v int64)`

SetEquinixAsn sets EquinixAsn field to given value.

### HasEquinixAsn

`func (o *PeeringProtocolData) HasEquinixAsn() bool`

HasEquinixAsn returns a boolean if a field has been set.

### GetState

`func (o *PeeringProtocolData) GetState() PeeringProtocolDataState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PeeringProtocolData) GetStateOk() (*PeeringProtocolDataState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PeeringProtocolData) SetState(v PeeringProtocolDataState)`

SetState sets State field to given value.

### HasState

`func (o *PeeringProtocolData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMacAddress

`func (o *PeeringProtocolData) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *PeeringProtocolData) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *PeeringProtocolData) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *PeeringProtocolData) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetBgpIpv4

`func (o *PeeringProtocolData) GetBgpIpv4() PeeringConnectionResIpv4`

GetBgpIpv4 returns the BgpIpv4 field if non-nil, zero value otherwise.

### GetBgpIpv4Ok

`func (o *PeeringProtocolData) GetBgpIpv4Ok() (*PeeringConnectionResIpv4, bool)`

GetBgpIpv4Ok returns a tuple with the BgpIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv4

`func (o *PeeringProtocolData) SetBgpIpv4(v PeeringConnectionResIpv4)`

SetBgpIpv4 sets BgpIpv4 field to given value.

### HasBgpIpv4

`func (o *PeeringProtocolData) HasBgpIpv4() bool`

HasBgpIpv4 returns a boolean if a field has been set.

### GetBgpIpv6

`func (o *PeeringProtocolData) GetBgpIpv6() PeeringConnectionResIpv6`

GetBgpIpv6 returns the BgpIpv6 field if non-nil, zero value otherwise.

### GetBgpIpv6Ok

`func (o *PeeringProtocolData) GetBgpIpv6Ok() (*PeeringConnectionResIpv6, bool)`

GetBgpIpv6Ok returns a tuple with the BgpIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv6

`func (o *PeeringProtocolData) SetBgpIpv6(v PeeringConnectionResIpv6)`

SetBgpIpv6 sets BgpIpv6 field to given value.

### HasBgpIpv6

`func (o *PeeringProtocolData) HasBgpIpv6() bool`

HasBgpIpv6 returns a boolean if a field has been set.

### GetRouteCollectors

`func (o *PeeringProtocolData) GetRouteCollectors() PeeringProtocolDataRouteCollectors`

GetRouteCollectors returns the RouteCollectors field if non-nil, zero value otherwise.

### GetRouteCollectorsOk

`func (o *PeeringProtocolData) GetRouteCollectorsOk() (*PeeringProtocolDataRouteCollectors, bool)`

GetRouteCollectorsOk returns a tuple with the RouteCollectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteCollectors

`func (o *PeeringProtocolData) SetRouteCollectors(v PeeringProtocolDataRouteCollectors)`

SetRouteCollectors sets RouteCollectors field to given value.

### HasRouteCollectors

`func (o *PeeringProtocolData) HasRouteCollectors() bool`

HasRouteCollectors returns a boolean if a field has been set.

### GetChangelog

`func (o *PeeringProtocolData) GetChangelog() Changelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *PeeringProtocolData) GetChangelogOk() (*Changelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *PeeringProtocolData) SetChangelog(v Changelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *PeeringProtocolData) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


