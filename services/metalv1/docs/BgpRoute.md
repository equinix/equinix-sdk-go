# BgpRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exact** | Pointer to **bool** |  | [optional] 
**Route** | Pointer to **string** |  | [optional] 

## Methods

### NewBgpRoute

`func NewBgpRoute() *BgpRoute`

NewBgpRoute instantiates a new BgpRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpRouteWithDefaults

`func NewBgpRouteWithDefaults() *BgpRoute`

NewBgpRouteWithDefaults instantiates a new BgpRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExact

`func (o *BgpRoute) GetExact() bool`

GetExact returns the Exact field if non-nil, zero value otherwise.

### GetExactOk

`func (o *BgpRoute) GetExactOk() (*bool, bool)`

GetExactOk returns a tuple with the Exact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExact

`func (o *BgpRoute) SetExact(v bool)`

SetExact sets Exact field to given value.

### HasExact

`func (o *BgpRoute) HasExact() bool`

HasExact returns a boolean if a field has been set.

### GetRoute

`func (o *BgpRoute) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *BgpRoute) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *BgpRoute) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *BgpRoute) HasRoute() bool`

HasRoute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


