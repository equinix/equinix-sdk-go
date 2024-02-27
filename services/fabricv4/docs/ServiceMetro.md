# ServiceMetro

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | metro code | [optional] 
**Name** | Pointer to **string** | metro name | [optional] 
**VcBandwidthMax** | Pointer to **int32** | max VC speed supported in Mbps | [optional] 
**Ibxs** | Pointer to **[]string** |  | [optional] 
**InTrail** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** | service metro display name | [optional] 
**SellerRegions** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceMetro

`func NewServiceMetro() *ServiceMetro`

NewServiceMetro instantiates a new ServiceMetro object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMetroWithDefaults

`func NewServiceMetroWithDefaults() *ServiceMetro`

NewServiceMetroWithDefaults instantiates a new ServiceMetro object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ServiceMetro) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServiceMetro) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServiceMetro) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ServiceMetro) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ServiceMetro) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceMetro) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceMetro) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceMetro) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVcBandwidthMax

`func (o *ServiceMetro) GetVcBandwidthMax() int32`

GetVcBandwidthMax returns the VcBandwidthMax field if non-nil, zero value otherwise.

### GetVcBandwidthMaxOk

`func (o *ServiceMetro) GetVcBandwidthMaxOk() (*int32, bool)`

GetVcBandwidthMaxOk returns a tuple with the VcBandwidthMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcBandwidthMax

`func (o *ServiceMetro) SetVcBandwidthMax(v int32)`

SetVcBandwidthMax sets VcBandwidthMax field to given value.

### HasVcBandwidthMax

`func (o *ServiceMetro) HasVcBandwidthMax() bool`

HasVcBandwidthMax returns a boolean if a field has been set.

### GetIbxs

`func (o *ServiceMetro) GetIbxs() []string`

GetIbxs returns the Ibxs field if non-nil, zero value otherwise.

### GetIbxsOk

`func (o *ServiceMetro) GetIbxsOk() (*[]string, bool)`

GetIbxsOk returns a tuple with the Ibxs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbxs

`func (o *ServiceMetro) SetIbxs(v []string)`

SetIbxs sets Ibxs field to given value.

### HasIbxs

`func (o *ServiceMetro) HasIbxs() bool`

HasIbxs returns a boolean if a field has been set.

### GetInTrail

`func (o *ServiceMetro) GetInTrail() bool`

GetInTrail returns the InTrail field if non-nil, zero value otherwise.

### GetInTrailOk

`func (o *ServiceMetro) GetInTrailOk() (*bool, bool)`

GetInTrailOk returns a tuple with the InTrail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTrail

`func (o *ServiceMetro) SetInTrail(v bool)`

SetInTrail sets InTrail field to given value.

### HasInTrail

`func (o *ServiceMetro) HasInTrail() bool`

HasInTrail returns a boolean if a field has been set.

### GetDisplayName

`func (o *ServiceMetro) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ServiceMetro) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ServiceMetro) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ServiceMetro) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSellerRegions

`func (o *ServiceMetro) GetSellerRegions() map[string]string`

GetSellerRegions returns the SellerRegions field if non-nil, zero value otherwise.

### GetSellerRegionsOk

`func (o *ServiceMetro) GetSellerRegionsOk() (*map[string]string, bool)`

GetSellerRegionsOk returns a tuple with the SellerRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerRegions

`func (o *ServiceMetro) SetSellerRegions(v map[string]string)`

SetSellerRegions sets SellerRegions field to given value.

### HasSellerRegions

`func (o *ServiceMetro) HasSellerRegions() bool`

HasSellerRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


