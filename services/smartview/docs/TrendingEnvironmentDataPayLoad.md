# TrendingEnvironmentDataPayLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | Pointer to **string** | customer account num | [optional] 
**Datapoint** | Pointer to **string** | data point for which the timeseries data is fetched | [optional] 
**Ibx** | Pointer to **string** | ibx code | [optional] 
**Interval** | Pointer to **string** | interval | [optional] 
**Series** | Pointer to [**[]DataValue**](DataValue.md) | Time series data for the data point | [optional] 
**Uom** | Pointer to **string** | unit of measure used for the datapoint | [optional] 

## Methods

### NewTrendingEnvironmentDataPayLoad

`func NewTrendingEnvironmentDataPayLoad() *TrendingEnvironmentDataPayLoad`

NewTrendingEnvironmentDataPayLoad instantiates a new TrendingEnvironmentDataPayLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendingEnvironmentDataPayLoadWithDefaults

`func NewTrendingEnvironmentDataPayLoadWithDefaults() *TrendingEnvironmentDataPayLoad`

NewTrendingEnvironmentDataPayLoadWithDefaults instantiates a new TrendingEnvironmentDataPayLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *TrendingEnvironmentDataPayLoad) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *TrendingEnvironmentDataPayLoad) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *TrendingEnvironmentDataPayLoad) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.

### HasAccountNo

`func (o *TrendingEnvironmentDataPayLoad) HasAccountNo() bool`

HasAccountNo returns a boolean if a field has been set.

### GetDatapoint

`func (o *TrendingEnvironmentDataPayLoad) GetDatapoint() string`

GetDatapoint returns the Datapoint field if non-nil, zero value otherwise.

### GetDatapointOk

`func (o *TrendingEnvironmentDataPayLoad) GetDatapointOk() (*string, bool)`

GetDatapointOk returns a tuple with the Datapoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatapoint

`func (o *TrendingEnvironmentDataPayLoad) SetDatapoint(v string)`

SetDatapoint sets Datapoint field to given value.

### HasDatapoint

`func (o *TrendingEnvironmentDataPayLoad) HasDatapoint() bool`

HasDatapoint returns a boolean if a field has been set.

### GetIbx

`func (o *TrendingEnvironmentDataPayLoad) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *TrendingEnvironmentDataPayLoad) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *TrendingEnvironmentDataPayLoad) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *TrendingEnvironmentDataPayLoad) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetInterval

`func (o *TrendingEnvironmentDataPayLoad) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *TrendingEnvironmentDataPayLoad) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *TrendingEnvironmentDataPayLoad) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *TrendingEnvironmentDataPayLoad) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetSeries

`func (o *TrendingEnvironmentDataPayLoad) GetSeries() []DataValue`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *TrendingEnvironmentDataPayLoad) GetSeriesOk() (*[]DataValue, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *TrendingEnvironmentDataPayLoad) SetSeries(v []DataValue)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *TrendingEnvironmentDataPayLoad) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetUom

`func (o *TrendingEnvironmentDataPayLoad) GetUom() string`

GetUom returns the Uom field if non-nil, zero value otherwise.

### GetUomOk

`func (o *TrendingEnvironmentDataPayLoad) GetUomOk() (*string, bool)`

GetUomOk returns a tuple with the Uom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUom

`func (o *TrendingEnvironmentDataPayLoad) SetUom(v string)`

SetUom sets Uom field to given value.

### HasUom

`func (o *TrendingEnvironmentDataPayLoad) HasUom() bool`

HasUom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


