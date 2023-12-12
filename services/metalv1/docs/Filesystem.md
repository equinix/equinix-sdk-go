# Filesystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mount** | Pointer to [**Mount**](Mount.md) |  | [optional] 

## Methods

### NewFilesystem

`func NewFilesystem() *Filesystem`

NewFilesystem instantiates a new Filesystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemWithDefaults

`func NewFilesystemWithDefaults() *Filesystem`

NewFilesystemWithDefaults instantiates a new Filesystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMount

`func (o *Filesystem) GetMount() Mount`

GetMount returns the Mount field if non-nil, zero value otherwise.

### GetMountOk

`func (o *Filesystem) GetMountOk() (*Mount, bool)`

GetMountOk returns a tuple with the Mount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMount

`func (o *Filesystem) SetMount(v Mount)`

SetMount sets Mount field to given value.

### HasMount

`func (o *Filesystem) HasMount() bool`

HasMount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


