# SupportRequestInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 
**Priority** | Pointer to [**SupportRequestInputPriority**](SupportRequestInputPriority.md) |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Subject** | **string** |  | 

## Methods

### NewSupportRequestInput

`func NewSupportRequestInput(message string, subject string, ) *SupportRequestInput`

NewSupportRequestInput instantiates a new SupportRequestInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportRequestInputWithDefaults

`func NewSupportRequestInputWithDefaults() *SupportRequestInput`

NewSupportRequestInputWithDefaults instantiates a new SupportRequestInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *SupportRequestInput) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SupportRequestInput) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SupportRequestInput) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SupportRequestInput) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetMessage

`func (o *SupportRequestInput) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SupportRequestInput) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SupportRequestInput) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPriority

`func (o *SupportRequestInput) GetPriority() SupportRequestInputPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SupportRequestInput) GetPriorityOk() (*SupportRequestInputPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SupportRequestInput) SetPriority(v SupportRequestInputPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SupportRequestInput) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProjectId

`func (o *SupportRequestInput) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SupportRequestInput) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SupportRequestInput) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SupportRequestInput) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSubject

`func (o *SupportRequestInput) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SupportRequestInput) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SupportRequestInput) SetSubject(v string)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


