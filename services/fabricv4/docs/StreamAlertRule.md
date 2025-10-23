# StreamAlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Stream Alert Rule URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Type** | Pointer to [**StreamAlertRuleType**](StreamAlertRuleType.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided stream alert rule name | [optional] 
**Description** | Pointer to **string** | Customer-provided stream alert rule description | [optional] 
**State** | Pointer to [**StreamAlertRuleState**](StreamAlertRuleState.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Stream alert rule enabled status | [optional] [default to true]
**MetricSelector** | Pointer to [**MetricSelectorResponse**](MetricSelectorResponse.md) |  | [optional] 
**ResourceSelector** | Pointer to [**ResourceSelectorResponse**](ResourceSelectorResponse.md) |  | [optional] 
**DetectionMethod** | Pointer to [**DetectionMethodResponse**](DetectionMethodResponse.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewStreamAlertRule

`func NewStreamAlertRule() *StreamAlertRule`

NewStreamAlertRule instantiates a new StreamAlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamAlertRuleWithDefaults

`func NewStreamAlertRuleWithDefaults() *StreamAlertRule`

NewStreamAlertRuleWithDefaults instantiates a new StreamAlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *StreamAlertRule) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *StreamAlertRule) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *StreamAlertRule) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *StreamAlertRule) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *StreamAlertRule) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StreamAlertRule) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StreamAlertRule) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StreamAlertRule) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *StreamAlertRule) GetType() StreamAlertRuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamAlertRule) GetTypeOk() (*StreamAlertRuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamAlertRule) SetType(v StreamAlertRuleType)`

SetType sets Type field to given value.

### HasType

`func (o *StreamAlertRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *StreamAlertRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamAlertRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamAlertRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamAlertRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StreamAlertRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StreamAlertRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StreamAlertRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StreamAlertRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *StreamAlertRule) GetState() StreamAlertRuleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StreamAlertRule) GetStateOk() (*StreamAlertRuleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StreamAlertRule) SetState(v StreamAlertRuleState)`

SetState sets State field to given value.

### HasState

`func (o *StreamAlertRule) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEnabled

`func (o *StreamAlertRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StreamAlertRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StreamAlertRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StreamAlertRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMetricSelector

`func (o *StreamAlertRule) GetMetricSelector() MetricSelectorResponse`

GetMetricSelector returns the MetricSelector field if non-nil, zero value otherwise.

### GetMetricSelectorOk

`func (o *StreamAlertRule) GetMetricSelectorOk() (*MetricSelectorResponse, bool)`

GetMetricSelectorOk returns a tuple with the MetricSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSelector

`func (o *StreamAlertRule) SetMetricSelector(v MetricSelectorResponse)`

SetMetricSelector sets MetricSelector field to given value.

### HasMetricSelector

`func (o *StreamAlertRule) HasMetricSelector() bool`

HasMetricSelector returns a boolean if a field has been set.

### GetResourceSelector

`func (o *StreamAlertRule) GetResourceSelector() ResourceSelectorResponse`

GetResourceSelector returns the ResourceSelector field if non-nil, zero value otherwise.

### GetResourceSelectorOk

`func (o *StreamAlertRule) GetResourceSelectorOk() (*ResourceSelectorResponse, bool)`

GetResourceSelectorOk returns a tuple with the ResourceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelector

`func (o *StreamAlertRule) SetResourceSelector(v ResourceSelectorResponse)`

SetResourceSelector sets ResourceSelector field to given value.

### HasResourceSelector

`func (o *StreamAlertRule) HasResourceSelector() bool`

HasResourceSelector returns a boolean if a field has been set.

### GetDetectionMethod

`func (o *StreamAlertRule) GetDetectionMethod() DetectionMethodResponse`

GetDetectionMethod returns the DetectionMethod field if non-nil, zero value otherwise.

### GetDetectionMethodOk

`func (o *StreamAlertRule) GetDetectionMethodOk() (*DetectionMethodResponse, bool)`

GetDetectionMethodOk returns a tuple with the DetectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionMethod

`func (o *StreamAlertRule) SetDetectionMethod(v DetectionMethodResponse)`

SetDetectionMethod sets DetectionMethod field to given value.

### HasDetectionMethod

`func (o *StreamAlertRule) HasDetectionMethod() bool`

HasDetectionMethod returns a boolean if a field has been set.

### GetChangeLog

`func (o *StreamAlertRule) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *StreamAlertRule) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *StreamAlertRule) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *StreamAlertRule) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


