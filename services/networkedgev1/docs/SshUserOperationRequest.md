# SshUserOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshUserUuid** | Pointer to **string** | Required for DELETE operation. | [optional] 
**Action** | [**SshUserOperationRequestAction**](SshUserOperationRequestAction.md) |  | 
**SshUsername** | Pointer to **string** | SSH User name | [optional] 
**SshPassword** | Pointer to **string** | SSH Password | [optional] 

## Methods

### NewSshUserOperationRequest

`func NewSshUserOperationRequest(action SshUserOperationRequestAction, ) *SshUserOperationRequest`

NewSshUserOperationRequest instantiates a new SshUserOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserOperationRequestWithDefaults

`func NewSshUserOperationRequestWithDefaults() *SshUserOperationRequest`

NewSshUserOperationRequestWithDefaults instantiates a new SshUserOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshUserUuid

`func (o *SshUserOperationRequest) GetSshUserUuid() string`

GetSshUserUuid returns the SshUserUuid field if non-nil, zero value otherwise.

### GetSshUserUuidOk

`func (o *SshUserOperationRequest) GetSshUserUuidOk() (*string, bool)`

GetSshUserUuidOk returns a tuple with the SshUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUserUuid

`func (o *SshUserOperationRequest) SetSshUserUuid(v string)`

SetSshUserUuid sets SshUserUuid field to given value.

### HasSshUserUuid

`func (o *SshUserOperationRequest) HasSshUserUuid() bool`

HasSshUserUuid returns a boolean if a field has been set.

### GetAction

`func (o *SshUserOperationRequest) GetAction() SshUserOperationRequestAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SshUserOperationRequest) GetActionOk() (*SshUserOperationRequestAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SshUserOperationRequest) SetAction(v SshUserOperationRequestAction)`

SetAction sets Action field to given value.


### GetSshUsername

`func (o *SshUserOperationRequest) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *SshUserOperationRequest) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *SshUserOperationRequest) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *SshUserOperationRequest) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *SshUserOperationRequest) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *SshUserOperationRequest) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *SshUserOperationRequest) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *SshUserOperationRequest) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


