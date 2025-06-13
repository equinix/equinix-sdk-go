# FabricConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Connection URI | [optional] [readonly] 
**Type** | [**FabricConnectionType**](FabricConnectionType.md) |  | 
**Uuid** | Pointer to **string** | Equinix-assigned connection identifier | [optional] 
**State** | [**DeploymentState**](DeploymentState.md) |  | 
**Bandwidth** | **int32** |  | 
**Name** | **string** |  | 
**Redundancy** | [**ConnectionRedundancy**](ConnectionRedundancy.md) |  | 
**ASide** | [**ConnectionSide**](ConnectionSide.md) |  | 
**ZSide** | [**ConnectionSide**](ConnectionSide.md) |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 

## Methods

### NewFabricConnectionResponse

`func NewFabricConnectionResponse(type_ FabricConnectionType, state DeploymentState, bandwidth int32, name string, redundancy ConnectionRedundancy, aSide ConnectionSide, zSide ConnectionSide, deploymentProperties TopologyProperties, ) *FabricConnectionResponse`

NewFabricConnectionResponse instantiates a new FabricConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricConnectionResponseWithDefaults

`func NewFabricConnectionResponseWithDefaults() *FabricConnectionResponse`

NewFabricConnectionResponseWithDefaults instantiates a new FabricConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *FabricConnectionResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FabricConnectionResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FabricConnectionResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FabricConnectionResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *FabricConnectionResponse) GetType() FabricConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricConnectionResponse) GetTypeOk() (*FabricConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricConnectionResponse) SetType(v FabricConnectionType)`

SetType sets Type field to given value.


### GetUuid

`func (o *FabricConnectionResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FabricConnectionResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FabricConnectionResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FabricConnectionResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *FabricConnectionResponse) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FabricConnectionResponse) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FabricConnectionResponse) SetState(v DeploymentState)`

SetState sets State field to given value.


### GetBandwidth

`func (o *FabricConnectionResponse) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *FabricConnectionResponse) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *FabricConnectionResponse) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetName

`func (o *FabricConnectionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricConnectionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricConnectionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRedundancy

`func (o *FabricConnectionResponse) GetRedundancy() ConnectionRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *FabricConnectionResponse) GetRedundancyOk() (*ConnectionRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *FabricConnectionResponse) SetRedundancy(v ConnectionRedundancy)`

SetRedundancy sets Redundancy field to given value.


### GetASide

`func (o *FabricConnectionResponse) GetASide() ConnectionSide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *FabricConnectionResponse) GetASideOk() (*ConnectionSide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *FabricConnectionResponse) SetASide(v ConnectionSide)`

SetASide sets ASide field to given value.


### GetZSide

`func (o *FabricConnectionResponse) GetZSide() ConnectionSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *FabricConnectionResponse) GetZSideOk() (*ConnectionSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *FabricConnectionResponse) SetZSide(v ConnectionSide)`

SetZSide sets ZSide field to given value.


### GetDeploymentProperties

`func (o *FabricConnectionResponse) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *FabricConnectionResponse) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *FabricConnectionResponse) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


