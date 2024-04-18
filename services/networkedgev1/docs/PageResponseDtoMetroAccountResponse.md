# PageResponseDtoMetroAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCreateUrl** | Pointer to **string** | accountCreateUrl | [optional] 
**Accounts** | Pointer to [**[]MetroAccountResponse**](MetroAccountResponse.md) |  | [optional] 
**ErrorMessage** | Pointer to **string** | error Message | [optional] 
**ErrorCode** | Pointer to **string** | error Code | [optional] 

## Methods

### NewPageResponseDtoMetroAccountResponse

`func NewPageResponseDtoMetroAccountResponse() *PageResponseDtoMetroAccountResponse`

NewPageResponseDtoMetroAccountResponse instantiates a new PageResponseDtoMetroAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageResponseDtoMetroAccountResponseWithDefaults

`func NewPageResponseDtoMetroAccountResponseWithDefaults() *PageResponseDtoMetroAccountResponse`

NewPageResponseDtoMetroAccountResponseWithDefaults instantiates a new PageResponseDtoMetroAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCreateUrl

`func (o *PageResponseDtoMetroAccountResponse) GetAccountCreateUrl() string`

GetAccountCreateUrl returns the AccountCreateUrl field if non-nil, zero value otherwise.

### GetAccountCreateUrlOk

`func (o *PageResponseDtoMetroAccountResponse) GetAccountCreateUrlOk() (*string, bool)`

GetAccountCreateUrlOk returns a tuple with the AccountCreateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreateUrl

`func (o *PageResponseDtoMetroAccountResponse) SetAccountCreateUrl(v string)`

SetAccountCreateUrl sets AccountCreateUrl field to given value.

### HasAccountCreateUrl

`func (o *PageResponseDtoMetroAccountResponse) HasAccountCreateUrl() bool`

HasAccountCreateUrl returns a boolean if a field has been set.

### GetAccounts

`func (o *PageResponseDtoMetroAccountResponse) GetAccounts() []MetroAccountResponse`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *PageResponseDtoMetroAccountResponse) GetAccountsOk() (*[]MetroAccountResponse, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *PageResponseDtoMetroAccountResponse) SetAccounts(v []MetroAccountResponse)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *PageResponseDtoMetroAccountResponse) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PageResponseDtoMetroAccountResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PageResponseDtoMetroAccountResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PageResponseDtoMetroAccountResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PageResponseDtoMetroAccountResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *PageResponseDtoMetroAccountResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *PageResponseDtoMetroAccountResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *PageResponseDtoMetroAccountResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *PageResponseDtoMetroAccountResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


