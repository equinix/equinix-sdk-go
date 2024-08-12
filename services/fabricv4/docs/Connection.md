# Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ConnectionType**](ConnectionType.md) |  | 
**Href** | Pointer to **string** | Connection URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned connection identifier | [optional] 
**Name** | **string** | Customer-provided connection name | 
**Description** | Pointer to **string** | Customer-provided connection description | [optional] 
**State** | Pointer to [**ConnectionState**](ConnectionState.md) |  | [optional] 
**Change** | Pointer to [**Change**](Change.md) |  | [optional] 
**Operation** | Pointer to [**ConnectionOperation**](ConnectionOperation.md) |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on connection configuration or status changes | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**Bandwidth** | **int32** | Connection bandwidth in Mbps | 
**GeoScope** | Pointer to [**GeoScopeType**](GeoScopeType.md) |  | [optional] 
**Redundancy** | Pointer to [**ConnectionRedundancy**](ConnectionRedundancy.md) |  | [optional] 
**IsRemote** | Pointer to **bool** | Connection property derived from access point locations | [optional] 
**Direction** | Pointer to [**ConnectionDirection**](ConnectionDirection.md) |  | [optional] 
**ASide** | [**ConnectionSide**](ConnectionSide.md) |  | 
**ZSide** | [**ConnectionSide**](ConnectionSide.md) |  | 
**MarketplaceSubscription** | Pointer to [**MarketplaceSubscription**](MarketplaceSubscription.md) |  | [optional] 
**AdditionalInfo** | Pointer to [**[]ConnectionSideAdditionalInfo**](ConnectionSideAdditionalInfo.md) | Connection additional information | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 

## Methods

### NewConnection

`func NewConnection(type_ ConnectionType, name string, bandwidth int32, aSide ConnectionSide, zSide ConnectionSide, ) *Connection`

NewConnection instantiates a new Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionWithDefaults

`func NewConnectionWithDefaults() *Connection`

NewConnectionWithDefaults instantiates a new Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Connection) GetType() ConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Connection) GetTypeOk() (*ConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Connection) SetType(v ConnectionType)`

SetType sets Type field to given value.


### GetHref

`func (o *Connection) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Connection) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Connection) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Connection) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *Connection) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Connection) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Connection) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Connection) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *Connection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Connection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Connection) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Connection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Connection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Connection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Connection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *Connection) GetState() ConnectionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Connection) GetStateOk() (*ConnectionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Connection) SetState(v ConnectionState)`

SetState sets State field to given value.

### HasState

`func (o *Connection) HasState() bool`

HasState returns a boolean if a field has been set.

### GetChange

`func (o *Connection) GetChange() Change`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *Connection) GetChangeOk() (*Change, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *Connection) SetChange(v Change)`

SetChange sets Change field to given value.

### HasChange

`func (o *Connection) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetOperation

`func (o *Connection) GetOperation() ConnectionOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Connection) GetOperationOk() (*ConnectionOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Connection) SetOperation(v ConnectionOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *Connection) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetOrder

`func (o *Connection) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Connection) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Connection) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Connection) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetNotifications

`func (o *Connection) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *Connection) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *Connection) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *Connection) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetAccount

`func (o *Connection) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Connection) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Connection) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Connection) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetChangeLog

`func (o *Connection) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *Connection) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *Connection) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *Connection) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetBandwidth

`func (o *Connection) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *Connection) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *Connection) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetGeoScope

`func (o *Connection) GetGeoScope() GeoScopeType`

GetGeoScope returns the GeoScope field if non-nil, zero value otherwise.

### GetGeoScopeOk

`func (o *Connection) GetGeoScopeOk() (*GeoScopeType, bool)`

GetGeoScopeOk returns a tuple with the GeoScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoScope

`func (o *Connection) SetGeoScope(v GeoScopeType)`

SetGeoScope sets GeoScope field to given value.

### HasGeoScope

`func (o *Connection) HasGeoScope() bool`

HasGeoScope returns a boolean if a field has been set.

### GetRedundancy

`func (o *Connection) GetRedundancy() ConnectionRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *Connection) GetRedundancyOk() (*ConnectionRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *Connection) SetRedundancy(v ConnectionRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *Connection) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetIsRemote

`func (o *Connection) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *Connection) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *Connection) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *Connection) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetDirection

`func (o *Connection) GetDirection() ConnectionDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Connection) GetDirectionOk() (*ConnectionDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Connection) SetDirection(v ConnectionDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Connection) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetASide

`func (o *Connection) GetASide() ConnectionSide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *Connection) GetASideOk() (*ConnectionSide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *Connection) SetASide(v ConnectionSide)`

SetASide sets ASide field to given value.


### GetZSide

`func (o *Connection) GetZSide() ConnectionSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *Connection) GetZSideOk() (*ConnectionSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *Connection) SetZSide(v ConnectionSide)`

SetZSide sets ZSide field to given value.


### GetMarketplaceSubscription

`func (o *Connection) GetMarketplaceSubscription() MarketplaceSubscription`

GetMarketplaceSubscription returns the MarketplaceSubscription field if non-nil, zero value otherwise.

### GetMarketplaceSubscriptionOk

`func (o *Connection) GetMarketplaceSubscriptionOk() (*MarketplaceSubscription, bool)`

GetMarketplaceSubscriptionOk returns a tuple with the MarketplaceSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceSubscription

`func (o *Connection) SetMarketplaceSubscription(v MarketplaceSubscription)`

SetMarketplaceSubscription sets MarketplaceSubscription field to given value.

### HasMarketplaceSubscription

`func (o *Connection) HasMarketplaceSubscription() bool`

HasMarketplaceSubscription returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *Connection) GetAdditionalInfo() []ConnectionSideAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *Connection) GetAdditionalInfoOk() (*[]ConnectionSideAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *Connection) SetAdditionalInfo(v []ConnectionSideAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *Connection) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetProject

`func (o *Connection) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Connection) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Connection) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Connection) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


