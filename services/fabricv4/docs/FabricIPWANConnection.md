# FabricIPWANConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Bandwidth** | **int32** |  | 
**Name** | **string** |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 

## Methods

### NewFabricIPWANConnection

`func NewFabricIPWANConnection(type_ string, bandwidth int32, name string, deploymentProperties TopologyProperties, ) *FabricIPWANConnection`

NewFabricIPWANConnection instantiates a new FabricIPWANConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricIPWANConnectionWithDefaults

`func NewFabricIPWANConnectionWithDefaults() *FabricIPWANConnection`

NewFabricIPWANConnectionWithDefaults instantiates a new FabricIPWANConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FabricIPWANConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricIPWANConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricIPWANConnection) SetType(v string)`

SetType sets Type field to given value.


### GetBandwidth

`func (o *FabricIPWANConnection) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *FabricIPWANConnection) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *FabricIPWANConnection) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetName

`func (o *FabricIPWANConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricIPWANConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricIPWANConnection) SetName(v string)`

SetName sets Name field to given value.


### GetDeploymentProperties

`func (o *FabricIPWANConnection) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *FabricIPWANConnection) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *FabricIPWANConnection) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


