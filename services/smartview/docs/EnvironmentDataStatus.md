# EnvironmentDataStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | [Ok|Invalid Level Type|Invalid Level Value|Invalid Account Details Or User is not Entitled|User Preference is not set as expected in the account|Invalid IBX|Invalid Account number|User is not entitled to the ibx|INVALID_SESSION|INVALID_SESSION_IBX|INTERNAL_ERROR|NO_DATA|UNAUTHORISZED_ACCESS] are the possible messages | [optional] 
**Statuscode** | Pointer to **float32** | [1000|3001|3002|3003|3004|3005|4000|5000|6001] are the possible status codes | [optional] 
**Type** | Pointer to [**AssetDetailResponseStatusType**](AssetDetailResponseStatusType.md) |  | [optional] 

## Methods

### NewEnvironmentDataStatus

`func NewEnvironmentDataStatus() *EnvironmentDataStatus`

NewEnvironmentDataStatus instantiates a new EnvironmentDataStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDataStatusWithDefaults

`func NewEnvironmentDataStatusWithDefaults() *EnvironmentDataStatus`

NewEnvironmentDataStatusWithDefaults instantiates a new EnvironmentDataStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *EnvironmentDataStatus) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *EnvironmentDataStatus) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *EnvironmentDataStatus) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *EnvironmentDataStatus) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatuscode

`func (o *EnvironmentDataStatus) GetStatuscode() float32`

GetStatuscode returns the Statuscode field if non-nil, zero value otherwise.

### GetStatuscodeOk

`func (o *EnvironmentDataStatus) GetStatuscodeOk() (*float32, bool)`

GetStatuscodeOk returns a tuple with the Statuscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuscode

`func (o *EnvironmentDataStatus) SetStatuscode(v float32)`

SetStatuscode sets Statuscode field to given value.

### HasStatuscode

`func (o *EnvironmentDataStatus) HasStatuscode() bool`

HasStatuscode returns a boolean if a field has been set.

### GetType

`func (o *EnvironmentDataStatus) GetType() AssetDetailResponseStatusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentDataStatus) GetTypeOk() (*AssetDetailResponseStatusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentDataStatus) SetType(v AssetDetailResponseStatusType)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentDataStatus) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


