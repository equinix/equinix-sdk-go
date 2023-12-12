# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**Balance** | Pointer to **float32** |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**CreditAmount** | Pointer to **float32** |  | [optional] 
**CreditsApplied** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DueOn** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]LineItem**](LineItem.md) |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**ProjectIdName**](ProjectIdName.md) |  | [optional] 
**ReferenceNumber** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TargetDate** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Invoice) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Invoice) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Invoice) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Invoice) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalance

`func (o *Invoice) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Invoice) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Invoice) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *Invoice) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Invoice) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Invoice) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Invoice) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Invoice) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCreditAmount

`func (o *Invoice) GetCreditAmount() float32`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *Invoice) GetCreditAmountOk() (*float32, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *Invoice) SetCreditAmount(v float32)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *Invoice) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetCreditsApplied

`func (o *Invoice) GetCreditsApplied() float32`

GetCreditsApplied returns the CreditsApplied field if non-nil, zero value otherwise.

### GetCreditsAppliedOk

`func (o *Invoice) GetCreditsAppliedOk() (*float32, bool)`

GetCreditsAppliedOk returns a tuple with the CreditsApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsApplied

`func (o *Invoice) SetCreditsApplied(v float32)`

SetCreditsApplied sets CreditsApplied field to given value.

### HasCreditsApplied

`func (o *Invoice) HasCreditsApplied() bool`

HasCreditsApplied returns a boolean if a field has been set.

### GetCurrency

`func (o *Invoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Invoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDueOn

`func (o *Invoice) GetDueOn() string`

GetDueOn returns the DueOn field if non-nil, zero value otherwise.

### GetDueOnOk

`func (o *Invoice) GetDueOnOk() (*string, bool)`

GetDueOnOk returns a tuple with the DueOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueOn

`func (o *Invoice) SetDueOn(v string)`

SetDueOn sets DueOn field to given value.

### HasDueOn

`func (o *Invoice) HasDueOn() bool`

HasDueOn returns a boolean if a field has been set.

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *Invoice) GetItems() []LineItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Invoice) GetItemsOk() (*[]LineItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Invoice) SetItems(v []LineItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Invoice) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNumber

`func (o *Invoice) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Invoice) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Invoice) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Invoice) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetProject

`func (o *Invoice) GetProject() ProjectIdName`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Invoice) GetProjectOk() (*ProjectIdName, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Invoice) SetProject(v ProjectIdName)`

SetProject sets Project field to given value.

### HasProject

`func (o *Invoice) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReferenceNumber

`func (o *Invoice) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *Invoice) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *Invoice) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *Invoice) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### GetStatus

`func (o *Invoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Invoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetDate

`func (o *Invoice) GetTargetDate() string`

GetTargetDate returns the TargetDate field if non-nil, zero value otherwise.

### GetTargetDateOk

`func (o *Invoice) GetTargetDateOk() (*string, bool)`

GetTargetDateOk returns a tuple with the TargetDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDate

`func (o *Invoice) SetTargetDate(v string)`

SetTargetDate sets TargetDate field to given value.

### HasTargetDate

`func (o *Invoice) HasTargetDate() bool`

HasTargetDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


