# DetectionMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DetectionMethodType**](DetectionMethodType.md) |  | [optional] 
**WindowSize** | Pointer to **string** | Stream alert rule metric window size | [optional] 
**Operand** | Pointer to [**DetectionMethodOperand**](DetectionMethodOperand.md) |  | [optional] 
**WarningThreshold** | Pointer to **string** | Stream alert rule metric warning threshold | [optional] 
**CriticalThreshold** | Pointer to **string** | Stream alert rule metric critical threshold | [optional] 

## Methods

### NewDetectionMethod

`func NewDetectionMethod() *DetectionMethod`

NewDetectionMethod instantiates a new DetectionMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectionMethodWithDefaults

`func NewDetectionMethodWithDefaults() *DetectionMethod`

NewDetectionMethodWithDefaults instantiates a new DetectionMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DetectionMethod) GetType() DetectionMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DetectionMethod) GetTypeOk() (*DetectionMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DetectionMethod) SetType(v DetectionMethodType)`

SetType sets Type field to given value.

### HasType

`func (o *DetectionMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWindowSize

`func (o *DetectionMethod) GetWindowSize() string`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *DetectionMethod) GetWindowSizeOk() (*string, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *DetectionMethod) SetWindowSize(v string)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *DetectionMethod) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetOperand

`func (o *DetectionMethod) GetOperand() DetectionMethodOperand`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *DetectionMethod) GetOperandOk() (*DetectionMethodOperand, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *DetectionMethod) SetOperand(v DetectionMethodOperand)`

SetOperand sets Operand field to given value.

### HasOperand

`func (o *DetectionMethod) HasOperand() bool`

HasOperand returns a boolean if a field has been set.

### GetWarningThreshold

`func (o *DetectionMethod) GetWarningThreshold() string`

GetWarningThreshold returns the WarningThreshold field if non-nil, zero value otherwise.

### GetWarningThresholdOk

`func (o *DetectionMethod) GetWarningThresholdOk() (*string, bool)`

GetWarningThresholdOk returns a tuple with the WarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThreshold

`func (o *DetectionMethod) SetWarningThreshold(v string)`

SetWarningThreshold sets WarningThreshold field to given value.

### HasWarningThreshold

`func (o *DetectionMethod) HasWarningThreshold() bool`

HasWarningThreshold returns a boolean if a field has been set.

### GetCriticalThreshold

`func (o *DetectionMethod) GetCriticalThreshold() string`

GetCriticalThreshold returns the CriticalThreshold field if non-nil, zero value otherwise.

### GetCriticalThresholdOk

`func (o *DetectionMethod) GetCriticalThresholdOk() (*string, bool)`

GetCriticalThresholdOk returns a tuple with the CriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalThreshold

`func (o *DetectionMethod) SetCriticalThreshold(v string)`

SetCriticalThreshold sets CriticalThreshold field to given value.

### HasCriticalThreshold

`func (o *DetectionMethod) HasCriticalThreshold() bool`

HasCriticalThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


