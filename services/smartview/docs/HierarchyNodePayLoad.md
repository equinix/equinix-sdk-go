# HierarchyNodePayLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cages** | Pointer to [**[]Cages**](Cages.md) |  | [optional] 
**Circuits** | Pointer to [**[]CircuitsMapWithCage**](CircuitsMapWithCage.md) |  | [optional] 

## Methods

### NewHierarchyNodePayLoad

`func NewHierarchyNodePayLoad() *HierarchyNodePayLoad`

NewHierarchyNodePayLoad instantiates a new HierarchyNodePayLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHierarchyNodePayLoadWithDefaults

`func NewHierarchyNodePayLoadWithDefaults() *HierarchyNodePayLoad`

NewHierarchyNodePayLoadWithDefaults instantiates a new HierarchyNodePayLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCages

`func (o *HierarchyNodePayLoad) GetCages() []Cages`

GetCages returns the Cages field if non-nil, zero value otherwise.

### GetCagesOk

`func (o *HierarchyNodePayLoad) GetCagesOk() (*[]Cages, bool)`

GetCagesOk returns a tuple with the Cages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCages

`func (o *HierarchyNodePayLoad) SetCages(v []Cages)`

SetCages sets Cages field to given value.

### HasCages

`func (o *HierarchyNodePayLoad) HasCages() bool`

HasCages returns a boolean if a field has been set.

### GetCircuits

`func (o *HierarchyNodePayLoad) GetCircuits() []CircuitsMapWithCage`

GetCircuits returns the Circuits field if non-nil, zero value otherwise.

### GetCircuitsOk

`func (o *HierarchyNodePayLoad) GetCircuitsOk() (*[]CircuitsMapWithCage, bool)`

GetCircuitsOk returns a tuple with the Circuits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuits

`func (o *HierarchyNodePayLoad) SetCircuits(v []CircuitsMapWithCage)`

SetCircuits sets Circuits field to given value.

### HasCircuits

`func (o *HierarchyNodePayLoad) HasCircuits() bool`

HasCircuits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


