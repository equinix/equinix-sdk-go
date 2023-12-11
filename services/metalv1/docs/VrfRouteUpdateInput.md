# VrfRouteUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** | The IPv4 prefix for the route, in CIDR-style notation. For a static default route, this will always be \&quot;0.0.0.0/0\&quot; | [optional] 
**NextHop** | Pointer to **string** | The IPv4 address within the VRF of the host that will handle this route | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVrfRouteUpdateInput

`func NewVrfRouteUpdateInput() *VrfRouteUpdateInput`

NewVrfRouteUpdateInput instantiates a new VrfRouteUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfRouteUpdateInputWithDefaults

`func NewVrfRouteUpdateInputWithDefaults() *VrfRouteUpdateInput`

NewVrfRouteUpdateInputWithDefaults instantiates a new VrfRouteUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *VrfRouteUpdateInput) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VrfRouteUpdateInput) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VrfRouteUpdateInput) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *VrfRouteUpdateInput) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetNextHop

`func (o *VrfRouteUpdateInput) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *VrfRouteUpdateInput) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *VrfRouteUpdateInput) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *VrfRouteUpdateInput) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetTags

`func (o *VrfRouteUpdateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VrfRouteUpdateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VrfRouteUpdateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VrfRouteUpdateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


