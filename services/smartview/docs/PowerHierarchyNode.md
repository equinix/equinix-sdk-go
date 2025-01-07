# PowerHierarchyNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]PowerHierarchyNode**](PowerHierarchyNode.md) | ibx, cage, cabinet nodes can have cage, cabinet and circuit,  circuit nodes as children respectively.  | [optional] 
**Label** | Pointer to **string** | ibx code, cage unique space id, cabinet unique space id, circuit label for levelType ibx, cage, cabinet, circuit resp.  | [optional] 
**LevelType** | Pointer to [**PowerHierarchyNodeLevelType**](PowerHierarchyNodeLevelType.md) |  | [optional] 
**LevelValue** | Pointer to **string** | ibx code, cage us id, cabinet us id, circuit number for  levelType ibx, cage, cabinet, circuit resp.  | [optional] 

## Methods

### NewPowerHierarchyNode

`func NewPowerHierarchyNode() *PowerHierarchyNode`

NewPowerHierarchyNode instantiates a new PowerHierarchyNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerHierarchyNodeWithDefaults

`func NewPowerHierarchyNodeWithDefaults() *PowerHierarchyNode`

NewPowerHierarchyNodeWithDefaults instantiates a new PowerHierarchyNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *PowerHierarchyNode) GetChildren() []PowerHierarchyNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *PowerHierarchyNode) GetChildrenOk() (*[]PowerHierarchyNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *PowerHierarchyNode) SetChildren(v []PowerHierarchyNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *PowerHierarchyNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetLabel

`func (o *PowerHierarchyNode) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PowerHierarchyNode) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PowerHierarchyNode) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PowerHierarchyNode) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLevelType

`func (o *PowerHierarchyNode) GetLevelType() PowerHierarchyNodeLevelType`

GetLevelType returns the LevelType field if non-nil, zero value otherwise.

### GetLevelTypeOk

`func (o *PowerHierarchyNode) GetLevelTypeOk() (*PowerHierarchyNodeLevelType, bool)`

GetLevelTypeOk returns a tuple with the LevelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelType

`func (o *PowerHierarchyNode) SetLevelType(v PowerHierarchyNodeLevelType)`

SetLevelType sets LevelType field to given value.

### HasLevelType

`func (o *PowerHierarchyNode) HasLevelType() bool`

HasLevelType returns a boolean if a field has been set.

### GetLevelValue

`func (o *PowerHierarchyNode) GetLevelValue() string`

GetLevelValue returns the LevelValue field if non-nil, zero value otherwise.

### GetLevelValueOk

`func (o *PowerHierarchyNode) GetLevelValueOk() (*string, bool)`

GetLevelValueOk returns a tuple with the LevelValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelValue

`func (o *PowerHierarchyNode) SetLevelValue(v string)`

SetLevelValue sets LevelValue field to given value.

### HasLevelValue

`func (o *PowerHierarchyNode) HasLevelValue() bool`

HasLevelValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


