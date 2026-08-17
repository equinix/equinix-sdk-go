# Interconnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Interconnect URI | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned interconnect identifier | [optional] 
**Type** | Pointer to [**InterconnectPostRequestType**](InterconnectPostRequestType.md) |  | [optional] 
**Name** | Pointer to **string** | Interconnect name | [optional] 
**Description** | Pointer to **string** | Interconnect description | [optional] 
**State** | Pointer to [**InterconnectState**](InterconnectState.md) |  | [optional] 
**Location** | Pointer to [**InterconnectLocation**](InterconnectLocation.md) |  | [optional] 
**UsedBandwidth** | Pointer to **int32** | Interconnect used bandwidth in Mbps | [optional] 
**Package** | Pointer to [**InterconnectPackage**](InterconnectPackage.md) |  | [optional] 
**Router** | Pointer to [**InterconnectRouter**](InterconnectRouter.md) |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Notifications** | Pointer to [**[]InterconnectNotification**](InterconnectNotification.md) | Interconnect notification preferences | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewInterconnect

`func NewInterconnect() *Interconnect`

NewInterconnect instantiates a new Interconnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectWithDefaults

`func NewInterconnectWithDefaults() *Interconnect`

NewInterconnectWithDefaults instantiates a new Interconnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Interconnect) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Interconnect) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Interconnect) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Interconnect) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *Interconnect) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Interconnect) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Interconnect) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Interconnect) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *Interconnect) GetType() InterconnectPostRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Interconnect) GetTypeOk() (*InterconnectPostRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Interconnect) SetType(v InterconnectPostRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *Interconnect) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Interconnect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Interconnect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Interconnect) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Interconnect) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Interconnect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Interconnect) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Interconnect) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Interconnect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *Interconnect) GetState() InterconnectState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Interconnect) GetStateOk() (*InterconnectState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Interconnect) SetState(v InterconnectState)`

SetState sets State field to given value.

### HasState

`func (o *Interconnect) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLocation

`func (o *Interconnect) GetLocation() InterconnectLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Interconnect) GetLocationOk() (*InterconnectLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Interconnect) SetLocation(v InterconnectLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Interconnect) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetUsedBandwidth

`func (o *Interconnect) GetUsedBandwidth() int32`

GetUsedBandwidth returns the UsedBandwidth field if non-nil, zero value otherwise.

### GetUsedBandwidthOk

`func (o *Interconnect) GetUsedBandwidthOk() (*int32, bool)`

GetUsedBandwidthOk returns a tuple with the UsedBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBandwidth

`func (o *Interconnect) SetUsedBandwidth(v int32)`

SetUsedBandwidth sets UsedBandwidth field to given value.

### HasUsedBandwidth

`func (o *Interconnect) HasUsedBandwidth() bool`

HasUsedBandwidth returns a boolean if a field has been set.

### GetPackage

`func (o *Interconnect) GetPackage() InterconnectPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *Interconnect) GetPackageOk() (*InterconnectPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *Interconnect) SetPackage(v InterconnectPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *Interconnect) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetRouter

`func (o *Interconnect) GetRouter() InterconnectRouter`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *Interconnect) GetRouterOk() (*InterconnectRouter, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *Interconnect) SetRouter(v InterconnectRouter)`

SetRouter sets Router field to given value.

### HasRouter

`func (o *Interconnect) HasRouter() bool`

HasRouter returns a boolean if a field has been set.

### GetOrder

`func (o *Interconnect) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Interconnect) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Interconnect) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Interconnect) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAccount

`func (o *Interconnect) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Interconnect) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Interconnect) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Interconnect) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetProject

`func (o *Interconnect) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Interconnect) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Interconnect) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Interconnect) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetNotifications

`func (o *Interconnect) GetNotifications() []InterconnectNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *Interconnect) GetNotificationsOk() (*[]InterconnectNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *Interconnect) SetNotifications(v []InterconnectNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *Interconnect) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetChangeLog

`func (o *Interconnect) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *Interconnect) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *Interconnect) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *Interconnect) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


