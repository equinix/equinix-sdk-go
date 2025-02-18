# AlertRulePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Customer-provided stream name | [optional] 
**Description** | Pointer to **string** | Customer-provided stream description | [optional] 
**Enabled** | Pointer to **bool** | Stream alert rule enabled status | [optional] [default to true]
**MetricName** | Pointer to [**StreamAlertRuleMetricName**](StreamAlertRuleMetricName.md) |  | [optional] 
**ResourceSelector** | Pointer to [**ResourceSelector**](ResourceSelector.md) |  | [optional] 
**Operand** | Pointer to [**StreamAlertRuleOperand**](StreamAlertRuleOperand.md) |  | [optional] 
**WindowSize** | Pointer to **string** | Stream alert rule metric window size | [optional] 
**WarningThreshold** | Pointer to **string** | Stream alert rule metric warning threshold | [optional] 
**CriticalThreshold** | Pointer to **string** | Stream alert rule metric critical threshold | [optional] 

## Methods

### NewAlertRulePutRequest

`func NewAlertRulePutRequest() *AlertRulePutRequest`

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

### HasName

`func (o *AlertRulePutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

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

### GetMetricName

`func (o *AlertRulePutRequest) GetMetricName() StreamAlertRuleMetricName`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *AlertRulePutRequest) GetMetricNameOk() (*StreamAlertRuleMetricName, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *AlertRulePutRequest) SetMetricName(v StreamAlertRuleMetricName)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *AlertRulePutRequest) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

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

### HasResourceSelector

`func (o *AlertRulePutRequest) HasResourceSelector() bool`

HasResourceSelector returns a boolean if a field has been set.

### GetOperand

`func (o *AlertRulePutRequest) GetOperand() StreamAlertRuleOperand`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *AlertRulePutRequest) GetOperandOk() (*StreamAlertRuleOperand, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *AlertRulePutRequest) SetOperand(v StreamAlertRuleOperand)`

SetOperand sets Operand field to given value.

### HasOperand

`func (o *AlertRulePutRequest) HasOperand() bool`

HasOperand returns a boolean if a field has been set.

### GetWindowSize

`func (o *AlertRulePutRequest) GetWindowSize() string`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *AlertRulePutRequest) GetWindowSizeOk() (*string, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *AlertRulePutRequest) SetWindowSize(v string)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *AlertRulePutRequest) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetWarningThreshold

`func (o *AlertRulePutRequest) GetWarningThreshold() string`

GetWarningThreshold returns the WarningThreshold field if non-nil, zero value otherwise.

### GetWarningThresholdOk

`func (o *AlertRulePutRequest) GetWarningThresholdOk() (*string, bool)`

GetWarningThresholdOk returns a tuple with the WarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThreshold

`func (o *AlertRulePutRequest) SetWarningThreshold(v string)`

SetWarningThreshold sets WarningThreshold field to given value.

### HasWarningThreshold

`func (o *AlertRulePutRequest) HasWarningThreshold() bool`

HasWarningThreshold returns a boolean if a field has been set.

### GetCriticalThreshold

`func (o *AlertRulePutRequest) GetCriticalThreshold() string`

GetCriticalThreshold returns the CriticalThreshold field if non-nil, zero value otherwise.

### GetCriticalThresholdOk

`func (o *AlertRulePutRequest) GetCriticalThresholdOk() (*string, bool)`

GetCriticalThresholdOk returns a tuple with the CriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalThreshold

`func (o *AlertRulePutRequest) SetCriticalThreshold(v string)`

SetCriticalThreshold sets CriticalThreshold field to given value.

### HasCriticalThreshold

`func (o *AlertRulePutRequest) HasCriticalThreshold() bool`

HasCriticalThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


