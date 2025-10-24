# MetroAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** | account Name | [optional] 
**AccountNumber** | Pointer to **int32** | account number | [optional] 
**AccountUcmId** | Pointer to **string** | account UcmId | [optional] 
**AccountStatus** | Pointer to **string** | account status | [optional] 
**Metros** | Pointer to **[]string** | An array of metros where the account is valid | [optional] 
**CreditHold** | Pointer to **bool** | Whether the account has a credit hold. You cannot use an account on credit hold to create a device. | [optional] 
**ReferenceId** | Pointer to **string** | referenceId | [optional] 

## Methods

### NewMetroAccountResponse

`func NewMetroAccountResponse() *MetroAccountResponse`

NewMetroAccountResponse instantiates a new MetroAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroAccountResponseWithDefaults

`func NewMetroAccountResponseWithDefaults() *MetroAccountResponse`

NewMetroAccountResponseWithDefaults instantiates a new MetroAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *MetroAccountResponse) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *MetroAccountResponse) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *MetroAccountResponse) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *MetroAccountResponse) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountNumber

`func (o *MetroAccountResponse) GetAccountNumber() int32`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *MetroAccountResponse) GetAccountNumberOk() (*int32, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *MetroAccountResponse) SetAccountNumber(v int32)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *MetroAccountResponse) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountUcmId

`func (o *MetroAccountResponse) GetAccountUcmId() string`

GetAccountUcmId returns the AccountUcmId field if non-nil, zero value otherwise.

### GetAccountUcmIdOk

`func (o *MetroAccountResponse) GetAccountUcmIdOk() (*string, bool)`

GetAccountUcmIdOk returns a tuple with the AccountUcmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUcmId

`func (o *MetroAccountResponse) SetAccountUcmId(v string)`

SetAccountUcmId sets AccountUcmId field to given value.

### HasAccountUcmId

`func (o *MetroAccountResponse) HasAccountUcmId() bool`

HasAccountUcmId returns a boolean if a field has been set.

### GetAccountStatus

`func (o *MetroAccountResponse) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *MetroAccountResponse) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *MetroAccountResponse) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *MetroAccountResponse) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetMetros

`func (o *MetroAccountResponse) GetMetros() []string`

GetMetros returns the Metros field if non-nil, zero value otherwise.

### GetMetrosOk

`func (o *MetroAccountResponse) GetMetrosOk() (*[]string, bool)`

GetMetrosOk returns a tuple with the Metros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetros

`func (o *MetroAccountResponse) SetMetros(v []string)`

SetMetros sets Metros field to given value.

### HasMetros

`func (o *MetroAccountResponse) HasMetros() bool`

HasMetros returns a boolean if a field has been set.

### GetCreditHold

`func (o *MetroAccountResponse) GetCreditHold() bool`

GetCreditHold returns the CreditHold field if non-nil, zero value otherwise.

### GetCreditHoldOk

`func (o *MetroAccountResponse) GetCreditHoldOk() (*bool, bool)`

GetCreditHoldOk returns a tuple with the CreditHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditHold

`func (o *MetroAccountResponse) SetCreditHold(v bool)`

SetCreditHold sets CreditHold field to given value.

### HasCreditHold

`func (o *MetroAccountResponse) HasCreditHold() bool`

HasCreditHold returns a boolean if a field has been set.

### GetReferenceId

`func (o *MetroAccountResponse) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *MetroAccountResponse) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *MetroAccountResponse) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *MetroAccountResponse) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


