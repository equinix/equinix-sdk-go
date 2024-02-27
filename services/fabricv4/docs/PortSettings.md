# PortSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** | Product name | [optional] 
**Buyout** | Pointer to **bool** |  | [optional] 
**ViewPortPermission** | Pointer to **bool** |  | [optional] 
**PlaceVcOrderPermission** | Pointer to **bool** |  | [optional] 
**Layer3Enabled** | Pointer to **bool** |  | [optional] 
**ProductCode** | Pointer to **string** |  | [optional] 
**SharedPortType** | Pointer to **bool** |  | [optional] 
**SharedPortProduct** | Pointer to [**PortSettingsSharedPortProduct**](PortSettingsSharedPortProduct.md) |  | [optional] 
**PackageType** | Pointer to [**PortSettingsPackageType**](PortSettingsPackageType.md) |  | [optional] 

## Methods

### NewPortSettings

`func NewPortSettings() *PortSettings`

NewPortSettings instantiates a new PortSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortSettingsWithDefaults

`func NewPortSettingsWithDefaults() *PortSettings`

NewPortSettingsWithDefaults instantiates a new PortSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *PortSettings) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PortSettings) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PortSettings) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PortSettings) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetBuyout

`func (o *PortSettings) GetBuyout() bool`

GetBuyout returns the Buyout field if non-nil, zero value otherwise.

### GetBuyoutOk

`func (o *PortSettings) GetBuyoutOk() (*bool, bool)`

GetBuyoutOk returns a tuple with the Buyout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyout

`func (o *PortSettings) SetBuyout(v bool)`

SetBuyout sets Buyout field to given value.

### HasBuyout

`func (o *PortSettings) HasBuyout() bool`

HasBuyout returns a boolean if a field has been set.

### GetViewPortPermission

`func (o *PortSettings) GetViewPortPermission() bool`

GetViewPortPermission returns the ViewPortPermission field if non-nil, zero value otherwise.

### GetViewPortPermissionOk

`func (o *PortSettings) GetViewPortPermissionOk() (*bool, bool)`

GetViewPortPermissionOk returns a tuple with the ViewPortPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPortPermission

`func (o *PortSettings) SetViewPortPermission(v bool)`

SetViewPortPermission sets ViewPortPermission field to given value.

### HasViewPortPermission

`func (o *PortSettings) HasViewPortPermission() bool`

HasViewPortPermission returns a boolean if a field has been set.

### GetPlaceVcOrderPermission

`func (o *PortSettings) GetPlaceVcOrderPermission() bool`

GetPlaceVcOrderPermission returns the PlaceVcOrderPermission field if non-nil, zero value otherwise.

### GetPlaceVcOrderPermissionOk

`func (o *PortSettings) GetPlaceVcOrderPermissionOk() (*bool, bool)`

GetPlaceVcOrderPermissionOk returns a tuple with the PlaceVcOrderPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceVcOrderPermission

`func (o *PortSettings) SetPlaceVcOrderPermission(v bool)`

SetPlaceVcOrderPermission sets PlaceVcOrderPermission field to given value.

### HasPlaceVcOrderPermission

`func (o *PortSettings) HasPlaceVcOrderPermission() bool`

HasPlaceVcOrderPermission returns a boolean if a field has been set.

### GetLayer3Enabled

`func (o *PortSettings) GetLayer3Enabled() bool`

GetLayer3Enabled returns the Layer3Enabled field if non-nil, zero value otherwise.

### GetLayer3EnabledOk

`func (o *PortSettings) GetLayer3EnabledOk() (*bool, bool)`

GetLayer3EnabledOk returns a tuple with the Layer3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer3Enabled

`func (o *PortSettings) SetLayer3Enabled(v bool)`

SetLayer3Enabled sets Layer3Enabled field to given value.

### HasLayer3Enabled

`func (o *PortSettings) HasLayer3Enabled() bool`

HasLayer3Enabled returns a boolean if a field has been set.

### GetProductCode

`func (o *PortSettings) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *PortSettings) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *PortSettings) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *PortSettings) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetSharedPortType

`func (o *PortSettings) GetSharedPortType() bool`

GetSharedPortType returns the SharedPortType field if non-nil, zero value otherwise.

### GetSharedPortTypeOk

`func (o *PortSettings) GetSharedPortTypeOk() (*bool, bool)`

GetSharedPortTypeOk returns a tuple with the SharedPortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedPortType

`func (o *PortSettings) SetSharedPortType(v bool)`

SetSharedPortType sets SharedPortType field to given value.

### HasSharedPortType

`func (o *PortSettings) HasSharedPortType() bool`

HasSharedPortType returns a boolean if a field has been set.

### GetSharedPortProduct

`func (o *PortSettings) GetSharedPortProduct() PortSettingsSharedPortProduct`

GetSharedPortProduct returns the SharedPortProduct field if non-nil, zero value otherwise.

### GetSharedPortProductOk

`func (o *PortSettings) GetSharedPortProductOk() (*PortSettingsSharedPortProduct, bool)`

GetSharedPortProductOk returns a tuple with the SharedPortProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedPortProduct

`func (o *PortSettings) SetSharedPortProduct(v PortSettingsSharedPortProduct)`

SetSharedPortProduct sets SharedPortProduct field to given value.

### HasSharedPortProduct

`func (o *PortSettings) HasSharedPortProduct() bool`

HasSharedPortProduct returns a boolean if a field has been set.

### GetPackageType

`func (o *PortSettings) GetPackageType() PortSettingsPackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *PortSettings) GetPackageTypeOk() (*PortSettingsPackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *PortSettings) SetPackageType(v PortSettingsPackageType)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *PortSettings) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


