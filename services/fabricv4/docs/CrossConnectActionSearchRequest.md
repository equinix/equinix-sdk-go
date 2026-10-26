# CrossConnectActionSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**CrossConnectActionSearchFilters**](CrossConnectActionSearchFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]CrossConnectActionSortCriteria**](CrossConnectActionSortCriteria.md) |  | [optional] 

## Methods

### NewCrossConnectActionSearchRequest

`func NewCrossConnectActionSearchRequest() *CrossConnectActionSearchRequest`

NewCrossConnectActionSearchRequest instantiates a new CrossConnectActionSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectActionSearchRequestWithDefaults

`func NewCrossConnectActionSearchRequestWithDefaults() *CrossConnectActionSearchRequest`

NewCrossConnectActionSearchRequestWithDefaults instantiates a new CrossConnectActionSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *CrossConnectActionSearchRequest) GetFilter() CrossConnectActionSearchFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CrossConnectActionSearchRequest) GetFilterOk() (*CrossConnectActionSearchFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CrossConnectActionSearchRequest) SetFilter(v CrossConnectActionSearchFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CrossConnectActionSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *CrossConnectActionSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CrossConnectActionSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CrossConnectActionSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *CrossConnectActionSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *CrossConnectActionSearchRequest) GetSort() []CrossConnectActionSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *CrossConnectActionSearchRequest) GetSortOk() (*[]CrossConnectActionSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *CrossConnectActionSearchRequest) SetSort(v []CrossConnectActionSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *CrossConnectActionSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


