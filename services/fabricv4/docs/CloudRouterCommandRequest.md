# CloudRouterCommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | **string** | Fabric Cloud Router Ping Command Destination | 
**SourceConnection** | Pointer to [**CloudRouterCommandRequestConnection**](CloudRouterCommandRequestConnection.md) |  | [optional] 
**Timeout** | Pointer to **int32** | Fabric Cloud Router Ping Command Timeout | [optional] 
**DataBytes** | Pointer to **int32** | Fabric Cloud Router Ping Command DataBytes | [optional] [default to 64]
**Interval** | Pointer to **int32** | Time in milliseconds between sending each packet | [optional] [readonly] [default to 1000]
**Count** | Pointer to **int32** | Total number of ping requests | [optional] [readonly] [default to 5]

## Methods

### NewCloudRouterCommandRequest

`func NewCloudRouterCommandRequest(destination string, ) *CloudRouterCommandRequest`

NewCloudRouterCommandRequest instantiates a new CloudRouterCommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterCommandRequestWithDefaults

`func NewCloudRouterCommandRequestWithDefaults() *CloudRouterCommandRequest`

NewCloudRouterCommandRequestWithDefaults instantiates a new CloudRouterCommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *CloudRouterCommandRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CloudRouterCommandRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CloudRouterCommandRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetSourceConnection

`func (o *CloudRouterCommandRequest) GetSourceConnection() CloudRouterCommandRequestConnection`

GetSourceConnection returns the SourceConnection field if non-nil, zero value otherwise.

### GetSourceConnectionOk

`func (o *CloudRouterCommandRequest) GetSourceConnectionOk() (*CloudRouterCommandRequestConnection, bool)`

GetSourceConnectionOk returns a tuple with the SourceConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConnection

`func (o *CloudRouterCommandRequest) SetSourceConnection(v CloudRouterCommandRequestConnection)`

SetSourceConnection sets SourceConnection field to given value.

### HasSourceConnection

`func (o *CloudRouterCommandRequest) HasSourceConnection() bool`

HasSourceConnection returns a boolean if a field has been set.

### GetTimeout

`func (o *CloudRouterCommandRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CloudRouterCommandRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CloudRouterCommandRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CloudRouterCommandRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetDataBytes

`func (o *CloudRouterCommandRequest) GetDataBytes() int32`

GetDataBytes returns the DataBytes field if non-nil, zero value otherwise.

### GetDataBytesOk

`func (o *CloudRouterCommandRequest) GetDataBytesOk() (*int32, bool)`

GetDataBytesOk returns a tuple with the DataBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBytes

`func (o *CloudRouterCommandRequest) SetDataBytes(v int32)`

SetDataBytes sets DataBytes field to given value.

### HasDataBytes

`func (o *CloudRouterCommandRequest) HasDataBytes() bool`

HasDataBytes returns a boolean if a field has been set.

### GetInterval

`func (o *CloudRouterCommandRequest) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CloudRouterCommandRequest) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CloudRouterCommandRequest) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CloudRouterCommandRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCount

`func (o *CloudRouterCommandRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloudRouterCommandRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloudRouterCommandRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CloudRouterCommandRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


