# SelfConfiguredConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Whether the device is EQUINIX-CONFIGURED or SELF-MANAGED. | [optional] 
**LicenseOptions** | Pointer to [**SelfConfiguredConfigLicenseOptions**](SelfConfiguredConfigLicenseOptions.md) |  | [optional] 
**SupportedServices** | Pointer to [**[]SupportedServicesConfig**](SupportedServicesConfig.md) |  | [optional] 
**AdditionalFields** | Pointer to [**[]AdditionalFieldsConfig**](AdditionalFieldsConfig.md) |  | [optional] 
**DefaultAcls** | Pointer to [**DefaultAclsConfig**](DefaultAclsConfig.md) |  | [optional] 
**ClusteringDetails** | Pointer to [**ClusteringDetails**](ClusteringDetails.md) |  | [optional] 

## Methods

### NewSelfConfiguredConfig

`func NewSelfConfiguredConfig() *SelfConfiguredConfig`

NewSelfConfiguredConfig instantiates a new SelfConfiguredConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfConfiguredConfigWithDefaults

`func NewSelfConfiguredConfigWithDefaults() *SelfConfiguredConfig`

NewSelfConfiguredConfigWithDefaults instantiates a new SelfConfiguredConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SelfConfiguredConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelfConfiguredConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelfConfiguredConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SelfConfiguredConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLicenseOptions

`func (o *SelfConfiguredConfig) GetLicenseOptions() SelfConfiguredConfigLicenseOptions`

GetLicenseOptions returns the LicenseOptions field if non-nil, zero value otherwise.

### GetLicenseOptionsOk

`func (o *SelfConfiguredConfig) GetLicenseOptionsOk() (*SelfConfiguredConfigLicenseOptions, bool)`

GetLicenseOptionsOk returns a tuple with the LicenseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseOptions

`func (o *SelfConfiguredConfig) SetLicenseOptions(v SelfConfiguredConfigLicenseOptions)`

SetLicenseOptions sets LicenseOptions field to given value.

### HasLicenseOptions

`func (o *SelfConfiguredConfig) HasLicenseOptions() bool`

HasLicenseOptions returns a boolean if a field has been set.

### GetSupportedServices

`func (o *SelfConfiguredConfig) GetSupportedServices() []SupportedServicesConfig`

GetSupportedServices returns the SupportedServices field if non-nil, zero value otherwise.

### GetSupportedServicesOk

`func (o *SelfConfiguredConfig) GetSupportedServicesOk() (*[]SupportedServicesConfig, bool)`

GetSupportedServicesOk returns a tuple with the SupportedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedServices

`func (o *SelfConfiguredConfig) SetSupportedServices(v []SupportedServicesConfig)`

SetSupportedServices sets SupportedServices field to given value.

### HasSupportedServices

`func (o *SelfConfiguredConfig) HasSupportedServices() bool`

HasSupportedServices returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *SelfConfiguredConfig) GetAdditionalFields() []AdditionalFieldsConfig`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *SelfConfiguredConfig) GetAdditionalFieldsOk() (*[]AdditionalFieldsConfig, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *SelfConfiguredConfig) SetAdditionalFields(v []AdditionalFieldsConfig)`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *SelfConfiguredConfig) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetDefaultAcls

`func (o *SelfConfiguredConfig) GetDefaultAcls() DefaultAclsConfig`

GetDefaultAcls returns the DefaultAcls field if non-nil, zero value otherwise.

### GetDefaultAclsOk

`func (o *SelfConfiguredConfig) GetDefaultAclsOk() (*DefaultAclsConfig, bool)`

GetDefaultAclsOk returns a tuple with the DefaultAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAcls

`func (o *SelfConfiguredConfig) SetDefaultAcls(v DefaultAclsConfig)`

SetDefaultAcls sets DefaultAcls field to given value.

### HasDefaultAcls

`func (o *SelfConfiguredConfig) HasDefaultAcls() bool`

HasDefaultAcls returns a boolean if a field has been set.

### GetClusteringDetails

`func (o *SelfConfiguredConfig) GetClusteringDetails() ClusteringDetails`

GetClusteringDetails returns the ClusteringDetails field if non-nil, zero value otherwise.

### GetClusteringDetailsOk

`func (o *SelfConfiguredConfig) GetClusteringDetailsOk() (*ClusteringDetails, bool)`

GetClusteringDetailsOk returns a tuple with the ClusteringDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusteringDetails

`func (o *SelfConfiguredConfig) SetClusteringDetails(v ClusteringDetails)`

SetClusteringDetails sets ClusteringDetails field to given value.

### HasClusteringDetails

`func (o *SelfConfiguredConfig) HasClusteringDetails() bool`

HasClusteringDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


