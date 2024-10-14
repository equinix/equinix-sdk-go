# ConnectionRouteSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**ConnectionRouteEntryFilters**](ConnectionRouteEntryFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]ConnectionRouteSortCriteria**](ConnectionRouteSortCriteria.md) |  | [optional] 

## Methods

### NewConnectionRouteSearchRequest

`func NewConnectionRouteSearchRequest() *ConnectionRouteSearchRequest`

NewConnectionRouteSearchRequest instantiates a new ConnectionRouteSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionRouteSearchRequestWithDefaults

`func NewConnectionRouteSearchRequestWithDefaults() *ConnectionRouteSearchRequest`

NewConnectionRouteSearchRequestWithDefaults instantiates a new ConnectionRouteSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ConnectionRouteSearchRequest) GetFilter() ConnectionRouteEntryFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ConnectionRouteSearchRequest) GetFilterOk() (*ConnectionRouteEntryFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ConnectionRouteSearchRequest) SetFilter(v ConnectionRouteEntryFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ConnectionRouteSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *ConnectionRouteSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ConnectionRouteSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ConnectionRouteSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ConnectionRouteSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *ConnectionRouteSearchRequest) GetSort() []ConnectionRouteSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ConnectionRouteSearchRequest) GetSortOk() (*[]ConnectionRouteSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ConnectionRouteSearchRequest) SetSort(v []ConnectionRouteSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ConnectionRouteSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


