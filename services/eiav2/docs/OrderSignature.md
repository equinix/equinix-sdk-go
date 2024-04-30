# OrderSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signatory** | [**OrderSignatureSignatory**](OrderSignatureSignatory.md) |  | 
**Delegate** | Pointer to [**OrderSignatureDelegate**](OrderSignatureDelegate.md) |  | [optional] 

## Methods

### NewOrderSignature

`func NewOrderSignature(signatory OrderSignatureSignatory, ) *OrderSignature`

NewOrderSignature instantiates a new OrderSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSignatureWithDefaults

`func NewOrderSignatureWithDefaults() *OrderSignature`

NewOrderSignatureWithDefaults instantiates a new OrderSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignatory

`func (o *OrderSignature) GetSignatory() OrderSignatureSignatory`

GetSignatory returns the Signatory field if non-nil, zero value otherwise.

### GetSignatoryOk

`func (o *OrderSignature) GetSignatoryOk() (*OrderSignatureSignatory, bool)`

GetSignatoryOk returns a tuple with the Signatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatory

`func (o *OrderSignature) SetSignatory(v OrderSignatureSignatory)`

SetSignatory sets Signatory field to given value.


### GetDelegate

`func (o *OrderSignature) GetDelegate() OrderSignatureDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *OrderSignature) GetDelegateOk() (*OrderSignatureDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *OrderSignature) SetDelegate(v OrderSignatureDelegate)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *OrderSignature) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


