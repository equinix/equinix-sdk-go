# ValueWithUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **map[string]interface{}** | Unit for value. | 
**Value** | **map[string]interface{}** | Specific value, to be read together with unit. | 

## Methods

### NewValueWithUnit

`func NewValueWithUnit(unit map[string]interface{}, value map[string]interface{}, ) *ValueWithUnit`

NewValueWithUnit instantiates a new ValueWithUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueWithUnitWithDefaults

`func NewValueWithUnitWithDefaults() *ValueWithUnit`

NewValueWithUnitWithDefaults instantiates a new ValueWithUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *ValueWithUnit) GetUnit() map[string]interface{}`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ValueWithUnit) GetUnitOk() (*map[string]interface{}, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ValueWithUnit) SetUnit(v map[string]interface{})`

SetUnit sets Unit field to given value.


### GetValue

`func (o *ValueWithUnit) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ValueWithUnit) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ValueWithUnit) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


