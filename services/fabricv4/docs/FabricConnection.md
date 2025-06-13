# FabricConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FabricConnectionType**](FabricConnectionType.md) |  | 
**Bandwidth** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Redundancy** | [**ConnectionRedundancy**](ConnectionRedundancy.md) |  | 
**ASide** | [**ConnectionSide**](ConnectionSide.md) |  | 
**ZSide** | [**ConnectionSide**](ConnectionSide.md) |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 

## Methods

### NewFabricConnection

`func NewFabricConnection(type_ FabricConnectionType, bandwidth int32, redundancy ConnectionRedundancy, aSide ConnectionSide, zSide ConnectionSide, deploymentProperties TopologyProperties, ) *FabricConnection`

NewFabricConnection instantiates a new FabricConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricConnectionWithDefaults

`func NewFabricConnectionWithDefaults() *FabricConnection`

NewFabricConnectionWithDefaults instantiates a new FabricConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FabricConnection) GetType() FabricConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricConnection) GetTypeOk() (*FabricConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricConnection) SetType(v FabricConnectionType)`

SetType sets Type field to given value.


### GetBandwidth

`func (o *FabricConnection) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *FabricConnection) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *FabricConnection) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetName

`func (o *FabricConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRedundancy

`func (o *FabricConnection) GetRedundancy() ConnectionRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *FabricConnection) GetRedundancyOk() (*ConnectionRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *FabricConnection) SetRedundancy(v ConnectionRedundancy)`

SetRedundancy sets Redundancy field to given value.


### GetASide

`func (o *FabricConnection) GetASide() ConnectionSide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *FabricConnection) GetASideOk() (*ConnectionSide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *FabricConnection) SetASide(v ConnectionSide)`

SetASide sets ASide field to given value.


### GetZSide

`func (o *FabricConnection) GetZSide() ConnectionSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *FabricConnection) GetZSideOk() (*ConnectionSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *FabricConnection) SetZSide(v ConnectionSide)`

SetZSide sets ZSide field to given value.


### GetDeploymentProperties

`func (o *FabricConnection) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *FabricConnection) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *FabricConnection) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


