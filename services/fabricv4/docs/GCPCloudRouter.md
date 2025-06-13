# GCPCloudRouter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**GCPCloudRouterType**](GCPCloudRouterType.md) |  | 
**VpcId** | **string** |  | 
**SubnetId** | **string** |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 

## Methods

### NewGCPCloudRouter

`func NewGCPCloudRouter(type_ GCPCloudRouterType, vpcId string, subnetId string, deploymentProperties TopologyProperties, ) *GCPCloudRouter`

NewGCPCloudRouter instantiates a new GCPCloudRouter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPCloudRouterWithDefaults

`func NewGCPCloudRouterWithDefaults() *GCPCloudRouter`

NewGCPCloudRouterWithDefaults instantiates a new GCPCloudRouter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GCPCloudRouter) GetType() GCPCloudRouterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCPCloudRouter) GetTypeOk() (*GCPCloudRouterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCPCloudRouter) SetType(v GCPCloudRouterType)`

SetType sets Type field to given value.


### GetVpcId

`func (o *GCPCloudRouter) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *GCPCloudRouter) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *GCPCloudRouter) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetSubnetId

`func (o *GCPCloudRouter) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *GCPCloudRouter) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *GCPCloudRouter) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetDeploymentProperties

`func (o *GCPCloudRouter) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *GCPCloudRouter) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *GCPCloudRouter) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


