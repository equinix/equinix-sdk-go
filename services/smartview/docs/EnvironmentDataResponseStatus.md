# EnvironmentDataResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | [Ok|Invalid Level Type|Invalid Level Value|Invalid Account Details Or User is not Entitled|User Preference is not set as expected in the account|Invalid IBX|Invalid Account number|User is not entitled to the ibx|INVALID_SESSION|INVALID_SESSION_IBX|INTERNAL_ERROR|NO_DATA|UNAUTHORISZED_ACCESS] are the possible messages         | [optional] 
**Statuscode** | Pointer to **float32** | [1000|3001|3002|3003|3004|3005|4000|5000|6001] are the possible status codes | [optional] 
**Type** | Pointer to [**AssetDetailResponseStatusType**](AssetDetailResponseStatusType.md) |  | [optional] 

## Methods

### NewEnvironmentDataResponseStatus

`func NewEnvironmentDataResponseStatus() *EnvironmentDataResponseStatus`

NewEnvironmentDataResponseStatus instantiates a new EnvironmentDataResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDataResponseStatusWithDefaults

`func NewEnvironmentDataResponseStatusWithDefaults() *EnvironmentDataResponseStatus`

NewEnvironmentDataResponseStatusWithDefaults instantiates a new EnvironmentDataResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *EnvironmentDataResponseStatus) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *EnvironmentDataResponseStatus) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *EnvironmentDataResponseStatus) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *EnvironmentDataResponseStatus) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatuscode

`func (o *EnvironmentDataResponseStatus) GetStatuscode() float32`

GetStatuscode returns the Statuscode field if non-nil, zero value otherwise.

### GetStatuscodeOk

`func (o *EnvironmentDataResponseStatus) GetStatuscodeOk() (*float32, bool)`

GetStatuscodeOk returns a tuple with the Statuscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuscode

`func (o *EnvironmentDataResponseStatus) SetStatuscode(v float32)`

SetStatuscode sets Statuscode field to given value.

### HasStatuscode

`func (o *EnvironmentDataResponseStatus) HasStatuscode() bool`

HasStatuscode returns a boolean if a field has been set.

### GetType

`func (o *EnvironmentDataResponseStatus) GetType() AssetDetailResponseStatusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentDataResponseStatus) GetTypeOk() (*AssetDetailResponseStatusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentDataResponseStatus) SetType(v AssetDetailResponseStatusType)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentDataResponseStatus) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


