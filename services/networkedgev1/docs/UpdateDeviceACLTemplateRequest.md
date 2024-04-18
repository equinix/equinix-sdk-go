# UpdateDeviceACLTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclDetails** | [**[]ACLDetails**](ACLDetails.md) | An array of ACLs. | 

## Methods

### NewUpdateDeviceACLTemplateRequest

`func NewUpdateDeviceACLTemplateRequest(aclDetails []ACLDetails, ) *UpdateDeviceACLTemplateRequest`

NewUpdateDeviceACLTemplateRequest instantiates a new UpdateDeviceACLTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceACLTemplateRequestWithDefaults

`func NewUpdateDeviceACLTemplateRequestWithDefaults() *UpdateDeviceACLTemplateRequest`

NewUpdateDeviceACLTemplateRequestWithDefaults instantiates a new UpdateDeviceACLTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclDetails

`func (o *UpdateDeviceACLTemplateRequest) GetAclDetails() []ACLDetails`

GetAclDetails returns the AclDetails field if non-nil, zero value otherwise.

### GetAclDetailsOk

`func (o *UpdateDeviceACLTemplateRequest) GetAclDetailsOk() (*[]ACLDetails, bool)`

GetAclDetailsOk returns a tuple with the AclDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclDetails

`func (o *UpdateDeviceACLTemplateRequest) SetAclDetails(v []ACLDetails)`

SetAclDetails sets AclDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


