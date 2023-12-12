# UserList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) |  | [optional] 

## Methods

### NewUserList

`func NewUserList() *UserList`

NewUserList instantiates a new UserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserListWithDefaults

`func NewUserListWithDefaults() *UserList`

NewUserListWithDefaults instantiates a new UserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *UserList) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UserList) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UserList) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UserList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUsers

`func (o *UserList) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UserList) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UserList) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UserList) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


