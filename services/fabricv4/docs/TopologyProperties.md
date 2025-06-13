# TopologyProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElementId** | **string** |  | 
**DependsOn** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTopologyProperties

`func NewTopologyProperties(elementId string, ) *TopologyProperties`

NewTopologyProperties instantiates a new TopologyProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyPropertiesWithDefaults

`func NewTopologyPropertiesWithDefaults() *TopologyProperties`

NewTopologyPropertiesWithDefaults instantiates a new TopologyProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElementId

`func (o *TopologyProperties) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *TopologyProperties) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *TopologyProperties) SetElementId(v string)`

SetElementId sets ElementId field to given value.


### GetDependsOn

`func (o *TopologyProperties) GetDependsOn() []string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *TopologyProperties) GetDependsOnOk() (*[]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *TopologyProperties) SetDependsOn(v []string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *TopologyProperties) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


