# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to [**[]InvitationRolesInner**](InvitationRolesInner.md) |  | [optional] 
**ProjectsCount** | Pointer to **int32** |  | [optional] 
**BoundRoles** | Pointer to **[]string** |  | [optional] 
**User** | Pointer to [**Href**](Href.md) |  | [optional] 
**Organization** | Pointer to [**Href**](Href.md) |  | [optional] 
**Projects** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 

## Methods

### NewMember

`func NewMember() *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Member) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Member) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Member) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Member) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoles

`func (o *Member) GetRoles() []InvitationRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Member) GetRolesOk() (*[]InvitationRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Member) SetRoles(v []InvitationRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Member) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetProjectsCount

`func (o *Member) GetProjectsCount() int32`

GetProjectsCount returns the ProjectsCount field if non-nil, zero value otherwise.

### GetProjectsCountOk

`func (o *Member) GetProjectsCountOk() (*int32, bool)`

GetProjectsCountOk returns a tuple with the ProjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCount

`func (o *Member) SetProjectsCount(v int32)`

SetProjectsCount sets ProjectsCount field to given value.

### HasProjectsCount

`func (o *Member) HasProjectsCount() bool`

HasProjectsCount returns a boolean if a field has been set.

### GetBoundRoles

`func (o *Member) GetBoundRoles() []string`

GetBoundRoles returns the BoundRoles field if non-nil, zero value otherwise.

### GetBoundRolesOk

`func (o *Member) GetBoundRolesOk() (*[]string, bool)`

GetBoundRolesOk returns a tuple with the BoundRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRoles

`func (o *Member) SetBoundRoles(v []string)`

SetBoundRoles sets BoundRoles field to given value.

### HasBoundRoles

`func (o *Member) HasBoundRoles() bool`

HasBoundRoles returns a boolean if a field has been set.

### GetUser

`func (o *Member) GetUser() Href`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Member) GetUserOk() (*Href, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Member) SetUser(v Href)`

SetUser sets User field to given value.

### HasUser

`func (o *Member) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetOrganization

`func (o *Member) GetOrganization() Href`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Member) GetOrganizationOk() (*Href, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Member) SetOrganization(v Href)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Member) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProjects

`func (o *Member) GetProjects() []Href`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *Member) GetProjectsOk() (*[]Href, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *Member) SetProjects(v []Href)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *Member) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetHref

`func (o *Member) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Member) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Member) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Member) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


