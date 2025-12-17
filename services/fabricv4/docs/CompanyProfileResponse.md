# CompanyProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**State** | Pointer to **map[string]interface{}** |  | [optional] 
**Logo** | Pointer to [**CompanyLogo**](CompanyLogo.md) |  | [optional] 
**Tags** | Pointer to [**[]TagResponse**](TagResponse.md) |  | [optional] 
**ServiceProfiles** | Pointer to [**[]CompanyServiceProfile**](CompanyServiceProfile.md) |  | [optional] 
**PrivateServices** | Pointer to [**[]PrivateService**](PrivateService.md) |  | [optional] 
**Notifications** | Pointer to **[]map[string]interface{}** |  | [optional] 
**WebUrl** | Pointer to **string** |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewCompanyProfileResponse

`func NewCompanyProfileResponse() *CompanyProfileResponse`

NewCompanyProfileResponse instantiates a new CompanyProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyProfileResponseWithDefaults

`func NewCompanyProfileResponseWithDefaults() *CompanyProfileResponse`

NewCompanyProfileResponseWithDefaults instantiates a new CompanyProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CompanyProfileResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CompanyProfileResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CompanyProfileResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CompanyProfileResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *CompanyProfileResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CompanyProfileResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CompanyProfileResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CompanyProfileResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *CompanyProfileResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyProfileResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyProfileResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CompanyProfileResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CompanyProfileResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyProfileResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyProfileResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyProfileResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSummary

`func (o *CompanyProfileResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CompanyProfileResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CompanyProfileResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CompanyProfileResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDescription

`func (o *CompanyProfileResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CompanyProfileResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CompanyProfileResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CompanyProfileResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *CompanyProfileResponse) GetState() map[string]interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CompanyProfileResponse) GetStateOk() (*map[string]interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CompanyProfileResponse) SetState(v map[string]interface{})`

SetState sets State field to given value.

### HasState

`func (o *CompanyProfileResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLogo

`func (o *CompanyProfileResponse) GetLogo() CompanyLogo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CompanyProfileResponse) GetLogoOk() (*CompanyLogo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CompanyProfileResponse) SetLogo(v CompanyLogo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CompanyProfileResponse) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetTags

`func (o *CompanyProfileResponse) GetTags() []TagResponse`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CompanyProfileResponse) GetTagsOk() (*[]TagResponse, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CompanyProfileResponse) SetTags(v []TagResponse)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CompanyProfileResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetServiceProfiles

`func (o *CompanyProfileResponse) GetServiceProfiles() []CompanyServiceProfile`

GetServiceProfiles returns the ServiceProfiles field if non-nil, zero value otherwise.

### GetServiceProfilesOk

`func (o *CompanyProfileResponse) GetServiceProfilesOk() (*[]CompanyServiceProfile, bool)`

GetServiceProfilesOk returns a tuple with the ServiceProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfiles

`func (o *CompanyProfileResponse) SetServiceProfiles(v []CompanyServiceProfile)`

SetServiceProfiles sets ServiceProfiles field to given value.

### HasServiceProfiles

`func (o *CompanyProfileResponse) HasServiceProfiles() bool`

HasServiceProfiles returns a boolean if a field has been set.

### GetPrivateServices

`func (o *CompanyProfileResponse) GetPrivateServices() []PrivateService`

GetPrivateServices returns the PrivateServices field if non-nil, zero value otherwise.

### GetPrivateServicesOk

`func (o *CompanyProfileResponse) GetPrivateServicesOk() (*[]PrivateService, bool)`

GetPrivateServicesOk returns a tuple with the PrivateServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServices

`func (o *CompanyProfileResponse) SetPrivateServices(v []PrivateService)`

SetPrivateServices sets PrivateServices field to given value.

### HasPrivateServices

`func (o *CompanyProfileResponse) HasPrivateServices() bool`

HasPrivateServices returns a boolean if a field has been set.

### GetNotifications

`func (o *CompanyProfileResponse) GetNotifications() []map[string]interface{}`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *CompanyProfileResponse) GetNotificationsOk() (*[]map[string]interface{}, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *CompanyProfileResponse) SetNotifications(v []map[string]interface{})`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *CompanyProfileResponse) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetWebUrl

`func (o *CompanyProfileResponse) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *CompanyProfileResponse) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *CompanyProfileResponse) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *CompanyProfileResponse) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetChangeLog

`func (o *CompanyProfileResponse) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *CompanyProfileResponse) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *CompanyProfileResponse) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *CompanyProfileResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


