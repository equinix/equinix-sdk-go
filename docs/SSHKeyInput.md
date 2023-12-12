# SSHKeyInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSSHKeyInput

`func NewSSHKeyInput() *SSHKeyInput`

NewSSHKeyInput instantiates a new SSHKeyInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHKeyInputWithDefaults

`func NewSSHKeyInputWithDefaults() *SSHKeyInput`

NewSSHKeyInputWithDefaults instantiates a new SSHKeyInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SSHKeyInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SSHKeyInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SSHKeyInput) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SSHKeyInput) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *SSHKeyInput) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SSHKeyInput) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SSHKeyInput) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SSHKeyInput) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetTags

`func (o *SSHKeyInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SSHKeyInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SSHKeyInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SSHKeyInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


