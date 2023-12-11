# VrfRouteCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | **string** | The IPv4 prefix for the route, in CIDR-style notation. For a static default route, this will always be \&quot;0.0.0.0/0\&quot; | 
**NextHop** | **string** | The IPv4 address within the VRF of the host that will handle this route | 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVrfRouteCreateInput

`func NewVrfRouteCreateInput(prefix string, nextHop string, ) *VrfRouteCreateInput`

NewVrfRouteCreateInput instantiates a new VrfRouteCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfRouteCreateInputWithDefaults

`func NewVrfRouteCreateInputWithDefaults() *VrfRouteCreateInput`

NewVrfRouteCreateInputWithDefaults instantiates a new VrfRouteCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *VrfRouteCreateInput) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VrfRouteCreateInput) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VrfRouteCreateInput) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetNextHop

`func (o *VrfRouteCreateInput) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *VrfRouteCreateInput) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *VrfRouteCreateInput) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.


### GetTags

`func (o *VrfRouteCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VrfRouteCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VrfRouteCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VrfRouteCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


