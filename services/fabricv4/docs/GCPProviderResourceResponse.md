# GCPProviderResourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**GCPCloudRouterType**](GCPCloudRouterType.md) |  | 
**ProjectId** | **string** |  | 
**ProviderId** | **string** |  | 
**PoolId** | **string** |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 
**Id** | Pointer to **string** | Cloud Router id. | [optional] 
**State** | Pointer to [**DeploymentState**](DeploymentState.md) |  | [optional] 
**VpcId** | Pointer to **string** |  | [optional] 
**SubnetId** | Pointer to **string** |  | [optional] 

## Methods

### NewGCPProviderResourceResponse

`func NewGCPProviderResourceResponse(type_ GCPCloudRouterType, projectId string, providerId string, poolId string, deploymentProperties TopologyProperties, ) *GCPProviderResourceResponse`

NewGCPProviderResourceResponse instantiates a new GCPProviderResourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPProviderResourceResponseWithDefaults

`func NewGCPProviderResourceResponseWithDefaults() *GCPProviderResourceResponse`

NewGCPProviderResourceResponseWithDefaults instantiates a new GCPProviderResourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GCPProviderResourceResponse) GetType() GCPCloudRouterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCPProviderResourceResponse) GetTypeOk() (*GCPCloudRouterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCPProviderResourceResponse) SetType(v GCPCloudRouterType)`

SetType sets Type field to given value.


### GetProjectId

`func (o *GCPProviderResourceResponse) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GCPProviderResourceResponse) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GCPProviderResourceResponse) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProviderId

`func (o *GCPProviderResourceResponse) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GCPProviderResourceResponse) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GCPProviderResourceResponse) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetPoolId

`func (o *GCPProviderResourceResponse) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *GCPProviderResourceResponse) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *GCPProviderResourceResponse) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetDeploymentProperties

`func (o *GCPProviderResourceResponse) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *GCPProviderResourceResponse) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *GCPProviderResourceResponse) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.


### GetId

`func (o *GCPProviderResourceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GCPProviderResourceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GCPProviderResourceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GCPProviderResourceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *GCPProviderResourceResponse) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GCPProviderResourceResponse) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GCPProviderResourceResponse) SetState(v DeploymentState)`

SetState sets State field to given value.

### HasState

`func (o *GCPProviderResourceResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVpcId

`func (o *GCPProviderResourceResponse) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *GCPProviderResourceResponse) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *GCPProviderResourceResponse) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *GCPProviderResourceResponse) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetSubnetId

`func (o *GCPProviderResourceResponse) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *GCPProviderResourceResponse) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *GCPProviderResourceResponse) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *GCPProviderResourceResponse) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


