# AgentTemplates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Agent Template URI | [optional] 
**Type** | Pointer to **string** | type | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Name** | Pointer to **string** | Equinix-provided agent template name | [optional] 
**Description** | Pointer to **string** | Equinix-provided agent template description | [optional] 
**State** | Pointer to [**AgentTemplatesState**](AgentTemplatesState.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Equinix-provided agent template enabled status | [optional] 
**AgentDefinition** | Pointer to [**AgentDefinition**](AgentDefinition.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewAgentTemplates

`func NewAgentTemplates() *AgentTemplates`

NewAgentTemplates instantiates a new AgentTemplates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTemplatesWithDefaults

`func NewAgentTemplatesWithDefaults() *AgentTemplates`

NewAgentTemplatesWithDefaults instantiates a new AgentTemplates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AgentTemplates) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AgentTemplates) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AgentTemplates) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AgentTemplates) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AgentTemplates) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentTemplates) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentTemplates) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AgentTemplates) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *AgentTemplates) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AgentTemplates) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AgentTemplates) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AgentTemplates) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *AgentTemplates) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentTemplates) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentTemplates) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentTemplates) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AgentTemplates) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentTemplates) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentTemplates) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentTemplates) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *AgentTemplates) GetState() AgentTemplatesState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AgentTemplates) GetStateOk() (*AgentTemplatesState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AgentTemplates) SetState(v AgentTemplatesState)`

SetState sets State field to given value.

### HasState

`func (o *AgentTemplates) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEnabled

`func (o *AgentTemplates) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentTemplates) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentTemplates) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgentTemplates) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAgentDefinition

`func (o *AgentTemplates) GetAgentDefinition() AgentDefinition`

GetAgentDefinition returns the AgentDefinition field if non-nil, zero value otherwise.

### GetAgentDefinitionOk

`func (o *AgentTemplates) GetAgentDefinitionOk() (*AgentDefinition, bool)`

GetAgentDefinitionOk returns a tuple with the AgentDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentDefinition

`func (o *AgentTemplates) SetAgentDefinition(v AgentDefinition)`

SetAgentDefinition sets AgentDefinition field to given value.

### HasAgentDefinition

`func (o *AgentTemplates) HasAgentDefinition() bool`

HasAgentDefinition returns a boolean if a field has been set.

### GetChangeLog

`func (o *AgentTemplates) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *AgentTemplates) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *AgentTemplates) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *AgentTemplates) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


