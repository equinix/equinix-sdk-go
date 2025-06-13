# GCPPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AWSPermissionType**](AWSPermissionType.md) |  | 
**ProjectId** | **string** |  | 
**ProviderId** | **string** |  | 
**PoolId** | **string** |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 

## Methods

### NewGCPPermission

`func NewGCPPermission(type_ AWSPermissionType, projectId string, providerId string, poolId string, deploymentProperties TopologyProperties, ) *GCPPermission`

NewGCPPermission instantiates a new GCPPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPPermissionWithDefaults

`func NewGCPPermissionWithDefaults() *GCPPermission`

NewGCPPermissionWithDefaults instantiates a new GCPPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GCPPermission) GetType() AWSPermissionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCPPermission) GetTypeOk() (*AWSPermissionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCPPermission) SetType(v AWSPermissionType)`

SetType sets Type field to given value.


### GetProjectId

`func (o *GCPPermission) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GCPPermission) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GCPPermission) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProviderId

`func (o *GCPPermission) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GCPPermission) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GCPPermission) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetPoolId

`func (o *GCPPermission) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *GCPPermission) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *GCPPermission) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetDeploymentProperties

`func (o *GCPPermission) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *GCPPermission) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *GCPPermission) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


