# DeviceRMAPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Any version you want. | 
**CloudInitFileId** | Pointer to **string** | For a C8KV device, this is the Id of the uploaded bootstrap file. Upload your Cisco bootstrap file by calling Upload File. In the response, you&#39;ll get a fileUuid that you can enter here as cloudInitFileId. This field may be required for some vendors. | [optional] 
**LicenseFileId** | Pointer to **string** | This is the Id of the uploaded license file. For a CSR1KV SDWAN device, upload your license file by calling Post License File. In the response, you&#39;ll get a fileId that you can enter here as licenseFileId. This field may be required for some vendors. | [optional] 
**Token** | Pointer to **string** | License token. For a cluster, you will need to provide license tokens for both node0 and node1. To get the exact payload for different vendors, check the Postman script on the API Reference page of online documentation. | [optional] 
**VendorConfig** | Pointer to [**RMAVendorConfig**](RMAVendorConfig.md) |  | [optional] 
**UserPublicKey** | Pointer to [**UserPublicKeyRequest**](UserPublicKeyRequest.md) |  | [optional] 

## Methods

### NewDeviceRMAPostRequest

`func NewDeviceRMAPostRequest(version string, ) *DeviceRMAPostRequest`

NewDeviceRMAPostRequest instantiates a new DeviceRMAPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceRMAPostRequestWithDefaults

`func NewDeviceRMAPostRequestWithDefaults() *DeviceRMAPostRequest`

NewDeviceRMAPostRequestWithDefaults instantiates a new DeviceRMAPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DeviceRMAPostRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceRMAPostRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceRMAPostRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetCloudInitFileId

`func (o *DeviceRMAPostRequest) GetCloudInitFileId() string`

GetCloudInitFileId returns the CloudInitFileId field if non-nil, zero value otherwise.

### GetCloudInitFileIdOk

`func (o *DeviceRMAPostRequest) GetCloudInitFileIdOk() (*string, bool)`

GetCloudInitFileIdOk returns a tuple with the CloudInitFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitFileId

`func (o *DeviceRMAPostRequest) SetCloudInitFileId(v string)`

SetCloudInitFileId sets CloudInitFileId field to given value.

### HasCloudInitFileId

`func (o *DeviceRMAPostRequest) HasCloudInitFileId() bool`

HasCloudInitFileId returns a boolean if a field has been set.

### GetLicenseFileId

`func (o *DeviceRMAPostRequest) GetLicenseFileId() string`

GetLicenseFileId returns the LicenseFileId field if non-nil, zero value otherwise.

### GetLicenseFileIdOk

`func (o *DeviceRMAPostRequest) GetLicenseFileIdOk() (*string, bool)`

GetLicenseFileIdOk returns a tuple with the LicenseFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseFileId

`func (o *DeviceRMAPostRequest) SetLicenseFileId(v string)`

SetLicenseFileId sets LicenseFileId field to given value.

### HasLicenseFileId

`func (o *DeviceRMAPostRequest) HasLicenseFileId() bool`

HasLicenseFileId returns a boolean if a field has been set.

### GetToken

`func (o *DeviceRMAPostRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeviceRMAPostRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeviceRMAPostRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DeviceRMAPostRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetVendorConfig

`func (o *DeviceRMAPostRequest) GetVendorConfig() RMAVendorConfig`

GetVendorConfig returns the VendorConfig field if non-nil, zero value otherwise.

### GetVendorConfigOk

`func (o *DeviceRMAPostRequest) GetVendorConfigOk() (*RMAVendorConfig, bool)`

GetVendorConfigOk returns a tuple with the VendorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorConfig

`func (o *DeviceRMAPostRequest) SetVendorConfig(v RMAVendorConfig)`

SetVendorConfig sets VendorConfig field to given value.

### HasVendorConfig

`func (o *DeviceRMAPostRequest) HasVendorConfig() bool`

HasVendorConfig returns a boolean if a field has been set.

### GetUserPublicKey

`func (o *DeviceRMAPostRequest) GetUserPublicKey() UserPublicKeyRequest`

GetUserPublicKey returns the UserPublicKey field if non-nil, zero value otherwise.

### GetUserPublicKeyOk

`func (o *DeviceRMAPostRequest) GetUserPublicKeyOk() (*UserPublicKeyRequest, bool)`

GetUserPublicKeyOk returns a tuple with the UserPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPublicKey

`func (o *DeviceRMAPostRequest) SetUserPublicKey(v UserPublicKeyRequest)`

SetUserPublicKey sets UserPublicKey field to given value.

### HasUserPublicKey

`func (o *DeviceRMAPostRequest) HasUserPublicKey() bool`

HasUserPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


