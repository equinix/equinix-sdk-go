# ValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** | Application error code. Format :EQ-&lt;error code&gt; | 
**ErrorMessage** | **string** | Application error message | 
**CorrelationId** | Pointer to **string** | Correlation ID without any sensitive or meaningful information. | [optional] 
**Details** | Pointer to **string** | More information on errors. | [optional] 
**Help** | Pointer to **string** | Help message or URL to a help page for the corresponding errorCode. | [optional] 
**AdditionalInfo** | Pointer to [**[]ValidationErrorAdditionalInfo**](ValidationErrorAdditionalInfo.md) | Contains application specific information for this error. The object inside this array can have any number of application specific properties. | [optional] 

## Methods

### NewValidationError

`func NewValidationError(errorCode string, errorMessage string, ) *ValidationError`

NewValidationError instantiates a new ValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorWithDefaults

`func NewValidationErrorWithDefaults() *ValidationError`

NewValidationErrorWithDefaults instantiates a new ValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *ValidationError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ValidationError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ValidationError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *ValidationError) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ValidationError) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ValidationError) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetCorrelationId

`func (o *ValidationError) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *ValidationError) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *ValidationError) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *ValidationError) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetDetails

`func (o *ValidationError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ValidationError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ValidationError) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ValidationError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetHelp

`func (o *ValidationError) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *ValidationError) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *ValidationError) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *ValidationError) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *ValidationError) GetAdditionalInfo() []ValidationErrorAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ValidationError) GetAdditionalInfoOk() (*[]ValidationErrorAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ValidationError) SetAdditionalInfo(v []ValidationErrorAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ValidationError) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


