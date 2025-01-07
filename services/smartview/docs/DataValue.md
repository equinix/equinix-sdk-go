# DataValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | timestamp since epoch. reading timestamp | [optional] 
**Modifiers** | Pointer to **[]string** | modifiers applicable for the reading value | [optional] 
**Value** | Pointer to **string** | reading value for the datapoint | [optional] 

## Methods

### NewDataValue

`func NewDataValue() *DataValue`

NewDataValue instantiates a new DataValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataValueWithDefaults

`func NewDataValueWithDefaults() *DataValue`

NewDataValueWithDefaults instantiates a new DataValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *DataValue) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *DataValue) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *DataValue) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *DataValue) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetModifiers

`func (o *DataValue) GetModifiers() []string`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *DataValue) GetModifiersOk() (*[]string, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *DataValue) SetModifiers(v []string)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *DataValue) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetValue

`func (o *DataValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DataValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


