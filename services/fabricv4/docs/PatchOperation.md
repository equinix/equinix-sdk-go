# PatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**PatchOperationOp**](PatchOperationOp.md) |  | 
**Path** | **string** | JSON Pointer path to the field to modify (e.g., /name, /description, /summary, /webUrl, /notifications) | 
**Value** | **map[string]interface{}** | The value to update the field to | 

## Methods

### NewPatchOperation

`func NewPatchOperation(op PatchOperationOp, path string, value map[string]interface{}, ) *PatchOperation`

NewPatchOperation instantiates a new PatchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchOperationWithDefaults

`func NewPatchOperationWithDefaults() *PatchOperation`

NewPatchOperationWithDefaults instantiates a new PatchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *PatchOperation) GetOp() PatchOperationOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *PatchOperation) GetOpOk() (*PatchOperationOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *PatchOperation) SetOp(v PatchOperationOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *PatchOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *PatchOperation) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchOperation) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchOperation) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


