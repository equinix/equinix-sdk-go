# AWSPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AWSPermissionType**](AWSPermissionType.md) |  | 
**RoleArn** | **string** |  | 
**Region** | **string** |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 

## Methods

### NewAWSPermission

`func NewAWSPermission(type_ AWSPermissionType, roleArn string, region string, deploymentProperties TopologyProperties, ) *AWSPermission`

NewAWSPermission instantiates a new AWSPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSPermissionWithDefaults

`func NewAWSPermissionWithDefaults() *AWSPermission`

NewAWSPermissionWithDefaults instantiates a new AWSPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AWSPermission) GetType() AWSPermissionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSPermission) GetTypeOk() (*AWSPermissionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSPermission) SetType(v AWSPermissionType)`

SetType sets Type field to given value.


### GetRoleArn

`func (o *AWSPermission) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AWSPermission) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AWSPermission) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetRegion

`func (o *AWSPermission) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AWSPermission) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AWSPermission) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetDeploymentProperties

`func (o *AWSPermission) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *AWSPermission) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *AWSPermission) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


