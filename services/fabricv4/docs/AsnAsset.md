# AsnAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Service** | Pointer to [**AsnAssetService**](AsnAssetService.md) |  | [optional] 
**AnnouncingMetro** | Pointer to **string** |  | [optional] 
**IpBlocks** | Pointer to [**[]AsnAssetIpBlockRef**](AsnAssetIpBlockRef.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewAsnAsset

`func NewAsnAsset() *AsnAsset`

NewAsnAsset instantiates a new AsnAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsnAssetWithDefaults

`func NewAsnAssetWithDefaults() *AsnAsset`

NewAsnAssetWithDefaults instantiates a new AsnAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AsnAsset) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AsnAsset) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AsnAsset) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AsnAsset) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetService

`func (o *AsnAsset) GetService() AsnAssetService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AsnAsset) GetServiceOk() (*AsnAssetService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AsnAsset) SetService(v AsnAssetService)`

SetService sets Service field to given value.

### HasService

`func (o *AsnAsset) HasService() bool`

HasService returns a boolean if a field has been set.

### GetAnnouncingMetro

`func (o *AsnAsset) GetAnnouncingMetro() string`

GetAnnouncingMetro returns the AnnouncingMetro field if non-nil, zero value otherwise.

### GetAnnouncingMetroOk

`func (o *AsnAsset) GetAnnouncingMetroOk() (*string, bool)`

GetAnnouncingMetroOk returns a tuple with the AnnouncingMetro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncingMetro

`func (o *AsnAsset) SetAnnouncingMetro(v string)`

SetAnnouncingMetro sets AnnouncingMetro field to given value.

### HasAnnouncingMetro

`func (o *AsnAsset) HasAnnouncingMetro() bool`

HasAnnouncingMetro returns a boolean if a field has been set.

### GetIpBlocks

`func (o *AsnAsset) GetIpBlocks() []AsnAssetIpBlockRef`

GetIpBlocks returns the IpBlocks field if non-nil, zero value otherwise.

### GetIpBlocksOk

`func (o *AsnAsset) GetIpBlocksOk() (*[]AsnAssetIpBlockRef, bool)`

GetIpBlocksOk returns a tuple with the IpBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlocks

`func (o *AsnAsset) SetIpBlocks(v []AsnAssetIpBlockRef)`

SetIpBlocks sets IpBlocks field to given value.

### HasIpBlocks

`func (o *AsnAsset) HasIpBlocks() bool`

HasIpBlocks returns a boolean if a field has been set.

### GetChangeLog

`func (o *AsnAsset) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *AsnAsset) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *AsnAsset) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *AsnAsset) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


