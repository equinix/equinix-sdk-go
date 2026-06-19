# RouteAggregationsSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**SearchFilter**](SearchFilter.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]RouteAggregationSortCriteria**](RouteAggregationSortCriteria.md) |  | [optional] 

## Methods

### NewRouteAggregationsSearchRequest

`func NewRouteAggregationsSearchRequest() *RouteAggregationsSearchRequest`

NewRouteAggregationsSearchRequest instantiates a new RouteAggregationsSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationsSearchRequestWithDefaults

`func NewRouteAggregationsSearchRequestWithDefaults() *RouteAggregationsSearchRequest`

NewRouteAggregationsSearchRequestWithDefaults instantiates a new RouteAggregationsSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *RouteAggregationsSearchRequest) GetFilter() SearchFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RouteAggregationsSearchRequest) GetFilterOk() (*SearchFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RouteAggregationsSearchRequest) SetFilter(v SearchFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RouteAggregationsSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *RouteAggregationsSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteAggregationsSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteAggregationsSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RouteAggregationsSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *RouteAggregationsSearchRequest) GetSort() []RouteAggregationSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RouteAggregationsSearchRequest) GetSortOk() (*[]RouteAggregationSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RouteAggregationsSearchRequest) SetSort(v []RouteAggregationSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RouteAggregationsSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


