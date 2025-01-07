# AssetDetailsResponsePayLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetDetails** | Pointer to [**[]AssetDetails**](AssetDetails.md) |  | [optional] 
**TotalCount** | Pointer to **int32** | total number of assets that match the request filters | [optional] 

## Methods

### NewAssetDetailsResponsePayLoad

`func NewAssetDetailsResponsePayLoad() *AssetDetailsResponsePayLoad`

NewAssetDetailsResponsePayLoad instantiates a new AssetDetailsResponsePayLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDetailsResponsePayLoadWithDefaults

`func NewAssetDetailsResponsePayLoadWithDefaults() *AssetDetailsResponsePayLoad`

NewAssetDetailsResponsePayLoadWithDefaults instantiates a new AssetDetailsResponsePayLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetDetails

`func (o *AssetDetailsResponsePayLoad) GetAssetDetails() []AssetDetails`

GetAssetDetails returns the AssetDetails field if non-nil, zero value otherwise.

### GetAssetDetailsOk

`func (o *AssetDetailsResponsePayLoad) GetAssetDetailsOk() (*[]AssetDetails, bool)`

GetAssetDetailsOk returns a tuple with the AssetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetDetails

`func (o *AssetDetailsResponsePayLoad) SetAssetDetails(v []AssetDetails)`

SetAssetDetails sets AssetDetails field to given value.

### HasAssetDetails

`func (o *AssetDetailsResponsePayLoad) HasAssetDetails() bool`

HasAssetDetails returns a boolean if a field has been set.

### GetTotalCount

`func (o *AssetDetailsResponsePayLoad) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AssetDetailsResponsePayLoad) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AssetDetailsResponsePayLoad) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AssetDetailsResponsePayLoad) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


