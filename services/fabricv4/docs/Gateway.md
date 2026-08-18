# Gateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Gateways URI | [optional] [readonly] 
**Type** | Pointer to [**GatewayType**](GatewayType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Name** | Pointer to **string** | Customer-provided Gateway name | [optional] 
**Description** | Pointer to **string** | Customer-provided Gateway description | [optional] 
**State** | Pointer to [**GatewayState**](GatewayState.md) |  | [optional] 
**Bandwidth** | Pointer to **int32** | Gateway bandwidth in Mbps | [optional] 
**LocalAsn** | Pointer to **int32** | Gateway local Autonomous System Number | [optional] 
**Router** | Pointer to [**Router**](Router.md) |  | [optional] 
**Ipv4** | Pointer to [**GatewayIpv4**](GatewayIpv4.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewGateway

`func NewGateway() *Gateway`

NewGateway instantiates a new Gateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayWithDefaults

`func NewGatewayWithDefaults() *Gateway`

NewGatewayWithDefaults instantiates a new Gateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Gateway) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Gateway) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Gateway) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Gateway) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *Gateway) GetType() GatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Gateway) GetTypeOk() (*GatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Gateway) SetType(v GatewayType)`

SetType sets Type field to given value.

### HasType

`func (o *Gateway) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *Gateway) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Gateway) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Gateway) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Gateway) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *Gateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Gateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Gateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Gateway) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Gateway) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Gateway) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Gateway) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Gateway) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *Gateway) GetState() GatewayState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Gateway) GetStateOk() (*GatewayState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Gateway) SetState(v GatewayState)`

SetState sets State field to given value.

### HasState

`func (o *Gateway) HasState() bool`

HasState returns a boolean if a field has been set.

### GetBandwidth

`func (o *Gateway) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *Gateway) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *Gateway) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *Gateway) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetLocalAsn

`func (o *Gateway) GetLocalAsn() int32`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *Gateway) GetLocalAsnOk() (*int32, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *Gateway) SetLocalAsn(v int32)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *Gateway) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetRouter

`func (o *Gateway) GetRouter() Router`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *Gateway) GetRouterOk() (*Router, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *Gateway) SetRouter(v Router)`

SetRouter sets Router field to given value.

### HasRouter

`func (o *Gateway) HasRouter() bool`

HasRouter returns a boolean if a field has been set.

### GetIpv4

`func (o *Gateway) GetIpv4() GatewayIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *Gateway) GetIpv4Ok() (*GatewayIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *Gateway) SetIpv4(v GatewayIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *Gateway) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetProject

`func (o *Gateway) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Gateway) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Gateway) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Gateway) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAccount

`func (o *Gateway) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Gateway) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Gateway) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Gateway) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOrder

`func (o *Gateway) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Gateway) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Gateway) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Gateway) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetChangeLog

`func (o *Gateway) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *Gateway) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *Gateway) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *Gateway) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


