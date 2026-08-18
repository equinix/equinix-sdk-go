# InterconnectNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**InterconnectNotificationType**](InterconnectNotificationType.md) |  | [optional] 
**Emails** | Pointer to **[]string** | Array of contact emails | [optional] 

## Methods

### NewInterconnectNotification

`func NewInterconnectNotification() *InterconnectNotification`

NewInterconnectNotification instantiates a new InterconnectNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectNotificationWithDefaults

`func NewInterconnectNotificationWithDefaults() *InterconnectNotification`

NewInterconnectNotificationWithDefaults instantiates a new InterconnectNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InterconnectNotification) GetType() InterconnectNotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterconnectNotification) GetTypeOk() (*InterconnectNotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterconnectNotification) SetType(v InterconnectNotificationType)`

SetType sets Type field to given value.

### HasType

`func (o *InterconnectNotification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmails

`func (o *InterconnectNotification) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *InterconnectNotification) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *InterconnectNotification) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *InterconnectNotification) HasEmails() bool`

HasEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


