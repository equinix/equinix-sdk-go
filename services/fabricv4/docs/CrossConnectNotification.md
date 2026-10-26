# CrossConnectNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**NotificationType**](NotificationType.md) |  | 
**FirstName** | Pointer to **string** | First name of the unregistered contact. Applicable only to the TECHNICAL_UNREGISTERED type.  | [optional] 
**LastName** | Pointer to **string** | Last name of the unregistered contact. Applicable only to the TECHNICAL_UNREGISTERED type.  | [optional] 
**Emails** | **[]string** | Email addresses for notifications. | 
**RegisteredUsers** | Pointer to **[]string** | Registered users for notifications. | [optional] 
**Phone** | Pointer to [**[]CrossConnectPhone**](CrossConnectPhone.md) | Contact phone numbers. | [optional] 

## Methods

### NewCrossConnectNotification

`func NewCrossConnectNotification(type_ NotificationType, emails []string, ) *CrossConnectNotification`

NewCrossConnectNotification instantiates a new CrossConnectNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectNotificationWithDefaults

`func NewCrossConnectNotificationWithDefaults() *CrossConnectNotification`

NewCrossConnectNotificationWithDefaults instantiates a new CrossConnectNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CrossConnectNotification) GetType() NotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossConnectNotification) GetTypeOk() (*NotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossConnectNotification) SetType(v NotificationType)`

SetType sets Type field to given value.


### GetFirstName

`func (o *CrossConnectNotification) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CrossConnectNotification) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CrossConnectNotification) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CrossConnectNotification) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CrossConnectNotification) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CrossConnectNotification) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CrossConnectNotification) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CrossConnectNotification) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmails

`func (o *CrossConnectNotification) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *CrossConnectNotification) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *CrossConnectNotification) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetRegisteredUsers

`func (o *CrossConnectNotification) GetRegisteredUsers() []string`

GetRegisteredUsers returns the RegisteredUsers field if non-nil, zero value otherwise.

### GetRegisteredUsersOk

`func (o *CrossConnectNotification) GetRegisteredUsersOk() (*[]string, bool)`

GetRegisteredUsersOk returns a tuple with the RegisteredUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredUsers

`func (o *CrossConnectNotification) SetRegisteredUsers(v []string)`

SetRegisteredUsers sets RegisteredUsers field to given value.

### HasRegisteredUsers

`func (o *CrossConnectNotification) HasRegisteredUsers() bool`

HasRegisteredUsers returns a boolean if a field has been set.

### GetPhone

`func (o *CrossConnectNotification) GetPhone() []CrossConnectPhone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CrossConnectNotification) GetPhoneOk() (*[]CrossConnectPhone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CrossConnectNotification) SetPhone(v []CrossConnectPhone)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CrossConnectNotification) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


