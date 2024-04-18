# CoresDisplayConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Core** | Pointer to **int32** | The number of cores. | [optional] 
**Memory** | Pointer to **int32** | The amount of memory. | [optional] 
**Unit** | Pointer to **string** | The unit of memory. | [optional] 
**Tier** | Pointer to **int32** | Tier is only relevant for Cisco8000V devices. | [optional] 

## Methods

### NewCoresDisplayConfig

`func NewCoresDisplayConfig() *CoresDisplayConfig`

NewCoresDisplayConfig instantiates a new CoresDisplayConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoresDisplayConfigWithDefaults

`func NewCoresDisplayConfigWithDefaults() *CoresDisplayConfig`

NewCoresDisplayConfigWithDefaults instantiates a new CoresDisplayConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCore

`func (o *CoresDisplayConfig) GetCore() int32`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *CoresDisplayConfig) GetCoreOk() (*int32, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *CoresDisplayConfig) SetCore(v int32)`

SetCore sets Core field to given value.

### HasCore

`func (o *CoresDisplayConfig) HasCore() bool`

HasCore returns a boolean if a field has been set.

### GetMemory

`func (o *CoresDisplayConfig) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CoresDisplayConfig) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CoresDisplayConfig) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CoresDisplayConfig) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetUnit

`func (o *CoresDisplayConfig) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CoresDisplayConfig) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CoresDisplayConfig) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CoresDisplayConfig) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetTier

`func (o *CoresDisplayConfig) GetTier() int32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CoresDisplayConfig) GetTierOk() (*int32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CoresDisplayConfig) SetTier(v int32)`

SetTier sets Tier field to given value.

### HasTier

`func (o *CoresDisplayConfig) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


