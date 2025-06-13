# AWSVirtualPrivateGatewayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Virtual private gateway id. | [optional] 
**Type** | **string** |  | 
**State** | [**DeploymentState**](DeploymentState.md) |  | 
**Required** | Pointer to **bool** |  | [optional] 
**VpcId** | **string** |  | 
**SubnetId** | **string** |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 

## Methods

### NewAWSVirtualPrivateGatewayResponse

`func NewAWSVirtualPrivateGatewayResponse(type_ string, state DeploymentState, vpcId string, subnetId string, deploymentProperties TopologyProperties, ) *AWSVirtualPrivateGatewayResponse`

NewAWSVirtualPrivateGatewayResponse instantiates a new AWSVirtualPrivateGatewayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSVirtualPrivateGatewayResponseWithDefaults

`func NewAWSVirtualPrivateGatewayResponseWithDefaults() *AWSVirtualPrivateGatewayResponse`

NewAWSVirtualPrivateGatewayResponseWithDefaults instantiates a new AWSVirtualPrivateGatewayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AWSVirtualPrivateGatewayResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AWSVirtualPrivateGatewayResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AWSVirtualPrivateGatewayResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AWSVirtualPrivateGatewayResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AWSVirtualPrivateGatewayResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSVirtualPrivateGatewayResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSVirtualPrivateGatewayResponse) SetType(v string)`

SetType sets Type field to given value.


### GetState

`func (o *AWSVirtualPrivateGatewayResponse) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AWSVirtualPrivateGatewayResponse) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AWSVirtualPrivateGatewayResponse) SetState(v DeploymentState)`

SetState sets State field to given value.


### GetRequired

`func (o *AWSVirtualPrivateGatewayResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AWSVirtualPrivateGatewayResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AWSVirtualPrivateGatewayResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *AWSVirtualPrivateGatewayResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetVpcId

`func (o *AWSVirtualPrivateGatewayResponse) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AWSVirtualPrivateGatewayResponse) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AWSVirtualPrivateGatewayResponse) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetSubnetId

`func (o *AWSVirtualPrivateGatewayResponse) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AWSVirtualPrivateGatewayResponse) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AWSVirtualPrivateGatewayResponse) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetDeploymentProperties

`func (o *AWSVirtualPrivateGatewayResponse) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *AWSVirtualPrivateGatewayResponse) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *AWSVirtualPrivateGatewayResponse) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


