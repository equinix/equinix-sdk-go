# OpticalConnectNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**OpticalConnectNotificationType**](OpticalConnectNotificationType.md) |  | 
**Emails** | **[]string** | Email addresses to notify. | 
**RegisteredUsers** | Pointer to **[]string** | Usernames of registered Equinix Portal users to notify. | [optional] 

## Methods

### NewOpticalConnectNotification

`func NewOpticalConnectNotification(type_ OpticalConnectNotificationType, emails []string, ) *OpticalConnectNotification`

NewOpticalConnectNotification instantiates a new OpticalConnectNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpticalConnectNotificationWithDefaults

`func NewOpticalConnectNotificationWithDefaults() *OpticalConnectNotification`

NewOpticalConnectNotificationWithDefaults instantiates a new OpticalConnectNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OpticalConnectNotification) GetType() OpticalConnectNotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpticalConnectNotification) GetTypeOk() (*OpticalConnectNotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpticalConnectNotification) SetType(v OpticalConnectNotificationType)`

SetType sets Type field to given value.


### GetEmails

`func (o *OpticalConnectNotification) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *OpticalConnectNotification) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *OpticalConnectNotification) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetRegisteredUsers

`func (o *OpticalConnectNotification) GetRegisteredUsers() []string`

GetRegisteredUsers returns the RegisteredUsers field if non-nil, zero value otherwise.

### GetRegisteredUsersOk

`func (o *OpticalConnectNotification) GetRegisteredUsersOk() (*[]string, bool)`

GetRegisteredUsersOk returns a tuple with the RegisteredUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredUsers

`func (o *OpticalConnectNotification) SetRegisteredUsers(v []string)`

SetRegisteredUsers sets RegisteredUsers field to given value.

### HasRegisteredUsers

`func (o *OpticalConnectNotification) HasRegisteredUsers() bool`

HasRegisteredUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


