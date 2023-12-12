# ProjectList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Projects** | Pointer to [**[]Project**](Project.md) |  | [optional] 

## Methods

### NewProjectList

`func NewProjectList() *ProjectList`

NewProjectList instantiates a new ProjectList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectListWithDefaults

`func NewProjectListWithDefaults() *ProjectList`

NewProjectListWithDefaults instantiates a new ProjectList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ProjectList) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProjectList) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProjectList) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProjectList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetProjects

`func (o *ProjectList) GetProjects() []Project`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ProjectList) GetProjectsOk() (*[]Project, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ProjectList) SetProjects(v []Project)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ProjectList) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


