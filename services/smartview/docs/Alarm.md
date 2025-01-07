# Alarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | [**AlarmAssetDetails**](AlarmAssetDetails.md) |  | 
**ConditionName** | Pointer to **string** | condition name | [optional] 
**Country** | Pointer to **string** | country | [optional] 
**CurrentValue** | [**AlarmCurrentValueDetails**](AlarmCurrentValueDetails.md) |  | 
**DataQuality** | Pointer to **string** | data quality: Good | Bad | Uncertain | [optional] 
**DefinitionId** | **string** | unique message id | 
**Heartbeat** | Pointer to **bool** | heartbeat | [optional] 
**Ibx** | **string** | ibx | 
**Metro** | Pointer to **string** | metro | [optional] 
**NormalProcessedTime** | Pointer to **string** | alarm normal processed time | [optional] 
**NormalTriggeredTime** | Pointer to **string** | alarm normal triggered time | [optional] 
**ProcessedTime** | Pointer to **string** | alarm processed time | [optional] 
**Region** | Pointer to **string** | region | [optional] 
**Severity** | Pointer to **int32** | severity | [optional] 
**Status** | [**AlarmStatusDetails**](AlarmStatusDetails.md) |  | 
**StreamId** | **string** | unique message id | 
**Tag** | [**AlarmTagDetails**](AlarmTagDetails.md) |  | 
**Threshold** | [**AlarmThresholdDetails**](AlarmThresholdDetails.md) |  | 
**TriggerRule** | Pointer to **string** | trigger rule | [optional] 
**TriggeredTime** | Pointer to **string** | alarm triggered time | [optional] 
**Type** | Pointer to **string** | type | [optional] 

## Methods

### NewAlarm

`func NewAlarm(asset AlarmAssetDetails, currentValue AlarmCurrentValueDetails, definitionId string, ibx string, status AlarmStatusDetails, streamId string, tag AlarmTagDetails, threshold AlarmThresholdDetails, ) *Alarm`

NewAlarm instantiates a new Alarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmWithDefaults

`func NewAlarmWithDefaults() *Alarm`

NewAlarmWithDefaults instantiates a new Alarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *Alarm) GetAsset() AlarmAssetDetails`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *Alarm) GetAssetOk() (*AlarmAssetDetails, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *Alarm) SetAsset(v AlarmAssetDetails)`

SetAsset sets Asset field to given value.


### GetConditionName

`func (o *Alarm) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *Alarm) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *Alarm) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.

### HasConditionName

`func (o *Alarm) HasConditionName() bool`

HasConditionName returns a boolean if a field has been set.

### GetCountry

`func (o *Alarm) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Alarm) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Alarm) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Alarm) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrentValue

`func (o *Alarm) GetCurrentValue() AlarmCurrentValueDetails`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *Alarm) GetCurrentValueOk() (*AlarmCurrentValueDetails, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *Alarm) SetCurrentValue(v AlarmCurrentValueDetails)`

SetCurrentValue sets CurrentValue field to given value.


### GetDataQuality

`func (o *Alarm) GetDataQuality() string`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *Alarm) GetDataQualityOk() (*string, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *Alarm) SetDataQuality(v string)`

SetDataQuality sets DataQuality field to given value.

### HasDataQuality

`func (o *Alarm) HasDataQuality() bool`

HasDataQuality returns a boolean if a field has been set.

### GetDefinitionId

`func (o *Alarm) GetDefinitionId() string`

GetDefinitionId returns the DefinitionId field if non-nil, zero value otherwise.

### GetDefinitionIdOk

`func (o *Alarm) GetDefinitionIdOk() (*string, bool)`

GetDefinitionIdOk returns a tuple with the DefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionId

`func (o *Alarm) SetDefinitionId(v string)`

SetDefinitionId sets DefinitionId field to given value.


### GetHeartbeat

`func (o *Alarm) GetHeartbeat() bool`

GetHeartbeat returns the Heartbeat field if non-nil, zero value otherwise.

### GetHeartbeatOk

`func (o *Alarm) GetHeartbeatOk() (*bool, bool)`

GetHeartbeatOk returns a tuple with the Heartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeat

`func (o *Alarm) SetHeartbeat(v bool)`

SetHeartbeat sets Heartbeat field to given value.

### HasHeartbeat

`func (o *Alarm) HasHeartbeat() bool`

HasHeartbeat returns a boolean if a field has been set.

### GetIbx

`func (o *Alarm) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *Alarm) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *Alarm) SetIbx(v string)`

SetIbx sets Ibx field to given value.


### GetMetro

`func (o *Alarm) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *Alarm) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *Alarm) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *Alarm) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNormalProcessedTime

`func (o *Alarm) GetNormalProcessedTime() string`

GetNormalProcessedTime returns the NormalProcessedTime field if non-nil, zero value otherwise.

### GetNormalProcessedTimeOk

`func (o *Alarm) GetNormalProcessedTimeOk() (*string, bool)`

GetNormalProcessedTimeOk returns a tuple with the NormalProcessedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalProcessedTime

`func (o *Alarm) SetNormalProcessedTime(v string)`

SetNormalProcessedTime sets NormalProcessedTime field to given value.

### HasNormalProcessedTime

`func (o *Alarm) HasNormalProcessedTime() bool`

HasNormalProcessedTime returns a boolean if a field has been set.

### GetNormalTriggeredTime

`func (o *Alarm) GetNormalTriggeredTime() string`

GetNormalTriggeredTime returns the NormalTriggeredTime field if non-nil, zero value otherwise.

### GetNormalTriggeredTimeOk

`func (o *Alarm) GetNormalTriggeredTimeOk() (*string, bool)`

GetNormalTriggeredTimeOk returns a tuple with the NormalTriggeredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalTriggeredTime

`func (o *Alarm) SetNormalTriggeredTime(v string)`

SetNormalTriggeredTime sets NormalTriggeredTime field to given value.

### HasNormalTriggeredTime

`func (o *Alarm) HasNormalTriggeredTime() bool`

HasNormalTriggeredTime returns a boolean if a field has been set.

### GetProcessedTime

`func (o *Alarm) GetProcessedTime() string`

GetProcessedTime returns the ProcessedTime field if non-nil, zero value otherwise.

### GetProcessedTimeOk

`func (o *Alarm) GetProcessedTimeOk() (*string, bool)`

GetProcessedTimeOk returns a tuple with the ProcessedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedTime

`func (o *Alarm) SetProcessedTime(v string)`

SetProcessedTime sets ProcessedTime field to given value.

### HasProcessedTime

`func (o *Alarm) HasProcessedTime() bool`

HasProcessedTime returns a boolean if a field has been set.

### GetRegion

`func (o *Alarm) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Alarm) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Alarm) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Alarm) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSeverity

`func (o *Alarm) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Alarm) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Alarm) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Alarm) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStatus

`func (o *Alarm) GetStatus() AlarmStatusDetails`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Alarm) GetStatusOk() (*AlarmStatusDetails, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Alarm) SetStatus(v AlarmStatusDetails)`

SetStatus sets Status field to given value.


### GetStreamId

`func (o *Alarm) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *Alarm) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *Alarm) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.


### GetTag

`func (o *Alarm) GetTag() AlarmTagDetails`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Alarm) GetTagOk() (*AlarmTagDetails, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Alarm) SetTag(v AlarmTagDetails)`

SetTag sets Tag field to given value.


### GetThreshold

`func (o *Alarm) GetThreshold() AlarmThresholdDetails`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Alarm) GetThresholdOk() (*AlarmThresholdDetails, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Alarm) SetThreshold(v AlarmThresholdDetails)`

SetThreshold sets Threshold field to given value.


### GetTriggerRule

`func (o *Alarm) GetTriggerRule() string`

GetTriggerRule returns the TriggerRule field if non-nil, zero value otherwise.

### GetTriggerRuleOk

`func (o *Alarm) GetTriggerRuleOk() (*string, bool)`

GetTriggerRuleOk returns a tuple with the TriggerRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerRule

`func (o *Alarm) SetTriggerRule(v string)`

SetTriggerRule sets TriggerRule field to given value.

### HasTriggerRule

`func (o *Alarm) HasTriggerRule() bool`

HasTriggerRule returns a boolean if a field has been set.

### GetTriggeredTime

`func (o *Alarm) GetTriggeredTime() string`

GetTriggeredTime returns the TriggeredTime field if non-nil, zero value otherwise.

### GetTriggeredTimeOk

`func (o *Alarm) GetTriggeredTimeOk() (*string, bool)`

GetTriggeredTimeOk returns a tuple with the TriggeredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredTime

`func (o *Alarm) SetTriggeredTime(v string)`

SetTriggeredTime sets TriggeredTime field to given value.

### HasTriggeredTime

`func (o *Alarm) HasTriggeredTime() bool`

HasTriggeredTime returns a boolean if a field has been set.

### GetType

`func (o *Alarm) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Alarm) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Alarm) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Alarm) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


