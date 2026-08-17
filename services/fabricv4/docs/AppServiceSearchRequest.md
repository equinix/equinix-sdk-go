# AppServiceSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**AppServiceFilters**](AppServiceFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]AppServiceSortCriteria**](AppServiceSortCriteria.md) |  | [optional] 

## Methods

### NewAppServiceSearchRequest

`func NewAppServiceSearchRequest() *AppServiceSearchRequest`

NewAppServiceSearchRequest instantiates a new AppServiceSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceSearchRequestWithDefaults

`func NewAppServiceSearchRequestWithDefaults() *AppServiceSearchRequest`

NewAppServiceSearchRequestWithDefaults instantiates a new AppServiceSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *AppServiceSearchRequest) GetFilter() AppServiceFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AppServiceSearchRequest) GetFilterOk() (*AppServiceFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AppServiceSearchRequest) SetFilter(v AppServiceFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AppServiceSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *AppServiceSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AppServiceSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AppServiceSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AppServiceSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *AppServiceSearchRequest) GetSort() []AppServiceSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *AppServiceSearchRequest) GetSortOk() (*[]AppServiceSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *AppServiceSearchRequest) SetSort(v []AppServiceSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *AppServiceSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


