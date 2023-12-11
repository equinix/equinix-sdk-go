# Storage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disks** | Pointer to [**[]Disk**](Disk.md) |  | [optional] 
**Raid** | Pointer to [**[]Raid**](Raid.md) |  | [optional] 
**Filesystems** | Pointer to [**[]Filesystem**](Filesystem.md) |  | [optional] 

## Methods

### NewStorage

`func NewStorage() *Storage`

NewStorage instantiates a new Storage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageWithDefaults

`func NewStorageWithDefaults() *Storage`

NewStorageWithDefaults instantiates a new Storage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisks

`func (o *Storage) GetDisks() []Disk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *Storage) GetDisksOk() (*[]Disk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *Storage) SetDisks(v []Disk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *Storage) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetRaid

`func (o *Storage) GetRaid() []Raid`

GetRaid returns the Raid field if non-nil, zero value otherwise.

### GetRaidOk

`func (o *Storage) GetRaidOk() (*[]Raid, bool)`

GetRaidOk returns a tuple with the Raid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaid

`func (o *Storage) SetRaid(v []Raid)`

SetRaid sets Raid field to given value.

### HasRaid

`func (o *Storage) HasRaid() bool`

HasRaid returns a boolean if a field has been set.

### GetFilesystems

`func (o *Storage) GetFilesystems() []Filesystem`

GetFilesystems returns the Filesystems field if non-nil, zero value otherwise.

### GetFilesystemsOk

`func (o *Storage) GetFilesystemsOk() (*[]Filesystem, bool)`

GetFilesystemsOk returns a tuple with the Filesystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystems

`func (o *Storage) SetFilesystems(v []Filesystem)`

SetFilesystems sets Filesystems field to given value.

### HasFilesystems

`func (o *Storage) HasFilesystems() bool`

HasFilesystems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


