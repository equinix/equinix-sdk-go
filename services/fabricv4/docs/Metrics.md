# Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntervalEndTimestamp** | Pointer to **time.Time** | Interval end timestamp | [optional] 
**Max** | Pointer to **float32** | Max bandwidth within statistics object time interval, represented in units specified by response \&quot;units\&quot; field | [optional] 
**Mean** | Pointer to **float32** | Mean bandwidth within statistics object time interval, represented in units specified by response \&quot;units\&quot; field | [optional] 

## Methods

### NewMetrics

`func NewMetrics() *Metrics`

NewMetrics instantiates a new Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsWithDefaults

`func NewMetricsWithDefaults() *Metrics`

NewMetricsWithDefaults instantiates a new Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervalEndTimestamp

`func (o *Metrics) GetIntervalEndTimestamp() time.Time`

GetIntervalEndTimestamp returns the IntervalEndTimestamp field if non-nil, zero value otherwise.

### GetIntervalEndTimestampOk

`func (o *Metrics) GetIntervalEndTimestampOk() (*time.Time, bool)`

GetIntervalEndTimestampOk returns a tuple with the IntervalEndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalEndTimestamp

`func (o *Metrics) SetIntervalEndTimestamp(v time.Time)`

SetIntervalEndTimestamp sets IntervalEndTimestamp field to given value.

### HasIntervalEndTimestamp

`func (o *Metrics) HasIntervalEndTimestamp() bool`

HasIntervalEndTimestamp returns a boolean if a field has been set.

### GetMax

`func (o *Metrics) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *Metrics) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *Metrics) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *Metrics) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMean

`func (o *Metrics) GetMean() float32`

GetMean returns the Mean field if non-nil, zero value otherwise.

### GetMeanOk

`func (o *Metrics) GetMeanOk() (*float32, bool)`

GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMean

`func (o *Metrics) SetMean(v float32)`

SetMean sets Mean field to given value.

### HasMean

`func (o *Metrics) HasMean() bool`

HasMean returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


