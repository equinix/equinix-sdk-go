# ExchangeServiceNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ExchangeServiceNotificationType**](ExchangeServiceNotificationType.md) |  | 
**RegisteredUsers** | **[]string** | Array of registered users | 

## Methods

### NewExchangeServiceNotification

`func NewExchangeServiceNotification(type_ ExchangeServiceNotificationType, registeredUsers []string, ) *ExchangeServiceNotification`

NewExchangeServiceNotification instantiates a new ExchangeServiceNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeServiceNotificationWithDefaults

`func NewExchangeServiceNotificationWithDefaults() *ExchangeServiceNotification`

NewExchangeServiceNotificationWithDefaults instantiates a new ExchangeServiceNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExchangeServiceNotification) GetType() ExchangeServiceNotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExchangeServiceNotification) GetTypeOk() (*ExchangeServiceNotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExchangeServiceNotification) SetType(v ExchangeServiceNotificationType)`

SetType sets Type field to given value.


### GetRegisteredUsers

`func (o *ExchangeServiceNotification) GetRegisteredUsers() []string`

GetRegisteredUsers returns the RegisteredUsers field if non-nil, zero value otherwise.

### GetRegisteredUsersOk

`func (o *ExchangeServiceNotification) GetRegisteredUsersOk() (*[]string, bool)`

GetRegisteredUsersOk returns a tuple with the RegisteredUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredUsers

`func (o *ExchangeServiceNotification) SetRegisteredUsers(v []string)`

SetRegisteredUsers sets RegisteredUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


