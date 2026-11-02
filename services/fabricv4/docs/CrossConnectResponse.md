# CrossConnectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Cross Connect URI. | [optional] 
**Type** | Pointer to [**CrossConnectType**](CrossConnectType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Cross Connect Identifier. | [optional] 
**Name** | Pointer to **string** | A short, descriptive name for this Cross Connect. | [optional] 
**State** | Pointer to [**CrossConnectState**](CrossConnectState.md) |  | [optional] 
**AllowedActions** | Pointer to [**[]CrossConnectResponseAllowedActionsInner**](CrossConnectResponseAllowedActionsInner.md) | List of actions that can be performed on the Cross Connect based on its current state. Use these values to determine which action types are valid to submit to the actions endpoint.  | [optional] [readonly] 
**BmmrType** | Pointer to [**BmmrType**](BmmrType.md) |  | [optional] 
**MediaType** | Pointer to **string** | Media type of the Cross Connect. | [optional] 
**MediaSubType** | Pointer to **string** | Media sub-type of the Cross Connect. | [optional] 
**CircuitId** | Pointer to **string** | Circuit identifier of the Cross Connect. | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the Cross Connect. | [optional] [readonly] 
**Redundancy** | Pointer to [**CrossConnectRedundancy**](CrossConnectRedundancy.md) |  | [optional] 
**ASide** | Pointer to [**CrossConnectASideResponse**](CrossConnectASideResponse.md) |  | [optional] 
**ZSide** | Pointer to [**CrossConnectZSideResponse**](CrossConnectZSideResponse.md) |  | [optional] 
**Account** | Pointer to [**CrossConnectAccount**](CrossConnectAccount.md) |  | [optional] 
**Notifications** | Pointer to [**[]CrossConnectNotification**](CrossConnectNotification.md) | Preferences for notifications on Cross Connect configuration or status changes.  | [optional] 
**Order** | Pointer to [**OrderResponse**](OrderResponse.md) |  | [optional] 
**Links** | Pointer to [**[]CrossConnectLink**](CrossConnectLink.md) | Links to the Cross Connect sub-resources - schedules, services, notes and attachments.  | [optional] [readonly] 
**ChangeLog** | Pointer to [**CrossConnectChangelog**](CrossConnectChangelog.md) |  | [optional] 

## Methods

### NewCrossConnectResponse

`func NewCrossConnectResponse() *CrossConnectResponse`

NewCrossConnectResponse instantiates a new CrossConnectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectResponseWithDefaults

`func NewCrossConnectResponseWithDefaults() *CrossConnectResponse`

NewCrossConnectResponseWithDefaults instantiates a new CrossConnectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CrossConnectResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CrossConnectResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CrossConnectResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CrossConnectResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *CrossConnectResponse) GetType() CrossConnectType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossConnectResponse) GetTypeOk() (*CrossConnectType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossConnectResponse) SetType(v CrossConnectType)`

SetType sets Type field to given value.

### HasType

`func (o *CrossConnectResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *CrossConnectResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CrossConnectResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CrossConnectResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CrossConnectResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *CrossConnectResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrossConnectResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrossConnectResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CrossConnectResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *CrossConnectResponse) GetState() CrossConnectState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CrossConnectResponse) GetStateOk() (*CrossConnectState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CrossConnectResponse) SetState(v CrossConnectState)`

SetState sets State field to given value.

### HasState

`func (o *CrossConnectResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAllowedActions

`func (o *CrossConnectResponse) GetAllowedActions() []CrossConnectResponseAllowedActionsInner`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *CrossConnectResponse) GetAllowedActionsOk() (*[]CrossConnectResponseAllowedActionsInner, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *CrossConnectResponse) SetAllowedActions(v []CrossConnectResponseAllowedActionsInner)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *CrossConnectResponse) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.

### GetBmmrType

`func (o *CrossConnectResponse) GetBmmrType() BmmrType`

GetBmmrType returns the BmmrType field if non-nil, zero value otherwise.

### GetBmmrTypeOk

`func (o *CrossConnectResponse) GetBmmrTypeOk() (*BmmrType, bool)`

GetBmmrTypeOk returns a tuple with the BmmrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmmrType

`func (o *CrossConnectResponse) SetBmmrType(v BmmrType)`

SetBmmrType sets BmmrType field to given value.

### HasBmmrType

`func (o *CrossConnectResponse) HasBmmrType() bool`

HasBmmrType returns a boolean if a field has been set.

### GetMediaType

`func (o *CrossConnectResponse) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *CrossConnectResponse) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *CrossConnectResponse) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *CrossConnectResponse) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetMediaSubType

`func (o *CrossConnectResponse) GetMediaSubType() string`

GetMediaSubType returns the MediaSubType field if non-nil, zero value otherwise.

### GetMediaSubTypeOk

`func (o *CrossConnectResponse) GetMediaSubTypeOk() (*string, bool)`

GetMediaSubTypeOk returns a tuple with the MediaSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSubType

`func (o *CrossConnectResponse) SetMediaSubType(v string)`

SetMediaSubType sets MediaSubType field to given value.

### HasMediaSubType

`func (o *CrossConnectResponse) HasMediaSubType() bool`

HasMediaSubType returns a boolean if a field has been set.

### GetCircuitId

`func (o *CrossConnectResponse) GetCircuitId() string`

GetCircuitId returns the CircuitId field if non-nil, zero value otherwise.

### GetCircuitIdOk

`func (o *CrossConnectResponse) GetCircuitIdOk() (*string, bool)`

GetCircuitIdOk returns a tuple with the CircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitId

`func (o *CrossConnectResponse) SetCircuitId(v string)`

SetCircuitId sets CircuitId field to given value.

### HasCircuitId

`func (o *CrossConnectResponse) HasCircuitId() bool`

HasCircuitId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *CrossConnectResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CrossConnectResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CrossConnectResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CrossConnectResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetRedundancy

`func (o *CrossConnectResponse) GetRedundancy() CrossConnectRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *CrossConnectResponse) GetRedundancyOk() (*CrossConnectRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *CrossConnectResponse) SetRedundancy(v CrossConnectRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *CrossConnectResponse) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetASide

`func (o *CrossConnectResponse) GetASide() CrossConnectASideResponse`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *CrossConnectResponse) GetASideOk() (*CrossConnectASideResponse, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *CrossConnectResponse) SetASide(v CrossConnectASideResponse)`

SetASide sets ASide field to given value.

### HasASide

`func (o *CrossConnectResponse) HasASide() bool`

HasASide returns a boolean if a field has been set.

### GetZSide

`func (o *CrossConnectResponse) GetZSide() CrossConnectZSideResponse`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *CrossConnectResponse) GetZSideOk() (*CrossConnectZSideResponse, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *CrossConnectResponse) SetZSide(v CrossConnectZSideResponse)`

SetZSide sets ZSide field to given value.

### HasZSide

`func (o *CrossConnectResponse) HasZSide() bool`

HasZSide returns a boolean if a field has been set.

### GetAccount

`func (o *CrossConnectResponse) GetAccount() CrossConnectAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CrossConnectResponse) GetAccountOk() (*CrossConnectAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CrossConnectResponse) SetAccount(v CrossConnectAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CrossConnectResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetNotifications

`func (o *CrossConnectResponse) GetNotifications() []CrossConnectNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *CrossConnectResponse) GetNotificationsOk() (*[]CrossConnectNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *CrossConnectResponse) SetNotifications(v []CrossConnectNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *CrossConnectResponse) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOrder

`func (o *CrossConnectResponse) GetOrder() OrderResponse`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CrossConnectResponse) GetOrderOk() (*OrderResponse, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CrossConnectResponse) SetOrder(v OrderResponse)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CrossConnectResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetLinks

`func (o *CrossConnectResponse) GetLinks() []CrossConnectLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CrossConnectResponse) GetLinksOk() (*[]CrossConnectLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CrossConnectResponse) SetLinks(v []CrossConnectLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CrossConnectResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetChangeLog

`func (o *CrossConnectResponse) GetChangeLog() CrossConnectChangelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *CrossConnectResponse) GetChangeLogOk() (*CrossConnectChangelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *CrossConnectResponse) SetChangeLog(v CrossConnectChangelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *CrossConnectResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


