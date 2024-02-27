# SubInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name of subinterafce of a port | [optional] 
**Unit** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubInterface

`func NewSubInterface() *SubInterface`

NewSubInterface instantiates a new SubInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubInterfaceWithDefaults

`func NewSubInterfaceWithDefaults() *SubInterface`

NewSubInterfaceWithDefaults instantiates a new SubInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SubInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnit

`func (o *SubInterface) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *SubInterface) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *SubInterface) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *SubInterface) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


