# SshUserInfoDissociateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshUserDeleted** | Pointer to **bool** | true &#x3D; the ssh user has been deleted since there are no more devices associated with this user; false &#x3D; the ssh user has not been deleted since associations with devices exist. | [optional] 
**SshUserToDeviceAssociationDeleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewSshUserInfoDissociateResponse

`func NewSshUserInfoDissociateResponse() *SshUserInfoDissociateResponse`

NewSshUserInfoDissociateResponse instantiates a new SshUserInfoDissociateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserInfoDissociateResponseWithDefaults

`func NewSshUserInfoDissociateResponseWithDefaults() *SshUserInfoDissociateResponse`

NewSshUserInfoDissociateResponseWithDefaults instantiates a new SshUserInfoDissociateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshUserDeleted

`func (o *SshUserInfoDissociateResponse) GetSshUserDeleted() bool`

GetSshUserDeleted returns the SshUserDeleted field if non-nil, zero value otherwise.

### GetSshUserDeletedOk

`func (o *SshUserInfoDissociateResponse) GetSshUserDeletedOk() (*bool, bool)`

GetSshUserDeletedOk returns a tuple with the SshUserDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUserDeleted

`func (o *SshUserInfoDissociateResponse) SetSshUserDeleted(v bool)`

SetSshUserDeleted sets SshUserDeleted field to given value.

### HasSshUserDeleted

`func (o *SshUserInfoDissociateResponse) HasSshUserDeleted() bool`

HasSshUserDeleted returns a boolean if a field has been set.

### GetSshUserToDeviceAssociationDeleted

`func (o *SshUserInfoDissociateResponse) GetSshUserToDeviceAssociationDeleted() bool`

GetSshUserToDeviceAssociationDeleted returns the SshUserToDeviceAssociationDeleted field if non-nil, zero value otherwise.

### GetSshUserToDeviceAssociationDeletedOk

`func (o *SshUserInfoDissociateResponse) GetSshUserToDeviceAssociationDeletedOk() (*bool, bool)`

GetSshUserToDeviceAssociationDeletedOk returns a tuple with the SshUserToDeviceAssociationDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUserToDeviceAssociationDeleted

`func (o *SshUserInfoDissociateResponse) SetSshUserToDeviceAssociationDeleted(v bool)`

SetSshUserToDeviceAssociationDeleted sets SshUserToDeviceAssociationDeleted field to given value.

### HasSshUserToDeviceAssociationDeleted

`func (o *SshUserInfoDissociateResponse) HasSshUserToDeviceAssociationDeleted() bool`

HasSshUserToDeviceAssociationDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


