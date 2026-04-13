# AgentGetActivities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]AgentActivities**](AgentActivities.md) | Data returned from the API call. | [optional] 

## Methods

### NewAgentGetActivities

`func NewAgentGetActivities() *AgentGetActivities`

NewAgentGetActivities instantiates a new AgentGetActivities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGetActivitiesWithDefaults

`func NewAgentGetActivitiesWithDefaults() *AgentGetActivities`

NewAgentGetActivitiesWithDefaults instantiates a new AgentGetActivities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *AgentGetActivities) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AgentGetActivities) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AgentGetActivities) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AgentGetActivities) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *AgentGetActivities) GetData() []AgentActivities`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AgentGetActivities) GetDataOk() (*[]AgentActivities, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AgentGetActivities) SetData(v []AgentActivities)`

SetData sets Data field to given value.

### HasData

`func (o *AgentGetActivities) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


