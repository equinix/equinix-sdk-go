# DeviceACLTemplatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the ACL template. | [optional] 
**Uuid** | Pointer to **string** | The unique Id of the ACL template. | [optional] 
**Description** | Pointer to **string** | The description of the ACL template. | [optional] 
**InboundRules** | Pointer to [**[]InboundRules**](InboundRules.md) | An array of inbound rules | [optional] 
**CreatedBy** | Pointer to **string** | Created by | [optional] 
**CreatedDate** | Pointer to **string** | Created date | [optional] 

## Methods

### NewDeviceACLTemplatesResponse

`func NewDeviceACLTemplatesResponse() *DeviceACLTemplatesResponse`

NewDeviceACLTemplatesResponse instantiates a new DeviceACLTemplatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceACLTemplatesResponseWithDefaults

`func NewDeviceACLTemplatesResponseWithDefaults() *DeviceACLTemplatesResponse`

NewDeviceACLTemplatesResponseWithDefaults instantiates a new DeviceACLTemplatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceACLTemplatesResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceACLTemplatesResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceACLTemplatesResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceACLTemplatesResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *DeviceACLTemplatesResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeviceACLTemplatesResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeviceACLTemplatesResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeviceACLTemplatesResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceACLTemplatesResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceACLTemplatesResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceACLTemplatesResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceACLTemplatesResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInboundRules

`func (o *DeviceACLTemplatesResponse) GetInboundRules() []InboundRules`

GetInboundRules returns the InboundRules field if non-nil, zero value otherwise.

### GetInboundRulesOk

`func (o *DeviceACLTemplatesResponse) GetInboundRulesOk() (*[]InboundRules, bool)`

GetInboundRulesOk returns a tuple with the InboundRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundRules

`func (o *DeviceACLTemplatesResponse) SetInboundRules(v []InboundRules)`

SetInboundRules sets InboundRules field to given value.

### HasInboundRules

`func (o *DeviceACLTemplatesResponse) HasInboundRules() bool`

HasInboundRules returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DeviceACLTemplatesResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DeviceACLTemplatesResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DeviceACLTemplatesResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DeviceACLTemplatesResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *DeviceACLTemplatesResponse) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *DeviceACLTemplatesResponse) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *DeviceACLTemplatesResponse) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *DeviceACLTemplatesResponse) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


