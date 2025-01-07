# AlarmThresholdDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | threshold message | [optional] 
**StateLimit** | Pointer to **string** | threshold state limit | [optional] 
**Unit** | Pointer to **string** | threshold unit | [optional] 

## Methods

### NewAlarmThresholdDetails

`func NewAlarmThresholdDetails() *AlarmThresholdDetails`

NewAlarmThresholdDetails instantiates a new AlarmThresholdDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmThresholdDetailsWithDefaults

`func NewAlarmThresholdDetailsWithDefaults() *AlarmThresholdDetails`

NewAlarmThresholdDetailsWithDefaults instantiates a new AlarmThresholdDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AlarmThresholdDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlarmThresholdDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlarmThresholdDetails) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AlarmThresholdDetails) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStateLimit

`func (o *AlarmThresholdDetails) GetStateLimit() string`

GetStateLimit returns the StateLimit field if non-nil, zero value otherwise.

### GetStateLimitOk

`func (o *AlarmThresholdDetails) GetStateLimitOk() (*string, bool)`

GetStateLimitOk returns a tuple with the StateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateLimit

`func (o *AlarmThresholdDetails) SetStateLimit(v string)`

SetStateLimit sets StateLimit field to given value.

### HasStateLimit

`func (o *AlarmThresholdDetails) HasStateLimit() bool`

HasStateLimit returns a boolean if a field has been set.

### GetUnit

`func (o *AlarmThresholdDetails) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AlarmThresholdDetails) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AlarmThresholdDetails) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *AlarmThresholdDetails) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


