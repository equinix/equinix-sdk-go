# Node0Details

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseFileId** | Pointer to **string** | License file id is required for Fortinet and Juniper clusters. | [optional] 
**LicenseToken** | Pointer to **string** | License token is required for Palo Alto clusters. | [optional] 
**VendorConfig** | Pointer to [**VendorConfigDetailsNode0**](VendorConfigDetailsNode0.md) |  | [optional] 

## Methods

### NewNode0Details

`func NewNode0Details() *Node0Details`

NewNode0Details instantiates a new Node0Details object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNode0DetailsWithDefaults

`func NewNode0DetailsWithDefaults() *Node0Details`

NewNode0DetailsWithDefaults instantiates a new Node0Details object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseFileId

`func (o *Node0Details) GetLicenseFileId() string`

GetLicenseFileId returns the LicenseFileId field if non-nil, zero value otherwise.

### GetLicenseFileIdOk

`func (o *Node0Details) GetLicenseFileIdOk() (*string, bool)`

GetLicenseFileIdOk returns a tuple with the LicenseFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseFileId

`func (o *Node0Details) SetLicenseFileId(v string)`

SetLicenseFileId sets LicenseFileId field to given value.

### HasLicenseFileId

`func (o *Node0Details) HasLicenseFileId() bool`

HasLicenseFileId returns a boolean if a field has been set.

### GetLicenseToken

`func (o *Node0Details) GetLicenseToken() string`

GetLicenseToken returns the LicenseToken field if non-nil, zero value otherwise.

### GetLicenseTokenOk

`func (o *Node0Details) GetLicenseTokenOk() (*string, bool)`

GetLicenseTokenOk returns a tuple with the LicenseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseToken

`func (o *Node0Details) SetLicenseToken(v string)`

SetLicenseToken sets LicenseToken field to given value.

### HasLicenseToken

`func (o *Node0Details) HasLicenseToken() bool`

HasLicenseToken returns a boolean if a field has been set.

### GetVendorConfig

`func (o *Node0Details) GetVendorConfig() VendorConfigDetailsNode0`

GetVendorConfig returns the VendorConfig field if non-nil, zero value otherwise.

### GetVendorConfigOk

`func (o *Node0Details) GetVendorConfigOk() (*VendorConfigDetailsNode0, bool)`

GetVendorConfigOk returns a tuple with the VendorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorConfig

`func (o *Node0Details) SetVendorConfig(v VendorConfigDetailsNode0)`

SetVendorConfig sets VendorConfig field to given value.

### HasVendorConfig

`func (o *Node0Details) HasVendorConfig() bool`

HasVendorConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


