# ConnectionPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ConnectionType**](ConnectionType.md) |  | 
**Name** | **string** | Customer-provided connection name | 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Notifications** | [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on connection configuration or status changes | 
**Bandwidth** | **int32** | Connection bandwidth in Mbps | 
**GeoScope** | Pointer to [**GeoScopeType**](GeoScopeType.md) |  | [optional] 
**Redundancy** | Pointer to [**ConnectionRedundancy**](ConnectionRedundancy.md) |  | [optional] 
**ASide** | [**ConnectionSide**](ConnectionSide.md) |  | 
**ZSide** | [**ConnectionSide**](ConnectionSide.md) |  | 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**AdditionalInfo** | Pointer to [**[]ConnectionSideAdditionalInfo**](ConnectionSideAdditionalInfo.md) | Connection additional information | [optional] 
**MarketplaceSubscription** | Pointer to [**MarketplaceSubscription**](MarketplaceSubscription.md) |  | [optional] 
**EndCustomer** | Pointer to [**EndCustomer**](EndCustomer.md) |  | [optional] 

## Methods

### NewConnectionPostRequest

`func NewConnectionPostRequest(type_ ConnectionType, name string, notifications []SimplifiedNotification, bandwidth int32, aSide ConnectionSide, zSide ConnectionSide, ) *ConnectionPostRequest`

NewConnectionPostRequest instantiates a new ConnectionPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionPostRequestWithDefaults

`func NewConnectionPostRequestWithDefaults() *ConnectionPostRequest`

NewConnectionPostRequestWithDefaults instantiates a new ConnectionPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectionPostRequest) GetType() ConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionPostRequest) GetTypeOk() (*ConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionPostRequest) SetType(v ConnectionType)`

SetType sets Type field to given value.


### GetName

`func (o *ConnectionPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOrder

`func (o *ConnectionPostRequest) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ConnectionPostRequest) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ConnectionPostRequest) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ConnectionPostRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetNotifications

`func (o *ConnectionPostRequest) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ConnectionPostRequest) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ConnectionPostRequest) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.


### GetBandwidth

`func (o *ConnectionPostRequest) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *ConnectionPostRequest) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *ConnectionPostRequest) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetGeoScope

`func (o *ConnectionPostRequest) GetGeoScope() GeoScopeType`

GetGeoScope returns the GeoScope field if non-nil, zero value otherwise.

### GetGeoScopeOk

`func (o *ConnectionPostRequest) GetGeoScopeOk() (*GeoScopeType, bool)`

GetGeoScopeOk returns a tuple with the GeoScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoScope

`func (o *ConnectionPostRequest) SetGeoScope(v GeoScopeType)`

SetGeoScope sets GeoScope field to given value.

### HasGeoScope

`func (o *ConnectionPostRequest) HasGeoScope() bool`

HasGeoScope returns a boolean if a field has been set.

### GetRedundancy

`func (o *ConnectionPostRequest) GetRedundancy() ConnectionRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *ConnectionPostRequest) GetRedundancyOk() (*ConnectionRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *ConnectionPostRequest) SetRedundancy(v ConnectionRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *ConnectionPostRequest) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetASide

`func (o *ConnectionPostRequest) GetASide() ConnectionSide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *ConnectionPostRequest) GetASideOk() (*ConnectionSide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *ConnectionPostRequest) SetASide(v ConnectionSide)`

SetASide sets ASide field to given value.


### GetZSide

`func (o *ConnectionPostRequest) GetZSide() ConnectionSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *ConnectionPostRequest) GetZSideOk() (*ConnectionSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *ConnectionPostRequest) SetZSide(v ConnectionSide)`

SetZSide sets ZSide field to given value.


### GetProject

`func (o *ConnectionPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ConnectionPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ConnectionPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *ConnectionPostRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *ConnectionPostRequest) GetAdditionalInfo() []ConnectionSideAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ConnectionPostRequest) GetAdditionalInfoOk() (*[]ConnectionSideAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ConnectionPostRequest) SetAdditionalInfo(v []ConnectionSideAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ConnectionPostRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetMarketplaceSubscription

`func (o *ConnectionPostRequest) GetMarketplaceSubscription() MarketplaceSubscription`

GetMarketplaceSubscription returns the MarketplaceSubscription field if non-nil, zero value otherwise.

### GetMarketplaceSubscriptionOk

`func (o *ConnectionPostRequest) GetMarketplaceSubscriptionOk() (*MarketplaceSubscription, bool)`

GetMarketplaceSubscriptionOk returns a tuple with the MarketplaceSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceSubscription

`func (o *ConnectionPostRequest) SetMarketplaceSubscription(v MarketplaceSubscription)`

SetMarketplaceSubscription sets MarketplaceSubscription field to given value.

### HasMarketplaceSubscription

`func (o *ConnectionPostRequest) HasMarketplaceSubscription() bool`

HasMarketplaceSubscription returns a boolean if a field has been set.

### GetEndCustomer

`func (o *ConnectionPostRequest) GetEndCustomer() EndCustomer`

GetEndCustomer returns the EndCustomer field if non-nil, zero value otherwise.

### GetEndCustomerOk

`func (o *ConnectionPostRequest) GetEndCustomerOk() (*EndCustomer, bool)`

GetEndCustomerOk returns a tuple with the EndCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomer

`func (o *ConnectionPostRequest) SetEndCustomer(v EndCustomer)`

SetEndCustomer sets EndCustomer field to given value.

### HasEndCustomer

`func (o *ConnectionPostRequest) HasEndCustomer() bool`

HasEndCustomer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


