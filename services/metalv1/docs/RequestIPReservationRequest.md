# RequestIPReservationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to **string** |  | [optional] 
**FailOnApprovalRequired** | Pointer to **bool** |  | [optional] 
**Metro** | Pointer to **string** | The code of the metro you are requesting the IP reservation in. | [optional] 
**Quantity** | **int32** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | **string** | Must be set to &#39;vrf&#39; | 
**Cidr** | **int32** | The size of the VRF IP Reservation&#39;s subnet. The following subnet sizes are supported: - IPv4: between 22 - 29 inclusive - IPv6: exactly 64  | 
**Network** | **string** | The starting address for this VRF IP Reservation&#39;s subnet. Both IPv4 and IPv6 are supported. | 
**VrfId** | **string** | The ID of the VRF in which this VRF IP Reservation is created. The VRF must have an existing IP Range that contains the requested subnet. This field may be aliased as just &#39;vrf&#39;. | 

## Methods

### NewRequestIPReservationRequest

`func NewRequestIPReservationRequest(quantity int32, type_ string, cidr int32, network string, vrfId string, ) *RequestIPReservationRequest`

NewRequestIPReservationRequest instantiates a new RequestIPReservationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestIPReservationRequestWithDefaults

`func NewRequestIPReservationRequestWithDefaults() *RequestIPReservationRequest`

NewRequestIPReservationRequestWithDefaults instantiates a new RequestIPReservationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *RequestIPReservationRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *RequestIPReservationRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *RequestIPReservationRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *RequestIPReservationRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCustomdata

`func (o *RequestIPReservationRequest) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *RequestIPReservationRequest) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *RequestIPReservationRequest) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *RequestIPReservationRequest) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDetails

`func (o *RequestIPReservationRequest) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RequestIPReservationRequest) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RequestIPReservationRequest) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RequestIPReservationRequest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFacility

`func (o *RequestIPReservationRequest) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *RequestIPReservationRequest) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *RequestIPReservationRequest) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *RequestIPReservationRequest) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetFailOnApprovalRequired

`func (o *RequestIPReservationRequest) GetFailOnApprovalRequired() bool`

GetFailOnApprovalRequired returns the FailOnApprovalRequired field if non-nil, zero value otherwise.

### GetFailOnApprovalRequiredOk

`func (o *RequestIPReservationRequest) GetFailOnApprovalRequiredOk() (*bool, bool)`

GetFailOnApprovalRequiredOk returns a tuple with the FailOnApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOnApprovalRequired

`func (o *RequestIPReservationRequest) SetFailOnApprovalRequired(v bool)`

SetFailOnApprovalRequired sets FailOnApprovalRequired field to given value.

### HasFailOnApprovalRequired

`func (o *RequestIPReservationRequest) HasFailOnApprovalRequired() bool`

HasFailOnApprovalRequired returns a boolean if a field has been set.

### GetMetro

`func (o *RequestIPReservationRequest) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *RequestIPReservationRequest) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *RequestIPReservationRequest) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *RequestIPReservationRequest) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetQuantity

`func (o *RequestIPReservationRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *RequestIPReservationRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *RequestIPReservationRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetTags

`func (o *RequestIPReservationRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RequestIPReservationRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RequestIPReservationRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RequestIPReservationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *RequestIPReservationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestIPReservationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestIPReservationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCidr

`func (o *RequestIPReservationRequest) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *RequestIPReservationRequest) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *RequestIPReservationRequest) SetCidr(v int32)`

SetCidr sets Cidr field to given value.


### GetNetwork

`func (o *RequestIPReservationRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RequestIPReservationRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RequestIPReservationRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetVrfId

`func (o *RequestIPReservationRequest) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *RequestIPReservationRequest) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *RequestIPReservationRequest) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


