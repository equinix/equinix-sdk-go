# ExchangeServiceCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ExchangeServiceCreateRequestType**](ExchangeServiceCreateRequestType.md) |  | 
**Name** | **string** | Customer-provided internet exchange service name | 
**Description** | Pointer to **string** | Customer-provided internet exchange service description | [optional] 
**PublicPeeringConnection** | [**PublicPeeringConnectionRequest**](PublicPeeringConnectionRequest.md) |  | 
**RoutingProtocol** | [**RoutingProtocolRequest**](RoutingProtocolRequest.md) |  | 
**Order** | Pointer to [**ExchangeServiceOrder**](ExchangeServiceOrder.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Notifications** | Pointer to [**[]ExchangeServiceNotification**](ExchangeServiceNotification.md) | Preferences for notifications on Internet Exchange Service configuration or status changes | [optional] 

## Methods

### NewExchangeServiceCreateRequest

`func NewExchangeServiceCreateRequest(type_ ExchangeServiceCreateRequestType, name string, publicPeeringConnection PublicPeeringConnectionRequest, routingProtocol RoutingProtocolRequest, ) *ExchangeServiceCreateRequest`

NewExchangeServiceCreateRequest instantiates a new ExchangeServiceCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeServiceCreateRequestWithDefaults

`func NewExchangeServiceCreateRequestWithDefaults() *ExchangeServiceCreateRequest`

NewExchangeServiceCreateRequestWithDefaults instantiates a new ExchangeServiceCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExchangeServiceCreateRequest) GetType() ExchangeServiceCreateRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExchangeServiceCreateRequest) GetTypeOk() (*ExchangeServiceCreateRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExchangeServiceCreateRequest) SetType(v ExchangeServiceCreateRequestType)`

SetType sets Type field to given value.


### GetName

`func (o *ExchangeServiceCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangeServiceCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangeServiceCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExchangeServiceCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExchangeServiceCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExchangeServiceCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExchangeServiceCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPublicPeeringConnection

`func (o *ExchangeServiceCreateRequest) GetPublicPeeringConnection() PublicPeeringConnectionRequest`

GetPublicPeeringConnection returns the PublicPeeringConnection field if non-nil, zero value otherwise.

### GetPublicPeeringConnectionOk

`func (o *ExchangeServiceCreateRequest) GetPublicPeeringConnectionOk() (*PublicPeeringConnectionRequest, bool)`

GetPublicPeeringConnectionOk returns a tuple with the PublicPeeringConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPeeringConnection

`func (o *ExchangeServiceCreateRequest) SetPublicPeeringConnection(v PublicPeeringConnectionRequest)`

SetPublicPeeringConnection sets PublicPeeringConnection field to given value.


### GetRoutingProtocol

`func (o *ExchangeServiceCreateRequest) GetRoutingProtocol() RoutingProtocolRequest`

GetRoutingProtocol returns the RoutingProtocol field if non-nil, zero value otherwise.

### GetRoutingProtocolOk

`func (o *ExchangeServiceCreateRequest) GetRoutingProtocolOk() (*RoutingProtocolRequest, bool)`

GetRoutingProtocolOk returns a tuple with the RoutingProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingProtocol

`func (o *ExchangeServiceCreateRequest) SetRoutingProtocol(v RoutingProtocolRequest)`

SetRoutingProtocol sets RoutingProtocol field to given value.


### GetOrder

`func (o *ExchangeServiceCreateRequest) GetOrder() ExchangeServiceOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ExchangeServiceCreateRequest) GetOrderOk() (*ExchangeServiceOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ExchangeServiceCreateRequest) SetOrder(v ExchangeServiceOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ExchangeServiceCreateRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProject

`func (o *ExchangeServiceCreateRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ExchangeServiceCreateRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ExchangeServiceCreateRequest) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *ExchangeServiceCreateRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAccount

`func (o *ExchangeServiceCreateRequest) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ExchangeServiceCreateRequest) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ExchangeServiceCreateRequest) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ExchangeServiceCreateRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetNotifications

`func (o *ExchangeServiceCreateRequest) GetNotifications() []ExchangeServiceNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ExchangeServiceCreateRequest) GetNotificationsOk() (*[]ExchangeServiceNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ExchangeServiceCreateRequest) SetNotifications(v []ExchangeServiceNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ExchangeServiceCreateRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


