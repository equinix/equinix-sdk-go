# VirtualPortPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Unique identifier assigned to the virtual port. Either the uuid or the remaining attributes must be supplied. | [optional] 
**Type** | Pointer to [**VirtualPortType**](VirtualPortType.md) |  | [optional] 
**Location** | Pointer to [**VirtualPortLocation**](VirtualPortLocation.md) |  | [optional] 
**Lag** | Pointer to [**LinkAggregationGroup**](LinkAggregationGroup.md) |  | [optional] 
**PhysicalPortsQuantity** | Pointer to **int32** | Number of physical ports requested. The defaults is 1. | [optional] [default to 1]
**Bandwidth** | Pointer to **int32** | Aggregated data transfer capacity,  expressed as follows &lt;br&gt; -&gt; Mbps, megabits (1 million bits) per second &lt;br&gt; -&gt; Gbps, gigabits (1 billion bits) per second &lt;br&gt; Bandwidth must be divisible by physicalPortsQuantity. | [optional] 
**Redundancy** | Pointer to [**VirtualPortRedundancy**](VirtualPortRedundancy.md) |  | [optional] 
**ConnectivitySource** | Pointer to [**ConnectivitySource**](ConnectivitySource.md) |  | [optional] 
**ServiceType** | Pointer to [**VirtualPortServiceType**](VirtualPortServiceType.md) |  | [optional] [default to VIRTUALPORTSERVICETYPE_MSP]
**Settings** | Pointer to [**VirtualPortConfiguration**](VirtualPortConfiguration.md) |  | [optional] 

## Methods

### NewVirtualPortPrice

`func NewVirtualPortPrice() *VirtualPortPrice`

NewVirtualPortPrice instantiates a new VirtualPortPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualPortPriceWithDefaults

`func NewVirtualPortPriceWithDefaults() *VirtualPortPrice`

NewVirtualPortPriceWithDefaults instantiates a new VirtualPortPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *VirtualPortPrice) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualPortPrice) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualPortPrice) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualPortPrice) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *VirtualPortPrice) GetType() VirtualPortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualPortPrice) GetTypeOk() (*VirtualPortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualPortPrice) SetType(v VirtualPortType)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualPortPrice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLocation

`func (o *VirtualPortPrice) GetLocation() VirtualPortLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VirtualPortPrice) GetLocationOk() (*VirtualPortLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VirtualPortPrice) SetLocation(v VirtualPortLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VirtualPortPrice) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLag

`func (o *VirtualPortPrice) GetLag() LinkAggregationGroup`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *VirtualPortPrice) GetLagOk() (*LinkAggregationGroup, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *VirtualPortPrice) SetLag(v LinkAggregationGroup)`

SetLag sets Lag field to given value.

### HasLag

`func (o *VirtualPortPrice) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetPhysicalPortsQuantity

`func (o *VirtualPortPrice) GetPhysicalPortsQuantity() int32`

GetPhysicalPortsQuantity returns the PhysicalPortsQuantity field if non-nil, zero value otherwise.

### GetPhysicalPortsQuantityOk

`func (o *VirtualPortPrice) GetPhysicalPortsQuantityOk() (*int32, bool)`

GetPhysicalPortsQuantityOk returns a tuple with the PhysicalPortsQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsQuantity

`func (o *VirtualPortPrice) SetPhysicalPortsQuantity(v int32)`

SetPhysicalPortsQuantity sets PhysicalPortsQuantity field to given value.

### HasPhysicalPortsQuantity

`func (o *VirtualPortPrice) HasPhysicalPortsQuantity() bool`

HasPhysicalPortsQuantity returns a boolean if a field has been set.

### GetBandwidth

`func (o *VirtualPortPrice) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *VirtualPortPrice) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *VirtualPortPrice) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *VirtualPortPrice) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetRedundancy

`func (o *VirtualPortPrice) GetRedundancy() VirtualPortRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *VirtualPortPrice) GetRedundancyOk() (*VirtualPortRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *VirtualPortPrice) SetRedundancy(v VirtualPortRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *VirtualPortPrice) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetConnectivitySource

`func (o *VirtualPortPrice) GetConnectivitySource() ConnectivitySource`

GetConnectivitySource returns the ConnectivitySource field if non-nil, zero value otherwise.

### GetConnectivitySourceOk

`func (o *VirtualPortPrice) GetConnectivitySourceOk() (*ConnectivitySource, bool)`

GetConnectivitySourceOk returns a tuple with the ConnectivitySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivitySource

`func (o *VirtualPortPrice) SetConnectivitySource(v ConnectivitySource)`

SetConnectivitySource sets ConnectivitySource field to given value.

### HasConnectivitySource

`func (o *VirtualPortPrice) HasConnectivitySource() bool`

HasConnectivitySource returns a boolean if a field has been set.

### GetServiceType

`func (o *VirtualPortPrice) GetServiceType() VirtualPortServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *VirtualPortPrice) GetServiceTypeOk() (*VirtualPortServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *VirtualPortPrice) SetServiceType(v VirtualPortServiceType)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *VirtualPortPrice) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSettings

`func (o *VirtualPortPrice) GetSettings() VirtualPortConfiguration`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *VirtualPortPrice) GetSettingsOk() (*VirtualPortConfiguration, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *VirtualPortPrice) SetSettings(v VirtualPortConfiguration)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *VirtualPortPrice) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


