# AlertRulePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AlertRulePostRequestType**](AlertRulePostRequestType.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided stream name | [optional] 
**Description** | Pointer to **string** | Customer-provided stream description | [optional] 
**Enabled** | Pointer to **bool** | Stream alert rule enabled status | [optional] [default to true]
**MetricName** | Pointer to [**AlertRulePostRequestMetricName**](AlertRulePostRequestMetricName.md) |  | [optional] 
**ResourceSelector** | Pointer to [**ResourceSelector**](ResourceSelector.md) |  | [optional] 
**WindowSize** | Pointer to **string** | Stream alert rule metric window size | [optional] 
**Operand** | Pointer to [**AlertRulePostRequestOperand**](AlertRulePostRequestOperand.md) |  | [optional] 
**WarningThreshold** | Pointer to **string** | Stream alert rule metric warning threshold | [optional] 
**CriticalThreshold** | Pointer to **string** | Stream alert rule metric critical threshold | [optional] 

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

### GetMetricName

`func (o *AlertRulePostRequest) GetMetricName() AlertRulePostRequestMetricName`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *AlertRulePostRequest) GetMetricNameOk() (*AlertRulePostRequestMetricName, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *AlertRulePostRequest) SetMetricName(v AlertRulePostRequestMetricName)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *AlertRulePostRequest) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

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

### GetWindowSize

`func (o *AlertRulePostRequest) GetWindowSize() string`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *AlertRulePostRequest) GetWindowSizeOk() (*string, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *AlertRulePostRequest) SetWindowSize(v string)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *AlertRulePostRequest) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetOperand

`func (o *AlertRulePostRequest) GetOperand() AlertRulePostRequestOperand`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *AlertRulePostRequest) GetOperandOk() (*AlertRulePostRequestOperand, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *AlertRulePostRequest) SetOperand(v AlertRulePostRequestOperand)`

SetOperand sets Operand field to given value.

### HasOperand

`func (o *AlertRulePostRequest) HasOperand() bool`

HasOperand returns a boolean if a field has been set.

### GetWarningThreshold

`func (o *AlertRulePostRequest) GetWarningThreshold() string`

GetWarningThreshold returns the WarningThreshold field if non-nil, zero value otherwise.

### GetWarningThresholdOk

`func (o *AlertRulePostRequest) GetWarningThresholdOk() (*string, bool)`

GetWarningThresholdOk returns a tuple with the WarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThreshold

`func (o *AlertRulePostRequest) SetWarningThreshold(v string)`

SetWarningThreshold sets WarningThreshold field to given value.

### HasWarningThreshold

`func (o *AlertRulePostRequest) HasWarningThreshold() bool`

HasWarningThreshold returns a boolean if a field has been set.

### GetCriticalThreshold

`func (o *AlertRulePostRequest) GetCriticalThreshold() string`

GetCriticalThreshold returns the CriticalThreshold field if non-nil, zero value otherwise.

### GetCriticalThresholdOk

`func (o *AlertRulePostRequest) GetCriticalThresholdOk() (*string, bool)`

GetCriticalThresholdOk returns a tuple with the CriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalThreshold

`func (o *AlertRulePostRequest) SetCriticalThreshold(v string)`

SetCriticalThreshold sets CriticalThreshold field to given value.

### HasCriticalThreshold

`func (o *AlertRulePostRequest) HasCriticalThreshold() bool`

HasCriticalThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


