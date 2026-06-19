# InternetAccessSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Filter** | Pointer to [**SearchExpression**](SearchExpression.md) |  | [optional] 
**Sort** | Pointer to [**[]SearchSortItem**](SearchSortItem.md) |  | [optional] 

## Methods

### NewInternetAccessSearchRequest

`func NewInternetAccessSearchRequest() *InternetAccessSearchRequest`

NewInternetAccessSearchRequest instantiates a new InternetAccessSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessSearchRequestWithDefaults

`func NewInternetAccessSearchRequestWithDefaults() *InternetAccessSearchRequest`

NewInternetAccessSearchRequestWithDefaults instantiates a new InternetAccessSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *InternetAccessSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *InternetAccessSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *InternetAccessSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *InternetAccessSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetFilter

`func (o *InternetAccessSearchRequest) GetFilter() SearchExpression`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InternetAccessSearchRequest) GetFilterOk() (*SearchExpression, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InternetAccessSearchRequest) SetFilter(v SearchExpression)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *InternetAccessSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *InternetAccessSearchRequest) GetSort() []SearchSortItem`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *InternetAccessSearchRequest) GetSortOk() (*[]SearchSortItem, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *InternetAccessSearchRequest) SetSort(v []SearchSortItem)`

SetSort sets Sort field to given value.

### HasSort

`func (o *InternetAccessSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


