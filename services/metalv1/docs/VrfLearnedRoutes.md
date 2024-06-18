# VrfLearnedRoutes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** |  | [optional] 
**OriginAs** | Pointer to **int64** | The ASN of the peer that advertised the prefix. | [optional] 

## Methods

### NewVrfLearnedRoutes

`func NewVrfLearnedRoutes() *VrfLearnedRoutes`

NewVrfLearnedRoutes instantiates a new VrfLearnedRoutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfLearnedRoutesWithDefaults

`func NewVrfLearnedRoutesWithDefaults() *VrfLearnedRoutes`

NewVrfLearnedRoutesWithDefaults instantiates a new VrfLearnedRoutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *VrfLearnedRoutes) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VrfLearnedRoutes) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VrfLearnedRoutes) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *VrfLearnedRoutes) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetOriginAs

`func (o *VrfLearnedRoutes) GetOriginAs() int64`

GetOriginAs returns the OriginAs field if non-nil, zero value otherwise.

### GetOriginAsOk

`func (o *VrfLearnedRoutes) GetOriginAsOk() (*int64, bool)`

GetOriginAsOk returns a tuple with the OriginAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAs

`func (o *VrfLearnedRoutes) SetOriginAs(v int64)`

SetOriginAs sets OriginAs field to given value.

### HasOriginAs

`func (o *VrfLearnedRoutes) HasOriginAs() bool`

HasOriginAs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


