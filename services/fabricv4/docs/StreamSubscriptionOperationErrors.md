# StreamSubscriptionOperationErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** | Equinix Observability error code | [optional] 
**ErrorMessage** | Pointer to **string** | Equinix Observability error message | [optional] 
**DateTime** | Pointer to **time.Time** | Equinix Observability error date time | [optional] 
**AdditionalInfo** | Pointer to [**StreamSubscriptionOperationAdditionalInfo**](StreamSubscriptionOperationAdditionalInfo.md) |  | [optional] 

## Methods

### NewStreamSubscriptionOperationErrors

`func NewStreamSubscriptionOperationErrors() *StreamSubscriptionOperationErrors`

NewStreamSubscriptionOperationErrors instantiates a new StreamSubscriptionOperationErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSubscriptionOperationErrorsWithDefaults

`func NewStreamSubscriptionOperationErrorsWithDefaults() *StreamSubscriptionOperationErrors`

NewStreamSubscriptionOperationErrorsWithDefaults instantiates a new StreamSubscriptionOperationErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *StreamSubscriptionOperationErrors) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *StreamSubscriptionOperationErrors) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *StreamSubscriptionOperationErrors) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *StreamSubscriptionOperationErrors) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *StreamSubscriptionOperationErrors) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *StreamSubscriptionOperationErrors) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *StreamSubscriptionOperationErrors) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *StreamSubscriptionOperationErrors) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetDateTime

`func (o *StreamSubscriptionOperationErrors) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *StreamSubscriptionOperationErrors) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *StreamSubscriptionOperationErrors) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *StreamSubscriptionOperationErrors) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *StreamSubscriptionOperationErrors) GetAdditionalInfo() StreamSubscriptionOperationAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *StreamSubscriptionOperationErrors) GetAdditionalInfoOk() (*StreamSubscriptionOperationAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *StreamSubscriptionOperationErrors) SetAdditionalInfo(v StreamSubscriptionOperationAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *StreamSubscriptionOperationErrors) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


