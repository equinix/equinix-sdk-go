# CloudRouter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Cloud Routers URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**State** | Pointer to [**CloudRouterAccessPointState**](CloudRouterAccessPointState.md) |  | [optional] 
**EquinixAsn** | Pointer to **int64** | Equinix ASN | [optional] 
**ConnectionsCount** | Pointer to **int32** | Number of connections associated with this Access point | [optional] 
**MarketplaceSubscription** | Pointer to [**MarketplaceSubscription**](MarketplaceSubscription.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**Change** | Pointer to [**CloudRouterChange**](CloudRouterChange.md) |  | [optional] 
**Type** | Pointer to [**CloudRouterPostRequestType**](CloudRouterPostRequestType.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided Cloud Router name | [optional] 
**Location** | Pointer to [**SimplifiedLocationWithoutIBX**](SimplifiedLocationWithoutIBX.md) |  | [optional] 
**Package** | Pointer to [**CloudRouterPostRequestPackage**](CloudRouterPostRequestPackage.md) |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on connection configuration or status changes | [optional] 

## Methods

### NewCloudRouter

`func NewCloudRouter() *CloudRouter`

NewCloudRouter instantiates a new CloudRouter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterWithDefaults

`func NewCloudRouterWithDefaults() *CloudRouter`

NewCloudRouterWithDefaults instantiates a new CloudRouter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CloudRouter) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CloudRouter) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CloudRouter) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CloudRouter) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *CloudRouter) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CloudRouter) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CloudRouter) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CloudRouter) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *CloudRouter) GetState() CloudRouterAccessPointState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudRouter) GetStateOk() (*CloudRouterAccessPointState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudRouter) SetState(v CloudRouterAccessPointState)`

SetState sets State field to given value.

### HasState

`func (o *CloudRouter) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEquinixAsn

`func (o *CloudRouter) GetEquinixAsn() int64`

GetEquinixAsn returns the EquinixAsn field if non-nil, zero value otherwise.

### GetEquinixAsnOk

`func (o *CloudRouter) GetEquinixAsnOk() (*int64, bool)`

GetEquinixAsnOk returns a tuple with the EquinixAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixAsn

`func (o *CloudRouter) SetEquinixAsn(v int64)`

SetEquinixAsn sets EquinixAsn field to given value.

### HasEquinixAsn

`func (o *CloudRouter) HasEquinixAsn() bool`

HasEquinixAsn returns a boolean if a field has been set.

### GetConnectionsCount

`func (o *CloudRouter) GetConnectionsCount() int32`

GetConnectionsCount returns the ConnectionsCount field if non-nil, zero value otherwise.

### GetConnectionsCountOk

`func (o *CloudRouter) GetConnectionsCountOk() (*int32, bool)`

GetConnectionsCountOk returns a tuple with the ConnectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsCount

`func (o *CloudRouter) SetConnectionsCount(v int32)`

SetConnectionsCount sets ConnectionsCount field to given value.

### HasConnectionsCount

`func (o *CloudRouter) HasConnectionsCount() bool`

HasConnectionsCount returns a boolean if a field has been set.

### GetMarketplaceSubscription

`func (o *CloudRouter) GetMarketplaceSubscription() MarketplaceSubscription`

GetMarketplaceSubscription returns the MarketplaceSubscription field if non-nil, zero value otherwise.

### GetMarketplaceSubscriptionOk

`func (o *CloudRouter) GetMarketplaceSubscriptionOk() (*MarketplaceSubscription, bool)`

GetMarketplaceSubscriptionOk returns a tuple with the MarketplaceSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceSubscription

`func (o *CloudRouter) SetMarketplaceSubscription(v MarketplaceSubscription)`

SetMarketplaceSubscription sets MarketplaceSubscription field to given value.

### HasMarketplaceSubscription

`func (o *CloudRouter) HasMarketplaceSubscription() bool`

HasMarketplaceSubscription returns a boolean if a field has been set.

### GetChangeLog

`func (o *CloudRouter) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *CloudRouter) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *CloudRouter) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *CloudRouter) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetChange

`func (o *CloudRouter) GetChange() CloudRouterChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *CloudRouter) GetChangeOk() (*CloudRouterChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *CloudRouter) SetChange(v CloudRouterChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *CloudRouter) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetType

`func (o *CloudRouter) GetType() CloudRouterPostRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudRouter) GetTypeOk() (*CloudRouterPostRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudRouter) SetType(v CloudRouterPostRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *CloudRouter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CloudRouter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudRouter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudRouter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudRouter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *CloudRouter) GetLocation() SimplifiedLocationWithoutIBX`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CloudRouter) GetLocationOk() (*SimplifiedLocationWithoutIBX, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CloudRouter) SetLocation(v SimplifiedLocationWithoutIBX)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CloudRouter) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPackage

`func (o *CloudRouter) GetPackage() CloudRouterPostRequestPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *CloudRouter) GetPackageOk() (*CloudRouterPostRequestPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *CloudRouter) SetPackage(v CloudRouterPostRequestPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *CloudRouter) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetOrder

`func (o *CloudRouter) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CloudRouter) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CloudRouter) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CloudRouter) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProject

`func (o *CloudRouter) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudRouter) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudRouter) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudRouter) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAccount

`func (o *CloudRouter) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudRouter) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudRouter) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudRouter) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetNotifications

`func (o *CloudRouter) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *CloudRouter) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *CloudRouter) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *CloudRouter) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


