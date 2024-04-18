# SupportedServicesConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of supported service. | [optional] 
**Required** | Pointer to **bool** | Whether or not this supported service is a required input at the time of device creation. | [optional] 
**PackageCodes** | Pointer to **[]string** |  | [optional] 
**SupportedForClustering** | Pointer to **bool** | Whether the service is available for cluster devices. | [optional] 

## Methods

### NewSupportedServicesConfig

`func NewSupportedServicesConfig() *SupportedServicesConfig`

NewSupportedServicesConfig instantiates a new SupportedServicesConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedServicesConfigWithDefaults

`func NewSupportedServicesConfigWithDefaults() *SupportedServicesConfig`

NewSupportedServicesConfigWithDefaults instantiates a new SupportedServicesConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SupportedServicesConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupportedServicesConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupportedServicesConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SupportedServicesConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *SupportedServicesConfig) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SupportedServicesConfig) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SupportedServicesConfig) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *SupportedServicesConfig) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetPackageCodes

`func (o *SupportedServicesConfig) GetPackageCodes() []string`

GetPackageCodes returns the PackageCodes field if non-nil, zero value otherwise.

### GetPackageCodesOk

`func (o *SupportedServicesConfig) GetPackageCodesOk() (*[]string, bool)`

GetPackageCodesOk returns a tuple with the PackageCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCodes

`func (o *SupportedServicesConfig) SetPackageCodes(v []string)`

SetPackageCodes sets PackageCodes field to given value.

### HasPackageCodes

`func (o *SupportedServicesConfig) HasPackageCodes() bool`

HasPackageCodes returns a boolean if a field has been set.

### GetSupportedForClustering

`func (o *SupportedServicesConfig) GetSupportedForClustering() bool`

GetSupportedForClustering returns the SupportedForClustering field if non-nil, zero value otherwise.

### GetSupportedForClusteringOk

`func (o *SupportedServicesConfig) GetSupportedForClusteringOk() (*bool, bool)`

GetSupportedForClusteringOk returns a tuple with the SupportedForClustering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedForClustering

`func (o *SupportedServicesConfig) SetSupportedForClustering(v bool)`

SetSupportedForClustering sets SupportedForClustering field to given value.

### HasSupportedForClustering

`func (o *SupportedServicesConfig) HasSupportedForClustering() bool`

HasSupportedForClustering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


