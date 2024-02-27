# Direction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Max** | Pointer to **float32** | Max bandwidth within request time range, represented in units specified by response \&quot;units\&quot; field | [optional] 
**Mean** | Pointer to **float32** | Mean bandwidth within request time range, represented in units specified by response \&quot;units\&quot; field | [optional] 
**Metrics** | Pointer to [**[]Metrics**](Metrics.md) | Bandwidth utilization statistics for a specified interval. | [optional] 

## Methods

### NewDirection

`func NewDirection() *Direction`

NewDirection instantiates a new Direction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectionWithDefaults

`func NewDirectionWithDefaults() *Direction`

NewDirectionWithDefaults instantiates a new Direction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMax

`func (o *Direction) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *Direction) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *Direction) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *Direction) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMean

`func (o *Direction) GetMean() float32`

GetMean returns the Mean field if non-nil, zero value otherwise.

### GetMeanOk

`func (o *Direction) GetMeanOk() (*float32, bool)`

GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMean

`func (o *Direction) SetMean(v float32)`

SetMean sets Mean field to given value.

### HasMean

`func (o *Direction) HasMean() bool`

HasMean returns a boolean if a field has been set.

### GetMetrics

`func (o *Direction) GetMetrics() []Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Direction) GetMetricsOk() (*[]Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Direction) SetMetrics(v []Metrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Direction) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


