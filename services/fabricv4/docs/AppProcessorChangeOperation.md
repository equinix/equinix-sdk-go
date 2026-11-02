# AppProcessorChangeOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**AppDomainChangeOperationOp**](AppDomainChangeOperationOp.md) |  | 
**Path** | [**AppProcessorChangeOperationPath**](AppProcessorChangeOperationPath.md) |  | 
**Value** | **interface{}** | new value for updated parameter | 

## Methods

### NewAppProcessorChangeOperation

`func NewAppProcessorChangeOperation(op AppDomainChangeOperationOp, path AppProcessorChangeOperationPath, value interface{}, ) *AppProcessorChangeOperation`

NewAppProcessorChangeOperation instantiates a new AppProcessorChangeOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppProcessorChangeOperationWithDefaults

`func NewAppProcessorChangeOperationWithDefaults() *AppProcessorChangeOperation`

NewAppProcessorChangeOperationWithDefaults instantiates a new AppProcessorChangeOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *AppProcessorChangeOperation) GetOp() AppDomainChangeOperationOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AppProcessorChangeOperation) GetOpOk() (*AppDomainChangeOperationOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AppProcessorChangeOperation) SetOp(v AppDomainChangeOperationOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *AppProcessorChangeOperation) GetPath() AppProcessorChangeOperationPath`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppProcessorChangeOperation) GetPathOk() (*AppProcessorChangeOperationPath, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppProcessorChangeOperation) SetPath(v AppProcessorChangeOperationPath)`

SetPath sets Path field to given value.


### GetValue

`func (o *AppProcessorChangeOperation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppProcessorChangeOperation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppProcessorChangeOperation) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *AppProcessorChangeOperation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AppProcessorChangeOperation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


