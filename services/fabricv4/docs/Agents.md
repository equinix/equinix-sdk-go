# Agents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Agent URI | [optional] [readonly] 
**Type** | Pointer to **string** | type | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Name** | Pointer to **string** | Customer-provided agent name | [optional] 
**Description** | Pointer to **string** | Customer-provided agent description | [optional] 
**State** | Pointer to [**AgentTemplatesState**](AgentTemplatesState.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Customer-provided agent enabled status | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**AgentTemplate** | Pointer to [**AgentTemplate**](AgentTemplate.md) |  | [optional] 
**Configuration** | Pointer to [**Configuration**](Configuration.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewAgents

`func NewAgents() *Agents`

NewAgents instantiates a new Agents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsWithDefaults

`func NewAgentsWithDefaults() *Agents`

NewAgentsWithDefaults instantiates a new Agents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Agents) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Agents) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Agents) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Agents) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *Agents) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Agents) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Agents) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Agents) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *Agents) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Agents) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Agents) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Agents) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *Agents) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agents) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Agents) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Agents) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Agents) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Agents) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Agents) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Agents) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *Agents) GetState() AgentTemplatesState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Agents) GetStateOk() (*AgentTemplatesState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Agents) SetState(v AgentTemplatesState)`

SetState sets State field to given value.

### HasState

`func (o *Agents) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEnabled

`func (o *Agents) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Agents) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Agents) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Agents) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProject

`func (o *Agents) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Agents) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Agents) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Agents) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAgentTemplate

`func (o *Agents) GetAgentTemplate() AgentTemplate`

GetAgentTemplate returns the AgentTemplate field if non-nil, zero value otherwise.

### GetAgentTemplateOk

`func (o *Agents) GetAgentTemplateOk() (*AgentTemplate, bool)`

GetAgentTemplateOk returns a tuple with the AgentTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTemplate

`func (o *Agents) SetAgentTemplate(v AgentTemplate)`

SetAgentTemplate sets AgentTemplate field to given value.

### HasAgentTemplate

`func (o *Agents) HasAgentTemplate() bool`

HasAgentTemplate returns a boolean if a field has been set.

### GetConfiguration

`func (o *Agents) GetConfiguration() Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Agents) GetConfigurationOk() (*Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Agents) SetConfiguration(v Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Agents) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetChangeLog

`func (o *Agents) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *Agents) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *Agents) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *Agents) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


