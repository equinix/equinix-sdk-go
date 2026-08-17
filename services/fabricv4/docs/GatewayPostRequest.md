# GatewayPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**GatewayType**](GatewayType.md) |  | 
**Name** | **string** | Customer-provided Gateway name | 
**Description** | Pointer to **string** | Customer-provided Gateway description | [optional] 
**Bandwidth** | **int32** | Gateway bandwidth in Mbps | 
**LocalAsn** | **int32** | Gateway local Autonomous System Number | 
**Router** | [**Router**](Router.md) |  | 
**HaEnabled** | Pointer to **bool** | High availability enabled | [optional] [default to false]
**Account** | [**SimplifiedAccount**](SimplifiedAccount.md) |  | 
**Project** | [**Project**](Project.md) |  | 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 

## Methods

### NewGatewayPostRequest

`func NewGatewayPostRequest(type_ GatewayType, name string, bandwidth int32, localAsn int32, router Router, account SimplifiedAccount, project Project, ) *GatewayPostRequest`

NewGatewayPostRequest instantiates a new GatewayPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPostRequestWithDefaults

`func NewGatewayPostRequestWithDefaults() *GatewayPostRequest`

NewGatewayPostRequestWithDefaults instantiates a new GatewayPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GatewayPostRequest) GetType() GatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayPostRequest) GetTypeOk() (*GatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayPostRequest) SetType(v GatewayType)`

SetType sets Type field to given value.


### GetName

`func (o *GatewayPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GatewayPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GatewayPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBandwidth

`func (o *GatewayPostRequest) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *GatewayPostRequest) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *GatewayPostRequest) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetLocalAsn

`func (o *GatewayPostRequest) GetLocalAsn() int32`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *GatewayPostRequest) GetLocalAsnOk() (*int32, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *GatewayPostRequest) SetLocalAsn(v int32)`

SetLocalAsn sets LocalAsn field to given value.


### GetRouter

`func (o *GatewayPostRequest) GetRouter() Router`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *GatewayPostRequest) GetRouterOk() (*Router, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *GatewayPostRequest) SetRouter(v Router)`

SetRouter sets Router field to given value.


### GetHaEnabled

`func (o *GatewayPostRequest) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *GatewayPostRequest) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *GatewayPostRequest) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *GatewayPostRequest) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetAccount

`func (o *GatewayPostRequest) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GatewayPostRequest) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GatewayPostRequest) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.


### GetProject

`func (o *GatewayPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GatewayPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GatewayPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.


### GetOrder

`func (o *GatewayPostRequest) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GatewayPostRequest) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GatewayPostRequest) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GatewayPostRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


