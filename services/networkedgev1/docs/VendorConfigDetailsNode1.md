# VendorConfigDetailsNode1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | The host name. Only a-z, A-Z, 0-9, and hyphen(-) are allowed. It should start with a letter and end with a letter or digit. The length should be between 2-30 characters. Exceptions - FTDv 2-14; Aruba 2-24. | [optional] 
**RootPassword** | Pointer to **string** |  | [optional] 
**AdminPassword** | Pointer to **string** | The administrative password of the device. You can use it to log in to the console. This field is not available for all device types. | [optional] 

## Methods

### NewVendorConfigDetailsNode1

`func NewVendorConfigDetailsNode1() *VendorConfigDetailsNode1`

NewVendorConfigDetailsNode1 instantiates a new VendorConfigDetailsNode1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorConfigDetailsNode1WithDefaults

`func NewVendorConfigDetailsNode1WithDefaults() *VendorConfigDetailsNode1`

NewVendorConfigDetailsNode1WithDefaults instantiates a new VendorConfigDetailsNode1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *VendorConfigDetailsNode1) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VendorConfigDetailsNode1) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VendorConfigDetailsNode1) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *VendorConfigDetailsNode1) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetRootPassword

`func (o *VendorConfigDetailsNode1) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *VendorConfigDetailsNode1) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *VendorConfigDetailsNode1) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *VendorConfigDetailsNode1) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetAdminPassword

`func (o *VendorConfigDetailsNode1) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *VendorConfigDetailsNode1) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *VendorConfigDetailsNode1) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### HasAdminPassword

`func (o *VendorConfigDetailsNode1) HasAdminPassword() bool`

HasAdminPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


