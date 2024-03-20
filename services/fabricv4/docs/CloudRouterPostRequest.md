# CloudRouterPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CloudRouterPostRequestType**](CloudRouterPostRequestType.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided Cloud Router name | [optional] 
**Location** | Pointer to [**SimplifiedLocationWithoutIBX**](SimplifiedLocationWithoutIBX.md) |  | [optional] 
**Package** | Pointer to [**CloudRouterPostRequestPackage**](CloudRouterPostRequestPackage.md) |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on connection configuration or status changes | [optional] 

## Methods

### NewCloudRouterPostRequest

`func NewCloudRouterPostRequest() *CloudRouterPostRequest`

NewCloudRouterPostRequest instantiates a new CloudRouterPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterPostRequestWithDefaults

`func NewCloudRouterPostRequestWithDefaults() *CloudRouterPostRequest`

NewCloudRouterPostRequestWithDefaults instantiates a new CloudRouterPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CloudRouterPostRequest) GetType() CloudRouterPostRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudRouterPostRequest) GetTypeOk() (*CloudRouterPostRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudRouterPostRequest) SetType(v CloudRouterPostRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *CloudRouterPostRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CloudRouterPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudRouterPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudRouterPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudRouterPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *CloudRouterPostRequest) GetLocation() SimplifiedLocationWithoutIBX`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CloudRouterPostRequest) GetLocationOk() (*SimplifiedLocationWithoutIBX, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CloudRouterPostRequest) SetLocation(v SimplifiedLocationWithoutIBX)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CloudRouterPostRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPackage

`func (o *CloudRouterPostRequest) GetPackage() CloudRouterPostRequestPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *CloudRouterPostRequest) GetPackageOk() (*CloudRouterPostRequestPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *CloudRouterPostRequest) SetPackage(v CloudRouterPostRequestPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *CloudRouterPostRequest) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetOrder

`func (o *CloudRouterPostRequest) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CloudRouterPostRequest) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CloudRouterPostRequest) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CloudRouterPostRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProject

`func (o *CloudRouterPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudRouterPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudRouterPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudRouterPostRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAccount

`func (o *CloudRouterPostRequest) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudRouterPostRequest) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudRouterPostRequest) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudRouterPostRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetNotifications

`func (o *CloudRouterPostRequest) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *CloudRouterPostRequest) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *CloudRouterPostRequest) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *CloudRouterPostRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


