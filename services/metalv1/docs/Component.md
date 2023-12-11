# Component

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Component UUID | [optional] [readonly] 
**Vendor** | Pointer to **string** | Component vendor | [optional] [readonly] 
**Model** | Pointer to **[]string** | List of models where this component version can be applied | [optional] [readonly] 
**Filename** | Pointer to **string** | name of the file | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the component | [optional] [readonly] 
**Component** | Pointer to **string** | Component type | [optional] [readonly] 
**Checksum** | Pointer to **string** | File checksum | [optional] [readonly] 
**UpstreamUrl** | Pointer to **string** | Location of the file | [optional] [readonly] 
**RepositoryUrl** | Pointer to **string** | Location of the file in the repository | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Datetime when the block was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Datetime when the block was updated. | [optional] [readonly] 

## Methods

### NewComponent

`func NewComponent() *Component`

NewComponent instantiates a new Component object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentWithDefaults

`func NewComponentWithDefaults() *Component`

NewComponentWithDefaults instantiates a new Component object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *Component) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Component) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Component) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Component) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendor

`func (o *Component) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Component) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Component) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Component) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetModel

`func (o *Component) GetModel() []string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Component) GetModelOk() (*[]string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Component) SetModel(v []string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Component) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetFilename

`func (o *Component) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Component) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Component) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Component) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetVersion

`func (o *Component) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Component) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Component) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Component) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetComponent

`func (o *Component) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *Component) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *Component) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *Component) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetChecksum

`func (o *Component) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *Component) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *Component) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *Component) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetUpstreamUrl

`func (o *Component) GetUpstreamUrl() string`

GetUpstreamUrl returns the UpstreamUrl field if non-nil, zero value otherwise.

### GetUpstreamUrlOk

`func (o *Component) GetUpstreamUrlOk() (*string, bool)`

GetUpstreamUrlOk returns a tuple with the UpstreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamUrl

`func (o *Component) SetUpstreamUrl(v string)`

SetUpstreamUrl sets UpstreamUrl field to given value.

### HasUpstreamUrl

`func (o *Component) HasUpstreamUrl() bool`

HasUpstreamUrl returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *Component) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *Component) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *Component) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *Component) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Component) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Component) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Component) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Component) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Component) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Component) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Component) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Component) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


