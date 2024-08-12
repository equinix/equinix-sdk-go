# ServiceTokenChangeOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**ServiceTokenChangeOperationOp**](ServiceTokenChangeOperationOp.md) |  | 
**Path** | **string** | path inside document leading to updated parameter | 
**Value** | **interface{}** | new value for updated parameter | 

## Methods

### NewServiceTokenChangeOperation

`func NewServiceTokenChangeOperation(op ServiceTokenChangeOperationOp, path string, value interface{}, ) *ServiceTokenChangeOperation`

NewServiceTokenChangeOperation instantiates a new ServiceTokenChangeOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTokenChangeOperationWithDefaults

`func NewServiceTokenChangeOperationWithDefaults() *ServiceTokenChangeOperation`

NewServiceTokenChangeOperationWithDefaults instantiates a new ServiceTokenChangeOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *ServiceTokenChangeOperation) GetOp() ServiceTokenChangeOperationOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ServiceTokenChangeOperation) GetOpOk() (*ServiceTokenChangeOperationOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ServiceTokenChangeOperation) SetOp(v ServiceTokenChangeOperationOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *ServiceTokenChangeOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ServiceTokenChangeOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ServiceTokenChangeOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *ServiceTokenChangeOperation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServiceTokenChangeOperation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServiceTokenChangeOperation) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ServiceTokenChangeOperation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ServiceTokenChangeOperation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


