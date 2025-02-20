# StatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **float32** | [1000|3001|3002|3003|4000] are the possible status codes | [optional] 
**Msg** | Pointer to **string** | [Ok|Permission Denied|Invalid Account number|Invalid IBX|Invalid LevelType|Invalid LevelValue|Invalid Interval|Invalid From/To Date|INVALID_SESSION|INVALID_SESSION_IBX|INTERNAL_ERROR] are the possible messages | [optional] 

## Methods

### NewStatusInfo

`func NewStatusInfo() *StatusInfo`

NewStatusInfo instantiates a new StatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusInfoWithDefaults

`func NewStatusInfoWithDefaults() *StatusInfo`

NewStatusInfoWithDefaults instantiates a new StatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *StatusInfo) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StatusInfo) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StatusInfo) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *StatusInfo) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *StatusInfo) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *StatusInfo) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *StatusInfo) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *StatusInfo) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


