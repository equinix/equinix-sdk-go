# Dimension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **int32** | Dimension Value | 
**Unit** | [**DimensionUnit**](DimensionUnit.md) |  | 

## Methods

### NewDimension

`func NewDimension(value int32, unit DimensionUnit, ) *Dimension`

NewDimension instantiates a new Dimension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDimensionWithDefaults

`func NewDimensionWithDefaults() *Dimension`

NewDimensionWithDefaults instantiates a new Dimension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Dimension) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Dimension) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Dimension) SetValue(v int32)`

SetValue sets Value field to given value.


### GetUnit

`func (o *Dimension) GetUnit() DimensionUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Dimension) GetUnitOk() (*DimensionUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Dimension) SetUnit(v DimensionUnit)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


