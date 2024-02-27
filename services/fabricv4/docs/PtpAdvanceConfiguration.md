# PtpAdvanceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeScale** | Pointer to [**PtpAdvanceConfigurationTimeScale**](PtpAdvanceConfigurationTimeScale.md) |  | [optional] 
**Domain** | Pointer to **int32** |  | [optional] 
**Priority1** | Pointer to **int32** |  | [optional] 
**Priority2** | Pointer to **int32** |  | [optional] 
**LogAnnounceInterval** | Pointer to **int32** | The mean time interval between Announce messages. A shorter interval makes ptp4l react faster to the changes in the master-slave hierarchy. The interval should be the same in the whole domain. It&#39;s specified as a power of two in seconds. The default is 1 (2 seconds). | [optional] 
**LogSyncInterval** | Pointer to **int32** | The mean time interval between Sync messages. A shorter interval may improve accuracy of the local clock. It&#39;s specified as a power of two in seconds. The default is 0 (1 second). | [optional] 
**LogDelayReqInterval** | Pointer to **int32** |  | [optional] 
**TransportMode** | Pointer to [**PtpAdvanceConfigurationTransportMode**](PtpAdvanceConfigurationTransportMode.md) |  | [optional] 
**GrantTime** | Pointer to **int32** | Unicast Grant Time in seconds. For Multicast and Hybrid transport modes, grant time defaults to 300 seconds. For Unicast mode, grant time can be between 30 to 7200. | [optional] 

## Methods

### NewPtpAdvanceConfiguration

`func NewPtpAdvanceConfiguration() *PtpAdvanceConfiguration`

NewPtpAdvanceConfiguration instantiates a new PtpAdvanceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPtpAdvanceConfigurationWithDefaults

`func NewPtpAdvanceConfigurationWithDefaults() *PtpAdvanceConfiguration`

NewPtpAdvanceConfigurationWithDefaults instantiates a new PtpAdvanceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeScale

`func (o *PtpAdvanceConfiguration) GetTimeScale() PtpAdvanceConfigurationTimeScale`

GetTimeScale returns the TimeScale field if non-nil, zero value otherwise.

### GetTimeScaleOk

`func (o *PtpAdvanceConfiguration) GetTimeScaleOk() (*PtpAdvanceConfigurationTimeScale, bool)`

GetTimeScaleOk returns a tuple with the TimeScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeScale

`func (o *PtpAdvanceConfiguration) SetTimeScale(v PtpAdvanceConfigurationTimeScale)`

SetTimeScale sets TimeScale field to given value.

### HasTimeScale

`func (o *PtpAdvanceConfiguration) HasTimeScale() bool`

HasTimeScale returns a boolean if a field has been set.

### GetDomain

`func (o *PtpAdvanceConfiguration) GetDomain() int32`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PtpAdvanceConfiguration) GetDomainOk() (*int32, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PtpAdvanceConfiguration) SetDomain(v int32)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PtpAdvanceConfiguration) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPriority1

`func (o *PtpAdvanceConfiguration) GetPriority1() int32`

GetPriority1 returns the Priority1 field if non-nil, zero value otherwise.

### GetPriority1Ok

`func (o *PtpAdvanceConfiguration) GetPriority1Ok() (*int32, bool)`

GetPriority1Ok returns a tuple with the Priority1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority1

`func (o *PtpAdvanceConfiguration) SetPriority1(v int32)`

SetPriority1 sets Priority1 field to given value.

### HasPriority1

`func (o *PtpAdvanceConfiguration) HasPriority1() bool`

HasPriority1 returns a boolean if a field has been set.

### GetPriority2

`func (o *PtpAdvanceConfiguration) GetPriority2() int32`

GetPriority2 returns the Priority2 field if non-nil, zero value otherwise.

### GetPriority2Ok

`func (o *PtpAdvanceConfiguration) GetPriority2Ok() (*int32, bool)`

GetPriority2Ok returns a tuple with the Priority2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority2

`func (o *PtpAdvanceConfiguration) SetPriority2(v int32)`

SetPriority2 sets Priority2 field to given value.

### HasPriority2

`func (o *PtpAdvanceConfiguration) HasPriority2() bool`

HasPriority2 returns a boolean if a field has been set.

### GetLogAnnounceInterval

`func (o *PtpAdvanceConfiguration) GetLogAnnounceInterval() int32`

GetLogAnnounceInterval returns the LogAnnounceInterval field if non-nil, zero value otherwise.

### GetLogAnnounceIntervalOk

`func (o *PtpAdvanceConfiguration) GetLogAnnounceIntervalOk() (*int32, bool)`

GetLogAnnounceIntervalOk returns a tuple with the LogAnnounceInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAnnounceInterval

`func (o *PtpAdvanceConfiguration) SetLogAnnounceInterval(v int32)`

SetLogAnnounceInterval sets LogAnnounceInterval field to given value.

### HasLogAnnounceInterval

`func (o *PtpAdvanceConfiguration) HasLogAnnounceInterval() bool`

HasLogAnnounceInterval returns a boolean if a field has been set.

### GetLogSyncInterval

`func (o *PtpAdvanceConfiguration) GetLogSyncInterval() int32`

GetLogSyncInterval returns the LogSyncInterval field if non-nil, zero value otherwise.

### GetLogSyncIntervalOk

`func (o *PtpAdvanceConfiguration) GetLogSyncIntervalOk() (*int32, bool)`

GetLogSyncIntervalOk returns a tuple with the LogSyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSyncInterval

`func (o *PtpAdvanceConfiguration) SetLogSyncInterval(v int32)`

SetLogSyncInterval sets LogSyncInterval field to given value.

### HasLogSyncInterval

`func (o *PtpAdvanceConfiguration) HasLogSyncInterval() bool`

HasLogSyncInterval returns a boolean if a field has been set.

### GetLogDelayReqInterval

`func (o *PtpAdvanceConfiguration) GetLogDelayReqInterval() int32`

GetLogDelayReqInterval returns the LogDelayReqInterval field if non-nil, zero value otherwise.

### GetLogDelayReqIntervalOk

`func (o *PtpAdvanceConfiguration) GetLogDelayReqIntervalOk() (*int32, bool)`

GetLogDelayReqIntervalOk returns a tuple with the LogDelayReqInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDelayReqInterval

`func (o *PtpAdvanceConfiguration) SetLogDelayReqInterval(v int32)`

SetLogDelayReqInterval sets LogDelayReqInterval field to given value.

### HasLogDelayReqInterval

`func (o *PtpAdvanceConfiguration) HasLogDelayReqInterval() bool`

HasLogDelayReqInterval returns a boolean if a field has been set.

### GetTransportMode

`func (o *PtpAdvanceConfiguration) GetTransportMode() PtpAdvanceConfigurationTransportMode`

GetTransportMode returns the TransportMode field if non-nil, zero value otherwise.

### GetTransportModeOk

`func (o *PtpAdvanceConfiguration) GetTransportModeOk() (*PtpAdvanceConfigurationTransportMode, bool)`

GetTransportModeOk returns a tuple with the TransportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportMode

`func (o *PtpAdvanceConfiguration) SetTransportMode(v PtpAdvanceConfigurationTransportMode)`

SetTransportMode sets TransportMode field to given value.

### HasTransportMode

`func (o *PtpAdvanceConfiguration) HasTransportMode() bool`

HasTransportMode returns a boolean if a field has been set.

### GetGrantTime

`func (o *PtpAdvanceConfiguration) GetGrantTime() int32`

GetGrantTime returns the GrantTime field if non-nil, zero value otherwise.

### GetGrantTimeOk

`func (o *PtpAdvanceConfiguration) GetGrantTimeOk() (*int32, bool)`

GetGrantTimeOk returns a tuple with the GrantTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTime

`func (o *PtpAdvanceConfiguration) SetGrantTime(v int32)`

SetGrantTime sets GrantTime field to given value.

### HasGrantTime

`func (o *PtpAdvanceConfiguration) HasGrantTime() bool`

HasGrantTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


