# AssetsListStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | [Ok|Permission Denied|Invalid Account number|Invalid IBX|Invalid LevelType|Invalid LevelValue|Invalid Interval|Invalid From/To Date|INVALID_SESSION|INVALID_SESSION_IBX|INTERNAL_ERROR] are the possible messages            possible messages | [optional] 
**Statuscode** | Pointer to **float32** | [1000|3001|3002|3003|4000] are the possible status codes | [optional] 
**Type** | Pointer to [**AssetDetailResponseStatusType**](AssetDetailResponseStatusType.md) |  | [optional] 

## Methods

### NewAssetsListStatus

`func NewAssetsListStatus() *AssetsListStatus`

NewAssetsListStatus instantiates a new AssetsListStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsListStatusWithDefaults

`func NewAssetsListStatusWithDefaults() *AssetsListStatus`

NewAssetsListStatusWithDefaults instantiates a new AssetsListStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *AssetsListStatus) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AssetsListStatus) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AssetsListStatus) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AssetsListStatus) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatuscode

`func (o *AssetsListStatus) GetStatuscode() float32`

GetStatuscode returns the Statuscode field if non-nil, zero value otherwise.

### GetStatuscodeOk

`func (o *AssetsListStatus) GetStatuscodeOk() (*float32, bool)`

GetStatuscodeOk returns a tuple with the Statuscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuscode

`func (o *AssetsListStatus) SetStatuscode(v float32)`

SetStatuscode sets Statuscode field to given value.

### HasStatuscode

`func (o *AssetsListStatus) HasStatuscode() bool`

HasStatuscode returns a boolean if a field has been set.

### GetType

`func (o *AssetsListStatus) GetType() AssetDetailResponseStatusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssetsListStatus) GetTypeOk() (*AssetDetailResponseStatusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssetsListStatus) SetType(v AssetDetailResponseStatusType)`

SetType sets Type field to given value.

### HasType

`func (o *AssetsListStatus) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


