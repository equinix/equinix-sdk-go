# AsnPatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**AsnPatchOperationOp**](AsnPatchOperationOp.md) |  | 
**Path** | **string** | JSON Pointer to the field to update | 
**Value** | **interface{}** | New value for add or replace operations | 

## Methods

### NewAsnPatchOperation

`func NewAsnPatchOperation(op AsnPatchOperationOp, path string, value interface{}, ) *AsnPatchOperation`

NewAsnPatchOperation instantiates a new AsnPatchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsnPatchOperationWithDefaults

`func NewAsnPatchOperationWithDefaults() *AsnPatchOperation`

NewAsnPatchOperationWithDefaults instantiates a new AsnPatchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *AsnPatchOperation) GetOp() AsnPatchOperationOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AsnPatchOperation) GetOpOk() (*AsnPatchOperationOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AsnPatchOperation) SetOp(v AsnPatchOperationOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *AsnPatchOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AsnPatchOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AsnPatchOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *AsnPatchOperation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AsnPatchOperation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AsnPatchOperation) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *AsnPatchOperation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AsnPatchOperation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


