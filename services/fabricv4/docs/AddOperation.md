# AddOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**OpEnum**](OpEnum.md) |  | 
**Path** | **string** | A JSON Pointer path. | 
**Value** | **map[string]interface{}** | value to add | 

## Methods

### NewAddOperation

`func NewAddOperation(op OpEnum, path string, value map[string]interface{}, ) *AddOperation`

NewAddOperation instantiates a new AddOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOperationWithDefaults

`func NewAddOperationWithDefaults() *AddOperation`

NewAddOperationWithDefaults instantiates a new AddOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *AddOperation) GetOp() OpEnum`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AddOperation) GetOpOk() (*OpEnum, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AddOperation) SetOp(v OpEnum)`

SetOp sets Op field to given value.


### GetPath

`func (o *AddOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AddOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AddOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *AddOperation) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddOperation) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddOperation) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


