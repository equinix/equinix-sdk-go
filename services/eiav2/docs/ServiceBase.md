# ServiceBase

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

## Methods

### NewServiceBase

`func NewServiceBase(type_ ServiceType, useCase UseCaseType, name string, bandwidth int32, ) *ServiceBase`

NewServiceBase instantiates a new ServiceBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBaseWithDefaults

`func NewServiceBaseWithDefaults() *ServiceBase`

NewServiceBaseWithDefaults instantiates a new ServiceBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ServiceBase) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceBase) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceBase) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceBase) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *ServiceBase) GetType() ServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceBase) GetTypeOk() (*ServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceBase) SetType(v ServiceType)`

SetType sets Type field to given value.


### GetUseCase

`func (o *ServiceBase) GetUseCase() UseCaseType`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *ServiceBase) GetUseCaseOk() (*UseCaseType, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *ServiceBase) SetUseCase(v UseCaseType)`

SetUseCase sets UseCase field to given value.


### GetName

`func (o *ServiceBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServiceBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBandwidth

`func (o *ServiceBase) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *ServiceBase) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *ServiceBase) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetMinBandwidthCommit

`func (o *ServiceBase) GetMinBandwidthCommit() int32`

GetMinBandwidthCommit returns the MinBandwidthCommit field if non-nil, zero value otherwise.

### GetMinBandwidthCommitOk

`func (o *ServiceBase) GetMinBandwidthCommitOk() (*int32, bool)`

GetMinBandwidthCommitOk returns a tuple with the MinBandwidthCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBandwidthCommit

`func (o *ServiceBase) SetMinBandwidthCommit(v int32)`

SetMinBandwidthCommit sets MinBandwidthCommit field to given value.

### HasMinBandwidthCommit

`func (o *ServiceBase) HasMinBandwidthCommit() bool`

HasMinBandwidthCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


