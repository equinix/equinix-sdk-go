# ServiceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**ServiceProfileStateEnum**](ServiceProfileStateEnum.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) | Seller Account for Service Profile. | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) | Seller Account for Service Profile. | [optional] 
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

### NewServiceProfile

`func NewServiceProfile() *ServiceProfile`

NewServiceProfile instantiates a new ServiceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileWithDefaults

`func NewServiceProfileWithDefaults() *ServiceProfile`

NewServiceProfileWithDefaults instantiates a new ServiceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ServiceProfile) GetState() ServiceProfileStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServiceProfile) GetStateOk() (*ServiceProfileStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServiceProfile) SetState(v ServiceProfileStateEnum)`

SetState sets State field to given value.

### HasState

`func (o *ServiceProfile) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAccount

`func (o *ServiceProfile) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ServiceProfile) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ServiceProfile) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ServiceProfile) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetProject

`func (o *ServiceProfile) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServiceProfile) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServiceProfile) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *ServiceProfile) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetChangeLog

`func (o *ServiceProfile) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *ServiceProfile) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *ServiceProfile) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *ServiceProfile) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetHref

`func (o *ServiceProfile) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ServiceProfile) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ServiceProfile) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ServiceProfile) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *ServiceProfile) GetType() ServiceProfileTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceProfile) GetTypeOk() (*ServiceProfileTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceProfile) SetType(v ServiceProfileTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ServiceProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *ServiceProfile) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceProfile) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceProfile) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceProfile) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotifications

`func (o *ServiceProfile) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ServiceProfile) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ServiceProfile) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ServiceProfile) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetTags

`func (o *ServiceProfile) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceProfile) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceProfile) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceProfile) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVisibility

`func (o *ServiceProfile) GetVisibility() ServiceProfileVisibilityEnum`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ServiceProfile) GetVisibilityOk() (*ServiceProfileVisibilityEnum, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ServiceProfile) SetVisibility(v ServiceProfileVisibilityEnum)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ServiceProfile) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAllowedEmails

`func (o *ServiceProfile) GetAllowedEmails() []string`

GetAllowedEmails returns the AllowedEmails field if non-nil, zero value otherwise.

### GetAllowedEmailsOk

`func (o *ServiceProfile) GetAllowedEmailsOk() (*[]string, bool)`

GetAllowedEmailsOk returns a tuple with the AllowedEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEmails

`func (o *ServiceProfile) SetAllowedEmails(v []string)`

SetAllowedEmails sets AllowedEmails field to given value.

### HasAllowedEmails

`func (o *ServiceProfile) HasAllowedEmails() bool`

HasAllowedEmails returns a boolean if a field has been set.

### GetAccessPointTypeConfigs

`func (o *ServiceProfile) GetAccessPointTypeConfigs() []ServiceProfileAccessPointType`

GetAccessPointTypeConfigs returns the AccessPointTypeConfigs field if non-nil, zero value otherwise.

### GetAccessPointTypeConfigsOk

`func (o *ServiceProfile) GetAccessPointTypeConfigsOk() (*[]ServiceProfileAccessPointType, bool)`

GetAccessPointTypeConfigsOk returns a tuple with the AccessPointTypeConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPointTypeConfigs

`func (o *ServiceProfile) SetAccessPointTypeConfigs(v []ServiceProfileAccessPointType)`

SetAccessPointTypeConfigs sets AccessPointTypeConfigs field to given value.

### HasAccessPointTypeConfigs

`func (o *ServiceProfile) HasAccessPointTypeConfigs() bool`

HasAccessPointTypeConfigs returns a boolean if a field has been set.

### GetCustomFields

`func (o *ServiceProfile) GetCustomFields() []CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ServiceProfile) GetCustomFieldsOk() (*[]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ServiceProfile) SetCustomFields(v []CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ServiceProfile) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetMarketingInfo

`func (o *ServiceProfile) GetMarketingInfo() MarketingInfo`

GetMarketingInfo returns the MarketingInfo field if non-nil, zero value otherwise.

### GetMarketingInfoOk

`func (o *ServiceProfile) GetMarketingInfoOk() (*MarketingInfo, bool)`

GetMarketingInfoOk returns a tuple with the MarketingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingInfo

`func (o *ServiceProfile) SetMarketingInfo(v MarketingInfo)`

SetMarketingInfo sets MarketingInfo field to given value.

### HasMarketingInfo

`func (o *ServiceProfile) HasMarketingInfo() bool`

HasMarketingInfo returns a boolean if a field has been set.

### GetPorts

`func (o *ServiceProfile) GetPorts() []ServiceProfileAccessPointCOLO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ServiceProfile) GetPortsOk() (*[]ServiceProfileAccessPointCOLO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ServiceProfile) SetPorts(v []ServiceProfileAccessPointCOLO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ServiceProfile) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetVirtualDevices

`func (o *ServiceProfile) GetVirtualDevices() []ServiceProfileAccessPointVD`

GetVirtualDevices returns the VirtualDevices field if non-nil, zero value otherwise.

### GetVirtualDevicesOk

`func (o *ServiceProfile) GetVirtualDevicesOk() (*[]ServiceProfileAccessPointVD, bool)`

GetVirtualDevicesOk returns a tuple with the VirtualDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDevices

`func (o *ServiceProfile) SetVirtualDevices(v []ServiceProfileAccessPointVD)`

SetVirtualDevices sets VirtualDevices field to given value.

### HasVirtualDevices

`func (o *ServiceProfile) HasVirtualDevices() bool`

HasVirtualDevices returns a boolean if a field has been set.

### GetMetros

`func (o *ServiceProfile) GetMetros() []ServiceMetro`

GetMetros returns the Metros field if non-nil, zero value otherwise.

### GetMetrosOk

`func (o *ServiceProfile) GetMetrosOk() (*[]ServiceMetro, bool)`

GetMetrosOk returns a tuple with the Metros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetros

`func (o *ServiceProfile) SetMetros(v []ServiceMetro)`

SetMetros sets Metros field to given value.

### HasMetros

`func (o *ServiceProfile) HasMetros() bool`

HasMetros returns a boolean if a field has been set.

### GetSelfProfile

`func (o *ServiceProfile) GetSelfProfile() bool`

GetSelfProfile returns the SelfProfile field if non-nil, zero value otherwise.

### GetSelfProfileOk

`func (o *ServiceProfile) GetSelfProfileOk() (*bool, bool)`

GetSelfProfileOk returns a tuple with the SelfProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfProfile

`func (o *ServiceProfile) SetSelfProfile(v bool)`

SetSelfProfile sets SelfProfile field to given value.

### HasSelfProfile

`func (o *ServiceProfile) HasSelfProfile() bool`

HasSelfProfile returns a boolean if a field has been set.

### GetProjectId

`func (o *ServiceProfile) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ServiceProfile) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ServiceProfile) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ServiceProfile) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


