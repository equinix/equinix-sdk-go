# PackageCodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageCode** | Pointer to **string** | The type of package. | [optional] 
**ExcludedVersions** | Pointer to **[]string** |  | [optional] 
**ExcludedClusterVersions** | Pointer to **[]string** |  | [optional] 
**Throughputs** | Pointer to [**[]ThroughputConfig**](ThroughputConfig.md) |  | [optional] 
**Supported** | Pointer to **bool** | Whether this software package is supported. | [optional] 

## Methods

### NewPackageCodes

`func NewPackageCodes() *PackageCodes`

NewPackageCodes instantiates a new PackageCodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageCodesWithDefaults

`func NewPackageCodesWithDefaults() *PackageCodes`

NewPackageCodesWithDefaults instantiates a new PackageCodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageCode

`func (o *PackageCodes) GetPackageCode() string`

GetPackageCode returns the PackageCode field if non-nil, zero value otherwise.

### GetPackageCodeOk

`func (o *PackageCodes) GetPackageCodeOk() (*string, bool)`

GetPackageCodeOk returns a tuple with the PackageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCode

`func (o *PackageCodes) SetPackageCode(v string)`

SetPackageCode sets PackageCode field to given value.

### HasPackageCode

`func (o *PackageCodes) HasPackageCode() bool`

HasPackageCode returns a boolean if a field has been set.

### GetExcludedVersions

`func (o *PackageCodes) GetExcludedVersions() []string`

GetExcludedVersions returns the ExcludedVersions field if non-nil, zero value otherwise.

### GetExcludedVersionsOk

`func (o *PackageCodes) GetExcludedVersionsOk() (*[]string, bool)`

GetExcludedVersionsOk returns a tuple with the ExcludedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedVersions

`func (o *PackageCodes) SetExcludedVersions(v []string)`

SetExcludedVersions sets ExcludedVersions field to given value.

### HasExcludedVersions

`func (o *PackageCodes) HasExcludedVersions() bool`

HasExcludedVersions returns a boolean if a field has been set.

### GetExcludedClusterVersions

`func (o *PackageCodes) GetExcludedClusterVersions() []string`

GetExcludedClusterVersions returns the ExcludedClusterVersions field if non-nil, zero value otherwise.

### GetExcludedClusterVersionsOk

`func (o *PackageCodes) GetExcludedClusterVersionsOk() (*[]string, bool)`

GetExcludedClusterVersionsOk returns a tuple with the ExcludedClusterVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedClusterVersions

`func (o *PackageCodes) SetExcludedClusterVersions(v []string)`

SetExcludedClusterVersions sets ExcludedClusterVersions field to given value.

### HasExcludedClusterVersions

`func (o *PackageCodes) HasExcludedClusterVersions() bool`

HasExcludedClusterVersions returns a boolean if a field has been set.

### GetThroughputs

`func (o *PackageCodes) GetThroughputs() []ThroughputConfig`

GetThroughputs returns the Throughputs field if non-nil, zero value otherwise.

### GetThroughputsOk

`func (o *PackageCodes) GetThroughputsOk() (*[]ThroughputConfig, bool)`

GetThroughputsOk returns a tuple with the Throughputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputs

`func (o *PackageCodes) SetThroughputs(v []ThroughputConfig)`

SetThroughputs sets Throughputs field to given value.

### HasThroughputs

`func (o *PackageCodes) HasThroughputs() bool`

HasThroughputs returns a boolean if a field has been set.

### GetSupported

`func (o *PackageCodes) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *PackageCodes) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *PackageCodes) SetSupported(v bool)`

SetSupported sets Supported field to given value.

### HasSupported

`func (o *PackageCodes) HasSupported() bool`

HasSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


