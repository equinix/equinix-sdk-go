# InvitationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invitations** | Pointer to [**[]Membership**](Membership.md) |  | [optional] 

## Methods

### NewInvitationList

`func NewInvitationList() *InvitationList`

NewInvitationList instantiates a new InvitationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationListWithDefaults

`func NewInvitationListWithDefaults() *InvitationList`

NewInvitationListWithDefaults instantiates a new InvitationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitations

`func (o *InvitationList) GetInvitations() []Membership`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *InvitationList) GetInvitationsOk() (*[]Membership, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *InvitationList) SetInvitations(v []Membership)`

SetInvitations sets Invitations field to given value.

### HasInvitations

`func (o *InvitationList) HasInvitations() bool`

HasInvitations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


