# AssetsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayLoad** | Pointer to [**AssetsListPayLoad**](AssetsListPayLoad.md) |  | [optional] 
**Status** | Pointer to [**AssetsListStatus**](AssetsListStatus.md) |  | [optional] 

## Methods

### NewAssetsList

`func NewAssetsList() *AssetsList`

NewAssetsList instantiates a new AssetsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsListWithDefaults

`func NewAssetsListWithDefaults() *AssetsList`

NewAssetsListWithDefaults instantiates a new AssetsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayLoad

`func (o *AssetsList) GetPayLoad() AssetsListPayLoad`

GetPayLoad returns the PayLoad field if non-nil, zero value otherwise.

### GetPayLoadOk

`func (o *AssetsList) GetPayLoadOk() (*AssetsListPayLoad, bool)`

GetPayLoadOk returns a tuple with the PayLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayLoad

`func (o *AssetsList) SetPayLoad(v AssetsListPayLoad)`

SetPayLoad sets PayLoad field to given value.

### HasPayLoad

`func (o *AssetsList) HasPayLoad() bool`

HasPayLoad returns a boolean if a field has been set.

### GetStatus

`func (o *AssetsList) GetStatus() AssetsListStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetsList) GetStatusOk() (*AssetsListStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetsList) SetStatus(v AssetsListStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetsList) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


