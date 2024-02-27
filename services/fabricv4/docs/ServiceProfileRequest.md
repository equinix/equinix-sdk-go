# ServiceProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Href** | Pointer to **string** | Service Profile URI response attribute | [optional] [readonly] 
**Type** | [**ServiceProfileTypeEnum**](ServiceProfileTypeEnum.md) |  | 
**Name** | **string** | Customer-assigned service profile name | 
**Uuid** | Pointer to **string** | Equinix-assigned service profile identifier | [optional] 
**Description** | **string** | User-provided service description should be of maximum length 375 | 
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

### NewServiceProfileRequest

`func NewServiceProfileRequest(type_ ServiceProfileTypeEnum, name string, description string, ) *ServiceProfileRequest`

NewServiceProfileRequest instantiates a new ServiceProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileRequestWithDefaults

`func NewServiceProfileRequestWithDefaults() *ServiceProfileRequest`

NewServiceProfileRequestWithDefaults instantiates a new ServiceProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ServiceProfileRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServiceProfileRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServiceProfileRequest) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *ServiceProfileRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetHref

`func (o *ServiceProfileRequest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ServiceProfileRequest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ServiceProfileRequest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ServiceProfileRequest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *ServiceProfileRequest) GetType() ServiceProfileTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceProfileRequest) GetTypeOk() (*ServiceProfileTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceProfileRequest) SetType(v ServiceProfileTypeEnum)`

SetType sets Type field to given value.


### GetName

`func (o *ServiceProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *ServiceProfileRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceProfileRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceProfileRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceProfileRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNotifications

`func (o *ServiceProfileRequest) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ServiceProfileRequest) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ServiceProfileRequest) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ServiceProfileRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetTags

`func (o *ServiceProfileRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceProfileRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceProfileRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVisibility

`func (o *ServiceProfileRequest) GetVisibility() ServiceProfileVisibilityEnum`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ServiceProfileRequest) GetVisibilityOk() (*ServiceProfileVisibilityEnum, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ServiceProfileRequest) SetVisibility(v ServiceProfileVisibilityEnum)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ServiceProfileRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAllowedEmails

`func (o *ServiceProfileRequest) GetAllowedEmails() []string`

GetAllowedEmails returns the AllowedEmails field if non-nil, zero value otherwise.

### GetAllowedEmailsOk

`func (o *ServiceProfileRequest) GetAllowedEmailsOk() (*[]string, bool)`

GetAllowedEmailsOk returns a tuple with the AllowedEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEmails

`func (o *ServiceProfileRequest) SetAllowedEmails(v []string)`

SetAllowedEmails sets AllowedEmails field to given value.

### HasAllowedEmails

`func (o *ServiceProfileRequest) HasAllowedEmails() bool`

HasAllowedEmails returns a boolean if a field has been set.

### GetAccessPointTypeConfigs

`func (o *ServiceProfileRequest) GetAccessPointTypeConfigs() []ServiceProfileAccessPointType`

GetAccessPointTypeConfigs returns the AccessPointTypeConfigs field if non-nil, zero value otherwise.

### GetAccessPointTypeConfigsOk

`func (o *ServiceProfileRequest) GetAccessPointTypeConfigsOk() (*[]ServiceProfileAccessPointType, bool)`

GetAccessPointTypeConfigsOk returns a tuple with the AccessPointTypeConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPointTypeConfigs

`func (o *ServiceProfileRequest) SetAccessPointTypeConfigs(v []ServiceProfileAccessPointType)`

SetAccessPointTypeConfigs sets AccessPointTypeConfigs field to given value.

### HasAccessPointTypeConfigs

`func (o *ServiceProfileRequest) HasAccessPointTypeConfigs() bool`

HasAccessPointTypeConfigs returns a boolean if a field has been set.

### GetCustomFields

`func (o *ServiceProfileRequest) GetCustomFields() []CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ServiceProfileRequest) GetCustomFieldsOk() (*[]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ServiceProfileRequest) SetCustomFields(v []CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ServiceProfileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetMarketingInfo

`func (o *ServiceProfileRequest) GetMarketingInfo() MarketingInfo`

GetMarketingInfo returns the MarketingInfo field if non-nil, zero value otherwise.

### GetMarketingInfoOk

`func (o *ServiceProfileRequest) GetMarketingInfoOk() (*MarketingInfo, bool)`

GetMarketingInfoOk returns a tuple with the MarketingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingInfo

`func (o *ServiceProfileRequest) SetMarketingInfo(v MarketingInfo)`

SetMarketingInfo sets MarketingInfo field to given value.

### HasMarketingInfo

`func (o *ServiceProfileRequest) HasMarketingInfo() bool`

HasMarketingInfo returns a boolean if a field has been set.

### GetPorts

`func (o *ServiceProfileRequest) GetPorts() []ServiceProfileAccessPointCOLO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ServiceProfileRequest) GetPortsOk() (*[]ServiceProfileAccessPointCOLO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ServiceProfileRequest) SetPorts(v []ServiceProfileAccessPointCOLO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ServiceProfileRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetVirtualDevices

`func (o *ServiceProfileRequest) GetVirtualDevices() []ServiceProfileAccessPointVD`

GetVirtualDevices returns the VirtualDevices field if non-nil, zero value otherwise.

### GetVirtualDevicesOk

`func (o *ServiceProfileRequest) GetVirtualDevicesOk() (*[]ServiceProfileAccessPointVD, bool)`

GetVirtualDevicesOk returns a tuple with the VirtualDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDevices

`func (o *ServiceProfileRequest) SetVirtualDevices(v []ServiceProfileAccessPointVD)`

SetVirtualDevices sets VirtualDevices field to given value.

### HasVirtualDevices

`func (o *ServiceProfileRequest) HasVirtualDevices() bool`

HasVirtualDevices returns a boolean if a field has been set.

### GetMetros

`func (o *ServiceProfileRequest) GetMetros() []ServiceMetro`

GetMetros returns the Metros field if non-nil, zero value otherwise.

### GetMetrosOk

`func (o *ServiceProfileRequest) GetMetrosOk() (*[]ServiceMetro, bool)`

GetMetrosOk returns a tuple with the Metros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetros

`func (o *ServiceProfileRequest) SetMetros(v []ServiceMetro)`

SetMetros sets Metros field to given value.

### HasMetros

`func (o *ServiceProfileRequest) HasMetros() bool`

HasMetros returns a boolean if a field has been set.

### GetSelfProfile

`func (o *ServiceProfileRequest) GetSelfProfile() bool`

GetSelfProfile returns the SelfProfile field if non-nil, zero value otherwise.

### GetSelfProfileOk

`func (o *ServiceProfileRequest) GetSelfProfileOk() (*bool, bool)`

GetSelfProfileOk returns a tuple with the SelfProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfProfile

`func (o *ServiceProfileRequest) SetSelfProfile(v bool)`

SetSelfProfile sets SelfProfile field to given value.

### HasSelfProfile

`func (o *ServiceProfileRequest) HasSelfProfile() bool`

HasSelfProfile returns a boolean if a field has been set.

### GetProjectId

`func (o *ServiceProfileRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ServiceProfileRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ServiceProfileRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ServiceProfileRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


