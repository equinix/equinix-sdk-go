# CompositePriceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | Pointer to [**PriceResponse**](PriceResponse.md) |  | [optional] 
**Secondary** | Pointer to [**PriceResponse**](PriceResponse.md) |  | [optional] 
**TermLength** | Pointer to **string** |  | [optional] 

## Methods

### NewCompositePriceResponse

`func NewCompositePriceResponse() *CompositePriceResponse`

NewCompositePriceResponse instantiates a new CompositePriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositePriceResponseWithDefaults

`func NewCompositePriceResponseWithDefaults() *CompositePriceResponse`

NewCompositePriceResponseWithDefaults instantiates a new CompositePriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *CompositePriceResponse) GetPrimary() PriceResponse`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *CompositePriceResponse) GetPrimaryOk() (*PriceResponse, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *CompositePriceResponse) SetPrimary(v PriceResponse)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *CompositePriceResponse) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetSecondary

`func (o *CompositePriceResponse) GetSecondary() PriceResponse`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *CompositePriceResponse) GetSecondaryOk() (*PriceResponse, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *CompositePriceResponse) SetSecondary(v PriceResponse)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *CompositePriceResponse) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### GetTermLength

`func (o *CompositePriceResponse) GetTermLength() string`

GetTermLength returns the TermLength field if non-nil, zero value otherwise.

### GetTermLengthOk

`func (o *CompositePriceResponse) GetTermLengthOk() (*string, bool)`

GetTermLengthOk returns a tuple with the TermLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermLength

`func (o *CompositePriceResponse) SetTermLength(v string)`

SetTermLength sets TermLength field to given value.

### HasTermLength

`func (o *CompositePriceResponse) HasTermLength() bool`

HasTermLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


