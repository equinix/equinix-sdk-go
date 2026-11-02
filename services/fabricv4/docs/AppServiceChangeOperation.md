# AppServiceChangeOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**AppDomainChangeOperationOp**](AppDomainChangeOperationOp.md) |  | 
**Path** | [**AppServiceChangeOperationPath**](AppServiceChangeOperationPath.md) |  | 
**Value** | **interface{}** | new value for updated parameter | 

## Methods

### NewAppServiceChangeOperation

`func NewAppServiceChangeOperation(op AppDomainChangeOperationOp, path AppServiceChangeOperationPath, value interface{}, ) *AppServiceChangeOperation`

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

`func (o *AppServiceChangeOperation) GetOp() AppDomainChangeOperationOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AppServiceChangeOperation) GetOpOk() (*AppDomainChangeOperationOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AppServiceChangeOperation) SetOp(v AppDomainChangeOperationOp)`

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

`func (o *AppServiceChangeOperation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppServiceChangeOperation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppServiceChangeOperation) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *AppServiceChangeOperation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AppServiceChangeOperation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


