# RouteFiltersPatchRequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** | Handy shortcut for operation name | 
**Path** | **string** | path to change | 
**Value** | **interface{}** | new value for updated parameter | 

## Methods

### NewRouteFiltersPatchRequestItem

`func NewRouteFiltersPatchRequestItem(op string, path string, value interface{}, ) *RouteFiltersPatchRequestItem`

NewRouteFiltersPatchRequestItem instantiates a new RouteFiltersPatchRequestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFiltersPatchRequestItemWithDefaults

`func NewRouteFiltersPatchRequestItemWithDefaults() *RouteFiltersPatchRequestItem`

NewRouteFiltersPatchRequestItemWithDefaults instantiates a new RouteFiltersPatchRequestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *RouteFiltersPatchRequestItem) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RouteFiltersPatchRequestItem) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RouteFiltersPatchRequestItem) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *RouteFiltersPatchRequestItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RouteFiltersPatchRequestItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RouteFiltersPatchRequestItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *RouteFiltersPatchRequestItem) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RouteFiltersPatchRequestItem) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RouteFiltersPatchRequestItem) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *RouteFiltersPatchRequestItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *RouteFiltersPatchRequestItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


