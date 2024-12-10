# StreamSubscriptionPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**StreamSubscriptionPostRequestType**](StreamSubscriptionPostRequestType.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided stream subscription name | [optional] 
**Description** | Pointer to **string** | Customer-provided stream subscription description | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Stream subscription enabled status | [optional] 
**Stream** | Pointer to [**StreamTarget**](StreamTarget.md) |  | [optional] 
**Filters** | Pointer to [**StreamSubscriptionFilter**](StreamSubscriptionFilter.md) |  | [optional] 
**MetricSelector** | Pointer to [**StreamSubscriptionSelector**](StreamSubscriptionSelector.md) |  | [optional] 
**EventSelector** | Pointer to [**StreamSubscriptionSelector**](StreamSubscriptionSelector.md) |  | [optional] 
**Sink** | Pointer to [**StreamSubscriptionSink**](StreamSubscriptionSink.md) |  | [optional] 

## Methods

### NewStreamSubscriptionPostRequest

`func NewStreamSubscriptionPostRequest() *StreamSubscriptionPostRequest`

NewStreamSubscriptionPostRequest instantiates a new StreamSubscriptionPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSubscriptionPostRequestWithDefaults

`func NewStreamSubscriptionPostRequestWithDefaults() *StreamSubscriptionPostRequest`

NewStreamSubscriptionPostRequestWithDefaults instantiates a new StreamSubscriptionPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StreamSubscriptionPostRequest) GetType() StreamSubscriptionPostRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamSubscriptionPostRequest) GetTypeOk() (*StreamSubscriptionPostRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamSubscriptionPostRequest) SetType(v StreamSubscriptionPostRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *StreamSubscriptionPostRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *StreamSubscriptionPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamSubscriptionPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamSubscriptionPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamSubscriptionPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StreamSubscriptionPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StreamSubscriptionPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StreamSubscriptionPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StreamSubscriptionPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProject

`func (o *StreamSubscriptionPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *StreamSubscriptionPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *StreamSubscriptionPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *StreamSubscriptionPostRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetEnabled

`func (o *StreamSubscriptionPostRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StreamSubscriptionPostRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StreamSubscriptionPostRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StreamSubscriptionPostRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStream

`func (o *StreamSubscriptionPostRequest) GetStream() StreamTarget`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *StreamSubscriptionPostRequest) GetStreamOk() (*StreamTarget, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *StreamSubscriptionPostRequest) SetStream(v StreamTarget)`

SetStream sets Stream field to given value.

### HasStream

`func (o *StreamSubscriptionPostRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetFilters

`func (o *StreamSubscriptionPostRequest) GetFilters() StreamSubscriptionFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *StreamSubscriptionPostRequest) GetFiltersOk() (*StreamSubscriptionFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *StreamSubscriptionPostRequest) SetFilters(v StreamSubscriptionFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *StreamSubscriptionPostRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMetricSelector

`func (o *StreamSubscriptionPostRequest) GetMetricSelector() StreamSubscriptionSelector`

GetMetricSelector returns the MetricSelector field if non-nil, zero value otherwise.

### GetMetricSelectorOk

`func (o *StreamSubscriptionPostRequest) GetMetricSelectorOk() (*StreamSubscriptionSelector, bool)`

GetMetricSelectorOk returns a tuple with the MetricSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSelector

`func (o *StreamSubscriptionPostRequest) SetMetricSelector(v StreamSubscriptionSelector)`

SetMetricSelector sets MetricSelector field to given value.

### HasMetricSelector

`func (o *StreamSubscriptionPostRequest) HasMetricSelector() bool`

HasMetricSelector returns a boolean if a field has been set.

### GetEventSelector

`func (o *StreamSubscriptionPostRequest) GetEventSelector() StreamSubscriptionSelector`

GetEventSelector returns the EventSelector field if non-nil, zero value otherwise.

### GetEventSelectorOk

`func (o *StreamSubscriptionPostRequest) GetEventSelectorOk() (*StreamSubscriptionSelector, bool)`

GetEventSelectorOk returns a tuple with the EventSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSelector

`func (o *StreamSubscriptionPostRequest) SetEventSelector(v StreamSubscriptionSelector)`

SetEventSelector sets EventSelector field to given value.

### HasEventSelector

`func (o *StreamSubscriptionPostRequest) HasEventSelector() bool`

HasEventSelector returns a boolean if a field has been set.

### GetSink

`func (o *StreamSubscriptionPostRequest) GetSink() StreamSubscriptionSink`

GetSink returns the Sink field if non-nil, zero value otherwise.

### GetSinkOk

`func (o *StreamSubscriptionPostRequest) GetSinkOk() (*StreamSubscriptionSink, bool)`

GetSinkOk returns a tuple with the Sink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSink

`func (o *StreamSubscriptionPostRequest) SetSink(v StreamSubscriptionSink)`

SetSink sets Sink field to given value.

### HasSink

`func (o *StreamSubscriptionPostRequest) HasSink() bool`

HasSink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


