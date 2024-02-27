# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**NetworkType**](NetworkType.md) |  | 
**Name** | **string** | Customer-provided network name | 
**Scope** | [**NetworkScope**](NetworkScope.md) |  | 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Notifications** | [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on network configuration or status changes | 
**Href** | **string** | Network URI | [readonly] 
**Uuid** | **string** | Equinix-assigned network identifier | 
**State** | [**NetworkState**](NetworkState.md) |  | 
**ConnectionsCount** | Pointer to **float32** | number of connections created on the network | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Change** | Pointer to [**SimplifiedNetworkChange**](SimplifiedNetworkChange.md) |  | [optional] 
**Operation** | Pointer to [**NetworkOperation**](NetworkOperation.md) |  | [optional] 
**ChangeLog** | [**Changelog**](Changelog.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) | Network sub-resources links | [optional] [readonly] 

## Methods

### NewNetwork

`func NewNetwork(type_ NetworkType, name string, scope NetworkScope, notifications []SimplifiedNotification, href string, uuid string, state NetworkState, changeLog Changelog, ) *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Network) GetType() NetworkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Network) GetTypeOk() (*NetworkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Network) SetType(v NetworkType)`

SetType sets Type field to given value.


### GetName

`func (o *Network) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Network) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Network) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *Network) GetScope() NetworkScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Network) GetScopeOk() (*NetworkScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Network) SetScope(v NetworkScope)`

SetScope sets Scope field to given value.


### GetLocation

`func (o *Network) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Network) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Network) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Network) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProject

`func (o *Network) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Network) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Network) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Network) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetNotifications

`func (o *Network) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *Network) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *Network) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.


### GetHref

`func (o *Network) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Network) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Network) SetHref(v string)`

SetHref sets Href field to given value.


### GetUuid

`func (o *Network) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Network) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Network) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetState

`func (o *Network) GetState() NetworkState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Network) GetStateOk() (*NetworkState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Network) SetState(v NetworkState)`

SetState sets State field to given value.


### GetConnectionsCount

`func (o *Network) GetConnectionsCount() float32`

GetConnectionsCount returns the ConnectionsCount field if non-nil, zero value otherwise.

### GetConnectionsCountOk

`func (o *Network) GetConnectionsCountOk() (*float32, bool)`

GetConnectionsCountOk returns a tuple with the ConnectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsCount

`func (o *Network) SetConnectionsCount(v float32)`

SetConnectionsCount sets ConnectionsCount field to given value.

### HasConnectionsCount

`func (o *Network) HasConnectionsCount() bool`

HasConnectionsCount returns a boolean if a field has been set.

### GetAccount

`func (o *Network) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Network) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Network) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Network) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetChange

`func (o *Network) GetChange() SimplifiedNetworkChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *Network) GetChangeOk() (*SimplifiedNetworkChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *Network) SetChange(v SimplifiedNetworkChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *Network) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetOperation

`func (o *Network) GetOperation() NetworkOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Network) GetOperationOk() (*NetworkOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Network) SetOperation(v NetworkOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *Network) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetChangeLog

`func (o *Network) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *Network) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *Network) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.


### GetLinks

`func (o *Network) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Network) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Network) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Network) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


