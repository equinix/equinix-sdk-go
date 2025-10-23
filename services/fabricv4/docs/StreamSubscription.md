# StreamSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Stream Subscription URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Type** | Pointer to [**StreamSubscriptionType**](StreamSubscriptionType.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided subscription name | [optional] 
**Description** | Pointer to **string** | Customer-provided subscription description | [optional] 
**State** | Pointer to [**StreamSubscriptionState**](StreamSubscriptionState.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Stream subscription enabled status | [optional] 
**MetricSelector** | Pointer to [**StreamSubscriptionSelector**](StreamSubscriptionSelector.md) |  | [optional] 
**EventSelector** | Pointer to [**StreamSubscriptionSelector**](StreamSubscriptionSelector.md) |  | [optional] 
**Sink** | Pointer to [**StreamSubscriptionSinkResponse**](StreamSubscriptionSinkResponse.md) |  | [optional] 
**Operation** | Pointer to [**StreamSubscriptionOperation**](StreamSubscriptionOperation.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewStreamSubscription

`func NewStreamSubscription() *StreamSubscription`

NewStreamSubscription instantiates a new StreamSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSubscriptionWithDefaults

`func NewStreamSubscriptionWithDefaults() *StreamSubscription`

NewStreamSubscriptionWithDefaults instantiates a new StreamSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *StreamSubscription) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *StreamSubscription) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *StreamSubscription) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *StreamSubscription) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *StreamSubscription) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StreamSubscription) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StreamSubscription) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StreamSubscription) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *StreamSubscription) GetType() StreamSubscriptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamSubscription) GetTypeOk() (*StreamSubscriptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamSubscription) SetType(v StreamSubscriptionType)`

SetType sets Type field to given value.

### HasType

`func (o *StreamSubscription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *StreamSubscription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamSubscription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamSubscription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamSubscription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StreamSubscription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StreamSubscription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StreamSubscription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StreamSubscription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *StreamSubscription) GetState() StreamSubscriptionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StreamSubscription) GetStateOk() (*StreamSubscriptionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StreamSubscription) SetState(v StreamSubscriptionState)`

SetState sets State field to given value.

### HasState

`func (o *StreamSubscription) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEnabled

`func (o *StreamSubscription) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StreamSubscription) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StreamSubscription) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StreamSubscription) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMetricSelector

`func (o *StreamSubscription) GetMetricSelector() StreamSubscriptionSelector`

GetMetricSelector returns the MetricSelector field if non-nil, zero value otherwise.

### GetMetricSelectorOk

`func (o *StreamSubscription) GetMetricSelectorOk() (*StreamSubscriptionSelector, bool)`

GetMetricSelectorOk returns a tuple with the MetricSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSelector

`func (o *StreamSubscription) SetMetricSelector(v StreamSubscriptionSelector)`

SetMetricSelector sets MetricSelector field to given value.

### HasMetricSelector

`func (o *StreamSubscription) HasMetricSelector() bool`

HasMetricSelector returns a boolean if a field has been set.

### GetEventSelector

`func (o *StreamSubscription) GetEventSelector() StreamSubscriptionSelector`

GetEventSelector returns the EventSelector field if non-nil, zero value otherwise.

### GetEventSelectorOk

`func (o *StreamSubscription) GetEventSelectorOk() (*StreamSubscriptionSelector, bool)`

GetEventSelectorOk returns a tuple with the EventSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSelector

`func (o *StreamSubscription) SetEventSelector(v StreamSubscriptionSelector)`

SetEventSelector sets EventSelector field to given value.

### HasEventSelector

`func (o *StreamSubscription) HasEventSelector() bool`

HasEventSelector returns a boolean if a field has been set.

### GetSink

`func (o *StreamSubscription) GetSink() StreamSubscriptionSinkResponse`

GetSink returns the Sink field if non-nil, zero value otherwise.

### GetSinkOk

`func (o *StreamSubscription) GetSinkOk() (*StreamSubscriptionSinkResponse, bool)`

GetSinkOk returns a tuple with the Sink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSink

`func (o *StreamSubscription) SetSink(v StreamSubscriptionSinkResponse)`

SetSink sets Sink field to given value.

### HasSink

`func (o *StreamSubscription) HasSink() bool`

HasSink returns a boolean if a field has been set.

### GetOperation

`func (o *StreamSubscription) GetOperation() StreamSubscriptionOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *StreamSubscription) GetOperationOk() (*StreamSubscriptionOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *StreamSubscription) SetOperation(v StreamSubscriptionOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *StreamSubscription) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetChangeLog

`func (o *StreamSubscription) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *StreamSubscription) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *StreamSubscription) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *StreamSubscription) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


