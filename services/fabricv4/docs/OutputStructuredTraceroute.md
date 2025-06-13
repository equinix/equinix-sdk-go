# OutputStructuredTraceroute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationIp** | Pointer to **string** |  | [optional] 
**DestinationName** | Pointer to **string** |  | [optional] 
**PacketBytes** | Pointer to **int32** |  | [optional] 
**HopsMax** | Pointer to **int32** |  | [optional] 
**Hops** | Pointer to [**[]Hop**](Hop.md) |  | [optional] 

## Methods

### NewOutputStructuredTraceroute

`func NewOutputStructuredTraceroute() *OutputStructuredTraceroute`

NewOutputStructuredTraceroute instantiates a new OutputStructuredTraceroute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputStructuredTracerouteWithDefaults

`func NewOutputStructuredTracerouteWithDefaults() *OutputStructuredTraceroute`

NewOutputStructuredTracerouteWithDefaults instantiates a new OutputStructuredTraceroute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationIp

`func (o *OutputStructuredTraceroute) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *OutputStructuredTraceroute) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *OutputStructuredTraceroute) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.

### HasDestinationIp

`func (o *OutputStructuredTraceroute) HasDestinationIp() bool`

HasDestinationIp returns a boolean if a field has been set.

### GetDestinationName

`func (o *OutputStructuredTraceroute) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *OutputStructuredTraceroute) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *OutputStructuredTraceroute) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *OutputStructuredTraceroute) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetPacketBytes

`func (o *OutputStructuredTraceroute) GetPacketBytes() int32`

GetPacketBytes returns the PacketBytes field if non-nil, zero value otherwise.

### GetPacketBytesOk

`func (o *OutputStructuredTraceroute) GetPacketBytesOk() (*int32, bool)`

GetPacketBytesOk returns a tuple with the PacketBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketBytes

`func (o *OutputStructuredTraceroute) SetPacketBytes(v int32)`

SetPacketBytes sets PacketBytes field to given value.

### HasPacketBytes

`func (o *OutputStructuredTraceroute) HasPacketBytes() bool`

HasPacketBytes returns a boolean if a field has been set.

### GetHopsMax

`func (o *OutputStructuredTraceroute) GetHopsMax() int32`

GetHopsMax returns the HopsMax field if non-nil, zero value otherwise.

### GetHopsMaxOk

`func (o *OutputStructuredTraceroute) GetHopsMaxOk() (*int32, bool)`

GetHopsMaxOk returns a tuple with the HopsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHopsMax

`func (o *OutputStructuredTraceroute) SetHopsMax(v int32)`

SetHopsMax sets HopsMax field to given value.

### HasHopsMax

`func (o *OutputStructuredTraceroute) HasHopsMax() bool`

HasHopsMax returns a boolean if a field has been set.

### GetHops

`func (o *OutputStructuredTraceroute) GetHops() []Hop`

GetHops returns the Hops field if non-nil, zero value otherwise.

### GetHopsOk

`func (o *OutputStructuredTraceroute) GetHopsOk() (*[]Hop, bool)`

GetHopsOk returns a tuple with the Hops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHops

`func (o *OutputStructuredTraceroute) SetHops(v []Hop)`

SetHops sets Hops field to given value.

### HasHops

`func (o *OutputStructuredTraceroute) HasHops() bool`

HasHops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


