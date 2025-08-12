# CloudRouterCommandRequestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | **string** | Fabric Cloud Router Ping or Traceroute Command Destination | 
**SourceConnection** | [**CloudRouterCommandRequestConnection**](CloudRouterCommandRequestConnection.md) |  | 
**Timeout** | Pointer to **int32** | Timeout in seconds for Fabric Cloud Router Command:   - For &#x60;PING_COMMAND&#x60;: Packet timeout duration. The default value is 5.   - For &#x60;TRACEROUTE_COMMAND&#x60;: Probe timeout duration.     The default value is 2 and it is not configurable.  | [optional] 
**DataBytes** | Pointer to **int32** | Ping Command DataBytes.  This field is only applicable for commands of type &#x60;PING_COMMAND&#x60;.  | [optional] [default to 64]
**Probes** | Pointer to **int32** | Number of probes for Fabric Cloud Router Traceroute Command. This field is only applicable for commands of type &#x60;TRACEROUTE_COMMAND&#x60; and is not configurable.  | [optional] [default to 3]
**HopsMax** | Pointer to **int32** | Maximum number of hops for the traceroute command. This field is only applicable for commands of type &#x60;TRACEROUTE_COMMAND&#x60;.  | [optional] [default to 20]

## Methods

### NewCloudRouterCommandRequestPayload

`func NewCloudRouterCommandRequestPayload(destination string, sourceConnection CloudRouterCommandRequestConnection, ) *CloudRouterCommandRequestPayload`

NewCloudRouterCommandRequestPayload instantiates a new CloudRouterCommandRequestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterCommandRequestPayloadWithDefaults

`func NewCloudRouterCommandRequestPayloadWithDefaults() *CloudRouterCommandRequestPayload`

NewCloudRouterCommandRequestPayloadWithDefaults instantiates a new CloudRouterCommandRequestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *CloudRouterCommandRequestPayload) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CloudRouterCommandRequestPayload) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CloudRouterCommandRequestPayload) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetSourceConnection

`func (o *CloudRouterCommandRequestPayload) GetSourceConnection() CloudRouterCommandRequestConnection`

GetSourceConnection returns the SourceConnection field if non-nil, zero value otherwise.

### GetSourceConnectionOk

`func (o *CloudRouterCommandRequestPayload) GetSourceConnectionOk() (*CloudRouterCommandRequestConnection, bool)`

GetSourceConnectionOk returns a tuple with the SourceConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConnection

`func (o *CloudRouterCommandRequestPayload) SetSourceConnection(v CloudRouterCommandRequestConnection)`

SetSourceConnection sets SourceConnection field to given value.


### GetTimeout

`func (o *CloudRouterCommandRequestPayload) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CloudRouterCommandRequestPayload) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CloudRouterCommandRequestPayload) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CloudRouterCommandRequestPayload) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetDataBytes

`func (o *CloudRouterCommandRequestPayload) GetDataBytes() int32`

GetDataBytes returns the DataBytes field if non-nil, zero value otherwise.

### GetDataBytesOk

`func (o *CloudRouterCommandRequestPayload) GetDataBytesOk() (*int32, bool)`

GetDataBytesOk returns a tuple with the DataBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBytes

`func (o *CloudRouterCommandRequestPayload) SetDataBytes(v int32)`

SetDataBytes sets DataBytes field to given value.

### HasDataBytes

`func (o *CloudRouterCommandRequestPayload) HasDataBytes() bool`

HasDataBytes returns a boolean if a field has been set.

### GetProbes

`func (o *CloudRouterCommandRequestPayload) GetProbes() int32`

GetProbes returns the Probes field if non-nil, zero value otherwise.

### GetProbesOk

`func (o *CloudRouterCommandRequestPayload) GetProbesOk() (*int32, bool)`

GetProbesOk returns a tuple with the Probes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbes

`func (o *CloudRouterCommandRequestPayload) SetProbes(v int32)`

SetProbes sets Probes field to given value.

### HasProbes

`func (o *CloudRouterCommandRequestPayload) HasProbes() bool`

HasProbes returns a boolean if a field has been set.

### GetHopsMax

`func (o *CloudRouterCommandRequestPayload) GetHopsMax() int32`

GetHopsMax returns the HopsMax field if non-nil, zero value otherwise.

### GetHopsMaxOk

`func (o *CloudRouterCommandRequestPayload) GetHopsMaxOk() (*int32, bool)`

GetHopsMaxOk returns a tuple with the HopsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHopsMax

`func (o *CloudRouterCommandRequestPayload) SetHopsMax(v int32)`

SetHopsMax sets HopsMax field to given value.

### HasHopsMax

`func (o *CloudRouterCommandRequestPayload) HasHopsMax() bool`

HasHopsMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


