# SearchStreamAssetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]StreamSearchAsset**](StreamSearchAsset.md) | Data returned from the API call. | [optional] 

## Methods

### NewSearchStreamAssetResponse

`func NewSearchStreamAssetResponse() *SearchStreamAssetResponse`

NewSearchStreamAssetResponse instantiates a new SearchStreamAssetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchStreamAssetResponseWithDefaults

`func NewSearchStreamAssetResponseWithDefaults() *SearchStreamAssetResponse`

NewSearchStreamAssetResponseWithDefaults instantiates a new SearchStreamAssetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *SearchStreamAssetResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SearchStreamAssetResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SearchStreamAssetResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *SearchStreamAssetResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *SearchStreamAssetResponse) GetData() []StreamSearchAsset`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SearchStreamAssetResponse) GetDataOk() (*[]StreamSearchAsset, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SearchStreamAssetResponse) SetData(v []StreamSearchAsset)`

SetData sets Data field to given value.

### HasData

`func (o *SearchStreamAssetResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


