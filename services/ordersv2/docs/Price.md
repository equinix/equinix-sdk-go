# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **float32** | To be used with &#x60;valueType&#x60; and &#x60;type&#x60;. Can be used to denote price or discount for an order or service | [optional] 
**ValueType** | Pointer to [**PriceValueType**](PriceValueType.md) |  | [optional] 
**Type** | Pointer to [**PriceType**](PriceType.md) |  | [optional] 

## Methods

### NewPrice

`func NewPrice() *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Price) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Price) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Price) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Price) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *Price) GetValueType() PriceValueType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *Price) GetValueTypeOk() (*PriceValueType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *Price) SetValueType(v PriceValueType)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *Price) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetType

`func (o *Price) GetType() PriceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Price) GetTypeOk() (*PriceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Price) SetType(v PriceType)`

SetType sets Type field to given value.

### HasType

`func (o *Price) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


