# PriceError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | [**PriceErrorErrorCode**](PriceErrorErrorCode.md) |  | 
**ErrorMessage** | [**PriceErrorErrorMessage**](PriceErrorErrorMessage.md) |  | 
**CorrelationId** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Help** | Pointer to **string** |  | [optional] 
**AdditionalInfo** | Pointer to [**[]PriceErrorAdditionalInfo**](PriceErrorAdditionalInfo.md) |  | [optional] 

## Methods

### NewPriceError

`func NewPriceError(errorCode PriceErrorErrorCode, errorMessage PriceErrorErrorMessage, ) *PriceError`

NewPriceError instantiates a new PriceError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceErrorWithDefaults

`func NewPriceErrorWithDefaults() *PriceError`

NewPriceErrorWithDefaults instantiates a new PriceError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *PriceError) GetErrorCode() PriceErrorErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *PriceError) GetErrorCodeOk() (*PriceErrorErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *PriceError) SetErrorCode(v PriceErrorErrorCode)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *PriceError) GetErrorMessage() PriceErrorErrorMessage`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PriceError) GetErrorMessageOk() (*PriceErrorErrorMessage, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PriceError) SetErrorMessage(v PriceErrorErrorMessage)`

SetErrorMessage sets ErrorMessage field to given value.


### GetCorrelationId

`func (o *PriceError) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *PriceError) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *PriceError) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *PriceError) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetDetails

`func (o *PriceError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PriceError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PriceError) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PriceError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetHelp

`func (o *PriceError) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *PriceError) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *PriceError) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *PriceError) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *PriceError) GetAdditionalInfo() []PriceErrorAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *PriceError) GetAdditionalInfoOk() (*[]PriceErrorAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *PriceError) SetAdditionalInfo(v []PriceErrorAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *PriceError) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


