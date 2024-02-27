# ServiceProfileSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**ServiceProfileFilter**](ServiceProfileFilter.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]ServiceProfileSortCriteria**](ServiceProfileSortCriteria.md) |  | [optional] 

## Methods

### NewServiceProfileSearchRequest

`func NewServiceProfileSearchRequest() *ServiceProfileSearchRequest`

NewServiceProfileSearchRequest instantiates a new ServiceProfileSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileSearchRequestWithDefaults

`func NewServiceProfileSearchRequestWithDefaults() *ServiceProfileSearchRequest`

NewServiceProfileSearchRequestWithDefaults instantiates a new ServiceProfileSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ServiceProfileSearchRequest) GetFilter() ServiceProfileFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ServiceProfileSearchRequest) GetFilterOk() (*ServiceProfileFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ServiceProfileSearchRequest) SetFilter(v ServiceProfileFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ServiceProfileSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *ServiceProfileSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ServiceProfileSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ServiceProfileSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ServiceProfileSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *ServiceProfileSearchRequest) GetSort() []ServiceProfileSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ServiceProfileSearchRequest) GetSortOk() (*[]ServiceProfileSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ServiceProfileSearchRequest) SetSort(v []ServiceProfileSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ServiceProfileSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


