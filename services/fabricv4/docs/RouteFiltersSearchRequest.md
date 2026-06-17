# RouteFiltersSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**SearchFilter**](SearchFilter.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]RouteFilterSortCriteria**](RouteFilterSortCriteria.md) |  | [optional] 

## Methods

### NewRouteFiltersSearchRequest

`func NewRouteFiltersSearchRequest() *RouteFiltersSearchRequest`

NewRouteFiltersSearchRequest instantiates a new RouteFiltersSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFiltersSearchRequestWithDefaults

`func NewRouteFiltersSearchRequestWithDefaults() *RouteFiltersSearchRequest`

NewRouteFiltersSearchRequestWithDefaults instantiates a new RouteFiltersSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *RouteFiltersSearchRequest) GetFilter() SearchFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RouteFiltersSearchRequest) GetFilterOk() (*SearchFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RouteFiltersSearchRequest) SetFilter(v SearchFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RouteFiltersSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *RouteFiltersSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteFiltersSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteFiltersSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RouteFiltersSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *RouteFiltersSearchRequest) GetSort() []RouteFilterSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RouteFiltersSearchRequest) GetSortOk() (*[]RouteFilterSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RouteFiltersSearchRequest) SetSort(v []RouteFilterSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RouteFiltersSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


