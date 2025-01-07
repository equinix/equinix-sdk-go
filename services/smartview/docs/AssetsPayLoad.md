# AssetsPayLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetsList** | Pointer to [**[]AssetsArray**](AssetsArray.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewAssetsPayLoad

`func NewAssetsPayLoad() *AssetsPayLoad`

NewAssetsPayLoad instantiates a new AssetsPayLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsPayLoadWithDefaults

`func NewAssetsPayLoadWithDefaults() *AssetsPayLoad`

NewAssetsPayLoadWithDefaults instantiates a new AssetsPayLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetsList

`func (o *AssetsPayLoad) GetAssetsList() []AssetsArray`

GetAssetsList returns the AssetsList field if non-nil, zero value otherwise.

### GetAssetsListOk

`func (o *AssetsPayLoad) GetAssetsListOk() (*[]AssetsArray, bool)`

GetAssetsListOk returns a tuple with the AssetsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetsList

`func (o *AssetsPayLoad) SetAssetsList(v []AssetsArray)`

SetAssetsList sets AssetsList field to given value.

### HasAssetsList

`func (o *AssetsPayLoad) HasAssetsList() bool`

HasAssetsList returns a boolean if a field has been set.

### GetTotalCount

`func (o *AssetsPayLoad) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AssetsPayLoad) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AssetsPayLoad) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AssetsPayLoad) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


