# AppDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | [**AppDomainType**](AppDomainType.md) |  | [default to APPDOMAINTYPE_APP_DOMAIN]
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Name** | **string** | Customer-provided App Domain name | 
**Description** | Pointer to **string** | Customer-provided App Domain description | [optional] 
**State** | Pointer to [**AppDomainState**](AppDomainState.md) |  | [optional] 
**Project** | [**Project**](Project.md) |  | 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**Change** | Pointer to [**AppDomainChange**](AppDomainChange.md) |  | [optional] 

## Methods

### NewAppDomain

`func NewAppDomain(type_ AppDomainType, name string, project Project, ) *AppDomain`

NewAppDomain instantiates a new AppDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDomainWithDefaults

`func NewAppDomainWithDefaults() *AppDomain`

NewAppDomainWithDefaults instantiates a new AppDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppDomain) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppDomain) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppDomain) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppDomain) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppDomain) GetType() AppDomainType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppDomain) GetTypeOk() (*AppDomainType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppDomain) SetType(v AppDomainType)`

SetType sets Type field to given value.


### GetUuid

`func (o *AppDomain) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppDomain) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppDomain) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AppDomain) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *AppDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDomain) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppDomain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppDomain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppDomain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppDomain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *AppDomain) GetState() AppDomainState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppDomain) GetStateOk() (*AppDomainState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppDomain) SetState(v AppDomainState)`

SetState sets State field to given value.

### HasState

`func (o *AppDomain) HasState() bool`

HasState returns a boolean if a field has been set.

### GetProject

`func (o *AppDomain) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AppDomain) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AppDomain) SetProject(v Project)`

SetProject sets Project field to given value.


### GetChangeLog

`func (o *AppDomain) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *AppDomain) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *AppDomain) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *AppDomain) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetChange

`func (o *AppDomain) GetChange() AppDomainChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *AppDomain) GetChangeOk() (*AppDomainChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *AppDomain) SetChange(v AppDomainChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *AppDomain) HasChange() bool`

HasChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


