# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpConfig** | Pointer to [**Href**](Href.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Devices** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Invitations** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**MaxDevices** | Pointer to **map[string]interface{}** |  | [optional] 
**Members** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Memberships** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the project. Cannot contain characters encoded in greater than 3 bytes such as emojis. | [optional] 
**NetworkStatus** | Pointer to **map[string]interface{}** |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**PaymentMethod** | Pointer to [**Href**](Href.md) |  | [optional] 
**SshKeys** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Type** | Pointer to [**ProjectType**](ProjectType.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProject

`func NewProject() *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpConfig

`func (o *Project) GetBgpConfig() Href`

GetBgpConfig returns the BgpConfig field if non-nil, zero value otherwise.

### GetBgpConfigOk

`func (o *Project) GetBgpConfigOk() (*Href, bool)`

GetBgpConfigOk returns a tuple with the BgpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfig

`func (o *Project) SetBgpConfig(v Href)`

SetBgpConfig sets BgpConfig field to given value.

### HasBgpConfig

`func (o *Project) HasBgpConfig() bool`

HasBgpConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Project) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Project) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Project) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Project) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomdata

`func (o *Project) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *Project) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *Project) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *Project) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDevices

`func (o *Project) GetDevices() []Href`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *Project) GetDevicesOk() (*[]Href, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *Project) SetDevices(v []Href)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *Project) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetHref

`func (o *Project) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Project) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Project) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Project) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Project) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Project) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvitations

`func (o *Project) GetInvitations() []Href`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *Project) GetInvitationsOk() (*[]Href, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *Project) SetInvitations(v []Href)`

SetInvitations sets Invitations field to given value.

### HasInvitations

`func (o *Project) HasInvitations() bool`

HasInvitations returns a boolean if a field has been set.

### GetMaxDevices

`func (o *Project) GetMaxDevices() map[string]interface{}`

GetMaxDevices returns the MaxDevices field if non-nil, zero value otherwise.

### GetMaxDevicesOk

`func (o *Project) GetMaxDevicesOk() (*map[string]interface{}, bool)`

GetMaxDevicesOk returns a tuple with the MaxDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDevices

`func (o *Project) SetMaxDevices(v map[string]interface{})`

SetMaxDevices sets MaxDevices field to given value.

### HasMaxDevices

`func (o *Project) HasMaxDevices() bool`

HasMaxDevices returns a boolean if a field has been set.

### GetMembers

`func (o *Project) GetMembers() []Href`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Project) GetMembersOk() (*[]Href, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Project) SetMembers(v []Href)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Project) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMemberships

`func (o *Project) GetMemberships() []Href`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *Project) GetMembershipsOk() (*[]Href, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *Project) SetMemberships(v []Href)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *Project) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.

### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Project) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkStatus

`func (o *Project) GetNetworkStatus() map[string]interface{}`

GetNetworkStatus returns the NetworkStatus field if non-nil, zero value otherwise.

### GetNetworkStatusOk

`func (o *Project) GetNetworkStatusOk() (*map[string]interface{}, bool)`

GetNetworkStatusOk returns a tuple with the NetworkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatus

`func (o *Project) SetNetworkStatus(v map[string]interface{})`

SetNetworkStatus sets NetworkStatus field to given value.

### HasNetworkStatus

`func (o *Project) HasNetworkStatus() bool`

HasNetworkStatus returns a boolean if a field has been set.

### GetOrganization

`func (o *Project) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Project) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Project) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Project) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *Project) GetPaymentMethod() Href`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *Project) GetPaymentMethodOk() (*Href, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *Project) SetPaymentMethod(v Href)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *Project) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetSshKeys

`func (o *Project) GetSshKeys() []Href`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *Project) GetSshKeysOk() (*[]Href, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *Project) SetSshKeys(v []Href)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *Project) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Project) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Project) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Project) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Project) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *Project) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Project) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Project) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Project) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVolumes

`func (o *Project) GetVolumes() []Href`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *Project) GetVolumesOk() (*[]Href, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *Project) SetVolumes(v []Href)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *Project) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetType

`func (o *Project) GetType() ProjectType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Project) GetTypeOk() (*ProjectType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Project) SetType(v ProjectType)`

SetType sets Type field to given value.

### HasType

`func (o *Project) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTags

`func (o *Project) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Project) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Project) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Project) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


