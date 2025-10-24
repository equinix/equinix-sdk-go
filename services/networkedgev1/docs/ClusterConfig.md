# ClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | The cluster name. | [optional] 
**ClusterNodeDetails** | Pointer to [**ClusterNodeDetails**](ClusterNodeDetails.md) |  | [optional] 

## Methods

### NewClusterConfig

`func NewClusterConfig() *ClusterConfig`

NewClusterConfig instantiates a new ClusterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigWithDefaults

`func NewClusterConfigWithDefaults() *ClusterConfig`

NewClusterConfigWithDefaults instantiates a new ClusterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ClusterConfig) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ClusterConfig) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ClusterConfig) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ClusterConfig) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetClusterNodeDetails

`func (o *ClusterConfig) GetClusterNodeDetails() ClusterNodeDetails`

GetClusterNodeDetails returns the ClusterNodeDetails field if non-nil, zero value otherwise.

### GetClusterNodeDetailsOk

`func (o *ClusterConfig) GetClusterNodeDetailsOk() (*ClusterNodeDetails, bool)`

GetClusterNodeDetailsOk returns a tuple with the ClusterNodeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeDetails

`func (o *ClusterConfig) SetClusterNodeDetails(v ClusterNodeDetails)`

SetClusterNodeDetails sets ClusterNodeDetails field to given value.

### HasClusterNodeDetails

`func (o *ClusterConfig) HasClusterNodeDetails() bool`

HasClusterNodeDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


