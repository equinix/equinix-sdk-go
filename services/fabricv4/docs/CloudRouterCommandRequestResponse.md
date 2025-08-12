# CloudRouterCommandRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | Fabric Cloud Router Ping or Traceroute Command Destination | [optional] 
**SourceConnection** | Pointer to [**CloudRouterCommandRequestConnection**](CloudRouterCommandRequestConnection.md) |  | [optional] 
**Timeout** | Pointer to **int32** | Timeout in seconds for Fabric Cloud Router Command:   - For &#x60;PING_COMMAND&#x60;: Packet timeout duration. The default value is 5.   - For &#x60;TRACEROUTE_COMMAND&#x60;: Probe timeout duration.     The default value is 2 and it is not configurable.  | [optional] 
**DataBytes** | Pointer to **int32** | Ping Command DataBytes.  This field is only applicable for commands of type &#x60;PING_COMMAND&#x60;.  | [optional] [default to 64]
**Interval** | Pointer to **int32** | Time in milliseconds between sending each packet. This field is only applicable for commands of type &#x60;PING_COMMAND&#x60;.  | [optional] [default to 1000]
**Count** | Pointer to **int32** | Total number of ping requests. This field is only applicable for commands of type &#x60;PING_COMMAND&#x60;.  | [optional] [default to 5]
**Probes** | Pointer to **int32** | Number of probes for Fabric Cloud Router Traceroute Command. This field is only applicable for commands of type &#x60;TRACEROUTE_COMMAND&#x60; and is not configurable.  | [optional] [default to 3]
**HopsMax** | Pointer to **int32** | Maximum number of hops for the traceroute command. This field is only applicable for commands of type &#x60;TRACEROUTE_COMMAND&#x60;.  | [optional] [default to 20]

## Methods

### NewCloudRouterCommandRequestResponse

`func NewCloudRouterCommandRequestResponse() *CloudRouterCommandRequestResponse`

NewCloudRouterCommandRequestResponse instantiates a new CloudRouterCommandRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterCommandRequestResponseWithDefaults

`func NewCloudRouterCommandRequestResponseWithDefaults() *CloudRouterCommandRequestResponse`

NewCloudRouterCommandRequestResponseWithDefaults instantiates a new CloudRouterCommandRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *CloudRouterCommandRequestResponse) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CloudRouterCommandRequestResponse) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CloudRouterCommandRequestResponse) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *CloudRouterCommandRequestResponse) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetSourceConnection

`func (o *CloudRouterCommandRequestResponse) GetSourceConnection() CloudRouterCommandRequestConnection`

GetSourceConnection returns the SourceConnection field if non-nil, zero value otherwise.

### GetSourceConnectionOk

`func (o *CloudRouterCommandRequestResponse) GetSourceConnectionOk() (*CloudRouterCommandRequestConnection, bool)`

GetSourceConnectionOk returns a tuple with the SourceConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConnection

`func (o *CloudRouterCommandRequestResponse) SetSourceConnection(v CloudRouterCommandRequestConnection)`

SetSourceConnection sets SourceConnection field to given value.

### HasSourceConnection

`func (o *CloudRouterCommandRequestResponse) HasSourceConnection() bool`

HasSourceConnection returns a boolean if a field has been set.

### GetTimeout

`func (o *CloudRouterCommandRequestResponse) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CloudRouterCommandRequestResponse) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CloudRouterCommandRequestResponse) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CloudRouterCommandRequestResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetDataBytes

`func (o *CloudRouterCommandRequestResponse) GetDataBytes() int32`

GetDataBytes returns the DataBytes field if non-nil, zero value otherwise.

### GetDataBytesOk

`func (o *CloudRouterCommandRequestResponse) GetDataBytesOk() (*int32, bool)`

GetDataBytesOk returns a tuple with the DataBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBytes

`func (o *CloudRouterCommandRequestResponse) SetDataBytes(v int32)`

SetDataBytes sets DataBytes field to given value.

### HasDataBytes

`func (o *CloudRouterCommandRequestResponse) HasDataBytes() bool`

HasDataBytes returns a boolean if a field has been set.

### GetInterval

`func (o *CloudRouterCommandRequestResponse) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CloudRouterCommandRequestResponse) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CloudRouterCommandRequestResponse) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CloudRouterCommandRequestResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCount

`func (o *CloudRouterCommandRequestResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloudRouterCommandRequestResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloudRouterCommandRequestResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CloudRouterCommandRequestResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetProbes

`func (o *CloudRouterCommandRequestResponse) GetProbes() int32`

GetProbes returns the Probes field if non-nil, zero value otherwise.

### GetProbesOk

`func (o *CloudRouterCommandRequestResponse) GetProbesOk() (*int32, bool)`

GetProbesOk returns a tuple with the Probes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbes

`func (o *CloudRouterCommandRequestResponse) SetProbes(v int32)`

SetProbes sets Probes field to given value.

### HasProbes

`func (o *CloudRouterCommandRequestResponse) HasProbes() bool`

HasProbes returns a boolean if a field has been set.

### GetHopsMax

`func (o *CloudRouterCommandRequestResponse) GetHopsMax() int32`

GetHopsMax returns the HopsMax field if non-nil, zero value otherwise.

### GetHopsMaxOk

`func (o *CloudRouterCommandRequestResponse) GetHopsMaxOk() (*int32, bool)`

GetHopsMaxOk returns a tuple with the HopsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHopsMax

`func (o *CloudRouterCommandRequestResponse) SetHopsMax(v int32)`

SetHopsMax sets HopsMax field to given value.

### HasHopsMax

`func (o *CloudRouterCommandRequestResponse) HasHopsMax() bool`

HasHopsMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


