# PortDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Device name | [optional] 
**Redundancy** | Pointer to [**PortDeviceRedundancy**](PortDeviceRedundancy.md) |  | [optional] 
**VcBandwidthMax** | Pointer to **int64** | Maximum bandwidth allowed for connection. | [optional] 

## Methods

### NewPortDevice

`func NewPortDevice() *PortDevice`

NewPortDevice instantiates a new PortDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortDeviceWithDefaults

`func NewPortDeviceWithDefaults() *PortDevice`

NewPortDeviceWithDefaults instantiates a new PortDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PortDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PortDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRedundancy

`func (o *PortDevice) GetRedundancy() PortDeviceRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *PortDevice) GetRedundancyOk() (*PortDeviceRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *PortDevice) SetRedundancy(v PortDeviceRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *PortDevice) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetVcBandwidthMax

`func (o *PortDevice) GetVcBandwidthMax() int64`

GetVcBandwidthMax returns the VcBandwidthMax field if non-nil, zero value otherwise.

### GetVcBandwidthMaxOk

`func (o *PortDevice) GetVcBandwidthMaxOk() (*int64, bool)`

GetVcBandwidthMaxOk returns a tuple with the VcBandwidthMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcBandwidthMax

`func (o *PortDevice) SetVcBandwidthMax(v int64)`

SetVcBandwidthMax sets VcBandwidthMax field to given value.

### HasVcBandwidthMax

`func (o *PortDevice) HasVcBandwidthMax() bool`

HasVcBandwidthMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


