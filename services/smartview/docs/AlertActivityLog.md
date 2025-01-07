# AlertActivityLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayLoad** | Pointer to [**[]AlertActivityLogObj**](AlertActivityLogObj.md) |  | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewAlertActivityLog

`func NewAlertActivityLog() *AlertActivityLog`

NewAlertActivityLog instantiates a new AlertActivityLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertActivityLogWithDefaults

`func NewAlertActivityLogWithDefaults() *AlertActivityLog`

NewAlertActivityLogWithDefaults instantiates a new AlertActivityLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayLoad

`func (o *AlertActivityLog) GetPayLoad() []AlertActivityLogObj`

GetPayLoad returns the PayLoad field if non-nil, zero value otherwise.

### GetPayLoadOk

`func (o *AlertActivityLog) GetPayLoadOk() (*[]AlertActivityLogObj, bool)`

GetPayLoadOk returns a tuple with the PayLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayLoad

`func (o *AlertActivityLog) SetPayLoad(v []AlertActivityLogObj)`

SetPayLoad sets PayLoad field to given value.

### HasPayLoad

`func (o *AlertActivityLog) HasPayLoad() bool`

HasPayLoad returns a boolean if a field has been set.

### GetStatus

`func (o *AlertActivityLog) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertActivityLog) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertActivityLog) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertActivityLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


