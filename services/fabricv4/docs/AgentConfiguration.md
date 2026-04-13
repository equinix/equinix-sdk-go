# AgentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | Pointer to **string** | Agent configuration prompt to be used for agent specification | [optional] 

## Methods

### NewAgentConfiguration

`func NewAgentConfiguration() *AgentConfiguration`

NewAgentConfiguration instantiates a new AgentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentConfigurationWithDefaults

`func NewAgentConfigurationWithDefaults() *AgentConfiguration`

NewAgentConfigurationWithDefaults instantiates a new AgentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompt

`func (o *AgentConfiguration) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AgentConfiguration) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AgentConfiguration) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AgentConfiguration) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


