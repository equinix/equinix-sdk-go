# LoaSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**LoaSearchFilters**](LoaSearchFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]LoaSortCriteria**](LoaSortCriteria.md) |  | [optional] 

## Methods

### NewLoaSearchRequest

`func NewLoaSearchRequest() *LoaSearchRequest`

NewLoaSearchRequest instantiates a new LoaSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaSearchRequestWithDefaults

`func NewLoaSearchRequestWithDefaults() *LoaSearchRequest`

NewLoaSearchRequestWithDefaults instantiates a new LoaSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *LoaSearchRequest) GetFilter() LoaSearchFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LoaSearchRequest) GetFilterOk() (*LoaSearchFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LoaSearchRequest) SetFilter(v LoaSearchFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LoaSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *LoaSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *LoaSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *LoaSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *LoaSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *LoaSearchRequest) GetSort() []LoaSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LoaSearchRequest) GetSortOk() (*[]LoaSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LoaSearchRequest) SetSort(v []LoaSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LoaSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


