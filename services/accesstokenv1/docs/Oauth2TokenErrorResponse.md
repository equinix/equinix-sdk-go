# Oauth2TokenErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorDomain** | **string** | The module associated with this error | 
**ErrorTitle** | **string** | The error title | 
**ErrorCode** | **string** | The code used to identify the error category | 
**DeveloperMessage** | **string** | The error message to be used for auditing purpose by the consuming application | 
**ErrorMessage** | **string** | The error message which could be displayed to the end user | 

## Methods

### NewOauth2TokenErrorResponse

`func NewOauth2TokenErrorResponse(errorDomain string, errorTitle string, errorCode string, developerMessage string, errorMessage string, ) *Oauth2TokenErrorResponse`

NewOauth2TokenErrorResponse instantiates a new Oauth2TokenErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2TokenErrorResponseWithDefaults

`func NewOauth2TokenErrorResponseWithDefaults() *Oauth2TokenErrorResponse`

NewOauth2TokenErrorResponseWithDefaults instantiates a new Oauth2TokenErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorDomain

`func (o *Oauth2TokenErrorResponse) GetErrorDomain() string`

GetErrorDomain returns the ErrorDomain field if non-nil, zero value otherwise.

### GetErrorDomainOk

`func (o *Oauth2TokenErrorResponse) GetErrorDomainOk() (*string, bool)`

GetErrorDomainOk returns a tuple with the ErrorDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDomain

`func (o *Oauth2TokenErrorResponse) SetErrorDomain(v string)`

SetErrorDomain sets ErrorDomain field to given value.


### GetErrorTitle

`func (o *Oauth2TokenErrorResponse) GetErrorTitle() string`

GetErrorTitle returns the ErrorTitle field if non-nil, zero value otherwise.

### GetErrorTitleOk

`func (o *Oauth2TokenErrorResponse) GetErrorTitleOk() (*string, bool)`

GetErrorTitleOk returns a tuple with the ErrorTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorTitle

`func (o *Oauth2TokenErrorResponse) SetErrorTitle(v string)`

SetErrorTitle sets ErrorTitle field to given value.


### GetErrorCode

`func (o *Oauth2TokenErrorResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Oauth2TokenErrorResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Oauth2TokenErrorResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetDeveloperMessage

`func (o *Oauth2TokenErrorResponse) GetDeveloperMessage() string`

GetDeveloperMessage returns the DeveloperMessage field if non-nil, zero value otherwise.

### GetDeveloperMessageOk

`func (o *Oauth2TokenErrorResponse) GetDeveloperMessageOk() (*string, bool)`

GetDeveloperMessageOk returns a tuple with the DeveloperMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperMessage

`func (o *Oauth2TokenErrorResponse) SetDeveloperMessage(v string)`

SetDeveloperMessage sets DeveloperMessage field to given value.


### GetErrorMessage

`func (o *Oauth2TokenErrorResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Oauth2TokenErrorResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Oauth2TokenErrorResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


