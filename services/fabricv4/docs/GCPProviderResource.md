# GCPProviderResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**GCPCloudRouterType**](GCPCloudRouterType.md) |  | 
**ProjectId** | **string** |  | 
**ProviderId** | **string** |  | 
**PoolId** | **string** |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 
**VpcId** | **string** |  | 
**SubnetId** | **string** |  | 

## Methods

### NewGCPProviderResource

`func NewGCPProviderResource(type_ GCPCloudRouterType, projectId string, providerId string, poolId string, deploymentProperties TopologyProperties, vpcId string, subnetId string, ) *GCPProviderResource`

NewGCPProviderResource instantiates a new GCPProviderResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPProviderResourceWithDefaults

`func NewGCPProviderResourceWithDefaults() *GCPProviderResource`

NewGCPProviderResourceWithDefaults instantiates a new GCPProviderResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GCPProviderResource) GetType() GCPCloudRouterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCPProviderResource) GetTypeOk() (*GCPCloudRouterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCPProviderResource) SetType(v GCPCloudRouterType)`

SetType sets Type field to given value.


### GetProjectId

`func (o *GCPProviderResource) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GCPProviderResource) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GCPProviderResource) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProviderId

`func (o *GCPProviderResource) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GCPProviderResource) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GCPProviderResource) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetPoolId

`func (o *GCPProviderResource) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *GCPProviderResource) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *GCPProviderResource) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetDeploymentProperties

`func (o *GCPProviderResource) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *GCPProviderResource) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *GCPProviderResource) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.


### GetVpcId

`func (o *GCPProviderResource) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *GCPProviderResource) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *GCPProviderResource) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetSubnetId

`func (o *GCPProviderResource) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *GCPProviderResource) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *GCPProviderResource) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


