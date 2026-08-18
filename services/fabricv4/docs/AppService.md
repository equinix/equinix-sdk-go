# AppService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | [**AppServiceType**](AppServiceType.md) |  | [default to APPSERVICETYPE_APP_SERVICE]
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Name** | **string** | Customer-provided App Service name | 
**Description** | Pointer to **string** | Customer-provided App Service description | [optional] 
**State** | Pointer to [**AppServiceState**](AppServiceState.md) |  | [optional] 
**Endpoint** | Pointer to **string** | Accessible endpoint through this service | [optional] 
**SourceDomains** | Pointer to **[]string** | List of source domains from where traffic is allowed | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**Change** | Pointer to [**AppServiceChange**](AppServiceChange.md) |  | [optional] 

## Methods

### NewAppService

`func NewAppService(type_ AppServiceType, name string, ) *AppService`

NewAppService instantiates a new AppService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceWithDefaults

`func NewAppServiceWithDefaults() *AppService`

NewAppServiceWithDefaults instantiates a new AppService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppService) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppService) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppService) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppService) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppService) GetType() AppServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppService) GetTypeOk() (*AppServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppService) SetType(v AppServiceType)`

SetType sets Type field to given value.


### GetUuid

`func (o *AppService) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppService) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppService) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AppService) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *AppService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppService) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppService) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppService) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppService) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppService) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *AppService) GetState() AppServiceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppService) GetStateOk() (*AppServiceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppService) SetState(v AppServiceState)`

SetState sets State field to given value.

### HasState

`func (o *AppService) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEndpoint

`func (o *AppService) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AppService) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AppService) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AppService) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetSourceDomains

`func (o *AppService) GetSourceDomains() []string`

GetSourceDomains returns the SourceDomains field if non-nil, zero value otherwise.

### GetSourceDomainsOk

`func (o *AppService) GetSourceDomainsOk() (*[]string, bool)`

GetSourceDomainsOk returns a tuple with the SourceDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDomains

`func (o *AppService) SetSourceDomains(v []string)`

SetSourceDomains sets SourceDomains field to given value.

### HasSourceDomains

`func (o *AppService) HasSourceDomains() bool`

HasSourceDomains returns a boolean if a field has been set.

### GetAccount

`func (o *AppService) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AppService) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AppService) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AppService) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetProject

`func (o *AppService) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AppService) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AppService) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *AppService) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetChangeLog

`func (o *AppService) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *AppService) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *AppService) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *AppService) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetChange

`func (o *AppService) GetChange() AppServiceChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *AppService) GetChangeOk() (*AppServiceChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *AppService) SetChange(v AppServiceChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *AppService) HasChange() bool`

HasChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


