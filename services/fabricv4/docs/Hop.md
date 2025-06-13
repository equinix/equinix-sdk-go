# Hop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hop** | Pointer to **int32** |  | [optional] 
**Probes** | Pointer to [**[]HopProbes**](HopProbes.md) |  | [optional] 

## Methods

### NewHop

`func NewHop() *Hop`

NewHop instantiates a new Hop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHopWithDefaults

`func NewHopWithDefaults() *Hop`

NewHopWithDefaults instantiates a new Hop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHop

`func (o *Hop) GetHop() int32`

GetHop returns the Hop field if non-nil, zero value otherwise.

### GetHopOk

`func (o *Hop) GetHopOk() (*int32, bool)`

GetHopOk returns a tuple with the Hop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHop

`func (o *Hop) SetHop(v int32)`

SetHop sets Hop field to given value.

### HasHop

`func (o *Hop) HasHop() bool`

HasHop returns a boolean if a field has been set.

### GetProbes

`func (o *Hop) GetProbes() []HopProbes`

GetProbes returns the Probes field if non-nil, zero value otherwise.

### GetProbesOk

`func (o *Hop) GetProbesOk() (*[]HopProbes, bool)`

GetProbesOk returns a tuple with the Probes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbes

`func (o *Hop) SetProbes(v []HopProbes)`

SetProbes sets Probes field to given value.

### HasProbes

`func (o *Hop) HasProbes() bool`

HasProbes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


