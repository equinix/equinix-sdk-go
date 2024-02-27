# SimplifiedServiceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Service Profile URI response attribute | [optional] [readonly] 
**Type** | Pointer to [**ServiceProfileTypeEnum**](ServiceProfileTypeEnum.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-assigned service profile name | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned service profile identifier | [optional] 
**Description** | Pointer to **string** | User-provided service description should be of maximum length 375 | [optional] 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Recipients of notifications on service profile change | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Visibility** | Pointer to [**ServiceProfileVisibilityEnum**](ServiceProfileVisibilityEnum.md) |  | [optional] 
**AllowedEmails** | Pointer to **[]string** |  | [optional] 
**AccessPointTypeConfigs** | Pointer to [**[]ServiceProfileAccessPointType**](ServiceProfileAccessPointType.md) |  | [optional] 
**CustomFields** | Pointer to [**[]CustomField**](CustomField.md) |  | [optional] 
**MarketingInfo** | Pointer to [**MarketingInfo**](MarketingInfo.md) |  | [optional] 
**Ports** | Pointer to [**[]ServiceProfileAccessPointCOLO**](ServiceProfileAccessPointCOLO.md) |  | [optional] 
**VirtualDevices** | Pointer to [**[]ServiceProfileAccessPointVD**](ServiceProfileAccessPointVD.md) |  | [optional] 
**Metros** | Pointer to [**[]ServiceMetro**](ServiceMetro.md) | Derived response attribute. | [optional] 
**SelfProfile** | Pointer to **bool** | response attribute indicates whether the profile belongs to the same organization as the api-invoker. | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 

## Methods

### NewSimplifiedServiceProfile

`func NewSimplifiedServiceProfile() *SimplifiedServiceProfile`

NewSimplifiedServiceProfile instantiates a new SimplifiedServiceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedServiceProfileWithDefaults

`func NewSimplifiedServiceProfileWithDefaults() *SimplifiedServiceProfile`

NewSimplifiedServiceProfileWithDefaults instantiates a new SimplifiedServiceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *SimplifiedServiceProfile) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SimplifiedServiceProfile) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SimplifiedServiceProfile) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SimplifiedServiceProfile) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *SimplifiedServiceProfile) GetType() ServiceProfileTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedServiceProfile) GetTypeOk() (*ServiceProfileTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedServiceProfile) SetType(v ServiceProfileTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *SimplifiedServiceProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *SimplifiedServiceProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimplifiedServiceProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimplifiedServiceProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimplifiedServiceProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *SimplifiedServiceProfile) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SimplifiedServiceProfile) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SimplifiedServiceProfile) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SimplifiedServiceProfile) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *SimplifiedServiceProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimplifiedServiceProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimplifiedServiceProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SimplifiedServiceProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotifications

`func (o *SimplifiedServiceProfile) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *SimplifiedServiceProfile) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *SimplifiedServiceProfile) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *SimplifiedServiceProfile) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetTags

`func (o *SimplifiedServiceProfile) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SimplifiedServiceProfile) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SimplifiedServiceProfile) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SimplifiedServiceProfile) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVisibility

`func (o *SimplifiedServiceProfile) GetVisibility() ServiceProfileVisibilityEnum`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SimplifiedServiceProfile) GetVisibilityOk() (*ServiceProfileVisibilityEnum, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SimplifiedServiceProfile) SetVisibility(v ServiceProfileVisibilityEnum)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SimplifiedServiceProfile) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAllowedEmails

`func (o *SimplifiedServiceProfile) GetAllowedEmails() []string`

GetAllowedEmails returns the AllowedEmails field if non-nil, zero value otherwise.

### GetAllowedEmailsOk

`func (o *SimplifiedServiceProfile) GetAllowedEmailsOk() (*[]string, bool)`

GetAllowedEmailsOk returns a tuple with the AllowedEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEmails

`func (o *SimplifiedServiceProfile) SetAllowedEmails(v []string)`

SetAllowedEmails sets AllowedEmails field to given value.

### HasAllowedEmails

`func (o *SimplifiedServiceProfile) HasAllowedEmails() bool`

HasAllowedEmails returns a boolean if a field has been set.

### GetAccessPointTypeConfigs

`func (o *SimplifiedServiceProfile) GetAccessPointTypeConfigs() []ServiceProfileAccessPointType`

GetAccessPointTypeConfigs returns the AccessPointTypeConfigs field if non-nil, zero value otherwise.

### GetAccessPointTypeConfigsOk

`func (o *SimplifiedServiceProfile) GetAccessPointTypeConfigsOk() (*[]ServiceProfileAccessPointType, bool)`

GetAccessPointTypeConfigsOk returns a tuple with the AccessPointTypeConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPointTypeConfigs

`func (o *SimplifiedServiceProfile) SetAccessPointTypeConfigs(v []ServiceProfileAccessPointType)`

SetAccessPointTypeConfigs sets AccessPointTypeConfigs field to given value.

### HasAccessPointTypeConfigs

`func (o *SimplifiedServiceProfile) HasAccessPointTypeConfigs() bool`

HasAccessPointTypeConfigs returns a boolean if a field has been set.

### GetCustomFields

`func (o *SimplifiedServiceProfile) GetCustomFields() []CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SimplifiedServiceProfile) GetCustomFieldsOk() (*[]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SimplifiedServiceProfile) SetCustomFields(v []CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SimplifiedServiceProfile) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetMarketingInfo

`func (o *SimplifiedServiceProfile) GetMarketingInfo() MarketingInfo`

GetMarketingInfo returns the MarketingInfo field if non-nil, zero value otherwise.

### GetMarketingInfoOk

`func (o *SimplifiedServiceProfile) GetMarketingInfoOk() (*MarketingInfo, bool)`

GetMarketingInfoOk returns a tuple with the MarketingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingInfo

`func (o *SimplifiedServiceProfile) SetMarketingInfo(v MarketingInfo)`

SetMarketingInfo sets MarketingInfo field to given value.

### HasMarketingInfo

`func (o *SimplifiedServiceProfile) HasMarketingInfo() bool`

HasMarketingInfo returns a boolean if a field has been set.

### GetPorts

`func (o *SimplifiedServiceProfile) GetPorts() []ServiceProfileAccessPointCOLO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *SimplifiedServiceProfile) GetPortsOk() (*[]ServiceProfileAccessPointCOLO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *SimplifiedServiceProfile) SetPorts(v []ServiceProfileAccessPointCOLO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *SimplifiedServiceProfile) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetVirtualDevices

`func (o *SimplifiedServiceProfile) GetVirtualDevices() []ServiceProfileAccessPointVD`

GetVirtualDevices returns the VirtualDevices field if non-nil, zero value otherwise.

### GetVirtualDevicesOk

`func (o *SimplifiedServiceProfile) GetVirtualDevicesOk() (*[]ServiceProfileAccessPointVD, bool)`

GetVirtualDevicesOk returns a tuple with the VirtualDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDevices

`func (o *SimplifiedServiceProfile) SetVirtualDevices(v []ServiceProfileAccessPointVD)`

SetVirtualDevices sets VirtualDevices field to given value.

### HasVirtualDevices

`func (o *SimplifiedServiceProfile) HasVirtualDevices() bool`

HasVirtualDevices returns a boolean if a field has been set.

### GetMetros

`func (o *SimplifiedServiceProfile) GetMetros() []ServiceMetro`

GetMetros returns the Metros field if non-nil, zero value otherwise.

### GetMetrosOk

`func (o *SimplifiedServiceProfile) GetMetrosOk() (*[]ServiceMetro, bool)`

GetMetrosOk returns a tuple with the Metros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetros

`func (o *SimplifiedServiceProfile) SetMetros(v []ServiceMetro)`

SetMetros sets Metros field to given value.

### HasMetros

`func (o *SimplifiedServiceProfile) HasMetros() bool`

HasMetros returns a boolean if a field has been set.

### GetSelfProfile

`func (o *SimplifiedServiceProfile) GetSelfProfile() bool`

GetSelfProfile returns the SelfProfile field if non-nil, zero value otherwise.

### GetSelfProfileOk

`func (o *SimplifiedServiceProfile) GetSelfProfileOk() (*bool, bool)`

GetSelfProfileOk returns a tuple with the SelfProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfProfile

`func (o *SimplifiedServiceProfile) SetSelfProfile(v bool)`

SetSelfProfile sets SelfProfile field to given value.

### HasSelfProfile

`func (o *SimplifiedServiceProfile) HasSelfProfile() bool`

HasSelfProfile returns a boolean if a field has been set.

### GetProjectId

`func (o *SimplifiedServiceProfile) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SimplifiedServiceProfile) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SimplifiedServiceProfile) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SimplifiedServiceProfile) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


