# DeleteCrossConnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**DeinstallScheduleRequest**](DeinstallScheduleRequest.md) | Removal schedule. The requestedDateTime is mandatory. Only the STANDARD schedule type is supported for a removal.  | 
**Notifications** | Pointer to [**[]CrossConnectNotification**](CrossConnectNotification.md) | Optional. The contacts are appended to the ones already held on the Cross Connect.  | [optional] 
**Order** | Pointer to [**OrderRequest**](OrderRequest.md) | Optional. The purchase order and the customer reference identifier replace the ones already held on the Cross Connect.  | [optional] 

## Methods

### NewDeleteCrossConnect

`func NewDeleteCrossConnect(schedule DeinstallScheduleRequest, ) *DeleteCrossConnect`

NewDeleteCrossConnect instantiates a new DeleteCrossConnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCrossConnectWithDefaults

`func NewDeleteCrossConnectWithDefaults() *DeleteCrossConnect`

NewDeleteCrossConnectWithDefaults instantiates a new DeleteCrossConnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *DeleteCrossConnect) GetSchedule() DeinstallScheduleRequest`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *DeleteCrossConnect) GetScheduleOk() (*DeinstallScheduleRequest, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *DeleteCrossConnect) SetSchedule(v DeinstallScheduleRequest)`

SetSchedule sets Schedule field to given value.


### GetNotifications

`func (o *DeleteCrossConnect) GetNotifications() []CrossConnectNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *DeleteCrossConnect) GetNotificationsOk() (*[]CrossConnectNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *DeleteCrossConnect) SetNotifications(v []CrossConnectNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *DeleteCrossConnect) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOrder

`func (o *DeleteCrossConnect) GetOrder() OrderRequest`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *DeleteCrossConnect) GetOrderOk() (*OrderRequest, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *DeleteCrossConnect) SetOrder(v OrderRequest)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *DeleteCrossConnect) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


