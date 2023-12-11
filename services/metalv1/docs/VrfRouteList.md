# VrfRouteList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | Pointer to [**[]VrfRoute**](VrfRoute.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewVrfRouteList

`func NewVrfRouteList() *VrfRouteList`

NewVrfRouteList instantiates a new VrfRouteList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfRouteListWithDefaults

`func NewVrfRouteListWithDefaults() *VrfRouteList`

NewVrfRouteListWithDefaults instantiates a new VrfRouteList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutes

`func (o *VrfRouteList) GetRoutes() []VrfRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *VrfRouteList) GetRoutesOk() (*[]VrfRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *VrfRouteList) SetRoutes(v []VrfRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *VrfRouteList) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetMeta

`func (o *VrfRouteList) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VrfRouteList) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VrfRouteList) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VrfRouteList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


