# VendorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **string** | Physical location within the Viptela overlay network, such as a branch office, or a campus (relevant only for Cisco SD-WANs) | [optional] 
**SystemIpAddress** | Pointer to **string** | IP assigned to CSRSDWAN router and vSmart controller (relevant only for Cisco SD-WANs) | [optional] 
**LicenseKey** | Pointer to **string** |  | [optional] 
**LicenseSecret** | Pointer to **string** |  | [optional] 
**LocalId** | Pointer to **string** |  | [optional] 
**RemoteId** | Pointer to **string** |  | [optional] 
**ManagementType** | Pointer to **string** | This is required for Cisco FTD Firewall devices. If you choose \&quot;FMC,\&quot; you must also provide the controller IP and the activation key. | [optional] 
**Controller1** | Pointer to **string** | For Fortinet devices, this is the System IP address. | [optional] 
**Controller2** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**AdminPassword** | Pointer to **string** | The administrative password of the device. You can use it to log in to the console. This field is not available for all device types. Should be at least 6 characters long and must include an uppercase letter and a number. This field may be required for some vendors. | [optional] 
**ActivationKey** | Pointer to **string** |  | [optional] 
**ControllerFqdn** | Pointer to **string** |  | [optional] 
**RootPassword** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** | The name of account. | [optional] 
**Hostname** | Pointer to **string** | The host name. | [optional] 
**AccountKey** | Pointer to **string** | The account key. | [optional] 
**ApplianceTag** | Pointer to **string** | The appliance tag. | [optional] 
**UserName** | Pointer to **string** | This field is rel | [optional] 
**ConnectToCloudVision** | Pointer to **bool** | Whether you want your Arista device to connect to Cloud Vision. Only relevant for Arista devices. | [optional] 
**CvpType** | Pointer to **string** | Either As-a-Service or On-Premise. Only relevant for Arista devices. | [optional] 
**CvpFqdn** | Pointer to **string** | Fully qualified domain name for Cloud Vision As-a-Service. Only relevant for Arista devices. | [optional] 
**CvpIpAddress** | Pointer to **string** | Only relevant for Arista devices. CvpIpAddress is required if connectToCloudVision&#x3D;true and cvpType&#x3D;On-Premise. | [optional] 
**CvaasPort** | Pointer to **string** | Only relevant for Arista devices. CvaasPort is required if connectToCloudVision&#x3D;true and cvpType&#x3D;As-a-Service. | [optional] 
**CvpPort** | Pointer to **string** | Only relevant for Arista devices. CvpPort is required if connectToCloudVision&#x3D;true and cvpType&#x3D;On-Premise. | [optional] 
**CvpToken** | Pointer to **string** | Only relevant for Arista devices. CvpToken is required if connectToCloudVision&#x3D;true and (cvpType&#x3D;On-Premise or cvpType&#x3D;As-a-Service). | [optional] 

## Methods

### NewVendorConfig

`func NewVendorConfig() *VendorConfig`

NewVendorConfig instantiates a new VendorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorConfigWithDefaults

`func NewVendorConfigWithDefaults() *VendorConfig`

NewVendorConfigWithDefaults instantiates a new VendorConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *VendorConfig) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VendorConfig) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VendorConfig) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *VendorConfig) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSystemIpAddress

`func (o *VendorConfig) GetSystemIpAddress() string`

GetSystemIpAddress returns the SystemIpAddress field if non-nil, zero value otherwise.

### GetSystemIpAddressOk

`func (o *VendorConfig) GetSystemIpAddressOk() (*string, bool)`

GetSystemIpAddressOk returns a tuple with the SystemIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIpAddress

`func (o *VendorConfig) SetSystemIpAddress(v string)`

SetSystemIpAddress sets SystemIpAddress field to given value.

### HasSystemIpAddress

`func (o *VendorConfig) HasSystemIpAddress() bool`

HasSystemIpAddress returns a boolean if a field has been set.

### GetLicenseKey

`func (o *VendorConfig) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *VendorConfig) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *VendorConfig) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *VendorConfig) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetLicenseSecret

`func (o *VendorConfig) GetLicenseSecret() string`

GetLicenseSecret returns the LicenseSecret field if non-nil, zero value otherwise.

### GetLicenseSecretOk

`func (o *VendorConfig) GetLicenseSecretOk() (*string, bool)`

GetLicenseSecretOk returns a tuple with the LicenseSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseSecret

`func (o *VendorConfig) SetLicenseSecret(v string)`

SetLicenseSecret sets LicenseSecret field to given value.

### HasLicenseSecret

`func (o *VendorConfig) HasLicenseSecret() bool`

HasLicenseSecret returns a boolean if a field has been set.

### GetLocalId

`func (o *VendorConfig) GetLocalId() string`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *VendorConfig) GetLocalIdOk() (*string, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *VendorConfig) SetLocalId(v string)`

SetLocalId sets LocalId field to given value.

### HasLocalId

`func (o *VendorConfig) HasLocalId() bool`

HasLocalId returns a boolean if a field has been set.

### GetRemoteId

`func (o *VendorConfig) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *VendorConfig) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *VendorConfig) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *VendorConfig) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetManagementType

`func (o *VendorConfig) GetManagementType() string`

GetManagementType returns the ManagementType field if non-nil, zero value otherwise.

### GetManagementTypeOk

`func (o *VendorConfig) GetManagementTypeOk() (*string, bool)`

GetManagementTypeOk returns a tuple with the ManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementType

`func (o *VendorConfig) SetManagementType(v string)`

SetManagementType sets ManagementType field to given value.

### HasManagementType

`func (o *VendorConfig) HasManagementType() bool`

HasManagementType returns a boolean if a field has been set.

### GetController1

`func (o *VendorConfig) GetController1() string`

GetController1 returns the Controller1 field if non-nil, zero value otherwise.

### GetController1Ok

`func (o *VendorConfig) GetController1Ok() (*string, bool)`

GetController1Ok returns a tuple with the Controller1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController1

`func (o *VendorConfig) SetController1(v string)`

SetController1 sets Controller1 field to given value.

### HasController1

`func (o *VendorConfig) HasController1() bool`

HasController1 returns a boolean if a field has been set.

### GetController2

`func (o *VendorConfig) GetController2() string`

GetController2 returns the Controller2 field if non-nil, zero value otherwise.

### GetController2Ok

`func (o *VendorConfig) GetController2Ok() (*string, bool)`

GetController2Ok returns a tuple with the Controller2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController2

`func (o *VendorConfig) SetController2(v string)`

SetController2 sets Controller2 field to given value.

### HasController2

`func (o *VendorConfig) HasController2() bool`

HasController2 returns a boolean if a field has been set.

### GetSerialNumber

`func (o *VendorConfig) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *VendorConfig) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *VendorConfig) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *VendorConfig) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetAdminPassword

`func (o *VendorConfig) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *VendorConfig) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *VendorConfig) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### HasAdminPassword

`func (o *VendorConfig) HasAdminPassword() bool`

HasAdminPassword returns a boolean if a field has been set.

### GetActivationKey

`func (o *VendorConfig) GetActivationKey() string`

GetActivationKey returns the ActivationKey field if non-nil, zero value otherwise.

### GetActivationKeyOk

`func (o *VendorConfig) GetActivationKeyOk() (*string, bool)`

GetActivationKeyOk returns a tuple with the ActivationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationKey

`func (o *VendorConfig) SetActivationKey(v string)`

SetActivationKey sets ActivationKey field to given value.

### HasActivationKey

`func (o *VendorConfig) HasActivationKey() bool`

HasActivationKey returns a boolean if a field has been set.

### GetControllerFqdn

`func (o *VendorConfig) GetControllerFqdn() string`

GetControllerFqdn returns the ControllerFqdn field if non-nil, zero value otherwise.

### GetControllerFqdnOk

`func (o *VendorConfig) GetControllerFqdnOk() (*string, bool)`

GetControllerFqdnOk returns a tuple with the ControllerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerFqdn

`func (o *VendorConfig) SetControllerFqdn(v string)`

SetControllerFqdn sets ControllerFqdn field to given value.

### HasControllerFqdn

`func (o *VendorConfig) HasControllerFqdn() bool`

HasControllerFqdn returns a boolean if a field has been set.

### GetRootPassword

`func (o *VendorConfig) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *VendorConfig) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *VendorConfig) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *VendorConfig) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetAccountName

`func (o *VendorConfig) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *VendorConfig) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *VendorConfig) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *VendorConfig) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetHostname

`func (o *VendorConfig) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VendorConfig) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VendorConfig) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *VendorConfig) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetAccountKey

`func (o *VendorConfig) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *VendorConfig) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *VendorConfig) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.

### HasAccountKey

`func (o *VendorConfig) HasAccountKey() bool`

HasAccountKey returns a boolean if a field has been set.

### GetApplianceTag

`func (o *VendorConfig) GetApplianceTag() string`

GetApplianceTag returns the ApplianceTag field if non-nil, zero value otherwise.

### GetApplianceTagOk

`func (o *VendorConfig) GetApplianceTagOk() (*string, bool)`

GetApplianceTagOk returns a tuple with the ApplianceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceTag

`func (o *VendorConfig) SetApplianceTag(v string)`

SetApplianceTag sets ApplianceTag field to given value.

### HasApplianceTag

`func (o *VendorConfig) HasApplianceTag() bool`

HasApplianceTag returns a boolean if a field has been set.

### GetUserName

`func (o *VendorConfig) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *VendorConfig) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *VendorConfig) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *VendorConfig) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetConnectToCloudVision

`func (o *VendorConfig) GetConnectToCloudVision() bool`

GetConnectToCloudVision returns the ConnectToCloudVision field if non-nil, zero value otherwise.

### GetConnectToCloudVisionOk

`func (o *VendorConfig) GetConnectToCloudVisionOk() (*bool, bool)`

GetConnectToCloudVisionOk returns a tuple with the ConnectToCloudVision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectToCloudVision

`func (o *VendorConfig) SetConnectToCloudVision(v bool)`

SetConnectToCloudVision sets ConnectToCloudVision field to given value.

### HasConnectToCloudVision

`func (o *VendorConfig) HasConnectToCloudVision() bool`

HasConnectToCloudVision returns a boolean if a field has been set.

### GetCvpType

`func (o *VendorConfig) GetCvpType() string`

GetCvpType returns the CvpType field if non-nil, zero value otherwise.

### GetCvpTypeOk

`func (o *VendorConfig) GetCvpTypeOk() (*string, bool)`

GetCvpTypeOk returns a tuple with the CvpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvpType

`func (o *VendorConfig) SetCvpType(v string)`

SetCvpType sets CvpType field to given value.

### HasCvpType

`func (o *VendorConfig) HasCvpType() bool`

HasCvpType returns a boolean if a field has been set.

### GetCvpFqdn

`func (o *VendorConfig) GetCvpFqdn() string`

GetCvpFqdn returns the CvpFqdn field if non-nil, zero value otherwise.

### GetCvpFqdnOk

`func (o *VendorConfig) GetCvpFqdnOk() (*string, bool)`

GetCvpFqdnOk returns a tuple with the CvpFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvpFqdn

`func (o *VendorConfig) SetCvpFqdn(v string)`

SetCvpFqdn sets CvpFqdn field to given value.

### HasCvpFqdn

`func (o *VendorConfig) HasCvpFqdn() bool`

HasCvpFqdn returns a boolean if a field has been set.

### GetCvpIpAddress

`func (o *VendorConfig) GetCvpIpAddress() string`

GetCvpIpAddress returns the CvpIpAddress field if non-nil, zero value otherwise.

### GetCvpIpAddressOk

`func (o *VendorConfig) GetCvpIpAddressOk() (*string, bool)`

GetCvpIpAddressOk returns a tuple with the CvpIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvpIpAddress

`func (o *VendorConfig) SetCvpIpAddress(v string)`

SetCvpIpAddress sets CvpIpAddress field to given value.

### HasCvpIpAddress

`func (o *VendorConfig) HasCvpIpAddress() bool`

HasCvpIpAddress returns a boolean if a field has been set.

### GetCvaasPort

`func (o *VendorConfig) GetCvaasPort() string`

GetCvaasPort returns the CvaasPort field if non-nil, zero value otherwise.

### GetCvaasPortOk

`func (o *VendorConfig) GetCvaasPortOk() (*string, bool)`

GetCvaasPortOk returns a tuple with the CvaasPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvaasPort

`func (o *VendorConfig) SetCvaasPort(v string)`

SetCvaasPort sets CvaasPort field to given value.

### HasCvaasPort

`func (o *VendorConfig) HasCvaasPort() bool`

HasCvaasPort returns a boolean if a field has been set.

### GetCvpPort

`func (o *VendorConfig) GetCvpPort() string`

GetCvpPort returns the CvpPort field if non-nil, zero value otherwise.

### GetCvpPortOk

`func (o *VendorConfig) GetCvpPortOk() (*string, bool)`

GetCvpPortOk returns a tuple with the CvpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvpPort

`func (o *VendorConfig) SetCvpPort(v string)`

SetCvpPort sets CvpPort field to given value.

### HasCvpPort

`func (o *VendorConfig) HasCvpPort() bool`

HasCvpPort returns a boolean if a field has been set.

### GetCvpToken

`func (o *VendorConfig) GetCvpToken() string`

GetCvpToken returns the CvpToken field if non-nil, zero value otherwise.

### GetCvpTokenOk

`func (o *VendorConfig) GetCvpTokenOk() (*string, bool)`

GetCvpTokenOk returns a tuple with the CvpToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvpToken

`func (o *VendorConfig) SetCvpToken(v string)`

SetCvpToken sets CvpToken field to given value.

### HasCvpToken

`func (o *VendorConfig) HasCvpToken() bool`

HasCvpToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


