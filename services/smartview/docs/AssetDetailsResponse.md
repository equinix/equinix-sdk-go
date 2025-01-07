# AssetDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayLoad** | Pointer to [**AssetDetailsResponsePayLoad**](AssetDetailsResponsePayLoad.md) |  | [optional] 
**Status** | Pointer to [**AssetDetailResponseStatus**](AssetDetailResponseStatus.md) |  | [optional] 

## Methods

### NewAssetDetailsResponse

`func NewAssetDetailsResponse() *AssetDetailsResponse`

NewAssetDetailsResponse instantiates a new AssetDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDetailsResponseWithDefaults

`func NewAssetDetailsResponseWithDefaults() *AssetDetailsResponse`

NewAssetDetailsResponseWithDefaults instantiates a new AssetDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayLoad

`func (o *AssetDetailsResponse) GetPayLoad() AssetDetailsResponsePayLoad`

GetPayLoad returns the PayLoad field if non-nil, zero value otherwise.

### GetPayLoadOk

`func (o *AssetDetailsResponse) GetPayLoadOk() (*AssetDetailsResponsePayLoad, bool)`

GetPayLoadOk returns a tuple with the PayLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayLoad

`func (o *AssetDetailsResponse) SetPayLoad(v AssetDetailsResponsePayLoad)`

SetPayLoad sets PayLoad field to given value.

### HasPayLoad

`func (o *AssetDetailsResponse) HasPayLoad() bool`

HasPayLoad returns a boolean if a field has been set.

### GetStatus

`func (o *AssetDetailsResponse) GetStatus() AssetDetailResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetDetailsResponse) GetStatusOk() (*AssetDetailResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetDetailsResponse) SetStatus(v AssetDetailResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetDetailsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


