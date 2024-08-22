# BaseConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ConnectionType**](ConnectionType.md) |  | [optional] 
**Href** | Pointer to **string** | Connection URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned connection identifier | [optional] 
**Name** | Pointer to **string** | Customer-provided connection name | [optional] 
**Description** | Pointer to **string** | Customer-provided connection description | [optional] 
**State** | Pointer to [**ConnectionState**](ConnectionState.md) |  | [optional] 
**Change** | Pointer to [**Change**](Change.md) |  | [optional] 
**Operation** | Pointer to [**ConnectionOperation**](ConnectionOperation.md) |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on connection configuration or status changes | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**Bandwidth** | Pointer to **int32** | Connection bandwidth in Mbps | [optional] 
**GeoScope** | Pointer to [**GeoScopeType**](GeoScopeType.md) |  | [optional] 
**Redundancy** | Pointer to [**ConnectionRedundancy**](ConnectionRedundancy.md) |  | [optional] 
**IsRemote** | Pointer to **bool** | Connection property derived from access point locations | [optional] 
**Direction** | Pointer to [**ConnectionDirection**](ConnectionDirection.md) |  | [optional] 
**ASide** | Pointer to [**ConnectionSide**](ConnectionSide.md) |  | [optional] 
**ZSide** | Pointer to [**ConnectionSide**](ConnectionSide.md) |  | [optional] 
**MarketplaceSubscription** | Pointer to [**MarketplaceSubscription**](MarketplaceSubscription.md) |  | [optional] 
**AdditionalInfo** | Pointer to [**[]ConnectionSideAdditionalInfo**](ConnectionSideAdditionalInfo.md) | Connection additional information | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 

## Methods

### NewBaseConnection

`func NewBaseConnection() *BaseConnection`

NewBaseConnection instantiates a new BaseConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseConnectionWithDefaults

`func NewBaseConnectionWithDefaults() *BaseConnection`

NewBaseConnectionWithDefaults instantiates a new BaseConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BaseConnection) GetType() ConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseConnection) GetTypeOk() (*ConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseConnection) SetType(v ConnectionType)`

SetType sets Type field to given value.

### HasType

`func (o *BaseConnection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *BaseConnection) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BaseConnection) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BaseConnection) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BaseConnection) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *BaseConnection) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *BaseConnection) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *BaseConnection) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *BaseConnection) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *BaseConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BaseConnection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseConnection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseConnection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseConnection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *BaseConnection) GetState() ConnectionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BaseConnection) GetStateOk() (*ConnectionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BaseConnection) SetState(v ConnectionState)`

SetState sets State field to given value.

### HasState

`func (o *BaseConnection) HasState() bool`

HasState returns a boolean if a field has been set.

### GetChange

`func (o *BaseConnection) GetChange() Change`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *BaseConnection) GetChangeOk() (*Change, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *BaseConnection) SetChange(v Change)`

SetChange sets Change field to given value.

### HasChange

`func (o *BaseConnection) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetOperation

`func (o *BaseConnection) GetOperation() ConnectionOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BaseConnection) GetOperationOk() (*ConnectionOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BaseConnection) SetOperation(v ConnectionOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *BaseConnection) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetOrder

`func (o *BaseConnection) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *BaseConnection) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *BaseConnection) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *BaseConnection) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetNotifications

`func (o *BaseConnection) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *BaseConnection) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *BaseConnection) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *BaseConnection) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetAccount

`func (o *BaseConnection) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BaseConnection) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BaseConnection) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *BaseConnection) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetChangeLog

`func (o *BaseConnection) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *BaseConnection) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *BaseConnection) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *BaseConnection) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetBandwidth

`func (o *BaseConnection) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *BaseConnection) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *BaseConnection) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *BaseConnection) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetGeoScope

`func (o *BaseConnection) GetGeoScope() GeoScopeType`

GetGeoScope returns the GeoScope field if non-nil, zero value otherwise.

### GetGeoScopeOk

`func (o *BaseConnection) GetGeoScopeOk() (*GeoScopeType, bool)`

GetGeoScopeOk returns a tuple with the GeoScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoScope

`func (o *BaseConnection) SetGeoScope(v GeoScopeType)`

SetGeoScope sets GeoScope field to given value.

### HasGeoScope

`func (o *BaseConnection) HasGeoScope() bool`

HasGeoScope returns a boolean if a field has been set.

### GetRedundancy

`func (o *BaseConnection) GetRedundancy() ConnectionRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *BaseConnection) GetRedundancyOk() (*ConnectionRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *BaseConnection) SetRedundancy(v ConnectionRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *BaseConnection) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetIsRemote

`func (o *BaseConnection) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *BaseConnection) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *BaseConnection) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *BaseConnection) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetDirection

`func (o *BaseConnection) GetDirection() ConnectionDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BaseConnection) GetDirectionOk() (*ConnectionDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BaseConnection) SetDirection(v ConnectionDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *BaseConnection) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetASide

`func (o *BaseConnection) GetASide() ConnectionSide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *BaseConnection) GetASideOk() (*ConnectionSide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *BaseConnection) SetASide(v ConnectionSide)`

SetASide sets ASide field to given value.

### HasASide

`func (o *BaseConnection) HasASide() bool`

HasASide returns a boolean if a field has been set.

### GetZSide

`func (o *BaseConnection) GetZSide() ConnectionSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *BaseConnection) GetZSideOk() (*ConnectionSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *BaseConnection) SetZSide(v ConnectionSide)`

SetZSide sets ZSide field to given value.

### HasZSide

`func (o *BaseConnection) HasZSide() bool`

HasZSide returns a boolean if a field has been set.

### GetMarketplaceSubscription

`func (o *BaseConnection) GetMarketplaceSubscription() MarketplaceSubscription`

GetMarketplaceSubscription returns the MarketplaceSubscription field if non-nil, zero value otherwise.

### GetMarketplaceSubscriptionOk

`func (o *BaseConnection) GetMarketplaceSubscriptionOk() (*MarketplaceSubscription, bool)`

GetMarketplaceSubscriptionOk returns a tuple with the MarketplaceSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceSubscription

`func (o *BaseConnection) SetMarketplaceSubscription(v MarketplaceSubscription)`

SetMarketplaceSubscription sets MarketplaceSubscription field to given value.

### HasMarketplaceSubscription

`func (o *BaseConnection) HasMarketplaceSubscription() bool`

HasMarketplaceSubscription returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *BaseConnection) GetAdditionalInfo() []ConnectionSideAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *BaseConnection) GetAdditionalInfoOk() (*[]ConnectionSideAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *BaseConnection) SetAdditionalInfo(v []ConnectionSideAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *BaseConnection) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetProject

`func (o *BaseConnection) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BaseConnection) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BaseConnection) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *BaseConnection) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


