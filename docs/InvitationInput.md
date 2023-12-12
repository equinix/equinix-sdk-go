# InvitationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invitee** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**ProjectsIds** | Pointer to **[]string** |  | [optional] 
**Roles** | Pointer to [**[]InvitationRolesInner**](InvitationRolesInner.md) |  | [optional] 

## Methods

### NewInvitationInput

`func NewInvitationInput(invitee string, ) *InvitationInput`

NewInvitationInput instantiates a new InvitationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationInputWithDefaults

`func NewInvitationInputWithDefaults() *InvitationInput`

NewInvitationInputWithDefaults instantiates a new InvitationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitee

`func (o *InvitationInput) GetInvitee() string`

GetInvitee returns the Invitee field if non-nil, zero value otherwise.

### GetInviteeOk

`func (o *InvitationInput) GetInviteeOk() (*string, bool)`

GetInviteeOk returns a tuple with the Invitee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitee

`func (o *InvitationInput) SetInvitee(v string)`

SetInvitee sets Invitee field to given value.


### GetMessage

`func (o *InvitationInput) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvitationInput) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvitationInput) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InvitationInput) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InvitationInput) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InvitationInput) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InvitationInput) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InvitationInput) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProjectsIds

`func (o *InvitationInput) GetProjectsIds() []string`

GetProjectsIds returns the ProjectsIds field if non-nil, zero value otherwise.

### GetProjectsIdsOk

`func (o *InvitationInput) GetProjectsIdsOk() (*[]string, bool)`

GetProjectsIdsOk returns a tuple with the ProjectsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsIds

`func (o *InvitationInput) SetProjectsIds(v []string)`

SetProjectsIds sets ProjectsIds field to given value.

### HasProjectsIds

`func (o *InvitationInput) HasProjectsIds() bool`

HasProjectsIds returns a boolean if a field has been set.

### GetRoles

`func (o *InvitationInput) GetRoles() []InvitationRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *InvitationInput) GetRolesOk() (*[]InvitationRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *InvitationInput) SetRoles(v []InvitationRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *InvitationInput) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


