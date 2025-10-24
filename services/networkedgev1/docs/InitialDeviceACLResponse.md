# InitialDeviceACLResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclTemplate** | Pointer to [**DeviceACLDetailsResponse**](DeviceACLDetailsResponse.md) |  | [optional] 
**MgmtAclTemplate** | Pointer to [**DeviceACLDetailsResponse**](DeviceACLDetailsResponse.md) |  | [optional] 

## Methods

### NewInitialDeviceACLResponse

`func NewInitialDeviceACLResponse() *InitialDeviceACLResponse`

NewInitialDeviceACLResponse instantiates a new InitialDeviceACLResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitialDeviceACLResponseWithDefaults

`func NewInitialDeviceACLResponseWithDefaults() *InitialDeviceACLResponse`

NewInitialDeviceACLResponseWithDefaults instantiates a new InitialDeviceACLResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclTemplate

`func (o *InitialDeviceACLResponse) GetAclTemplate() DeviceACLDetailsResponse`

GetAclTemplate returns the AclTemplate field if non-nil, zero value otherwise.

### GetAclTemplateOk

`func (o *InitialDeviceACLResponse) GetAclTemplateOk() (*DeviceACLDetailsResponse, bool)`

GetAclTemplateOk returns a tuple with the AclTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclTemplate

`func (o *InitialDeviceACLResponse) SetAclTemplate(v DeviceACLDetailsResponse)`

SetAclTemplate sets AclTemplate field to given value.

### HasAclTemplate

`func (o *InitialDeviceACLResponse) HasAclTemplate() bool`

HasAclTemplate returns a boolean if a field has been set.

### GetMgmtAclTemplate

`func (o *InitialDeviceACLResponse) GetMgmtAclTemplate() DeviceACLDetailsResponse`

GetMgmtAclTemplate returns the MgmtAclTemplate field if non-nil, zero value otherwise.

### GetMgmtAclTemplateOk

`func (o *InitialDeviceACLResponse) GetMgmtAclTemplateOk() (*DeviceACLDetailsResponse, bool)`

GetMgmtAclTemplateOk returns a tuple with the MgmtAclTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtAclTemplate

`func (o *InitialDeviceACLResponse) SetMgmtAclTemplate(v DeviceACLDetailsResponse)`

SetMgmtAclTemplate sets MgmtAclTemplate field to given value.

### HasMgmtAclTemplate

`func (o *InitialDeviceACLResponse) HasMgmtAclTemplate() bool`

HasMgmtAclTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


