# VrfBGPNeighbors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerIp** | Pointer to **string** |  | [optional] 
**PeerAs** | Pointer to **int64** | The ASN of the peer that advertised the prefix. | [optional] 
**State** | Pointer to **string** | The current status of the connection to the BGP peer. State is either up or down. | [optional] 

## Methods

### NewVrfBGPNeighbors

`func NewVrfBGPNeighbors() *VrfBGPNeighbors`

NewVrfBGPNeighbors instantiates a new VrfBGPNeighbors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfBGPNeighborsWithDefaults

`func NewVrfBGPNeighborsWithDefaults() *VrfBGPNeighbors`

NewVrfBGPNeighborsWithDefaults instantiates a new VrfBGPNeighbors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerIp

`func (o *VrfBGPNeighbors) GetPeerIp() string`

GetPeerIp returns the PeerIp field if non-nil, zero value otherwise.

### GetPeerIpOk

`func (o *VrfBGPNeighbors) GetPeerIpOk() (*string, bool)`

GetPeerIpOk returns a tuple with the PeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIp

`func (o *VrfBGPNeighbors) SetPeerIp(v string)`

SetPeerIp sets PeerIp field to given value.

### HasPeerIp

`func (o *VrfBGPNeighbors) HasPeerIp() bool`

HasPeerIp returns a boolean if a field has been set.

### GetPeerAs

`func (o *VrfBGPNeighbors) GetPeerAs() int64`

GetPeerAs returns the PeerAs field if non-nil, zero value otherwise.

### GetPeerAsOk

`func (o *VrfBGPNeighbors) GetPeerAsOk() (*int64, bool)`

GetPeerAsOk returns a tuple with the PeerAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAs

`func (o *VrfBGPNeighbors) SetPeerAs(v int64)`

SetPeerAs sets PeerAs field to given value.

### HasPeerAs

`func (o *VrfBGPNeighbors) HasPeerAs() bool`

HasPeerAs returns a boolean if a field has been set.

### GetState

`func (o *VrfBGPNeighbors) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VrfBGPNeighbors) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VrfBGPNeighbors) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VrfBGPNeighbors) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


