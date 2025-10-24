# PricingSiebelConfigPrimary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency of the charges. | [optional] 
**Charges** | Pointer to [**[]Charges**](Charges.md) |  | [optional] 

## Methods

### NewPricingSiebelConfigPrimary

`func NewPricingSiebelConfigPrimary() *PricingSiebelConfigPrimary`

NewPricingSiebelConfigPrimary instantiates a new PricingSiebelConfigPrimary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingSiebelConfigPrimaryWithDefaults

`func NewPricingSiebelConfigPrimaryWithDefaults() *PricingSiebelConfigPrimary`

NewPricingSiebelConfigPrimaryWithDefaults instantiates a new PricingSiebelConfigPrimary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PricingSiebelConfigPrimary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PricingSiebelConfigPrimary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PricingSiebelConfigPrimary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PricingSiebelConfigPrimary) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCharges

`func (o *PricingSiebelConfigPrimary) GetCharges() []Charges`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *PricingSiebelConfigPrimary) GetChargesOk() (*[]Charges, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *PricingSiebelConfigPrimary) SetCharges(v []Charges)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *PricingSiebelConfigPrimary) HasCharges() bool`

HasCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


