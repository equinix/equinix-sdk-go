# MetroConnectPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Metro Connect Type | 
**Bandwidth** | **int32** | Bandwidth in Mbps | 
**ConnectivityDestinationType** | [**MetroConnectPostRequestConnectivityDestinationType**](MetroConnectPostRequestConnectivityDestinationType.md) |  | 
**ASide** | [**MetroConnectASide**](MetroConnectASide.md) |  | 
**ZSide** | [**MetroConnectZSide**](MetroConnectZSide.md) |  | 
**Order** | Pointer to [**MetroConnectOrder**](MetroConnectOrder.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Account** | [**SimplifiedAccount**](SimplifiedAccount.md) |  | 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on connection configuration or status changes | [optional] 

## Methods

### NewMetroConnectPostRequest

`func NewMetroConnectPostRequest(type_ string, bandwidth int32, connectivityDestinationType MetroConnectPostRequestConnectivityDestinationType, aSide MetroConnectASide, zSide MetroConnectZSide, account SimplifiedAccount, ) *MetroConnectPostRequest`

NewMetroConnectPostRequest instantiates a new MetroConnectPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroConnectPostRequestWithDefaults

`func NewMetroConnectPostRequestWithDefaults() *MetroConnectPostRequest`

NewMetroConnectPostRequestWithDefaults instantiates a new MetroConnectPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetroConnectPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetroConnectPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetroConnectPostRequest) SetType(v string)`

SetType sets Type field to given value.


### GetBandwidth

`func (o *MetroConnectPostRequest) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *MetroConnectPostRequest) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *MetroConnectPostRequest) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetConnectivityDestinationType

`func (o *MetroConnectPostRequest) GetConnectivityDestinationType() MetroConnectPostRequestConnectivityDestinationType`

GetConnectivityDestinationType returns the ConnectivityDestinationType field if non-nil, zero value otherwise.

### GetConnectivityDestinationTypeOk

`func (o *MetroConnectPostRequest) GetConnectivityDestinationTypeOk() (*MetroConnectPostRequestConnectivityDestinationType, bool)`

GetConnectivityDestinationTypeOk returns a tuple with the ConnectivityDestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityDestinationType

`func (o *MetroConnectPostRequest) SetConnectivityDestinationType(v MetroConnectPostRequestConnectivityDestinationType)`

SetConnectivityDestinationType sets ConnectivityDestinationType field to given value.


### GetASide

`func (o *MetroConnectPostRequest) GetASide() MetroConnectASide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *MetroConnectPostRequest) GetASideOk() (*MetroConnectASide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *MetroConnectPostRequest) SetASide(v MetroConnectASide)`

SetASide sets ASide field to given value.


### GetZSide

`func (o *MetroConnectPostRequest) GetZSide() MetroConnectZSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *MetroConnectPostRequest) GetZSideOk() (*MetroConnectZSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *MetroConnectPostRequest) SetZSide(v MetroConnectZSide)`

SetZSide sets ZSide field to given value.


### GetOrder

`func (o *MetroConnectPostRequest) GetOrder() MetroConnectOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MetroConnectPostRequest) GetOrderOk() (*MetroConnectOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MetroConnectPostRequest) SetOrder(v MetroConnectOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MetroConnectPostRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProject

`func (o *MetroConnectPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *MetroConnectPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *MetroConnectPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *MetroConnectPostRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAccount

`func (o *MetroConnectPostRequest) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *MetroConnectPostRequest) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *MetroConnectPostRequest) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.


### GetNotifications

`func (o *MetroConnectPostRequest) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *MetroConnectPostRequest) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *MetroConnectPostRequest) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *MetroConnectPostRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


