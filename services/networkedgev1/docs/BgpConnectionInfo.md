# BgpConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpStatus** | Pointer to **string** |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] 
**Metro** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProviderStatus** | Pointer to **string** |  | [optional] 
**RedundantConnection** | Pointer to [**BgpConnectionInfo**](BgpConnectionInfo.md) |  | [optional] 
**RedundantUUID** | Pointer to **string** |  | [optional] 
**SellerOrganizationName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewBgpConnectionInfo

`func NewBgpConnectionInfo() *BgpConnectionInfo`

NewBgpConnectionInfo instantiates a new BgpConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpConnectionInfoWithDefaults

`func NewBgpConnectionInfoWithDefaults() *BgpConnectionInfo`

NewBgpConnectionInfoWithDefaults instantiates a new BgpConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpStatus

`func (o *BgpConnectionInfo) GetBgpStatus() string`

GetBgpStatus returns the BgpStatus field if non-nil, zero value otherwise.

### GetBgpStatusOk

`func (o *BgpConnectionInfo) GetBgpStatusOk() (*string, bool)`

GetBgpStatusOk returns a tuple with the BgpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpStatus

`func (o *BgpConnectionInfo) SetBgpStatus(v string)`

SetBgpStatus sets BgpStatus field to given value.

### HasBgpStatus

`func (o *BgpConnectionInfo) HasBgpStatus() bool`

HasBgpStatus returns a boolean if a field has been set.

### GetIsPrimary

`func (o *BgpConnectionInfo) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *BgpConnectionInfo) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *BgpConnectionInfo) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *BgpConnectionInfo) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetMetro

`func (o *BgpConnectionInfo) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *BgpConnectionInfo) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *BgpConnectionInfo) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *BgpConnectionInfo) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetName

`func (o *BgpConnectionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BgpConnectionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BgpConnectionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BgpConnectionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProviderStatus

`func (o *BgpConnectionInfo) GetProviderStatus() string`

GetProviderStatus returns the ProviderStatus field if non-nil, zero value otherwise.

### GetProviderStatusOk

`func (o *BgpConnectionInfo) GetProviderStatusOk() (*string, bool)`

GetProviderStatusOk returns a tuple with the ProviderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderStatus

`func (o *BgpConnectionInfo) SetProviderStatus(v string)`

SetProviderStatus sets ProviderStatus field to given value.

### HasProviderStatus

`func (o *BgpConnectionInfo) HasProviderStatus() bool`

HasProviderStatus returns a boolean if a field has been set.

### GetRedundantConnection

`func (o *BgpConnectionInfo) GetRedundantConnection() BgpConnectionInfo`

GetRedundantConnection returns the RedundantConnection field if non-nil, zero value otherwise.

### GetRedundantConnectionOk

`func (o *BgpConnectionInfo) GetRedundantConnectionOk() (*BgpConnectionInfo, bool)`

GetRedundantConnectionOk returns a tuple with the RedundantConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantConnection

`func (o *BgpConnectionInfo) SetRedundantConnection(v BgpConnectionInfo)`

SetRedundantConnection sets RedundantConnection field to given value.

### HasRedundantConnection

`func (o *BgpConnectionInfo) HasRedundantConnection() bool`

HasRedundantConnection returns a boolean if a field has been set.

### GetRedundantUUID

`func (o *BgpConnectionInfo) GetRedundantUUID() string`

GetRedundantUUID returns the RedundantUUID field if non-nil, zero value otherwise.

### GetRedundantUUIDOk

`func (o *BgpConnectionInfo) GetRedundantUUIDOk() (*string, bool)`

GetRedundantUUIDOk returns a tuple with the RedundantUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantUUID

`func (o *BgpConnectionInfo) SetRedundantUUID(v string)`

SetRedundantUUID sets RedundantUUID field to given value.

### HasRedundantUUID

`func (o *BgpConnectionInfo) HasRedundantUUID() bool`

HasRedundantUUID returns a boolean if a field has been set.

### GetSellerOrganizationName

`func (o *BgpConnectionInfo) GetSellerOrganizationName() string`

GetSellerOrganizationName returns the SellerOrganizationName field if non-nil, zero value otherwise.

### GetSellerOrganizationNameOk

`func (o *BgpConnectionInfo) GetSellerOrganizationNameOk() (*string, bool)`

GetSellerOrganizationNameOk returns a tuple with the SellerOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerOrganizationName

`func (o *BgpConnectionInfo) SetSellerOrganizationName(v string)`

SetSellerOrganizationName sets SellerOrganizationName field to given value.

### HasSellerOrganizationName

`func (o *BgpConnectionInfo) HasSellerOrganizationName() bool`

HasSellerOrganizationName returns a boolean if a field has been set.

### GetStatus

`func (o *BgpConnectionInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BgpConnectionInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BgpConnectionInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BgpConnectionInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUuid

`func (o *BgpConnectionInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *BgpConnectionInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *BgpConnectionInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *BgpConnectionInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


