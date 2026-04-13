# AgentPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** | Customer-provided agent name | 
**Description** | Pointer to **string** | Customer-provided agent description | [optional] 
**Enabled** | Pointer to **bool** | Customer-provided agent enabled status | [optional] 
**Project** | [**Project**](Project.md) |  | 
**AgentTemplate** | [**AgentTemplate**](AgentTemplate.md) |  | 
**Configuration** | Pointer to [**AgentConfiguration**](AgentConfiguration.md) |  | [optional] 

## Methods

### NewAgentPostRequest

`func NewAgentPostRequest(type_ string, name string, project Project, agentTemplate AgentTemplate, ) *AgentPostRequest`

NewAgentPostRequest instantiates a new AgentPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPostRequestWithDefaults

`func NewAgentPostRequestWithDefaults() *AgentPostRequest`

NewAgentPostRequestWithDefaults instantiates a new AgentPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AgentPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentPostRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AgentPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AgentPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AgentPostRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentPostRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentPostRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgentPostRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProject

`func (o *AgentPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AgentPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AgentPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.


### GetAgentTemplate

`func (o *AgentPostRequest) GetAgentTemplate() AgentTemplate`

GetAgentTemplate returns the AgentTemplate field if non-nil, zero value otherwise.

### GetAgentTemplateOk

`func (o *AgentPostRequest) GetAgentTemplateOk() (*AgentTemplate, bool)`

GetAgentTemplateOk returns a tuple with the AgentTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTemplate

`func (o *AgentPostRequest) SetAgentTemplate(v AgentTemplate)`

SetAgentTemplate sets AgentTemplate field to given value.


### GetConfiguration

`func (o *AgentPostRequest) GetConfiguration() AgentConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AgentPostRequest) GetConfigurationOk() (*AgentConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AgentPostRequest) SetConfiguration(v AgentConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *AgentPostRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


