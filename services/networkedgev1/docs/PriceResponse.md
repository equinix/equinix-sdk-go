# PriceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCommencementDate** | Pointer to **string** |  | [optional] 
**BillingEnabled** | Pointer to **bool** |  | [optional] 
**Charges** | Pointer to [**[]Charges**](Charges.md) | An array of the monthly recurring charges. | [optional] 
**Currency** | Pointer to **string** |  | [optional] 

## Methods

### NewPriceResponse

`func NewPriceResponse() *PriceResponse`

NewPriceResponse instantiates a new PriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceResponseWithDefaults

`func NewPriceResponseWithDefaults() *PriceResponse`

NewPriceResponseWithDefaults instantiates a new PriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCommencementDate

`func (o *PriceResponse) GetBillingCommencementDate() string`

GetBillingCommencementDate returns the BillingCommencementDate field if non-nil, zero value otherwise.

### GetBillingCommencementDateOk

`func (o *PriceResponse) GetBillingCommencementDateOk() (*string, bool)`

GetBillingCommencementDateOk returns a tuple with the BillingCommencementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCommencementDate

`func (o *PriceResponse) SetBillingCommencementDate(v string)`

SetBillingCommencementDate sets BillingCommencementDate field to given value.

### HasBillingCommencementDate

`func (o *PriceResponse) HasBillingCommencementDate() bool`

HasBillingCommencementDate returns a boolean if a field has been set.

### GetBillingEnabled

`func (o *PriceResponse) GetBillingEnabled() bool`

GetBillingEnabled returns the BillingEnabled field if non-nil, zero value otherwise.

### GetBillingEnabledOk

`func (o *PriceResponse) GetBillingEnabledOk() (*bool, bool)`

GetBillingEnabledOk returns a tuple with the BillingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEnabled

`func (o *PriceResponse) SetBillingEnabled(v bool)`

SetBillingEnabled sets BillingEnabled field to given value.

### HasBillingEnabled

`func (o *PriceResponse) HasBillingEnabled() bool`

HasBillingEnabled returns a boolean if a field has been set.

### GetCharges

`func (o *PriceResponse) GetCharges() []Charges`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *PriceResponse) GetChargesOk() (*[]Charges, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *PriceResponse) SetCharges(v []Charges)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *PriceResponse) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetCurrency

`func (o *PriceResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PriceResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PriceResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PriceResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


