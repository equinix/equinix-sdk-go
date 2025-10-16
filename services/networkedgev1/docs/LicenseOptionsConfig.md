# LicenseOptionsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the license. | [optional] 
**Name** | Pointer to **string** | The name of the license. | [optional] 
**FileUploadSupportedCluster** | Pointer to **bool** | Whether you can upload a license file for cluster devices. | [optional] 
**Cores** | Pointer to [**[]CoresConfig**](CoresConfig.md) |  | [optional] 

## Methods

### NewLicenseOptionsConfig

`func NewLicenseOptionsConfig() *LicenseOptionsConfig`

NewLicenseOptionsConfig instantiates a new LicenseOptionsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseOptionsConfigWithDefaults

`func NewLicenseOptionsConfigWithDefaults() *LicenseOptionsConfig`

NewLicenseOptionsConfigWithDefaults instantiates a new LicenseOptionsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LicenseOptionsConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LicenseOptionsConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LicenseOptionsConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LicenseOptionsConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *LicenseOptionsConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LicenseOptionsConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LicenseOptionsConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LicenseOptionsConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFileUploadSupportedCluster

`func (o *LicenseOptionsConfig) GetFileUploadSupportedCluster() bool`

GetFileUploadSupportedCluster returns the FileUploadSupportedCluster field if non-nil, zero value otherwise.

### GetFileUploadSupportedClusterOk

`func (o *LicenseOptionsConfig) GetFileUploadSupportedClusterOk() (*bool, bool)`

GetFileUploadSupportedClusterOk returns a tuple with the FileUploadSupportedCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadSupportedCluster

`func (o *LicenseOptionsConfig) SetFileUploadSupportedCluster(v bool)`

SetFileUploadSupportedCluster sets FileUploadSupportedCluster field to given value.

### HasFileUploadSupportedCluster

`func (o *LicenseOptionsConfig) HasFileUploadSupportedCluster() bool`

HasFileUploadSupportedCluster returns a boolean if a field has been set.

### GetCores

`func (o *LicenseOptionsConfig) GetCores() []CoresConfig`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *LicenseOptionsConfig) GetCoresOk() (*[]CoresConfig, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *LicenseOptionsConfig) SetCores(v []CoresConfig)`

SetCores sets Cores field to given value.

### HasCores

`func (o *LicenseOptionsConfig) HasCores() bool`

HasCores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


