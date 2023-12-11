# SSHKeyCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstancesIds** | Pointer to **[]string** | List of instance UUIDs to associate SSH key with, when empty array is sent all instances belonging       to entity will be included | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSSHKeyCreateInput

`func NewSSHKeyCreateInput() *SSHKeyCreateInput`

NewSSHKeyCreateInput instantiates a new SSHKeyCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHKeyCreateInputWithDefaults

`func NewSSHKeyCreateInputWithDefaults() *SSHKeyCreateInput`

NewSSHKeyCreateInputWithDefaults instantiates a new SSHKeyCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstancesIds

`func (o *SSHKeyCreateInput) GetInstancesIds() []string`

GetInstancesIds returns the InstancesIds field if non-nil, zero value otherwise.

### GetInstancesIdsOk

`func (o *SSHKeyCreateInput) GetInstancesIdsOk() (*[]string, bool)`

GetInstancesIdsOk returns a tuple with the InstancesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesIds

`func (o *SSHKeyCreateInput) SetInstancesIds(v []string)`

SetInstancesIds sets InstancesIds field to given value.

### HasInstancesIds

`func (o *SSHKeyCreateInput) HasInstancesIds() bool`

HasInstancesIds returns a boolean if a field has been set.

### GetKey

`func (o *SSHKeyCreateInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SSHKeyCreateInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SSHKeyCreateInput) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SSHKeyCreateInput) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *SSHKeyCreateInput) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SSHKeyCreateInput) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SSHKeyCreateInput) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SSHKeyCreateInput) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetTags

`func (o *SSHKeyCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SSHKeyCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SSHKeyCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SSHKeyCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


