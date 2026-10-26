# ScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ScheduleRequestType**](ScheduleRequestType.md) |  | [optional] 
**Code** | [**ScheduleType**](ScheduleType.md) |  | 
**AutoConvertToStandard** | Pointer to **bool** | Applicable only when the code is EXPEDITE. When true, Equinix schedules the Cross Connect as STANDARD when no EXPEDITE scheduling is available.  | [optional] 
**ReadinessDateTime** | Pointer to **time.Time** | Carrier Circuit delivery date is the date when the Carrier commits to enable the circuit. Your order will be placed on hold, pending the date you enter. Equinix will wait for the Carrier to deliver the service and then will complete the installation.  | [optional] 
**RequestedDateTime** | Pointer to **time.Time** | Requested Installation Date and Time. | [optional] 

## Methods

### NewScheduleRequest

`func NewScheduleRequest(code ScheduleType, ) *ScheduleRequest`

NewScheduleRequest instantiates a new ScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleRequestWithDefaults

`func NewScheduleRequestWithDefaults() *ScheduleRequest`

NewScheduleRequestWithDefaults instantiates a new ScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScheduleRequest) GetType() ScheduleRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduleRequest) GetTypeOk() (*ScheduleRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduleRequest) SetType(v ScheduleRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *ScheduleRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *ScheduleRequest) GetCode() ScheduleType`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ScheduleRequest) GetCodeOk() (*ScheduleType, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ScheduleRequest) SetCode(v ScheduleType)`

SetCode sets Code field to given value.


### GetAutoConvertToStandard

`func (o *ScheduleRequest) GetAutoConvertToStandard() bool`

GetAutoConvertToStandard returns the AutoConvertToStandard field if non-nil, zero value otherwise.

### GetAutoConvertToStandardOk

`func (o *ScheduleRequest) GetAutoConvertToStandardOk() (*bool, bool)`

GetAutoConvertToStandardOk returns a tuple with the AutoConvertToStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConvertToStandard

`func (o *ScheduleRequest) SetAutoConvertToStandard(v bool)`

SetAutoConvertToStandard sets AutoConvertToStandard field to given value.

### HasAutoConvertToStandard

`func (o *ScheduleRequest) HasAutoConvertToStandard() bool`

HasAutoConvertToStandard returns a boolean if a field has been set.

### GetReadinessDateTime

`func (o *ScheduleRequest) GetReadinessDateTime() time.Time`

GetReadinessDateTime returns the ReadinessDateTime field if non-nil, zero value otherwise.

### GetReadinessDateTimeOk

`func (o *ScheduleRequest) GetReadinessDateTimeOk() (*time.Time, bool)`

GetReadinessDateTimeOk returns a tuple with the ReadinessDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessDateTime

`func (o *ScheduleRequest) SetReadinessDateTime(v time.Time)`

SetReadinessDateTime sets ReadinessDateTime field to given value.

### HasReadinessDateTime

`func (o *ScheduleRequest) HasReadinessDateTime() bool`

HasReadinessDateTime returns a boolean if a field has been set.

### GetRequestedDateTime

`func (o *ScheduleRequest) GetRequestedDateTime() time.Time`

GetRequestedDateTime returns the RequestedDateTime field if non-nil, zero value otherwise.

### GetRequestedDateTimeOk

`func (o *ScheduleRequest) GetRequestedDateTimeOk() (*time.Time, bool)`

GetRequestedDateTimeOk returns a tuple with the RequestedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDateTime

`func (o *ScheduleRequest) SetRequestedDateTime(v time.Time)`

SetRequestedDateTime sets RequestedDateTime field to given value.

### HasRequestedDateTime

`func (o *ScheduleRequest) HasRequestedDateTime() bool`

HasRequestedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


