# OpticalConnectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | URI of this Optical Connect resource. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Unique identifier of this Optical Connect. | [optional] [readonly] 
**Type** | Pointer to [**OpticalConnectResponseType**](OpticalConnectResponseType.md) |  | [optional] 
**Name** | Pointer to **string** | Equinix-assigned name, derived from the account number and the two IBX locations.  | [optional] [readonly] 
**State** | Pointer to [**OpticalConnectState**](OpticalConnectState.md) |  | [optional] 
**Bandwidth** | Pointer to **int32** | Provisioned connection bandwidth in Mbps. | [optional] 
**ConnectionDestinationType** | Pointer to [**OpticalConnectResponseConnectionDestinationType**](OpticalConnectResponseConnectionDestinationType.md) |  | [optional] 
**PathType** | Pointer to [**OpticalConnectResponsePathType**](OpticalConnectResponsePathType.md) |  | [optional] 
**BmmrType** | Pointer to [**OpticalConnectResponseBmmrType**](OpticalConnectResponseBmmrType.md) |  | [optional] 
**Redundancy** | Pointer to [**OpticalConnectRedundancy**](OpticalConnectRedundancy.md) |  | [optional] 
**ASide** | Pointer to [**OpticalConnectASideResponse**](OpticalConnectASideResponse.md) |  | [optional] 
**ZSide** | Pointer to [**OpticalConnectZSideResponse**](OpticalConnectZSideResponse.md) |  | [optional] 
**Order** | Pointer to [**OpticalConnectOrder**](OpticalConnectOrder.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Notifications** | Pointer to [**[]OpticalConnectNotification**](OpticalConnectNotification.md) | Preferences for notifications on connection configuration or status changes | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewOpticalConnectResponse

`func NewOpticalConnectResponse() *OpticalConnectResponse`

NewOpticalConnectResponse instantiates a new OpticalConnectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpticalConnectResponseWithDefaults

`func NewOpticalConnectResponseWithDefaults() *OpticalConnectResponse`

NewOpticalConnectResponseWithDefaults instantiates a new OpticalConnectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *OpticalConnectResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *OpticalConnectResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *OpticalConnectResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *OpticalConnectResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *OpticalConnectResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OpticalConnectResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OpticalConnectResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OpticalConnectResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *OpticalConnectResponse) GetType() OpticalConnectResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpticalConnectResponse) GetTypeOk() (*OpticalConnectResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpticalConnectResponse) SetType(v OpticalConnectResponseType)`

SetType sets Type field to given value.

### HasType

`func (o *OpticalConnectResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *OpticalConnectResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpticalConnectResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpticalConnectResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpticalConnectResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *OpticalConnectResponse) GetState() OpticalConnectState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OpticalConnectResponse) GetStateOk() (*OpticalConnectState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OpticalConnectResponse) SetState(v OpticalConnectState)`

SetState sets State field to given value.

### HasState

`func (o *OpticalConnectResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetBandwidth

`func (o *OpticalConnectResponse) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *OpticalConnectResponse) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *OpticalConnectResponse) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *OpticalConnectResponse) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetConnectionDestinationType

`func (o *OpticalConnectResponse) GetConnectionDestinationType() OpticalConnectResponseConnectionDestinationType`

GetConnectionDestinationType returns the ConnectionDestinationType field if non-nil, zero value otherwise.

### GetConnectionDestinationTypeOk

`func (o *OpticalConnectResponse) GetConnectionDestinationTypeOk() (*OpticalConnectResponseConnectionDestinationType, bool)`

GetConnectionDestinationTypeOk returns a tuple with the ConnectionDestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDestinationType

`func (o *OpticalConnectResponse) SetConnectionDestinationType(v OpticalConnectResponseConnectionDestinationType)`

SetConnectionDestinationType sets ConnectionDestinationType field to given value.

### HasConnectionDestinationType

`func (o *OpticalConnectResponse) HasConnectionDestinationType() bool`

HasConnectionDestinationType returns a boolean if a field has been set.

### GetPathType

`func (o *OpticalConnectResponse) GetPathType() OpticalConnectResponsePathType`

GetPathType returns the PathType field if non-nil, zero value otherwise.

### GetPathTypeOk

`func (o *OpticalConnectResponse) GetPathTypeOk() (*OpticalConnectResponsePathType, bool)`

GetPathTypeOk returns a tuple with the PathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathType

`func (o *OpticalConnectResponse) SetPathType(v OpticalConnectResponsePathType)`

SetPathType sets PathType field to given value.

### HasPathType

`func (o *OpticalConnectResponse) HasPathType() bool`

HasPathType returns a boolean if a field has been set.

### GetBmmrType

`func (o *OpticalConnectResponse) GetBmmrType() OpticalConnectResponseBmmrType`

GetBmmrType returns the BmmrType field if non-nil, zero value otherwise.

### GetBmmrTypeOk

`func (o *OpticalConnectResponse) GetBmmrTypeOk() (*OpticalConnectResponseBmmrType, bool)`

GetBmmrTypeOk returns a tuple with the BmmrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmmrType

`func (o *OpticalConnectResponse) SetBmmrType(v OpticalConnectResponseBmmrType)`

SetBmmrType sets BmmrType field to given value.

### HasBmmrType

`func (o *OpticalConnectResponse) HasBmmrType() bool`

HasBmmrType returns a boolean if a field has been set.

### GetRedundancy

`func (o *OpticalConnectResponse) GetRedundancy() OpticalConnectRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *OpticalConnectResponse) GetRedundancyOk() (*OpticalConnectRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *OpticalConnectResponse) SetRedundancy(v OpticalConnectRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *OpticalConnectResponse) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetASide

`func (o *OpticalConnectResponse) GetASide() OpticalConnectASideResponse`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *OpticalConnectResponse) GetASideOk() (*OpticalConnectASideResponse, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *OpticalConnectResponse) SetASide(v OpticalConnectASideResponse)`

SetASide sets ASide field to given value.

### HasASide

`func (o *OpticalConnectResponse) HasASide() bool`

HasASide returns a boolean if a field has been set.

### GetZSide

`func (o *OpticalConnectResponse) GetZSide() OpticalConnectZSideResponse`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *OpticalConnectResponse) GetZSideOk() (*OpticalConnectZSideResponse, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *OpticalConnectResponse) SetZSide(v OpticalConnectZSideResponse)`

SetZSide sets ZSide field to given value.

### HasZSide

`func (o *OpticalConnectResponse) HasZSide() bool`

HasZSide returns a boolean if a field has been set.

### GetOrder

`func (o *OpticalConnectResponse) GetOrder() OpticalConnectOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OpticalConnectResponse) GetOrderOk() (*OpticalConnectOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OpticalConnectResponse) SetOrder(v OpticalConnectOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OpticalConnectResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAccount

`func (o *OpticalConnectResponse) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OpticalConnectResponse) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OpticalConnectResponse) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OpticalConnectResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetNotifications

`func (o *OpticalConnectResponse) GetNotifications() []OpticalConnectNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *OpticalConnectResponse) GetNotificationsOk() (*[]OpticalConnectNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *OpticalConnectResponse) SetNotifications(v []OpticalConnectNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *OpticalConnectResponse) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetChangeLog

`func (o *OpticalConnectResponse) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *OpticalConnectResponse) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *OpticalConnectResponse) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *OpticalConnectResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


