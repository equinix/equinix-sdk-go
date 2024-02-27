# RouteTableEntrySearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**RouteTableEntryFilters**](RouteTableEntryFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]RouteTableEntrySortCriteria**](RouteTableEntrySortCriteria.md) |  | [optional] 

## Methods

### NewRouteTableEntrySearchRequest

`func NewRouteTableEntrySearchRequest() *RouteTableEntrySearchRequest`

NewRouteTableEntrySearchRequest instantiates a new RouteTableEntrySearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteTableEntrySearchRequestWithDefaults

`func NewRouteTableEntrySearchRequestWithDefaults() *RouteTableEntrySearchRequest`

NewRouteTableEntrySearchRequestWithDefaults instantiates a new RouteTableEntrySearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *RouteTableEntrySearchRequest) GetFilter() RouteTableEntryFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RouteTableEntrySearchRequest) GetFilterOk() (*RouteTableEntryFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RouteTableEntrySearchRequest) SetFilter(v RouteTableEntryFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RouteTableEntrySearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *RouteTableEntrySearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteTableEntrySearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteTableEntrySearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RouteTableEntrySearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *RouteTableEntrySearchRequest) GetSort() []RouteTableEntrySortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RouteTableEntrySearchRequest) GetSortOk() (*[]RouteTableEntrySortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RouteTableEntrySearchRequest) SetSort(v []RouteTableEntrySortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RouteTableEntrySearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


