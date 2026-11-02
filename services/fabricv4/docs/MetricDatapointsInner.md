# MetricDatapointsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDateTime** | Pointer to **time.Time** | Datapoint end date and time | [optional] 
**StartDateTime** | Pointer to **time.Time** | Datapoint start date and time | [optional] 
**Value** | Pointer to **float32** | Datapoint value | [optional] 

## Methods

### NewMetricDatapointsInner

`func NewMetricDatapointsInner() *MetricDatapointsInner`

NewMetricDatapointsInner instantiates a new MetricDatapointsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDatapointsInnerWithDefaults

`func NewMetricDatapointsInnerWithDefaults() *MetricDatapointsInner`

NewMetricDatapointsInnerWithDefaults instantiates a new MetricDatapointsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDateTime

`func (o *MetricDatapointsInner) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MetricDatapointsInner) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MetricDatapointsInner) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MetricDatapointsInner) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetStartDateTime

`func (o *MetricDatapointsInner) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MetricDatapointsInner) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MetricDatapointsInner) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MetricDatapointsInner) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetValue

`func (o *MetricDatapointsInner) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricDatapointsInner) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricDatapointsInner) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *MetricDatapointsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


