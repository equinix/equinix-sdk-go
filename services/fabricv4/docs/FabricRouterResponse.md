# FabricRouterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Cloud Routers URI | [optional] [readonly] 
**Type** | [**FabricRouterType**](FabricRouterType.md) |  | 
**Uuid** | Pointer to **string** | Equinix-assigned cloud router identifier | [optional] 
**State** | [**DeploymentState**](DeploymentState.md) |  | 
**Name** | **string** |  | 
**Location** | [**SimplifiedLocationWithoutIBX**](SimplifiedLocationWithoutIBX.md) |  | 
**Package** | [**CloudRouterPostRequestPackage**](CloudRouterPostRequestPackage.md) |  | 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 

## Methods

### NewFabricRouterResponse

`func NewFabricRouterResponse(type_ FabricRouterType, state DeploymentState, name string, location SimplifiedLocationWithoutIBX, package_ CloudRouterPostRequestPackage, deploymentProperties TopologyProperties, ) *FabricRouterResponse`

NewFabricRouterResponse instantiates a new FabricRouterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricRouterResponseWithDefaults

`func NewFabricRouterResponseWithDefaults() *FabricRouterResponse`

NewFabricRouterResponseWithDefaults instantiates a new FabricRouterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *FabricRouterResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FabricRouterResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FabricRouterResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FabricRouterResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *FabricRouterResponse) GetType() FabricRouterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricRouterResponse) GetTypeOk() (*FabricRouterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricRouterResponse) SetType(v FabricRouterType)`

SetType sets Type field to given value.


### GetUuid

`func (o *FabricRouterResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FabricRouterResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FabricRouterResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FabricRouterResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *FabricRouterResponse) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FabricRouterResponse) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FabricRouterResponse) SetState(v DeploymentState)`

SetState sets State field to given value.


### GetName

`func (o *FabricRouterResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricRouterResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricRouterResponse) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *FabricRouterResponse) GetLocation() SimplifiedLocationWithoutIBX`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FabricRouterResponse) GetLocationOk() (*SimplifiedLocationWithoutIBX, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FabricRouterResponse) SetLocation(v SimplifiedLocationWithoutIBX)`

SetLocation sets Location field to given value.


### GetPackage

`func (o *FabricRouterResponse) GetPackage() CloudRouterPostRequestPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *FabricRouterResponse) GetPackageOk() (*CloudRouterPostRequestPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *FabricRouterResponse) SetPackage(v CloudRouterPostRequestPackage)`

SetPackage sets Package field to given value.


### GetDeploymentProperties

`func (o *FabricRouterResponse) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *FabricRouterResponse) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *FabricRouterResponse) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


