# AlertRulePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Customer-provided stream name | 
**Type** | [**AlertRulePostRequestType**](AlertRulePostRequestType.md) |  | 
**Description** | Pointer to **string** | Customer-provided stream description | [optional] 
**Enabled** | Pointer to **bool** | Stream alert rule enabled status | [optional] [default to true]
**MetricSelector** | [**MetricSelector**](MetricSelector.md) |  | 
**ResourceSelector** | [**ResourceSelector**](ResourceSelector.md) |  | 
**DetectionMethod** | [**DetectionMethod**](DetectionMethod.md) |  | 

## Methods

### NewAlertRulePutRequest

`func NewAlertRulePutRequest(name string, type_ AlertRulePostRequestType, metricSelector MetricSelector, resourceSelector ResourceSelector, detectionMethod DetectionMethod, ) *AlertRulePutRequest`

NewAlertRulePutRequest instantiates a new AlertRulePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRulePutRequestWithDefaults

`func NewAlertRulePutRequestWithDefaults() *AlertRulePutRequest`

NewAlertRulePutRequestWithDefaults instantiates a new AlertRulePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlertRulePutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertRulePutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertRulePutRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AlertRulePutRequest) GetType() AlertRulePostRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertRulePutRequest) GetTypeOk() (*AlertRulePostRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertRulePutRequest) SetType(v AlertRulePostRequestType)`

SetType sets Type field to given value.


### GetDescription

`func (o *AlertRulePutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertRulePutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertRulePutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlertRulePutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AlertRulePutRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertRulePutRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertRulePutRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AlertRulePutRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMetricSelector

`func (o *AlertRulePutRequest) GetMetricSelector() MetricSelector`

GetMetricSelector returns the MetricSelector field if non-nil, zero value otherwise.

### GetMetricSelectorOk

`func (o *AlertRulePutRequest) GetMetricSelectorOk() (*MetricSelector, bool)`

GetMetricSelectorOk returns a tuple with the MetricSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSelector

`func (o *AlertRulePutRequest) SetMetricSelector(v MetricSelector)`

SetMetricSelector sets MetricSelector field to given value.


### GetResourceSelector

`func (o *AlertRulePutRequest) GetResourceSelector() ResourceSelector`

GetResourceSelector returns the ResourceSelector field if non-nil, zero value otherwise.

### GetResourceSelectorOk

`func (o *AlertRulePutRequest) GetResourceSelectorOk() (*ResourceSelector, bool)`

GetResourceSelectorOk returns a tuple with the ResourceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelector

`func (o *AlertRulePutRequest) SetResourceSelector(v ResourceSelector)`

SetResourceSelector sets ResourceSelector field to given value.


### GetDetectionMethod

`func (o *AlertRulePutRequest) GetDetectionMethod() DetectionMethod`

GetDetectionMethod returns the DetectionMethod field if non-nil, zero value otherwise.

### GetDetectionMethodOk

`func (o *AlertRulePutRequest) GetDetectionMethodOk() (*DetectionMethod, bool)`

GetDetectionMethodOk returns a tuple with the DetectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionMethod

`func (o *AlertRulePutRequest) SetDetectionMethod(v DetectionMethod)`

SetDetectionMethod sets DetectionMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


