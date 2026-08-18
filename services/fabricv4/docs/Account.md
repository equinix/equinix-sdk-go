# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | Account identifier | [optional] 
**AccountNumber** | Pointer to **int32** | Account number | [optional] 
**AccountName** | Pointer to **string** | Account name | [optional] 
**UcmId** | Pointer to **string** | Account ucm id | [optional] 
**GlobalCustId** | Pointer to **string** | Global customer organization id | [optional] 
**OrgId** | Pointer to **int32** | Customer organization id | [optional] 
**OrganizationName** | Pointer to **string** | Customer organization name | [optional] 
**SubCustomes** | Pointer to [**[]Account**](Account.md) | All sub customer accounts | [optional] 
**CountryCode** | Pointer to **string** | Account country code | [optional] 
**OperationalUnit** | Pointer to **string** | Account operational unit | [optional] 
**OperationalUnitMetros** | Pointer to **[]string** | Account operational unit metros | [optional] 
**SignatureRequired** | Pointer to **bool** | Is signature required | [optional] 
**PoBearing** | Pointer to **bool** | Purchase order bearing | [optional] 
**Default** | Pointer to **bool** | Default account or not | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Account) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Account) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Account) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Account) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

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

### HasAccountNumber

`func (o *Account) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountName

`func (o *Account) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Account) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Account) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *Account) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetUcmId

`func (o *Account) GetUcmId() string`

GetUcmId returns the UcmId field if non-nil, zero value otherwise.

### GetUcmIdOk

`func (o *Account) GetUcmIdOk() (*string, bool)`

GetUcmIdOk returns a tuple with the UcmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcmId

`func (o *Account) SetUcmId(v string)`

SetUcmId sets UcmId field to given value.

### HasUcmId

`func (o *Account) HasUcmId() bool`

HasUcmId returns a boolean if a field has been set.

### GetGlobalCustId

`func (o *Account) GetGlobalCustId() string`

GetGlobalCustId returns the GlobalCustId field if non-nil, zero value otherwise.

### GetGlobalCustIdOk

`func (o *Account) GetGlobalCustIdOk() (*string, bool)`

GetGlobalCustIdOk returns a tuple with the GlobalCustId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCustId

`func (o *Account) SetGlobalCustId(v string)`

SetGlobalCustId sets GlobalCustId field to given value.

### HasGlobalCustId

`func (o *Account) HasGlobalCustId() bool`

HasGlobalCustId returns a boolean if a field has been set.

### GetOrgId

`func (o *Account) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Account) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Account) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Account) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *Account) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *Account) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *Account) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *Account) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetSubCustomes

`func (o *Account) GetSubCustomes() []Account`

GetSubCustomes returns the SubCustomes field if non-nil, zero value otherwise.

### GetSubCustomesOk

`func (o *Account) GetSubCustomesOk() (*[]Account, bool)`

GetSubCustomesOk returns a tuple with the SubCustomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCustomes

`func (o *Account) SetSubCustomes(v []Account)`

SetSubCustomes sets SubCustomes field to given value.

### HasSubCustomes

`func (o *Account) HasSubCustomes() bool`

HasSubCustomes returns a boolean if a field has been set.

### GetCountryCode

`func (o *Account) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Account) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Account) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Account) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetOperationalUnit

`func (o *Account) GetOperationalUnit() string`

GetOperationalUnit returns the OperationalUnit field if non-nil, zero value otherwise.

### GetOperationalUnitOk

`func (o *Account) GetOperationalUnitOk() (*string, bool)`

GetOperationalUnitOk returns a tuple with the OperationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalUnit

`func (o *Account) SetOperationalUnit(v string)`

SetOperationalUnit sets OperationalUnit field to given value.

### HasOperationalUnit

`func (o *Account) HasOperationalUnit() bool`

HasOperationalUnit returns a boolean if a field has been set.

### GetOperationalUnitMetros

`func (o *Account) GetOperationalUnitMetros() []string`

GetOperationalUnitMetros returns the OperationalUnitMetros field if non-nil, zero value otherwise.

### GetOperationalUnitMetrosOk

`func (o *Account) GetOperationalUnitMetrosOk() (*[]string, bool)`

GetOperationalUnitMetrosOk returns a tuple with the OperationalUnitMetros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalUnitMetros

`func (o *Account) SetOperationalUnitMetros(v []string)`

SetOperationalUnitMetros sets OperationalUnitMetros field to given value.

### HasOperationalUnitMetros

`func (o *Account) HasOperationalUnitMetros() bool`

HasOperationalUnitMetros returns a boolean if a field has been set.

### GetSignatureRequired

`func (o *Account) GetSignatureRequired() bool`

GetSignatureRequired returns the SignatureRequired field if non-nil, zero value otherwise.

### GetSignatureRequiredOk

`func (o *Account) GetSignatureRequiredOk() (*bool, bool)`

GetSignatureRequiredOk returns a tuple with the SignatureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureRequired

`func (o *Account) SetSignatureRequired(v bool)`

SetSignatureRequired sets SignatureRequired field to given value.

### HasSignatureRequired

`func (o *Account) HasSignatureRequired() bool`

HasSignatureRequired returns a boolean if a field has been set.

### GetPoBearing

`func (o *Account) GetPoBearing() bool`

GetPoBearing returns the PoBearing field if non-nil, zero value otherwise.

### GetPoBearingOk

`func (o *Account) GetPoBearingOk() (*bool, bool)`

GetPoBearingOk returns a tuple with the PoBearing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoBearing

`func (o *Account) SetPoBearing(v bool)`

SetPoBearing sets PoBearing field to given value.

### HasPoBearing

`func (o *Account) HasPoBearing() bool`

HasPoBearing returns a boolean if a field has been set.

### GetDefault

`func (o *Account) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Account) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Account) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Account) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


