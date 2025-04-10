# OutputStructuredPing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationIp** | Pointer to **string** |  | [optional] 
**DestinationName** | Pointer to **string** |  | [optional] 
**DataBytes** | Pointer to **int32** |  | [optional] 
**PacketsTransmitted** | Pointer to **int32** |  | [optional] 
**PacketsReceived** | Pointer to **int32** |  | [optional] 
**PacketsLossPercent** | Pointer to **float32** |  | [optional] 
**RttMin** | Pointer to **float32** |  | [optional] 
**RttAvg** | Pointer to **float32** |  | [optional] 
**RttMax** | Pointer to **float32** |  | [optional] 
**RttStdDev** | Pointer to **float32** |  | [optional] 
**Responses** | Pointer to [**[]OutputStructuredPingResponseItem**](OutputStructuredPingResponseItem.md) |  | [optional] 

## Methods

### NewOutputStructuredPing

`func NewOutputStructuredPing() *OutputStructuredPing`

NewOutputStructuredPing instantiates a new OutputStructuredPing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputStructuredPingWithDefaults

`func NewOutputStructuredPingWithDefaults() *OutputStructuredPing`

NewOutputStructuredPingWithDefaults instantiates a new OutputStructuredPing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationIp

`func (o *OutputStructuredPing) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *OutputStructuredPing) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *OutputStructuredPing) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.

### HasDestinationIp

`func (o *OutputStructuredPing) HasDestinationIp() bool`

HasDestinationIp returns a boolean if a field has been set.

### GetDestinationName

`func (o *OutputStructuredPing) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *OutputStructuredPing) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *OutputStructuredPing) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *OutputStructuredPing) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetDataBytes

`func (o *OutputStructuredPing) GetDataBytes() int32`

GetDataBytes returns the DataBytes field if non-nil, zero value otherwise.

### GetDataBytesOk

`func (o *OutputStructuredPing) GetDataBytesOk() (*int32, bool)`

GetDataBytesOk returns a tuple with the DataBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBytes

`func (o *OutputStructuredPing) SetDataBytes(v int32)`

SetDataBytes sets DataBytes field to given value.

### HasDataBytes

`func (o *OutputStructuredPing) HasDataBytes() bool`

HasDataBytes returns a boolean if a field has been set.

### GetPacketsTransmitted

`func (o *OutputStructuredPing) GetPacketsTransmitted() int32`

GetPacketsTransmitted returns the PacketsTransmitted field if non-nil, zero value otherwise.

### GetPacketsTransmittedOk

`func (o *OutputStructuredPing) GetPacketsTransmittedOk() (*int32, bool)`

GetPacketsTransmittedOk returns a tuple with the PacketsTransmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsTransmitted

`func (o *OutputStructuredPing) SetPacketsTransmitted(v int32)`

SetPacketsTransmitted sets PacketsTransmitted field to given value.

### HasPacketsTransmitted

`func (o *OutputStructuredPing) HasPacketsTransmitted() bool`

HasPacketsTransmitted returns a boolean if a field has been set.

### GetPacketsReceived

`func (o *OutputStructuredPing) GetPacketsReceived() int32`

GetPacketsReceived returns the PacketsReceived field if non-nil, zero value otherwise.

### GetPacketsReceivedOk

`func (o *OutputStructuredPing) GetPacketsReceivedOk() (*int32, bool)`

GetPacketsReceivedOk returns a tuple with the PacketsReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsReceived

`func (o *OutputStructuredPing) SetPacketsReceived(v int32)`

SetPacketsReceived sets PacketsReceived field to given value.

### HasPacketsReceived

`func (o *OutputStructuredPing) HasPacketsReceived() bool`

HasPacketsReceived returns a boolean if a field has been set.

### GetPacketsLossPercent

`func (o *OutputStructuredPing) GetPacketsLossPercent() float32`

GetPacketsLossPercent returns the PacketsLossPercent field if non-nil, zero value otherwise.

### GetPacketsLossPercentOk

`func (o *OutputStructuredPing) GetPacketsLossPercentOk() (*float32, bool)`

GetPacketsLossPercentOk returns a tuple with the PacketsLossPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsLossPercent

`func (o *OutputStructuredPing) SetPacketsLossPercent(v float32)`

SetPacketsLossPercent sets PacketsLossPercent field to given value.

### HasPacketsLossPercent

`func (o *OutputStructuredPing) HasPacketsLossPercent() bool`

HasPacketsLossPercent returns a boolean if a field has been set.

### GetRttMin

`func (o *OutputStructuredPing) GetRttMin() float32`

GetRttMin returns the RttMin field if non-nil, zero value otherwise.

### GetRttMinOk

`func (o *OutputStructuredPing) GetRttMinOk() (*float32, bool)`

GetRttMinOk returns a tuple with the RttMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRttMin

`func (o *OutputStructuredPing) SetRttMin(v float32)`

SetRttMin sets RttMin field to given value.

### HasRttMin

`func (o *OutputStructuredPing) HasRttMin() bool`

HasRttMin returns a boolean if a field has been set.

### GetRttAvg

`func (o *OutputStructuredPing) GetRttAvg() float32`

GetRttAvg returns the RttAvg field if non-nil, zero value otherwise.

### GetRttAvgOk

`func (o *OutputStructuredPing) GetRttAvgOk() (*float32, bool)`

GetRttAvgOk returns a tuple with the RttAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRttAvg

`func (o *OutputStructuredPing) SetRttAvg(v float32)`

SetRttAvg sets RttAvg field to given value.

### HasRttAvg

`func (o *OutputStructuredPing) HasRttAvg() bool`

HasRttAvg returns a boolean if a field has been set.

### GetRttMax

`func (o *OutputStructuredPing) GetRttMax() float32`

GetRttMax returns the RttMax field if non-nil, zero value otherwise.

### GetRttMaxOk

`func (o *OutputStructuredPing) GetRttMaxOk() (*float32, bool)`

GetRttMaxOk returns a tuple with the RttMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRttMax

`func (o *OutputStructuredPing) SetRttMax(v float32)`

SetRttMax sets RttMax field to given value.

### HasRttMax

`func (o *OutputStructuredPing) HasRttMax() bool`

HasRttMax returns a boolean if a field has been set.

### GetRttStdDev

`func (o *OutputStructuredPing) GetRttStdDev() float32`

GetRttStdDev returns the RttStdDev field if non-nil, zero value otherwise.

### GetRttStdDevOk

`func (o *OutputStructuredPing) GetRttStdDevOk() (*float32, bool)`

GetRttStdDevOk returns a tuple with the RttStdDev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRttStdDev

`func (o *OutputStructuredPing) SetRttStdDev(v float32)`

SetRttStdDev sets RttStdDev field to given value.

### HasRttStdDev

`func (o *OutputStructuredPing) HasRttStdDev() bool`

HasRttStdDev returns a boolean if a field has been set.

### GetResponses

`func (o *OutputStructuredPing) GetResponses() []OutputStructuredPingResponseItem`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *OutputStructuredPing) GetResponsesOk() (*[]OutputStructuredPingResponseItem, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *OutputStructuredPing) SetResponses(v []OutputStructuredPingResponseItem)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *OutputStructuredPing) HasResponses() bool`

HasResponses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


