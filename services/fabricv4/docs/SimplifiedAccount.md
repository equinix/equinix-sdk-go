# SimplifiedAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **int64** | Account number | [optional] 
**AccountName** | Pointer to **string** | Account name | [optional] 
**OrgId** | Pointer to **int64** | Customer organization identifier | [optional] 
**OrganizationName** | Pointer to **string** | Customer organization name | [optional] 
**GlobalOrgId** | Pointer to **string** | Global organization identifier | [optional] 
**GlobalOrganizationName** | Pointer to **string** | Global organization name | [optional] 
**UcmId** | Pointer to **string** | Account ucmId | [optional] 
**GlobalCustId** | Pointer to **string** | Account name | [optional] 
**ResellerAccountNumber** | Pointer to **int64** | Reseller account number | [optional] 
**ResellerAccountName** | Pointer to **string** | Reseller account name | [optional] 
**ResellerUcmId** | Pointer to **string** | Reseller account ucmId | [optional] 
**ResellerOrgId** | Pointer to **int64** | Reseller customer organization identifier | [optional] 

## Methods

### NewSimplifiedAccount

`func NewSimplifiedAccount() *SimplifiedAccount`

NewSimplifiedAccount instantiates a new SimplifiedAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedAccountWithDefaults

`func NewSimplifiedAccountWithDefaults() *SimplifiedAccount`

NewSimplifiedAccountWithDefaults instantiates a new SimplifiedAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *SimplifiedAccount) GetAccountNumber() int64`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *SimplifiedAccount) GetAccountNumberOk() (*int64, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *SimplifiedAccount) SetAccountNumber(v int64)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *SimplifiedAccount) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountName

`func (o *SimplifiedAccount) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *SimplifiedAccount) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *SimplifiedAccount) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *SimplifiedAccount) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetOrgId

`func (o *SimplifiedAccount) GetOrgId() int64`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SimplifiedAccount) GetOrgIdOk() (*int64, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SimplifiedAccount) SetOrgId(v int64)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SimplifiedAccount) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *SimplifiedAccount) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *SimplifiedAccount) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *SimplifiedAccount) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *SimplifiedAccount) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetGlobalOrgId

`func (o *SimplifiedAccount) GetGlobalOrgId() string`

GetGlobalOrgId returns the GlobalOrgId field if non-nil, zero value otherwise.

### GetGlobalOrgIdOk

`func (o *SimplifiedAccount) GetGlobalOrgIdOk() (*string, bool)`

GetGlobalOrgIdOk returns a tuple with the GlobalOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalOrgId

`func (o *SimplifiedAccount) SetGlobalOrgId(v string)`

SetGlobalOrgId sets GlobalOrgId field to given value.

### HasGlobalOrgId

`func (o *SimplifiedAccount) HasGlobalOrgId() bool`

HasGlobalOrgId returns a boolean if a field has been set.

### GetGlobalOrganizationName

`func (o *SimplifiedAccount) GetGlobalOrganizationName() string`

GetGlobalOrganizationName returns the GlobalOrganizationName field if non-nil, zero value otherwise.

### GetGlobalOrganizationNameOk

`func (o *SimplifiedAccount) GetGlobalOrganizationNameOk() (*string, bool)`

GetGlobalOrganizationNameOk returns a tuple with the GlobalOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalOrganizationName

`func (o *SimplifiedAccount) SetGlobalOrganizationName(v string)`

SetGlobalOrganizationName sets GlobalOrganizationName field to given value.

### HasGlobalOrganizationName

`func (o *SimplifiedAccount) HasGlobalOrganizationName() bool`

HasGlobalOrganizationName returns a boolean if a field has been set.

### GetUcmId

`func (o *SimplifiedAccount) GetUcmId() string`

GetUcmId returns the UcmId field if non-nil, zero value otherwise.

### GetUcmIdOk

`func (o *SimplifiedAccount) GetUcmIdOk() (*string, bool)`

GetUcmIdOk returns a tuple with the UcmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcmId

`func (o *SimplifiedAccount) SetUcmId(v string)`

SetUcmId sets UcmId field to given value.

### HasUcmId

`func (o *SimplifiedAccount) HasUcmId() bool`

HasUcmId returns a boolean if a field has been set.

### GetGlobalCustId

`func (o *SimplifiedAccount) GetGlobalCustId() string`

GetGlobalCustId returns the GlobalCustId field if non-nil, zero value otherwise.

### GetGlobalCustIdOk

`func (o *SimplifiedAccount) GetGlobalCustIdOk() (*string, bool)`

GetGlobalCustIdOk returns a tuple with the GlobalCustId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCustId

`func (o *SimplifiedAccount) SetGlobalCustId(v string)`

SetGlobalCustId sets GlobalCustId field to given value.

### HasGlobalCustId

`func (o *SimplifiedAccount) HasGlobalCustId() bool`

HasGlobalCustId returns a boolean if a field has been set.

### GetResellerAccountNumber

`func (o *SimplifiedAccount) GetResellerAccountNumber() int64`

GetResellerAccountNumber returns the ResellerAccountNumber field if non-nil, zero value otherwise.

### GetResellerAccountNumberOk

`func (o *SimplifiedAccount) GetResellerAccountNumberOk() (*int64, bool)`

GetResellerAccountNumberOk returns a tuple with the ResellerAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerAccountNumber

`func (o *SimplifiedAccount) SetResellerAccountNumber(v int64)`

SetResellerAccountNumber sets ResellerAccountNumber field to given value.

### HasResellerAccountNumber

`func (o *SimplifiedAccount) HasResellerAccountNumber() bool`

HasResellerAccountNumber returns a boolean if a field has been set.

### GetResellerAccountName

`func (o *SimplifiedAccount) GetResellerAccountName() string`

GetResellerAccountName returns the ResellerAccountName field if non-nil, zero value otherwise.

### GetResellerAccountNameOk

`func (o *SimplifiedAccount) GetResellerAccountNameOk() (*string, bool)`

GetResellerAccountNameOk returns a tuple with the ResellerAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerAccountName

`func (o *SimplifiedAccount) SetResellerAccountName(v string)`

SetResellerAccountName sets ResellerAccountName field to given value.

### HasResellerAccountName

`func (o *SimplifiedAccount) HasResellerAccountName() bool`

HasResellerAccountName returns a boolean if a field has been set.

### GetResellerUcmId

`func (o *SimplifiedAccount) GetResellerUcmId() string`

GetResellerUcmId returns the ResellerUcmId field if non-nil, zero value otherwise.

### GetResellerUcmIdOk

`func (o *SimplifiedAccount) GetResellerUcmIdOk() (*string, bool)`

GetResellerUcmIdOk returns a tuple with the ResellerUcmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerUcmId

`func (o *SimplifiedAccount) SetResellerUcmId(v string)`

SetResellerUcmId sets ResellerUcmId field to given value.

### HasResellerUcmId

`func (o *SimplifiedAccount) HasResellerUcmId() bool`

HasResellerUcmId returns a boolean if a field has been set.

### GetResellerOrgId

`func (o *SimplifiedAccount) GetResellerOrgId() int64`

GetResellerOrgId returns the ResellerOrgId field if non-nil, zero value otherwise.

### GetResellerOrgIdOk

`func (o *SimplifiedAccount) GetResellerOrgIdOk() (*int64, bool)`

GetResellerOrgIdOk returns a tuple with the ResellerOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerOrgId

`func (o *SimplifiedAccount) SetResellerOrgId(v int64)`

SetResellerOrgId sets ResellerOrgId field to given value.

### HasResellerOrgId

`func (o *SimplifiedAccount) HasResellerOrgId() bool`

HasResellerOrgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


