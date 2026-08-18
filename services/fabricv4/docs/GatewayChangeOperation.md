# GatewayChangeOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**PrecisionTimeChangeOperationOp**](PrecisionTimeChangeOperationOp.md) |  | 
**Path** | **string** | path inside document leading to updated parameter | 
**Value** | **map[string]interface{}** | new value for updated parameter | 

## Methods

### NewGatewayChangeOperation

`func NewGatewayChangeOperation(op PrecisionTimeChangeOperationOp, path string, value map[string]interface{}, ) *GatewayChangeOperation`

NewGatewayChangeOperation instantiates a new GatewayChangeOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayChangeOperationWithDefaults

`func NewGatewayChangeOperationWithDefaults() *GatewayChangeOperation`

NewGatewayChangeOperationWithDefaults instantiates a new GatewayChangeOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *GatewayChangeOperation) GetOp() PrecisionTimeChangeOperationOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *GatewayChangeOperation) GetOpOk() (*PrecisionTimeChangeOperationOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *GatewayChangeOperation) SetOp(v PrecisionTimeChangeOperationOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *GatewayChangeOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GatewayChangeOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GatewayChangeOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *GatewayChangeOperation) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GatewayChangeOperation) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GatewayChangeOperation) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


