# CoresConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Core** | Pointer to **int32** | The number of cores. | [optional] 
**Memory** | Pointer to **int32** | The amount of memory. | [optional] 
**Unit** | Pointer to **string** | The unit of memory. | [optional] 
**Flavor** | Pointer to **string** | Small, medium or large. | [optional] 
**PackageCodes** | Pointer to [**[]PackageCodes**](PackageCodes.md) | An array that has all the software packages and throughput options. | [optional] 
**Supported** | Pointer to **bool** | Whether or not this core is supported. | [optional] 
**Tier** | Pointer to **int32** | Tier is relevant only for Cisco 8000V devices | [optional] 

## Methods

### NewCoresConfig

`func NewCoresConfig() *CoresConfig`

NewCoresConfig instantiates a new CoresConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoresConfigWithDefaults

`func NewCoresConfigWithDefaults() *CoresConfig`

NewCoresConfigWithDefaults instantiates a new CoresConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCore

`func (o *CoresConfig) GetCore() int32`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *CoresConfig) GetCoreOk() (*int32, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *CoresConfig) SetCore(v int32)`

SetCore sets Core field to given value.

### HasCore

`func (o *CoresConfig) HasCore() bool`

HasCore returns a boolean if a field has been set.

### GetMemory

`func (o *CoresConfig) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CoresConfig) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CoresConfig) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CoresConfig) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetUnit

`func (o *CoresConfig) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CoresConfig) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CoresConfig) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CoresConfig) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetFlavor

`func (o *CoresConfig) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *CoresConfig) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *CoresConfig) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *CoresConfig) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetPackageCodes

`func (o *CoresConfig) GetPackageCodes() []PackageCodes`

GetPackageCodes returns the PackageCodes field if non-nil, zero value otherwise.

### GetPackageCodesOk

`func (o *CoresConfig) GetPackageCodesOk() (*[]PackageCodes, bool)`

GetPackageCodesOk returns a tuple with the PackageCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCodes

`func (o *CoresConfig) SetPackageCodes(v []PackageCodes)`

SetPackageCodes sets PackageCodes field to given value.

### HasPackageCodes

`func (o *CoresConfig) HasPackageCodes() bool`

HasPackageCodes returns a boolean if a field has been set.

### GetSupported

`func (o *CoresConfig) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *CoresConfig) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *CoresConfig) SetSupported(v bool)`

SetSupported sets Supported field to given value.

### HasSupported

`func (o *CoresConfig) HasSupported() bool`

HasSupported returns a boolean if a field has been set.

### GetTier

`func (o *CoresConfig) GetTier() int32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CoresConfig) GetTierOk() (*int32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CoresConfig) SetTier(v int32)`

SetTier sets Tier field to given value.

### HasTier

`func (o *CoresConfig) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


