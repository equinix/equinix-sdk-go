# OrderSignatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signatory** | [**OrderSignatureSignatory**](OrderSignatureSignatory.md) |  | 
**Delegate** | Pointer to [**OrderSignatureDelegateRequest**](OrderSignatureDelegateRequest.md) |  | [optional] 

## Methods

### NewOrderSignatureRequest

`func NewOrderSignatureRequest(signatory OrderSignatureSignatory, ) *OrderSignatureRequest`

NewOrderSignatureRequest instantiates a new OrderSignatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSignatureRequestWithDefaults

`func NewOrderSignatureRequestWithDefaults() *OrderSignatureRequest`

NewOrderSignatureRequestWithDefaults instantiates a new OrderSignatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignatory

`func (o *OrderSignatureRequest) GetSignatory() OrderSignatureSignatory`

GetSignatory returns the Signatory field if non-nil, zero value otherwise.

### GetSignatoryOk

`func (o *OrderSignatureRequest) GetSignatoryOk() (*OrderSignatureSignatory, bool)`

GetSignatoryOk returns a tuple with the Signatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatory

`func (o *OrderSignatureRequest) SetSignatory(v OrderSignatureSignatory)`

SetSignatory sets Signatory field to given value.


### GetDelegate

`func (o *OrderSignatureRequest) GetDelegate() OrderSignatureDelegateRequest`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *OrderSignatureRequest) GetDelegateOk() (*OrderSignatureDelegateRequest, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *OrderSignatureRequest) SetDelegate(v OrderSignatureDelegateRequest)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *OrderSignatureRequest) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


