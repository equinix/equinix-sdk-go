# VersionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cause** | Pointer to [**VersionDetails**](VersionDetails.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**VersionDate** | Pointer to **string** | The date the software was released | [optional] 
**RetireDate** | Pointer to **string** | The date the software will no longer be available for new devices. This field will not show if the software does not have a retire date. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StableVersion** | Pointer to **string** |  | [optional] 
**AllowedUpgradableVersions** | Pointer to **[]string** |  | [optional] 
**SupportedLicenseTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVersionDetails

`func NewVersionDetails() *VersionDetails`

NewVersionDetails instantiates a new VersionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionDetailsWithDefaults

`func NewVersionDetailsWithDefaults() *VersionDetails`

NewVersionDetailsWithDefaults instantiates a new VersionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCause

`func (o *VersionDetails) GetCause() VersionDetails`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *VersionDetails) GetCauseOk() (*VersionDetails, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *VersionDetails) SetCause(v VersionDetails)`

SetCause sets Cause field to given value.

### HasCause

`func (o *VersionDetails) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetVersion

`func (o *VersionDetails) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionDetails) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionDetails) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VersionDetails) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetImageName

`func (o *VersionDetails) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *VersionDetails) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *VersionDetails) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *VersionDetails) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetVersionDate

`func (o *VersionDetails) GetVersionDate() string`

GetVersionDate returns the VersionDate field if non-nil, zero value otherwise.

### GetVersionDateOk

`func (o *VersionDetails) GetVersionDateOk() (*string, bool)`

GetVersionDateOk returns a tuple with the VersionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionDate

`func (o *VersionDetails) SetVersionDate(v string)`

SetVersionDate sets VersionDate field to given value.

### HasVersionDate

`func (o *VersionDetails) HasVersionDate() bool`

HasVersionDate returns a boolean if a field has been set.

### GetRetireDate

`func (o *VersionDetails) GetRetireDate() string`

GetRetireDate returns the RetireDate field if non-nil, zero value otherwise.

### GetRetireDateOk

`func (o *VersionDetails) GetRetireDateOk() (*string, bool)`

GetRetireDateOk returns a tuple with the RetireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetireDate

`func (o *VersionDetails) SetRetireDate(v string)`

SetRetireDate sets RetireDate field to given value.

### HasRetireDate

`func (o *VersionDetails) HasRetireDate() bool`

HasRetireDate returns a boolean if a field has been set.

### GetStatus

`func (o *VersionDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VersionDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VersionDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VersionDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStableVersion

`func (o *VersionDetails) GetStableVersion() string`

GetStableVersion returns the StableVersion field if non-nil, zero value otherwise.

### GetStableVersionOk

`func (o *VersionDetails) GetStableVersionOk() (*string, bool)`

GetStableVersionOk returns a tuple with the StableVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableVersion

`func (o *VersionDetails) SetStableVersion(v string)`

SetStableVersion sets StableVersion field to given value.

### HasStableVersion

`func (o *VersionDetails) HasStableVersion() bool`

HasStableVersion returns a boolean if a field has been set.

### GetAllowedUpgradableVersions

`func (o *VersionDetails) GetAllowedUpgradableVersions() []string`

GetAllowedUpgradableVersions returns the AllowedUpgradableVersions field if non-nil, zero value otherwise.

### GetAllowedUpgradableVersionsOk

`func (o *VersionDetails) GetAllowedUpgradableVersionsOk() (*[]string, bool)`

GetAllowedUpgradableVersionsOk returns a tuple with the AllowedUpgradableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUpgradableVersions

`func (o *VersionDetails) SetAllowedUpgradableVersions(v []string)`

SetAllowedUpgradableVersions sets AllowedUpgradableVersions field to given value.

### HasAllowedUpgradableVersions

`func (o *VersionDetails) HasAllowedUpgradableVersions() bool`

HasAllowedUpgradableVersions returns a boolean if a field has been set.

### GetSupportedLicenseTypes

`func (o *VersionDetails) GetSupportedLicenseTypes() []string`

GetSupportedLicenseTypes returns the SupportedLicenseTypes field if non-nil, zero value otherwise.

### GetSupportedLicenseTypesOk

`func (o *VersionDetails) GetSupportedLicenseTypesOk() (*[]string, bool)`

GetSupportedLicenseTypesOk returns a tuple with the SupportedLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedLicenseTypes

`func (o *VersionDetails) SetSupportedLicenseTypes(v []string)`

SetSupportedLicenseTypes sets SupportedLicenseTypes field to given value.

### HasSupportedLicenseTypes

`func (o *VersionDetails) HasSupportedLicenseTypes() bool`

HasSupportedLicenseTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


