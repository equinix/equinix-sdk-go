# EquinixConfiguredConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Whether the device is EQUINIX-CONFIGURED or SELF-MANAGED. | [optional] 
**LicenseOptions** | Pointer to [**EquinixConfiguredConfigLicenseOptions**](EquinixConfiguredConfigLicenseOptions.md) |  | [optional] 
**SupportedServices** | Pointer to [**[]SupportedServicesConfig**](SupportedServicesConfig.md) |  | [optional] 
**AdditionalFields** | Pointer to [**[]AdditionalFieldsConfig**](AdditionalFieldsConfig.md) |  | [optional] 
**ClusteringDetails** | Pointer to [**ClusteringDetails**](ClusteringDetails.md) |  | [optional] 

## Methods

### NewEquinixConfiguredConfig

`func NewEquinixConfiguredConfig() *EquinixConfiguredConfig`

NewEquinixConfiguredConfig instantiates a new EquinixConfiguredConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquinixConfiguredConfigWithDefaults

`func NewEquinixConfiguredConfigWithDefaults() *EquinixConfiguredConfig`

NewEquinixConfiguredConfigWithDefaults instantiates a new EquinixConfiguredConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EquinixConfiguredConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EquinixConfiguredConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EquinixConfiguredConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EquinixConfiguredConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLicenseOptions

`func (o *EquinixConfiguredConfig) GetLicenseOptions() EquinixConfiguredConfigLicenseOptions`

GetLicenseOptions returns the LicenseOptions field if non-nil, zero value otherwise.

### GetLicenseOptionsOk

`func (o *EquinixConfiguredConfig) GetLicenseOptionsOk() (*EquinixConfiguredConfigLicenseOptions, bool)`

GetLicenseOptionsOk returns a tuple with the LicenseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseOptions

`func (o *EquinixConfiguredConfig) SetLicenseOptions(v EquinixConfiguredConfigLicenseOptions)`

SetLicenseOptions sets LicenseOptions field to given value.

### HasLicenseOptions

`func (o *EquinixConfiguredConfig) HasLicenseOptions() bool`

HasLicenseOptions returns a boolean if a field has been set.

### GetSupportedServices

`func (o *EquinixConfiguredConfig) GetSupportedServices() []SupportedServicesConfig`

GetSupportedServices returns the SupportedServices field if non-nil, zero value otherwise.

### GetSupportedServicesOk

`func (o *EquinixConfiguredConfig) GetSupportedServicesOk() (*[]SupportedServicesConfig, bool)`

GetSupportedServicesOk returns a tuple with the SupportedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedServices

`func (o *EquinixConfiguredConfig) SetSupportedServices(v []SupportedServicesConfig)`

SetSupportedServices sets SupportedServices field to given value.

### HasSupportedServices

`func (o *EquinixConfiguredConfig) HasSupportedServices() bool`

HasSupportedServices returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *EquinixConfiguredConfig) GetAdditionalFields() []AdditionalFieldsConfig`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *EquinixConfiguredConfig) GetAdditionalFieldsOk() (*[]AdditionalFieldsConfig, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *EquinixConfiguredConfig) SetAdditionalFields(v []AdditionalFieldsConfig)`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *EquinixConfiguredConfig) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetClusteringDetails

`func (o *EquinixConfiguredConfig) GetClusteringDetails() ClusteringDetails`

GetClusteringDetails returns the ClusteringDetails field if non-nil, zero value otherwise.

### GetClusteringDetailsOk

`func (o *EquinixConfiguredConfig) GetClusteringDetailsOk() (*ClusteringDetails, bool)`

GetClusteringDetailsOk returns a tuple with the ClusteringDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusteringDetails

`func (o *EquinixConfiguredConfig) SetClusteringDetails(v ClusteringDetails)`

SetClusteringDetails sets ClusteringDetails field to given value.

### HasClusteringDetails

`func (o *EquinixConfiguredConfig) HasClusteringDetails() bool`

HasClusteringDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


