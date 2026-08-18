# ExchangeServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Internet Exchange Service URI | [optional] 
**Uuid** | Pointer to **string** | Internet Exchange Service identifier | [optional] 
**Type** | Pointer to [**ExchangeServiceResponseType**](ExchangeServiceResponseType.md) |  | [optional] [default to EXCHANGESERVICERESPONSETYPE_IX]
**Name** | Pointer to **string** | Name | [optional] 
**Bandwidth** | Pointer to **int32** | bandwidth in Mbps | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**State** | Pointer to [**ExchangeServiceResponseState**](ExchangeServiceResponseState.md) |  | [optional] 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 
**PublicPeeringConnection** | Pointer to [**PublicPeeringConnectionResponse**](PublicPeeringConnectionResponse.md) |  | [optional] 
**RoutingProtocol** | Pointer to [**RoutingProtocolResponse**](RoutingProtocolResponse.md) |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Notifications** | Pointer to [**[]ExchangeServiceNotification**](ExchangeServiceNotification.md) | Preferences for notifications on Internet Exchange Service configuration or status changes | [optional] 
**Changelog** | Pointer to [**PlatformChangelog**](PlatformChangelog.md) |  | [optional] 

## Methods

### NewExchangeServiceResponse

`func NewExchangeServiceResponse() *ExchangeServiceResponse`

NewExchangeServiceResponse instantiates a new ExchangeServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeServiceResponseWithDefaults

`func NewExchangeServiceResponseWithDefaults() *ExchangeServiceResponse`

NewExchangeServiceResponseWithDefaults instantiates a new ExchangeServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ExchangeServiceResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ExchangeServiceResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ExchangeServiceResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ExchangeServiceResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *ExchangeServiceResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ExchangeServiceResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ExchangeServiceResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ExchangeServiceResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *ExchangeServiceResponse) GetType() ExchangeServiceResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExchangeServiceResponse) GetTypeOk() (*ExchangeServiceResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExchangeServiceResponse) SetType(v ExchangeServiceResponseType)`

SetType sets Type field to given value.

### HasType

`func (o *ExchangeServiceResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ExchangeServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangeServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangeServiceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExchangeServiceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBandwidth

`func (o *ExchangeServiceResponse) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *ExchangeServiceResponse) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *ExchangeServiceResponse) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *ExchangeServiceResponse) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetDescription

`func (o *ExchangeServiceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExchangeServiceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExchangeServiceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExchangeServiceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *ExchangeServiceResponse) GetState() ExchangeServiceResponseState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ExchangeServiceResponse) GetStateOk() (*ExchangeServiceResponseState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ExchangeServiceResponse) SetState(v ExchangeServiceResponseState)`

SetState sets State field to given value.

### HasState

`func (o *ExchangeServiceResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLocation

`func (o *ExchangeServiceResponse) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ExchangeServiceResponse) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ExchangeServiceResponse) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ExchangeServiceResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPublicPeeringConnection

`func (o *ExchangeServiceResponse) GetPublicPeeringConnection() PublicPeeringConnectionResponse`

GetPublicPeeringConnection returns the PublicPeeringConnection field if non-nil, zero value otherwise.

### GetPublicPeeringConnectionOk

`func (o *ExchangeServiceResponse) GetPublicPeeringConnectionOk() (*PublicPeeringConnectionResponse, bool)`

GetPublicPeeringConnectionOk returns a tuple with the PublicPeeringConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPeeringConnection

`func (o *ExchangeServiceResponse) SetPublicPeeringConnection(v PublicPeeringConnectionResponse)`

SetPublicPeeringConnection sets PublicPeeringConnection field to given value.

### HasPublicPeeringConnection

`func (o *ExchangeServiceResponse) HasPublicPeeringConnection() bool`

HasPublicPeeringConnection returns a boolean if a field has been set.

### GetRoutingProtocol

`func (o *ExchangeServiceResponse) GetRoutingProtocol() RoutingProtocolResponse`

GetRoutingProtocol returns the RoutingProtocol field if non-nil, zero value otherwise.

### GetRoutingProtocolOk

`func (o *ExchangeServiceResponse) GetRoutingProtocolOk() (*RoutingProtocolResponse, bool)`

GetRoutingProtocolOk returns a tuple with the RoutingProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingProtocol

`func (o *ExchangeServiceResponse) SetRoutingProtocol(v RoutingProtocolResponse)`

SetRoutingProtocol sets RoutingProtocol field to given value.

### HasRoutingProtocol

`func (o *ExchangeServiceResponse) HasRoutingProtocol() bool`

HasRoutingProtocol returns a boolean if a field has been set.

### GetOrder

`func (o *ExchangeServiceResponse) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ExchangeServiceResponse) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ExchangeServiceResponse) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ExchangeServiceResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProject

`func (o *ExchangeServiceResponse) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ExchangeServiceResponse) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ExchangeServiceResponse) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *ExchangeServiceResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAccount

`func (o *ExchangeServiceResponse) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ExchangeServiceResponse) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ExchangeServiceResponse) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ExchangeServiceResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetNotifications

`func (o *ExchangeServiceResponse) GetNotifications() []ExchangeServiceNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ExchangeServiceResponse) GetNotificationsOk() (*[]ExchangeServiceNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ExchangeServiceResponse) SetNotifications(v []ExchangeServiceNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ExchangeServiceResponse) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetChangelog

`func (o *ExchangeServiceResponse) GetChangelog() PlatformChangelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *ExchangeServiceResponse) GetChangelogOk() (*PlatformChangelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *ExchangeServiceResponse) SetChangelog(v PlatformChangelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *ExchangeServiceResponse) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


