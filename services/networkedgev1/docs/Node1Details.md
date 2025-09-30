# Node1Details

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseFileId** | Pointer to **string** | License file id is required for Fortinet and Juniper clusters. | [optional] 
**LicenseToken** | Pointer to **string** | License token is required for Palo Alto clusters. | [optional] 
**VendorConfig** | Pointer to [**VendorConfigDetailsNode1**](VendorConfigDetailsNode1.md) |  | [optional] 

## Methods

### NewNode1Details

`func NewNode1Details() *Node1Details`

NewNode1Details instantiates a new Node1Details object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNode1DetailsWithDefaults

`func NewNode1DetailsWithDefaults() *Node1Details`

NewNode1DetailsWithDefaults instantiates a new Node1Details object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseFileId

`func (o *Node1Details) GetLicenseFileId() string`

GetLicenseFileId returns the LicenseFileId field if non-nil, zero value otherwise.

### GetLicenseFileIdOk

`func (o *Node1Details) GetLicenseFileIdOk() (*string, bool)`

GetLicenseFileIdOk returns a tuple with the LicenseFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseFileId

`func (o *Node1Details) SetLicenseFileId(v string)`

SetLicenseFileId sets LicenseFileId field to given value.

### HasLicenseFileId

`func (o *Node1Details) HasLicenseFileId() bool`

HasLicenseFileId returns a boolean if a field has been set.

### GetLicenseToken

`func (o *Node1Details) GetLicenseToken() string`

GetLicenseToken returns the LicenseToken field if non-nil, zero value otherwise.

### GetLicenseTokenOk

`func (o *Node1Details) GetLicenseTokenOk() (*string, bool)`

GetLicenseTokenOk returns a tuple with the LicenseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseToken

`func (o *Node1Details) SetLicenseToken(v string)`

SetLicenseToken sets LicenseToken field to given value.

### HasLicenseToken

`func (o *Node1Details) HasLicenseToken() bool`

HasLicenseToken returns a boolean if a field has been set.

### GetVendorConfig

`func (o *Node1Details) GetVendorConfig() VendorConfigDetailsNode1`

GetVendorConfig returns the VendorConfig field if non-nil, zero value otherwise.

### GetVendorConfigOk

`func (o *Node1Details) GetVendorConfigOk() (*VendorConfigDetailsNode1, bool)`

GetVendorConfigOk returns a tuple with the VendorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorConfig

`func (o *Node1Details) SetVendorConfig(v VendorConfigDetailsNode1)`

SetVendorConfig sets VendorConfig field to given value.

### HasVendorConfig

`func (o *Node1Details) HasVendorConfig() bool`

HasVendorConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


