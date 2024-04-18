# PolledThroughputMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntervalDateTime** | Pointer to **string** | The end of a polled time period. | [optional] 
**Mean** | Pointer to **float32** | Mean traffic throughput. | [optional] 

## Methods

### NewPolledThroughputMetrics

`func NewPolledThroughputMetrics() *PolledThroughputMetrics`

NewPolledThroughputMetrics instantiates a new PolledThroughputMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolledThroughputMetricsWithDefaults

`func NewPolledThroughputMetricsWithDefaults() *PolledThroughputMetrics`

NewPolledThroughputMetricsWithDefaults instantiates a new PolledThroughputMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervalDateTime

`func (o *PolledThroughputMetrics) GetIntervalDateTime() string`

GetIntervalDateTime returns the IntervalDateTime field if non-nil, zero value otherwise.

### GetIntervalDateTimeOk

`func (o *PolledThroughputMetrics) GetIntervalDateTimeOk() (*string, bool)`

GetIntervalDateTimeOk returns a tuple with the IntervalDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalDateTime

`func (o *PolledThroughputMetrics) SetIntervalDateTime(v string)`

SetIntervalDateTime sets IntervalDateTime field to given value.

### HasIntervalDateTime

`func (o *PolledThroughputMetrics) HasIntervalDateTime() bool`

HasIntervalDateTime returns a boolean if a field has been set.

### GetMean

`func (o *PolledThroughputMetrics) GetMean() float32`

GetMean returns the Mean field if non-nil, zero value otherwise.

### GetMeanOk

`func (o *PolledThroughputMetrics) GetMeanOk() (*float32, bool)`

GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMean

`func (o *PolledThroughputMetrics) SetMean(v float32)`

SetMean sets Mean field to given value.

### HasMean

`func (o *PolledThroughputMetrics) HasMean() bool`

HasMean returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


