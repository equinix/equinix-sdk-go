# UserCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **string** |  | [optional] 
**CompanyUrl** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Emails** | [**[]EmailInput**](EmailInput.md) |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Level** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**SocialAccounts** | Pointer to **map[string]interface{}** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TwoFactorAuth** | Pointer to **string** |  | [optional] 
**VerifiedAt** | Pointer to **time.Time** |  | [optional] 
**InvitationId** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 

## Methods

### NewUserCreateInput

`func NewUserCreateInput(emails []EmailInput, firstName string, lastName string, ) *UserCreateInput`

NewUserCreateInput instantiates a new UserCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateInputWithDefaults

`func NewUserCreateInputWithDefaults() *UserCreateInput`

NewUserCreateInputWithDefaults instantiates a new UserCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *UserCreateInput) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UserCreateInput) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UserCreateInput) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UserCreateInput) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompanyUrl

`func (o *UserCreateInput) GetCompanyUrl() string`

GetCompanyUrl returns the CompanyUrl field if non-nil, zero value otherwise.

### GetCompanyUrlOk

`func (o *UserCreateInput) GetCompanyUrlOk() (*string, bool)`

GetCompanyUrlOk returns a tuple with the CompanyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyUrl

`func (o *UserCreateInput) SetCompanyUrl(v string)`

SetCompanyUrl sets CompanyUrl field to given value.

### HasCompanyUrl

`func (o *UserCreateInput) HasCompanyUrl() bool`

HasCompanyUrl returns a boolean if a field has been set.

### GetCustomdata

`func (o *UserCreateInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *UserCreateInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *UserCreateInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *UserCreateInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetEmails

`func (o *UserCreateInput) GetEmails() []EmailInput`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UserCreateInput) GetEmailsOk() (*[]EmailInput, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UserCreateInput) SetEmails(v []EmailInput)`

SetEmails sets Emails field to given value.


### GetFirstName

`func (o *UserCreateInput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserCreateInput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserCreateInput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UserCreateInput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserCreateInput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserCreateInput) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetLevel

`func (o *UserCreateInput) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *UserCreateInput) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *UserCreateInput) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *UserCreateInput) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetPassword

`func (o *UserCreateInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreateInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreateInput) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCreateInput) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *UserCreateInput) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserCreateInput) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserCreateInput) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserCreateInput) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetSocialAccounts

`func (o *UserCreateInput) GetSocialAccounts() map[string]interface{}`

GetSocialAccounts returns the SocialAccounts field if non-nil, zero value otherwise.

### GetSocialAccountsOk

`func (o *UserCreateInput) GetSocialAccountsOk() (*map[string]interface{}, bool)`

GetSocialAccountsOk returns a tuple with the SocialAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialAccounts

`func (o *UserCreateInput) SetSocialAccounts(v map[string]interface{})`

SetSocialAccounts sets SocialAccounts field to given value.

### HasSocialAccounts

`func (o *UserCreateInput) HasSocialAccounts() bool`

HasSocialAccounts returns a boolean if a field has been set.

### GetTimezone

`func (o *UserCreateInput) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserCreateInput) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserCreateInput) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserCreateInput) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTitle

`func (o *UserCreateInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserCreateInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserCreateInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserCreateInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTwoFactorAuth

`func (o *UserCreateInput) GetTwoFactorAuth() string`

GetTwoFactorAuth returns the TwoFactorAuth field if non-nil, zero value otherwise.

### GetTwoFactorAuthOk

`func (o *UserCreateInput) GetTwoFactorAuthOk() (*string, bool)`

GetTwoFactorAuthOk returns a tuple with the TwoFactorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuth

`func (o *UserCreateInput) SetTwoFactorAuth(v string)`

SetTwoFactorAuth sets TwoFactorAuth field to given value.

### HasTwoFactorAuth

`func (o *UserCreateInput) HasTwoFactorAuth() bool`

HasTwoFactorAuth returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *UserCreateInput) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *UserCreateInput) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *UserCreateInput) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *UserCreateInput) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### GetInvitationId

`func (o *UserCreateInput) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *UserCreateInput) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *UserCreateInput) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.

### HasInvitationId

`func (o *UserCreateInput) HasInvitationId() bool`

HasInvitationId returns a boolean if a field has been set.

### GetNonce

`func (o *UserCreateInput) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *UserCreateInput) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *UserCreateInput) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *UserCreateInput) HasNonce() bool`

HasNonce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


