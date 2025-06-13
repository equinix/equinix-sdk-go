# ActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AWSPermissionType**](AWSPermissionType.md) |  | 
**RoleArn** | **string** |  | 
**Region** | **string** |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 
**ProjectId** | **string** |  | 
**ProviderId** | **string** |  | 
**PoolId** | **string** |  | 

## Methods

### NewActionRequest

`func NewActionRequest(type_ AWSPermissionType, roleArn string, region string, deploymentProperties TopologyProperties, projectId string, providerId string, poolId string, ) *ActionRequest`

NewActionRequest instantiates a new ActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRequestWithDefaults

`func NewActionRequestWithDefaults() *ActionRequest`

NewActionRequestWithDefaults instantiates a new ActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActionRequest) GetType() AWSPermissionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionRequest) GetTypeOk() (*AWSPermissionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionRequest) SetType(v AWSPermissionType)`

SetType sets Type field to given value.


### GetRoleArn

`func (o *ActionRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *ActionRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *ActionRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetRegion

`func (o *ActionRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ActionRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ActionRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetDeploymentProperties

`func (o *ActionRequest) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *ActionRequest) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *ActionRequest) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.


### GetProjectId

`func (o *ActionRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ActionRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ActionRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProviderId

`func (o *ActionRequest) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ActionRequest) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ActionRequest) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetPoolId

`func (o *ActionRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *ActionRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *ActionRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


