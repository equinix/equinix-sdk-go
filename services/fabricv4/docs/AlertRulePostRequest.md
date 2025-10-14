# AlertRulePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AlertRulePostRequestType**](AlertRulePostRequestType.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided stream name | [optional] 
**Description** | Pointer to **string** | Customer-provided stream description | [optional] 
**Enabled** | Pointer to **bool** | Stream alert rule enabled status | [optional] [default to true]
**MetricSelector** | Pointer to [**MetricSelector**](MetricSelector.md) |  | [optional] 
**ResourceSelector** | Pointer to [**ResourceSelector**](ResourceSelector.md) |  | [optional] 
**DetectionMethod** | Pointer to [**DetectionMethod**](DetectionMethod.md) |  | [optional] 

## Methods

### NewAlertRulePostRequest

`func NewAlertRulePostRequest() *AlertRulePostRequest`

NewAlertRulePostRequest instantiates a new AlertRulePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRulePostRequestWithDefaults

`func NewAlertRulePostRequestWithDefaults() *AlertRulePostRequest`

NewAlertRulePostRequestWithDefaults instantiates a new AlertRulePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AlertRulePostRequest) GetType() AlertRulePostRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertRulePostRequest) GetTypeOk() (*AlertRulePostRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertRulePostRequest) SetType(v AlertRulePostRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *AlertRulePostRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AlertRulePostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertRulePostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertRulePostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertRulePostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AlertRulePostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertRulePostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertRulePostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlertRulePostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AlertRulePostRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertRulePostRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertRulePostRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AlertRulePostRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMetricSelector

`func (o *AlertRulePostRequest) GetMetricSelector() MetricSelector`

GetMetricSelector returns the MetricSelector field if non-nil, zero value otherwise.

### GetMetricSelectorOk

`func (o *AlertRulePostRequest) GetMetricSelectorOk() (*MetricSelector, bool)`

GetMetricSelectorOk returns a tuple with the MetricSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSelector

`func (o *AlertRulePostRequest) SetMetricSelector(v MetricSelector)`

SetMetricSelector sets MetricSelector field to given value.

### HasMetricSelector

`func (o *AlertRulePostRequest) HasMetricSelector() bool`

HasMetricSelector returns a boolean if a field has been set.

### GetResourceSelector

`func (o *AlertRulePostRequest) GetResourceSelector() ResourceSelector`

GetResourceSelector returns the ResourceSelector field if non-nil, zero value otherwise.

### GetResourceSelectorOk

`func (o *AlertRulePostRequest) GetResourceSelectorOk() (*ResourceSelector, bool)`

GetResourceSelectorOk returns a tuple with the ResourceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelector

`func (o *AlertRulePostRequest) SetResourceSelector(v ResourceSelector)`

SetResourceSelector sets ResourceSelector field to given value.

### HasResourceSelector

`func (o *AlertRulePostRequest) HasResourceSelector() bool`

HasResourceSelector returns a boolean if a field has been set.

### GetDetectionMethod

`func (o *AlertRulePostRequest) GetDetectionMethod() DetectionMethod`

GetDetectionMethod returns the DetectionMethod field if non-nil, zero value otherwise.

### GetDetectionMethodOk

`func (o *AlertRulePostRequest) GetDetectionMethodOk() (*DetectionMethod, bool)`

GetDetectionMethodOk returns a tuple with the DetectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionMethod

`func (o *AlertRulePostRequest) SetDetectionMethod(v DetectionMethod)`

SetDetectionMethod sets DetectionMethod field to given value.

### HasDetectionMethod

`func (o *AlertRulePostRequest) HasDetectionMethod() bool`

HasDetectionMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


