# RouteFiltersChangeOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**RoutingProtocolChangeOperationOp**](RoutingProtocolChangeOperationOp.md) |  | 
**Path** | **string** | path inside document leading to updated parameter | 
**Value** | [**RouteFiltersBase**](RouteFiltersBase.md) |  | 

## Methods

### NewRouteFiltersChangeOperation

`func NewRouteFiltersChangeOperation(op RoutingProtocolChangeOperationOp, path string, value RouteFiltersBase, ) *RouteFiltersChangeOperation`

NewRouteFiltersChangeOperation instantiates a new RouteFiltersChangeOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFiltersChangeOperationWithDefaults

`func NewRouteFiltersChangeOperationWithDefaults() *RouteFiltersChangeOperation`

NewRouteFiltersChangeOperationWithDefaults instantiates a new RouteFiltersChangeOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *RouteFiltersChangeOperation) GetOp() RoutingProtocolChangeOperationOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RouteFiltersChangeOperation) GetOpOk() (*RoutingProtocolChangeOperationOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RouteFiltersChangeOperation) SetOp(v RoutingProtocolChangeOperationOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *RouteFiltersChangeOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RouteFiltersChangeOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RouteFiltersChangeOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *RouteFiltersChangeOperation) GetValue() RouteFiltersBase`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RouteFiltersChangeOperation) GetValueOk() (*RouteFiltersBase, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RouteFiltersChangeOperation) SetValue(v RouteFiltersBase)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


