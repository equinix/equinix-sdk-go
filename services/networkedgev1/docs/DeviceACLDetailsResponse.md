# DeviceACLDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the ACL template. | [optional] 
**Uuid** | Pointer to **string** | The unique Id of the ACL template. | [optional] 
**Description** | Pointer to **string** | The description of the ACL template. | [optional] 
**InboundRules** | Pointer to [**[]InboundRules**](InboundRules.md) | An array of inbound rules | [optional] 
**CreatedBy** | Pointer to **string** | Created by | [optional] 
**CreatedDate** | Pointer to **string** | Created date | [optional] 
**Status** | Pointer to **string** | The status of the ACL template on the device. Possible values are PROVISIONED, DEPROVISIONED, DEVICE_NOT_READY, FAILED, NOT_APPLIED, PROVISIONING. | [optional] 

## Methods

### NewDeviceACLDetailsResponse

`func NewDeviceACLDetailsResponse() *DeviceACLDetailsResponse`

NewDeviceACLDetailsResponse instantiates a new DeviceACLDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceACLDetailsResponseWithDefaults

`func NewDeviceACLDetailsResponseWithDefaults() *DeviceACLDetailsResponse`

NewDeviceACLDetailsResponseWithDefaults instantiates a new DeviceACLDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceACLDetailsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceACLDetailsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceACLDetailsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceACLDetailsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *DeviceACLDetailsResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeviceACLDetailsResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeviceACLDetailsResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeviceACLDetailsResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceACLDetailsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceACLDetailsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceACLDetailsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceACLDetailsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInboundRules

`func (o *DeviceACLDetailsResponse) GetInboundRules() []InboundRules`

GetInboundRules returns the InboundRules field if non-nil, zero value otherwise.

### GetInboundRulesOk

`func (o *DeviceACLDetailsResponse) GetInboundRulesOk() (*[]InboundRules, bool)`

GetInboundRulesOk returns a tuple with the InboundRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundRules

`func (o *DeviceACLDetailsResponse) SetInboundRules(v []InboundRules)`

SetInboundRules sets InboundRules field to given value.

### HasInboundRules

`func (o *DeviceACLDetailsResponse) HasInboundRules() bool`

HasInboundRules returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DeviceACLDetailsResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DeviceACLDetailsResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DeviceACLDetailsResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DeviceACLDetailsResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *DeviceACLDetailsResponse) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *DeviceACLDetailsResponse) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *DeviceACLDetailsResponse) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *DeviceACLDetailsResponse) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceACLDetailsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceACLDetailsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceACLDetailsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceACLDetailsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


