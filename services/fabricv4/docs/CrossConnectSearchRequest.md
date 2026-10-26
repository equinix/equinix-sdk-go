# CrossConnectSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**CrossConnectSearchFilters**](CrossConnectSearchFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]CrossConnectSortCriteria**](CrossConnectSortCriteria.md) |  | [optional] 

## Methods

### NewCrossConnectSearchRequest

`func NewCrossConnectSearchRequest() *CrossConnectSearchRequest`

NewCrossConnectSearchRequest instantiates a new CrossConnectSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectSearchRequestWithDefaults

`func NewCrossConnectSearchRequestWithDefaults() *CrossConnectSearchRequest`

NewCrossConnectSearchRequestWithDefaults instantiates a new CrossConnectSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *CrossConnectSearchRequest) GetFilter() CrossConnectSearchFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CrossConnectSearchRequest) GetFilterOk() (*CrossConnectSearchFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CrossConnectSearchRequest) SetFilter(v CrossConnectSearchFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CrossConnectSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *CrossConnectSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CrossConnectSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CrossConnectSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *CrossConnectSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *CrossConnectSearchRequest) GetSort() []CrossConnectSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *CrossConnectSearchRequest) GetSortOk() (*[]CrossConnectSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *CrossConnectSearchRequest) SetSort(v []CrossConnectSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *CrossConnectSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


