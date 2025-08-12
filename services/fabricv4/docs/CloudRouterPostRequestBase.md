# CloudRouterPostRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CloudRouterPostRequestBaseType**](CloudRouterPostRequestBaseType.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided Cloud Router name | [optional] 
**Location** | Pointer to [**SimplifiedLocationWithoutIBX**](SimplifiedLocationWithoutIBX.md) |  | [optional] 
**Package** | Pointer to [**CloudRouterPostRequestPackage**](CloudRouterPostRequestPackage.md) |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on connection configuration or status changes | [optional] 
**MarketplaceSubscription** | Pointer to [**MarketplaceSubscription**](MarketplaceSubscription.md) |  | [optional] 

## Methods

### NewCloudRouterPostRequestBase

`func NewCloudRouterPostRequestBase() *CloudRouterPostRequestBase`

NewCloudRouterPostRequestBase instantiates a new CloudRouterPostRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterPostRequestBaseWithDefaults

`func NewCloudRouterPostRequestBaseWithDefaults() *CloudRouterPostRequestBase`

NewCloudRouterPostRequestBaseWithDefaults instantiates a new CloudRouterPostRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CloudRouterPostRequestBase) GetType() CloudRouterPostRequestBaseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudRouterPostRequestBase) GetTypeOk() (*CloudRouterPostRequestBaseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudRouterPostRequestBase) SetType(v CloudRouterPostRequestBaseType)`

SetType sets Type field to given value.

### HasType

`func (o *CloudRouterPostRequestBase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CloudRouterPostRequestBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudRouterPostRequestBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudRouterPostRequestBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudRouterPostRequestBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *CloudRouterPostRequestBase) GetLocation() SimplifiedLocationWithoutIBX`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CloudRouterPostRequestBase) GetLocationOk() (*SimplifiedLocationWithoutIBX, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CloudRouterPostRequestBase) SetLocation(v SimplifiedLocationWithoutIBX)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CloudRouterPostRequestBase) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPackage

`func (o *CloudRouterPostRequestBase) GetPackage() CloudRouterPostRequestPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *CloudRouterPostRequestBase) GetPackageOk() (*CloudRouterPostRequestPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *CloudRouterPostRequestBase) SetPackage(v CloudRouterPostRequestPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *CloudRouterPostRequestBase) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetOrder

`func (o *CloudRouterPostRequestBase) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CloudRouterPostRequestBase) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CloudRouterPostRequestBase) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CloudRouterPostRequestBase) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProject

`func (o *CloudRouterPostRequestBase) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudRouterPostRequestBase) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudRouterPostRequestBase) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudRouterPostRequestBase) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAccount

`func (o *CloudRouterPostRequestBase) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudRouterPostRequestBase) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudRouterPostRequestBase) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudRouterPostRequestBase) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetNotifications

`func (o *CloudRouterPostRequestBase) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *CloudRouterPostRequestBase) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *CloudRouterPostRequestBase) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *CloudRouterPostRequestBase) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetMarketplaceSubscription

`func (o *CloudRouterPostRequestBase) GetMarketplaceSubscription() MarketplaceSubscription`

GetMarketplaceSubscription returns the MarketplaceSubscription field if non-nil, zero value otherwise.

### GetMarketplaceSubscriptionOk

`func (o *CloudRouterPostRequestBase) GetMarketplaceSubscriptionOk() (*MarketplaceSubscription, bool)`

GetMarketplaceSubscriptionOk returns a tuple with the MarketplaceSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceSubscription

`func (o *CloudRouterPostRequestBase) SetMarketplaceSubscription(v MarketplaceSubscription)`

SetMarketplaceSubscription sets MarketplaceSubscription field to given value.

### HasMarketplaceSubscription

`func (o *CloudRouterPostRequestBase) HasMarketplaceSubscription() bool`

HasMarketplaceSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


