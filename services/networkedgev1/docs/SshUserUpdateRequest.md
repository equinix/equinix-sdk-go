# SshUserUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | At least 6 and upto a maximum of 12 alphanumeric characters. The only special characters allowed are - _ $ @ | 

## Methods

### NewSshUserUpdateRequest

`func NewSshUserUpdateRequest(password string, ) *SshUserUpdateRequest`

NewSshUserUpdateRequest instantiates a new SshUserUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserUpdateRequestWithDefaults

`func NewSshUserUpdateRequestWithDefaults() *SshUserUpdateRequest`

NewSshUserUpdateRequestWithDefaults instantiates a new SshUserUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *SshUserUpdateRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SshUserUpdateRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SshUserUpdateRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


