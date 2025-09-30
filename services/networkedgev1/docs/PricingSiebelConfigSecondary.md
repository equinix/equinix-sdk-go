# PricingSiebelConfigSecondary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency of the charges. | [optional] 
**Charges** | Pointer to [**[]Charges**](Charges.md) |  | [optional] 

## Methods

### NewPricingSiebelConfigSecondary

`func NewPricingSiebelConfigSecondary() *PricingSiebelConfigSecondary`

NewPricingSiebelConfigSecondary instantiates a new PricingSiebelConfigSecondary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingSiebelConfigSecondaryWithDefaults

`func NewPricingSiebelConfigSecondaryWithDefaults() *PricingSiebelConfigSecondary`

NewPricingSiebelConfigSecondaryWithDefaults instantiates a new PricingSiebelConfigSecondary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PricingSiebelConfigSecondary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PricingSiebelConfigSecondary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PricingSiebelConfigSecondary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PricingSiebelConfigSecondary) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCharges

`func (o *PricingSiebelConfigSecondary) GetCharges() []Charges`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *PricingSiebelConfigSecondary) GetChargesOk() (*[]Charges, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *PricingSiebelConfigSecondary) SetCharges(v []Charges)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *PricingSiebelConfigSecondary) HasCharges() bool`

HasCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


