# BandwidthUtilization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to [**BandwidthUtilizationUnit**](BandwidthUtilizationUnit.md) |  | [optional] 
**MetricInterval** | Pointer to **string** | An interval formatted value, indicating the time-interval the metric objects within the response represent | [optional] 
**Inbound** | Pointer to [**Direction**](Direction.md) |  | [optional] 
**Outbound** | Pointer to [**Direction**](Direction.md) |  | [optional] 

## Methods

### NewBandwidthUtilization

`func NewBandwidthUtilization() *BandwidthUtilization`

NewBandwidthUtilization instantiates a new BandwidthUtilization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthUtilizationWithDefaults

`func NewBandwidthUtilizationWithDefaults() *BandwidthUtilization`

NewBandwidthUtilizationWithDefaults instantiates a new BandwidthUtilization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *BandwidthUtilization) GetUnit() BandwidthUtilizationUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BandwidthUtilization) GetUnitOk() (*BandwidthUtilizationUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BandwidthUtilization) SetUnit(v BandwidthUtilizationUnit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BandwidthUtilization) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetMetricInterval

`func (o *BandwidthUtilization) GetMetricInterval() string`

GetMetricInterval returns the MetricInterval field if non-nil, zero value otherwise.

### GetMetricIntervalOk

`func (o *BandwidthUtilization) GetMetricIntervalOk() (*string, bool)`

GetMetricIntervalOk returns a tuple with the MetricInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricInterval

`func (o *BandwidthUtilization) SetMetricInterval(v string)`

SetMetricInterval sets MetricInterval field to given value.

### HasMetricInterval

`func (o *BandwidthUtilization) HasMetricInterval() bool`

HasMetricInterval returns a boolean if a field has been set.

### GetInbound

`func (o *BandwidthUtilization) GetInbound() Direction`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *BandwidthUtilization) GetInboundOk() (*Direction, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *BandwidthUtilization) SetInbound(v Direction)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *BandwidthUtilization) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *BandwidthUtilization) GetOutbound() Direction`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *BandwidthUtilization) GetOutboundOk() (*Direction, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *BandwidthUtilization) SetOutbound(v Direction)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *BandwidthUtilization) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


