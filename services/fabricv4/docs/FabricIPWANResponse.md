# FabricIPWANResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | IPWAN URI | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned IPWAN identifier | [optional] 
**State** | Pointer to [**DeploymentState**](DeploymentState.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 
**Scope** | Pointer to [**NetworkScope**](NetworkScope.md) |  | [optional] 
**DeploymentProperties** | Pointer to [**TopologyProperties**](TopologyProperties.md) |  | [optional] 

## Methods

### NewFabricIPWANResponse

`func NewFabricIPWANResponse() *FabricIPWANResponse`

NewFabricIPWANResponse instantiates a new FabricIPWANResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricIPWANResponseWithDefaults

`func NewFabricIPWANResponseWithDefaults() *FabricIPWANResponse`

NewFabricIPWANResponseWithDefaults instantiates a new FabricIPWANResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *FabricIPWANResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FabricIPWANResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FabricIPWANResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FabricIPWANResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *FabricIPWANResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricIPWANResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricIPWANResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FabricIPWANResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *FabricIPWANResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FabricIPWANResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FabricIPWANResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FabricIPWANResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *FabricIPWANResponse) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FabricIPWANResponse) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FabricIPWANResponse) SetState(v DeploymentState)`

SetState sets State field to given value.

### HasState

`func (o *FabricIPWANResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *FabricIPWANResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricIPWANResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricIPWANResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricIPWANResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *FabricIPWANResponse) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FabricIPWANResponse) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FabricIPWANResponse) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FabricIPWANResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetScope

`func (o *FabricIPWANResponse) GetScope() NetworkScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FabricIPWANResponse) GetScopeOk() (*NetworkScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FabricIPWANResponse) SetScope(v NetworkScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *FabricIPWANResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDeploymentProperties

`func (o *FabricIPWANResponse) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *FabricIPWANResponse) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *FabricIPWANResponse) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.

### HasDeploymentProperties

`func (o *FabricIPWANResponse) HasDeploymentProperties() bool`

HasDeploymentProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


