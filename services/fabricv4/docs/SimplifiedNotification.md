# SimplifiedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SimplifiedNotificationType**](SimplifiedNotificationType.md) |  | 
**SendInterval** | Pointer to **string** |  | [optional] 
**Emails** | **[]string** | Array of contact emails | 
**RegisteredUsers** | Pointer to **[]string** | Array of registered users | [optional] 

## Methods

### NewSimplifiedNotification

`func NewSimplifiedNotification(type_ SimplifiedNotificationType, emails []string, ) *SimplifiedNotification`

NewSimplifiedNotification instantiates a new SimplifiedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedNotificationWithDefaults

`func NewSimplifiedNotificationWithDefaults() *SimplifiedNotification`

NewSimplifiedNotificationWithDefaults instantiates a new SimplifiedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SimplifiedNotification) GetType() SimplifiedNotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedNotification) GetTypeOk() (*SimplifiedNotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedNotification) SetType(v SimplifiedNotificationType)`

SetType sets Type field to given value.


### GetSendInterval

`func (o *SimplifiedNotification) GetSendInterval() string`

GetSendInterval returns the SendInterval field if non-nil, zero value otherwise.

### GetSendIntervalOk

`func (o *SimplifiedNotification) GetSendIntervalOk() (*string, bool)`

GetSendIntervalOk returns a tuple with the SendInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInterval

`func (o *SimplifiedNotification) SetSendInterval(v string)`

SetSendInterval sets SendInterval field to given value.

### HasSendInterval

`func (o *SimplifiedNotification) HasSendInterval() bool`

HasSendInterval returns a boolean if a field has been set.

### GetEmails

`func (o *SimplifiedNotification) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *SimplifiedNotification) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *SimplifiedNotification) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetRegisteredUsers

`func (o *SimplifiedNotification) GetRegisteredUsers() []string`

GetRegisteredUsers returns the RegisteredUsers field if non-nil, zero value otherwise.

### GetRegisteredUsersOk

`func (o *SimplifiedNotification) GetRegisteredUsersOk() (*[]string, bool)`

GetRegisteredUsersOk returns a tuple with the RegisteredUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredUsers

`func (o *SimplifiedNotification) SetRegisteredUsers(v []string)`

SetRegisteredUsers sets RegisteredUsers field to given value.

### HasRegisteredUsers

`func (o *SimplifiedNotification) HasRegisteredUsers() bool`

HasRegisteredUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


