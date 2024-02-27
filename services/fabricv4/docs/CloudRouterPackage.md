# CloudRouterPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Cloud Router package URI | [optional] [readonly] 
**Type** | Pointer to [**CloudRouterPackageType**](CloudRouterPackageType.md) |  | [optional] 
**Code** | Pointer to [**Code**](Code.md) |  | [optional] 
**Description** | Pointer to **string** | Fabric Cloud Router Package description | [optional] 
**TotalIPv4RoutesMax** | Pointer to **int32** | Cloud Router package BGP IPv4 routes limit | [optional] 
**TotalIPv6RoutesMax** | Pointer to **int32** | Cloud Router package BGP IPv6 routes limit | [optional] 
**StaticIPv4RoutesMax** | Pointer to **int32** | CloudRouter package static IPv4 routes limit | [optional] 
**StaticIPv6RoutesMax** | Pointer to **int32** | CloudRouter package static IPv6 routes limit | [optional] 
**NaclsMax** | Pointer to **int32** | CloudRouter package NACLs limit | [optional] 
**NaclRulesMax** | Pointer to **int32** | CloudRouter package NACLs rules limit | [optional] 
**HaSupported** | Pointer to **bool** | CloudRouter package high-available configuration support | [optional] 
**RouteFilterSupported** | Pointer to **bool** | CloudRouter package route filter support | [optional] 
**NatType** | Pointer to [**CloudRouterPackageNatType**](CloudRouterPackageNatType.md) |  | [optional] 
**VcCountMax** | Pointer to **int32** | CloudRouter package Max Connection limit | [optional] 
**CrCountMax** | Pointer to **int32** | CloudRouter package Max CloudRouter limit | [optional] 
**VcBandwidthMax** | Pointer to **int32** | CloudRouter package Max Bandwidth limit | [optional] 
**ChangeLog** | Pointer to [**PackageChangeLog**](PackageChangeLog.md) |  | [optional] 

## Methods

### NewCloudRouterPackage

`func NewCloudRouterPackage() *CloudRouterPackage`

NewCloudRouterPackage instantiates a new CloudRouterPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterPackageWithDefaults

`func NewCloudRouterPackageWithDefaults() *CloudRouterPackage`

NewCloudRouterPackageWithDefaults instantiates a new CloudRouterPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CloudRouterPackage) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CloudRouterPackage) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CloudRouterPackage) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CloudRouterPackage) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *CloudRouterPackage) GetType() CloudRouterPackageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudRouterPackage) GetTypeOk() (*CloudRouterPackageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudRouterPackage) SetType(v CloudRouterPackageType)`

SetType sets Type field to given value.

### HasType

`func (o *CloudRouterPackage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *CloudRouterPackage) GetCode() Code`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CloudRouterPackage) GetCodeOk() (*Code, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CloudRouterPackage) SetCode(v Code)`

SetCode sets Code field to given value.

### HasCode

`func (o *CloudRouterPackage) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *CloudRouterPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudRouterPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudRouterPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudRouterPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTotalIPv4RoutesMax

`func (o *CloudRouterPackage) GetTotalIPv4RoutesMax() int32`

GetTotalIPv4RoutesMax returns the TotalIPv4RoutesMax field if non-nil, zero value otherwise.

### GetTotalIPv4RoutesMaxOk

`func (o *CloudRouterPackage) GetTotalIPv4RoutesMaxOk() (*int32, bool)`

GetTotalIPv4RoutesMaxOk returns a tuple with the TotalIPv4RoutesMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIPv4RoutesMax

`func (o *CloudRouterPackage) SetTotalIPv4RoutesMax(v int32)`

SetTotalIPv4RoutesMax sets TotalIPv4RoutesMax field to given value.

### HasTotalIPv4RoutesMax

`func (o *CloudRouterPackage) HasTotalIPv4RoutesMax() bool`

HasTotalIPv4RoutesMax returns a boolean if a field has been set.

### GetTotalIPv6RoutesMax

`func (o *CloudRouterPackage) GetTotalIPv6RoutesMax() int32`

GetTotalIPv6RoutesMax returns the TotalIPv6RoutesMax field if non-nil, zero value otherwise.

### GetTotalIPv6RoutesMaxOk

`func (o *CloudRouterPackage) GetTotalIPv6RoutesMaxOk() (*int32, bool)`

GetTotalIPv6RoutesMaxOk returns a tuple with the TotalIPv6RoutesMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIPv6RoutesMax

`func (o *CloudRouterPackage) SetTotalIPv6RoutesMax(v int32)`

SetTotalIPv6RoutesMax sets TotalIPv6RoutesMax field to given value.

### HasTotalIPv6RoutesMax

`func (o *CloudRouterPackage) HasTotalIPv6RoutesMax() bool`

HasTotalIPv6RoutesMax returns a boolean if a field has been set.

### GetStaticIPv4RoutesMax

`func (o *CloudRouterPackage) GetStaticIPv4RoutesMax() int32`

GetStaticIPv4RoutesMax returns the StaticIPv4RoutesMax field if non-nil, zero value otherwise.

### GetStaticIPv4RoutesMaxOk

`func (o *CloudRouterPackage) GetStaticIPv4RoutesMaxOk() (*int32, bool)`

GetStaticIPv4RoutesMaxOk returns a tuple with the StaticIPv4RoutesMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIPv4RoutesMax

`func (o *CloudRouterPackage) SetStaticIPv4RoutesMax(v int32)`

SetStaticIPv4RoutesMax sets StaticIPv4RoutesMax field to given value.

### HasStaticIPv4RoutesMax

`func (o *CloudRouterPackage) HasStaticIPv4RoutesMax() bool`

HasStaticIPv4RoutesMax returns a boolean if a field has been set.

### GetStaticIPv6RoutesMax

`func (o *CloudRouterPackage) GetStaticIPv6RoutesMax() int32`

GetStaticIPv6RoutesMax returns the StaticIPv6RoutesMax field if non-nil, zero value otherwise.

### GetStaticIPv6RoutesMaxOk

`func (o *CloudRouterPackage) GetStaticIPv6RoutesMaxOk() (*int32, bool)`

GetStaticIPv6RoutesMaxOk returns a tuple with the StaticIPv6RoutesMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIPv6RoutesMax

`func (o *CloudRouterPackage) SetStaticIPv6RoutesMax(v int32)`

SetStaticIPv6RoutesMax sets StaticIPv6RoutesMax field to given value.

### HasStaticIPv6RoutesMax

`func (o *CloudRouterPackage) HasStaticIPv6RoutesMax() bool`

HasStaticIPv6RoutesMax returns a boolean if a field has been set.

### GetNaclsMax

`func (o *CloudRouterPackage) GetNaclsMax() int32`

GetNaclsMax returns the NaclsMax field if non-nil, zero value otherwise.

### GetNaclsMaxOk

`func (o *CloudRouterPackage) GetNaclsMaxOk() (*int32, bool)`

GetNaclsMaxOk returns a tuple with the NaclsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaclsMax

`func (o *CloudRouterPackage) SetNaclsMax(v int32)`

SetNaclsMax sets NaclsMax field to given value.

### HasNaclsMax

`func (o *CloudRouterPackage) HasNaclsMax() bool`

HasNaclsMax returns a boolean if a field has been set.

### GetNaclRulesMax

`func (o *CloudRouterPackage) GetNaclRulesMax() int32`

GetNaclRulesMax returns the NaclRulesMax field if non-nil, zero value otherwise.

### GetNaclRulesMaxOk

`func (o *CloudRouterPackage) GetNaclRulesMaxOk() (*int32, bool)`

GetNaclRulesMaxOk returns a tuple with the NaclRulesMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaclRulesMax

`func (o *CloudRouterPackage) SetNaclRulesMax(v int32)`

SetNaclRulesMax sets NaclRulesMax field to given value.

### HasNaclRulesMax

`func (o *CloudRouterPackage) HasNaclRulesMax() bool`

HasNaclRulesMax returns a boolean if a field has been set.

### GetHaSupported

`func (o *CloudRouterPackage) GetHaSupported() bool`

GetHaSupported returns the HaSupported field if non-nil, zero value otherwise.

### GetHaSupportedOk

`func (o *CloudRouterPackage) GetHaSupportedOk() (*bool, bool)`

GetHaSupportedOk returns a tuple with the HaSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaSupported

`func (o *CloudRouterPackage) SetHaSupported(v bool)`

SetHaSupported sets HaSupported field to given value.

### HasHaSupported

`func (o *CloudRouterPackage) HasHaSupported() bool`

HasHaSupported returns a boolean if a field has been set.

### GetRouteFilterSupported

`func (o *CloudRouterPackage) GetRouteFilterSupported() bool`

GetRouteFilterSupported returns the RouteFilterSupported field if non-nil, zero value otherwise.

### GetRouteFilterSupportedOk

`func (o *CloudRouterPackage) GetRouteFilterSupportedOk() (*bool, bool)`

GetRouteFilterSupportedOk returns a tuple with the RouteFilterSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilterSupported

`func (o *CloudRouterPackage) SetRouteFilterSupported(v bool)`

SetRouteFilterSupported sets RouteFilterSupported field to given value.

### HasRouteFilterSupported

`func (o *CloudRouterPackage) HasRouteFilterSupported() bool`

HasRouteFilterSupported returns a boolean if a field has been set.

### GetNatType

`func (o *CloudRouterPackage) GetNatType() CloudRouterPackageNatType`

GetNatType returns the NatType field if non-nil, zero value otherwise.

### GetNatTypeOk

`func (o *CloudRouterPackage) GetNatTypeOk() (*CloudRouterPackageNatType, bool)`

GetNatTypeOk returns a tuple with the NatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatType

`func (o *CloudRouterPackage) SetNatType(v CloudRouterPackageNatType)`

SetNatType sets NatType field to given value.

### HasNatType

`func (o *CloudRouterPackage) HasNatType() bool`

HasNatType returns a boolean if a field has been set.

### GetVcCountMax

`func (o *CloudRouterPackage) GetVcCountMax() int32`

GetVcCountMax returns the VcCountMax field if non-nil, zero value otherwise.

### GetVcCountMaxOk

`func (o *CloudRouterPackage) GetVcCountMaxOk() (*int32, bool)`

GetVcCountMaxOk returns a tuple with the VcCountMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcCountMax

`func (o *CloudRouterPackage) SetVcCountMax(v int32)`

SetVcCountMax sets VcCountMax field to given value.

### HasVcCountMax

`func (o *CloudRouterPackage) HasVcCountMax() bool`

HasVcCountMax returns a boolean if a field has been set.

### GetCrCountMax

`func (o *CloudRouterPackage) GetCrCountMax() int32`

GetCrCountMax returns the CrCountMax field if non-nil, zero value otherwise.

### GetCrCountMaxOk

`func (o *CloudRouterPackage) GetCrCountMaxOk() (*int32, bool)`

GetCrCountMaxOk returns a tuple with the CrCountMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrCountMax

`func (o *CloudRouterPackage) SetCrCountMax(v int32)`

SetCrCountMax sets CrCountMax field to given value.

### HasCrCountMax

`func (o *CloudRouterPackage) HasCrCountMax() bool`

HasCrCountMax returns a boolean if a field has been set.

### GetVcBandwidthMax

`func (o *CloudRouterPackage) GetVcBandwidthMax() int32`

GetVcBandwidthMax returns the VcBandwidthMax field if non-nil, zero value otherwise.

### GetVcBandwidthMaxOk

`func (o *CloudRouterPackage) GetVcBandwidthMaxOk() (*int32, bool)`

GetVcBandwidthMaxOk returns a tuple with the VcBandwidthMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcBandwidthMax

`func (o *CloudRouterPackage) SetVcBandwidthMax(v int32)`

SetVcBandwidthMax sets VcBandwidthMax field to given value.

### HasVcBandwidthMax

`func (o *CloudRouterPackage) HasVcBandwidthMax() bool`

HasVcBandwidthMax returns a boolean if a field has been set.

### GetChangeLog

`func (o *CloudRouterPackage) GetChangeLog() PackageChangeLog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *CloudRouterPackage) GetChangeLogOk() (*PackageChangeLog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *CloudRouterPackage) SetChangeLog(v PackageChangeLog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *CloudRouterPackage) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


