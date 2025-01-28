# StreamSubscriptionPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Customer-provided stream subscription name | [optional] 
**Description** | Pointer to **string** | Customer-provided stream subscription description | [optional] 
**Enabled** | Pointer to **bool** | Stream subscription enabled status | [optional] 
**Filters** | Pointer to [**StreamSubscriptionFilter**](StreamSubscriptionFilter.md) |  | [optional] 
**MetricSelector** | Pointer to [**StreamSubscriptionSelector**](StreamSubscriptionSelector.md) |  | [optional] 
**EventSelector** | Pointer to [**StreamSubscriptionSelector**](StreamSubscriptionSelector.md) |  | [optional] 
**Sink** | Pointer to [**StreamSubscriptionSink**](StreamSubscriptionSink.md) |  | [optional] 

## Methods

### NewStreamSubscriptionPutRequest

`func NewStreamSubscriptionPutRequest() *StreamSubscriptionPutRequest`

NewStreamSubscriptionPutRequest instantiates a new StreamSubscriptionPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSubscriptionPutRequestWithDefaults

`func NewStreamSubscriptionPutRequestWithDefaults() *StreamSubscriptionPutRequest`

NewStreamSubscriptionPutRequestWithDefaults instantiates a new StreamSubscriptionPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StreamSubscriptionPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamSubscriptionPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamSubscriptionPutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamSubscriptionPutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StreamSubscriptionPutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StreamSubscriptionPutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StreamSubscriptionPutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StreamSubscriptionPutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *StreamSubscriptionPutRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StreamSubscriptionPutRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StreamSubscriptionPutRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StreamSubscriptionPutRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFilters

`func (o *StreamSubscriptionPutRequest) GetFilters() StreamSubscriptionFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *StreamSubscriptionPutRequest) GetFiltersOk() (*StreamSubscriptionFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *StreamSubscriptionPutRequest) SetFilters(v StreamSubscriptionFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *StreamSubscriptionPutRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMetricSelector

`func (o *StreamSubscriptionPutRequest) GetMetricSelector() StreamSubscriptionSelector`

GetMetricSelector returns the MetricSelector field if non-nil, zero value otherwise.

### GetMetricSelectorOk

`func (o *StreamSubscriptionPutRequest) GetMetricSelectorOk() (*StreamSubscriptionSelector, bool)`

GetMetricSelectorOk returns a tuple with the MetricSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSelector

`func (o *StreamSubscriptionPutRequest) SetMetricSelector(v StreamSubscriptionSelector)`

SetMetricSelector sets MetricSelector field to given value.

### HasMetricSelector

`func (o *StreamSubscriptionPutRequest) HasMetricSelector() bool`

HasMetricSelector returns a boolean if a field has been set.

### GetEventSelector

`func (o *StreamSubscriptionPutRequest) GetEventSelector() StreamSubscriptionSelector`

GetEventSelector returns the EventSelector field if non-nil, zero value otherwise.

### GetEventSelectorOk

`func (o *StreamSubscriptionPutRequest) GetEventSelectorOk() (*StreamSubscriptionSelector, bool)`

GetEventSelectorOk returns a tuple with the EventSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSelector

`func (o *StreamSubscriptionPutRequest) SetEventSelector(v StreamSubscriptionSelector)`

SetEventSelector sets EventSelector field to given value.

### HasEventSelector

`func (o *StreamSubscriptionPutRequest) HasEventSelector() bool`

HasEventSelector returns a boolean if a field has been set.

### GetSink

`func (o *StreamSubscriptionPutRequest) GetSink() StreamSubscriptionSink`

GetSink returns the Sink field if non-nil, zero value otherwise.

### GetSinkOk

`func (o *StreamSubscriptionPutRequest) GetSinkOk() (*StreamSubscriptionSink, bool)`

GetSinkOk returns a tuple with the Sink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSink

`func (o *StreamSubscriptionPutRequest) SetSink(v StreamSubscriptionSink)`

SetSink sets Sink field to given value.

### HasSink

`func (o *StreamSubscriptionPutRequest) HasSink() bool`

HasSink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


