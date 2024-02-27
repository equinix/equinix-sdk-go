# Statistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDateTime** | Pointer to **time.Time** | Start and duration of the statistical analysis interval. | [optional] 
**EndDateTime** | Pointer to **time.Time** | End and duration of the statistical analysis interval. | [optional] 
**ViewPoint** | Pointer to [**StatisticsViewPoint**](StatisticsViewPoint.md) |  | [optional] 
**BandwidthUtilization** | Pointer to [**BandwidthUtilization**](BandwidthUtilization.md) |  | [optional] 

## Methods

### NewStatistics

`func NewStatistics() *Statistics`

NewStatistics instantiates a new Statistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsWithDefaults

`func NewStatisticsWithDefaults() *Statistics`

NewStatisticsWithDefaults instantiates a new Statistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDateTime

`func (o *Statistics) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *Statistics) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *Statistics) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *Statistics) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *Statistics) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *Statistics) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *Statistics) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *Statistics) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetViewPoint

`func (o *Statistics) GetViewPoint() StatisticsViewPoint`

GetViewPoint returns the ViewPoint field if non-nil, zero value otherwise.

### GetViewPointOk

`func (o *Statistics) GetViewPointOk() (*StatisticsViewPoint, bool)`

GetViewPointOk returns a tuple with the ViewPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPoint

`func (o *Statistics) SetViewPoint(v StatisticsViewPoint)`

SetViewPoint sets ViewPoint field to given value.

### HasViewPoint

`func (o *Statistics) HasViewPoint() bool`

HasViewPoint returns a boolean if a field has been set.

### GetBandwidthUtilization

`func (o *Statistics) GetBandwidthUtilization() BandwidthUtilization`

GetBandwidthUtilization returns the BandwidthUtilization field if non-nil, zero value otherwise.

### GetBandwidthUtilizationOk

`func (o *Statistics) GetBandwidthUtilizationOk() (*BandwidthUtilization, bool)`

GetBandwidthUtilizationOk returns a tuple with the BandwidthUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthUtilization

`func (o *Statistics) SetBandwidthUtilization(v BandwidthUtilization)`

SetBandwidthUtilization sets BandwidthUtilization field to given value.

### HasBandwidthUtilization

`func (o *Statistics) HasBandwidthUtilization() bool`

HasBandwidthUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


