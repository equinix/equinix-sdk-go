# AppLinkAttachedAppDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | [**AppDomainType**](AppDomainType.md) |  | [default to APPDOMAINTYPE_APP_DOMAIN]
**Uuid** | **string** | Equinix-assigned access point identifier | 
**Name** | Pointer to **string** | Customer-provided App Domain name | [optional] 
**Description** | Pointer to **string** | Customer-provided App Domain description | [optional] 
**State** | Pointer to [**AppDomainState**](AppDomainState.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**AttachmentStatus** | Pointer to [**AppLinkAttachState**](AppLinkAttachState.md) |  | [optional] 

## Methods

### NewAppLinkAttachedAppDomain

`func NewAppLinkAttachedAppDomain(type_ AppDomainType, uuid string, ) *AppLinkAttachedAppDomain`

NewAppLinkAttachedAppDomain instantiates a new AppLinkAttachedAppDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkAttachedAppDomainWithDefaults

`func NewAppLinkAttachedAppDomainWithDefaults() *AppLinkAttachedAppDomain`

NewAppLinkAttachedAppDomainWithDefaults instantiates a new AppLinkAttachedAppDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppLinkAttachedAppDomain) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppLinkAttachedAppDomain) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppLinkAttachedAppDomain) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppLinkAttachedAppDomain) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppLinkAttachedAppDomain) GetType() AppDomainType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppLinkAttachedAppDomain) GetTypeOk() (*AppDomainType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppLinkAttachedAppDomain) SetType(v AppDomainType)`

SetType sets Type field to given value.


### GetUuid

`func (o *AppLinkAttachedAppDomain) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppLinkAttachedAppDomain) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppLinkAttachedAppDomain) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *AppLinkAttachedAppDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppLinkAttachedAppDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppLinkAttachedAppDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppLinkAttachedAppDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AppLinkAttachedAppDomain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppLinkAttachedAppDomain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppLinkAttachedAppDomain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppLinkAttachedAppDomain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *AppLinkAttachedAppDomain) GetState() AppDomainState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppLinkAttachedAppDomain) GetStateOk() (*AppDomainState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppLinkAttachedAppDomain) SetState(v AppDomainState)`

SetState sets State field to given value.

### HasState

`func (o *AppLinkAttachedAppDomain) HasState() bool`

HasState returns a boolean if a field has been set.

### GetProject

`func (o *AppLinkAttachedAppDomain) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AppLinkAttachedAppDomain) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AppLinkAttachedAppDomain) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *AppLinkAttachedAppDomain) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAttachmentStatus

`func (o *AppLinkAttachedAppDomain) GetAttachmentStatus() AppLinkAttachState`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *AppLinkAttachedAppDomain) GetAttachmentStatusOk() (*AppLinkAttachState, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *AppLinkAttachedAppDomain) SetAttachmentStatus(v AppLinkAttachState)`

SetAttachmentStatus sets AttachmentStatus field to given value.

### HasAttachmentStatus

`func (o *AppLinkAttachedAppDomain) HasAttachmentStatus() bool`

HasAttachmentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


