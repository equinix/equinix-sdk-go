# ErrorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**CorrelationId** | Pointer to **string** |  | [optional] 
**Help** | Pointer to **string** |  | [optional] 
**AdditionalInfo** | Pointer to [**ErrorDetails**](ErrorDetails.md) |  | [optional] 

## Methods

### NewErrorDetails

`func NewErrorDetails() *ErrorDetails`

NewErrorDetails instantiates a new ErrorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailsWithDefaults

`func NewErrorDetailsWithDefaults() *ErrorDetails`

NewErrorDetailsWithDefaults instantiates a new ErrorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *ErrorDetails) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ErrorDetails) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ErrorDetails) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ErrorDetails) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ErrorDetails) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ErrorDetails) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ErrorDetails) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ErrorDetails) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetDetails

`func (o *ErrorDetails) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorDetails) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorDetails) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorDetails) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetCorrelationId

`func (o *ErrorDetails) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *ErrorDetails) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *ErrorDetails) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *ErrorDetails) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetHelp

`func (o *ErrorDetails) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *ErrorDetails) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *ErrorDetails) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *ErrorDetails) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *ErrorDetails) GetAdditionalInfo() ErrorDetails`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ErrorDetails) GetAdditionalInfoOk() (*ErrorDetails, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ErrorDetails) SetAdditionalInfo(v ErrorDetails)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ErrorDetails) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


