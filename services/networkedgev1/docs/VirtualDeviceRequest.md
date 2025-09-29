# VirtualDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | Account number. Either an account number or accountReferenceId is required. | [optional] 
**AccountReferenceId** | Pointer to **string** | AccountReferenceId. This is a temporary ID that can be used to create a device when the account status is still pending, not active. Either an account number or accountReferenceId is required. | [optional] 
**ProjectId** | Pointer to **string** | Customer project Id. Check your projectId under Resource Management on Equinix Portal. You should have access to a project to see or create assets under it. Equinix will assign a projectId if you do not provide one. | [optional] 
**Version** | **string** | Version. | 
**DeviceTypeCode** | **string** | Virtual device type (device type code) | 
**HostNamePrefix** | Pointer to **string** | Host name prefix for identification. Only a-z, A-Z, 0-9, and hyphen(-) are allowed. It should start with a letter and end with a letter or digit. The length should be between 2-30 characters. Exceptions - FTDv 2-14 characters; Aruba 2-24 characters. | [optional] 
**AgreeOrderTerms** | **bool** | To create a device, you must accept the order terms. Call Get Order Terms to review the details. If you are creating an Equinix-Configured device, read your vendor&#39;s terms by calling Get Vendor Terms. | 
**LicenseMode** | **string** | License type. One of SUB (Subscription) or BYOL (Bring Your Own License) | 
**LicenseCategory** | Pointer to **string** | This field will be deprecated in the future. | [optional] 
**LicenseFileId** | Pointer to **string** | For Juniper devices you need to provide a licenseFileId if you want to BYOL (Bring Your Own License). You get a licenseFileId when you upload a license file by calling license upload API (Upload a license file before creating a virtual device). For Cisco devices, you do not need to provide a licenseFileId at the time of device creation. Once the device is provisioned, you can get the deviceSerialNo by calling Get virtual device {uuid} API. With the deviceSerialNo you can generate a license file on Cisco site. Afterward, you can upload the license file by calling license upload API (Upload a license file after creating a virtual device). | [optional] 
**LicenseToken** | Pointer to **string** | In case you want to BYOL (Bring Your Own License) for a Palo Alto device, you must provide a license token. This field must have 8 alphanumeric characters. | [optional] 
**Day0TextFileId** | Pointer to **string** | Some devices require a day0TextFileId. Upload your license file by calling Upload File API. You&#39;ll get a fileUuid in the response. You can enter the value in the day0TextFileId field of the create payload to create a virtual device. Check the payloads of individual devices (provided as Postman Scripts on our API doc site) for details. | [optional] 
**Termlength** | Pointer to **string** | Term length in months. | [optional] 
**MetroCode** | **string** | Metro code | 
**PackageCode** | Pointer to **string** | Software package code | [optional] 
**SshUsers** | Pointer to [**[]SshUsers**](SshUsers.md) | An array of sshUsernames and passwords | [optional] 
**Throughput** | Pointer to **int32** | Numeric throughput. For Cisco8KV self-configured devices, you can choose either numeric or tier throughput options. | [optional] 
**ThroughputUnit** | Pointer to **string** | Throughput unit. | [optional] 
**Tier** | Pointer to **int32** | Tier throughput. Relevant for Cisco8KV devices. For Cisco8KV self-configured devices, you can choose either numeric or tier throughput options. Possible values - 0, 1, 2, 3. Default - 2 | [optional] 
**VirtualDeviceName** | **string** | Virtual device name for identification. This should be minimum 3 and maximum 50 characters long. | 
**OrderingContact** | Pointer to **string** |  | [optional] 
**Notifications** | **[]string** | Email addresses for notification. We need a minimum of 1 and no more than 5 email addresses. | 
**AclDetails** | Pointer to [**[]ACLDetails**](ACLDetails.md) | An array of ACLs | [optional] 
**AdditionalBandwidth** | Pointer to **int32** | Secondary additional bandwidth to be configured (in Mbps for HA). Default bandwidth provided is 15 Mbps. | [optional] 
**DeviceManagementType** | **string** | Whether the device is SELF-CONFIGURED or EQUINIX-CONFIGURED | 
**Core** | **int32** |  | 
**InterfaceCount** | Pointer to **int32** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**SystemIpAddress** | Pointer to **string** |  | [optional] 
**VendorConfig** | Pointer to [**VendorConfig**](VendorConfig.md) |  | [optional] 
**UserPublicKey** | Pointer to [**UserPublicKeyRequest**](UserPublicKeyRequest.md) |  | [optional] 
**IpType** | Pointer to **string** | This field is deprecated. The ipType value always defaults to STATIC. | [optional] 
**SshInterfaceId** | Pointer to **string** | You may specify any available interface on the device as the sshInterfaceId. This field is only applicable to self-configured devices. | [optional] 
**SmartLicenseUrl** | Pointer to **string** | License URL. This field is only relevant for Ciso ASAv devices. | [optional] 
**DiverseFromDeviceUuid** | Pointer to **string** | Unique ID of an existing device. Use this field to let Equinix know if you want  your new device to be in a different location from any existing virtual device. This field is only meaningful for single devices. | [optional] 
**ClusterDetails** | Pointer to [**ClusterConfig**](ClusterConfig.md) |  | [optional] 
**PrimaryDeviceUuid** | Pointer to **string** | This field is mandatory if you are using this API to add a secondary device to an existing primary device. | [optional] 
**Connectivity** | Pointer to **string** | Specifies the connectivity on the device. You can have INTERNET-ACCESS, PRIVATE, or INTERNET-ACCESS-WITH-PRVT-MGMT. Private devices don&#39;t have ACLs or bandwidth. | [optional] 
**ChannelPartner** | Pointer to **string** | The name of the channel partner. | [optional] 
**CloudInitFileId** | Pointer to **string** | The Id of a previously uploaded license or cloud_init file. | [optional] 
**PurchaseOrderNumber** | Pointer to **string** | Purchase Order information will be included in your order confirmation email. | [optional] 
**OrderReference** | Pointer to **string** | Enter a short name/number to identify this order on the invoice. | [optional] 
**Secondary** | Pointer to [**VirtualDevicHARequest**](VirtualDevicHARequest.md) |  | [optional] 

## Methods

### NewVirtualDeviceRequest

`func NewVirtualDeviceRequest(version string, deviceTypeCode string, agreeOrderTerms bool, licenseMode string, metroCode string, virtualDeviceName string, notifications []string, deviceManagementType string, core int32, ) *VirtualDeviceRequest`

NewVirtualDeviceRequest instantiates a new VirtualDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDeviceRequestWithDefaults

`func NewVirtualDeviceRequestWithDefaults() *VirtualDeviceRequest`

NewVirtualDeviceRequestWithDefaults instantiates a new VirtualDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *VirtualDeviceRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *VirtualDeviceRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *VirtualDeviceRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *VirtualDeviceRequest) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountReferenceId

`func (o *VirtualDeviceRequest) GetAccountReferenceId() string`

GetAccountReferenceId returns the AccountReferenceId field if non-nil, zero value otherwise.

### GetAccountReferenceIdOk

`func (o *VirtualDeviceRequest) GetAccountReferenceIdOk() (*string, bool)`

GetAccountReferenceIdOk returns a tuple with the AccountReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReferenceId

`func (o *VirtualDeviceRequest) SetAccountReferenceId(v string)`

SetAccountReferenceId sets AccountReferenceId field to given value.

### HasAccountReferenceId

`func (o *VirtualDeviceRequest) HasAccountReferenceId() bool`

HasAccountReferenceId returns a boolean if a field has been set.

### GetProjectId

`func (o *VirtualDeviceRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *VirtualDeviceRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *VirtualDeviceRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *VirtualDeviceRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetVersion

`func (o *VirtualDeviceRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualDeviceRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualDeviceRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDeviceTypeCode

`func (o *VirtualDeviceRequest) GetDeviceTypeCode() string`

GetDeviceTypeCode returns the DeviceTypeCode field if non-nil, zero value otherwise.

### GetDeviceTypeCodeOk

`func (o *VirtualDeviceRequest) GetDeviceTypeCodeOk() (*string, bool)`

GetDeviceTypeCodeOk returns a tuple with the DeviceTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeCode

`func (o *VirtualDeviceRequest) SetDeviceTypeCode(v string)`

SetDeviceTypeCode sets DeviceTypeCode field to given value.


### GetHostNamePrefix

`func (o *VirtualDeviceRequest) GetHostNamePrefix() string`

GetHostNamePrefix returns the HostNamePrefix field if non-nil, zero value otherwise.

### GetHostNamePrefixOk

`func (o *VirtualDeviceRequest) GetHostNamePrefixOk() (*string, bool)`

GetHostNamePrefixOk returns a tuple with the HostNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamePrefix

`func (o *VirtualDeviceRequest) SetHostNamePrefix(v string)`

SetHostNamePrefix sets HostNamePrefix field to given value.

### HasHostNamePrefix

`func (o *VirtualDeviceRequest) HasHostNamePrefix() bool`

HasHostNamePrefix returns a boolean if a field has been set.

### GetAgreeOrderTerms

`func (o *VirtualDeviceRequest) GetAgreeOrderTerms() bool`

GetAgreeOrderTerms returns the AgreeOrderTerms field if non-nil, zero value otherwise.

### GetAgreeOrderTermsOk

`func (o *VirtualDeviceRequest) GetAgreeOrderTermsOk() (*bool, bool)`

GetAgreeOrderTermsOk returns a tuple with the AgreeOrderTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreeOrderTerms

`func (o *VirtualDeviceRequest) SetAgreeOrderTerms(v bool)`

SetAgreeOrderTerms sets AgreeOrderTerms field to given value.


### GetLicenseMode

`func (o *VirtualDeviceRequest) GetLicenseMode() string`

GetLicenseMode returns the LicenseMode field if non-nil, zero value otherwise.

### GetLicenseModeOk

`func (o *VirtualDeviceRequest) GetLicenseModeOk() (*string, bool)`

GetLicenseModeOk returns a tuple with the LicenseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseMode

`func (o *VirtualDeviceRequest) SetLicenseMode(v string)`

SetLicenseMode sets LicenseMode field to given value.


### GetLicenseCategory

`func (o *VirtualDeviceRequest) GetLicenseCategory() string`

GetLicenseCategory returns the LicenseCategory field if non-nil, zero value otherwise.

### GetLicenseCategoryOk

`func (o *VirtualDeviceRequest) GetLicenseCategoryOk() (*string, bool)`

GetLicenseCategoryOk returns a tuple with the LicenseCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCategory

`func (o *VirtualDeviceRequest) SetLicenseCategory(v string)`

SetLicenseCategory sets LicenseCategory field to given value.

### HasLicenseCategory

`func (o *VirtualDeviceRequest) HasLicenseCategory() bool`

HasLicenseCategory returns a boolean if a field has been set.

### GetLicenseFileId

`func (o *VirtualDeviceRequest) GetLicenseFileId() string`

GetLicenseFileId returns the LicenseFileId field if non-nil, zero value otherwise.

### GetLicenseFileIdOk

`func (o *VirtualDeviceRequest) GetLicenseFileIdOk() (*string, bool)`

GetLicenseFileIdOk returns a tuple with the LicenseFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseFileId

`func (o *VirtualDeviceRequest) SetLicenseFileId(v string)`

SetLicenseFileId sets LicenseFileId field to given value.

### HasLicenseFileId

`func (o *VirtualDeviceRequest) HasLicenseFileId() bool`

HasLicenseFileId returns a boolean if a field has been set.

### GetLicenseToken

`func (o *VirtualDeviceRequest) GetLicenseToken() string`

GetLicenseToken returns the LicenseToken field if non-nil, zero value otherwise.

### GetLicenseTokenOk

`func (o *VirtualDeviceRequest) GetLicenseTokenOk() (*string, bool)`

GetLicenseTokenOk returns a tuple with the LicenseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseToken

`func (o *VirtualDeviceRequest) SetLicenseToken(v string)`

SetLicenseToken sets LicenseToken field to given value.

### HasLicenseToken

`func (o *VirtualDeviceRequest) HasLicenseToken() bool`

HasLicenseToken returns a boolean if a field has been set.

### GetDay0TextFileId

`func (o *VirtualDeviceRequest) GetDay0TextFileId() string`

GetDay0TextFileId returns the Day0TextFileId field if non-nil, zero value otherwise.

### GetDay0TextFileIdOk

`func (o *VirtualDeviceRequest) GetDay0TextFileIdOk() (*string, bool)`

GetDay0TextFileIdOk returns a tuple with the Day0TextFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay0TextFileId

`func (o *VirtualDeviceRequest) SetDay0TextFileId(v string)`

SetDay0TextFileId sets Day0TextFileId field to given value.

### HasDay0TextFileId

`func (o *VirtualDeviceRequest) HasDay0TextFileId() bool`

HasDay0TextFileId returns a boolean if a field has been set.

### GetTermlength

`func (o *VirtualDeviceRequest) GetTermlength() string`

GetTermlength returns the Termlength field if non-nil, zero value otherwise.

### GetTermlengthOk

`func (o *VirtualDeviceRequest) GetTermlengthOk() (*string, bool)`

GetTermlengthOk returns a tuple with the Termlength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermlength

`func (o *VirtualDeviceRequest) SetTermlength(v string)`

SetTermlength sets Termlength field to given value.

### HasTermlength

`func (o *VirtualDeviceRequest) HasTermlength() bool`

HasTermlength returns a boolean if a field has been set.

### GetMetroCode

`func (o *VirtualDeviceRequest) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *VirtualDeviceRequest) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *VirtualDeviceRequest) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.


### GetPackageCode

`func (o *VirtualDeviceRequest) GetPackageCode() string`

GetPackageCode returns the PackageCode field if non-nil, zero value otherwise.

### GetPackageCodeOk

`func (o *VirtualDeviceRequest) GetPackageCodeOk() (*string, bool)`

GetPackageCodeOk returns a tuple with the PackageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCode

`func (o *VirtualDeviceRequest) SetPackageCode(v string)`

SetPackageCode sets PackageCode field to given value.

### HasPackageCode

`func (o *VirtualDeviceRequest) HasPackageCode() bool`

HasPackageCode returns a boolean if a field has been set.

### GetSshUsers

`func (o *VirtualDeviceRequest) GetSshUsers() []SshUsers`

GetSshUsers returns the SshUsers field if non-nil, zero value otherwise.

### GetSshUsersOk

`func (o *VirtualDeviceRequest) GetSshUsersOk() (*[]SshUsers, bool)`

GetSshUsersOk returns a tuple with the SshUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsers

`func (o *VirtualDeviceRequest) SetSshUsers(v []SshUsers)`

SetSshUsers sets SshUsers field to given value.

### HasSshUsers

`func (o *VirtualDeviceRequest) HasSshUsers() bool`

HasSshUsers returns a boolean if a field has been set.

### GetThroughput

`func (o *VirtualDeviceRequest) GetThroughput() int32`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *VirtualDeviceRequest) GetThroughputOk() (*int32, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *VirtualDeviceRequest) SetThroughput(v int32)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *VirtualDeviceRequest) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetThroughputUnit

`func (o *VirtualDeviceRequest) GetThroughputUnit() string`

GetThroughputUnit returns the ThroughputUnit field if non-nil, zero value otherwise.

### GetThroughputUnitOk

`func (o *VirtualDeviceRequest) GetThroughputUnitOk() (*string, bool)`

GetThroughputUnitOk returns a tuple with the ThroughputUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputUnit

`func (o *VirtualDeviceRequest) SetThroughputUnit(v string)`

SetThroughputUnit sets ThroughputUnit field to given value.

### HasThroughputUnit

`func (o *VirtualDeviceRequest) HasThroughputUnit() bool`

HasThroughputUnit returns a boolean if a field has been set.

### GetTier

`func (o *VirtualDeviceRequest) GetTier() int32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *VirtualDeviceRequest) GetTierOk() (*int32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *VirtualDeviceRequest) SetTier(v int32)`

SetTier sets Tier field to given value.

### HasTier

`func (o *VirtualDeviceRequest) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetVirtualDeviceName

`func (o *VirtualDeviceRequest) GetVirtualDeviceName() string`

GetVirtualDeviceName returns the VirtualDeviceName field if non-nil, zero value otherwise.

### GetVirtualDeviceNameOk

`func (o *VirtualDeviceRequest) GetVirtualDeviceNameOk() (*string, bool)`

GetVirtualDeviceNameOk returns a tuple with the VirtualDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceName

`func (o *VirtualDeviceRequest) SetVirtualDeviceName(v string)`

SetVirtualDeviceName sets VirtualDeviceName field to given value.


### GetOrderingContact

`func (o *VirtualDeviceRequest) GetOrderingContact() string`

GetOrderingContact returns the OrderingContact field if non-nil, zero value otherwise.

### GetOrderingContactOk

`func (o *VirtualDeviceRequest) GetOrderingContactOk() (*string, bool)`

GetOrderingContactOk returns a tuple with the OrderingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderingContact

`func (o *VirtualDeviceRequest) SetOrderingContact(v string)`

SetOrderingContact sets OrderingContact field to given value.

### HasOrderingContact

`func (o *VirtualDeviceRequest) HasOrderingContact() bool`

HasOrderingContact returns a boolean if a field has been set.

### GetNotifications

`func (o *VirtualDeviceRequest) GetNotifications() []string`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *VirtualDeviceRequest) GetNotificationsOk() (*[]string, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *VirtualDeviceRequest) SetNotifications(v []string)`

SetNotifications sets Notifications field to given value.


### GetAclDetails

`func (o *VirtualDeviceRequest) GetAclDetails() []ACLDetails`

GetAclDetails returns the AclDetails field if non-nil, zero value otherwise.

### GetAclDetailsOk

`func (o *VirtualDeviceRequest) GetAclDetailsOk() (*[]ACLDetails, bool)`

GetAclDetailsOk returns a tuple with the AclDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclDetails

`func (o *VirtualDeviceRequest) SetAclDetails(v []ACLDetails)`

SetAclDetails sets AclDetails field to given value.

### HasAclDetails

`func (o *VirtualDeviceRequest) HasAclDetails() bool`

HasAclDetails returns a boolean if a field has been set.

### GetAdditionalBandwidth

`func (o *VirtualDeviceRequest) GetAdditionalBandwidth() int32`

GetAdditionalBandwidth returns the AdditionalBandwidth field if non-nil, zero value otherwise.

### GetAdditionalBandwidthOk

`func (o *VirtualDeviceRequest) GetAdditionalBandwidthOk() (*int32, bool)`

GetAdditionalBandwidthOk returns a tuple with the AdditionalBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalBandwidth

`func (o *VirtualDeviceRequest) SetAdditionalBandwidth(v int32)`

SetAdditionalBandwidth sets AdditionalBandwidth field to given value.

### HasAdditionalBandwidth

`func (o *VirtualDeviceRequest) HasAdditionalBandwidth() bool`

HasAdditionalBandwidth returns a boolean if a field has been set.

### GetDeviceManagementType

`func (o *VirtualDeviceRequest) GetDeviceManagementType() string`

GetDeviceManagementType returns the DeviceManagementType field if non-nil, zero value otherwise.

### GetDeviceManagementTypeOk

`func (o *VirtualDeviceRequest) GetDeviceManagementTypeOk() (*string, bool)`

GetDeviceManagementTypeOk returns a tuple with the DeviceManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceManagementType

`func (o *VirtualDeviceRequest) SetDeviceManagementType(v string)`

SetDeviceManagementType sets DeviceManagementType field to given value.


### GetCore

`func (o *VirtualDeviceRequest) GetCore() int32`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *VirtualDeviceRequest) GetCoreOk() (*int32, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *VirtualDeviceRequest) SetCore(v int32)`

SetCore sets Core field to given value.


### GetInterfaceCount

`func (o *VirtualDeviceRequest) GetInterfaceCount() int32`

GetInterfaceCount returns the InterfaceCount field if non-nil, zero value otherwise.

### GetInterfaceCountOk

`func (o *VirtualDeviceRequest) GetInterfaceCountOk() (*int32, bool)`

GetInterfaceCountOk returns a tuple with the InterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceCount

`func (o *VirtualDeviceRequest) SetInterfaceCount(v int32)`

SetInterfaceCount sets InterfaceCount field to given value.

### HasInterfaceCount

`func (o *VirtualDeviceRequest) HasInterfaceCount() bool`

HasInterfaceCount returns a boolean if a field has been set.

### GetSiteId

`func (o *VirtualDeviceRequest) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VirtualDeviceRequest) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VirtualDeviceRequest) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *VirtualDeviceRequest) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSystemIpAddress

`func (o *VirtualDeviceRequest) GetSystemIpAddress() string`

GetSystemIpAddress returns the SystemIpAddress field if non-nil, zero value otherwise.

### GetSystemIpAddressOk

`func (o *VirtualDeviceRequest) GetSystemIpAddressOk() (*string, bool)`

GetSystemIpAddressOk returns a tuple with the SystemIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIpAddress

`func (o *VirtualDeviceRequest) SetSystemIpAddress(v string)`

SetSystemIpAddress sets SystemIpAddress field to given value.

### HasSystemIpAddress

`func (o *VirtualDeviceRequest) HasSystemIpAddress() bool`

HasSystemIpAddress returns a boolean if a field has been set.

### GetVendorConfig

`func (o *VirtualDeviceRequest) GetVendorConfig() VendorConfig`

GetVendorConfig returns the VendorConfig field if non-nil, zero value otherwise.

### GetVendorConfigOk

`func (o *VirtualDeviceRequest) GetVendorConfigOk() (*VendorConfig, bool)`

GetVendorConfigOk returns a tuple with the VendorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorConfig

`func (o *VirtualDeviceRequest) SetVendorConfig(v VendorConfig)`

SetVendorConfig sets VendorConfig field to given value.

### HasVendorConfig

`func (o *VirtualDeviceRequest) HasVendorConfig() bool`

HasVendorConfig returns a boolean if a field has been set.

### GetUserPublicKey

`func (o *VirtualDeviceRequest) GetUserPublicKey() UserPublicKeyRequest`

GetUserPublicKey returns the UserPublicKey field if non-nil, zero value otherwise.

### GetUserPublicKeyOk

`func (o *VirtualDeviceRequest) GetUserPublicKeyOk() (*UserPublicKeyRequest, bool)`

GetUserPublicKeyOk returns a tuple with the UserPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPublicKey

`func (o *VirtualDeviceRequest) SetUserPublicKey(v UserPublicKeyRequest)`

SetUserPublicKey sets UserPublicKey field to given value.

### HasUserPublicKey

`func (o *VirtualDeviceRequest) HasUserPublicKey() bool`

HasUserPublicKey returns a boolean if a field has been set.

### GetIpType

`func (o *VirtualDeviceRequest) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *VirtualDeviceRequest) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *VirtualDeviceRequest) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *VirtualDeviceRequest) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetSshInterfaceId

`func (o *VirtualDeviceRequest) GetSshInterfaceId() string`

GetSshInterfaceId returns the SshInterfaceId field if non-nil, zero value otherwise.

### GetSshInterfaceIdOk

`func (o *VirtualDeviceRequest) GetSshInterfaceIdOk() (*string, bool)`

GetSshInterfaceIdOk returns a tuple with the SshInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshInterfaceId

`func (o *VirtualDeviceRequest) SetSshInterfaceId(v string)`

SetSshInterfaceId sets SshInterfaceId field to given value.

### HasSshInterfaceId

`func (o *VirtualDeviceRequest) HasSshInterfaceId() bool`

HasSshInterfaceId returns a boolean if a field has been set.

### GetSmartLicenseUrl

`func (o *VirtualDeviceRequest) GetSmartLicenseUrl() string`

GetSmartLicenseUrl returns the SmartLicenseUrl field if non-nil, zero value otherwise.

### GetSmartLicenseUrlOk

`func (o *VirtualDeviceRequest) GetSmartLicenseUrlOk() (*string, bool)`

GetSmartLicenseUrlOk returns a tuple with the SmartLicenseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartLicenseUrl

`func (o *VirtualDeviceRequest) SetSmartLicenseUrl(v string)`

SetSmartLicenseUrl sets SmartLicenseUrl field to given value.

### HasSmartLicenseUrl

`func (o *VirtualDeviceRequest) HasSmartLicenseUrl() bool`

HasSmartLicenseUrl returns a boolean if a field has been set.

### GetDiverseFromDeviceUuid

`func (o *VirtualDeviceRequest) GetDiverseFromDeviceUuid() string`

GetDiverseFromDeviceUuid returns the DiverseFromDeviceUuid field if non-nil, zero value otherwise.

### GetDiverseFromDeviceUuidOk

`func (o *VirtualDeviceRequest) GetDiverseFromDeviceUuidOk() (*string, bool)`

GetDiverseFromDeviceUuidOk returns a tuple with the DiverseFromDeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiverseFromDeviceUuid

`func (o *VirtualDeviceRequest) SetDiverseFromDeviceUuid(v string)`

SetDiverseFromDeviceUuid sets DiverseFromDeviceUuid field to given value.

### HasDiverseFromDeviceUuid

`func (o *VirtualDeviceRequest) HasDiverseFromDeviceUuid() bool`

HasDiverseFromDeviceUuid returns a boolean if a field has been set.

### GetClusterDetails

`func (o *VirtualDeviceRequest) GetClusterDetails() ClusterConfig`

GetClusterDetails returns the ClusterDetails field if non-nil, zero value otherwise.

### GetClusterDetailsOk

`func (o *VirtualDeviceRequest) GetClusterDetailsOk() (*ClusterConfig, bool)`

GetClusterDetailsOk returns a tuple with the ClusterDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDetails

`func (o *VirtualDeviceRequest) SetClusterDetails(v ClusterConfig)`

SetClusterDetails sets ClusterDetails field to given value.

### HasClusterDetails

`func (o *VirtualDeviceRequest) HasClusterDetails() bool`

HasClusterDetails returns a boolean if a field has been set.

### GetPrimaryDeviceUuid

`func (o *VirtualDeviceRequest) GetPrimaryDeviceUuid() string`

GetPrimaryDeviceUuid returns the PrimaryDeviceUuid field if non-nil, zero value otherwise.

### GetPrimaryDeviceUuidOk

`func (o *VirtualDeviceRequest) GetPrimaryDeviceUuidOk() (*string, bool)`

GetPrimaryDeviceUuidOk returns a tuple with the PrimaryDeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDeviceUuid

`func (o *VirtualDeviceRequest) SetPrimaryDeviceUuid(v string)`

SetPrimaryDeviceUuid sets PrimaryDeviceUuid field to given value.

### HasPrimaryDeviceUuid

`func (o *VirtualDeviceRequest) HasPrimaryDeviceUuid() bool`

HasPrimaryDeviceUuid returns a boolean if a field has been set.

### GetConnectivity

`func (o *VirtualDeviceRequest) GetConnectivity() string`

GetConnectivity returns the Connectivity field if non-nil, zero value otherwise.

### GetConnectivityOk

`func (o *VirtualDeviceRequest) GetConnectivityOk() (*string, bool)`

GetConnectivityOk returns a tuple with the Connectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivity

`func (o *VirtualDeviceRequest) SetConnectivity(v string)`

SetConnectivity sets Connectivity field to given value.

### HasConnectivity

`func (o *VirtualDeviceRequest) HasConnectivity() bool`

HasConnectivity returns a boolean if a field has been set.

### GetChannelPartner

`func (o *VirtualDeviceRequest) GetChannelPartner() string`

GetChannelPartner returns the ChannelPartner field if non-nil, zero value otherwise.

### GetChannelPartnerOk

`func (o *VirtualDeviceRequest) GetChannelPartnerOk() (*string, bool)`

GetChannelPartnerOk returns a tuple with the ChannelPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPartner

`func (o *VirtualDeviceRequest) SetChannelPartner(v string)`

SetChannelPartner sets ChannelPartner field to given value.

### HasChannelPartner

`func (o *VirtualDeviceRequest) HasChannelPartner() bool`

HasChannelPartner returns a boolean if a field has been set.

### GetCloudInitFileId

`func (o *VirtualDeviceRequest) GetCloudInitFileId() string`

GetCloudInitFileId returns the CloudInitFileId field if non-nil, zero value otherwise.

### GetCloudInitFileIdOk

`func (o *VirtualDeviceRequest) GetCloudInitFileIdOk() (*string, bool)`

GetCloudInitFileIdOk returns a tuple with the CloudInitFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitFileId

`func (o *VirtualDeviceRequest) SetCloudInitFileId(v string)`

SetCloudInitFileId sets CloudInitFileId field to given value.

### HasCloudInitFileId

`func (o *VirtualDeviceRequest) HasCloudInitFileId() bool`

HasCloudInitFileId returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *VirtualDeviceRequest) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *VirtualDeviceRequest) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *VirtualDeviceRequest) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *VirtualDeviceRequest) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetOrderReference

`func (o *VirtualDeviceRequest) GetOrderReference() string`

GetOrderReference returns the OrderReference field if non-nil, zero value otherwise.

### GetOrderReferenceOk

`func (o *VirtualDeviceRequest) GetOrderReferenceOk() (*string, bool)`

GetOrderReferenceOk returns a tuple with the OrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReference

`func (o *VirtualDeviceRequest) SetOrderReference(v string)`

SetOrderReference sets OrderReference field to given value.

### HasOrderReference

`func (o *VirtualDeviceRequest) HasOrderReference() bool`

HasOrderReference returns a boolean if a field has been set.

### GetSecondary

`func (o *VirtualDeviceRequest) GetSecondary() VirtualDevicHARequest`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *VirtualDeviceRequest) GetSecondaryOk() (*VirtualDevicHARequest, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *VirtualDeviceRequest) SetSecondary(v VirtualDevicHARequest)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *VirtualDeviceRequest) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


