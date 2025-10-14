# FabricIPWAN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Scope** | [**NetworkScope**](NetworkScope.md) |  | 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 
**DeploymentProperties** | [**TopologyProperties**](TopologyProperties.md) |  | 

## Methods

### NewFabricIPWAN

`func NewFabricIPWAN(type_ string, name string, scope NetworkScope, deploymentProperties TopologyProperties, ) *FabricIPWAN`

NewFabricIPWAN instantiates a new FabricIPWAN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricIPWANWithDefaults

`func NewFabricIPWANWithDefaults() *FabricIPWAN`

NewFabricIPWANWithDefaults instantiates a new FabricIPWAN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FabricIPWAN) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricIPWAN) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricIPWAN) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *FabricIPWAN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricIPWAN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricIPWAN) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *FabricIPWAN) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FabricIPWAN) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FabricIPWAN) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FabricIPWAN) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetScope

`func (o *FabricIPWAN) GetScope() NetworkScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FabricIPWAN) GetScopeOk() (*NetworkScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FabricIPWAN) SetScope(v NetworkScope)`

SetScope sets Scope field to given value.


### GetLocation

`func (o *FabricIPWAN) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FabricIPWAN) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FabricIPWAN) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FabricIPWAN) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDeploymentProperties

`func (o *FabricIPWAN) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *FabricIPWAN) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *FabricIPWAN) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


