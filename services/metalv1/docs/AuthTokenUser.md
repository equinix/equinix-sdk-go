# AuthTokenUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarThumbUrl** | Pointer to **string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultOrganizationId** | Pointer to **string** |  | [optional] 
**DefaultProjectId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**FraudScore** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastLoginAt** | Pointer to **time.Time** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**MaxOrganizations** | Pointer to **int32** |  | [optional] 
**MaxProjects** | Pointer to **int32** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**ShortId** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**TwoFactorAuth** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAuthTokenUser

`func NewAuthTokenUser() *AuthTokenUser`

NewAuthTokenUser instantiates a new AuthTokenUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenUserWithDefaults

`func NewAuthTokenUserWithDefaults() *AuthTokenUser`

NewAuthTokenUserWithDefaults instantiates a new AuthTokenUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarThumbUrl

`func (o *AuthTokenUser) GetAvatarThumbUrl() string`

GetAvatarThumbUrl returns the AvatarThumbUrl field if non-nil, zero value otherwise.

### GetAvatarThumbUrlOk

`func (o *AuthTokenUser) GetAvatarThumbUrlOk() (*string, bool)`

GetAvatarThumbUrlOk returns a tuple with the AvatarThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarThumbUrl

`func (o *AuthTokenUser) SetAvatarThumbUrl(v string)`

SetAvatarThumbUrl sets AvatarThumbUrl field to given value.

### HasAvatarThumbUrl

`func (o *AuthTokenUser) HasAvatarThumbUrl() bool`

HasAvatarThumbUrl returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *AuthTokenUser) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *AuthTokenUser) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *AuthTokenUser) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *AuthTokenUser) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthTokenUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthTokenUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthTokenUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthTokenUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomdata

`func (o *AuthTokenUser) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *AuthTokenUser) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *AuthTokenUser) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *AuthTokenUser) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDefaultOrganizationId

`func (o *AuthTokenUser) GetDefaultOrganizationId() string`

GetDefaultOrganizationId returns the DefaultOrganizationId field if non-nil, zero value otherwise.

### GetDefaultOrganizationIdOk

`func (o *AuthTokenUser) GetDefaultOrganizationIdOk() (*string, bool)`

GetDefaultOrganizationIdOk returns a tuple with the DefaultOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOrganizationId

`func (o *AuthTokenUser) SetDefaultOrganizationId(v string)`

SetDefaultOrganizationId sets DefaultOrganizationId field to given value.

### HasDefaultOrganizationId

`func (o *AuthTokenUser) HasDefaultOrganizationId() bool`

HasDefaultOrganizationId returns a boolean if a field has been set.

### GetDefaultProjectId

`func (o *AuthTokenUser) GetDefaultProjectId() string`

GetDefaultProjectId returns the DefaultProjectId field if non-nil, zero value otherwise.

### GetDefaultProjectIdOk

`func (o *AuthTokenUser) GetDefaultProjectIdOk() (*string, bool)`

GetDefaultProjectIdOk returns a tuple with the DefaultProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProjectId

`func (o *AuthTokenUser) SetDefaultProjectId(v string)`

SetDefaultProjectId sets DefaultProjectId field to given value.

### HasDefaultProjectId

`func (o *AuthTokenUser) HasDefaultProjectId() bool`

HasDefaultProjectId returns a boolean if a field has been set.

### GetEmail

`func (o *AuthTokenUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthTokenUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthTokenUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthTokenUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmails

`func (o *AuthTokenUser) GetEmails() []Href`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *AuthTokenUser) GetEmailsOk() (*[]Href, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *AuthTokenUser) SetEmails(v []Href)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *AuthTokenUser) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetFirstName

`func (o *AuthTokenUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AuthTokenUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AuthTokenUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AuthTokenUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetFraudScore

`func (o *AuthTokenUser) GetFraudScore() string`

GetFraudScore returns the FraudScore field if non-nil, zero value otherwise.

### GetFraudScoreOk

`func (o *AuthTokenUser) GetFraudScoreOk() (*string, bool)`

GetFraudScoreOk returns a tuple with the FraudScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudScore

`func (o *AuthTokenUser) SetFraudScore(v string)`

SetFraudScore sets FraudScore field to given value.

### HasFraudScore

`func (o *AuthTokenUser) HasFraudScore() bool`

HasFraudScore returns a boolean if a field has been set.

### GetFullName

`func (o *AuthTokenUser) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AuthTokenUser) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AuthTokenUser) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AuthTokenUser) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetHref

`func (o *AuthTokenUser) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AuthTokenUser) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AuthTokenUser) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AuthTokenUser) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *AuthTokenUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthTokenUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthTokenUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthTokenUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *AuthTokenUser) GetLastLoginAt() time.Time`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *AuthTokenUser) GetLastLoginAtOk() (*time.Time, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *AuthTokenUser) SetLastLoginAt(v time.Time)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *AuthTokenUser) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### GetLastName

`func (o *AuthTokenUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AuthTokenUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AuthTokenUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AuthTokenUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMaxOrganizations

`func (o *AuthTokenUser) GetMaxOrganizations() int32`

GetMaxOrganizations returns the MaxOrganizations field if non-nil, zero value otherwise.

### GetMaxOrganizationsOk

`func (o *AuthTokenUser) GetMaxOrganizationsOk() (*int32, bool)`

GetMaxOrganizationsOk returns a tuple with the MaxOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrganizations

`func (o *AuthTokenUser) SetMaxOrganizations(v int32)`

SetMaxOrganizations sets MaxOrganizations field to given value.

### HasMaxOrganizations

`func (o *AuthTokenUser) HasMaxOrganizations() bool`

HasMaxOrganizations returns a boolean if a field has been set.

### GetMaxProjects

`func (o *AuthTokenUser) GetMaxProjects() int32`

GetMaxProjects returns the MaxProjects field if non-nil, zero value otherwise.

### GetMaxProjectsOk

`func (o *AuthTokenUser) GetMaxProjectsOk() (*int32, bool)`

GetMaxProjectsOk returns a tuple with the MaxProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProjects

`func (o *AuthTokenUser) SetMaxProjects(v int32)`

SetMaxProjects sets MaxProjects field to given value.

### HasMaxProjects

`func (o *AuthTokenUser) HasMaxProjects() bool`

HasMaxProjects returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *AuthTokenUser) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *AuthTokenUser) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *AuthTokenUser) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *AuthTokenUser) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetShortId

`func (o *AuthTokenUser) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *AuthTokenUser) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *AuthTokenUser) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *AuthTokenUser) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetTimezone

`func (o *AuthTokenUser) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AuthTokenUser) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AuthTokenUser) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AuthTokenUser) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTwoFactorAuth

`func (o *AuthTokenUser) GetTwoFactorAuth() string`

GetTwoFactorAuth returns the TwoFactorAuth field if non-nil, zero value otherwise.

### GetTwoFactorAuthOk

`func (o *AuthTokenUser) GetTwoFactorAuthOk() (*string, bool)`

GetTwoFactorAuthOk returns a tuple with the TwoFactorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuth

`func (o *AuthTokenUser) SetTwoFactorAuth(v string)`

SetTwoFactorAuth sets TwoFactorAuth field to given value.

### HasTwoFactorAuth

`func (o *AuthTokenUser) HasTwoFactorAuth() bool`

HasTwoFactorAuth returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthTokenUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthTokenUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthTokenUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthTokenUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


