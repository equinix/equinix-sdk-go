# FabricCloudRouterPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Unique identifier assigned to the Cloud Router | [optional] 
**Location** | Pointer to [**PriceLocation**](PriceLocation.md) |  | [optional] 
**Package** | Pointer to [**FabricCloudRouterPackages**](FabricCloudRouterPackages.md) |  | [optional] 

## Methods

### NewFabricCloudRouterPrice

`func NewFabricCloudRouterPrice() *FabricCloudRouterPrice`

NewFabricCloudRouterPrice instantiates a new FabricCloudRouterPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricCloudRouterPriceWithDefaults

`func NewFabricCloudRouterPriceWithDefaults() *FabricCloudRouterPrice`

NewFabricCloudRouterPriceWithDefaults instantiates a new FabricCloudRouterPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *FabricCloudRouterPrice) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FabricCloudRouterPrice) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FabricCloudRouterPrice) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FabricCloudRouterPrice) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetLocation

`func (o *FabricCloudRouterPrice) GetLocation() PriceLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FabricCloudRouterPrice) GetLocationOk() (*PriceLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FabricCloudRouterPrice) SetLocation(v PriceLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FabricCloudRouterPrice) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPackage

`func (o *FabricCloudRouterPrice) GetPackage() FabricCloudRouterPackages`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *FabricCloudRouterPrice) GetPackageOk() (*FabricCloudRouterPackages, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *FabricCloudRouterPrice) SetPackage(v FabricCloudRouterPackages)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *FabricCloudRouterPrice) HasPackage() bool`

HasPackage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


