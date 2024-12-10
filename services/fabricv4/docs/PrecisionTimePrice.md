# PrecisionTimePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | offering price currency | [optional] 
**Charges** | Pointer to [**[]PriceCharge**](PriceCharge.md) |  | [optional] 

## Methods

### NewPrecisionTimePrice

`func NewPrecisionTimePrice() *PrecisionTimePrice`

NewPrecisionTimePrice instantiates a new PrecisionTimePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrecisionTimePriceWithDefaults

`func NewPrecisionTimePriceWithDefaults() *PrecisionTimePrice`

NewPrecisionTimePriceWithDefaults instantiates a new PrecisionTimePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PrecisionTimePrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PrecisionTimePrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PrecisionTimePrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PrecisionTimePrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCharges

`func (o *PrecisionTimePrice) GetCharges() []PriceCharge`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *PrecisionTimePrice) GetChargesOk() (*[]PriceCharge, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *PrecisionTimePrice) SetCharges(v []PriceCharge)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *PrecisionTimePrice) HasCharges() bool`

HasCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


