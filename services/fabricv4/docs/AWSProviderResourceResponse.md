# AWSProviderResourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**RoleArn** | **string** |  | 
**Region** | **string** |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 
**Id** | Pointer to **string** | Virtual private gateway id. | [optional] 
**State** | [**DeploymentState**](DeploymentState.md) |  | 
**Required** | Pointer to **bool** |  | [optional] 
**VpcId** | **string** |  | 
**SubnetId** | **string** |  | 

## Methods

### NewAWSProviderResourceResponse

`func NewAWSProviderResourceResponse(type_ string, roleArn string, region string, deploymentProperties TopologyProperties, state DeploymentState, vpcId string, subnetId string, ) *AWSProviderResourceResponse`

NewAWSProviderResourceResponse instantiates a new AWSProviderResourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSProviderResourceResponseWithDefaults

`func NewAWSProviderResourceResponseWithDefaults() *AWSProviderResourceResponse`

NewAWSProviderResourceResponseWithDefaults instantiates a new AWSProviderResourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AWSProviderResourceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSProviderResourceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSProviderResourceResponse) SetType(v string)`

SetType sets Type field to given value.


### GetRoleArn

`func (o *AWSProviderResourceResponse) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AWSProviderResourceResponse) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AWSProviderResourceResponse) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetRegion

`func (o *AWSProviderResourceResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AWSProviderResourceResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AWSProviderResourceResponse) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetDeploymentProperties

`func (o *AWSProviderResourceResponse) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *AWSProviderResourceResponse) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *AWSProviderResourceResponse) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.


### GetId

`func (o *AWSProviderResourceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AWSProviderResourceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AWSProviderResourceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AWSProviderResourceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *AWSProviderResourceResponse) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AWSProviderResourceResponse) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AWSProviderResourceResponse) SetState(v DeploymentState)`

SetState sets State field to given value.


### GetRequired

`func (o *AWSProviderResourceResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AWSProviderResourceResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AWSProviderResourceResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *AWSProviderResourceResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetVpcId

`func (o *AWSProviderResourceResponse) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AWSProviderResourceResponse) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AWSProviderResourceResponse) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetSubnetId

`func (o *AWSProviderResourceResponse) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AWSProviderResourceResponse) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AWSProviderResourceResponse) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


