# StreamSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**StreamSearchFilters**](StreamSearchFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]StreamSearchSortCriteria**](StreamSearchSortCriteria.md) |  | [optional] 

## Methods

### NewStreamSearchRequest

`func NewStreamSearchRequest() *StreamSearchRequest`

NewStreamSearchRequest instantiates a new StreamSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSearchRequestWithDefaults

`func NewStreamSearchRequestWithDefaults() *StreamSearchRequest`

NewStreamSearchRequestWithDefaults instantiates a new StreamSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *StreamSearchRequest) GetFilter() StreamSearchFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *StreamSearchRequest) GetFilterOk() (*StreamSearchFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *StreamSearchRequest) SetFilter(v StreamSearchFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *StreamSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *StreamSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *StreamSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *StreamSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *StreamSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *StreamSearchRequest) GetSort() []StreamSearchSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *StreamSearchRequest) GetSortOk() (*[]StreamSearchSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *StreamSearchRequest) SetSort(v []StreamSearchSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *StreamSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


