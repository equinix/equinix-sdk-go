# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | [**AlertAssetDetails**](AlertAssetDetails.md) |  | 
**Conditional** | Pointer to **string** | conditional | [optional] 
**DataQuality** | Pointer to **string** | data quality: Good | Bad | Uncertain | [optional] 
**EventType** | Pointer to **string** | event type | [optional] 
**Heartbeat** | Pointer to **bool** | heartbeat | [optional] 
**Ibx** | **string** | ibx | 
**Id** | Pointer to **string** | id | [optional] 
**Region** | Pointer to **string** | region | [optional] 
**StreamId** | **string** | unique message id | 
**Tag** | [**AlertTagDetails**](AlertTagDetails.md) |  | 
**Threshold** | [**AlertThresholdDetails**](AlertThresholdDetails.md) |  | 
**TriggeredTime** | Pointer to **string** | alert triggered time | [optional] 
**Type** | Pointer to **string** | type | [optional] 
**TypeId** | Pointer to **string** | type id | [optional] 

## Methods

### NewAlert

`func NewAlert(asset AlertAssetDetails, ibx string, streamId string, tag AlertTagDetails, threshold AlertThresholdDetails, ) *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *Alert) GetAsset() AlertAssetDetails`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *Alert) GetAssetOk() (*AlertAssetDetails, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *Alert) SetAsset(v AlertAssetDetails)`

SetAsset sets Asset field to given value.


### GetConditional

`func (o *Alert) GetConditional() string`

GetConditional returns the Conditional field if non-nil, zero value otherwise.

### GetConditionalOk

`func (o *Alert) GetConditionalOk() (*string, bool)`

GetConditionalOk returns a tuple with the Conditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditional

`func (o *Alert) SetConditional(v string)`

SetConditional sets Conditional field to given value.

### HasConditional

`func (o *Alert) HasConditional() bool`

HasConditional returns a boolean if a field has been set.

### GetDataQuality

`func (o *Alert) GetDataQuality() string`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *Alert) GetDataQualityOk() (*string, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *Alert) SetDataQuality(v string)`

SetDataQuality sets DataQuality field to given value.

### HasDataQuality

`func (o *Alert) HasDataQuality() bool`

HasDataQuality returns a boolean if a field has been set.

### GetEventType

`func (o *Alert) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Alert) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Alert) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *Alert) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetHeartbeat

`func (o *Alert) GetHeartbeat() bool`

GetHeartbeat returns the Heartbeat field if non-nil, zero value otherwise.

### GetHeartbeatOk

`func (o *Alert) GetHeartbeatOk() (*bool, bool)`

GetHeartbeatOk returns a tuple with the Heartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeat

`func (o *Alert) SetHeartbeat(v bool)`

SetHeartbeat sets Heartbeat field to given value.

### HasHeartbeat

`func (o *Alert) HasHeartbeat() bool`

HasHeartbeat returns a boolean if a field has been set.

### GetIbx

`func (o *Alert) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *Alert) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *Alert) SetIbx(v string)`

SetIbx sets Ibx field to given value.


### GetId

`func (o *Alert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Alert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Alert) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Alert) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRegion

`func (o *Alert) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Alert) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Alert) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Alert) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStreamId

`func (o *Alert) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *Alert) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *Alert) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.


### GetTag

`func (o *Alert) GetTag() AlertTagDetails`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Alert) GetTagOk() (*AlertTagDetails, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Alert) SetTag(v AlertTagDetails)`

SetTag sets Tag field to given value.


### GetThreshold

`func (o *Alert) GetThreshold() AlertThresholdDetails`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Alert) GetThresholdOk() (*AlertThresholdDetails, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Alert) SetThreshold(v AlertThresholdDetails)`

SetThreshold sets Threshold field to given value.


### GetTriggeredTime

`func (o *Alert) GetTriggeredTime() string`

GetTriggeredTime returns the TriggeredTime field if non-nil, zero value otherwise.

### GetTriggeredTimeOk

`func (o *Alert) GetTriggeredTimeOk() (*string, bool)`

GetTriggeredTimeOk returns a tuple with the TriggeredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredTime

`func (o *Alert) SetTriggeredTime(v string)`

SetTriggeredTime sets TriggeredTime field to given value.

### HasTriggeredTime

`func (o *Alert) HasTriggeredTime() bool`

HasTriggeredTime returns a boolean if a field has been set.

### GetType

`func (o *Alert) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Alert) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Alert) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Alert) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeId

`func (o *Alert) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *Alert) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *Alert) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *Alert) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


