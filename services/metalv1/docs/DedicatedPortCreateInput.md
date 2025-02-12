# DedicatedPortCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAccountName** | Pointer to **string** | The billing account name of the Equinix Fabric account. | [optional] 
**ContactEmail** | Pointer to **string** | The preferred email used for communication and notifications about the Equinix Fabric interconnection. Optional and defaults to the primary user email address when using a User API key or the organization owner email address when using a Project API key. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Metro** | **string** | A Metro ID or code. For interconnections with Dedicated Ports, this will be the location of the issued Dedicated Ports. | 
**Mode** | Pointer to [**DedicatedPortCreateInputMode**](DedicatedPortCreateInputMode.md) |  | [optional] 
**Name** | **string** |  | 
**Project** | Pointer to **string** |  | [optional] 
**Redundancy** | **string** | Either &#39;primary&#39; or &#39;redundant&#39;. | 
**Speed** | Pointer to **string** | A interconnection speed, in bps, mbps, or gbps. For Dedicated Ports, this can be 10Gbps or 100Gbps. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | [**DedicatedPortCreateInputType**](DedicatedPortCreateInputType.md) |  | 
**UseCase** | Pointer to **string** | The intended use case of the dedicated port. | [optional] 

## Methods

### NewDedicatedPortCreateInput

`func NewDedicatedPortCreateInput(metro string, name string, redundancy string, type_ DedicatedPortCreateInputType, ) *DedicatedPortCreateInput`

NewDedicatedPortCreateInput instantiates a new DedicatedPortCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedPortCreateInputWithDefaults

`func NewDedicatedPortCreateInputWithDefaults() *DedicatedPortCreateInput`

NewDedicatedPortCreateInputWithDefaults instantiates a new DedicatedPortCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAccountName

`func (o *DedicatedPortCreateInput) GetBillingAccountName() string`

GetBillingAccountName returns the BillingAccountName field if non-nil, zero value otherwise.

### GetBillingAccountNameOk

`func (o *DedicatedPortCreateInput) GetBillingAccountNameOk() (*string, bool)`

GetBillingAccountNameOk returns a tuple with the BillingAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountName

`func (o *DedicatedPortCreateInput) SetBillingAccountName(v string)`

SetBillingAccountName sets BillingAccountName field to given value.

### HasBillingAccountName

`func (o *DedicatedPortCreateInput) HasBillingAccountName() bool`

HasBillingAccountName returns a boolean if a field has been set.

### GetContactEmail

`func (o *DedicatedPortCreateInput) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *DedicatedPortCreateInput) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *DedicatedPortCreateInput) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *DedicatedPortCreateInput) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetDescription

`func (o *DedicatedPortCreateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DedicatedPortCreateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DedicatedPortCreateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DedicatedPortCreateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetro

`func (o *DedicatedPortCreateInput) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *DedicatedPortCreateInput) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *DedicatedPortCreateInput) SetMetro(v string)`

SetMetro sets Metro field to given value.


### GetMode

`func (o *DedicatedPortCreateInput) GetMode() DedicatedPortCreateInputMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DedicatedPortCreateInput) GetModeOk() (*DedicatedPortCreateInputMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DedicatedPortCreateInput) SetMode(v DedicatedPortCreateInputMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DedicatedPortCreateInput) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *DedicatedPortCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DedicatedPortCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DedicatedPortCreateInput) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *DedicatedPortCreateInput) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *DedicatedPortCreateInput) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *DedicatedPortCreateInput) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *DedicatedPortCreateInput) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRedundancy

`func (o *DedicatedPortCreateInput) GetRedundancy() string`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *DedicatedPortCreateInput) GetRedundancyOk() (*string, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *DedicatedPortCreateInput) SetRedundancy(v string)`

SetRedundancy sets Redundancy field to given value.


### GetSpeed

`func (o *DedicatedPortCreateInput) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *DedicatedPortCreateInput) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *DedicatedPortCreateInput) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *DedicatedPortCreateInput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTags

`func (o *DedicatedPortCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DedicatedPortCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DedicatedPortCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DedicatedPortCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *DedicatedPortCreateInput) GetType() DedicatedPortCreateInputType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DedicatedPortCreateInput) GetTypeOk() (*DedicatedPortCreateInputType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DedicatedPortCreateInput) SetType(v DedicatedPortCreateInputType)`

SetType sets Type field to given value.


### GetUseCase

`func (o *DedicatedPortCreateInput) GetUseCase() string`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *DedicatedPortCreateInput) GetUseCaseOk() (*string, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *DedicatedPortCreateInput) SetUseCase(v string)`

SetUseCase sets UseCase field to given value.

### HasUseCase

`func (o *DedicatedPortCreateInput) HasUseCase() bool`

HasUseCase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


