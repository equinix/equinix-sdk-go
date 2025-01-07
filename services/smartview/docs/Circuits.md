# Circuits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**CircuitsName**](CircuitsName.md) |  | [optional] 
**Type** | Pointer to [**CircuitsType**](CircuitsType.md) |  | [optional] 

## Methods

### NewCircuits

`func NewCircuits() *Circuits`

NewCircuits instantiates a new Circuits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitsWithDefaults

`func NewCircuitsWithDefaults() *Circuits`

NewCircuitsWithDefaults instantiates a new Circuits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Circuits) GetName() CircuitsName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Circuits) GetNameOk() (*CircuitsName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Circuits) SetName(v CircuitsName)`

SetName sets Name field to given value.

### HasName

`func (o *Circuits) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Circuits) GetType() CircuitsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Circuits) GetTypeOk() (*CircuitsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Circuits) SetType(v CircuitsType)`

SetType sets Type field to given value.

### HasType

`func (o *Circuits) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


