# Disk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** |  | [optional] 
**WipeTable** | Pointer to **bool** |  | [optional] 
**Partitions** | Pointer to [**[]Partition**](Partition.md) |  | [optional] 

## Methods

### NewDisk

`func NewDisk() *Disk`

NewDisk instantiates a new Disk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskWithDefaults

`func NewDiskWithDefaults() *Disk`

NewDiskWithDefaults instantiates a new Disk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *Disk) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Disk) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Disk) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *Disk) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetWipeTable

`func (o *Disk) GetWipeTable() bool`

GetWipeTable returns the WipeTable field if non-nil, zero value otherwise.

### GetWipeTableOk

`func (o *Disk) GetWipeTableOk() (*bool, bool)`

GetWipeTableOk returns a tuple with the WipeTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWipeTable

`func (o *Disk) SetWipeTable(v bool)`

SetWipeTable sets WipeTable field to given value.

### HasWipeTable

`func (o *Disk) HasWipeTable() bool`

HasWipeTable returns a boolean if a field has been set.

### GetPartitions

`func (o *Disk) GetPartitions() []Partition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *Disk) GetPartitionsOk() (*[]Partition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *Disk) SetPartitions(v []Partition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *Disk) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


