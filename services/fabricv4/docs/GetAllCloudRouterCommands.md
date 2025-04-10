# GetAllCloudRouterCommands

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]CloudRouterCommand**](CloudRouterCommand.md) |  | [optional] 

## Methods

### NewGetAllCloudRouterCommands

`func NewGetAllCloudRouterCommands() *GetAllCloudRouterCommands`

NewGetAllCloudRouterCommands instantiates a new GetAllCloudRouterCommands object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllCloudRouterCommandsWithDefaults

`func NewGetAllCloudRouterCommandsWithDefaults() *GetAllCloudRouterCommands`

NewGetAllCloudRouterCommandsWithDefaults instantiates a new GetAllCloudRouterCommands object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *GetAllCloudRouterCommands) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetAllCloudRouterCommands) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetAllCloudRouterCommands) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetAllCloudRouterCommands) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *GetAllCloudRouterCommands) GetData() []CloudRouterCommand`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllCloudRouterCommands) GetDataOk() (*[]CloudRouterCommand, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllCloudRouterCommands) SetData(v []CloudRouterCommand)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllCloudRouterCommands) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


