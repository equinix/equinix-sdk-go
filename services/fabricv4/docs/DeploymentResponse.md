# DeploymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Topology** | Pointer to [**DeploymentTopology**](DeploymentTopology.md) |  | [optional] 
**Providers** | Pointer to [**[]ProviderResponse**](ProviderResponse.md) |  | [optional] 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on deployement status changes | [optional] 

## Methods

### NewDeploymentResponse

`func NewDeploymentResponse() *DeploymentResponse`

NewDeploymentResponse instantiates a new DeploymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentResponseWithDefaults

`func NewDeploymentResponseWithDefaults() *DeploymentResponse`

NewDeploymentResponseWithDefaults instantiates a new DeploymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *DeploymentResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DeploymentResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DeploymentResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DeploymentResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *DeploymentResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeploymentResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeploymentResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeploymentResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *DeploymentResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DeploymentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DeploymentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeploymentResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *DeploymentResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DeploymentResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTopology

`func (o *DeploymentResponse) GetTopology() DeploymentTopology`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *DeploymentResponse) GetTopologyOk() (*DeploymentTopology, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *DeploymentResponse) SetTopology(v DeploymentTopology)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *DeploymentResponse) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetProviders

`func (o *DeploymentResponse) GetProviders() []ProviderResponse`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *DeploymentResponse) GetProvidersOk() (*[]ProviderResponse, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *DeploymentResponse) SetProviders(v []ProviderResponse)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *DeploymentResponse) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetNotifications

`func (o *DeploymentResponse) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *DeploymentResponse) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *DeploymentResponse) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *DeploymentResponse) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


