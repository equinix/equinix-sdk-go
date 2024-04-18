# LicenseOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The license name | [optional] 
**Type** | Pointer to **string** | The license type | [optional] 
**MetroCodes** | Pointer to **[]string** | The metros where the license is available | [optional] 

## Methods

### NewLicenseOptions

`func NewLicenseOptions() *LicenseOptions`

NewLicenseOptions instantiates a new LicenseOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseOptionsWithDefaults

`func NewLicenseOptionsWithDefaults() *LicenseOptions`

NewLicenseOptionsWithDefaults instantiates a new LicenseOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LicenseOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LicenseOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LicenseOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LicenseOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *LicenseOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LicenseOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LicenseOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LicenseOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetroCodes

`func (o *LicenseOptions) GetMetroCodes() []string`

GetMetroCodes returns the MetroCodes field if non-nil, zero value otherwise.

### GetMetroCodesOk

`func (o *LicenseOptions) GetMetroCodesOk() (*[]string, bool)`

GetMetroCodesOk returns a tuple with the MetroCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCodes

`func (o *LicenseOptions) SetMetroCodes(v []string)`

SetMetroCodes sets MetroCodes field to given value.

### HasMetroCodes

`func (o *LicenseOptions) HasMetroCodes() bool`

HasMetroCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


