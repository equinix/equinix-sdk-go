# PaymentMethodCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Nonce** | **string** |  | 

## Methods

### NewPaymentMethodCreateInput

`func NewPaymentMethodCreateInput(name string, nonce string, ) *PaymentMethodCreateInput`

NewPaymentMethodCreateInput instantiates a new PaymentMethodCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCreateInputWithDefaults

`func NewPaymentMethodCreateInputWithDefaults() *PaymentMethodCreateInput`

NewPaymentMethodCreateInputWithDefaults instantiates a new PaymentMethodCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *PaymentMethodCreateInput) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PaymentMethodCreateInput) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PaymentMethodCreateInput) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PaymentMethodCreateInput) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetName

`func (o *PaymentMethodCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentMethodCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentMethodCreateInput) SetName(v string)`

SetName sets Name field to given value.


### GetNonce

`func (o *PaymentMethodCreateInput) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *PaymentMethodCreateInput) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *PaymentMethodCreateInput) SetNonce(v string)`

SetNonce sets Nonce field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


