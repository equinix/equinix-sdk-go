# SSHKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Entity** | Pointer to [**Href**](Href.md) |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSSHKey

`func NewSSHKey() *SSHKey`

NewSSHKey instantiates a new SSHKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHKeyWithDefaults

`func NewSSHKeyWithDefaults() *SSHKey`

NewSSHKeyWithDefaults instantiates a new SSHKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SSHKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SSHKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SSHKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SSHKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEntity

`func (o *SSHKey) GetEntity() Href`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SSHKey) GetEntityOk() (*Href, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SSHKey) SetEntity(v Href)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *SSHKey) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetFingerprint

`func (o *SSHKey) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *SSHKey) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *SSHKey) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *SSHKey) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetHref

`func (o *SSHKey) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SSHKey) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SSHKey) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SSHKey) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *SSHKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SSHKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SSHKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SSHKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *SSHKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SSHKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SSHKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SSHKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *SSHKey) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SSHKey) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SSHKey) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SSHKey) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SSHKey) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SSHKey) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SSHKey) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SSHKey) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTags

`func (o *SSHKey) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SSHKey) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SSHKey) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SSHKey) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


