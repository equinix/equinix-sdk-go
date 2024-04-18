# RMAVendorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **string** | Physical location within the Viptela overlay network, such as a branch office, or a campus (relevant only for Cisco SD-WANs) | [optional] 
**SystemIpAddress** | Pointer to **string** | IP assigned to CSRSDWAN router and vSmart controller (relevant only for Cisco SD-WANs) | [optional] 
**LicenseKey** | Pointer to **string** | License key. Mandatory for some devices. | [optional] 
**LicenseSecret** | Pointer to **string** | License secret (Secret key). Mandatory for some devices. | [optional] 
**Controller1** | Pointer to **string** | For Fortinet devices, this is the System IP address. | [optional] 
**AdminPassword** | Pointer to **string** | The administrative password of the device. You can use it to log in to the console. This field is not available for all device types. Should be at least 6 characters long and must include an uppercase letter and a number. This field may be required for some vendors. | [optional] 
**ActivationKey** | Pointer to **string** | Available on VMware Orchestration Portal | [optional] 
**ProvisioningKey** | Pointer to **string** | Mandatory for Zscaler devices | [optional] 

## Methods

### NewRMAVendorConfig

`func NewRMAVendorConfig() *RMAVendorConfig`

NewRMAVendorConfig instantiates a new RMAVendorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRMAVendorConfigWithDefaults

`func NewRMAVendorConfigWithDefaults() *RMAVendorConfig`

NewRMAVendorConfigWithDefaults instantiates a new RMAVendorConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *RMAVendorConfig) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *RMAVendorConfig) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *RMAVendorConfig) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *RMAVendorConfig) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSystemIpAddress

`func (o *RMAVendorConfig) GetSystemIpAddress() string`

GetSystemIpAddress returns the SystemIpAddress field if non-nil, zero value otherwise.

### GetSystemIpAddressOk

`func (o *RMAVendorConfig) GetSystemIpAddressOk() (*string, bool)`

GetSystemIpAddressOk returns a tuple with the SystemIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIpAddress

`func (o *RMAVendorConfig) SetSystemIpAddress(v string)`

SetSystemIpAddress sets SystemIpAddress field to given value.

### HasSystemIpAddress

`func (o *RMAVendorConfig) HasSystemIpAddress() bool`

HasSystemIpAddress returns a boolean if a field has been set.

### GetLicenseKey

`func (o *RMAVendorConfig) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *RMAVendorConfig) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *RMAVendorConfig) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *RMAVendorConfig) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetLicenseSecret

`func (o *RMAVendorConfig) GetLicenseSecret() string`

GetLicenseSecret returns the LicenseSecret field if non-nil, zero value otherwise.

### GetLicenseSecretOk

`func (o *RMAVendorConfig) GetLicenseSecretOk() (*string, bool)`

GetLicenseSecretOk returns a tuple with the LicenseSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseSecret

`func (o *RMAVendorConfig) SetLicenseSecret(v string)`

SetLicenseSecret sets LicenseSecret field to given value.

### HasLicenseSecret

`func (o *RMAVendorConfig) HasLicenseSecret() bool`

HasLicenseSecret returns a boolean if a field has been set.

### GetController1

`func (o *RMAVendorConfig) GetController1() string`

GetController1 returns the Controller1 field if non-nil, zero value otherwise.

### GetController1Ok

`func (o *RMAVendorConfig) GetController1Ok() (*string, bool)`

GetController1Ok returns a tuple with the Controller1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController1

`func (o *RMAVendorConfig) SetController1(v string)`

SetController1 sets Controller1 field to given value.

### HasController1

`func (o *RMAVendorConfig) HasController1() bool`

HasController1 returns a boolean if a field has been set.

### GetAdminPassword

`func (o *RMAVendorConfig) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *RMAVendorConfig) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *RMAVendorConfig) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### HasAdminPassword

`func (o *RMAVendorConfig) HasAdminPassword() bool`

HasAdminPassword returns a boolean if a field has been set.

### GetActivationKey

`func (o *RMAVendorConfig) GetActivationKey() string`

GetActivationKey returns the ActivationKey field if non-nil, zero value otherwise.

### GetActivationKeyOk

`func (o *RMAVendorConfig) GetActivationKeyOk() (*string, bool)`

GetActivationKeyOk returns a tuple with the ActivationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationKey

`func (o *RMAVendorConfig) SetActivationKey(v string)`

SetActivationKey sets ActivationKey field to given value.

### HasActivationKey

`func (o *RMAVendorConfig) HasActivationKey() bool`

HasActivationKey returns a boolean if a field has been set.

### GetProvisioningKey

`func (o *RMAVendorConfig) GetProvisioningKey() string`

GetProvisioningKey returns the ProvisioningKey field if non-nil, zero value otherwise.

### GetProvisioningKeyOk

`func (o *RMAVendorConfig) GetProvisioningKeyOk() (*string, bool)`

GetProvisioningKeyOk returns a tuple with the ProvisioningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningKey

`func (o *RMAVendorConfig) SetProvisioningKey(v string)`

SetProvisioningKey sets ProvisioningKey field to given value.

### HasProvisioningKey

`func (o *RMAVendorConfig) HasProvisioningKey() bool`

HasProvisioningKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


