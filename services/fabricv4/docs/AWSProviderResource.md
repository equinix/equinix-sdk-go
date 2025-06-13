# AWSProviderResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AWSVirtualPrivateGatewayType**](AWSVirtualPrivateGatewayType.md) |  | 
**RoleArn** | **string** |  | 
**Region** | **string** |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 
**Id** | Pointer to **string** |  | [optional] 
**Required** | **bool** |  | 
**VpcId** | Pointer to **string** |  | [optional] 
**SubnetId** | Pointer to **string** |  | [optional] 

## Methods

### NewAWSProviderResource

`func NewAWSProviderResource(type_ AWSVirtualPrivateGatewayType, roleArn string, region string, deploymentProperties TopologyProperties, required bool, ) *AWSProviderResource`

NewAWSProviderResource instantiates a new AWSProviderResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSProviderResourceWithDefaults

`func NewAWSProviderResourceWithDefaults() *AWSProviderResource`

NewAWSProviderResourceWithDefaults instantiates a new AWSProviderResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AWSProviderResource) GetType() AWSVirtualPrivateGatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSProviderResource) GetTypeOk() (*AWSVirtualPrivateGatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSProviderResource) SetType(v AWSVirtualPrivateGatewayType)`

SetType sets Type field to given value.


### GetRoleArn

`func (o *AWSProviderResource) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AWSProviderResource) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AWSProviderResource) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetRegion

`func (o *AWSProviderResource) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AWSProviderResource) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AWSProviderResource) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetDeploymentProperties

`func (o *AWSProviderResource) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *AWSProviderResource) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *AWSProviderResource) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.


### GetId

`func (o *AWSProviderResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AWSProviderResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AWSProviderResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AWSProviderResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequired

`func (o *AWSProviderResource) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AWSProviderResource) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AWSProviderResource) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetVpcId

`func (o *AWSProviderResource) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AWSProviderResource) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AWSProviderResource) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AWSProviderResource) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetSubnetId

`func (o *AWSProviderResource) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AWSProviderResource) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AWSProviderResource) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *AWSProviderResource) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


