# CloudRouterReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CloudRouterReadResponseType**](CloudRouterReadResponseType.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided Cloud Router name | [optional] 
**Location** | Pointer to [**SimplifiedLocationWithoutIBX**](SimplifiedLocationWithoutIBX.md) |  | [optional] 
**Package** | Pointer to [**CloudRouterPostRequestPackage**](CloudRouterPostRequestPackage.md) |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on connection configuration or status changes | [optional] 
**Href** | Pointer to **string** | Cloud Routers URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**State** | Pointer to [**CloudRouterAccessPointState**](CloudRouterAccessPointState.md) |  | [optional] 
**EquinixAsn** | Pointer to **int64** | Equinix ASN | [optional] 
**ConnectionsCount** | Pointer to **int32** | Number of connections associated with this Access point | [optional] 
**MarketplaceSubscription** | Pointer to [**MarketplaceSubscription**](MarketplaceSubscription.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**Change** | Pointer to [**CloudRouterChange**](CloudRouterChange.md) |  | [optional] 

## Methods

### NewCloudRouterReadResponse

`func NewCloudRouterReadResponse() *CloudRouterReadResponse`

NewCloudRouterReadResponse instantiates a new CloudRouterReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterReadResponseWithDefaults

`func NewCloudRouterReadResponseWithDefaults() *CloudRouterReadResponse`

NewCloudRouterReadResponseWithDefaults instantiates a new CloudRouterReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CloudRouterReadResponse) GetType() CloudRouterReadResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudRouterReadResponse) GetTypeOk() (*CloudRouterReadResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudRouterReadResponse) SetType(v CloudRouterReadResponseType)`

SetType sets Type field to given value.

### HasType

`func (o *CloudRouterReadResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CloudRouterReadResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudRouterReadResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudRouterReadResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudRouterReadResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *CloudRouterReadResponse) GetLocation() SimplifiedLocationWithoutIBX`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CloudRouterReadResponse) GetLocationOk() (*SimplifiedLocationWithoutIBX, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CloudRouterReadResponse) SetLocation(v SimplifiedLocationWithoutIBX)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CloudRouterReadResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPackage

`func (o *CloudRouterReadResponse) GetPackage() CloudRouterPostRequestPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *CloudRouterReadResponse) GetPackageOk() (*CloudRouterPostRequestPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *CloudRouterReadResponse) SetPackage(v CloudRouterPostRequestPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *CloudRouterReadResponse) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetOrder

`func (o *CloudRouterReadResponse) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CloudRouterReadResponse) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CloudRouterReadResponse) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CloudRouterReadResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProject

`func (o *CloudRouterReadResponse) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudRouterReadResponse) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudRouterReadResponse) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudRouterReadResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAccount

`func (o *CloudRouterReadResponse) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudRouterReadResponse) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudRouterReadResponse) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudRouterReadResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetNotifications

`func (o *CloudRouterReadResponse) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *CloudRouterReadResponse) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *CloudRouterReadResponse) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *CloudRouterReadResponse) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetHref

`func (o *CloudRouterReadResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CloudRouterReadResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CloudRouterReadResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CloudRouterReadResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *CloudRouterReadResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CloudRouterReadResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CloudRouterReadResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CloudRouterReadResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *CloudRouterReadResponse) GetState() CloudRouterAccessPointState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudRouterReadResponse) GetStateOk() (*CloudRouterAccessPointState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudRouterReadResponse) SetState(v CloudRouterAccessPointState)`

SetState sets State field to given value.

### HasState

`func (o *CloudRouterReadResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEquinixAsn

`func (o *CloudRouterReadResponse) GetEquinixAsn() int64`

GetEquinixAsn returns the EquinixAsn field if non-nil, zero value otherwise.

### GetEquinixAsnOk

`func (o *CloudRouterReadResponse) GetEquinixAsnOk() (*int64, bool)`

GetEquinixAsnOk returns a tuple with the EquinixAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixAsn

`func (o *CloudRouterReadResponse) SetEquinixAsn(v int64)`

SetEquinixAsn sets EquinixAsn field to given value.

### HasEquinixAsn

`func (o *CloudRouterReadResponse) HasEquinixAsn() bool`

HasEquinixAsn returns a boolean if a field has been set.

### GetConnectionsCount

`func (o *CloudRouterReadResponse) GetConnectionsCount() int32`

GetConnectionsCount returns the ConnectionsCount field if non-nil, zero value otherwise.

### GetConnectionsCountOk

`func (o *CloudRouterReadResponse) GetConnectionsCountOk() (*int32, bool)`

GetConnectionsCountOk returns a tuple with the ConnectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsCount

`func (o *CloudRouterReadResponse) SetConnectionsCount(v int32)`

SetConnectionsCount sets ConnectionsCount field to given value.

### HasConnectionsCount

`func (o *CloudRouterReadResponse) HasConnectionsCount() bool`

HasConnectionsCount returns a boolean if a field has been set.

### GetMarketplaceSubscription

`func (o *CloudRouterReadResponse) GetMarketplaceSubscription() MarketplaceSubscription`

GetMarketplaceSubscription returns the MarketplaceSubscription field if non-nil, zero value otherwise.

### GetMarketplaceSubscriptionOk

`func (o *CloudRouterReadResponse) GetMarketplaceSubscriptionOk() (*MarketplaceSubscription, bool)`

GetMarketplaceSubscriptionOk returns a tuple with the MarketplaceSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceSubscription

`func (o *CloudRouterReadResponse) SetMarketplaceSubscription(v MarketplaceSubscription)`

SetMarketplaceSubscription sets MarketplaceSubscription field to given value.

### HasMarketplaceSubscription

`func (o *CloudRouterReadResponse) HasMarketplaceSubscription() bool`

HasMarketplaceSubscription returns a boolean if a field has been set.

### GetChangeLog

`func (o *CloudRouterReadResponse) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *CloudRouterReadResponse) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *CloudRouterReadResponse) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *CloudRouterReadResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetChange

`func (o *CloudRouterReadResponse) GetChange() CloudRouterChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *CloudRouterReadResponse) GetChangeOk() (*CloudRouterChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *CloudRouterReadResponse) SetChange(v CloudRouterChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *CloudRouterReadResponse) HasChange() bool`

HasChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


