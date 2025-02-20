# ComparisonData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datapoint** | Pointer to [**ComparisonDataDatapoint**](ComparisonDataDatapoint.md) |  | [optional] 
**LastMonth** | Pointer to **float32** | comparison for the current value of the datapoint with the last  month&#39;s value  | [optional] 
**LastQuarter** | Pointer to **float32** | comparison for the current value of the datapoint with the last  quarter&#39;s value   | [optional] 
**LastWeek** | Pointer to **float32** | comparison for the current value of the datapoint with last week&#39;s  value  | [optional] 
**Yesterday** | Pointer to **float32** | comparison for the current value of the datapoint with yesterday&#39;s  value  | [optional] 

## Methods

### NewComparisonData

`func NewComparisonData() *ComparisonData`

NewComparisonData instantiates a new ComparisonData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComparisonDataWithDefaults

`func NewComparisonDataWithDefaults() *ComparisonData`

NewComparisonDataWithDefaults instantiates a new ComparisonData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatapoint

`func (o *ComparisonData) GetDatapoint() ComparisonDataDatapoint`

GetDatapoint returns the Datapoint field if non-nil, zero value otherwise.

### GetDatapointOk

`func (o *ComparisonData) GetDatapointOk() (*ComparisonDataDatapoint, bool)`

GetDatapointOk returns a tuple with the Datapoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatapoint

`func (o *ComparisonData) SetDatapoint(v ComparisonDataDatapoint)`

SetDatapoint sets Datapoint field to given value.

### HasDatapoint

`func (o *ComparisonData) HasDatapoint() bool`

HasDatapoint returns a boolean if a field has been set.

### GetLastMonth

`func (o *ComparisonData) GetLastMonth() float32`

GetLastMonth returns the LastMonth field if non-nil, zero value otherwise.

### GetLastMonthOk

`func (o *ComparisonData) GetLastMonthOk() (*float32, bool)`

GetLastMonthOk returns a tuple with the LastMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMonth

`func (o *ComparisonData) SetLastMonth(v float32)`

SetLastMonth sets LastMonth field to given value.

### HasLastMonth

`func (o *ComparisonData) HasLastMonth() bool`

HasLastMonth returns a boolean if a field has been set.

### GetLastQuarter

`func (o *ComparisonData) GetLastQuarter() float32`

GetLastQuarter returns the LastQuarter field if non-nil, zero value otherwise.

### GetLastQuarterOk

`func (o *ComparisonData) GetLastQuarterOk() (*float32, bool)`

GetLastQuarterOk returns a tuple with the LastQuarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQuarter

`func (o *ComparisonData) SetLastQuarter(v float32)`

SetLastQuarter sets LastQuarter field to given value.

### HasLastQuarter

`func (o *ComparisonData) HasLastQuarter() bool`

HasLastQuarter returns a boolean if a field has been set.

### GetLastWeek

`func (o *ComparisonData) GetLastWeek() float32`

GetLastWeek returns the LastWeek field if non-nil, zero value otherwise.

### GetLastWeekOk

`func (o *ComparisonData) GetLastWeekOk() (*float32, bool)`

GetLastWeekOk returns a tuple with the LastWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWeek

`func (o *ComparisonData) SetLastWeek(v float32)`

SetLastWeek sets LastWeek field to given value.

### HasLastWeek

`func (o *ComparisonData) HasLastWeek() bool`

HasLastWeek returns a boolean if a field has been set.

### GetYesterday

`func (o *ComparisonData) GetYesterday() float32`

GetYesterday returns the Yesterday field if non-nil, zero value otherwise.

### GetYesterdayOk

`func (o *ComparisonData) GetYesterdayOk() (*float32, bool)`

GetYesterdayOk returns a tuple with the Yesterday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesterday

`func (o *ComparisonData) SetYesterday(v float32)`

SetYesterday sets Yesterday field to given value.

### HasYesterday

`func (o *ComparisonData) HasYesterday() bool`

HasYesterday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


