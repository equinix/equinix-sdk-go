# PrecisionTimeChangeOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**PrecisionTimeChangeOperationOp**](PrecisionTimeChangeOperationOp.md) |  | 
**Path** | [**PrecisionTimeChangeOperationPath**](PrecisionTimeChangeOperationPath.md) |  | 
**Value** | **map[string]interface{}** | new value for updated parameter | 

## Methods

### NewPrecisionTimeChangeOperation

`func NewPrecisionTimeChangeOperation(op PrecisionTimeChangeOperationOp, path PrecisionTimeChangeOperationPath, value map[string]interface{}, ) *PrecisionTimeChangeOperation`

NewPrecisionTimeChangeOperation instantiates a new PrecisionTimeChangeOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrecisionTimeChangeOperationWithDefaults

`func NewPrecisionTimeChangeOperationWithDefaults() *PrecisionTimeChangeOperation`

NewPrecisionTimeChangeOperationWithDefaults instantiates a new PrecisionTimeChangeOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *PrecisionTimeChangeOperation) GetOp() PrecisionTimeChangeOperationOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *PrecisionTimeChangeOperation) GetOpOk() (*PrecisionTimeChangeOperationOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *PrecisionTimeChangeOperation) SetOp(v PrecisionTimeChangeOperationOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *PrecisionTimeChangeOperation) GetPath() PrecisionTimeChangeOperationPath`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PrecisionTimeChangeOperation) GetPathOk() (*PrecisionTimeChangeOperationPath, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PrecisionTimeChangeOperation) SetPath(v PrecisionTimeChangeOperationPath)`

SetPath sets Path field to given value.


### GetValue

`func (o *PrecisionTimeChangeOperation) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PrecisionTimeChangeOperation) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PrecisionTimeChangeOperation) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


