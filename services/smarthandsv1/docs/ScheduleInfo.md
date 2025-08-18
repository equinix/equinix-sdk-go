# ScheduleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleType** | [**ScheduleInfoScheduleType**](ScheduleInfoScheduleType.md) |  | 
**RequestedStartDate** | Pointer to **time.Time** | Requested Start Date Time (ISO Date)&lt;br&gt; Cannot Past Date. | [optional] 
**RequestedCompletionDate** | Pointer to **time.Time** | Requested Completion Date Time (ISO Date)&lt;br&gt;Cannot Past Date. | [optional] 

## Methods

### NewScheduleInfo

`func NewScheduleInfo(scheduleType ScheduleInfoScheduleType, ) *ScheduleInfo`

NewScheduleInfo instantiates a new ScheduleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleInfoWithDefaults

`func NewScheduleInfoWithDefaults() *ScheduleInfo`

NewScheduleInfoWithDefaults instantiates a new ScheduleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleType

`func (o *ScheduleInfo) GetScheduleType() ScheduleInfoScheduleType`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *ScheduleInfo) GetScheduleTypeOk() (*ScheduleInfoScheduleType, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *ScheduleInfo) SetScheduleType(v ScheduleInfoScheduleType)`

SetScheduleType sets ScheduleType field to given value.


### GetRequestedStartDate

`func (o *ScheduleInfo) GetRequestedStartDate() time.Time`

GetRequestedStartDate returns the RequestedStartDate field if non-nil, zero value otherwise.

### GetRequestedStartDateOk

`func (o *ScheduleInfo) GetRequestedStartDateOk() (*time.Time, bool)`

GetRequestedStartDateOk returns a tuple with the RequestedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedStartDate

`func (o *ScheduleInfo) SetRequestedStartDate(v time.Time)`

SetRequestedStartDate sets RequestedStartDate field to given value.

### HasRequestedStartDate

`func (o *ScheduleInfo) HasRequestedStartDate() bool`

HasRequestedStartDate returns a boolean if a field has been set.

### GetRequestedCompletionDate

`func (o *ScheduleInfo) GetRequestedCompletionDate() time.Time`

GetRequestedCompletionDate returns the RequestedCompletionDate field if non-nil, zero value otherwise.

### GetRequestedCompletionDateOk

`func (o *ScheduleInfo) GetRequestedCompletionDateOk() (*time.Time, bool)`

GetRequestedCompletionDateOk returns a tuple with the RequestedCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCompletionDate

`func (o *ScheduleInfo) SetRequestedCompletionDate(v time.Time)`

SetRequestedCompletionDate sets RequestedCompletionDate field to given value.

### HasRequestedCompletionDate

`func (o *ScheduleInfo) HasRequestedCompletionDate() bool`

HasRequestedCompletionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


