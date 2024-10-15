# ServiceToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ServiceTokenType**](ServiceTokenType.md) |  | [optional] 
**Href** | Pointer to **string** | An absolute URL that is the subject of the link&#39;s context. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned service token identifier | [optional] 
**IssuerSide** | Pointer to [**ServiceTokenIssuerSide**](ServiceTokenIssuerSide.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided service token name | [optional] 
**Description** | Pointer to **string** | Customer-provided service token description | [optional] 
**ExpirationDateTime** | Pointer to **time.Time** | Expiration date and time of the service token. | [optional] 
**Connection** | Pointer to [**ServiceTokenConnection**](ServiceTokenConnection.md) |  | [optional] 
**State** | Pointer to [**ServiceTokenState**](ServiceTokenState.md) |  | [optional] 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Service token related notifications | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 

## Methods

### NewServiceToken

`func NewServiceToken() *ServiceToken`

NewServiceToken instantiates a new ServiceToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTokenWithDefaults

`func NewServiceTokenWithDefaults() *ServiceToken`

NewServiceTokenWithDefaults instantiates a new ServiceToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceToken) GetType() ServiceTokenType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceToken) GetTypeOk() (*ServiceTokenType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceToken) SetType(v ServiceTokenType)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceToken) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *ServiceToken) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ServiceToken) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ServiceToken) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ServiceToken) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *ServiceToken) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceToken) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceToken) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceToken) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetIssuerSide

`func (o *ServiceToken) GetIssuerSide() ServiceTokenIssuerSide`

GetIssuerSide returns the IssuerSide field if non-nil, zero value otherwise.

### GetIssuerSideOk

`func (o *ServiceToken) GetIssuerSideOk() (*ServiceTokenIssuerSide, bool)`

GetIssuerSideOk returns a tuple with the IssuerSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerSide

`func (o *ServiceToken) SetIssuerSide(v ServiceTokenIssuerSide)`

SetIssuerSide sets IssuerSide field to given value.

### HasIssuerSide

`func (o *ServiceToken) HasIssuerSide() bool`

HasIssuerSide returns a boolean if a field has been set.

### GetName

`func (o *ServiceToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpirationDateTime

`func (o *ServiceToken) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *ServiceToken) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *ServiceToken) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *ServiceToken) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetConnection

`func (o *ServiceToken) GetConnection() ServiceTokenConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ServiceToken) GetConnectionOk() (*ServiceTokenConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ServiceToken) SetConnection(v ServiceTokenConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ServiceToken) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetState

`func (o *ServiceToken) GetState() ServiceTokenState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServiceToken) GetStateOk() (*ServiceTokenState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServiceToken) SetState(v ServiceTokenState)`

SetState sets State field to given value.

### HasState

`func (o *ServiceToken) HasState() bool`

HasState returns a boolean if a field has been set.

### GetNotifications

`func (o *ServiceToken) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ServiceToken) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ServiceToken) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ServiceToken) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetAccount

`func (o *ServiceToken) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ServiceToken) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ServiceToken) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ServiceToken) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetChangelog

`func (o *ServiceToken) GetChangelog() Changelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *ServiceToken) GetChangelogOk() (*Changelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *ServiceToken) SetChangelog(v Changelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *ServiceToken) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.

### GetProject

`func (o *ServiceToken) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServiceToken) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServiceToken) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *ServiceToken) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


