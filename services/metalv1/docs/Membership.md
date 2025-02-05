# Membership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**Href**](Href.md) |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to [**Href**](Href.md) |  | [optional] 
**BoundRoles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMembership

`func NewMembership() *Membership`

NewMembership instantiates a new Membership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipWithDefaults

`func NewMembershipWithDefaults() *Membership`

NewMembershipWithDefaults instantiates a new Membership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Membership) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Membership) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Membership) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Membership) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *Membership) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Membership) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Membership) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Membership) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Membership) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Membership) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Membership) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Membership) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *Membership) GetProject() Href`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Membership) GetProjectOk() (*Href, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Membership) SetProject(v Href)`

SetProject sets Project field to given value.

### HasProject

`func (o *Membership) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRoles

`func (o *Membership) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Membership) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Membership) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Membership) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Membership) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Membership) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Membership) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Membership) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *Membership) GetUser() Href`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Membership) GetUserOk() (*Href, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Membership) SetUser(v Href)`

SetUser sets User field to given value.

### HasUser

`func (o *Membership) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetBoundRoles

`func (o *Membership) GetBoundRoles() []string`

GetBoundRoles returns the BoundRoles field if non-nil, zero value otherwise.

### GetBoundRolesOk

`func (o *Membership) GetBoundRolesOk() (*[]string, bool)`

GetBoundRolesOk returns a tuple with the BoundRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundRoles

`func (o *Membership) SetBoundRoles(v []string)`

SetBoundRoles sets BoundRoles field to given value.

### HasBoundRoles

`func (o *Membership) HasBoundRoles() bool`

HasBoundRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


