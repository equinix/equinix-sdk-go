# OpticalConnectPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**OpticalConnectPostRequestType**](OpticalConnectPostRequestType.md) |  | 
**Bandwidth** | **int32** | Connection bandwidth Mbps. &lt;br&gt; Available bandwidths depend on the IBX pair. &lt;br&gt; 1000 - 1 Gbps. &lt;br&gt; 10000 - 10 Gbps. &lt;br&gt; 100000 - 100 Gbps. &lt;br&gt;  | 
**ConnectionDestinationType** | [**OpticalConnectPostRequestConnectionDestinationType**](OpticalConnectPostRequestConnectionDestinationType.md) |  | 
**PathType** | [**OpticalConnectPostRequestPathType**](OpticalConnectPostRequestPathType.md) |  | 
**Redundancy** | Pointer to [**OpticalConnectRedundancy**](OpticalConnectRedundancy.md) |  | [optional] 
**ASide** | [**OpticalConnectASideRequest**](OpticalConnectASideRequest.md) |  | 
**ZSide** | [**OpticalConnectZSideRequest**](OpticalConnectZSideRequest.md) |  | 
**Order** | Pointer to [**OpticalConnectOrder**](OpticalConnectOrder.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Notifications** | Pointer to [**[]OpticalConnectNotification**](OpticalConnectNotification.md) | Contacts to notify about connection configuration and status changes | [optional] 
**BmmrType** | Pointer to [**OpticalConnectPostRequestBmmrType**](OpticalConnectPostRequestBmmrType.md) |  | [optional] 

## Methods

### NewOpticalConnectPostRequest

`func NewOpticalConnectPostRequest(type_ OpticalConnectPostRequestType, bandwidth int32, connectionDestinationType OpticalConnectPostRequestConnectionDestinationType, pathType OpticalConnectPostRequestPathType, aSide OpticalConnectASideRequest, zSide OpticalConnectZSideRequest, ) *OpticalConnectPostRequest`

NewOpticalConnectPostRequest instantiates a new OpticalConnectPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpticalConnectPostRequestWithDefaults

`func NewOpticalConnectPostRequestWithDefaults() *OpticalConnectPostRequest`

NewOpticalConnectPostRequestWithDefaults instantiates a new OpticalConnectPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OpticalConnectPostRequest) GetType() OpticalConnectPostRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpticalConnectPostRequest) GetTypeOk() (*OpticalConnectPostRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpticalConnectPostRequest) SetType(v OpticalConnectPostRequestType)`

SetType sets Type field to given value.


### GetBandwidth

`func (o *OpticalConnectPostRequest) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *OpticalConnectPostRequest) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *OpticalConnectPostRequest) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetConnectionDestinationType

`func (o *OpticalConnectPostRequest) GetConnectionDestinationType() OpticalConnectPostRequestConnectionDestinationType`

GetConnectionDestinationType returns the ConnectionDestinationType field if non-nil, zero value otherwise.

### GetConnectionDestinationTypeOk

`func (o *OpticalConnectPostRequest) GetConnectionDestinationTypeOk() (*OpticalConnectPostRequestConnectionDestinationType, bool)`

GetConnectionDestinationTypeOk returns a tuple with the ConnectionDestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDestinationType

`func (o *OpticalConnectPostRequest) SetConnectionDestinationType(v OpticalConnectPostRequestConnectionDestinationType)`

SetConnectionDestinationType sets ConnectionDestinationType field to given value.


### GetPathType

`func (o *OpticalConnectPostRequest) GetPathType() OpticalConnectPostRequestPathType`

GetPathType returns the PathType field if non-nil, zero value otherwise.

### GetPathTypeOk

`func (o *OpticalConnectPostRequest) GetPathTypeOk() (*OpticalConnectPostRequestPathType, bool)`

GetPathTypeOk returns a tuple with the PathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathType

`func (o *OpticalConnectPostRequest) SetPathType(v OpticalConnectPostRequestPathType)`

SetPathType sets PathType field to given value.


### GetRedundancy

`func (o *OpticalConnectPostRequest) GetRedundancy() OpticalConnectRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *OpticalConnectPostRequest) GetRedundancyOk() (*OpticalConnectRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *OpticalConnectPostRequest) SetRedundancy(v OpticalConnectRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *OpticalConnectPostRequest) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetASide

`func (o *OpticalConnectPostRequest) GetASide() OpticalConnectASideRequest`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *OpticalConnectPostRequest) GetASideOk() (*OpticalConnectASideRequest, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *OpticalConnectPostRequest) SetASide(v OpticalConnectASideRequest)`

SetASide sets ASide field to given value.


### GetZSide

`func (o *OpticalConnectPostRequest) GetZSide() OpticalConnectZSideRequest`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *OpticalConnectPostRequest) GetZSideOk() (*OpticalConnectZSideRequest, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *OpticalConnectPostRequest) SetZSide(v OpticalConnectZSideRequest)`

SetZSide sets ZSide field to given value.


### GetOrder

`func (o *OpticalConnectPostRequest) GetOrder() OpticalConnectOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OpticalConnectPostRequest) GetOrderOk() (*OpticalConnectOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OpticalConnectPostRequest) SetOrder(v OpticalConnectOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OpticalConnectPostRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAccount

`func (o *OpticalConnectPostRequest) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OpticalConnectPostRequest) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OpticalConnectPostRequest) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OpticalConnectPostRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetNotifications

`func (o *OpticalConnectPostRequest) GetNotifications() []OpticalConnectNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *OpticalConnectPostRequest) GetNotificationsOk() (*[]OpticalConnectNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *OpticalConnectPostRequest) SetNotifications(v []OpticalConnectNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *OpticalConnectPostRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetBmmrType

`func (o *OpticalConnectPostRequest) GetBmmrType() OpticalConnectPostRequestBmmrType`

GetBmmrType returns the BmmrType field if non-nil, zero value otherwise.

### GetBmmrTypeOk

`func (o *OpticalConnectPostRequest) GetBmmrTypeOk() (*OpticalConnectPostRequestBmmrType, bool)`

GetBmmrTypeOk returns a tuple with the BmmrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmmrType

`func (o *OpticalConnectPostRequest) SetBmmrType(v OpticalConnectPostRequestBmmrType)`

SetBmmrType sets BmmrType field to given value.

### HasBmmrType

`func (o *OpticalConnectPostRequest) HasBmmrType() bool`

HasBmmrType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


