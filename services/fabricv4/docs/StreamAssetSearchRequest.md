# StreamAssetSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**StreamAssetFilters**](StreamAssetFilters.md) |  | 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]StreamAssetSortCriteria**](StreamAssetSortCriteria.md) |  | [optional] 

## Methods

### NewStreamAssetSearchRequest

`func NewStreamAssetSearchRequest(filter StreamAssetFilters, ) *StreamAssetSearchRequest`

NewStreamAssetSearchRequest instantiates a new StreamAssetSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamAssetSearchRequestWithDefaults

`func NewStreamAssetSearchRequestWithDefaults() *StreamAssetSearchRequest`

NewStreamAssetSearchRequestWithDefaults instantiates a new StreamAssetSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *StreamAssetSearchRequest) GetFilter() StreamAssetFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *StreamAssetSearchRequest) GetFilterOk() (*StreamAssetFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *StreamAssetSearchRequest) SetFilter(v StreamAssetFilters)`

SetFilter sets Filter field to given value.


### GetPagination

`func (o *StreamAssetSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *StreamAssetSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *StreamAssetSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *StreamAssetSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *StreamAssetSearchRequest) GetSort() []StreamAssetSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *StreamAssetSearchRequest) GetSortOk() (*[]StreamAssetSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *StreamAssetSearchRequest) SetSort(v []StreamAssetSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *StreamAssetSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


