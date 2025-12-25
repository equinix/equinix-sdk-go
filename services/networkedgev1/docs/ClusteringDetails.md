# ClusteringDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusteringEnabled** | Pointer to **bool** | Whether this device management type supports clustering. | [optional] 
**MaxAllowedNodes** | Pointer to **int32** | The number of nodes you can have for a cluster device. | [optional] 

## Methods

### NewClusteringDetails

`func NewClusteringDetails() *ClusteringDetails`

NewClusteringDetails instantiates a new ClusteringDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusteringDetailsWithDefaults

`func NewClusteringDetailsWithDefaults() *ClusteringDetails`

NewClusteringDetailsWithDefaults instantiates a new ClusteringDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusteringEnabled

`func (o *ClusteringDetails) GetClusteringEnabled() bool`

GetClusteringEnabled returns the ClusteringEnabled field if non-nil, zero value otherwise.

### GetClusteringEnabledOk

`func (o *ClusteringDetails) GetClusteringEnabledOk() (*bool, bool)`

GetClusteringEnabledOk returns a tuple with the ClusteringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusteringEnabled

`func (o *ClusteringDetails) SetClusteringEnabled(v bool)`

SetClusteringEnabled sets ClusteringEnabled field to given value.

### HasClusteringEnabled

`func (o *ClusteringDetails) HasClusteringEnabled() bool`

HasClusteringEnabled returns a boolean if a field has been set.

### GetMaxAllowedNodes

`func (o *ClusteringDetails) GetMaxAllowedNodes() int32`

GetMaxAllowedNodes returns the MaxAllowedNodes field if non-nil, zero value otherwise.

### GetMaxAllowedNodesOk

`func (o *ClusteringDetails) GetMaxAllowedNodesOk() (*int32, bool)`

GetMaxAllowedNodesOk returns a tuple with the MaxAllowedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedNodes

`func (o *ClusteringDetails) SetMaxAllowedNodes(v int32)`

SetMaxAllowedNodes sets MaxAllowedNodes field to given value.

### HasMaxAllowedNodes

`func (o *ClusteringDetails) HasMaxAllowedNodes() bool`

HasMaxAllowedNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


