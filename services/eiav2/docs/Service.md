# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | [**ServiceType**](ServiceType.md) |  | 
**UseCase** | [**UseCaseType**](UseCaseType.md) |  | 
**Name** | **string** | Name of the service instance  | 
**Description** | Pointer to **string** | Description of the service instance  | [optional] 
**Bandwidth** | **int32** | Service bandwidth in Mbps  | 
**MinBandwidthCommit** | Pointer to **int32** | Service min bandwidth commit in Mbps  | [optional] 
**Uuid** | **string** |  | 
**Account** | Pointer to [**CustomerBillingAccount**](CustomerBillingAccount.md) |  | [optional] 
**Billing** | [**BillingType**](BillingType.md) |  | [default to BILLINGTYPE_FIXED]
**ChangeLog** | Pointer to [**ServiceChangeLog**](ServiceChangeLog.md) |  | [optional] 
**Draft** | **bool** |  | [default to false]
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 
**Order** | Pointer to [**ServiceOrderReference**](ServiceOrderReference.md) |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**State** | [**ServiceState**](ServiceState.md) |  | 

## Methods

### NewService

`func NewService(type_ ServiceType, useCase UseCaseType, name string, bandwidth int32, uuid string, billing BillingType, draft bool, state ServiceState, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *Service) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Service) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Service) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Service) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *Service) GetType() ServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Service) GetTypeOk() (*ServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Service) SetType(v ServiceType)`

SetType sets Type field to given value.


### GetUseCase

`func (o *Service) GetUseCase() UseCaseType`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *Service) GetUseCaseOk() (*UseCaseType, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *Service) SetUseCase(v UseCaseType)`

SetUseCase sets UseCase field to given value.


### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Service) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Service) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Service) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Service) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBandwidth

`func (o *Service) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *Service) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *Service) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetMinBandwidthCommit

`func (o *Service) GetMinBandwidthCommit() int32`

GetMinBandwidthCommit returns the MinBandwidthCommit field if non-nil, zero value otherwise.

### GetMinBandwidthCommitOk

`func (o *Service) GetMinBandwidthCommitOk() (*int32, bool)`

GetMinBandwidthCommitOk returns a tuple with the MinBandwidthCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBandwidthCommit

`func (o *Service) SetMinBandwidthCommit(v int32)`

SetMinBandwidthCommit sets MinBandwidthCommit field to given value.

### HasMinBandwidthCommit

`func (o *Service) HasMinBandwidthCommit() bool`

HasMinBandwidthCommit returns a boolean if a field has been set.

### GetUuid

`func (o *Service) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Service) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Service) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetAccount

`func (o *Service) GetAccount() CustomerBillingAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Service) GetAccountOk() (*CustomerBillingAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Service) SetAccount(v CustomerBillingAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Service) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBilling

`func (o *Service) GetBilling() BillingType`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *Service) GetBillingOk() (*BillingType, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *Service) SetBilling(v BillingType)`

SetBilling sets Billing field to given value.


### GetChangeLog

`func (o *Service) GetChangeLog() ServiceChangeLog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *Service) GetChangeLogOk() (*ServiceChangeLog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *Service) SetChangeLog(v ServiceChangeLog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *Service) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetDraft

`func (o *Service) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *Service) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *Service) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetLinks

`func (o *Service) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Service) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Service) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Service) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOrder

`func (o *Service) GetOrder() ServiceOrderReference`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Service) GetOrderOk() (*ServiceOrderReference, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Service) SetOrder(v ServiceOrderReference)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Service) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProject

`func (o *Service) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Service) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Service) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *Service) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *Service) GetState() ServiceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Service) GetStateOk() (*ServiceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Service) SetState(v ServiceState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


