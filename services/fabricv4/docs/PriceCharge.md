# PriceCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**PriceChargeType**](PriceChargeType.md) |  | [optional] 
**Price** | Pointer to **float64** | Offering price | [optional] 

## Methods

### NewPriceCharge

`func NewPriceCharge() *PriceCharge`

NewPriceCharge instantiates a new PriceCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceChargeWithDefaults

`func NewPriceChargeWithDefaults() *PriceCharge`

NewPriceChargeWithDefaults instantiates a new PriceCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceCharge) GetType() PriceChargeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceCharge) GetTypeOk() (*PriceChargeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceCharge) SetType(v PriceChargeType)`

SetType sets Type field to given value.

### HasType

`func (o *PriceCharge) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrice

`func (o *PriceCharge) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PriceCharge) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PriceCharge) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PriceCharge) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


