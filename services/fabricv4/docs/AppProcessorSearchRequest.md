# AppProcessorSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**AppProcessorFilters**](AppProcessorFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]AppProcessorSortCriteria**](AppProcessorSortCriteria.md) |  | [optional] 

## Methods

### NewAppProcessorSearchRequest

`func NewAppProcessorSearchRequest() *AppProcessorSearchRequest`

NewAppProcessorSearchRequest instantiates a new AppProcessorSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppProcessorSearchRequestWithDefaults

`func NewAppProcessorSearchRequestWithDefaults() *AppProcessorSearchRequest`

NewAppProcessorSearchRequestWithDefaults instantiates a new AppProcessorSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *AppProcessorSearchRequest) GetFilter() AppProcessorFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AppProcessorSearchRequest) GetFilterOk() (*AppProcessorFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AppProcessorSearchRequest) SetFilter(v AppProcessorFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AppProcessorSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *AppProcessorSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AppProcessorSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AppProcessorSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AppProcessorSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *AppProcessorSearchRequest) GetSort() []AppProcessorSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *AppProcessorSearchRequest) GetSortOk() (*[]AppProcessorSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *AppProcessorSearchRequest) SetSort(v []AppProcessorSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *AppProcessorSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


