# AWSDirectConnectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SearchDirectConnectType**](SearchDirectConnectType.md) |  | 
**Id** | Pointer to **string** |  | [optional] 
**State** | [**DeploymentState**](DeploymentState.md) |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 

## Methods

### NewAWSDirectConnectResponse

`func NewAWSDirectConnectResponse(type_ SearchDirectConnectType, state DeploymentState, deploymentProperties TopologyProperties, ) *AWSDirectConnectResponse`

NewAWSDirectConnectResponse instantiates a new AWSDirectConnectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSDirectConnectResponseWithDefaults

`func NewAWSDirectConnectResponseWithDefaults() *AWSDirectConnectResponse`

NewAWSDirectConnectResponseWithDefaults instantiates a new AWSDirectConnectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AWSDirectConnectResponse) GetType() SearchDirectConnectType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSDirectConnectResponse) GetTypeOk() (*SearchDirectConnectType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSDirectConnectResponse) SetType(v SearchDirectConnectType)`

SetType sets Type field to given value.


### GetId

`func (o *AWSDirectConnectResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AWSDirectConnectResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AWSDirectConnectResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AWSDirectConnectResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *AWSDirectConnectResponse) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AWSDirectConnectResponse) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AWSDirectConnectResponse) SetState(v DeploymentState)`

SetState sets State field to given value.


### GetDeploymentProperties

`func (o *AWSDirectConnectResponse) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *AWSDirectConnectResponse) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *AWSDirectConnectResponse) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


