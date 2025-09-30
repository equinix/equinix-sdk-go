# ACLTemplateDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the ACL template. | [optional] 
**Uuid** | Pointer to **string** | The unique Id of the ACL template. | [optional] 
**Description** | Pointer to **string** | The description of the ACL template. | [optional] 
**InboundRules** | Pointer to [**[]InboundRules**](InboundRules.md) | An array of inbound rules | [optional] 
**VirtualDeviceDetails** | Pointer to [**[]VirtualDeviceACLDetails**](VirtualDeviceACLDetails.md) | The array of devices associated with this ACL template | [optional] 
**CreatedBy** | Pointer to **string** | Created by | [optional] 
**CreatedDate** | Pointer to **string** | Created date | [optional] 

## Methods

### NewACLTemplateDetailsResponse

`func NewACLTemplateDetailsResponse() *ACLTemplateDetailsResponse`

NewACLTemplateDetailsResponse instantiates a new ACLTemplateDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACLTemplateDetailsResponseWithDefaults

`func NewACLTemplateDetailsResponseWithDefaults() *ACLTemplateDetailsResponse`

NewACLTemplateDetailsResponseWithDefaults instantiates a new ACLTemplateDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ACLTemplateDetailsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ACLTemplateDetailsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ACLTemplateDetailsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ACLTemplateDetailsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *ACLTemplateDetailsResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ACLTemplateDetailsResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ACLTemplateDetailsResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ACLTemplateDetailsResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *ACLTemplateDetailsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ACLTemplateDetailsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ACLTemplateDetailsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ACLTemplateDetailsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInboundRules

`func (o *ACLTemplateDetailsResponse) GetInboundRules() []InboundRules`

GetInboundRules returns the InboundRules field if non-nil, zero value otherwise.

### GetInboundRulesOk

`func (o *ACLTemplateDetailsResponse) GetInboundRulesOk() (*[]InboundRules, bool)`

GetInboundRulesOk returns a tuple with the InboundRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundRules

`func (o *ACLTemplateDetailsResponse) SetInboundRules(v []InboundRules)`

SetInboundRules sets InboundRules field to given value.

### HasInboundRules

`func (o *ACLTemplateDetailsResponse) HasInboundRules() bool`

HasInboundRules returns a boolean if a field has been set.

### GetVirtualDeviceDetails

`func (o *ACLTemplateDetailsResponse) GetVirtualDeviceDetails() []VirtualDeviceACLDetails`

GetVirtualDeviceDetails returns the VirtualDeviceDetails field if non-nil, zero value otherwise.

### GetVirtualDeviceDetailsOk

`func (o *ACLTemplateDetailsResponse) GetVirtualDeviceDetailsOk() (*[]VirtualDeviceACLDetails, bool)`

GetVirtualDeviceDetailsOk returns a tuple with the VirtualDeviceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceDetails

`func (o *ACLTemplateDetailsResponse) SetVirtualDeviceDetails(v []VirtualDeviceACLDetails)`

SetVirtualDeviceDetails sets VirtualDeviceDetails field to given value.

### HasVirtualDeviceDetails

`func (o *ACLTemplateDetailsResponse) HasVirtualDeviceDetails() bool`

HasVirtualDeviceDetails returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ACLTemplateDetailsResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ACLTemplateDetailsResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ACLTemplateDetailsResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ACLTemplateDetailsResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ACLTemplateDetailsResponse) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ACLTemplateDetailsResponse) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ACLTemplateDetailsResponse) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ACLTemplateDetailsResponse) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


