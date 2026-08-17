# AppLinkAppDomainAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | Pointer to [**AppDomainType**](AppDomainType.md) |  | [optional] [default to APPDOMAINTYPE_APP_DOMAIN]
**Uuid** | **string** | Equinix-assigned access point identifier | 
**Name** | Pointer to **string** | Customer-provided App Domain name | [optional] 
**AttachmentStatus** | Pointer to [**AppLinkAttachState**](AppLinkAttachState.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewAppLinkAppDomainAttachment

`func NewAppLinkAppDomainAttachment(uuid string, ) *AppLinkAppDomainAttachment`

NewAppLinkAppDomainAttachment instantiates a new AppLinkAppDomainAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkAppDomainAttachmentWithDefaults

`func NewAppLinkAppDomainAttachmentWithDefaults() *AppLinkAppDomainAttachment`

NewAppLinkAppDomainAttachmentWithDefaults instantiates a new AppLinkAppDomainAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppLinkAppDomainAttachment) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppLinkAppDomainAttachment) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppLinkAppDomainAttachment) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppLinkAppDomainAttachment) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppLinkAppDomainAttachment) GetType() AppDomainType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppLinkAppDomainAttachment) GetTypeOk() (*AppDomainType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppLinkAppDomainAttachment) SetType(v AppDomainType)`

SetType sets Type field to given value.

### HasType

`func (o *AppLinkAppDomainAttachment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *AppLinkAppDomainAttachment) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppLinkAppDomainAttachment) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppLinkAppDomainAttachment) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *AppLinkAppDomainAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppLinkAppDomainAttachment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppLinkAppDomainAttachment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppLinkAppDomainAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAttachmentStatus

`func (o *AppLinkAppDomainAttachment) GetAttachmentStatus() AppLinkAttachState`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *AppLinkAppDomainAttachment) GetAttachmentStatusOk() (*AppLinkAttachState, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *AppLinkAppDomainAttachment) SetAttachmentStatus(v AppLinkAttachState)`

SetAttachmentStatus sets AttachmentStatus field to given value.

### HasAttachmentStatus

`func (o *AppLinkAppDomainAttachment) HasAttachmentStatus() bool`

HasAttachmentStatus returns a boolean if a field has been set.

### GetChangeLog

`func (o *AppLinkAppDomainAttachment) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *AppLinkAppDomainAttachment) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *AppLinkAppDomainAttachment) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *AppLinkAppDomainAttachment) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


