# RoutingProtocolChangeOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**RoutingProtocolChangeOperationOp**](RoutingProtocolChangeOperationOp.md) |  | 
**Path** | **string** | path inside document leading to updated parameter | 
**Value** | [**RoutingProtocolBase**](RoutingProtocolBase.md) |  | 

## Methods

### NewRoutingProtocolChangeOperation

`func NewRoutingProtocolChangeOperation(op RoutingProtocolChangeOperationOp, path string, value RoutingProtocolBase, ) *RoutingProtocolChangeOperation`

NewRoutingProtocolChangeOperation instantiates a new RoutingProtocolChangeOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolChangeOperationWithDefaults

`func NewRoutingProtocolChangeOperationWithDefaults() *RoutingProtocolChangeOperation`

NewRoutingProtocolChangeOperationWithDefaults instantiates a new RoutingProtocolChangeOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *RoutingProtocolChangeOperation) GetOp() RoutingProtocolChangeOperationOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RoutingProtocolChangeOperation) GetOpOk() (*RoutingProtocolChangeOperationOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RoutingProtocolChangeOperation) SetOp(v RoutingProtocolChangeOperationOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *RoutingProtocolChangeOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RoutingProtocolChangeOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RoutingProtocolChangeOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *RoutingProtocolChangeOperation) GetValue() RoutingProtocolBase`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RoutingProtocolChangeOperation) GetValueOk() (*RoutingProtocolBase, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RoutingProtocolChangeOperation) SetValue(v RoutingProtocolBase)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


