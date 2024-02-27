# PortNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PortNotificationType**](PortNotificationType.md) |  | 
**RegisteredUsers** | **[]string** | Array of registered users | 

## Methods

### NewPortNotification

`func NewPortNotification(type_ PortNotificationType, registeredUsers []string, ) *PortNotification`

NewPortNotification instantiates a new PortNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortNotificationWithDefaults

`func NewPortNotificationWithDefaults() *PortNotification`

NewPortNotificationWithDefaults instantiates a new PortNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PortNotification) GetType() PortNotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortNotification) GetTypeOk() (*PortNotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortNotification) SetType(v PortNotificationType)`

SetType sets Type field to given value.


### GetRegisteredUsers

`func (o *PortNotification) GetRegisteredUsers() []string`

GetRegisteredUsers returns the RegisteredUsers field if non-nil, zero value otherwise.

### GetRegisteredUsersOk

`func (o *PortNotification) GetRegisteredUsersOk() (*[]string, bool)`

GetRegisteredUsersOk returns a tuple with the RegisteredUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredUsers

`func (o *PortNotification) SetRegisteredUsers(v []string)`

SetRegisteredUsers sets RegisteredUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


