# SshUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshUsername** | Pointer to **string** | sshUsername. This should be minimum 3 and maximum 32 characters and include alphanumeric characters, dash, and underscore. | [optional] 
**SshPassword** | Pointer to **string** | sshPassword | [optional] 
**SshUserUuid** | Pointer to **string** | sshUserUuid | [optional] 
**Action** | Pointer to **string** | action | [optional] 

## Methods

### NewSshUsers

`func NewSshUsers() *SshUsers`

NewSshUsers instantiates a new SshUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUsersWithDefaults

`func NewSshUsersWithDefaults() *SshUsers`

NewSshUsersWithDefaults instantiates a new SshUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshUsername

`func (o *SshUsers) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *SshUsers) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *SshUsers) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *SshUsers) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *SshUsers) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *SshUsers) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *SshUsers) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *SshUsers) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshUserUuid

`func (o *SshUsers) GetSshUserUuid() string`

GetSshUserUuid returns the SshUserUuid field if non-nil, zero value otherwise.

### GetSshUserUuidOk

`func (o *SshUsers) GetSshUserUuidOk() (*string, bool)`

GetSshUserUuidOk returns a tuple with the SshUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUserUuid

`func (o *SshUsers) SetSshUserUuid(v string)`

SetSshUserUuid sets SshUserUuid field to given value.

### HasSshUserUuid

`func (o *SshUsers) HasSshUserUuid() bool`

HasSshUserUuid returns a boolean if a field has been set.

### GetAction

`func (o *SshUsers) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SshUsers) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SshUsers) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SshUsers) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


