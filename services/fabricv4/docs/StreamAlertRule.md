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
**MetricName** | Pointer to [**AlertRulePostRequestMetricName**](AlertRulePostRequestMetricName.md) |  | [optional] 
**ResourceSelector** | Pointer to [**ResourceSelector**](ResourceSelector.md) |  | [optional] 
**WindowSize** | Pointer to **string** | Stream alert rule metric window size | [optional] 
**Operand** | Pointer to [**AlertRulePostRequestOperand**](AlertRulePostRequestOperand.md) |  | [optional] 
**WarningThreshold** | Pointer to **string** | Stream alert rule metric warning threshold | [optional] 
**CriticalThreshold** | Pointer to **string** | Stream alert rule metric critical threshold | [optional] 
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

### GetMetricName

`func (o *StreamAlertRule) GetMetricName() AlertRulePostRequestMetricName`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *StreamAlertRule) GetMetricNameOk() (*AlertRulePostRequestMetricName, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *StreamAlertRule) SetMetricName(v AlertRulePostRequestMetricName)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *StreamAlertRule) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetResourceSelector

`func (o *StreamAlertRule) GetResourceSelector() ResourceSelector`

GetResourceSelector returns the ResourceSelector field if non-nil, zero value otherwise.

### GetResourceSelectorOk

`func (o *StreamAlertRule) GetResourceSelectorOk() (*ResourceSelector, bool)`

GetResourceSelectorOk returns a tuple with the ResourceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelector

`func (o *StreamAlertRule) SetResourceSelector(v ResourceSelector)`

SetResourceSelector sets ResourceSelector field to given value.

### HasResourceSelector

`func (o *StreamAlertRule) HasResourceSelector() bool`

HasResourceSelector returns a boolean if a field has been set.

### GetWindowSize

`func (o *StreamAlertRule) GetWindowSize() string`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *StreamAlertRule) GetWindowSizeOk() (*string, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *StreamAlertRule) SetWindowSize(v string)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *StreamAlertRule) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetOperand

`func (o *StreamAlertRule) GetOperand() AlertRulePostRequestOperand`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *StreamAlertRule) GetOperandOk() (*AlertRulePostRequestOperand, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *StreamAlertRule) SetOperand(v AlertRulePostRequestOperand)`

SetOperand sets Operand field to given value.

### HasOperand

`func (o *StreamAlertRule) HasOperand() bool`

HasOperand returns a boolean if a field has been set.

### GetWarningThreshold

`func (o *StreamAlertRule) GetWarningThreshold() string`

GetWarningThreshold returns the WarningThreshold field if non-nil, zero value otherwise.

### GetWarningThresholdOk

`func (o *StreamAlertRule) GetWarningThresholdOk() (*string, bool)`

GetWarningThresholdOk returns a tuple with the WarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThreshold

`func (o *StreamAlertRule) SetWarningThreshold(v string)`

SetWarningThreshold sets WarningThreshold field to given value.

### HasWarningThreshold

`func (o *StreamAlertRule) HasWarningThreshold() bool`

HasWarningThreshold returns a boolean if a field has been set.

### GetCriticalThreshold

`func (o *StreamAlertRule) GetCriticalThreshold() string`

GetCriticalThreshold returns the CriticalThreshold field if non-nil, zero value otherwise.

### GetCriticalThresholdOk

`func (o *StreamAlertRule) GetCriticalThresholdOk() (*string, bool)`

GetCriticalThresholdOk returns a tuple with the CriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalThreshold

`func (o *StreamAlertRule) SetCriticalThreshold(v string)`

SetCriticalThreshold sets CriticalThreshold field to given value.

### HasCriticalThreshold

`func (o *StreamAlertRule) HasCriticalThreshold() bool`

HasCriticalThreshold returns a boolean if a field has been set.

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


