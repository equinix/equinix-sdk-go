# AppDomainSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**AppDomainFilters**](AppDomainFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]AppDomainSortCriteria**](AppDomainSortCriteria.md) |  | [optional] 

## Methods

### NewAppDomainSearchRequest

`func NewAppDomainSearchRequest() *AppDomainSearchRequest`

NewAppDomainSearchRequest instantiates a new AppDomainSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDomainSearchRequestWithDefaults

`func NewAppDomainSearchRequestWithDefaults() *AppDomainSearchRequest`

NewAppDomainSearchRequestWithDefaults instantiates a new AppDomainSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *AppDomainSearchRequest) GetFilter() AppDomainFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AppDomainSearchRequest) GetFilterOk() (*AppDomainFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AppDomainSearchRequest) SetFilter(v AppDomainFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AppDomainSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *AppDomainSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AppDomainSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AppDomainSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AppDomainSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *AppDomainSearchRequest) GetSort() []AppDomainSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *AppDomainSearchRequest) GetSortOk() (*[]AppDomainSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *AppDomainSearchRequest) SetSort(v []AppDomainSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *AppDomainSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


