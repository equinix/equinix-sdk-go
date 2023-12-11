# VrfIpReservationCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **int32** | The size of the VRF IP Reservation&#39;s subnet | 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Network** | **string** | The starting address for this VRF IP Reservation&#39;s subnet | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | **string** | Must be set to &#39;vrf&#39; | 
**VrfId** | **string** | The ID of the VRF in which this VRF IP Reservation is created. The VRF must have an existing IP Range that contains the requested subnet. This field may be aliased as just &#39;vrf&#39;. | 

## Methods

### NewVrfIpReservationCreateInput

`func NewVrfIpReservationCreateInput(cidr int32, network string, type_ string, vrfId string, ) *VrfIpReservationCreateInput`

NewVrfIpReservationCreateInput instantiates a new VrfIpReservationCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfIpReservationCreateInputWithDefaults

`func NewVrfIpReservationCreateInputWithDefaults() *VrfIpReservationCreateInput`

NewVrfIpReservationCreateInputWithDefaults instantiates a new VrfIpReservationCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *VrfIpReservationCreateInput) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *VrfIpReservationCreateInput) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *VrfIpReservationCreateInput) SetCidr(v int32)`

SetCidr sets Cidr field to given value.


### GetCustomdata

`func (o *VrfIpReservationCreateInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *VrfIpReservationCreateInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *VrfIpReservationCreateInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *VrfIpReservationCreateInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDetails

`func (o *VrfIpReservationCreateInput) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VrfIpReservationCreateInput) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VrfIpReservationCreateInput) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VrfIpReservationCreateInput) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetNetwork

`func (o *VrfIpReservationCreateInput) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VrfIpReservationCreateInput) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VrfIpReservationCreateInput) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTags

`func (o *VrfIpReservationCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VrfIpReservationCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VrfIpReservationCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VrfIpReservationCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *VrfIpReservationCreateInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VrfIpReservationCreateInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VrfIpReservationCreateInput) SetType(v string)`

SetType sets Type field to given value.


### GetVrfId

`func (o *VrfIpReservationCreateInput) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *VrfIpReservationCreateInput) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *VrfIpReservationCreateInput) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


