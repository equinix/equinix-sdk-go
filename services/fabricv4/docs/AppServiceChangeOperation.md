# AppServiceChangeOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**PrecisionTimeChangeOperationOp**](PrecisionTimeChangeOperationOp.md) |  | 
**Path** | [**AppServiceChangeOperationPath**](AppServiceChangeOperationPath.md) |  | 
**Value** | **map[string]interface{}** | new value for updated parameter | 

## Methods

### NewAppServiceChangeOperation

`func NewAppServiceChangeOperation(op PrecisionTimeChangeOperationOp, path AppServiceChangeOperationPath, value map[string]interface{}, ) *AppServiceChangeOperation`

NewAppServiceChangeOperation instantiates a new AppServiceChangeOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceChangeOperationWithDefaults

`func NewAppServiceChangeOperationWithDefaults() *AppServiceChangeOperation`

NewAppServiceChangeOperationWithDefaults instantiates a new AppServiceChangeOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *AppServiceChangeOperation) GetOp() PrecisionTimeChangeOperationOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AppServiceChangeOperation) GetOpOk() (*PrecisionTimeChangeOperationOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AppServiceChangeOperation) SetOp(v PrecisionTimeChangeOperationOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *AppServiceChangeOperation) GetPath() AppServiceChangeOperationPath`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppServiceChangeOperation) GetPathOk() (*AppServiceChangeOperationPath, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppServiceChangeOperation) SetPath(v AppServiceChangeOperationPath)`

SetPath sets Path field to given value.


### GetValue

`func (o *AppServiceChangeOperation) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppServiceChangeOperation) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppServiceChangeOperation) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


