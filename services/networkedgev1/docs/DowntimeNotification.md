# DowntimeNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | Pointer to **string** | Type of notification, whether planned or unplanned. | [optional] 
**StartTime** | Pointer to **string** | Start of the downtime. | [optional] 
**EndTime** | Pointer to **string** | End of the downtime. | [optional] 
**ImpactedServices** | Pointer to [**[]ImpactedServices**](ImpactedServices.md) | An array of services impacted by the downtime. | [optional] 
**AdditionalMessage** | Pointer to **string** | Any additional messages. | [optional] 

## Methods

### NewDowntimeNotification

`func NewDowntimeNotification() *DowntimeNotification`

NewDowntimeNotification instantiates a new DowntimeNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDowntimeNotificationWithDefaults

`func NewDowntimeNotificationWithDefaults() *DowntimeNotification`

NewDowntimeNotificationWithDefaults instantiates a new DowntimeNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *DowntimeNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *DowntimeNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *DowntimeNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *DowntimeNotification) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.

### GetStartTime

`func (o *DowntimeNotification) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DowntimeNotification) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DowntimeNotification) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DowntimeNotification) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *DowntimeNotification) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DowntimeNotification) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DowntimeNotification) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DowntimeNotification) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetImpactedServices

`func (o *DowntimeNotification) GetImpactedServices() []ImpactedServices`

GetImpactedServices returns the ImpactedServices field if non-nil, zero value otherwise.

### GetImpactedServicesOk

`func (o *DowntimeNotification) GetImpactedServicesOk() (*[]ImpactedServices, bool)`

GetImpactedServicesOk returns a tuple with the ImpactedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedServices

`func (o *DowntimeNotification) SetImpactedServices(v []ImpactedServices)`

SetImpactedServices sets ImpactedServices field to given value.

### HasImpactedServices

`func (o *DowntimeNotification) HasImpactedServices() bool`

HasImpactedServices returns a boolean if a field has been set.

### GetAdditionalMessage

`func (o *DowntimeNotification) GetAdditionalMessage() string`

GetAdditionalMessage returns the AdditionalMessage field if non-nil, zero value otherwise.

### GetAdditionalMessageOk

`func (o *DowntimeNotification) GetAdditionalMessageOk() (*string, bool)`

GetAdditionalMessageOk returns a tuple with the AdditionalMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMessage

`func (o *DowntimeNotification) SetAdditionalMessage(v string)`

SetAdditionalMessage sets AdditionalMessage field to given value.

### HasAdditionalMessage

`func (o *DowntimeNotification) HasAdditionalMessage() bool`

HasAdditionalMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


