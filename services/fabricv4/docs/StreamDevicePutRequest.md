# StreamDevicePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricsEnabled** | Pointer to **bool** | enable metric. Must be true when interfaces are provided | [optional] [default to false]
**Interfaces** | Pointer to [**[]StreamDeviceInterfaceRequest**](StreamDeviceInterfaceRequest.md) | Complete desired set of Network Edge Device interfaces attached to the stream. Interfaces not in the list are detached; an empty or absent list detaches all interfaces.  | [optional] 

## Methods

### NewStreamDevicePutRequest

`func NewStreamDevicePutRequest() *StreamDevicePutRequest`

NewStreamDevicePutRequest instantiates a new StreamDevicePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamDevicePutRequestWithDefaults

`func NewStreamDevicePutRequestWithDefaults() *StreamDevicePutRequest`

NewStreamDevicePutRequestWithDefaults instantiates a new StreamDevicePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricsEnabled

`func (o *StreamDevicePutRequest) GetMetricsEnabled() bool`

GetMetricsEnabled returns the MetricsEnabled field if non-nil, zero value otherwise.

### GetMetricsEnabledOk

`func (o *StreamDevicePutRequest) GetMetricsEnabledOk() (*bool, bool)`

GetMetricsEnabledOk returns a tuple with the MetricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEnabled

`func (o *StreamDevicePutRequest) SetMetricsEnabled(v bool)`

SetMetricsEnabled sets MetricsEnabled field to given value.

### HasMetricsEnabled

`func (o *StreamDevicePutRequest) HasMetricsEnabled() bool`

HasMetricsEnabled returns a boolean if a field has been set.

### GetInterfaces

`func (o *StreamDevicePutRequest) GetInterfaces() []StreamDeviceInterfaceRequest`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *StreamDevicePutRequest) GetInterfacesOk() (*[]StreamDeviceInterfaceRequest, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *StreamDevicePutRequest) SetInterfaces(v []StreamDeviceInterfaceRequest)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *StreamDevicePutRequest) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


