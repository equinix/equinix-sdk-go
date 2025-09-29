# DeviceACLTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The ACL template name. | 
**Description** | **string** | The ACL template description. | 
**ProjectId** | Pointer to **string** | Customer project Id. Check your projectId under Resource Management on Equinix Portal. You should have access to a project to see or create assets under it. Equinix will assign a projectId if you do not provide one. | [optional] 
**InboundRules** | [**[]InboundRules**](InboundRules.md) | An array of inbound rules. | 

## Methods

### NewDeviceACLTemplateRequest

`func NewDeviceACLTemplateRequest(name string, description string, inboundRules []InboundRules, ) *DeviceACLTemplateRequest`

NewDeviceACLTemplateRequest instantiates a new DeviceACLTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceACLTemplateRequestWithDefaults

`func NewDeviceACLTemplateRequestWithDefaults() *DeviceACLTemplateRequest`

NewDeviceACLTemplateRequestWithDefaults instantiates a new DeviceACLTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceACLTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceACLTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceACLTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DeviceACLTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceACLTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceACLTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProjectId

`func (o *DeviceACLTemplateRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeviceACLTemplateRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeviceACLTemplateRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DeviceACLTemplateRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetInboundRules

`func (o *DeviceACLTemplateRequest) GetInboundRules() []InboundRules`

GetInboundRules returns the InboundRules field if non-nil, zero value otherwise.

### GetInboundRulesOk

`func (o *DeviceACLTemplateRequest) GetInboundRulesOk() (*[]InboundRules, bool)`

GetInboundRulesOk returns a tuple with the InboundRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundRules

`func (o *DeviceACLTemplateRequest) SetInboundRules(v []InboundRules)`

SetInboundRules sets InboundRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


