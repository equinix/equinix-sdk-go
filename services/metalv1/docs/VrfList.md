# VrfList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Vrfs** | Pointer to [**[]Vrf**](Vrf.md) |  | [optional] 

## Methods

### NewVrfList

`func NewVrfList() *VrfList`

NewVrfList instantiates a new VrfList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfListWithDefaults

`func NewVrfListWithDefaults() *VrfList`

NewVrfListWithDefaults instantiates a new VrfList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *VrfList) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VrfList) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VrfList) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VrfList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetVrfs

`func (o *VrfList) GetVrfs() []Vrf`

GetVrfs returns the Vrfs field if non-nil, zero value otherwise.

### GetVrfsOk

`func (o *VrfList) GetVrfsOk() (*[]Vrf, bool)`

GetVrfsOk returns a tuple with the Vrfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfs

`func (o *VrfList) SetVrfs(v []Vrf)`

SetVrfs sets Vrfs field to given value.

### HasVrfs

`func (o *VrfList) HasVrfs() bool`

HasVrfs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


