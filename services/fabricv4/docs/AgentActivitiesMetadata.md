# AgentActivitiesMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatMessage** | Pointer to [**ChatMessage**](ChatMessage.md) |  | [optional] 
**ToolCallInformation** | Pointer to [**[]ToolCallInformationInner**](ToolCallInformationInner.md) | List of tools called during the agent operation | [optional] 

## Methods

### NewAgentActivitiesMetadata

`func NewAgentActivitiesMetadata() *AgentActivitiesMetadata`

NewAgentActivitiesMetadata instantiates a new AgentActivitiesMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivitiesMetadataWithDefaults

`func NewAgentActivitiesMetadataWithDefaults() *AgentActivitiesMetadata`

NewAgentActivitiesMetadataWithDefaults instantiates a new AgentActivitiesMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatMessage

`func (o *AgentActivitiesMetadata) GetChatMessage() ChatMessage`

GetChatMessage returns the ChatMessage field if non-nil, zero value otherwise.

### GetChatMessageOk

`func (o *AgentActivitiesMetadata) GetChatMessageOk() (*ChatMessage, bool)`

GetChatMessageOk returns a tuple with the ChatMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatMessage

`func (o *AgentActivitiesMetadata) SetChatMessage(v ChatMessage)`

SetChatMessage sets ChatMessage field to given value.

### HasChatMessage

`func (o *AgentActivitiesMetadata) HasChatMessage() bool`

HasChatMessage returns a boolean if a field has been set.

### GetToolCallInformation

`func (o *AgentActivitiesMetadata) GetToolCallInformation() []ToolCallInformationInner`

GetToolCallInformation returns the ToolCallInformation field if non-nil, zero value otherwise.

### GetToolCallInformationOk

`func (o *AgentActivitiesMetadata) GetToolCallInformationOk() (*[]ToolCallInformationInner, bool)`

GetToolCallInformationOk returns a tuple with the ToolCallInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallInformation

`func (o *AgentActivitiesMetadata) SetToolCallInformation(v []ToolCallInformationInner)`

SetToolCallInformation sets ToolCallInformation field to given value.

### HasToolCallInformation

`func (o *AgentActivitiesMetadata) HasToolCallInformation() bool`

HasToolCallInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


