# VendorConfigDetailsNode0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | The host name. Only a-z, A-Z, 0-9, and hyphen(-) are allowed. It should start with a letter and end with a letter or digit. The length should be between 2-30 characters. Exceptions - FTDv 2-14; Aruba 2-24. | [optional] 
**ActivationKey** | Pointer to **string** |  | [optional] 
**ControllerFqdn** | Pointer to **string** |  | [optional] 
**RootPassword** | Pointer to **string** |  | [optional] 
**AdminPassword** | Pointer to **string** | The administrative password of the device. You can use it to log in to the console. This field is not available for all device types. | [optional] 
**Controller1** | Pointer to **string** |  | [optional] 

## Methods

### NewVendorConfigDetailsNode0

`func NewVendorConfigDetailsNode0() *VendorConfigDetailsNode0`

NewVendorConfigDetailsNode0 instantiates a new VendorConfigDetailsNode0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorConfigDetailsNode0WithDefaults

`func NewVendorConfigDetailsNode0WithDefaults() *VendorConfigDetailsNode0`

NewVendorConfigDetailsNode0WithDefaults instantiates a new VendorConfigDetailsNode0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *VendorConfigDetailsNode0) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VendorConfigDetailsNode0) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VendorConfigDetailsNode0) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *VendorConfigDetailsNode0) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetActivationKey

`func (o *VendorConfigDetailsNode0) GetActivationKey() string`

GetActivationKey returns the ActivationKey field if non-nil, zero value otherwise.

### GetActivationKeyOk

`func (o *VendorConfigDetailsNode0) GetActivationKeyOk() (*string, bool)`

GetActivationKeyOk returns a tuple with the ActivationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationKey

`func (o *VendorConfigDetailsNode0) SetActivationKey(v string)`

SetActivationKey sets ActivationKey field to given value.

### HasActivationKey

`func (o *VendorConfigDetailsNode0) HasActivationKey() bool`

HasActivationKey returns a boolean if a field has been set.

### GetControllerFqdn

`func (o *VendorConfigDetailsNode0) GetControllerFqdn() string`

GetControllerFqdn returns the ControllerFqdn field if non-nil, zero value otherwise.

### GetControllerFqdnOk

`func (o *VendorConfigDetailsNode0) GetControllerFqdnOk() (*string, bool)`

GetControllerFqdnOk returns a tuple with the ControllerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerFqdn

`func (o *VendorConfigDetailsNode0) SetControllerFqdn(v string)`

SetControllerFqdn sets ControllerFqdn field to given value.

### HasControllerFqdn

`func (o *VendorConfigDetailsNode0) HasControllerFqdn() bool`

HasControllerFqdn returns a boolean if a field has been set.

### GetRootPassword

`func (o *VendorConfigDetailsNode0) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *VendorConfigDetailsNode0) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *VendorConfigDetailsNode0) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *VendorConfigDetailsNode0) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetAdminPassword

`func (o *VendorConfigDetailsNode0) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *VendorConfigDetailsNode0) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *VendorConfigDetailsNode0) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### HasAdminPassword

`func (o *VendorConfigDetailsNode0) HasAdminPassword() bool`

HasAdminPassword returns a boolean if a field has been set.

### GetController1

`func (o *VendorConfigDetailsNode0) GetController1() string`

GetController1 returns the Controller1 field if non-nil, zero value otherwise.

### GetController1Ok

`func (o *VendorConfigDetailsNode0) GetController1Ok() (*string, bool)`

GetController1Ok returns a tuple with the Controller1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController1

`func (o *VendorConfigDetailsNode0) SetController1(v string)`

SetController1 sets Controller1 field to given value.

### HasController1

`func (o *VendorConfigDetailsNode0) HasController1() bool`

HasController1 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


