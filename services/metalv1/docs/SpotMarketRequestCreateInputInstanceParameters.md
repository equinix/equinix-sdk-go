# SpotMarketRequestCreateInputInstanceParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysPxe** | Pointer to **bool** |  | [optional] 
**BillingCycle** | Pointer to **string** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Hostnames** | Pointer to **[]string** |  | [optional] 
**Locked** | Pointer to **bool** | Whether the device should be locked, preventing accidental deletion. | [optional] 
**NoSshKeys** | Pointer to **bool** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**PrivateIpv4SubnetSize** | Pointer to **int32** |  | [optional] 
**ProjectSshKeys** | Pointer to **[]string** |  | [optional] 
**PublicIpv4SubnetSize** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TerminationTime** | Pointer to **time.Time** |  | [optional] 
**UserSshKeys** | Pointer to **[]string** | The UUIDs of users whose SSH keys should be included on the provisioned device. | [optional] 
**Userdata** | Pointer to **string** |  | [optional] 

## Methods

### NewSpotMarketRequestCreateInputInstanceParameters

`func NewSpotMarketRequestCreateInputInstanceParameters() *SpotMarketRequestCreateInputInstanceParameters`

NewSpotMarketRequestCreateInputInstanceParameters instantiates a new SpotMarketRequestCreateInputInstanceParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotMarketRequestCreateInputInstanceParametersWithDefaults

`func NewSpotMarketRequestCreateInputInstanceParametersWithDefaults() *SpotMarketRequestCreateInputInstanceParameters`

NewSpotMarketRequestCreateInputInstanceParametersWithDefaults instantiates a new SpotMarketRequestCreateInputInstanceParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysPxe

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetAlwaysPxe() bool`

GetAlwaysPxe returns the AlwaysPxe field if non-nil, zero value otherwise.

### GetAlwaysPxeOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetAlwaysPxeOk() (*bool, bool)`

GetAlwaysPxeOk returns a tuple with the AlwaysPxe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysPxe

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetAlwaysPxe(v bool)`

SetAlwaysPxe sets AlwaysPxe field to given value.

### HasAlwaysPxe

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasAlwaysPxe() bool`

HasAlwaysPxe returns a boolean if a field has been set.

### GetBillingCycle

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetCustomdata

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDescription

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetHostname

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetHostnames

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetLocked

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetNoSshKeys

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetNoSshKeys() bool`

GetNoSshKeys returns the NoSshKeys field if non-nil, zero value otherwise.

### GetNoSshKeysOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetNoSshKeysOk() (*bool, bool)`

GetNoSshKeysOk returns a tuple with the NoSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSshKeys

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetNoSshKeys(v bool)`

SetNoSshKeys sets NoSshKeys field to given value.

### HasNoSshKeys

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasNoSshKeys() bool`

HasNoSshKeys returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetPlan

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPrivateIpv4SubnetSize

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetPrivateIpv4SubnetSize() int32`

GetPrivateIpv4SubnetSize returns the PrivateIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPrivateIpv4SubnetSizeOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetPrivateIpv4SubnetSizeOk() (*int32, bool)`

GetPrivateIpv4SubnetSizeOk returns a tuple with the PrivateIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpv4SubnetSize

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetPrivateIpv4SubnetSize(v int32)`

SetPrivateIpv4SubnetSize sets PrivateIpv4SubnetSize field to given value.

### HasPrivateIpv4SubnetSize

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasPrivateIpv4SubnetSize() bool`

HasPrivateIpv4SubnetSize returns a boolean if a field has been set.

### GetProjectSshKeys

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetProjectSshKeys() []string`

GetProjectSshKeys returns the ProjectSshKeys field if non-nil, zero value otherwise.

### GetProjectSshKeysOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetProjectSshKeysOk() (*[]string, bool)`

GetProjectSshKeysOk returns a tuple with the ProjectSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSshKeys

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetProjectSshKeys(v []string)`

SetProjectSshKeys sets ProjectSshKeys field to given value.

### HasProjectSshKeys

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasProjectSshKeys() bool`

HasProjectSshKeys returns a boolean if a field has been set.

### GetPublicIpv4SubnetSize

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetPublicIpv4SubnetSize() int32`

GetPublicIpv4SubnetSize returns the PublicIpv4SubnetSize field if non-nil, zero value otherwise.

### GetPublicIpv4SubnetSizeOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetPublicIpv4SubnetSizeOk() (*int32, bool)`

GetPublicIpv4SubnetSizeOk returns a tuple with the PublicIpv4SubnetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv4SubnetSize

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetPublicIpv4SubnetSize(v int32)`

SetPublicIpv4SubnetSize sets PublicIpv4SubnetSize field to given value.

### HasPublicIpv4SubnetSize

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasPublicIpv4SubnetSize() bool`

HasPublicIpv4SubnetSize returns a boolean if a field has been set.

### GetTags

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationTime

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.

### GetUserSshKeys

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetUserSshKeys() []string`

GetUserSshKeys returns the UserSshKeys field if non-nil, zero value otherwise.

### GetUserSshKeysOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetUserSshKeysOk() (*[]string, bool)`

GetUserSshKeysOk returns a tuple with the UserSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSshKeys

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetUserSshKeys(v []string)`

SetUserSshKeys sets UserSshKeys field to given value.

### HasUserSshKeys

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasUserSshKeys() bool`

HasUserSshKeys returns a boolean if a field has been set.

### GetUserdata

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *SpotMarketRequestCreateInputInstanceParameters) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *SpotMarketRequestCreateInputInstanceParameters) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *SpotMarketRequestCreateInputInstanceParameters) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


