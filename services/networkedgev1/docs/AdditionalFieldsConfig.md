# AdditionalFieldsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of field. | [optional] 
**Required** | Pointer to **bool** | Whether or not the field is required at the time of device creation. | [optional] 
**IsSameValueAllowedForPrimaryAndSecondary** | Pointer to **bool** | Whether or not you need two distinct values for primary and secondary devices at the time of device creation. This field is only useful for HA devices. | [optional] 

## Methods

### NewAdditionalFieldsConfig

`func NewAdditionalFieldsConfig() *AdditionalFieldsConfig`

NewAdditionalFieldsConfig instantiates a new AdditionalFieldsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalFieldsConfigWithDefaults

`func NewAdditionalFieldsConfigWithDefaults() *AdditionalFieldsConfig`

NewAdditionalFieldsConfigWithDefaults instantiates a new AdditionalFieldsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdditionalFieldsConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdditionalFieldsConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdditionalFieldsConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdditionalFieldsConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *AdditionalFieldsConfig) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AdditionalFieldsConfig) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AdditionalFieldsConfig) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *AdditionalFieldsConfig) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetIsSameValueAllowedForPrimaryAndSecondary

`func (o *AdditionalFieldsConfig) GetIsSameValueAllowedForPrimaryAndSecondary() bool`

GetIsSameValueAllowedForPrimaryAndSecondary returns the IsSameValueAllowedForPrimaryAndSecondary field if non-nil, zero value otherwise.

### GetIsSameValueAllowedForPrimaryAndSecondaryOk

`func (o *AdditionalFieldsConfig) GetIsSameValueAllowedForPrimaryAndSecondaryOk() (*bool, bool)`

GetIsSameValueAllowedForPrimaryAndSecondaryOk returns a tuple with the IsSameValueAllowedForPrimaryAndSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSameValueAllowedForPrimaryAndSecondary

`func (o *AdditionalFieldsConfig) SetIsSameValueAllowedForPrimaryAndSecondary(v bool)`

SetIsSameValueAllowedForPrimaryAndSecondary sets IsSameValueAllowedForPrimaryAndSecondary field to given value.

### HasIsSameValueAllowedForPrimaryAndSecondary

`func (o *AdditionalFieldsConfig) HasIsSameValueAllowedForPrimaryAndSecondary() bool`

HasIsSameValueAllowedForPrimaryAndSecondary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


