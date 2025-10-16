# Accounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | Customer account number | [optional] 
**AccountName** | Pointer to **string** | Customer account name | [optional] 
**AllowOrder** | Pointer to **bool** | When set to &#x60;true&#x60;, User is allowed to place an orders on this account | [optional] [default to true]

## Methods

### NewAccounts

`func NewAccounts() *Accounts`

NewAccounts instantiates a new Accounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsWithDefaults

`func NewAccountsWithDefaults() *Accounts`

NewAccountsWithDefaults instantiates a new Accounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *Accounts) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Accounts) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Accounts) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Accounts) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountName

`func (o *Accounts) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Accounts) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Accounts) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *Accounts) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAllowOrder

`func (o *Accounts) GetAllowOrder() bool`

GetAllowOrder returns the AllowOrder field if non-nil, zero value otherwise.

### GetAllowOrderOk

`func (o *Accounts) GetAllowOrderOk() (*bool, bool)`

GetAllowOrderOk returns a tuple with the AllowOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOrder

`func (o *Accounts) SetAllowOrder(v bool)`

SetAllowOrder sets AllowOrder field to given value.

### HasAllowOrder

`func (o *Accounts) HasAllowOrder() bool`

HasAllowOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


