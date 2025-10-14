# FabricIPWANConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | IPWAN Connection URI | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned ipwan connection identifier | [optional] 
**State** | Pointer to [**DeploymentState**](DeploymentState.md) |  | [optional] 
**Bandwidth** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DeploymentProperties** | Pointer to [**TopologyProperties**](TopologyProperties.md) |  | [optional] 

## Methods

### NewFabricIPWANConnectionResponse

`func NewFabricIPWANConnectionResponse() *FabricIPWANConnectionResponse`

NewFabricIPWANConnectionResponse instantiates a new FabricIPWANConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricIPWANConnectionResponseWithDefaults

`func NewFabricIPWANConnectionResponseWithDefaults() *FabricIPWANConnectionResponse`

NewFabricIPWANConnectionResponseWithDefaults instantiates a new FabricIPWANConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *FabricIPWANConnectionResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FabricIPWANConnectionResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FabricIPWANConnectionResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FabricIPWANConnectionResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *FabricIPWANConnectionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricIPWANConnectionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricIPWANConnectionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FabricIPWANConnectionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *FabricIPWANConnectionResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FabricIPWANConnectionResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FabricIPWANConnectionResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FabricIPWANConnectionResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *FabricIPWANConnectionResponse) GetState() DeploymentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FabricIPWANConnectionResponse) GetStateOk() (*DeploymentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FabricIPWANConnectionResponse) SetState(v DeploymentState)`

SetState sets State field to given value.

### HasState

`func (o *FabricIPWANConnectionResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetBandwidth

`func (o *FabricIPWANConnectionResponse) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *FabricIPWANConnectionResponse) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *FabricIPWANConnectionResponse) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *FabricIPWANConnectionResponse) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetName

`func (o *FabricIPWANConnectionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricIPWANConnectionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricIPWANConnectionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricIPWANConnectionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeploymentProperties

`func (o *FabricIPWANConnectionResponse) GetDeploymentProperties() TopologyProperties`

GetDeploymentProperties returns the DeploymentProperties field if non-nil, zero value otherwise.

### GetDeploymentPropertiesOk

`func (o *FabricIPWANConnectionResponse) GetDeploymentPropertiesOk() (*TopologyProperties, bool)`

GetDeploymentPropertiesOk returns a tuple with the DeploymentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentProperties

`func (o *FabricIPWANConnectionResponse) SetDeploymentProperties(v TopologyProperties)`

SetDeploymentProperties sets DeploymentProperties field to given value.

### HasDeploymentProperties

`func (o *FabricIPWANConnectionResponse) HasDeploymentProperties() bool`

HasDeploymentProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


