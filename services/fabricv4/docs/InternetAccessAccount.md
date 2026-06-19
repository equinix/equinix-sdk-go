# InternetAccessAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Account number | 
**Href** | Pointer to **string** | Account URL path | [optional] 

## Methods

### NewInternetAccessAccount

`func NewInternetAccessAccount(accountNumber string, ) *InternetAccessAccount`

NewInternetAccessAccount instantiates a new InternetAccessAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessAccountWithDefaults

`func NewInternetAccessAccountWithDefaults() *InternetAccessAccount`

NewInternetAccessAccountWithDefaults instantiates a new InternetAccessAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *InternetAccessAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *InternetAccessAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *InternetAccessAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetHref

`func (o *InternetAccessAccount) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *InternetAccessAccount) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *InternetAccessAccount) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *InternetAccessAccount) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


