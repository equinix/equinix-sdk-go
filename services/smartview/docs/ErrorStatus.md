# ErrorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | [Ok|Invalid Level Type|Invalid Level Value|Invalid DataPoint|Invalid Interval|Invalid FromDate|Invalid ToDate|ToDate interval cannot be greater  than 1 year|Invalid Account Details Or User is not Entitled|User Preference is not set as expected in the account|Invalid IBX|Invalid Account number|User is not entitled to the ibx|INVALID_SESSION|INVALID_SESSION_IBX|INTERNAL_ERROR|NO_DATA|UNAUTHORISZED_ACCESS] are the possible messages | [optional] 
**Statuscode** | Pointer to **float32** | [1000|3001|3002|3003|3004|3005|4000|5000|6001] are the possible status codes | [optional] 
**Type** | Pointer to [**ErrorStatusType**](ErrorStatusType.md) |  | [optional] 

## Methods

### NewErrorStatus

`func NewErrorStatus() *ErrorStatus`

NewErrorStatus instantiates a new ErrorStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorStatusWithDefaults

`func NewErrorStatusWithDefaults() *ErrorStatus`

NewErrorStatusWithDefaults instantiates a new ErrorStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *ErrorStatus) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ErrorStatus) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ErrorStatus) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ErrorStatus) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatuscode

`func (o *ErrorStatus) GetStatuscode() float32`

GetStatuscode returns the Statuscode field if non-nil, zero value otherwise.

### GetStatuscodeOk

`func (o *ErrorStatus) GetStatuscodeOk() (*float32, bool)`

GetStatuscodeOk returns a tuple with the Statuscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuscode

`func (o *ErrorStatus) SetStatuscode(v float32)`

SetStatuscode sets Statuscode field to given value.

### HasStatuscode

`func (o *ErrorStatus) HasStatuscode() bool`

HasStatuscode returns a boolean if a field has been set.

### GetType

`func (o *ErrorStatus) GetType() ErrorStatusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorStatus) GetTypeOk() (*ErrorStatusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorStatus) SetType(v ErrorStatusType)`

SetType sets Type field to given value.

### HasType

`func (o *ErrorStatus) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


