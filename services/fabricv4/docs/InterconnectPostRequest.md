# InterconnectPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**InterconnectPostRequestType**](InterconnectPostRequestType.md) |  | 
**Name** | **string** | Customer-provided interconnect name | 
**Description** | Pointer to **string** | Customer-provided interconnect description | [optional] 
**Location** | [**InterconnectLocationRequest**](InterconnectLocationRequest.md) |  | 
**Package** | [**InterconnectPackage**](InterconnectPackage.md) |  | 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Account** | [**SimplifiedAccount**](SimplifiedAccount.md) |  | 
**Project** | [**Project**](Project.md) |  | 
**Notifications** | [**[]InterconnectNotification**](InterconnectNotification.md) | Preferences for notifications on interconnect configuration or status changes | 

## Methods

### NewInterconnectPostRequest

`func NewInterconnectPostRequest(type_ InterconnectPostRequestType, name string, location InterconnectLocationRequest, package_ InterconnectPackage, account SimplifiedAccount, project Project, notifications []InterconnectNotification, ) *InterconnectPostRequest`

NewInterconnectPostRequest instantiates a new InterconnectPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectPostRequestWithDefaults

`func NewInterconnectPostRequestWithDefaults() *InterconnectPostRequest`

NewInterconnectPostRequestWithDefaults instantiates a new InterconnectPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InterconnectPostRequest) GetType() InterconnectPostRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterconnectPostRequest) GetTypeOk() (*InterconnectPostRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterconnectPostRequest) SetType(v InterconnectPostRequestType)`

SetType sets Type field to given value.


### GetName

`func (o *InterconnectPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterconnectPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterconnectPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InterconnectPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterconnectPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterconnectPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterconnectPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *InterconnectPostRequest) GetLocation() InterconnectLocationRequest`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InterconnectPostRequest) GetLocationOk() (*InterconnectLocationRequest, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InterconnectPostRequest) SetLocation(v InterconnectLocationRequest)`

SetLocation sets Location field to given value.


### GetPackage

`func (o *InterconnectPostRequest) GetPackage() InterconnectPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *InterconnectPostRequest) GetPackageOk() (*InterconnectPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *InterconnectPostRequest) SetPackage(v InterconnectPackage)`

SetPackage sets Package field to given value.


### GetOrder

`func (o *InterconnectPostRequest) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InterconnectPostRequest) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InterconnectPostRequest) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *InterconnectPostRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAccount

`func (o *InterconnectPostRequest) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InterconnectPostRequest) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InterconnectPostRequest) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.


### GetProject

`func (o *InterconnectPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *InterconnectPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *InterconnectPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.


### GetNotifications

`func (o *InterconnectPostRequest) GetNotifications() []InterconnectNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *InterconnectPostRequest) GetNotificationsOk() (*[]InterconnectNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *InterconnectPostRequest) SetNotifications(v []InterconnectNotification)`

SetNotifications sets Notifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


