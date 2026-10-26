# AppProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | [**AppProcessorType**](AppProcessorType.md) |  | [default to APPPROCESSORTYPE_APP_PROCESSOR]
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Name** | **string** | Customer-provided App Processor name | 
**Description** | Pointer to **string** | Customer-provided App Processor description | [optional] 
**State** | Pointer to [**AppProcessorState**](AppProcessorState.md) |  | [optional] 
**RegistryId** | **string** | Container image reference for the processor runtime | 
**Marketplace** | [**AppProcessorMarketplace**](AppProcessorMarketplace.md) |  | 
**Project** | [**Project**](Project.md) |  | 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**Change** | Pointer to [**AppProcessorChange**](AppProcessorChange.md) |  | [optional] 

## Methods

### NewAppProcessor

`func NewAppProcessor(type_ AppProcessorType, name string, registryId string, marketplace AppProcessorMarketplace, project Project, ) *AppProcessor`

NewAppProcessor instantiates a new AppProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppProcessorWithDefaults

`func NewAppProcessorWithDefaults() *AppProcessor`

NewAppProcessorWithDefaults instantiates a new AppProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppProcessor) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppProcessor) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppProcessor) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppProcessor) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppProcessor) GetType() AppProcessorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppProcessor) GetTypeOk() (*AppProcessorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppProcessor) SetType(v AppProcessorType)`

SetType sets Type field to given value.


### GetUuid

`func (o *AppProcessor) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppProcessor) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppProcessor) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AppProcessor) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *AppProcessor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppProcessor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppProcessor) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppProcessor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppProcessor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppProcessor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppProcessor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *AppProcessor) GetState() AppProcessorState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppProcessor) GetStateOk() (*AppProcessorState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppProcessor) SetState(v AppProcessorState)`

SetState sets State field to given value.

### HasState

`func (o *AppProcessor) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRegistryId

`func (o *AppProcessor) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *AppProcessor) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *AppProcessor) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.


### GetMarketplace

`func (o *AppProcessor) GetMarketplace() AppProcessorMarketplace`

GetMarketplace returns the Marketplace field if non-nil, zero value otherwise.

### GetMarketplaceOk

`func (o *AppProcessor) GetMarketplaceOk() (*AppProcessorMarketplace, bool)`

GetMarketplaceOk returns a tuple with the Marketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplace

`func (o *AppProcessor) SetMarketplace(v AppProcessorMarketplace)`

SetMarketplace sets Marketplace field to given value.


### GetProject

`func (o *AppProcessor) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AppProcessor) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AppProcessor) SetProject(v Project)`

SetProject sets Project field to given value.


### GetChangeLog

`func (o *AppProcessor) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *AppProcessor) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *AppProcessor) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *AppProcessor) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetChange

`func (o *AppProcessor) GetChange() AppProcessorChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *AppProcessor) GetChangeOk() (*AppProcessorChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *AppProcessor) SetChange(v AppProcessorChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *AppProcessor) HasChange() bool`

HasChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


