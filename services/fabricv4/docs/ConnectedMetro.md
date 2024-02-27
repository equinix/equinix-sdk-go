# ConnectedMetro

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | The Canonical URL at which the resource resides. | [optional] 
**Code** | Pointer to **string** | Code assigned to an Equinix International Business Exchange (IBX) data center in a specified metropolitan area. | [optional] 
**AvgLatency** | Pointer to **float32** | Average latency (in milliseconds[ms]) between two specified metros. | [optional] 
**RemoteVCBandwidthMax** | Pointer to **int64** | This field holds the Max Connection speed with connected metros | [optional] 

## Methods

### NewConnectedMetro

`func NewConnectedMetro() *ConnectedMetro`

NewConnectedMetro instantiates a new ConnectedMetro object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedMetroWithDefaults

`func NewConnectedMetroWithDefaults() *ConnectedMetro`

NewConnectedMetroWithDefaults instantiates a new ConnectedMetro object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ConnectedMetro) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConnectedMetro) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConnectedMetro) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ConnectedMetro) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetCode

`func (o *ConnectedMetro) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ConnectedMetro) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ConnectedMetro) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ConnectedMetro) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAvgLatency

`func (o *ConnectedMetro) GetAvgLatency() float32`

GetAvgLatency returns the AvgLatency field if non-nil, zero value otherwise.

### GetAvgLatencyOk

`func (o *ConnectedMetro) GetAvgLatencyOk() (*float32, bool)`

GetAvgLatencyOk returns a tuple with the AvgLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgLatency

`func (o *ConnectedMetro) SetAvgLatency(v float32)`

SetAvgLatency sets AvgLatency field to given value.

### HasAvgLatency

`func (o *ConnectedMetro) HasAvgLatency() bool`

HasAvgLatency returns a boolean if a field has been set.

### GetRemoteVCBandwidthMax

`func (o *ConnectedMetro) GetRemoteVCBandwidthMax() int64`

GetRemoteVCBandwidthMax returns the RemoteVCBandwidthMax field if non-nil, zero value otherwise.

### GetRemoteVCBandwidthMaxOk

`func (o *ConnectedMetro) GetRemoteVCBandwidthMaxOk() (*int64, bool)`

GetRemoteVCBandwidthMaxOk returns a tuple with the RemoteVCBandwidthMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVCBandwidthMax

`func (o *ConnectedMetro) SetRemoteVCBandwidthMax(v int64)`

SetRemoteVCBandwidthMax sets RemoteVCBandwidthMax field to given value.

### HasRemoteVCBandwidthMax

`func (o *ConnectedMetro) HasRemoteVCBandwidthMax() bool`

HasRemoteVCBandwidthMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


