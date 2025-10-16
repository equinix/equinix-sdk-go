# PortPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Port Package URI | [optional] [readonly] 
**Type** | [**PortPackageType**](PortPackageType.md) |  | 
**Code** | **string** | Port Package code | 
**VcBandwidthMax** | Pointer to **int32** | Maximum virtual connection bandwidth in Mbps | [optional] 
**VcRemoteSupported** | Pointer to **bool** | Indicates if remote virtual connections are supported | [optional] 
**SupportedServiceTypes** | Pointer to [**[]PortPackageSupportedServiceTypesInner**](PortPackageSupportedServiceTypesInner.md) | List of supported service types | [optional] 
**SupportedSourceTypes** | Pointer to [**[]PortPackageSourceType**](PortPackageSourceType.md) | List of supported source types | [optional] 
**SupportedMetros** | Pointer to **[]string** | List of supported metros | [optional] 

## Methods

### NewPortPackage

`func NewPortPackage(type_ PortPackageType, code string, ) *PortPackage`

NewPortPackage instantiates a new PortPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortPackageWithDefaults

`func NewPortPackageWithDefaults() *PortPackage`

NewPortPackageWithDefaults instantiates a new PortPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PortPackage) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PortPackage) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PortPackage) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PortPackage) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *PortPackage) GetType() PortPackageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortPackage) GetTypeOk() (*PortPackageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortPackage) SetType(v PortPackageType)`

SetType sets Type field to given value.


### GetCode

`func (o *PortPackage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PortPackage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PortPackage) SetCode(v string)`

SetCode sets Code field to given value.


### GetVcBandwidthMax

`func (o *PortPackage) GetVcBandwidthMax() int32`

GetVcBandwidthMax returns the VcBandwidthMax field if non-nil, zero value otherwise.

### GetVcBandwidthMaxOk

`func (o *PortPackage) GetVcBandwidthMaxOk() (*int32, bool)`

GetVcBandwidthMaxOk returns a tuple with the VcBandwidthMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcBandwidthMax

`func (o *PortPackage) SetVcBandwidthMax(v int32)`

SetVcBandwidthMax sets VcBandwidthMax field to given value.

### HasVcBandwidthMax

`func (o *PortPackage) HasVcBandwidthMax() bool`

HasVcBandwidthMax returns a boolean if a field has been set.

### GetVcRemoteSupported

`func (o *PortPackage) GetVcRemoteSupported() bool`

GetVcRemoteSupported returns the VcRemoteSupported field if non-nil, zero value otherwise.

### GetVcRemoteSupportedOk

`func (o *PortPackage) GetVcRemoteSupportedOk() (*bool, bool)`

GetVcRemoteSupportedOk returns a tuple with the VcRemoteSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcRemoteSupported

`func (o *PortPackage) SetVcRemoteSupported(v bool)`

SetVcRemoteSupported sets VcRemoteSupported field to given value.

### HasVcRemoteSupported

`func (o *PortPackage) HasVcRemoteSupported() bool`

HasVcRemoteSupported returns a boolean if a field has been set.

### GetSupportedServiceTypes

`func (o *PortPackage) GetSupportedServiceTypes() []PortPackageSupportedServiceTypesInner`

GetSupportedServiceTypes returns the SupportedServiceTypes field if non-nil, zero value otherwise.

### GetSupportedServiceTypesOk

`func (o *PortPackage) GetSupportedServiceTypesOk() (*[]PortPackageSupportedServiceTypesInner, bool)`

GetSupportedServiceTypesOk returns a tuple with the SupportedServiceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedServiceTypes

`func (o *PortPackage) SetSupportedServiceTypes(v []PortPackageSupportedServiceTypesInner)`

SetSupportedServiceTypes sets SupportedServiceTypes field to given value.

### HasSupportedServiceTypes

`func (o *PortPackage) HasSupportedServiceTypes() bool`

HasSupportedServiceTypes returns a boolean if a field has been set.

### GetSupportedSourceTypes

`func (o *PortPackage) GetSupportedSourceTypes() []PortPackageSourceType`

GetSupportedSourceTypes returns the SupportedSourceTypes field if non-nil, zero value otherwise.

### GetSupportedSourceTypesOk

`func (o *PortPackage) GetSupportedSourceTypesOk() (*[]PortPackageSourceType, bool)`

GetSupportedSourceTypesOk returns a tuple with the SupportedSourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSourceTypes

`func (o *PortPackage) SetSupportedSourceTypes(v []PortPackageSourceType)`

SetSupportedSourceTypes sets SupportedSourceTypes field to given value.

### HasSupportedSourceTypes

`func (o *PortPackage) HasSupportedSourceTypes() bool`

HasSupportedSourceTypes returns a boolean if a field has been set.

### GetSupportedMetros

`func (o *PortPackage) GetSupportedMetros() []string`

GetSupportedMetros returns the SupportedMetros field if non-nil, zero value otherwise.

### GetSupportedMetrosOk

`func (o *PortPackage) GetSupportedMetrosOk() (*[]string, bool)`

GetSupportedMetrosOk returns a tuple with the SupportedMetros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMetros

`func (o *PortPackage) SetSupportedMetros(v []string)`

SetSupportedMetros sets SupportedMetros field to given value.

### HasSupportedMetros

`func (o *PortPackage) HasSupportedMetros() bool`

HasSupportedMetros returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


