# TrendingPowerDataStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | [Ok|Permission Denied|Invalid Account number|Invalid IBX|Invalid LevelType|Invalid LevelValue|Invalid Interval|Invalid From/To Date|From/ToDate interval cannot be greater  than 1 year|INVALID_SESSION|INVALID_SESSION_IBX|INTERNAL_ERROR] are the possible messages | [optional] 
**Statuscode** | Pointer to **float32** | [1000|3001|3002|3003|4000] are the possible status codes | [optional] 
**Type** | Pointer to [**AssetDetailResponseStatusType**](AssetDetailResponseStatusType.md) |  | [optional] 

## Methods

### NewTrendingPowerDataStatus

`func NewTrendingPowerDataStatus() *TrendingPowerDataStatus`

NewTrendingPowerDataStatus instantiates a new TrendingPowerDataStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendingPowerDataStatusWithDefaults

`func NewTrendingPowerDataStatusWithDefaults() *TrendingPowerDataStatus`

NewTrendingPowerDataStatusWithDefaults instantiates a new TrendingPowerDataStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *TrendingPowerDataStatus) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *TrendingPowerDataStatus) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *TrendingPowerDataStatus) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *TrendingPowerDataStatus) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatuscode

`func (o *TrendingPowerDataStatus) GetStatuscode() float32`

GetStatuscode returns the Statuscode field if non-nil, zero value otherwise.

### GetStatuscodeOk

`func (o *TrendingPowerDataStatus) GetStatuscodeOk() (*float32, bool)`

GetStatuscodeOk returns a tuple with the Statuscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuscode

`func (o *TrendingPowerDataStatus) SetStatuscode(v float32)`

SetStatuscode sets Statuscode field to given value.

### HasStatuscode

`func (o *TrendingPowerDataStatus) HasStatuscode() bool`

HasStatuscode returns a boolean if a field has been set.

### GetType

`func (o *TrendingPowerDataStatus) GetType() AssetDetailResponseStatusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrendingPowerDataStatus) GetTypeOk() (*AssetDetailResponseStatusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrendingPowerDataStatus) SetType(v AssetDetailResponseStatusType)`

SetType sets Type field to given value.

### HasType

`func (o *TrendingPowerDataStatus) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


