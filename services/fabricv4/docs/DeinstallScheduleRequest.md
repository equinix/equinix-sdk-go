# DeinstallScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DeinstallScheduleRequestType**](DeinstallScheduleRequestType.md) |  | [optional] 
**Code** | Pointer to [**DeinstallScheduleRequestCode**](DeinstallScheduleRequestCode.md) |  | [optional] [default to DEINSTALLSCHEDULEREQUESTCODE_STANDARD]
**RequestedDateTime** | **time.Time** | Requested Removal Date and Time. | 
**DeinstallationDateTime** | Pointer to **time.Time** | Deinstallation Date and Time. | [optional] 

## Methods

### NewDeinstallScheduleRequest

`func NewDeinstallScheduleRequest(requestedDateTime time.Time, ) *DeinstallScheduleRequest`

NewDeinstallScheduleRequest instantiates a new DeinstallScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeinstallScheduleRequestWithDefaults

`func NewDeinstallScheduleRequestWithDefaults() *DeinstallScheduleRequest`

NewDeinstallScheduleRequestWithDefaults instantiates a new DeinstallScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeinstallScheduleRequest) GetType() DeinstallScheduleRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeinstallScheduleRequest) GetTypeOk() (*DeinstallScheduleRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeinstallScheduleRequest) SetType(v DeinstallScheduleRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *DeinstallScheduleRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *DeinstallScheduleRequest) GetCode() DeinstallScheduleRequestCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DeinstallScheduleRequest) GetCodeOk() (*DeinstallScheduleRequestCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DeinstallScheduleRequest) SetCode(v DeinstallScheduleRequestCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *DeinstallScheduleRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRequestedDateTime

`func (o *DeinstallScheduleRequest) GetRequestedDateTime() time.Time`

GetRequestedDateTime returns the RequestedDateTime field if non-nil, zero value otherwise.

### GetRequestedDateTimeOk

`func (o *DeinstallScheduleRequest) GetRequestedDateTimeOk() (*time.Time, bool)`

GetRequestedDateTimeOk returns a tuple with the RequestedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDateTime

`func (o *DeinstallScheduleRequest) SetRequestedDateTime(v time.Time)`

SetRequestedDateTime sets RequestedDateTime field to given value.


### GetDeinstallationDateTime

`func (o *DeinstallScheduleRequest) GetDeinstallationDateTime() time.Time`

GetDeinstallationDateTime returns the DeinstallationDateTime field if non-nil, zero value otherwise.

### GetDeinstallationDateTimeOk

`func (o *DeinstallScheduleRequest) GetDeinstallationDateTimeOk() (*time.Time, bool)`

GetDeinstallationDateTimeOk returns a tuple with the DeinstallationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeinstallationDateTime

`func (o *DeinstallScheduleRequest) SetDeinstallationDateTime(v time.Time)`

SetDeinstallationDateTime sets DeinstallationDateTime field to given value.

### HasDeinstallationDateTime

`func (o *DeinstallScheduleRequest) HasDeinstallationDateTime() bool`

HasDeinstallationDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


