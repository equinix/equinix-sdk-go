# Partition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 

## Methods

### NewPartition

`func NewPartition() *Partition`

NewPartition instantiates a new Partition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionWithDefaults

`func NewPartitionWithDefaults() *Partition`

NewPartitionWithDefaults instantiates a new Partition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *Partition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Partition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Partition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Partition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetNumber

`func (o *Partition) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Partition) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Partition) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Partition) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetSize

`func (o *Partition) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Partition) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Partition) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Partition) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


