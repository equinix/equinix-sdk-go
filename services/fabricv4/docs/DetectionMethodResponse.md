# DetectionMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DetectionMethodType**](DetectionMethodType.md) |  | [optional] 
**WindowSize** | Pointer to **string** | Stream alert rule metric window size | [optional] 
**Operand** | Pointer to [**DetectionMethodOperand**](DetectionMethodOperand.md) |  | [optional] 
**WarningThreshold** | Pointer to **string** | Stream alert rule metric warning threshold | [optional] 
**CriticalThreshold** | Pointer to **string** | Stream alert rule metric critical threshold | [optional] 

## Methods

### NewDetectionMethodResponse

`func NewDetectionMethodResponse() *DetectionMethodResponse`

NewDetectionMethodResponse instantiates a new DetectionMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectionMethodResponseWithDefaults

`func NewDetectionMethodResponseWithDefaults() *DetectionMethodResponse`

NewDetectionMethodResponseWithDefaults instantiates a new DetectionMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DetectionMethodResponse) GetType() DetectionMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DetectionMethodResponse) GetTypeOk() (*DetectionMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DetectionMethodResponse) SetType(v DetectionMethodType)`

SetType sets Type field to given value.

### HasType

`func (o *DetectionMethodResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWindowSize

`func (o *DetectionMethodResponse) GetWindowSize() string`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *DetectionMethodResponse) GetWindowSizeOk() (*string, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *DetectionMethodResponse) SetWindowSize(v string)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *DetectionMethodResponse) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetOperand

`func (o *DetectionMethodResponse) GetOperand() DetectionMethodOperand`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *DetectionMethodResponse) GetOperandOk() (*DetectionMethodOperand, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *DetectionMethodResponse) SetOperand(v DetectionMethodOperand)`

SetOperand sets Operand field to given value.

### HasOperand

`func (o *DetectionMethodResponse) HasOperand() bool`

HasOperand returns a boolean if a field has been set.

### GetWarningThreshold

`func (o *DetectionMethodResponse) GetWarningThreshold() string`

GetWarningThreshold returns the WarningThreshold field if non-nil, zero value otherwise.

### GetWarningThresholdOk

`func (o *DetectionMethodResponse) GetWarningThresholdOk() (*string, bool)`

GetWarningThresholdOk returns a tuple with the WarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThreshold

`func (o *DetectionMethodResponse) SetWarningThreshold(v string)`

SetWarningThreshold sets WarningThreshold field to given value.

### HasWarningThreshold

`func (o *DetectionMethodResponse) HasWarningThreshold() bool`

HasWarningThreshold returns a boolean if a field has been set.

### GetCriticalThreshold

`func (o *DetectionMethodResponse) GetCriticalThreshold() string`

GetCriticalThreshold returns the CriticalThreshold field if non-nil, zero value otherwise.

### GetCriticalThresholdOk

`func (o *DetectionMethodResponse) GetCriticalThresholdOk() (*string, bool)`

GetCriticalThresholdOk returns a tuple with the CriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalThreshold

`func (o *DetectionMethodResponse) SetCriticalThreshold(v string)`

SetCriticalThreshold sets CriticalThreshold field to given value.

### HasCriticalThreshold

`func (o *DetectionMethodResponse) HasCriticalThreshold() bool`

HasCriticalThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


