# Cage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cage** | **string** | Cage or Suite | 
**CageTypes** | Pointer to **[]string** | Cage Types | [optional] 
**Accounts** | Pointer to [**[]Account**](Account.md) | List of Accounts | [optional] 

## Methods

### NewCage

`func NewCage(cage string, ) *Cage`

NewCage instantiates a new Cage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCageWithDefaults

`func NewCageWithDefaults() *Cage`

NewCageWithDefaults instantiates a new Cage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCage

`func (o *Cage) GetCage() string`

GetCage returns the Cage field if non-nil, zero value otherwise.

### GetCageOk

`func (o *Cage) GetCageOk() (*string, bool)`

GetCageOk returns a tuple with the Cage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCage

`func (o *Cage) SetCage(v string)`

SetCage sets Cage field to given value.


### GetCageTypes

`func (o *Cage) GetCageTypes() []string`

GetCageTypes returns the CageTypes field if non-nil, zero value otherwise.

### GetCageTypesOk

`func (o *Cage) GetCageTypesOk() (*[]string, bool)`

GetCageTypesOk returns a tuple with the CageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCageTypes

`func (o *Cage) SetCageTypes(v []string)`

SetCageTypes sets CageTypes field to given value.

### HasCageTypes

`func (o *Cage) HasCageTypes() bool`

HasCageTypes returns a boolean if a field has been set.

### GetAccounts

`func (o *Cage) GetAccounts() []Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Cage) GetAccountsOk() (*[]Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Cage) SetAccounts(v []Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Cage) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


