# InterfaceStatsofTraffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Max** | Pointer to **float32** | Max throughput during the time interval. | [optional] 
**Mean** | Pointer to **float32** | Mean throughput during the time interval. | [optional] 
**LastPolled** | Pointer to **float32** | The throughput of the last polled data. | [optional] 
**Metrics** | Pointer to [**[]PolledThroughputMetrics**](PolledThroughputMetrics.md) |  | [optional] 

## Methods

### NewInterfaceStatsofTraffic

`func NewInterfaceStatsofTraffic() *InterfaceStatsofTraffic`

NewInterfaceStatsofTraffic instantiates a new InterfaceStatsofTraffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceStatsofTrafficWithDefaults

`func NewInterfaceStatsofTrafficWithDefaults() *InterfaceStatsofTraffic`

NewInterfaceStatsofTrafficWithDefaults instantiates a new InterfaceStatsofTraffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMax

`func (o *InterfaceStatsofTraffic) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *InterfaceStatsofTraffic) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *InterfaceStatsofTraffic) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *InterfaceStatsofTraffic) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMean

`func (o *InterfaceStatsofTraffic) GetMean() float32`

GetMean returns the Mean field if non-nil, zero value otherwise.

### GetMeanOk

`func (o *InterfaceStatsofTraffic) GetMeanOk() (*float32, bool)`

GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMean

`func (o *InterfaceStatsofTraffic) SetMean(v float32)`

SetMean sets Mean field to given value.

### HasMean

`func (o *InterfaceStatsofTraffic) HasMean() bool`

HasMean returns a boolean if a field has been set.

### GetLastPolled

`func (o *InterfaceStatsofTraffic) GetLastPolled() float32`

GetLastPolled returns the LastPolled field if non-nil, zero value otherwise.

### GetLastPolledOk

`func (o *InterfaceStatsofTraffic) GetLastPolledOk() (*float32, bool)`

GetLastPolledOk returns a tuple with the LastPolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPolled

`func (o *InterfaceStatsofTraffic) SetLastPolled(v float32)`

SetLastPolled sets LastPolled field to given value.

### HasLastPolled

`func (o *InterfaceStatsofTraffic) HasLastPolled() bool`

HasLastPolled returns a boolean if a field has been set.

### GetMetrics

`func (o *InterfaceStatsofTraffic) GetMetrics() []PolledThroughputMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *InterfaceStatsofTraffic) GetMetricsOk() (*[]PolledThroughputMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *InterfaceStatsofTraffic) SetMetrics(v []PolledThroughputMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *InterfaceStatsofTraffic) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


