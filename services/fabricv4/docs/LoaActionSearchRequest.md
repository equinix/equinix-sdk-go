# LoaActionSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**LoaActionSearchFilters**](LoaActionSearchFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]LoaActionSortCriteria**](LoaActionSortCriteria.md) |  | [optional] 

## Methods

### NewLoaActionSearchRequest

`func NewLoaActionSearchRequest() *LoaActionSearchRequest`

NewLoaActionSearchRequest instantiates a new LoaActionSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaActionSearchRequestWithDefaults

`func NewLoaActionSearchRequestWithDefaults() *LoaActionSearchRequest`

NewLoaActionSearchRequestWithDefaults instantiates a new LoaActionSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *LoaActionSearchRequest) GetFilter() LoaActionSearchFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LoaActionSearchRequest) GetFilterOk() (*LoaActionSearchFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LoaActionSearchRequest) SetFilter(v LoaActionSearchFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LoaActionSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *LoaActionSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *LoaActionSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *LoaActionSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *LoaActionSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *LoaActionSearchRequest) GetSort() []LoaActionSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LoaActionSearchRequest) GetSortOk() (*[]LoaActionSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LoaActionSearchRequest) SetSort(v []LoaActionSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LoaActionSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


