# AWSVirtualPrivateGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AWSVirtualPrivateGatewayType**](AWSVirtualPrivateGatewayType.md) |  | 
**Required** | **bool** |  | 
**VpcId** | Pointer to **string** |  | [optional] 
**SubnetId** | Pointer to **string** |  | [optional] 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 

## Methods

### NewAWSVirtualPrivateGateway

`func NewAWSVirtualPrivateGateway(type_ AWSVirtualPrivateGatewayType, required bool, deploymentProperties TopologyProperties, ) *AWSVirtualPrivateGateway`

NewAWSVirtualPrivateGateway instantiates a new AWSVirtualPrivateGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSVirtualPrivateGatewayWithDefaults

`func NewAWSVirtualPrivateGatewayWithDefaults() *AWSVirtualPrivateGateway`

NewAWSVirtualPrivateGatewayWithDefaults instantiates a new AWSVirtualPrivateGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AWSVirtualPrivateGateway) GetType() AWSVirtualPrivateGatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSVirtualPrivateGateway) GetTypeOk() (*AWSVirtualPrivateGatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSVirtualPrivateGateway) SetType(v AWSVirtualPrivateGatewayType)`

SetType sets Type field to given value.


### GetRequired

`func (o *AWSVirtualPrivateGateway) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AWSVirtualPrivateGateway) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AWSVirtualPrivateGateway) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetVpcId

`func (o *AWSVirtualPrivateGateway) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AWSVirtualPrivateGateway) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AWSVirtualPrivateGateway) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AWSVirtualPrivateGateway) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetSubnetId

`func (o *AWSVirtualPrivateGateway) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AWSVirtualPrivateGateway) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AWSVirtualPrivateGateway) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *AWSVirtualPrivateGateway) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetDeploymentProperties

`func (o *AWSVirtualPrivateGateway) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *AWSVirtualPrivateGateway) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *AWSVirtualPrivateGateway) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


