# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Topology** | Pointer to [**DeploymentTopology**](DeploymentTopology.md) |  | [optional] 
**Providers** | [**[]OrchestratorProviders**](OrchestratorProviders.md) |  | 
**Project** | [**Project**](Project.md) |  | 
**Account** | [**SimplifiedAccount**](SimplifiedAccount.md) |  | 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on deployement status changes | [optional] 

## Methods

### NewDeployment

`func NewDeployment(type_ string, name string, providers []OrchestratorProviders, project Project, account SimplifiedAccount, ) *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Deployment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Deployment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Deployment) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *Deployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Deployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Deployment) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Deployment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Deployment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Deployment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Deployment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTopology

`func (o *Deployment) GetTopology() DeploymentTopology`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *Deployment) GetTopologyOk() (*DeploymentTopology, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *Deployment) SetTopology(v DeploymentTopology)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *Deployment) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetProviders

`func (o *Deployment) GetProviders() []OrchestratorProviders`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *Deployment) GetProvidersOk() (*[]OrchestratorProviders, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *Deployment) SetProviders(v []OrchestratorProviders)`

SetProviders sets Providers field to given value.


### GetProject

`func (o *Deployment) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Deployment) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Deployment) SetProject(v Project)`

SetProject sets Project field to given value.


### GetAccount

`func (o *Deployment) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Deployment) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Deployment) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.


### GetOrder

`func (o *Deployment) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Deployment) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Deployment) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Deployment) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetNotifications

`func (o *Deployment) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *Deployment) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *Deployment) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *Deployment) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


