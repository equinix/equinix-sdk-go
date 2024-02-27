# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **int32** |  | 
**IsResellerAccount** | Pointer to **bool** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**GlobalOrgId** | Pointer to **string** |  | [optional] 

## Methods

### NewAccount

`func NewAccount(accountNumber int32, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *Account) GetAccountNumber() int32`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Account) GetAccountNumberOk() (*int32, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Account) SetAccountNumber(v int32)`

SetAccountNumber sets AccountNumber field to given value.


### GetIsResellerAccount

`func (o *Account) GetIsResellerAccount() bool`

GetIsResellerAccount returns the IsResellerAccount field if non-nil, zero value otherwise.

### GetIsResellerAccountOk

`func (o *Account) GetIsResellerAccountOk() (*bool, bool)`

GetIsResellerAccountOk returns a tuple with the IsResellerAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResellerAccount

`func (o *Account) SetIsResellerAccount(v bool)`

SetIsResellerAccount sets IsResellerAccount field to given value.

### HasIsResellerAccount

`func (o *Account) HasIsResellerAccount() bool`

HasIsResellerAccount returns a boolean if a field has been set.

### GetOrgId

`func (o *Account) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Account) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Account) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Account) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetGlobalOrgId

`func (o *Account) GetGlobalOrgId() string`

GetGlobalOrgId returns the GlobalOrgId field if non-nil, zero value otherwise.

### GetGlobalOrgIdOk

`func (o *Account) GetGlobalOrgIdOk() (*string, bool)`

GetGlobalOrgIdOk returns a tuple with the GlobalOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalOrgId

`func (o *Account) SetGlobalOrgId(v string)`

SetGlobalOrgId sets GlobalOrgId field to given value.

### HasGlobalOrgId

`func (o *Account) HasGlobalOrgId() bool`

HasGlobalOrgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


