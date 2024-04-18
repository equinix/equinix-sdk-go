# VirtualDeviceDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **string** |  | [optional] 
**DeviceSerialNo** | Pointer to **string** |  | [optional] 
**DeviceTypeCategory** | Pointer to **string** |  | [optional] 
**DiverseFromDeviceName** | Pointer to **string** | The name of a device that is in a location different from this device. | [optional] 
**DiverseFromDeviceUuid** | Pointer to **string** | The unique ID of a device that is in a location different from this device. | [optional] 
**DeviceTypeCode** | Pointer to **string** |  | [optional] 
**DeviceTypeName** | Pointer to **string** |  | [optional] 
**Expiry** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**DeviceTypeVendor** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] 
**LastUpdatedDate** | Pointer to **string** |  | [optional] 
**LicenseFileId** | Pointer to **string** |  | [optional] 
**LicenseName** | Pointer to **string** |  | [optional] 
**LicenseStatus** | Pointer to **string** |  | [optional] 
**LicenseType** | Pointer to **string** |  | [optional] 
**MetroCode** | Pointer to **string** |  | [optional] 
**MetroName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notifications** | Pointer to **[]string** |  | [optional] 
**PackageCode** | Pointer to **string** |  | [optional] 
**PackageName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**PurchaseOrderNumber** | Pointer to **string** |  | [optional] 
**RedundancyType** | Pointer to **string** |  | [optional] 
**RedundantUuid** | Pointer to **string** |  | [optional] 
**Connectivity** | Pointer to **string** |  | [optional] 
**SshIpAddress** | Pointer to **string** |  | [optional] 
**SshIpFqdn** | Pointer to **string** |  | [optional] 
**SshIpFqdnStatus** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Throughput** | Pointer to **int32** |  | [optional] 
**ThroughputUnit** | Pointer to **string** |  | [optional] 
**Core** | Pointer to [**CoresDisplayConfig**](CoresDisplayConfig.md) |  | [optional] 
**PricingDetails** | Pointer to [**PricingSiebelConfig**](PricingSiebelConfig.md) |  | [optional] 
**InterfaceCount** | Pointer to **int32** |  | [optional] 
**DeviceManagementType** | Pointer to **string** |  | [optional] 
**UserPublicKey** | Pointer to [**UserPublicKeyConfig**](UserPublicKeyConfig.md) |  | [optional] 
**ManagementIp** | Pointer to **string** |  | [optional] 
**ManagementGatewayIp** | Pointer to **string** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**PublicGatewayIp** | Pointer to **string** |  | [optional] 
**PrimaryDnsName** | Pointer to **string** |  | [optional] 
**SecondaryDnsName** | Pointer to **string** |  | [optional] 
**TermLength** | Pointer to **string** |  | [optional] 
**AdditionalBandwidth** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**SystemIpAddress** | Pointer to **string** |  | [optional] 
**VendorConfig** | Pointer to [**VendorConfig**](VendorConfig.md) |  | [optional] 
**Interfaces** | Pointer to [**[]InterfaceBasicInfoResponse**](InterfaceBasicInfoResponse.md) |  | [optional] 
**Asn** | Pointer to **float32** | The ASN number. | [optional] 
**ChannelPartner** | Pointer to **string** | The name of the channel partner. | [optional] 

## Methods

### NewVirtualDeviceDetailsResponse

`func NewVirtualDeviceDetailsResponse() *VirtualDeviceDetailsResponse`

NewVirtualDeviceDetailsResponse instantiates a new VirtualDeviceDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDeviceDetailsResponseWithDefaults

`func NewVirtualDeviceDetailsResponseWithDefaults() *VirtualDeviceDetailsResponse`

NewVirtualDeviceDetailsResponseWithDefaults instantiates a new VirtualDeviceDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *VirtualDeviceDetailsResponse) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *VirtualDeviceDetailsResponse) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *VirtualDeviceDetailsResponse) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *VirtualDeviceDetailsResponse) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountNumber

`func (o *VirtualDeviceDetailsResponse) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *VirtualDeviceDetailsResponse) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *VirtualDeviceDetailsResponse) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *VirtualDeviceDetailsResponse) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VirtualDeviceDetailsResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VirtualDeviceDetailsResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VirtualDeviceDetailsResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VirtualDeviceDetailsResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *VirtualDeviceDetailsResponse) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *VirtualDeviceDetailsResponse) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *VirtualDeviceDetailsResponse) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *VirtualDeviceDetailsResponse) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDeviceSerialNo

`func (o *VirtualDeviceDetailsResponse) GetDeviceSerialNo() string`

GetDeviceSerialNo returns the DeviceSerialNo field if non-nil, zero value otherwise.

### GetDeviceSerialNoOk

`func (o *VirtualDeviceDetailsResponse) GetDeviceSerialNoOk() (*string, bool)`

GetDeviceSerialNoOk returns a tuple with the DeviceSerialNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerialNo

`func (o *VirtualDeviceDetailsResponse) SetDeviceSerialNo(v string)`

SetDeviceSerialNo sets DeviceSerialNo field to given value.

### HasDeviceSerialNo

`func (o *VirtualDeviceDetailsResponse) HasDeviceSerialNo() bool`

HasDeviceSerialNo returns a boolean if a field has been set.

### GetDeviceTypeCategory

`func (o *VirtualDeviceDetailsResponse) GetDeviceTypeCategory() string`

GetDeviceTypeCategory returns the DeviceTypeCategory field if non-nil, zero value otherwise.

### GetDeviceTypeCategoryOk

`func (o *VirtualDeviceDetailsResponse) GetDeviceTypeCategoryOk() (*string, bool)`

GetDeviceTypeCategoryOk returns a tuple with the DeviceTypeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeCategory

`func (o *VirtualDeviceDetailsResponse) SetDeviceTypeCategory(v string)`

SetDeviceTypeCategory sets DeviceTypeCategory field to given value.

### HasDeviceTypeCategory

`func (o *VirtualDeviceDetailsResponse) HasDeviceTypeCategory() bool`

HasDeviceTypeCategory returns a boolean if a field has been set.

### GetDiverseFromDeviceName

`func (o *VirtualDeviceDetailsResponse) GetDiverseFromDeviceName() string`

GetDiverseFromDeviceName returns the DiverseFromDeviceName field if non-nil, zero value otherwise.

### GetDiverseFromDeviceNameOk

`func (o *VirtualDeviceDetailsResponse) GetDiverseFromDeviceNameOk() (*string, bool)`

GetDiverseFromDeviceNameOk returns a tuple with the DiverseFromDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiverseFromDeviceName

`func (o *VirtualDeviceDetailsResponse) SetDiverseFromDeviceName(v string)`

SetDiverseFromDeviceName sets DiverseFromDeviceName field to given value.

### HasDiverseFromDeviceName

`func (o *VirtualDeviceDetailsResponse) HasDiverseFromDeviceName() bool`

HasDiverseFromDeviceName returns a boolean if a field has been set.

### GetDiverseFromDeviceUuid

`func (o *VirtualDeviceDetailsResponse) GetDiverseFromDeviceUuid() string`

GetDiverseFromDeviceUuid returns the DiverseFromDeviceUuid field if non-nil, zero value otherwise.

### GetDiverseFromDeviceUuidOk

`func (o *VirtualDeviceDetailsResponse) GetDiverseFromDeviceUuidOk() (*string, bool)`

GetDiverseFromDeviceUuidOk returns a tuple with the DiverseFromDeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiverseFromDeviceUuid

`func (o *VirtualDeviceDetailsResponse) SetDiverseFromDeviceUuid(v string)`

SetDiverseFromDeviceUuid sets DiverseFromDeviceUuid field to given value.

### HasDiverseFromDeviceUuid

`func (o *VirtualDeviceDetailsResponse) HasDiverseFromDeviceUuid() bool`

HasDiverseFromDeviceUuid returns a boolean if a field has been set.

### GetDeviceTypeCode

`func (o *VirtualDeviceDetailsResponse) GetDeviceTypeCode() string`

GetDeviceTypeCode returns the DeviceTypeCode field if non-nil, zero value otherwise.

### GetDeviceTypeCodeOk

`func (o *VirtualDeviceDetailsResponse) GetDeviceTypeCodeOk() (*string, bool)`

GetDeviceTypeCodeOk returns a tuple with the DeviceTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeCode

`func (o *VirtualDeviceDetailsResponse) SetDeviceTypeCode(v string)`

SetDeviceTypeCode sets DeviceTypeCode field to given value.

### HasDeviceTypeCode

`func (o *VirtualDeviceDetailsResponse) HasDeviceTypeCode() bool`

HasDeviceTypeCode returns a boolean if a field has been set.

### GetDeviceTypeName

`func (o *VirtualDeviceDetailsResponse) GetDeviceTypeName() string`

GetDeviceTypeName returns the DeviceTypeName field if non-nil, zero value otherwise.

### GetDeviceTypeNameOk

`func (o *VirtualDeviceDetailsResponse) GetDeviceTypeNameOk() (*string, bool)`

GetDeviceTypeNameOk returns a tuple with the DeviceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeName

`func (o *VirtualDeviceDetailsResponse) SetDeviceTypeName(v string)`

SetDeviceTypeName sets DeviceTypeName field to given value.

### HasDeviceTypeName

`func (o *VirtualDeviceDetailsResponse) HasDeviceTypeName() bool`

HasDeviceTypeName returns a boolean if a field has been set.

### GetExpiry

`func (o *VirtualDeviceDetailsResponse) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *VirtualDeviceDetailsResponse) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *VirtualDeviceDetailsResponse) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *VirtualDeviceDetailsResponse) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetRegion

`func (o *VirtualDeviceDetailsResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VirtualDeviceDetailsResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VirtualDeviceDetailsResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VirtualDeviceDetailsResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDeviceTypeVendor

`func (o *VirtualDeviceDetailsResponse) GetDeviceTypeVendor() string`

GetDeviceTypeVendor returns the DeviceTypeVendor field if non-nil, zero value otherwise.

### GetDeviceTypeVendorOk

`func (o *VirtualDeviceDetailsResponse) GetDeviceTypeVendorOk() (*string, bool)`

GetDeviceTypeVendorOk returns a tuple with the DeviceTypeVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeVendor

`func (o *VirtualDeviceDetailsResponse) SetDeviceTypeVendor(v string)`

SetDeviceTypeVendor sets DeviceTypeVendor field to given value.

### HasDeviceTypeVendor

`func (o *VirtualDeviceDetailsResponse) HasDeviceTypeVendor() bool`

HasDeviceTypeVendor returns a boolean if a field has been set.

### GetHostName

`func (o *VirtualDeviceDetailsResponse) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *VirtualDeviceDetailsResponse) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *VirtualDeviceDetailsResponse) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *VirtualDeviceDetailsResponse) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetUuid

`func (o *VirtualDeviceDetailsResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualDeviceDetailsResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualDeviceDetailsResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualDeviceDetailsResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *VirtualDeviceDetailsResponse) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *VirtualDeviceDetailsResponse) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *VirtualDeviceDetailsResponse) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *VirtualDeviceDetailsResponse) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *VirtualDeviceDetailsResponse) GetLastUpdatedDate() string`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *VirtualDeviceDetailsResponse) GetLastUpdatedDateOk() (*string, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *VirtualDeviceDetailsResponse) SetLastUpdatedDate(v string)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *VirtualDeviceDetailsResponse) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetLicenseFileId

`func (o *VirtualDeviceDetailsResponse) GetLicenseFileId() string`

GetLicenseFileId returns the LicenseFileId field if non-nil, zero value otherwise.

### GetLicenseFileIdOk

`func (o *VirtualDeviceDetailsResponse) GetLicenseFileIdOk() (*string, bool)`

GetLicenseFileIdOk returns a tuple with the LicenseFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseFileId

`func (o *VirtualDeviceDetailsResponse) SetLicenseFileId(v string)`

SetLicenseFileId sets LicenseFileId field to given value.

### HasLicenseFileId

`func (o *VirtualDeviceDetailsResponse) HasLicenseFileId() bool`

HasLicenseFileId returns a boolean if a field has been set.

### GetLicenseName

`func (o *VirtualDeviceDetailsResponse) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *VirtualDeviceDetailsResponse) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *VirtualDeviceDetailsResponse) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *VirtualDeviceDetailsResponse) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *VirtualDeviceDetailsResponse) GetLicenseStatus() string`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *VirtualDeviceDetailsResponse) GetLicenseStatusOk() (*string, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *VirtualDeviceDetailsResponse) SetLicenseStatus(v string)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *VirtualDeviceDetailsResponse) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetLicenseType

`func (o *VirtualDeviceDetailsResponse) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *VirtualDeviceDetailsResponse) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *VirtualDeviceDetailsResponse) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *VirtualDeviceDetailsResponse) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetMetroCode

`func (o *VirtualDeviceDetailsResponse) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *VirtualDeviceDetailsResponse) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *VirtualDeviceDetailsResponse) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.

### HasMetroCode

`func (o *VirtualDeviceDetailsResponse) HasMetroCode() bool`

HasMetroCode returns a boolean if a field has been set.

### GetMetroName

`func (o *VirtualDeviceDetailsResponse) GetMetroName() string`

GetMetroName returns the MetroName field if non-nil, zero value otherwise.

### GetMetroNameOk

`func (o *VirtualDeviceDetailsResponse) GetMetroNameOk() (*string, bool)`

GetMetroNameOk returns a tuple with the MetroName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroName

`func (o *VirtualDeviceDetailsResponse) SetMetroName(v string)`

SetMetroName sets MetroName field to given value.

### HasMetroName

`func (o *VirtualDeviceDetailsResponse) HasMetroName() bool`

HasMetroName returns a boolean if a field has been set.

### GetName

`func (o *VirtualDeviceDetailsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualDeviceDetailsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualDeviceDetailsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualDeviceDetailsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifications

`func (o *VirtualDeviceDetailsResponse) GetNotifications() []string`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *VirtualDeviceDetailsResponse) GetNotificationsOk() (*[]string, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *VirtualDeviceDetailsResponse) SetNotifications(v []string)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *VirtualDeviceDetailsResponse) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetPackageCode

`func (o *VirtualDeviceDetailsResponse) GetPackageCode() string`

GetPackageCode returns the PackageCode field if non-nil, zero value otherwise.

### GetPackageCodeOk

`func (o *VirtualDeviceDetailsResponse) GetPackageCodeOk() (*string, bool)`

GetPackageCodeOk returns a tuple with the PackageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCode

`func (o *VirtualDeviceDetailsResponse) SetPackageCode(v string)`

SetPackageCode sets PackageCode field to given value.

### HasPackageCode

`func (o *VirtualDeviceDetailsResponse) HasPackageCode() bool`

HasPackageCode returns a boolean if a field has been set.

### GetPackageName

`func (o *VirtualDeviceDetailsResponse) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *VirtualDeviceDetailsResponse) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *VirtualDeviceDetailsResponse) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *VirtualDeviceDetailsResponse) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetVersion

`func (o *VirtualDeviceDetailsResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualDeviceDetailsResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualDeviceDetailsResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VirtualDeviceDetailsResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *VirtualDeviceDetailsResponse) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *VirtualDeviceDetailsResponse) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *VirtualDeviceDetailsResponse) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *VirtualDeviceDetailsResponse) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetRedundancyType

`func (o *VirtualDeviceDetailsResponse) GetRedundancyType() string`

GetRedundancyType returns the RedundancyType field if non-nil, zero value otherwise.

### GetRedundancyTypeOk

`func (o *VirtualDeviceDetailsResponse) GetRedundancyTypeOk() (*string, bool)`

GetRedundancyTypeOk returns a tuple with the RedundancyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyType

`func (o *VirtualDeviceDetailsResponse) SetRedundancyType(v string)`

SetRedundancyType sets RedundancyType field to given value.

### HasRedundancyType

`func (o *VirtualDeviceDetailsResponse) HasRedundancyType() bool`

HasRedundancyType returns a boolean if a field has been set.

### GetRedundantUuid

`func (o *VirtualDeviceDetailsResponse) GetRedundantUuid() string`

GetRedundantUuid returns the RedundantUuid field if non-nil, zero value otherwise.

### GetRedundantUuidOk

`func (o *VirtualDeviceDetailsResponse) GetRedundantUuidOk() (*string, bool)`

GetRedundantUuidOk returns a tuple with the RedundantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantUuid

`func (o *VirtualDeviceDetailsResponse) SetRedundantUuid(v string)`

SetRedundantUuid sets RedundantUuid field to given value.

### HasRedundantUuid

`func (o *VirtualDeviceDetailsResponse) HasRedundantUuid() bool`

HasRedundantUuid returns a boolean if a field has been set.

### GetConnectivity

`func (o *VirtualDeviceDetailsResponse) GetConnectivity() string`

GetConnectivity returns the Connectivity field if non-nil, zero value otherwise.

### GetConnectivityOk

`func (o *VirtualDeviceDetailsResponse) GetConnectivityOk() (*string, bool)`

GetConnectivityOk returns a tuple with the Connectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivity

`func (o *VirtualDeviceDetailsResponse) SetConnectivity(v string)`

SetConnectivity sets Connectivity field to given value.

### HasConnectivity

`func (o *VirtualDeviceDetailsResponse) HasConnectivity() bool`

HasConnectivity returns a boolean if a field has been set.

### GetSshIpAddress

`func (o *VirtualDeviceDetailsResponse) GetSshIpAddress() string`

GetSshIpAddress returns the SshIpAddress field if non-nil, zero value otherwise.

### GetSshIpAddressOk

`func (o *VirtualDeviceDetailsResponse) GetSshIpAddressOk() (*string, bool)`

GetSshIpAddressOk returns a tuple with the SshIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshIpAddress

`func (o *VirtualDeviceDetailsResponse) SetSshIpAddress(v string)`

SetSshIpAddress sets SshIpAddress field to given value.

### HasSshIpAddress

`func (o *VirtualDeviceDetailsResponse) HasSshIpAddress() bool`

HasSshIpAddress returns a boolean if a field has been set.

### GetSshIpFqdn

`func (o *VirtualDeviceDetailsResponse) GetSshIpFqdn() string`

GetSshIpFqdn returns the SshIpFqdn field if non-nil, zero value otherwise.

### GetSshIpFqdnOk

`func (o *VirtualDeviceDetailsResponse) GetSshIpFqdnOk() (*string, bool)`

GetSshIpFqdnOk returns a tuple with the SshIpFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshIpFqdn

`func (o *VirtualDeviceDetailsResponse) SetSshIpFqdn(v string)`

SetSshIpFqdn sets SshIpFqdn field to given value.

### HasSshIpFqdn

`func (o *VirtualDeviceDetailsResponse) HasSshIpFqdn() bool`

HasSshIpFqdn returns a boolean if a field has been set.

### GetSshIpFqdnStatus

`func (o *VirtualDeviceDetailsResponse) GetSshIpFqdnStatus() string`

GetSshIpFqdnStatus returns the SshIpFqdnStatus field if non-nil, zero value otherwise.

### GetSshIpFqdnStatusOk

`func (o *VirtualDeviceDetailsResponse) GetSshIpFqdnStatusOk() (*string, bool)`

GetSshIpFqdnStatusOk returns a tuple with the SshIpFqdnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshIpFqdnStatus

`func (o *VirtualDeviceDetailsResponse) SetSshIpFqdnStatus(v string)`

SetSshIpFqdnStatus sets SshIpFqdnStatus field to given value.

### HasSshIpFqdnStatus

`func (o *VirtualDeviceDetailsResponse) HasSshIpFqdnStatus() bool`

HasSshIpFqdnStatus returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualDeviceDetailsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualDeviceDetailsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualDeviceDetailsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualDeviceDetailsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThroughput

`func (o *VirtualDeviceDetailsResponse) GetThroughput() int32`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *VirtualDeviceDetailsResponse) GetThroughputOk() (*int32, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *VirtualDeviceDetailsResponse) SetThroughput(v int32)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *VirtualDeviceDetailsResponse) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetThroughputUnit

`func (o *VirtualDeviceDetailsResponse) GetThroughputUnit() string`

GetThroughputUnit returns the ThroughputUnit field if non-nil, zero value otherwise.

### GetThroughputUnitOk

`func (o *VirtualDeviceDetailsResponse) GetThroughputUnitOk() (*string, bool)`

GetThroughputUnitOk returns a tuple with the ThroughputUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputUnit

`func (o *VirtualDeviceDetailsResponse) SetThroughputUnit(v string)`

SetThroughputUnit sets ThroughputUnit field to given value.

### HasThroughputUnit

`func (o *VirtualDeviceDetailsResponse) HasThroughputUnit() bool`

HasThroughputUnit returns a boolean if a field has been set.

### GetCore

`func (o *VirtualDeviceDetailsResponse) GetCore() CoresDisplayConfig`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *VirtualDeviceDetailsResponse) GetCoreOk() (*CoresDisplayConfig, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *VirtualDeviceDetailsResponse) SetCore(v CoresDisplayConfig)`

SetCore sets Core field to given value.

### HasCore

`func (o *VirtualDeviceDetailsResponse) HasCore() bool`

HasCore returns a boolean if a field has been set.

### GetPricingDetails

`func (o *VirtualDeviceDetailsResponse) GetPricingDetails() PricingSiebelConfig`

GetPricingDetails returns the PricingDetails field if non-nil, zero value otherwise.

### GetPricingDetailsOk

`func (o *VirtualDeviceDetailsResponse) GetPricingDetailsOk() (*PricingSiebelConfig, bool)`

GetPricingDetailsOk returns a tuple with the PricingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingDetails

`func (o *VirtualDeviceDetailsResponse) SetPricingDetails(v PricingSiebelConfig)`

SetPricingDetails sets PricingDetails field to given value.

### HasPricingDetails

`func (o *VirtualDeviceDetailsResponse) HasPricingDetails() bool`

HasPricingDetails returns a boolean if a field has been set.

### GetInterfaceCount

`func (o *VirtualDeviceDetailsResponse) GetInterfaceCount() int32`

GetInterfaceCount returns the InterfaceCount field if non-nil, zero value otherwise.

### GetInterfaceCountOk

`func (o *VirtualDeviceDetailsResponse) GetInterfaceCountOk() (*int32, bool)`

GetInterfaceCountOk returns a tuple with the InterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceCount

`func (o *VirtualDeviceDetailsResponse) SetInterfaceCount(v int32)`

SetInterfaceCount sets InterfaceCount field to given value.

### HasInterfaceCount

`func (o *VirtualDeviceDetailsResponse) HasInterfaceCount() bool`

HasInterfaceCount returns a boolean if a field has been set.

### GetDeviceManagementType

`func (o *VirtualDeviceDetailsResponse) GetDeviceManagementType() string`

GetDeviceManagementType returns the DeviceManagementType field if non-nil, zero value otherwise.

### GetDeviceManagementTypeOk

`func (o *VirtualDeviceDetailsResponse) GetDeviceManagementTypeOk() (*string, bool)`

GetDeviceManagementTypeOk returns a tuple with the DeviceManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceManagementType

`func (o *VirtualDeviceDetailsResponse) SetDeviceManagementType(v string)`

SetDeviceManagementType sets DeviceManagementType field to given value.

### HasDeviceManagementType

`func (o *VirtualDeviceDetailsResponse) HasDeviceManagementType() bool`

HasDeviceManagementType returns a boolean if a field has been set.

### GetUserPublicKey

`func (o *VirtualDeviceDetailsResponse) GetUserPublicKey() UserPublicKeyConfig`

GetUserPublicKey returns the UserPublicKey field if non-nil, zero value otherwise.

### GetUserPublicKeyOk

`func (o *VirtualDeviceDetailsResponse) GetUserPublicKeyOk() (*UserPublicKeyConfig, bool)`

GetUserPublicKeyOk returns a tuple with the UserPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPublicKey

`func (o *VirtualDeviceDetailsResponse) SetUserPublicKey(v UserPublicKeyConfig)`

SetUserPublicKey sets UserPublicKey field to given value.

### HasUserPublicKey

`func (o *VirtualDeviceDetailsResponse) HasUserPublicKey() bool`

HasUserPublicKey returns a boolean if a field has been set.

### GetManagementIp

`func (o *VirtualDeviceDetailsResponse) GetManagementIp() string`

GetManagementIp returns the ManagementIp field if non-nil, zero value otherwise.

### GetManagementIpOk

`func (o *VirtualDeviceDetailsResponse) GetManagementIpOk() (*string, bool)`

GetManagementIpOk returns a tuple with the ManagementIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementIp

`func (o *VirtualDeviceDetailsResponse) SetManagementIp(v string)`

SetManagementIp sets ManagementIp field to given value.

### HasManagementIp

`func (o *VirtualDeviceDetailsResponse) HasManagementIp() bool`

HasManagementIp returns a boolean if a field has been set.

### GetManagementGatewayIp

`func (o *VirtualDeviceDetailsResponse) GetManagementGatewayIp() string`

GetManagementGatewayIp returns the ManagementGatewayIp field if non-nil, zero value otherwise.

### GetManagementGatewayIpOk

`func (o *VirtualDeviceDetailsResponse) GetManagementGatewayIpOk() (*string, bool)`

GetManagementGatewayIpOk returns a tuple with the ManagementGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementGatewayIp

`func (o *VirtualDeviceDetailsResponse) SetManagementGatewayIp(v string)`

SetManagementGatewayIp sets ManagementGatewayIp field to given value.

### HasManagementGatewayIp

`func (o *VirtualDeviceDetailsResponse) HasManagementGatewayIp() bool`

HasManagementGatewayIp returns a boolean if a field has been set.

### GetPublicIp

`func (o *VirtualDeviceDetailsResponse) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *VirtualDeviceDetailsResponse) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *VirtualDeviceDetailsResponse) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *VirtualDeviceDetailsResponse) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPublicGatewayIp

`func (o *VirtualDeviceDetailsResponse) GetPublicGatewayIp() string`

GetPublicGatewayIp returns the PublicGatewayIp field if non-nil, zero value otherwise.

### GetPublicGatewayIpOk

`func (o *VirtualDeviceDetailsResponse) GetPublicGatewayIpOk() (*string, bool)`

GetPublicGatewayIpOk returns a tuple with the PublicGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicGatewayIp

`func (o *VirtualDeviceDetailsResponse) SetPublicGatewayIp(v string)`

SetPublicGatewayIp sets PublicGatewayIp field to given value.

### HasPublicGatewayIp

`func (o *VirtualDeviceDetailsResponse) HasPublicGatewayIp() bool`

HasPublicGatewayIp returns a boolean if a field has been set.

### GetPrimaryDnsName

`func (o *VirtualDeviceDetailsResponse) GetPrimaryDnsName() string`

GetPrimaryDnsName returns the PrimaryDnsName field if non-nil, zero value otherwise.

### GetPrimaryDnsNameOk

`func (o *VirtualDeviceDetailsResponse) GetPrimaryDnsNameOk() (*string, bool)`

GetPrimaryDnsNameOk returns a tuple with the PrimaryDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDnsName

`func (o *VirtualDeviceDetailsResponse) SetPrimaryDnsName(v string)`

SetPrimaryDnsName sets PrimaryDnsName field to given value.

### HasPrimaryDnsName

`func (o *VirtualDeviceDetailsResponse) HasPrimaryDnsName() bool`

HasPrimaryDnsName returns a boolean if a field has been set.

### GetSecondaryDnsName

`func (o *VirtualDeviceDetailsResponse) GetSecondaryDnsName() string`

GetSecondaryDnsName returns the SecondaryDnsName field if non-nil, zero value otherwise.

### GetSecondaryDnsNameOk

`func (o *VirtualDeviceDetailsResponse) GetSecondaryDnsNameOk() (*string, bool)`

GetSecondaryDnsNameOk returns a tuple with the SecondaryDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDnsName

`func (o *VirtualDeviceDetailsResponse) SetSecondaryDnsName(v string)`

SetSecondaryDnsName sets SecondaryDnsName field to given value.

### HasSecondaryDnsName

`func (o *VirtualDeviceDetailsResponse) HasSecondaryDnsName() bool`

HasSecondaryDnsName returns a boolean if a field has been set.

### GetTermLength

`func (o *VirtualDeviceDetailsResponse) GetTermLength() string`

GetTermLength returns the TermLength field if non-nil, zero value otherwise.

### GetTermLengthOk

`func (o *VirtualDeviceDetailsResponse) GetTermLengthOk() (*string, bool)`

GetTermLengthOk returns a tuple with the TermLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermLength

`func (o *VirtualDeviceDetailsResponse) SetTermLength(v string)`

SetTermLength sets TermLength field to given value.

### HasTermLength

`func (o *VirtualDeviceDetailsResponse) HasTermLength() bool`

HasTermLength returns a boolean if a field has been set.

### GetAdditionalBandwidth

`func (o *VirtualDeviceDetailsResponse) GetAdditionalBandwidth() string`

GetAdditionalBandwidth returns the AdditionalBandwidth field if non-nil, zero value otherwise.

### GetAdditionalBandwidthOk

`func (o *VirtualDeviceDetailsResponse) GetAdditionalBandwidthOk() (*string, bool)`

GetAdditionalBandwidthOk returns a tuple with the AdditionalBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalBandwidth

`func (o *VirtualDeviceDetailsResponse) SetAdditionalBandwidth(v string)`

SetAdditionalBandwidth sets AdditionalBandwidth field to given value.

### HasAdditionalBandwidth

`func (o *VirtualDeviceDetailsResponse) HasAdditionalBandwidth() bool`

HasAdditionalBandwidth returns a boolean if a field has been set.

### GetSiteId

`func (o *VirtualDeviceDetailsResponse) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VirtualDeviceDetailsResponse) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VirtualDeviceDetailsResponse) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *VirtualDeviceDetailsResponse) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSystemIpAddress

`func (o *VirtualDeviceDetailsResponse) GetSystemIpAddress() string`

GetSystemIpAddress returns the SystemIpAddress field if non-nil, zero value otherwise.

### GetSystemIpAddressOk

`func (o *VirtualDeviceDetailsResponse) GetSystemIpAddressOk() (*string, bool)`

GetSystemIpAddressOk returns a tuple with the SystemIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIpAddress

`func (o *VirtualDeviceDetailsResponse) SetSystemIpAddress(v string)`

SetSystemIpAddress sets SystemIpAddress field to given value.

### HasSystemIpAddress

`func (o *VirtualDeviceDetailsResponse) HasSystemIpAddress() bool`

HasSystemIpAddress returns a boolean if a field has been set.

### GetVendorConfig

`func (o *VirtualDeviceDetailsResponse) GetVendorConfig() VendorConfig`

GetVendorConfig returns the VendorConfig field if non-nil, zero value otherwise.

### GetVendorConfigOk

`func (o *VirtualDeviceDetailsResponse) GetVendorConfigOk() (*VendorConfig, bool)`

GetVendorConfigOk returns a tuple with the VendorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorConfig

`func (o *VirtualDeviceDetailsResponse) SetVendorConfig(v VendorConfig)`

SetVendorConfig sets VendorConfig field to given value.

### HasVendorConfig

`func (o *VirtualDeviceDetailsResponse) HasVendorConfig() bool`

HasVendorConfig returns a boolean if a field has been set.

### GetInterfaces

`func (o *VirtualDeviceDetailsResponse) GetInterfaces() []InterfaceBasicInfoResponse`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *VirtualDeviceDetailsResponse) GetInterfacesOk() (*[]InterfaceBasicInfoResponse, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *VirtualDeviceDetailsResponse) SetInterfaces(v []InterfaceBasicInfoResponse)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *VirtualDeviceDetailsResponse) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetAsn

`func (o *VirtualDeviceDetailsResponse) GetAsn() float32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *VirtualDeviceDetailsResponse) GetAsnOk() (*float32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *VirtualDeviceDetailsResponse) SetAsn(v float32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *VirtualDeviceDetailsResponse) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetChannelPartner

`func (o *VirtualDeviceDetailsResponse) GetChannelPartner() string`

GetChannelPartner returns the ChannelPartner field if non-nil, zero value otherwise.

### GetChannelPartnerOk

`func (o *VirtualDeviceDetailsResponse) GetChannelPartnerOk() (*string, bool)`

GetChannelPartnerOk returns a tuple with the ChannelPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPartner

`func (o *VirtualDeviceDetailsResponse) SetChannelPartner(v string)`

SetChannelPartner sets ChannelPartner field to given value.

### HasChannelPartner

`func (o *VirtualDeviceDetailsResponse) HasChannelPartner() bool`

HasChannelPartner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


