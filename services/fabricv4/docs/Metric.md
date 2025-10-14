# Metric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Equinix supported metric type | [optional] 
**Name** | Pointer to **string** | Metric name | [optional] 
**Unit** | Pointer to **string** | Metric unit | [optional] 
**Interval** | Pointer to **string** | Metric interval (set automatically based on search range) | [optional] 
**Resource** | Pointer to [**MetricResource**](MetricResource.md) |  | [optional] 
**Summary** | Pointer to **string** | Metric summary | [optional] 
**Datapoints** | Pointer to [**[]MetricDatapoints**](MetricDatapoints.md) | Metric data points | [optional] 

## Methods

### NewMetric

`func NewMetric() *Metric`

NewMetric instantiates a new Metric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricWithDefaults

`func NewMetricWithDefaults() *Metric`

NewMetricWithDefaults instantiates a new Metric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Metric) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Metric) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Metric) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Metric) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Metric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Metric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Metric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Metric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnit

`func (o *Metric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Metric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Metric) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Metric) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetInterval

`func (o *Metric) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Metric) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Metric) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Metric) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetResource

`func (o *Metric) GetResource() MetricResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Metric) GetResourceOk() (*MetricResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Metric) SetResource(v MetricResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Metric) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSummary

`func (o *Metric) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Metric) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Metric) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Metric) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDatapoints

`func (o *Metric) GetDatapoints() []MetricDatapoints`

GetDatapoints returns the Datapoints field if non-nil, zero value otherwise.

### GetDatapointsOk

`func (o *Metric) GetDatapointsOk() (*[]MetricDatapoints, bool)`

GetDatapointsOk returns a tuple with the Datapoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatapoints

`func (o *Metric) SetDatapoints(v []MetricDatapoints)`

SetDatapoints sets Datapoints field to given value.

### HasDatapoints

`func (o *Metric) HasDatapoints() bool`

HasDatapoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


