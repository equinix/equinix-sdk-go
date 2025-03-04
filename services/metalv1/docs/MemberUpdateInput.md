# MemberUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**[]InvitationRolesInner**](InvitationRolesInner.md) | Primary role for the user within the organization | [optional] 
**BoundRoles** | Pointer to **[]string** | Additional roles that can be bound to the user to grant extra permissions. | [optional] 
**ProjectIds** | Pointer to **[]string** | Project IDs the user should be able to access. This field is only required when role is set to &#x60;collaborator&#x60; or &#x60;limited_collaborator&#x60;. | [optional] 

## Methods

### NewMemberUpdateInput

`func NewMemberUpdateInput() *MemberUpdateInput`

NewMemberUpdateInput instantiates a new MemberUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberUpdateInputWithDefaults

`func NewMemberUpdateInputWithDefaults() *MemberUpdateInput`

NewMemberUpdateInputWithDefaults instantiates a new MemberUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *MemberUpdateInput) GetRole() []InvitationRolesInner`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberUpdateInput) GetRoleOk() (*[]InvitationRolesInner, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberUpdateInput) SetRole(v []InvitationRolesInner)`

SetRole sets Role field to given value.

### HasRole

`func (o *MemberUpdateInput) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetBoundRoles

`func (o *MemberUpdateInput) GetBoundRoles() []string`

GetBoundRoles returns the BoundRoles field if non-nil, zero value otherwise.

### GetBoundRolesOk

`func (o *MemberUpdateInput) GetBoundRolesOk() (*[]string, bool)`

GetBoundRolesOk returns a tuple with the BoundRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRoles

`func (o *MemberUpdateInput) SetBoundRoles(v []string)`

SetBoundRoles sets BoundRoles field to given value.

### HasBoundRoles

`func (o *MemberUpdateInput) HasBoundRoles() bool`

HasBoundRoles returns a boolean if a field has been set.

### GetProjectIds

`func (o *MemberUpdateInput) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *MemberUpdateInput) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *MemberUpdateInput) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *MemberUpdateInput) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


