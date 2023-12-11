# FirmwareSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Firmware Set UUID | [readonly] 
**Name** | **string** | Firmware Set Name | [readonly] 
**CreatedAt** | Pointer to **time.Time** | Datetime when the block was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Datetime when the block was updated. | [optional] [readonly] 
**Attributes** | Pointer to [**[]Attribute**](Attribute.md) | Represents a list of attributes | [optional] 
**ComponentFirmware** | Pointer to [**[]Component**](Component.md) | List of components versions | [optional] 

## Methods

### NewFirmwareSet

`func NewFirmwareSet(uuid string, name string, ) *FirmwareSet`

NewFirmwareSet instantiates a new FirmwareSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareSetWithDefaults

`func NewFirmwareSetWithDefaults() *FirmwareSet`

NewFirmwareSetWithDefaults instantiates a new FirmwareSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *FirmwareSet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FirmwareSet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FirmwareSet) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *FirmwareSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirmwareSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirmwareSet) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *FirmwareSet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FirmwareSet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FirmwareSet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FirmwareSet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FirmwareSet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FirmwareSet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FirmwareSet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FirmwareSet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAttributes

`func (o *FirmwareSet) GetAttributes() []Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FirmwareSet) GetAttributesOk() (*[]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FirmwareSet) SetAttributes(v []Attribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *FirmwareSet) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetComponentFirmware

`func (o *FirmwareSet) GetComponentFirmware() []Component`

GetComponentFirmware returns the ComponentFirmware field if non-nil, zero value otherwise.

### GetComponentFirmwareOk

`func (o *FirmwareSet) GetComponentFirmwareOk() (*[]Component, bool)`

GetComponentFirmwareOk returns a tuple with the ComponentFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentFirmware

`func (o *FirmwareSet) SetComponentFirmware(v []Component)`

SetComponentFirmware sets ComponentFirmware field to given value.

### HasComponentFirmware

`func (o *FirmwareSet) HasComponentFirmware() bool`

HasComponentFirmware returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


