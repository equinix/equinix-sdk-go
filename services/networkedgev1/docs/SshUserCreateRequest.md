# SshUserCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | At least 3 and upto a maximum of 32 alphanumeric characters. The only special characters allowed are - _ | 
**Password** | **string** | At least 6 and upto a maximum of 12 alphanumeric characters. The only special characters allowed are - _ $ @ | 
**DeviceUuid** | **string** | The unique Id of a virtual device. | 

## Methods

### NewSshUserCreateRequest

`func NewSshUserCreateRequest(username string, password string, deviceUuid string, ) *SshUserCreateRequest`

NewSshUserCreateRequest instantiates a new SshUserCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserCreateRequestWithDefaults

`func NewSshUserCreateRequestWithDefaults() *SshUserCreateRequest`

NewSshUserCreateRequestWithDefaults instantiates a new SshUserCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *SshUserCreateRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SshUserCreateRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SshUserCreateRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *SshUserCreateRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SshUserCreateRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SshUserCreateRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDeviceUuid

`func (o *SshUserCreateRequest) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *SshUserCreateRequest) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *SshUserCreateRequest) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


