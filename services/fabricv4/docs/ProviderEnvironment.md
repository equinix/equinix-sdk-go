# ProviderEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Provider Environment URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned provider environment identifier | [optional] 
**Type** | Pointer to [**ProviderEnvironmentTypeEnum**](ProviderEnvironmentTypeEnum.md) |  | [optional] 
**Name** | Pointer to **string** | Provider environment name | [optional] 
**Description** | Pointer to **string** | Provider environment description | [optional] 
**Region** | Pointer to **string** | Cloud provider region identifier | [optional] 
**SupportedBandwidths** | Pointer to **[]int32** | Supported bandwidths in Mbps | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewProviderEnvironment

`func NewProviderEnvironment() *ProviderEnvironment`

NewProviderEnvironment instantiates a new ProviderEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderEnvironmentWithDefaults

`func NewProviderEnvironmentWithDefaults() *ProviderEnvironment`

NewProviderEnvironmentWithDefaults instantiates a new ProviderEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ProviderEnvironment) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ProviderEnvironment) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ProviderEnvironment) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ProviderEnvironment) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *ProviderEnvironment) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProviderEnvironment) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProviderEnvironment) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProviderEnvironment) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *ProviderEnvironment) GetType() ProviderEnvironmentTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProviderEnvironment) GetTypeOk() (*ProviderEnvironmentTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProviderEnvironment) SetType(v ProviderEnvironmentTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *ProviderEnvironment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ProviderEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProviderEnvironment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProviderEnvironment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProviderEnvironment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProviderEnvironment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRegion

`func (o *ProviderEnvironment) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ProviderEnvironment) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ProviderEnvironment) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ProviderEnvironment) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSupportedBandwidths

`func (o *ProviderEnvironment) GetSupportedBandwidths() []int32`

GetSupportedBandwidths returns the SupportedBandwidths field if non-nil, zero value otherwise.

### GetSupportedBandwidthsOk

`func (o *ProviderEnvironment) GetSupportedBandwidthsOk() (*[]int32, bool)`

GetSupportedBandwidthsOk returns a tuple with the SupportedBandwidths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedBandwidths

`func (o *ProviderEnvironment) SetSupportedBandwidths(v []int32)`

SetSupportedBandwidths sets SupportedBandwidths field to given value.

### HasSupportedBandwidths

`func (o *ProviderEnvironment) HasSupportedBandwidths() bool`

HasSupportedBandwidths returns a boolean if a field has been set.

### GetChangeLog

`func (o *ProviderEnvironment) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *ProviderEnvironment) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *ProviderEnvironment) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *ProviderEnvironment) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


