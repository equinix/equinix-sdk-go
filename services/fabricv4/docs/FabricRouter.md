# FabricRouter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FabricRouterType**](FabricRouterType.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Location** | Pointer to [**SimplifiedLocationWithoutIBX**](SimplifiedLocationWithoutIBX.md) |  | [optional] 
**Package** | Pointer to [**CloudRouterPostRequestPackage**](CloudRouterPostRequestPackage.md) |  | [optional] 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 

## Methods

### NewFabricRouter

`func NewFabricRouter(type_ FabricRouterType, deploymentProperties TopologyProperties, ) *FabricRouter`

NewFabricRouter instantiates a new FabricRouter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricRouterWithDefaults

`func NewFabricRouterWithDefaults() *FabricRouter`

NewFabricRouterWithDefaults instantiates a new FabricRouter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FabricRouter) GetType() FabricRouterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricRouter) GetTypeOk() (*FabricRouterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricRouter) SetType(v FabricRouterType)`

SetType sets Type field to given value.


### GetName

`func (o *FabricRouter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricRouter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricRouter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricRouter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *FabricRouter) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FabricRouter) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FabricRouter) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FabricRouter) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetLocation

`func (o *FabricRouter) GetLocation() SimplifiedLocationWithoutIBX`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FabricRouter) GetLocationOk() (*SimplifiedLocationWithoutIBX, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FabricRouter) SetLocation(v SimplifiedLocationWithoutIBX)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FabricRouter) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPackage

`func (o *FabricRouter) GetPackage() CloudRouterPostRequestPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *FabricRouter) GetPackageOk() (*CloudRouterPostRequestPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *FabricRouter) SetPackage(v CloudRouterPostRequestPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *FabricRouter) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetDeploymentProperties

`func (o *FabricRouter) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *FabricRouter) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *FabricRouter) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


