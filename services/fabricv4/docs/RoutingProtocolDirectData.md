# RoutingProtocolDirectData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**RoutingProtocolDirectTypeType**](RoutingProtocolDirectTypeType.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DirectIpv4** | Pointer to [**DirectConnectionIpv4**](DirectConnectionIpv4.md) |  | [optional] 
**DirectIpv6** | Pointer to [**DirectConnectionIpv6**](DirectConnectionIpv6.md) |  | [optional] 
**Href** | Pointer to **string** | Routing Protocol URI | [optional] 
**Uuid** | Pointer to **string** | Routing protocol identifier | [optional] 
**State** | Pointer to [**RoutingProtocolBGPDataState**](RoutingProtocolBGPDataState.md) |  | [optional] 
**Operation** | Pointer to [**RoutingProtocolOperation**](RoutingProtocolOperation.md) |  | [optional] 
**Change** | Pointer to [**RoutingProtocolChange**](RoutingProtocolChange.md) |  | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewRoutingProtocolDirectData

`func NewRoutingProtocolDirectData() *RoutingProtocolDirectData`

NewRoutingProtocolDirectData instantiates a new RoutingProtocolDirectData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolDirectDataWithDefaults

`func NewRoutingProtocolDirectDataWithDefaults() *RoutingProtocolDirectData`

NewRoutingProtocolDirectDataWithDefaults instantiates a new RoutingProtocolDirectData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutingProtocolDirectData) GetType() RoutingProtocolDirectTypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolDirectData) GetTypeOk() (*RoutingProtocolDirectTypeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolDirectData) SetType(v RoutingProtocolDirectTypeType)`

SetType sets Type field to given value.

### HasType

`func (o *RoutingProtocolDirectData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *RoutingProtocolDirectData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolDirectData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolDirectData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoutingProtocolDirectData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDirectIpv4

`func (o *RoutingProtocolDirectData) GetDirectIpv4() DirectConnectionIpv4`

GetDirectIpv4 returns the DirectIpv4 field if non-nil, zero value otherwise.

### GetDirectIpv4Ok

`func (o *RoutingProtocolDirectData) GetDirectIpv4Ok() (*DirectConnectionIpv4, bool)`

GetDirectIpv4Ok returns a tuple with the DirectIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectIpv4

`func (o *RoutingProtocolDirectData) SetDirectIpv4(v DirectConnectionIpv4)`

SetDirectIpv4 sets DirectIpv4 field to given value.

### HasDirectIpv4

`func (o *RoutingProtocolDirectData) HasDirectIpv4() bool`

HasDirectIpv4 returns a boolean if a field has been set.

### GetDirectIpv6

`func (o *RoutingProtocolDirectData) GetDirectIpv6() DirectConnectionIpv6`

GetDirectIpv6 returns the DirectIpv6 field if non-nil, zero value otherwise.

### GetDirectIpv6Ok

`func (o *RoutingProtocolDirectData) GetDirectIpv6Ok() (*DirectConnectionIpv6, bool)`

GetDirectIpv6Ok returns a tuple with the DirectIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectIpv6

`func (o *RoutingProtocolDirectData) SetDirectIpv6(v DirectConnectionIpv6)`

SetDirectIpv6 sets DirectIpv6 field to given value.

### HasDirectIpv6

`func (o *RoutingProtocolDirectData) HasDirectIpv6() bool`

HasDirectIpv6 returns a boolean if a field has been set.

### GetHref

`func (o *RoutingProtocolDirectData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RoutingProtocolDirectData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RoutingProtocolDirectData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RoutingProtocolDirectData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *RoutingProtocolDirectData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoutingProtocolDirectData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoutingProtocolDirectData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RoutingProtocolDirectData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *RoutingProtocolDirectData) GetState() RoutingProtocolBGPDataState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RoutingProtocolDirectData) GetStateOk() (*RoutingProtocolBGPDataState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RoutingProtocolDirectData) SetState(v RoutingProtocolBGPDataState)`

SetState sets State field to given value.

### HasState

`func (o *RoutingProtocolDirectData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOperation

`func (o *RoutingProtocolDirectData) GetOperation() RoutingProtocolOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *RoutingProtocolDirectData) GetOperationOk() (*RoutingProtocolOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *RoutingProtocolDirectData) SetOperation(v RoutingProtocolOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *RoutingProtocolDirectData) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetChange

`func (o *RoutingProtocolDirectData) GetChange() RoutingProtocolChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *RoutingProtocolDirectData) GetChangeOk() (*RoutingProtocolChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *RoutingProtocolDirectData) SetChange(v RoutingProtocolChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *RoutingProtocolDirectData) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetChangelog

`func (o *RoutingProtocolDirectData) GetChangelog() Changelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *RoutingProtocolDirectData) GetChangelogOk() (*Changelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *RoutingProtocolDirectData) SetChangelog(v Changelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *RoutingProtocolDirectData) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


