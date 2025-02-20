# HierarchyNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayLoad** | Pointer to [**HierarchyNodePayLoad**](HierarchyNodePayLoad.md) |  | [optional] 
**Status** | Pointer to [**AssetDetailResponseStatus**](AssetDetailResponseStatus.md) |  | [optional] 

## Methods

### NewHierarchyNode

`func NewHierarchyNode() *HierarchyNode`

NewHierarchyNode instantiates a new HierarchyNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHierarchyNodeWithDefaults

`func NewHierarchyNodeWithDefaults() *HierarchyNode`

NewHierarchyNodeWithDefaults instantiates a new HierarchyNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayLoad

`func (o *HierarchyNode) GetPayLoad() HierarchyNodePayLoad`

GetPayLoad returns the PayLoad field if non-nil, zero value otherwise.

### GetPayLoadOk

`func (o *HierarchyNode) GetPayLoadOk() (*HierarchyNodePayLoad, bool)`

GetPayLoadOk returns a tuple with the PayLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayLoad

`func (o *HierarchyNode) SetPayLoad(v HierarchyNodePayLoad)`

SetPayLoad sets PayLoad field to given value.

### HasPayLoad

`func (o *HierarchyNode) HasPayLoad() bool`

HasPayLoad returns a boolean if a field has been set.

### GetStatus

`func (o *HierarchyNode) GetStatus() AssetDetailResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HierarchyNode) GetStatusOk() (*AssetDetailResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HierarchyNode) SetStatus(v AssetDetailResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HierarchyNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


