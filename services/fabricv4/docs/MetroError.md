# MetroError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | [**MetroErrorErrorCode**](MetroErrorErrorCode.md) |  | 
**ErrorMessage** | [**MetroErrorErrorMessage**](MetroErrorErrorMessage.md) |  | 
**CorrelationId** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Help** | Pointer to **string** |  | [optional] 
**AdditionalInfo** | Pointer to [**[]PriceErrorAdditionalInfo**](PriceErrorAdditionalInfo.md) |  | [optional] 

## Methods

### NewMetroError

`func NewMetroError(errorCode MetroErrorErrorCode, errorMessage MetroErrorErrorMessage, ) *MetroError`

NewMetroError instantiates a new MetroError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroErrorWithDefaults

`func NewMetroErrorWithDefaults() *MetroError`

NewMetroErrorWithDefaults instantiates a new MetroError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *MetroError) GetErrorCode() MetroErrorErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MetroError) GetErrorCodeOk() (*MetroErrorErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MetroError) SetErrorCode(v MetroErrorErrorCode)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *MetroError) GetErrorMessage() MetroErrorErrorMessage`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *MetroError) GetErrorMessageOk() (*MetroErrorErrorMessage, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *MetroError) SetErrorMessage(v MetroErrorErrorMessage)`

SetErrorMessage sets ErrorMessage field to given value.


### GetCorrelationId

`func (o *MetroError) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MetroError) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *MetroError) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *MetroError) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetDetails

`func (o *MetroError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MetroError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MetroError) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MetroError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetHelp

`func (o *MetroError) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *MetroError) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *MetroError) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *MetroError) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *MetroError) GetAdditionalInfo() []PriceErrorAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *MetroError) GetAdditionalInfoOk() (*[]PriceErrorAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *MetroError) SetAdditionalInfo(v []PriceErrorAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *MetroError) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


