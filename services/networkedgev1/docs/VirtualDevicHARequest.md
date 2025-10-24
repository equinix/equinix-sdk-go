# VirtualDevicHARequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** |  | [optional] 
**AccountReferenceId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** | You can only choose a version for the secondary device when adding a secondary device to an existing device. | [optional] 
**AdditionalBandwidth** | Pointer to **int32** | Secondary additional bandwidth to be configured (in Mbps for HA). Default bandwidth provided is 15 Mbps. | [optional] 
**LicenseFileId** | Pointer to **string** |  | [optional] 
**LicenseToken** | Pointer to **string** |  | [optional] 
**Day0TextFileId** | Pointer to **string** | Some devices require a day0TextFileId. Upload your license file by calling Upload File API. You&#39;ll get a fileUuid in the response. You can enter the value in the day0TextFileId field of the create payload to create a virtual device. Check the payloads of individual devices (provided as Postman Scripts on our API doc site) for details. | [optional] 
**MetroCode** | **string** |  | 
**Notifications** | [**[]VirtualDevicHARequestNotificationsInner**](VirtualDevicHARequestNotificationsInner.md) |  | 
**AclDetails** | Pointer to [**[]ACLDetails**](ACLDetails.md) | An array of ACLs | [optional] 
**SshUsers** | Pointer to [**[]SshUserOperationRequest**](SshUserOperationRequest.md) |  | [optional] 
**VirtualDeviceName** | **string** | Virtual Device Name | 
**HostNamePrefix** | Pointer to **string** | Host name prefix for identification. Only a-z, A-Z, 0-9, and hyphen(-) are allowed. It should start with a letter and end with a letter or digit. The length should be between 2-30 characters. Exceptions - FTDv 2-14 characters; Aruba 2-24 characters. | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**SystemIpAddress** | Pointer to **string** |  | [optional] 
**VendorConfig** | Pointer to [**VendorConfig**](VendorConfig.md) |  | [optional] 
**SshInterfaceId** | Pointer to **string** | You may specify any available interface on the device as the sshInterfaceId. This field is only applicable to self-configured devices. | [optional] 
**SmartLicenseUrl** | Pointer to **string** | License URL. This field is only relevant for Ciso ASAv devices. | [optional] 
**CloudInitFileId** | Pointer to **string** | The Id of a previously uploaded license or cloud_init file. | [optional] 

## Methods

### NewVirtualDevicHARequest

`func NewVirtualDevicHARequest(metroCode string, notifications []VirtualDevicHARequestNotificationsInner, virtualDeviceName string, ) *VirtualDevicHARequest`

NewVirtualDevicHARequest instantiates a new VirtualDevicHARequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDevicHARequestWithDefaults

`func NewVirtualDevicHARequestWithDefaults() *VirtualDevicHARequest`

NewVirtualDevicHARequestWithDefaults instantiates a new VirtualDevicHARequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *VirtualDevicHARequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *VirtualDevicHARequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *VirtualDevicHARequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *VirtualDevicHARequest) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountReferenceId

`func (o *VirtualDevicHARequest) GetAccountReferenceId() string`

GetAccountReferenceId returns the AccountReferenceId field if non-nil, zero value otherwise.

### GetAccountReferenceIdOk

`func (o *VirtualDevicHARequest) GetAccountReferenceIdOk() (*string, bool)`

GetAccountReferenceIdOk returns a tuple with the AccountReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReferenceId

`func (o *VirtualDevicHARequest) SetAccountReferenceId(v string)`

SetAccountReferenceId sets AccountReferenceId field to given value.

### HasAccountReferenceId

`func (o *VirtualDevicHARequest) HasAccountReferenceId() bool`

HasAccountReferenceId returns a boolean if a field has been set.

### GetVersion

`func (o *VirtualDevicHARequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualDevicHARequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualDevicHARequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VirtualDevicHARequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAdditionalBandwidth

`func (o *VirtualDevicHARequest) GetAdditionalBandwidth() int32`

GetAdditionalBandwidth returns the AdditionalBandwidth field if non-nil, zero value otherwise.

### GetAdditionalBandwidthOk

`func (o *VirtualDevicHARequest) GetAdditionalBandwidthOk() (*int32, bool)`

GetAdditionalBandwidthOk returns a tuple with the AdditionalBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalBandwidth

`func (o *VirtualDevicHARequest) SetAdditionalBandwidth(v int32)`

SetAdditionalBandwidth sets AdditionalBandwidth field to given value.

### HasAdditionalBandwidth

`func (o *VirtualDevicHARequest) HasAdditionalBandwidth() bool`

HasAdditionalBandwidth returns a boolean if a field has been set.

### GetLicenseFileId

`func (o *VirtualDevicHARequest) GetLicenseFileId() string`

GetLicenseFileId returns the LicenseFileId field if non-nil, zero value otherwise.

### GetLicenseFileIdOk

`func (o *VirtualDevicHARequest) GetLicenseFileIdOk() (*string, bool)`

GetLicenseFileIdOk returns a tuple with the LicenseFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseFileId

`func (o *VirtualDevicHARequest) SetLicenseFileId(v string)`

SetLicenseFileId sets LicenseFileId field to given value.

### HasLicenseFileId

`func (o *VirtualDevicHARequest) HasLicenseFileId() bool`

HasLicenseFileId returns a boolean if a field has been set.

### GetLicenseToken

`func (o *VirtualDevicHARequest) GetLicenseToken() string`

GetLicenseToken returns the LicenseToken field if non-nil, zero value otherwise.

### GetLicenseTokenOk

`func (o *VirtualDevicHARequest) GetLicenseTokenOk() (*string, bool)`

GetLicenseTokenOk returns a tuple with the LicenseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseToken

`func (o *VirtualDevicHARequest) SetLicenseToken(v string)`

SetLicenseToken sets LicenseToken field to given value.

### HasLicenseToken

`func (o *VirtualDevicHARequest) HasLicenseToken() bool`

HasLicenseToken returns a boolean if a field has been set.

### GetDay0TextFileId

`func (o *VirtualDevicHARequest) GetDay0TextFileId() string`

GetDay0TextFileId returns the Day0TextFileId field if non-nil, zero value otherwise.

### GetDay0TextFileIdOk

`func (o *VirtualDevicHARequest) GetDay0TextFileIdOk() (*string, bool)`

GetDay0TextFileIdOk returns a tuple with the Day0TextFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay0TextFileId

`func (o *VirtualDevicHARequest) SetDay0TextFileId(v string)`

SetDay0TextFileId sets Day0TextFileId field to given value.

### HasDay0TextFileId

`func (o *VirtualDevicHARequest) HasDay0TextFileId() bool`

HasDay0TextFileId returns a boolean if a field has been set.

### GetMetroCode

`func (o *VirtualDevicHARequest) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *VirtualDevicHARequest) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *VirtualDevicHARequest) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.


### GetNotifications

`func (o *VirtualDevicHARequest) GetNotifications() []VirtualDevicHARequestNotificationsInner`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *VirtualDevicHARequest) GetNotificationsOk() (*[]VirtualDevicHARequestNotificationsInner, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *VirtualDevicHARequest) SetNotifications(v []VirtualDevicHARequestNotificationsInner)`

SetNotifications sets Notifications field to given value.


### GetAclDetails

`func (o *VirtualDevicHARequest) GetAclDetails() []ACLDetails`

GetAclDetails returns the AclDetails field if non-nil, zero value otherwise.

### GetAclDetailsOk

`func (o *VirtualDevicHARequest) GetAclDetailsOk() (*[]ACLDetails, bool)`

GetAclDetailsOk returns a tuple with the AclDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclDetails

`func (o *VirtualDevicHARequest) SetAclDetails(v []ACLDetails)`

SetAclDetails sets AclDetails field to given value.

### HasAclDetails

`func (o *VirtualDevicHARequest) HasAclDetails() bool`

HasAclDetails returns a boolean if a field has been set.

### GetSshUsers

`func (o *VirtualDevicHARequest) GetSshUsers() []SshUserOperationRequest`

GetSshUsers returns the SshUsers field if non-nil, zero value otherwise.

### GetSshUsersOk

`func (o *VirtualDevicHARequest) GetSshUsersOk() (*[]SshUserOperationRequest, bool)`

GetSshUsersOk returns a tuple with the SshUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsers

`func (o *VirtualDevicHARequest) SetSshUsers(v []SshUserOperationRequest)`

SetSshUsers sets SshUsers field to given value.

### HasSshUsers

`func (o *VirtualDevicHARequest) HasSshUsers() bool`

HasSshUsers returns a boolean if a field has been set.

### GetVirtualDeviceName

`func (o *VirtualDevicHARequest) GetVirtualDeviceName() string`

GetVirtualDeviceName returns the VirtualDeviceName field if non-nil, zero value otherwise.

### GetVirtualDeviceNameOk

`func (o *VirtualDevicHARequest) GetVirtualDeviceNameOk() (*string, bool)`

GetVirtualDeviceNameOk returns a tuple with the VirtualDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceName

`func (o *VirtualDevicHARequest) SetVirtualDeviceName(v string)`

SetVirtualDeviceName sets VirtualDeviceName field to given value.


### GetHostNamePrefix

`func (o *VirtualDevicHARequest) GetHostNamePrefix() string`

GetHostNamePrefix returns the HostNamePrefix field if non-nil, zero value otherwise.

### GetHostNamePrefixOk

`func (o *VirtualDevicHARequest) GetHostNamePrefixOk() (*string, bool)`

GetHostNamePrefixOk returns a tuple with the HostNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamePrefix

`func (o *VirtualDevicHARequest) SetHostNamePrefix(v string)`

SetHostNamePrefix sets HostNamePrefix field to given value.

### HasHostNamePrefix

`func (o *VirtualDevicHARequest) HasHostNamePrefix() bool`

HasHostNamePrefix returns a boolean if a field has been set.

### GetSiteId

`func (o *VirtualDevicHARequest) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VirtualDevicHARequest) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VirtualDevicHARequest) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *VirtualDevicHARequest) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSystemIpAddress

`func (o *VirtualDevicHARequest) GetSystemIpAddress() string`

GetSystemIpAddress returns the SystemIpAddress field if non-nil, zero value otherwise.

### GetSystemIpAddressOk

`func (o *VirtualDevicHARequest) GetSystemIpAddressOk() (*string, bool)`

GetSystemIpAddressOk returns a tuple with the SystemIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIpAddress

`func (o *VirtualDevicHARequest) SetSystemIpAddress(v string)`

SetSystemIpAddress sets SystemIpAddress field to given value.

### HasSystemIpAddress

`func (o *VirtualDevicHARequest) HasSystemIpAddress() bool`

HasSystemIpAddress returns a boolean if a field has been set.

### GetVendorConfig

`func (o *VirtualDevicHARequest) GetVendorConfig() VendorConfig`

GetVendorConfig returns the VendorConfig field if non-nil, zero value otherwise.

### GetVendorConfigOk

`func (o *VirtualDevicHARequest) GetVendorConfigOk() (*VendorConfig, bool)`

GetVendorConfigOk returns a tuple with the VendorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorConfig

`func (o *VirtualDevicHARequest) SetVendorConfig(v VendorConfig)`

SetVendorConfig sets VendorConfig field to given value.

### HasVendorConfig

`func (o *VirtualDevicHARequest) HasVendorConfig() bool`

HasVendorConfig returns a boolean if a field has been set.

### GetSshInterfaceId

`func (o *VirtualDevicHARequest) GetSshInterfaceId() string`

GetSshInterfaceId returns the SshInterfaceId field if non-nil, zero value otherwise.

### GetSshInterfaceIdOk

`func (o *VirtualDevicHARequest) GetSshInterfaceIdOk() (*string, bool)`

GetSshInterfaceIdOk returns a tuple with the SshInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshInterfaceId

`func (o *VirtualDevicHARequest) SetSshInterfaceId(v string)`

SetSshInterfaceId sets SshInterfaceId field to given value.

### HasSshInterfaceId

`func (o *VirtualDevicHARequest) HasSshInterfaceId() bool`

HasSshInterfaceId returns a boolean if a field has been set.

### GetSmartLicenseUrl

`func (o *VirtualDevicHARequest) GetSmartLicenseUrl() string`

GetSmartLicenseUrl returns the SmartLicenseUrl field if non-nil, zero value otherwise.

### GetSmartLicenseUrlOk

`func (o *VirtualDevicHARequest) GetSmartLicenseUrlOk() (*string, bool)`

GetSmartLicenseUrlOk returns a tuple with the SmartLicenseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartLicenseUrl

`func (o *VirtualDevicHARequest) SetSmartLicenseUrl(v string)`

SetSmartLicenseUrl sets SmartLicenseUrl field to given value.

### HasSmartLicenseUrl

`func (o *VirtualDevicHARequest) HasSmartLicenseUrl() bool`

HasSmartLicenseUrl returns a boolean if a field has been set.

### GetCloudInitFileId

`func (o *VirtualDevicHARequest) GetCloudInitFileId() string`

GetCloudInitFileId returns the CloudInitFileId field if non-nil, zero value otherwise.

### GetCloudInitFileIdOk

`func (o *VirtualDevicHARequest) GetCloudInitFileIdOk() (*string, bool)`

GetCloudInitFileIdOk returns a tuple with the CloudInitFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitFileId

`func (o *VirtualDevicHARequest) SetCloudInitFileId(v string)`

SetCloudInitFileId sets CloudInitFileId field to given value.

### HasCloudInitFileId

`func (o *VirtualDevicHARequest) HasCloudInitFileId() bool`

HasCloudInitFileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


