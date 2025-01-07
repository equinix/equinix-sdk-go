# AlarmMessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Alarm**](Alarm.md) |  | 
**Type** | **string** | message type | [default to "system-alert"]

## Methods

### NewAlarmMessageData

`func NewAlarmMessageData(data Alarm, type_ string, ) *AlarmMessageData`

NewAlarmMessageData instantiates a new AlarmMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmMessageDataWithDefaults

`func NewAlarmMessageDataWithDefaults() *AlarmMessageData`

NewAlarmMessageDataWithDefaults instantiates a new AlarmMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AlarmMessageData) GetData() Alarm`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlarmMessageData) GetDataOk() (*Alarm, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlarmMessageData) SetData(v Alarm)`

SetData sets Data field to given value.


### GetType

`func (o *AlarmMessageData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlarmMessageData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlarmMessageData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


