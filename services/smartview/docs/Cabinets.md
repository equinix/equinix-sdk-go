# Cabinets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Circuits** | Pointer to [**[]Circuits**](Circuits.md) |  | [optional] 
**Name** | Pointer to [**CabinetsName**](CabinetsName.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCabinets

`func NewCabinets() *Cabinets`

NewCabinets instantiates a new Cabinets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCabinetsWithDefaults

`func NewCabinetsWithDefaults() *Cabinets`

NewCabinetsWithDefaults instantiates a new Cabinets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuits

`func (o *Cabinets) GetCircuits() []Circuits`

GetCircuits returns the Circuits field if non-nil, zero value otherwise.

### GetCircuitsOk

`func (o *Cabinets) GetCircuitsOk() (*[]Circuits, bool)`

GetCircuitsOk returns a tuple with the Circuits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuits

`func (o *Cabinets) SetCircuits(v []Circuits)`

SetCircuits sets Circuits field to given value.

### HasCircuits

`func (o *Cabinets) HasCircuits() bool`

HasCircuits returns a boolean if a field has been set.

### GetName

`func (o *Cabinets) GetName() CabinetsName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cabinets) GetNameOk() (*CabinetsName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cabinets) SetName(v CabinetsName)`

SetName sets Name field to given value.

### HasName

`func (o *Cabinets) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Cabinets) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cabinets) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cabinets) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Cabinets) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


