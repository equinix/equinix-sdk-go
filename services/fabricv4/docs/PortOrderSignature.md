# PortOrderSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signatory** | Pointer to [**PortOrderSignatureSignatory**](PortOrderSignatureSignatory.md) |  | [optional] 
**Delegate** | Pointer to [**PortOrderSignatureDelegate**](PortOrderSignatureDelegate.md) |  | [optional] 

## Methods

### NewPortOrderSignature

`func NewPortOrderSignature() *PortOrderSignature`

NewPortOrderSignature instantiates a new PortOrderSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortOrderSignatureWithDefaults

`func NewPortOrderSignatureWithDefaults() *PortOrderSignature`

NewPortOrderSignatureWithDefaults instantiates a new PortOrderSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignatory

`func (o *PortOrderSignature) GetSignatory() PortOrderSignatureSignatory`

GetSignatory returns the Signatory field if non-nil, zero value otherwise.

### GetSignatoryOk

`func (o *PortOrderSignature) GetSignatoryOk() (*PortOrderSignatureSignatory, bool)`

GetSignatoryOk returns a tuple with the Signatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatory

`func (o *PortOrderSignature) SetSignatory(v PortOrderSignatureSignatory)`

SetSignatory sets Signatory field to given value.

### HasSignatory

`func (o *PortOrderSignature) HasSignatory() bool`

HasSignatory returns a boolean if a field has been set.

### GetDelegate

`func (o *PortOrderSignature) GetDelegate() PortOrderSignatureDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *PortOrderSignature) GetDelegateOk() (*PortOrderSignatureDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *PortOrderSignature) SetDelegate(v PortOrderSignatureDelegate)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *PortOrderSignature) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


