# SoftwarePackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Software package name | [optional] 
**PackageCode** | Pointer to **string** | Software package code | [optional] 
**LicenseType** | Pointer to **string** | Software package license type | [optional] 
**VersionDetails** | Pointer to [**[]VersionDetails**](VersionDetails.md) |  | [optional] 

## Methods

### NewSoftwarePackage

`func NewSoftwarePackage() *SoftwarePackage`

NewSoftwarePackage instantiates a new SoftwarePackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarePackageWithDefaults

`func NewSoftwarePackageWithDefaults() *SoftwarePackage`

NewSoftwarePackageWithDefaults instantiates a new SoftwarePackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SoftwarePackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwarePackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwarePackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoftwarePackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackageCode

`func (o *SoftwarePackage) GetPackageCode() string`

GetPackageCode returns the PackageCode field if non-nil, zero value otherwise.

### GetPackageCodeOk

`func (o *SoftwarePackage) GetPackageCodeOk() (*string, bool)`

GetPackageCodeOk returns a tuple with the PackageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCode

`func (o *SoftwarePackage) SetPackageCode(v string)`

SetPackageCode sets PackageCode field to given value.

### HasPackageCode

`func (o *SoftwarePackage) HasPackageCode() bool`

HasPackageCode returns a boolean if a field has been set.

### GetLicenseType

`func (o *SoftwarePackage) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *SoftwarePackage) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *SoftwarePackage) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *SoftwarePackage) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetVersionDetails

`func (o *SoftwarePackage) GetVersionDetails() []VersionDetails`

GetVersionDetails returns the VersionDetails field if non-nil, zero value otherwise.

### GetVersionDetailsOk

`func (o *SoftwarePackage) GetVersionDetailsOk() (*[]VersionDetails, bool)`

GetVersionDetailsOk returns a tuple with the VersionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionDetails

`func (o *SoftwarePackage) SetVersionDetails(v []VersionDetails)`

SetVersionDetails sets VersionDetails field to given value.

### HasVersionDetails

`func (o *SoftwarePackage) HasVersionDetails() bool`

HasVersionDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


