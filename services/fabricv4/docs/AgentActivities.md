# AgentActivities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Agent Activities URI | [optional] 
**Type** | Pointer to **string** | type | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned agent operation identifier | [optional] [readonly] 
**Agent** | Pointer to [**Agent**](Agent.md) |  | [optional] 
**Status** | Pointer to **string** | Agent activities state COMPLETED, PENDING, PENDING_USER_INPUT, FAILED | [optional] 
**Metadata** | Pointer to [**AgentActivitiesMetadata**](AgentActivitiesMetadata.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewAgentActivities

`func NewAgentActivities() *AgentActivities`

NewAgentActivities instantiates a new AgentActivities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivitiesWithDefaults

`func NewAgentActivitiesWithDefaults() *AgentActivities`

NewAgentActivitiesWithDefaults instantiates a new AgentActivities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AgentActivities) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AgentActivities) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AgentActivities) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AgentActivities) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AgentActivities) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentActivities) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentActivities) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AgentActivities) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *AgentActivities) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AgentActivities) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AgentActivities) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AgentActivities) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAgent

`func (o *AgentActivities) GetAgent() Agent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentActivities) GetAgentOk() (*Agent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentActivities) SetAgent(v Agent)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *AgentActivities) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetStatus

`func (o *AgentActivities) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentActivities) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentActivities) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentActivities) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetadata

`func (o *AgentActivities) GetMetadata() AgentActivitiesMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AgentActivities) GetMetadataOk() (*AgentActivitiesMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AgentActivities) SetMetadata(v AgentActivitiesMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AgentActivities) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetChangeLog

`func (o *AgentActivities) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *AgentActivities) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *AgentActivities) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *AgentActivities) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


