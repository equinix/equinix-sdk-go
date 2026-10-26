# StreamPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**StreamPostRequestType**](StreamPostRequestType.md) |  | 
**Name** | **string** | Customer-provided stream name | 
**Description** | Pointer to **string** | Customer-provided stream description | [optional] 
**Project** | [**Project**](Project.md) |  | 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for stream notifications. Optional; at most one entry, of type ALL, with up to 12 emails. Set at creation only. | [optional] 

## Methods

### NewStreamPostRequest

`func NewStreamPostRequest(type_ StreamPostRequestType, name string, project Project, ) *StreamPostRequest`

NewStreamPostRequest instantiates a new StreamPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamPostRequestWithDefaults

`func NewStreamPostRequestWithDefaults() *StreamPostRequest`

NewStreamPostRequestWithDefaults instantiates a new StreamPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StreamPostRequest) GetType() StreamPostRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamPostRequest) GetTypeOk() (*StreamPostRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamPostRequest) SetType(v StreamPostRequestType)`

SetType sets Type field to given value.


### GetName

`func (o *StreamPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *StreamPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StreamPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StreamPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StreamPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProject

`func (o *StreamPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *StreamPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *StreamPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.


### GetAccount

`func (o *StreamPostRequest) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *StreamPostRequest) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *StreamPostRequest) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *StreamPostRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetNotifications

`func (o *StreamPostRequest) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *StreamPostRequest) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *StreamPostRequest) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *StreamPostRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


