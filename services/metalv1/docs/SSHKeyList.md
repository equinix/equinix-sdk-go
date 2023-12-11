# SSHKeyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshKeys** | Pointer to [**[]SSHKey**](SSHKey.md) |  | [optional] 

## Methods

### NewSSHKeyList

`func NewSSHKeyList() *SSHKeyList`

NewSSHKeyList instantiates a new SSHKeyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHKeyListWithDefaults

`func NewSSHKeyListWithDefaults() *SSHKeyList`

NewSSHKeyListWithDefaults instantiates a new SSHKeyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshKeys

`func (o *SSHKeyList) GetSshKeys() []SSHKey`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *SSHKeyList) GetSshKeysOk() (*[]SSHKey, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *SSHKeyList) SetSshKeys(v []SSHKey)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *SSHKeyList) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


