# AlarmStatusDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledged** | Pointer to **bool** | status acknowledged | [optional] 
**AcknowledgementTime** | Pointer to **string** | status acknowledgement time | [optional] 
**Active** | Pointer to **bool** | status active | [optional] 
**Cleared** | Pointer to **bool** | status cleared | [optional] 

## Methods

### NewAlarmStatusDetails

`func NewAlarmStatusDetails() *AlarmStatusDetails`

NewAlarmStatusDetails instantiates a new AlarmStatusDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmStatusDetailsWithDefaults

`func NewAlarmStatusDetailsWithDefaults() *AlarmStatusDetails`

NewAlarmStatusDetailsWithDefaults instantiates a new AlarmStatusDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledged

`func (o *AlarmStatusDetails) GetAcknowledged() bool`

GetAcknowledged returns the Acknowledged field if non-nil, zero value otherwise.

### GetAcknowledgedOk

`func (o *AlarmStatusDetails) GetAcknowledgedOk() (*bool, bool)`

GetAcknowledgedOk returns a tuple with the Acknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledged

`func (o *AlarmStatusDetails) SetAcknowledged(v bool)`

SetAcknowledged sets Acknowledged field to given value.

### HasAcknowledged

`func (o *AlarmStatusDetails) HasAcknowledged() bool`

HasAcknowledged returns a boolean if a field has been set.

### GetAcknowledgementTime

`func (o *AlarmStatusDetails) GetAcknowledgementTime() string`

GetAcknowledgementTime returns the AcknowledgementTime field if non-nil, zero value otherwise.

### GetAcknowledgementTimeOk

`func (o *AlarmStatusDetails) GetAcknowledgementTimeOk() (*string, bool)`

GetAcknowledgementTimeOk returns a tuple with the AcknowledgementTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgementTime

`func (o *AlarmStatusDetails) SetAcknowledgementTime(v string)`

SetAcknowledgementTime sets AcknowledgementTime field to given value.

### HasAcknowledgementTime

`func (o *AlarmStatusDetails) HasAcknowledgementTime() bool`

HasAcknowledgementTime returns a boolean if a field has been set.

### GetActive

`func (o *AlarmStatusDetails) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AlarmStatusDetails) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AlarmStatusDetails) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AlarmStatusDetails) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCleared

`func (o *AlarmStatusDetails) GetCleared() bool`

GetCleared returns the Cleared field if non-nil, zero value otherwise.

### GetClearedOk

`func (o *AlarmStatusDetails) GetClearedOk() (*bool, bool)`

GetClearedOk returns a tuple with the Cleared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleared

`func (o *AlarmStatusDetails) SetCleared(v bool)`

SetCleared sets Cleared field to given value.

### HasCleared

`func (o *AlarmStatusDetails) HasCleared() bool`

HasCleared returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


